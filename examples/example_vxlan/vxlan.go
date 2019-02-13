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

	"github.com/beluganos/go-opennsl/opennsl"

	log "github.com/sirupsen/logrus"
)

func do_vxlan_global_setting(unit int) error {
	var err error

	err = opennsl.SwitchControlsSet(
		unit,
		opennsl.SwitchL3EgressMode.Arg(opennsl.TRUE),
		opennsl.SwitchVxlanUdpDestPortSet.Arg(opennsl.VXLAN_L4PORT),
		opennsl.SwitchVxlanEntropyEnable.Arg(opennsl.TRUE),
		opennsl.SwitchVxlanTunnelMissToCpu.Arg(opennsl.TRUE),
		opennsl.SwitchVxlanVnIdMissToCpu.Arg(opennsl.TRUE),
	)
	if err != nil {
		return err
	}

	err = opennsl.VlanControlsSet(
		unit,
		opennsl.VlanTranslate.Arg(opennsl.TRUE),
	)
	if err != nil {
		return err
	}

	return nil
}

func do_vxlan_access_port_settings(unit int, port opennsl.Port) error {
	return port.PortControlsSet(
		unit,
		opennsl.PortControlVxlanEnable.Arg(opennsl.FALSE),
		opennsl.PortControlVxlanTunnelbasedVnId.Arg(opennsl.FALSE),
	)
}

func do_vxlan_net_port_settings(unit int, port opennsl.Port) error {
	log.Debugf("do_vxlan_net_port_settings(%d)", port)
	return port.PortControlsSet(
		unit,
		opennsl.PortControlVxlanEnable.Arg(opennsl.TRUE),
		opennsl.PortControlVxlanTunnelbasedVnId.Arg(opennsl.FALSE),
		opennsl.PortControlVxlanDefaultTunnelEnable.Arg(opennsl.FALSE),
	)
}

func create_vxlan_vpn(unit int, vpn opennsl.Vpn, vnid opennsl.VNID, mcast opennsl.Multicast) error {
	log.Debugf("create_vxlan_vpn(%d, %d, %d)", vpn, vnid, mcast)
	cfg := opennsl.NewVxlanVpnConfig()
	cfg.SetFlags(opennsl.NewVxlanVpnFlags(
		opennsl.VXLAN_VPN_ELAN,
		opennsl.VXLAN_VPN_WITH_ID,
		opennsl.VXLAN_VPN_WITH_VPNID,
	))
	cfg.SetVpn(vpn)
	cfg.SetVNID(vnid)
	cfg.SetBroadcastGroup(mcast)
	cfg.SetUnknownMuticastGroup(mcast)
	cfg.SetUnknownUnicastGroup(mcast)

	return cfg.Create(unit)
}

func create_vxlan_acc_vp(unit int, vpn opennsl.Vpn, matchPort opennsl.GPort, criteria opennsl.VxlanPortMatch, l3egr opennsl.L3EgressID, vid opennsl.Vlan, flags ...opennsl.VxlanPortFlags) (opennsl.GPort, error) {
	vport := opennsl.NewVxlanPort()
	vport.SetFlags(opennsl.NewVxlanPortFlags(flags...) | opennsl.VXLAN_PORT_SERVICE_TAGGED)
	vport.SetCriteria(criteria)
	vport.SetMatchPort(matchPort)
	vport.SetMatchVlan(vid)
	vport.SetEgress(l3egr)

	if err := vport.Add(unit, vpn); err != nil {
		return 0, err
	}

	return vport.VxlanPortID(), nil
}

func create_vxlan_net_vp(unit int, vpn opennsl.Vpn, matchPort opennsl.GPort, criteria opennsl.VxlanPortMatch, l3egr opennsl.L3EgressID, tunInit opennsl.TunnelID, tunTerm opennsl.TunnelID, flags ...opennsl.VxlanPortFlags) (opennsl.GPort, error) {
	flags = append(
		flags,
		opennsl.VXLAN_PORT_NETWORK,
		opennsl.VXLAN_PORT_EGRESS_TUNNEL,
		opennsl.VXLAN_PORT_SERVICE_TAGGED,
	)
	vport := opennsl.NewVxlanPort()
	vport.SetFlags(opennsl.NewVxlanPortFlags(flags...))
	vport.SetCriteria(criteria)
	vport.SetMatchPort(matchPort)
	vport.SetMatchTunnelID(tunTerm)
	vport.SetEgressTunnelID(tunInit)
	vport.SetEgress(l3egr)

	log.Debugf("create_vxlan_net_vp(%d, %d, %d, %d/%d, %d)", vpn, matchPort, l3egr, tunInit, tunTerm, criteria)
	if err := vport.Add(unit, vpn); err != nil {
		return 0, err
	}

	return vport.VxlanPortID(), nil
}

