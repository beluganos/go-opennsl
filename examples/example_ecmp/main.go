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

func main() {
	log.SetLevel(log.DebugLevel)

	unit := 0

	port1 := opennsl.Port(1)
	port2 := opennsl.Port(2)
	mac1, _ := net.ParseMAC("00:00:00:00:00:11")
	mac2, _ := net.ParseMAC("00:00:00:00:00:22")
	nh1, _ := net.ParseMAC("00:00:00:00:11:11")
	nh2, _ := net.ParseMAC("00:00:00:00:22:22")

	// Driver initialize.
	if err := sal.DriverInit(); err != nil {
		log.Errorf("driver_init error. %s", err)
		return
	}
	defer sal.DriverExit()

	// Port configuration
	if err := util.PortDefaultConfig(unit); err != nil {
		log.Errorf("%s", err)
		return
	}

	if err := util.SwitchDefaultVlanConfig(unit); err != nil {
		log.Errorf("%s", err)
		return
	}

	// Switch configuration
	if err := opennsl.SwitchL3EgressMode.Set(unit, opennsl.TRUE); err != nil {
		log.Errorf("%s", err)
		return
	}

	if err := opennsl.SwitchEcmpMacroFlowHashEnable.Set(unit, opennsl.TRUE); err != nil {
		log.Errorf("%s", err)
		return
	}

	// L2 Address
	l2Flags := opennsl.L2_L3LOOKUP
	if _, err := util.AddL2Addr(unit, port1, opennsl.VLAN_ID_DEFAULT, mac1, l2Flags); err != nil {
		log.Errorf("%s", err)
		return
	}
	if _, err := util.AddL2Addr(unit, port2, opennsl.VLAN_ID_DEFAULT, mac2, l2Flags); err != nil {
		log.Errorf("%s", err)
		return
	}

	// L3 Interface
	l3intf1, err := util.NewL3IfaceObj(unit, port1, opennsl.VLAN_ID_DEFAULT, mac1)
	if err != nil {
		log.Errorf("%s", err)
		return
	}
	log.Infof("L3intf1:%d", l3intf1.IfaceID())

	l3intf2, err := util.NewL3IfaceObj(unit, port2, opennsl.VLAN_ID_DEFAULT, mac2)
	if err != nil {
		log.Errorf("%s", err)
		return
	}
	log.Infof("L3intf2:%d", l3intf2.IfaceID())

	// L3 Egress
	l3egr1, err := util.NewL3Egress(unit, port1, opennsl.VLAN_ID_DEFAULT, l3intf1.IfaceID(), nh1)
	if err != nil {
		log.Errorf("%s", err)
		return
	}
	l3egr2, err := util.NewL3Egress(unit, port2, opennsl.VLAN_ID_DEFAULT, l3intf2.IfaceID(), nh2)
	if err != nil {
		log.Errorf("%s", err)
		return
	}

	// L3 ECMP
	l3ecmp := opennsl.NewL3EgressEcmp()
	l3ecmp.SetDynamicMode(opennsl.L3_ECMP_DYNAMIC_MODE_NORMAL)
	if err := l3ecmp.Create(unit, []opennsl.L3EgressID{l3egr1, l3egr2}); err != nil {
		log.Errorf("%s", err)
		return
	}

	l3ecmpID := l3ecmp.EgressEcmp()
	log.Infof("ECMP: %d", l3ecmpID)

	// L3 Routing
	_, route, _ := net.ParseCIDR("100.0.0.0/24")
	if _, err := util.AddRoute(unit, route, l3ecmpID.L3EgressID(), opennsl.VRF_NONE); err != nil {
		log.Errorf("%s", err)
		return
	}

	done := make(chan struct{})
	go util.WatchSignal(done)

	<-done
}
