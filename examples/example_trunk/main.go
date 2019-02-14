// -*- coding: utf-8 -*-

// Copyright (C) 2019 Nippon Telegraph and Telephone Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"net"

	"github.com/beluganos/go-opennsl/examples/util"
	"github.com/beluganos/go-opennsl/opennsl"
	"github.com/beluganos/go-opennsl/sal"

	log "github.com/sirupsen/logrus"
)

func switchControlSet(unit int) error {
	if err := opennsl.SwitchL3EgressMode.Set(unit, opennsl.TRUE); err != nil {
		log.Errorf("%s", err)
		return err
	}

	hc, err := opennsl.SwitchHashControl.Get(unit)
	if err != nil {
		log.Errorf("%s", err)
		return err
	}

	hashControl := opennsl.HashControls(hc)
	hashControl = opennsl.NewHashControls(
		hashControl,
		opennsl.HASH_CONTROL_TRUNK_NUC_DST,
		opennsl.HASH_CONTROL_TRUNK_NUC_SRC,
		opennsl.HASH_CONTROL_TRUNK_UC_SRCPORT,
	)

	if err := opennsl.SwitchHashControl.Set(unit, int(hashControl)); err != nil {
		log.Errorf("%s", err)
		return err
	}

	hashControl = opennsl.NewHashControls(
		hashControl,
		opennsl.HASH_CONTROL_MULTIPATH_L4PORTS,
		opennsl.HASH_CONTROL_MULTIPATH_DIP,
	)

	if err := opennsl.SwitchHashControl.Set(unit, int(hashControl)); err != nil {
		log.Errorf("%s", err)
		return err
	}

	return nil
}

func newTrunk(unit int, ports ...opennsl.Port) (opennsl.Trunk, error) {
	trunk, err := opennsl.TrunkCreate(unit, opennsl.TRUNK_FLAG_NONE)
	if err != nil {
		return 0, err
	}

	for _, port := range ports {
		gport, err := port.GPortGet(unit)
		if err != nil {
			trunk.Destroy(unit)
			return 0, err
		}

		member := opennsl.NewTrunkMember()
		member.SetGPort(gport)

		if err := trunk.MemberAdd(unit, member); err != nil {
			trunk.Destroy(unit)
			return 0, err
		}
	}

	return trunk, nil
}

func main() {
	log.SetLevel(log.DebugLevel)

	unit := 0
	mymac, _ := net.ParseMAC("00:99:99:99:99:99")

	port1 := opennsl.Port(51)
	port2 := opennsl.Port(52)
	nh, _ := net.ParseMAC("00:11:11:11:11:11")

	if err := sal.DriverInit(); err != nil {
		log.Errorf("driver_init error. %s", err)
		return
	}
	defer sal.DriverExit()

	if err := util.PortDefaultConfig(unit); err != nil {
		log.Errorf("%s", err)
		return
	}

	if err := util.SwitchDefaultVlanConfig(unit); err != nil {
		log.Errorf("%s", err)
		return
	}

	if err := switchControlSet(unit); err != nil {
		log.Errorf("%s", err)
		return
	}

	trunk, err := newTrunk(unit, port1, port2)
	if err != nil {
		log.Errorf("%s", err)
		return
	}

	log.Infof("Trunk: %d", trunk)

	l2Flags := opennsl.L2_L3LOOKUP
	if _, err := util.AddL2Addr(unit, port1, opennsl.VLAN_ID_DEFAULT, mymac, l2Flags); err != nil {
		log.Errorf("%s", err)
		return
	}
	if _, err := util.AddL2Addr(unit, port2, opennsl.VLAN_ID_DEFAULT, mymac, l2Flags); err != nil {
		log.Errorf("%s", err)
		return
	}

	l3iface, err := util.NewL3IfaceObj(unit, 0, opennsl.VLAN_ID_DEFAULT, mymac)
	if err != nil {
		log.Errorf("%s", err)
		return
	}

	l3egr, err := util.NewL3EgressTrunk(unit, trunk, opennsl.VLAN_ID_DEFAULT, l3iface.IfaceID(), nh)
	if err != nil {
		log.Errorf("%s", err)
		return
	}

	_, route, _ := net.ParseCIDR("100.0.0.0/24")
	if _, err := util.AddRoute(unit, route, l3egr, opennsl.VRF_NONE); err != nil {
		log.Errorf("%s", err)
		return
	}

	done := make(chan struct{})
	go util.WatchSignal(done)
	<-done
}
