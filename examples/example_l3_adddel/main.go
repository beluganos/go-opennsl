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

func makeIfaces(unit int) []opennsl.L3IfaceID {
	var l3iface *opennsl.L3Iface
	l3ifaces := []opennsl.L3IfaceID{}
	l3iface, _ = util.NewL3IfaceObj(unit, 50, 11, util.ParseMAC("00:00:00:00:00:11"))
	l3ifaces = append(l3ifaces, l3iface.IfaceID())
	l3iface, _ = util.NewL3IfaceObj(unit, 50, 12, util.ParseMAC("00:00:00:00:00:11"))
	l3ifaces = append(l3ifaces, l3iface.IfaceID())
	l3iface, _ = util.NewL3IfaceObj(unit, 51, 11, util.ParseMAC("00:00:00:00:00:22"))
	l3ifaces = append(l3ifaces, l3iface.IfaceID())
	l3iface, _ = util.NewL3IfaceObj(unit, 51, 12, util.ParseMAC("00:00:00:00:00:22"))
	l3ifaces = append(l3ifaces, l3iface.IfaceID())

	return l3ifaces
}

func testL2Addr(unit int) {
	log.Infof("testL2Addr START")

	flags := opennsl.L2_L3LOOKUP
	var l2addr *opennsl.L2Addr
	l2addrs := []*opennsl.L2Addr{}
	l2addr, _ = util.AddL2Addr(unit, 50, 11, util.ParseMAC("00:00:00:00:00:11"), flags)
	l2addrs = append(l2addrs, l2addr)
	l2addr, _ = util.AddL2Addr(unit, 50, 12, util.ParseMAC("00:00:00:00:00:11"), flags)
	l2addrs = append(l2addrs, l2addr)
	l2addr, _ = util.AddL2Addr(unit, 51, 11, util.ParseMAC("00:00:00:00:00:22"), flags)
	l2addrs = append(l2addrs, l2addr)
	l2addr, _ = util.AddL2Addr(unit, 51, 12, util.ParseMAC("00:00:00:00:00:22"), flags)
	l2addrs = append(l2addrs, l2addr)

	for _, l2addr := range l2addrs {
		util.TraverseL2Addrs(unit)
		if err := l2addr.Delete(unit); err != nil {
			log.Errorf("L2Addr.Delete error. %s", err)
		}
	}
	util.TraverseL2Addrs(unit)

	log.Infof("testL2Addr END")
}

func testL3Egress(unit int, ifaces []opennsl.L3IfaceID) {
	log.Infof("testL3Egress START")

	var l3egrID opennsl.L3EgressID
	l3egrs := []opennsl.L3EgressID{}
	l3egrID, _ = util.NewL3Egress(unit, 50, 11, ifaces[0], util.ParseMAC("00:00:01:00:00:11"))
	l3egrs = append(l3egrs, l3egrID)
	l3egrID, _ = util.NewL3Egress(unit, 50, 12, ifaces[1], util.ParseMAC("00:00:01:00:00:11"))
	l3egrs = append(l3egrs, l3egrID)
	l3egrID, _ = util.NewL3Egress(unit, 51, 11, ifaces[2], util.ParseMAC("00:00:01:00:00:22"))
	l3egrs = append(l3egrs, l3egrID)
	l3egrID, _ = util.NewL3Egress(unit, 51, 12, ifaces[3], util.ParseMAC("00:00:01:00:00:22"))
	l3egrs = append(l3egrs, l3egrID)

	for _, l3egrID := range l3egrs {
		util.TraverseL3Egresses(unit)
		if err := l3egrID.Destroy(unit); err != nil {
			log.Errorf("L3EgressID.Destroy error. %s", err)
		}
	}
	util.TraverseL3Egresses(unit)

	log.Infof("testL3Egress END")
}

func testL3Host(unit int, ifaces []opennsl.L3IfaceID) {
	log.Infof("testL3Host START")

	l3egrID1, _ := util.NewL3Egress(unit, 50, 11, ifaces[0], util.ParseMAC("00:00:01:00:00:11"))
	l3egrID2, _ := util.NewL3Egress(unit, 51, 12, ifaces[1], util.ParseMAC("00:00:01:00:00:22"))
	defer l3egrID1.Destroy(unit)
	defer l3egrID2.Destroy(unit)

	var l3host *opennsl.L3Host
	l3hosts := []*opennsl.L3Host{}
	l3host, _ = util.AddHost(unit, net.ParseIP("10.1.1.1"), l3egrID1, 101)
	l3hosts = append(l3hosts, l3host)
	l3host, _ = util.AddHost(unit, net.ParseIP("10.1.1.1"), l3egrID2, 102)
	l3hosts = append(l3hosts, l3host)
	l3host, _ = util.AddHost(unit, net.ParseIP("10.1.1.2"), l3egrID1, 101)
	l3hosts = append(l3hosts, l3host)
	l3host, _ = util.AddHost(unit, net.ParseIP("10.1.1.2"), l3egrID2, 102)
	l3hosts = append(l3hosts, l3host)

	for _, l3host := range l3hosts {
		util.TraverseL3Hosts(unit)
		if err := l3host.Delete(unit); err != nil {
			log.Errorf("L3Host.Delete error. %s", err)
		}
	}
	util.TraverseL3Hosts(unit)

	log.Infof("testL3Host END")
}

func testL3Route(unit int, ifaces []opennsl.L3IfaceID) {
	log.Infof("testL3Route START")

	l3egrID1, _ := util.NewL3Egress(unit, 50, 11, ifaces[0], util.ParseMAC("00:00:01:00:00:11"))
	l3egrID2, _ := util.NewL3Egress(unit, 51, 12, ifaces[1], util.ParseMAC("00:00:01:00:00:22"))
	defer l3egrID1.Destroy(unit)
	defer l3egrID2.Destroy(unit)

	var l3route *opennsl.L3Route
	l3routes := []*opennsl.L3Route{}
	l3route, _ = util.AddRoute(unit, util.ParseIPNet("10.1.1.0/24"), l3egrID1, 101)
	l3routes = append(l3routes, l3route)
	l3route, _ = util.AddRoute(unit, util.ParseIPNet("10.1.1.0/24"), l3egrID2, 102)
	l3routes = append(l3routes, l3route)
	l3route, _ = util.AddRoute(unit, util.ParseIPNet("10.1.2.0/24"), l3egrID1, 101)
	l3routes = append(l3routes, l3route)
	l3route, _ = util.AddRoute(unit, util.ParseIPNet("10.1.2.0/24"), l3egrID2, 102)
	l3routes = append(l3routes, l3route)

	for _, l3route := range l3routes {
		util.TraverseL3Routes(unit)
		if err := l3route.Delete(unit); err != nil {
			log.Errorf("L3route.Delete error. %s", err)
		}
	}
	util.TraverseL3Routes(unit)

	log.Infof("testL3Route END")
}

func testMain(unit int) {
	testL2Addr(unit)

	ifaces := makeIfaces(unit)
	testL3Egress(unit, ifaces)
	testL3Host(unit, ifaces)
	testL3Route(unit, ifaces)
}

func main() {
	log.SetLevel(log.DebugLevel)

	unit := 0

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

	testMain(unit)

	done := make(chan struct{})
	go util.WatchSignal(done)

	<-done
}