func add_to_l2_station(unit int, mac net.HardwareAddr, vid opennsl.Vlan) (opennsl.L2StationID, error) {
	l2st := opennsl.NewL2Station()
	l2st.SetFlags(opennsl.L2_STATION_IPV4)
	l2st.SetDstMAC(mac)
	l2st.SetDstMACMask(opennsl.NewL2AddrMaskExact())
	l2st.SetVID(vid)
	l2st.SetVIDMask(opennsl.VLAN_MASK_EXACT)

	stID, err := opennsl.L2_STATION.Add(unit, l2st)
	if err != nil {
		log.Errorf("L2Station.Add error. %s", err)
		return opennsl.L2_STATION, err
	}

	return stID, err
}

func add_to_l2_table(unit int, mac net.HardwareAddr, vid opennsl.Vlan, gport opennsl.GPort) error {
	port, err := gport.LocalPortGet(unit)
	if err != nil {
		log.Errorf("gport.LocalPortGet(%d) error. %s", gport, err)
		return err
	}

	l2addr := opennsl.NewL2Addr(mac, vid)
	l2addr.SetFlags(opennsl.L2_STATIC)
	l2addr.SetPort(port)

	if err := l2addr.Add(unit); err != nil {
		log.Errorf("L2Addr.Add error. %s", err)
		return err
	}

	log.Debugf("AddL2Addr port %d, vlan %d, mac '%s'", gport, vid, mac)
	return nil
}

func vlan_create_add_port(unit int, vid opennsl.Vlan, port opennsl.Port) error {
	if _, err := vid.Create(unit); err != nil {
		log.Errorf("Vlan.Create error. %s", err)
		return err
	}

	pbmp := opennsl.NewPBmp()
	upbmp := opennsl.NewPBmp()
	pbmp.Add(port)

	if err := vid.PortAdd(unit, pbmp, upbmp); err != nil {
		log.Errorf("Vlan.PortAdd error. %s", err)
		return err
	}

	return nil
}

func create_l3_interface(unit int, local_mac net.HardwareAddr, vid opennsl.Vlan) (opennsl.L3IfaceID, error) {
	l3if := opennsl.NewL3Iface()
	l3if.SetMAC(local_mac)
	l3if.SetVID(vid)
	if err := l3if.Create(unit); err != nil {
		log.Errorf("NewL3Iface error. %s", err)
		return opennsl.L3_IFACE, err
	}

	return l3if.IfaceID(), nil
}

func create_egr_obj(unit int, l3_if opennsl.L3IfaceID, mac net.HardwareAddr, port opennsl.Port, vid opennsl.Vlan, flags ...opennsl.L3Flags) (opennsl.L3EgressID, error) {
	flags2 := append(flags, opennsl.L3_VXLAN_ONLY)

	l3eg := opennsl.NewL3Egress()
	l3eg.SetFlags(opennsl.NewL3Flags(flags2...))
	l3eg.SetIfaceID(l3_if)
	l3eg.SetMAC(mac)
	l3eg.SetVID(vid)
	l3eg.SetPort(port)

	return opennsl.L3_EGRESS.Create(unit, opennsl.NewL3Flags(flags...), l3eg)
}

func tunnel_initiator_setup(unit int, localIp net.IP, remoteIp net.IP) (opennsl.TunnelID, error) {
	log.Debugf("tunnel_initiator_setup(%s, %s)", localIp, remoteIp)

	tun := opennsl.NewVxlanTunnelInitiator()
	tun.SetTTL(16)
	tun.SetSrcIP4(localIp)
	tun.SetDstIP4(remoteIp)
	tun.SetUdpSrcPort(0xffff)
	tun.SetUdpDstPort(opennsl.VXLAN_L4PORT)

	if err := tun.Create(unit); err != nil {
		return opennsl.TUNNEL, err
	}

	return tun.TunnelID(), nil
}

func tunnel_terminator_setup(unit int, remoteIp net.IP, localIp net.IP, vid opennsl.Vlan, tunInitID opennsl.TunnelID) (opennsl.TunnelID, error) {
	log.Debugf("tunnel_terminator_setup(%s, %s, %d, %d)", remoteIp, localIp, vid, tunInitID)
	tun := opennsl.NewVxlanTunnelTerminator()
	tun.SetFlags(opennsl.TUNNEL_TERM_TUNNEL_WITH_ID)
	tun.SetDstIP4(localIp)
	tun.SetSrcIP4(remoteIp)
	tun.SetTunnelID(tunInitID)
	if vid.Valid() {
		tun.SetVID(vid)
	}

	if err := tun.Create(unit); err != nil {
		return opennsl.TUNNEL, err
	}

	return tun.TunnelID(), nil
}

//
// VxlanIface
//
type VxlanIface struct {
	Port opennsl.Port
	Vlan opennsl.Vlan
	MAC  net.HardwareAddr

	IfaceID opennsl.L3IfaceID
}

