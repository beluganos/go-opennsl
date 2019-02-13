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

func tunInitiator4(unit int, dst, src net.IP, iface *opennsl.L3Iface) error {

	tun := opennsl.NewTunnelInitiator(opennsl.TunnelTypeIPIP4encap)
	tun.SetTTL(64)
	tun.SetDstIP4(dst)
	tun.SetSrcIP4(src)
	tun.SetVID(opennsl.VLAN_ID_DEFAULT)
	tun.SetL3IfaceID(iface.IfaceID())
	if err := tun.Create(unit, iface); err != nil {
		log.Errorf("%s", err)
		return err
	}

	return nil
}

func tunInitiator6(unit int, dst, src net.IP, iface *opennsl.L3Iface) error {

	tun := opennsl.NewTunnelInitiator(opennsl.TunnelTypeIPIP6encap)
	tun.SetTTL(64)
	tun.SetDstIP6(dst)
	tun.SetSrcIP6(src)
	tun.SetVID(opennsl.VLAN_ID_DEFAULT)
	tun.SetL3IfaceID(iface.IfaceID())
	if err := tun.Create(unit, iface); err != nil {
		log.Errorf("%s", err)
		return err
	}

	return nil
}

func initDriver(unit int) error {
	if err := sal.DriverInit(); err != nil {
		log.Errorf("driver_init error. %s", err)
		return err
	}

	if err := util.PortDefaultConfig(unit); err != nil {
		log.Errorf("%s", err)
		return err
	}

	//if err := util.SwitchDefaultVlanConfig(unit); err != nil {
	//  log.Errorf("%s", err)
	//  return err
	//}

	if err := opennsl.SwitchL3EgressMode.Set(unit, 1); err != nil {
		log.Errorf("SwitchL3EgressMode.Set error. %s", err)
		return err
	}

	return nil
}

func main() {
	log.SetLevel(log.DebugLevel)

	unit := 0
	vid := opennsl.VLAN_ID_DEFAULT
	vid0 := opennsl.Vlan(11)
	vid1 := opennsl.Vlan(12)
	vid2 := opennsl.Vlan(13)
	vid3 := opennsl.Vlan(14)

	port0 := opennsl.Port(54)
	port1 := opennsl.Port(55)
	port2 := opennsl.Port(56)
	port3 := opennsl.Port(57)

	my_mac0, _ := net.ParseMAC("00:11:22:33:44:00")
	my_mac1, _ := net.ParseMAC("00:11:22:33:44:01")
	my_mac2, _ := net.ParseMAC("00:11:22:33:44:02")
	my_mac3, _ := net.ParseMAC("00:11:22:33:44:03")

	nh_mac0, _ := net.ParseMAC("00:66:77:88:99:00")
	nh_mac1, _ := net.ParseMAC("00:66:77:88:99:01")
	nh_mac2, _ := net.ParseMAC("00:66:77:88:99:02")
	nh_mac3, _ := net.ParseMAC("00:66:77:88:99:03")

	tun_local4 := net.ParseIP("100.0.1.1")
	tun_remot4 := net.ParseIP("100.0.1.2")
	tun_local6 := net.ParseIP("2010:2020::1")
	tun_remot6 := net.ParseIP("2010:2010::4")

	host4 := net.ParseIP("10.0.1.1")
	_, route4, _ := net.ParseCIDR("55.0.0.0/8")
	host6 := net.ParseIP("2010:2010::3")
	_, route6, _ := net.ParseCIDR("2010:2011::/64")

	log.Infof("DEVAULT_VLAN=%d", vid)

	//
	// Driver Initialize
	//
	if err := initDriver(unit); err != nil {
		log.Errorf("initDriver error. %s", err)
	}
	defer sal.DriverExit()

	//
	// L3 Interface
	//
	iface0, err := util.NewL3IfaceObjUntagged(unit, port0, vid0, my_mac0)
	if err != nil {
		log.Errorf("%s", err)
		return
	}

	iface1, err := util.NewL3IfaceObjUntagged(unit, port1, vid1, my_mac1)
	if err != nil {
		log.Errorf("%s", err)
		return
	}

	iface2, err := util.NewL3IfaceObjUntagged(unit, port2, vid2, my_mac2)
	if err != nil {
		log.Errorf("%s", err)
		return
	}

	iface3, err := util.NewL3IfaceObjUntagged(unit, port3, vid3, my_mac3)
	if err != nil {
		log.Errorf("%s", err)
		return
	}

	//
	// L3 Egress
	//
	if _, err := util.NewL3Egress(unit, port0, vid, iface0.IfaceID(), nh_mac0); err != nil {
		log.Errorf("%s", err)
		return
	}

	if _, err := util.NewL3Egress(unit, port1, vid, iface1.IfaceID(), nh_mac1); err != nil {
		log.Errorf("%s", err)
		return
	}

	l3eg2, err := util.NewL3Egress(unit, port2, vid, iface2.IfaceID(), nh_mac2)
	if err != nil {
		log.Errorf("%s", err)
		return
	}

	l3eg3, err := util.NewL3Egress(unit, port3, vid, iface3.IfaceID(), nh_mac3)
	if err != nil {
		log.Errorf("%s", err)
		return
	}

	//
	// L2 Address
	//
	if _, err := util.AddL2Addr(unit, port0, vid, my_mac0, opennsl.L2_L3LOOKUP); err != nil {
		log.Errorf("%s", err)
		return
	}
	if _, err := util.AddL2Addr(unit, port1, vid, my_mac1, opennsl.L2_L3LOOKUP); err != nil {
		log.Errorf("%s", err)
		return
	}
	if _, err := util.AddL2Addr(unit, port2, vid, my_mac2, opennsl.L2_L3LOOKUP); err != nil {
		log.Errorf("%s", err)
		return
	}
	if _, err := util.AddL2Addr(unit, port3, vid, my_mac3, opennsl.L2_L3LOOKUP); err != nil {
		log.Errorf("%s", err)
		return
	}

	//
	// Tunnel Initiator
	//
	if err := tunInitiator4(unit, tun_remot4, tun_local4, iface2); err != nil {
		log.Errorf("%s", err)
		return
	}

	if err := tunInitiator6(unit, tun_remot6, tun_local6, iface3); err != nil {
		log.Errorf("%s", err)
		return
	}

	//
	// Host
	//
	if _, err := util.AddHost(unit, host4, l3eg2, opennsl.VRF_NONE); err != nil {
		log.Errorf("%s", err)
		return
	}

	if _, err := util.AddHost6(unit, host6, l3eg2, opennsl.VRF_NONE); err != nil {
		log.Errorf("%s", err)
		return
	}

	//
	// Route
	//
	if _, err := util.AddRoute(unit, route4, l3eg3, opennsl.VRF_NONE); err != nil {
		log.Errorf("%s", err)
		return
	}

	if _, err := util.AddRoute6(unit, route6, l3eg3, opennsl.VRF_NONE); err != nil {
		log.Errorf("%s", err)
		return
	}

	//
	// Wait signal
	//
	done := make(chan struct{})
	go util.WatchSignal(done)

	<-done
}
