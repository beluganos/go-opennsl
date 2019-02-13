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

func tunInitSetup(unit int, mac net.HardwareAddr, vlan opennsl.Vlan, iface opennsl.L3IfaceID, label opennsl.MplsLabel) error {
	if err := opennsl.L2TunnelAdd(unit, mac, vlan); err != nil {
		log.Errorf("opennsl.L2TunnelAdd error %s", err)
		return err
	}

	tunLabel := opennsl.NewMplsEgressLabel()
	tunLabel.SetFlags(opennsl.NewMplsEgressLabelFlags(
		opennsl.MPLS_EGRESS_LABEL_TTL_SET,
	))
	tunLabel.SetLabel(label)
	tunLabel.SetTTL(64)

	return iface.MplsTunInitiatorCreate(unit, tunLabel)
}

func tunPopSetup(unit int, label opennsl.MplsLabel, l3eg opennsl.L3EgressID, vpn opennsl.Vpn) error {
	tunsw := opennsl.NewMplsTunnelSwitch()
	tunsw.SetLabel(label)
	tunsw.SetPort(opennsl.GPORT_INVALID)
	tunsw.SetAction(opennsl.MPLS_SWITCH_ACTION_POP)
	//tunsw.SetActionIfBos(opennsl.MPLS_SWITCH_ACTION_POP)
	tunsw.SetVpn(vpn)
	tunsw.SetEgress(l3eg)

	return tunsw.Create(unit)
}

func vpnSetup(unit int, vpn opennsl.Vpn) (opennsl.Vpn, error) {
	vpncfg := opennsl.NewMplsVpnConfig()
	vpncfg.SetFlags(opennsl.NewMplsVpnFlags(
		opennsl.MPLS_VPN_L3,
		opennsl.MPLS_VPN_WITH_ID,
	))
	vpncfg.SetVpn(vpn)

	if err := vpncfg.Create(unit); err != nil {
		return 0, err
	}

	return vpncfg.Vpn(), nil
}

func main() {
	log.SetLevel(log.DebugLevel)

	unit := 0
	port1 := opennsl.Port(52)
	vlan1 := opennsl.Vlan(10)
	vpn1 := opennsl.Vpn(7001)
	port2 := opennsl.Port(53)
	vlan2 := opennsl.Vlan(20)
	vpn2 := opennsl.Vpn(7002)

	myMAC, _ := net.ParseMAC("00:11:22:33:44:55")
	nhMAC1, _ := net.ParseMAC("00:00:00:00:00:11")
	nhMAC2, _ := net.ParseMAC("00:00:00:00:00:22")

	host1 := net.ParseIP("10.0.1.2")
	host2 := net.ParseIP("10.0.2.2")
	_, route1, _ := net.ParseCIDR("55.0.0.0/8")
	_, route2, _ := net.ParseCIDR("66.0.0.0/8")

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

	//
	// Create L3 Interfaces
	//
	iface1, err := util.NewL3IfaceObjUntagged(unit, port1, vlan1, myMAC)
	if err != nil {
		log.Errorf("%s", err)
		return
	}

	iface2, err := util.NewL3IfaceObjUntagged(unit, port2, vlan2, myMAC)
	if err != nil {
		log.Errorf("%s", err)
		return
	}

	log.Infof("Iface1: id=%d port=%d vlan=%d mac=%s", iface1.IfaceID(), port1, vlan1, myMAC)
	log.Infof("Iface2: id=%d port=%d vlan=%d mac=%s", iface2.IfaceID(), port2, vlan2, myMAC)

	//
	// Add L2 Table
	//
	l2Flags := opennsl.L2_L3LOOKUP
	if _, err := util.AddL2Addr(unit, port1, vlan1, myMAC, l2Flags); err != nil {
		log.Errorf("%s", err)
		return
	}
	if _, err := util.AddL2Addr(unit, port2, vlan2, myMAC, l2Flags); err != nil {
		log.Errorf("%s", err)
		return
	}

	//
	// Create L3 Egress
	//
	l3eg1, err := util.NewL3Egress(unit, port1, vlan1, iface1.IfaceID(), nhMAC1)
	if err != nil {
		log.Errorf("%s", err)
		return
	}
	l3eg2, err := util.NewL3Egress(unit, port2, vlan2, iface2.IfaceID(), nhMAC2)
	if err != nil {
		log.Errorf("%s", err)
		return
	}

	log.Infof("L3Egress1: id=%d ifaceID=%d", l3eg1, iface1)
	log.Infof("L3Egress2: id=%d ifaceID=%d", l3eg2, iface2)

	//
	// Add route to host.
	//
	if _, err := util.AddHost(unit, host1, l3eg1, opennsl.VRF_NONE); err != nil {
		log.Errorf("hostAdd error. %s", err)
		return
	}

	log.Infof("Host: host=%s egress=%d", host1, l3eg1)

	if _, err := util.AddHost(unit, host2, l3eg2, opennsl.VRF_NONE); err != nil {
		log.Errorf("hostAdd error. %s", err)
		return
	}

	log.Infof("Host: host=%s egress=%d", host2, l3eg2)

	//
	// Add route to subnet
	//
	if _, err := util.AddRoute(unit, route1, l3eg1, opennsl.VRF_NONE); err != nil {
		log.Errorf("routeAdd error. %s", err)
		return
	}

	log.Infof("routeAdd: route=%s egress=%d", route1, l3eg1)

	if _, err := util.AddRoute(unit, route2, l3eg2, opennsl.VRF_NONE); err != nil {
		log.Errorf("routeAdd error. %s", err)
		return
	}

	log.Infof("routeAdd: route=%s egress=%d", route2, l3eg2)

	//
	// MPLS tunnel
	//
	if err := tunInitSetup(unit, myMAC, vlan2, iface2.IfaceID(), 400); err != nil {
		log.Errorf("tunInitSetup error. %s", err)
		return
	}

	log.Infof("tunInitSetup: mac=%s iface=%d", myMAC, iface2)

	//
	// MPLS VPN
	//
	if _, err := vpnSetup(unit, vpn1); err != nil {
		log.Errorf("vpnSetup error. %s", err)
		return
	}
	if _, err := vpnSetup(unit, vpn2); err != nil {
		log.Errorf("vpnSetup error. %s", err)
		return
	}

	//
	// MPLS POP
	//
	if err := tunPopSetup(unit, 100, l3eg2, vpn2); err != nil {
		log.Errorf("tunPopSetup error. %s", err)
		return
	}

	//
	// Wait CTRL-c.
	//
	done := make(chan struct{})
	go util.WatchSignal(done)
	<-done
}