func NewVxlanIface(port opennsl.Port, vid opennsl.Vlan, mac net.HardwareAddr) *VxlanIface {
	return &VxlanIface{
		Port: port,
		Vlan: vid,
		MAC:  mac,
	}
}

func NewVxlanAccessIface(port opennsl.Port, vid opennsl.Vlan) *VxlanIface {
	dummyMAC, _ := net.ParseMAC("00:00:01:00:00:01")
	return NewVxlanIface(port, vid, dummyMAC)
}

func NewVxlanNetworkIface(port opennsl.Port, vid opennsl.Vlan, localMAC net.HardwareAddr) *VxlanIface {
	return NewVxlanIface(port, vid, localMAC)
}

func (v *VxlanIface) Init(unit int) error {
	if err := vlan_create_add_port(unit, v.Vlan, v.Port); err != nil {
		return err
	}

	ifaceID, err := create_l3_interface(unit, v.MAC, v.Vlan)
	if err != nil {
		return err
	}

	v.IfaceID = ifaceID

	return nil
}

//
// VxlanTunnel
//
type VxlanTunnel struct {
	LocalMAC    net.HardwareAddr
	RemoteMAC   net.HardwareAddr
	LocalIp     net.IP
	RemoteIp    net.IP
	Vlan        opennsl.Vlan
	Criteria    opennsl.VxlanPortMatch
	EgressFlags []opennsl.L3Flags
	VxlanFlags  []opennsl.VxlanPortFlags
}

func NewVxlanNetworkTunnel(localMAC, remoteMAC net.HardwareAddr, localIp, remoteIp net.IP) *VxlanTunnel {
	return &VxlanTunnel{
		LocalMAC:    localMAC,
		RemoteMAC:   remoteMAC,
		LocalIp:     localIp,
		RemoteIp:    remoteIp,
		Vlan:        opennsl.VLAN_ID_INVALID,
		Criteria:    opennsl.VXLAN_PORT_MATCH_VN_ID,
		EgressFlags: []opennsl.L3Flags{},
		VxlanFlags:  []opennsl.VxlanPortFlags{},
	}
}

func NewVxlanNetworkMCTunnel(localMAC, remoteMAC net.HardwareAddr, localIp, remoteIp net.IP, vid opennsl.Vlan) *VxlanTunnel {
	return &VxlanTunnel{
		LocalMAC:    localMAC,
		RemoteMAC:   remoteMAC,
		LocalIp:     localIp,
		RemoteIp:    remoteIp,
		Vlan:        vid,
		Criteria:    opennsl.VXLAN_PORT_MATCH_NONE,
		EgressFlags: []opennsl.L3Flags{opennsl.L3_IPMC},
		VxlanFlags:  []opennsl.VxlanPortFlags{opennsl.VXLAN_PORT_MULTICAST},
	}
}

//
// Vxlan Virtual Port
//
func NewVxlanAccessVPort(unit int, iface VxlanIface, vpn opennsl.Vpn) (opennsl.GPort, error) {
	remoteMAC, _ := net.ParseMAC("00:00:01:00:00:01")
	l3eg, err := create_egr_obj(unit, iface.IfaceID, remoteMAC, iface.Port, iface.Vlan)
	if err != nil {
		return 0, err
	}

	gport, err := iface.Port.GPortGet(unit)
	if err != nil {
		return 0, err
	}

	vport, err := create_vxlan_acc_vp(unit, vpn, gport, opennsl.VXLAN_PORT_MATCH_PORT, l3eg, iface.Vlan)
	if err != nil {
		return 0, err
	}

	return vport, nil
}

func NewVxlanNetworkVPort(unit int, iface VxlanIface, vpn opennsl.Vpn, tun *VxlanTunnel) (opennsl.GPort, error) {
	l3eg, err := create_egr_obj(unit, iface.IfaceID, tun.RemoteMAC, iface.Port, iface.Vlan, tun.EgressFlags...)
	if err != nil {
		return 0, err
	}

	gport, err := iface.Port.GPortGet(unit)
	if err != nil {
		return 0, err
	}

	tun_init, err := tunnel_initiator_setup(unit, tun.LocalIp, tun.RemoteIp)
	if err != nil {
		return 0, err
	}

	tun_term, err := tunnel_terminator_setup(unit, tun.LocalIp, tun.RemoteIp, tun.Vlan, tun_init)
	if err != nil {
		return 0, err
	}

	vport, err := create_vxlan_net_vp(unit, vpn, gport, tun.Criteria, l3eg, tun_init, tun_term, tun.VxlanFlags...)
	if err != nil {
		return 0, err
	}

	if _, err = add_to_l2_station(unit, tun.LocalMAC, iface.Vlan); err != nil {
		log.Errorf("add_to_l2_station() error. %s", err)
	}

	return vport, nil
}
