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

package util

import (
	"net"

	"github.com/beluganos/go-opennsl/opennsl"

	log "github.com/sirupsen/logrus"
)

func ParseMAC(mac string) net.HardwareAddr {
	hwaddr, err := net.ParseMAC(mac)
	if err != nil {
		log.Errorf("net.ParseMAC(%s) error. %s", mac, err)
		return nil
	}
	return hwaddr
}

func ParseIPNet(ip string) *net.IPNet {
	_, ipnet, err := net.ParseCIDR(ip)
	if err != nil {
		log.Errorf("net.ParseCIDR(%s) error. %s", ip, err)
		return nil
	}
	return ipnet
}

func AddL2Station4(unit int, mac net.HardwareAddr, vid opennsl.Vlan, flags ...opennsl.L2StationFlags) (opennsl.L2StationID, error) {
	flags = append(flags, opennsl.L2_STATION_IPV4)
	l2st := opennsl.NewL2Station()
	l2st.SetFlags(opennsl.NewL2StationFlags(flags...))
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

func AddL2Station6(unit int, mac net.HardwareAddr, vid opennsl.Vlan, flags ...opennsl.L2StationFlags) (opennsl.L2StationID, error) {
	flags = append(flags, opennsl.L2_STATION_IPV6)
	l2st := opennsl.NewL2Station()
	l2st.SetFlags(opennsl.NewL2StationFlags(flags...))
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

func AddVlanPort(unit int, vid opennsl.Vlan, port opennsl.Port) error {
	untag := (vid == opennsl.VLAN_ID_DEFAULT)
	return addVlanPort(unit, vid, port, untag)
}

func AddUntaggedVlanPort(unit int, vid opennsl.Vlan, port opennsl.Port) error {
	return addVlanPort(unit, vid, port, true)
}

func addVlanPort(unit int, vid opennsl.Vlan, port opennsl.Port, untag bool) error {
	if _, err := vid.Create(unit); err != nil {
		log.Errorf("Vlan.Create error. %s", err)
		return err
	}

	pbmp := opennsl.NewPBmp()
	upbmp := opennsl.NewPBmp()
	pbmp.Add(port)
	if untag {
		upbmp.Add(port)
	}

	if err := vid.PortAdd(unit, pbmp, upbmp); err != nil {
		log.Errorf("Vlan.PortAdd error. %s", err)
		return err
	}

	return nil
}

func AddL2Addr(unit int, port opennsl.Port, vid opennsl.Vlan, mac net.HardwareAddr, flags ...opennsl.L2Flags) (*opennsl.L2Addr, error) {
	flags = append(flags, opennsl.L2_STATIC)
	l2addr := opennsl.NewL2Addr(mac, vid)
	l2addr.SetFlags(opennsl.NewL2Flags(flags...))
	l2addr.SetPort(port)

	if err := l2addr.Add(unit); err != nil {
		log.Errorf("L2Addr.Add error. %s", err)
		return nil, err
	}

	log.Debugf("AddL2Addr port %d, vlan %d, mac '%s'", port, vid, mac)
	return l2addr, nil
}

func AddHost(unit int, addr net.IP, l3eg opennsl.L3EgressID, vrf opennsl.Vrf) (*opennsl.L3Host, error) {
	host := opennsl.NewL3Host()
	host.SetIPAddr(addr)
	host.SetEgressID(l3eg)
	if vrf.IsValid() {
		host.SetVRF(vrf)
	}

	if err := host.Add(unit); err != nil {
		log.Errorf("L3Host.Add error. %s", err)
		return nil, err
	}

	log.Debugf("AddHost4 addr '%s', egr %d, vrf %d", addr, l3eg, vrf)
	return host, nil
}

func AddHost6(unit int, addr net.IP, l3eg opennsl.L3EgressID, vrf opennsl.Vrf) (*opennsl.L3Host, error) {
	host := opennsl.NewL3Host()
	host.SetFlags(opennsl.NewL3Flags(
		opennsl.L3_IP6,
	))
	if err := host.SetIP6Addr(addr); err != nil {
		log.Errorf("L3Host.SetIPv6Addr error. %s %s", addr, err)
		return nil, err
	}
	host.SetEgressID(l3eg)
	if vrf.IsValid() {
		host.SetVRF(vrf)
	}

	if err := host.Add(unit); err != nil {
		log.Errorf("L3Host.Add error. %s", err)
		return nil, err
	}

	log.Debugf("AddHost6 addr '%s', egr %d, vrf %d", addr, l3eg, vrf)
	return host, nil
}

func AddRoute(unit int, dest *net.IPNet, l3eg opennsl.L3EgressID, vrf opennsl.Vrf) (*opennsl.L3Route, error) {
	route := opennsl.NewL3Route()
	route.SetIP4Net(dest)
	route.SetEgressID(l3eg)
	if vrf.IsValid() {
		route.SetVRF(vrf)
	}

	if err := route.Add(unit); err != nil {
		log.Errorf("L3Route.Add(%s, %d, %d) error. %s", dest, l3eg, vrf, err)
		return nil, err
	}

	log.Debugf("AddRoute4 dst '%s', egr %d, vrf %d", dest, l3eg, vrf)
	return route, nil
}

func AddRoute6(unit int, dest *net.IPNet, l3eg opennsl.L3EgressID, vrf opennsl.Vrf) (*opennsl.L3Route, error) {
	route := opennsl.NewL3Route()
	route.SetFlags(opennsl.NewL3Flags(
		opennsl.L3_IP6,
	))
	if err := route.SetIP6Net(dest); err != nil {
		log.Errorf("L3Route.SetIP6Net error. %s %s", dest, err)
		return nil, err
	}
	route.SetEgressID(l3eg)
	if vrf.IsValid() {
		route.SetVRF(vrf)
	}

	if err := route.Add(unit); err != nil {
		log.Errorf("L3Route.Add(%s, %d, %d) error. %s", dest, l3eg, vrf, err)
		return nil, err
	}

	log.Debugf("AddRoute6 dst '%s', egr %d, vrf %d", dest, l3eg, vrf)
	return route, nil
}

func NewL3IfaceObj(unit int, port opennsl.Port, vid opennsl.Vlan, mac net.HardwareAddr) (*opennsl.L3Iface, error) {
	untag := (vid == opennsl.VLAN_ID_DEFAULT)
	return newL3IfaceObj(unit, port, vid, mac, untag)
}

func NewL3IfaceObjUntagged(unit int, port opennsl.Port, vid opennsl.Vlan, mac net.HardwareAddr) (*opennsl.L3Iface, error) {
	return newL3IfaceObj(unit, port, vid, mac, true)
}

func newL3IfaceObj(unit int, port opennsl.Port, vid opennsl.Vlan, mac net.HardwareAddr, untag bool) (*opennsl.L3Iface, error) {
	if err := addVlanPort(unit, vid, port, untag); err != nil {
		log.Errorf("NewL3Iface error. %s", err)
		return nil, err
	}

	l3if := opennsl.NewL3Iface()
	l3if.SetMAC(mac)
	l3if.SetVID(vid)
	if err := l3if.Create(unit); err != nil {
		log.Errorf("NewL3Iface error. %s", err)
		return nil, err
	}

	log.Debugf("NewL3Iface port %d, vid %d, iface %d", port, vid, l3if.IfaceID())
	return l3if, nil
}

func NewL3Egress(unit int, outPort opennsl.Port, vid opennsl.Vlan, iface opennsl.L3IfaceID, nh net.HardwareAddr, flags ...opennsl.L3Flags) (opennsl.L3EgressID, error) {
	l3eg := opennsl.NewL3Egress()
	l3eg.SetIfaceID(iface)
	l3eg.SetMAC(nh)
	l3eg.SetVID(vid)
	l3eg.SetPort(outPort)
	if len(flags) > 0 {
		l3eg.SetFlags(opennsl.NewL3Flags(flags...))
	}

	l3egrId, err := opennsl.L3_EGRESS.Create(unit, opennsl.NewL3Flags(flags...), l3eg)
	if err != nil {
		log.Errorf("NewL3Egress error. %s", err)
		return 0, err
	}

	log.Debugf("NewL3Egress port %d, vid %d, iface %d, mac %s, l3egr %d", outPort, vid, iface, nh, l3egrId)

	return l3egrId, nil
}

func TraverseL2Addrs(unit int) {
	log.Infof("TraverseL2Addrs START")
	err := opennsl.L2Traverse(unit, func(l2Unit int, l2addr *opennsl.L2Addr) opennsl.OpenNSLError {
		log.Infof("L2Addr: flg %08x, %s, vid %d, port %d", l2addr.Flags(), l2addr.MAC(), l2addr.VID(), l2addr.Port())
		log.Debugf("L2Addr: %v", l2addr)
		return opennsl.E_NONE
	})
	if err != nil {
		log.Errorf("L2Traverse error. %s", err)
	}
	log.Infof("TraverseL2Addrs END")
}

func TraverseL3Egresses(unit int) {
	log.Infof("TraverseL3Egresses START")
	err := opennsl.L3EgressTraverse(unit, func(l3Unit int, l3egrID opennsl.L3EgressID, l3egr *opennsl.L3Egress) opennsl.OpenNSLError {
		log.Infof("L3Egress[%d]: flg %08x/%08x, if %d, %s, vid %d, port %d", l3egrID, l3egr.Flags(), l3egr.Flags2(), l3egr.IfaceID(), l3egr.MAC(), l3egr.VID(), l3egr.Port())
		log.Debugf("L3Egress[%d]: %v", l3egrID, l3egr)
		return opennsl.E_NONE
	})
	if err != nil {
		log.Errorf("L3EgressTraverse error. %s", err)
	}
	log.Infof("TraverseL3Egresses END")
}

func TraverseL3Hosts(unit int) {
	log.Infof("TraverseHosts START")
	err := opennsl.L3HostTraverse(unit, 0, 0, 128, func(hostUnit int, hostIndex int, host *opennsl.L3Host) opennsl.OpenNSLError {
		log.Infof("Host[%d]: flg %08x, vrf %d, %s, l3egr %d", hostIndex, host.Flags(), host.VRF(), host.IPAddr(), host.EgressID())
		log.Debugf("Host[%d]: %v", hostIndex, host)
		return opennsl.E_NONE
	})
	if err != nil {
		log.Errorf("L3HostTraverse error. %s", err)
	}
	log.Infof("TraverseHosts END")
}

func TraverseL3Routes(unit int) {
	log.Infof("TraverseL3Route START")
	err := opennsl.L3RouteTraverse(unit, 0, 0, 128, func(rtUnit int, rtIndex int, route *opennsl.L3Route) opennsl.OpenNSLError {
		log.Infof("Route[%d]: flg %08x, vrf %d, %s, l3egr %d", rtIndex, route.Flags(), route.VRF(), route.IP4Net(), route.EgressID())
		log.Debugf("Route[%d]: %v", rtIndex, route)
		return opennsl.E_NONE
	})
	if err != nil {
		log.Errorf("L3RouteTraverse error. %s", err)
	}
	log.Infof("TraverseL3Route END")
}
