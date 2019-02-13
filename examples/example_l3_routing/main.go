// -*- coding: utf-8 -*-

// Copyright (C) 2018 Nippon Telegraph and Telephone Corporation.
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
	inPort := opennsl.Port(1)
	inVid := opennsl.Vlan(10)
	outPort := opennsl.Port(53)
	outVid := opennsl.Vlan(11)
	mac, _ := net.ParseMAC("00:11:22:33:44:55")
	nh, _ := net.ParseMAC("00:66:77:88:99:aa")

	if err := sal.DriverInit(); err != nil {
		log.Errorf("driver_init error. %s", err)
		return
	}
	defer sal.DriverExit()

	bootflag := sal.DriverBootFlagsGet()
	log.Infof("boot flag = %x", bootflag)

	if err := util.PortDefaultConfig(unit); err != nil {
		log.Errorf("%s", err)
		return
	}

	if err := opennsl.SwitchL3EgressMode.Set(unit, 1); err != nil {
		log.Errorf("%s", err)
		return
	}

	inIface, err := util.NewL3IfaceObj(unit, inPort, inVid, mac)
	if err != nil {
		log.Errorf("%s", err)
		return
	}
	log.Infof("Iface In:%d", inIface.IfaceID())

	outIface, err := util.NewL3IfaceObj(unit, outPort, outVid, mac)
	if err != nil {
		log.Errorf("%s", err)
		return
	}

	log.Infof("Iface Out:%d", outIface.IfaceID())

	l2Flags := opennsl.L2_L3LOOKUP

	if _, err := util.AddL2Addr(unit, inPort, 1, mac, l2Flags); err != nil {
		log.Errorf("%s", err)
		return
	}
	if _, err := util.AddL2Addr(unit, inPort, inVid, mac, l2Flags); err != nil {
		log.Errorf("%s", err)
		return
	}
	if _, err := util.AddL2Addr(unit, outPort, outVid, mac, l2Flags); err != nil {
		log.Errorf("%s", err)
		return
	}

	l3eg, err := util.NewL3Egress(unit, outPort, outVid, outIface.IfaceID(), nh)
	if err != nil {
		log.Errorf("%s", err)
		return
	}

	log.Infof("L3Egress: %d", l3eg)

	host4 := net.ParseIP("10.0.1.1")
	_, route4, _ := net.ParseCIDR("55.0.0.0/8")
	if _, err := util.AddHost(unit, host4, l3eg, opennsl.VRF_NONE); err != nil {
		log.Errorf("%s", err)
		return
	}

	if _, err := util.AddRoute(unit, route4, l3eg, opennsl.VRF_NONE); err != nil {
		log.Errorf("%s", err)
		return
	}

	host6 := net.ParseIP("2001:db8::10:0:1:1")
	_, route6, _ := net.ParseCIDR("2001:db8:55::/64")
	if _, err := util.AddHost6(unit, host6, l3eg, opennsl.VRF_NONE); err != nil {
		log.Errorf("%s", err)
		return
	}

	if _, err := util.AddRoute6(unit, route6, l3eg, opennsl.VRF_NONE); err != nil {
		log.Errorf("%s", err)
		return
	}

	util.TraverseL2Addrs(unit)
	util.TraverseL3Egresses(unit)
	util.TraverseL3Hosts(unit)
	util.TraverseL3Routes(unit)

	done := make(chan struct{})
	go util.WatchSignal(done)

	<-done
}
