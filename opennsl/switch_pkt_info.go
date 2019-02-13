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

package opennsl

/*
#include <opennsl/switch.h>
*/
import "C"

import (
	"net"
)

//
// SwitchPktInfo
//
type SwitchPktInfo C.opennsl_switch_pkt_info_t

func (v *SwitchPktInfo) C() *C.opennsl_switch_pkt_info_t {
	return (*C.opennsl_switch_pkt_info_t)(v)
}

func (v *SwitchPktInfo) Flags() SwitchPktInfoFlags {
	return SwitchPktInfoFlags(v.flags)
}

func (v *SwitchPktInfo) SetFlags(flags SwitchPktInfoFlags) {
	v.flags = flags.C()
}

func (v *SwitchPktInfo) SrcGPort() int {
	return int(v.src_gport)
}

func (v *SwitchPktInfo) SetSrcGPort(gport GPort) {
	v.src_gport = gport.C()
}

func (v *SwitchPktInfo) VID() Vlan {
	return Vlan(v.vid)
}

func (v *SwitchPktInfo) SetVID(vid Vlan) {
	v.vid = vid.C()
}

func (v *SwitchPktInfo) Ethertype() Ethertype {
	return Ethertype(v.ethertype)
}

func (v *SwitchPktInfo) SetEthertype(ethertype Ethertype) {
	v.ethertype = ethertype.C()
}

func (v *SwitchPktInfo) MAC() (net.HardwareAddr, net.HardwareAddr) {
	return ParseMAC(v.dst_mac), ParseMAC(v.src_mac)
}

func (v *SwitchPktInfo) SetSrcMAC(mac net.HardwareAddr) {
	v.src_mac = NewMAC(mac)
}

func (v *SwitchPktInfo) SetDestMAC(mac net.HardwareAddr) {
	v.dst_mac = NewMAC(mac)
}

func (v *SwitchPktInfo) IP4() (net.IP, net.IP) {
	return ParseIP4(v.dip), ParseIP4(v.sip)
}

func (v *SwitchPktInfo) SetSrcIP4(ip net.IP) error {
	ip4, err := NewIP4(ip)
	if err == nil {
		v.sip = ip4
	}
	return err
}

func (v *SwitchPktInfo) SetDestIP4(ip net.IP) error {
	ip4, err := NewIP4(ip)
	if err == nil {
		v.dip = ip4
	}
	return err
}

func (v *SwitchPktInfo) IP6() (net.IP, net.IP) {
	return ParseIP6(v.dip6), ParseIP6(v.sip6)
}

func (v *SwitchPktInfo) SetSrcIP6(ip net.IP) error {
	ip6, err := NewIP6(ip)
	if err == nil {
		v.sip6 = ip6
	}
	return err
}

func (v *SwitchPktInfo) SetDestIP6(ip net.IP) error {
	ip6, err := NewIP6(ip)
	if err == nil {
		v.dip6 = ip6
	}
	return err
}

func (v *SwitchPktInfo) Protocol() uint8 {
	return uint8(v.protocol)
}

func (v *SwitchPktInfo) SetProtocol(protocol uint8) {
	v.protocol = C.uint8(protocol)
}

func (v *SwitchPktInfo) L4Port() (uint32, uint32) {
	return uint32(v.dst_l4_port), uint32(v.src_l4_port)
}

func (v *SwitchPktInfo) SetSrcL4Port(port uint32) {
	v.src_l4_port = C.uint32(port)
}

func (v *SwitchPktInfo) SetDestL4Port(port uint32) {
	v.dst_l4_port = C.uint32(port)
}

func (v *SwitchPktInfo) TrunkGPort() int {
	return int(v.trunk_gport)
}

func (v *SwitchPktInfo) SetTrunkGPort(gport GPort) {
	v.trunk_gport = gport.C()
}

func (v *SwitchPktInfo) MPIface() Iface {
	return Iface(v.mpintf)
}

func (v *SwitchPktInfo) SetMPIface(mpintf Iface) {
	v.mpintf = mpintf.C()
}

func (v *SwitchPktInfo) FwdReason() SwitchPktHashInfoFwdReason {
	return SwitchPktHashInfoFwdReason(v.fwd_reason)
}

func (v *SwitchPktInfo) SetFwdReason(reason SwitchPktHashInfoFwdReason) {
	v.fwd_reason = reason.C()
}

func (v *SwitchPktInfo) Init() {
	C.opennsl_switch_pkt_info_t_init(v.C())
}

func (v *SwitchPktInfo) HashGet(unit int) (GPort, Iface, error) {
	var gport C.opennsl_gport_t = 0
	var iface C.opennsl_if_t = 0
	rc := C.opennsl_switch_pkt_info_hash_get(C.int(unit), v.C(), &gport, &iface)
	return GPort(gport), Iface(iface), ParseError(rc)
}
