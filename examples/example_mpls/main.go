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
	"os"
	"os/signal"

	"github.com/beluganos/go-opennsl/examples/util"
	"github.com/beluganos/go-opennsl/opennsl"
	"github.com/beluganos/go-opennsl/sal"

	log "github.com/sirupsen/logrus"
)

func vlanSetup(unit int, vid opennsl.Vlan) error {
	_, err := vid.Create(unit)
	return err
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

func portSetup(unit int, port opennsl.Port, vid opennsl.Vlan) error {
	if err := opennsl.PortControlDoNotCheckVlan.Set(unit, port, 1); err != nil {
		return err
	}

	pbmp := opennsl.NewPBmp()
	ubmp := opennsl.NewPBmp()

	pbmp.Add(port)

	return vid.PortAdd(unit, pbmp, ubmp)
}

func l3ifaceSetup(unit int, vid opennsl.Vlan, srcMac net.HardwareAddr, ifaceID opennsl.L3IfaceID) (opennsl.L3IfaceID, error) {
	iface := opennsl.NewL3Iface()
	iface.SetFlags(opennsl.NewL3Flags(
		opennsl.L3_WITH_ID,
	))
	iface.SetIfaceID(ifaceID)
	iface.SetVID(vid)
	iface.SetVRF(0)
	iface.SetMAC(srcMac)
	if err := iface.Create(unit); err != nil {
		return opennsl.L3_IFACE, err
	}

	return iface.IfaceID(), nil
}

func l3egSetup(unit int, port opennsl.Port, vid opennsl.Vlan, dstMac net.HardwareAddr, iface opennsl.L3IfaceID) (opennsl.L3EgressID, error) {
	l3eg := opennsl.NewL3Egress()
	l3eg.SetIfaceID(iface)
	l3eg.SetMAC(dstMac)
	l3eg.SetVID(vid)
	l3eg.SetPort(port)
	return l3eg.Create(unit, 0, 0)
}

func tunPHPSetup(unit int, l3eg opennsl.L3EgressID) error {
	tunsw := opennsl.NewMplsTunnelSwitch()
	tunsw.SetLabel(100)
	tunsw.SetPort(opennsl.GPORT_INVALID)
	tunsw.SetAction(opennsl.MPLS_SWITCH_ACTION_PHP)
	tunsw.SetEgress(l3eg)
	tunsw.EgressLabel().SetFlags(opennsl.NewMplsEgressLabelFlags(
		opennsl.MPLS_EGRESS_LABEL_TTL_SET,
	))

	return tunsw.Create(unit)
}

func tunPopSetup(unit int, vpn opennsl.Vpn) error {
	tunsw := opennsl.NewMplsTunnelSwitch()
	tunsw.SetLabel(101)
	tunsw.SetPort(opennsl.GPORT_INVALID)
	tunsw.SetAction(opennsl.MPLS_SWITCH_ACTION_POP)
	tunsw.SetVpn(vpn)
	tunsw.EgressLabel().SetFlags(opennsl.NewMplsEgressLabelFlags(
		opennsl.MPLS_EGRESS_LABEL_TTL_SET,
	))

	return tunsw.Create(unit)
}

func routeSetup(unit int, l3eg opennsl.L3EgressID, vpn opennsl.Vpn) error {
	_, route, _ := net.ParseCIDR("55.0.0.0/24")
	if _, err := util.AddRoute(unit, route, l3eg, opennsl.VRF_NONE); err != nil {
		return err
	}

	return nil

}

func tunPushSetup(unit int, l3eg opennsl.L3EgressID, vpn opennsl.Vpn) error {
	tunsw := opennsl.NewMplsTunnelSwitch()
	tunsw.SetLabel(200)
	tunsw.SetPort(opennsl.GPORT_INVALID)
	tunsw.SetAction(opennsl.MPLS_SWITCH_ACTION_SWAP)
	tunsw.SetVpn(vpn)
	tunsw.SetEgress(l3eg)
	tunsw.EgressLabel().SetFlags(opennsl.NewMplsEgressLabelFlags(
		opennsl.MPLS_EGRESS_LABEL_TTL_SET,
	))
	tunsw.EgressLabel().SetLabel(200)
	tunsw.EgressLabel().SetExp(0)
	tunsw.EgressLabel().SetTTL(34)
	tunsw.EgressLabel().SetPktPri(0)
	tunsw.EgressLabel().SetPktCfi(0)

	return tunsw.Create(unit)
}

func tunSwapSetup(unit int, l3eg opennsl.L3EgressID, vpn opennsl.Vpn) error {
	tunsw := opennsl.NewMplsTunnelSwitch()
	tunsw.SetLabel(300)
	tunsw.SetPort(opennsl.GPORT_INVALID)
	tunsw.SetAction(opennsl.MPLS_SWITCH_ACTION_SWAP)
	tunsw.SetVpn(vpn)
	tunsw.SetEgress(l3eg)
	tunsw.EgressLabel().SetFlags(opennsl.NewMplsEgressLabelFlags(
		opennsl.MPLS_EGRESS_LABEL_TTL_SET,
	))
	tunsw.EgressLabel().SetLabel(350)
	tunsw.EgressLabel().SetExp(0)
	tunsw.EgressLabel().SetTTL(35)
	tunsw.EgressLabel().SetPktPri(0)
	tunsw.EgressLabel().SetPktCfi(0)

	return tunsw.Create(unit)
}

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

func mplsSetup(unit int, inPort, outPort opennsl.Port) error {

	vpn := opennsl.Vpn(1 << 12)
	inVid := opennsl.Vlan(10)
	outVid := opennsl.Vlan(20)
	my_mac, _ := net.ParseMAC("00:00:00:00:11:11")
	srcmac, _ := net.ParseMAC("00:11:11:11:11:0e")
	dstmac, _ := net.ParseMAC("00:11:11:11:11:0f")

	if err := opennsl.SwitchL3EgressMode.Set(unit, 1); err != nil {
		return err
	}

	vpnId, err := vpnSetup(unit, vpn)
	if err != nil {
		log.Errorf("vpnSetup error. %d %s", vpn, err)
		return err
	}

	if err := vlanSetup(unit, inVid); err != nil {
		log.Errorf("vlanSetup error. %d %s", inVid, err)
		return err
	}

	if err := portSetup(unit, inPort, inVid); err != nil {
		log.Errorf("portSetup error. %d %d %s", inPort, inVid, err)
		return err
	}

	if err := vlanSetup(unit, outVid); err != nil {
		log.Errorf("vpnSetup error. %d %s", outVid, err)
		return err
	}

	if err := portSetup(unit, outPort, outVid); err != nil {
		log.Errorf("portSetup error. %d %d %s", outPort, outVid, err)
		return err
	}

	ifaceId, err := l3ifaceSetup(unit, outVid, srcmac, 31)
	if err != nil {
		log.Errorf("l3ifaceSetup error. %d %s %s", outVid, srcmac, err)
		return err
	}

	log.Debugf("L3Iface %d", ifaceId)

	l3egId, err := l3egSetup(unit, outPort, outVid, dstmac, ifaceId)
	if err != nil {
		log.Errorf("l3legSetup error. port=%d vlan=%d iface=%d %s", outPort, outVid, ifaceId, err)
		return err
	}

	log.Debugf("L3Egress %d", l3egId)

	ifaceId2, err := l3ifaceSetup(unit, outVid, srcmac, 32)
	if err != nil {
		log.Errorf("l3ifaceSetup error. %d %s %s", outVid, srcmac, err)
		return err
	}

	log.Debugf("L3Iface(2) %d", ifaceId2)

	l3egId2, err := l3egSetup(unit, outPort, outVid, dstmac, ifaceId2)
	if err != nil {
		log.Errorf("l3legSetup error. port=%d vlan=%d iface=%d %s", outPort, outVid, ifaceId2, err)
		return err
	}

	log.Debugf("L3Egress(2) %d", l3egId2)

	if err := tunInitSetup(unit, my_mac, inVid, ifaceId, 400); err != nil {
		log.Errorf("tunInitSetup error. %s", err)
		return err
	}

	if err := tunSwapSetup(unit, l3egId, vpnId); err != nil {
		log.Errorf("tunSwapSetup error %s", err)
		return err
	}

	log.Debugf("tunSwapSetup ok.")

	if err := tunPushSetup(unit, l3egId, vpnId); err != nil {
		log.Errorf("tunPushSetup error %s", err)
		return err
	}

	log.Debugf("tunPushSetup ok")

	if err := routeSetup(unit, l3egId2, vpnId); err != nil {
		log.Errorf("routeSetup error. %s", err)
		return err
	}

	if err := tunPopSetup(unit, vpnId); err != nil {
		log.Errorf("tunPopSetup error %s", err)
		return err
	}

	log.Debugf("tunPopSetup ok")

	if err := tunPHPSetup(unit, l3egId2); err != nil {
		log.Errorf("tunPHPSetup error %s", err)
		return err
	}

	log.Debugf("tunPHPSetup ok")

	return nil
}

func watchSignal(done chan struct{}) {

	ch := make(chan os.Signal, 1)
	defer close(ch)

	signal.Notify(ch, os.Interrupt)
	<-ch

	log.Infof("Interrupt signal.")

	close(done)
}

func main() {
	log.SetLevel(log.DebugLevel)

	unit := int(0)
	inPort := opennsl.Port(50)
	outPort := opennsl.Port(51)

	if err := sal.DriverInit(); err != nil {
		log.Errorf("driver_init error. %s", err)
		return
	}
	defer sal.DriverExit()

	if err := util.PortDefaultConfig(unit); err != nil {
		log.Errorf("PortDefaultConfig. %s", err)
		return
	}

	if err := util.SwitchDefaultVlanConfig(unit); err != nil {
		log.Errorf("SwitchDefaultVlanConfig. %s", err)
		return
	}

	if err := opennsl.SwitchL3EgressMode.Set(unit, 1); err != nil {
		log.Errorf("%s", err)
		return
	}

	if err := mplsSetup(unit, inPort, outPort); err != nil {
		log.Errorf("mplsSetup error. %s", err)
		return
	}

	done := make(chan struct{})
	go watchSignal(done)

	<-done
}
