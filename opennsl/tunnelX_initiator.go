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
#include <opennsl/types.h>
#include <opennsl/tunnelX.h>
#include "helper.h"
*/
import "C"

import (
	"net"
	"unsafe"
)

//
// TunelInitiator
//
type TunnelInitiator C.opennsl_tunnel_initiator_t

func (v *TunnelInitiator) C() *C.opennsl_tunnel_initiator_t {
	return (*C.opennsl_tunnel_initiator_t)(v)
}

func (v *TunnelInitiator) Flags() TunnelFlags {
	return TunnelFlags(v.flags)
}

func (v *TunnelInitiator) SetFlags(flags TunnelFlags) {
	v.flags = flags.C()
}

func (v *TunnelInitiator) Type() TunnelType {
	return TunnelType(v._type)
}

func (v *TunnelInitiator) SetType(t TunnelType) {
	v._type = t.C()
}

func (v *TunnelInitiator) TTL() int {
	return int(v.ttl)
}

func (v *TunnelInitiator) SetTTL(ttl int) {
	v.ttl = C.int(ttl)
}

func (v *TunnelInitiator) DstMAC() net.HardwareAddr {
	return ParseMAC(v.dmac)
}

func (v *TunnelInitiator) SetDstMAC(mac net.HardwareAddr) {
	v.dmac = NewMAC(mac)
}

func (v *TunnelInitiator) SrcIP4() net.IP {
	return ParseIP4(v.sip)
}

func (v *TunnelInitiator) SetSrcIP4(ip net.IP) error {
	sip, err := NewIP4(ip)
	if err != nil {
		return err
	}
	v.sip = sip
	return nil
}

func (v *TunnelInitiator) DstIP4() net.IP {
	return ParseIP4(v.dip)
}

func (v *TunnelInitiator) SetDstIP4(ip net.IP) error {
	dip, err := NewIP4(ip)
	if err != nil {
		return err
	}
	v.dip = dip
	return nil
}

func (v *TunnelInitiator) SrcIP6() net.IP {
	return ParseIP6(v.sip6)
}

func (v *TunnelInitiator) SetSrcIP6(ip net.IP) error {
	sip, err := NewIP6(ip)
	if err != nil {
		return err
	}
	v.sip6 = sip
	return nil
}

func (v *TunnelInitiator) DstIP6() net.IP {
	return ParseIP6(v.dip6)
}

func (v *TunnelInitiator) SetDstIP6(ip net.IP) error {
	dip, err := NewIP6(ip)
	if err != nil {
		return err
	}
	v.dip6 = dip
	return nil
}

func (v *TunnelInitiator) FlowLabel() uint32 {
	return uint32(v.flow_label)
}

func (v *TunnelInitiator) SetFlowLabel(label uint32) {
	v.flow_label = C.uint32(label)
}

func (v *TunnelInitiator) DSCPSelect() TunnelDSCPSelect {
	return TunnelDSCPSelect(v.dscp_sel)
}

func (v *TunnelInitiator) SetDSCPSelect(s TunnelDSCPSelect) {
	v.dscp_sel = s.C()
}

func (v *TunnelInitiator) DSCP() int {
	return int(v.dscp)
}

func (v *TunnelInitiator) SetDSCP(dscp int) {
	v.dscp = C.int(dscp)
}

func (v *TunnelInitiator) DSCPMap() int {
	return int(v.dscp_map)
}

func (v *TunnelInitiator) SetDSCPMap(dscp int) {
	v.dscp_map = C.int(dscp)
}

func (v *TunnelInitiator) TunnelID() TunnelID {
	return TunnelID(v.tunnel_id)
}

func (v *TunnelInitiator) SetTunnelID(tunID TunnelID) {
	v.tunnel_id = tunID.C()
}

func (v *TunnelInitiator) UdpDstPort() L4Port {
	return L4Port(v.udp_dst_port)
}

func (v *TunnelInitiator) SetUdpDstPort(port L4Port) {
	v.udp_dst_port = C.uint16(port)
}

func (v *TunnelInitiator) UdpSrcPort() L4Port {
	return L4Port(v.udp_src_port)
}

func (v *TunnelInitiator) SetUdpSrcPort(port L4Port) {
	v.udp_src_port = C.uint16(port)
}

func (v *TunnelInitiator) SrcMAC() net.HardwareAddr {
	return ParseMAC(v.smac)
}

func (v *TunnelInitiator) SetSrcMAC(mac net.HardwareAddr) {
	v.smac = NewMAC(mac)
}

func (v *TunnelInitiator) MTU() int {
	return int(v.mtu)
}

func (v *TunnelInitiator) SetMTU(mtu int) {
	v.mtu = C.int(mtu)
}

func (v *TunnelInitiator) VID() Vlan {
	return Vlan(v.vlan)
}

func (v *TunnelInitiator) SetVID(vlan Vlan) {
	v.vlan = vlan.C()
}

func (v *TunnelInitiator) TPID() uint16 {
	return uint16(v.tpid)
}

func (v *TunnelInitiator) SetTPID(tpid uint16) {
	v.tpid = C.uint16(tpid)
}

func (v *TunnelInitiator) PktPri() uint8 {
	return uint8(v.pkt_pri)
}

func (v *TunnelInitiator) SetPktPri(pri uint8) {
	v.pkt_pri = C.uint8(pri)
}

func (v *TunnelInitiator) PktCfi() uint8 {
	return uint8(v.pkt_cfi)
}

func (v *TunnelInitiator) SetPktCfi(cfi uint8) {
	v.pkt_cfi = C.uint8(cfi)
}

func (v *TunnelInitiator) IPv4ID() uint16 {
	return uint16(v.ip4_id)
}

func (v *TunnelInitiator) SetIPv4ID(ip4ID uint16) {
	v.ip4_id = C.uint16(ip4ID)
}

func (v *TunnelInitiator) L3IfaceID() L3IfaceID {
	return L3IfaceID(v.l3_intf_id)
}

func (v *TunnelInitiator) SetL3IfaceID(l3iface L3IfaceID) {
	v.l3_intf_id = l3iface.C()
}

func (v *TunnelInitiator) SpanID() uint16 {
	return uint16(v.span_id)
}

func (v *TunnelInitiator) SetSpanID(span uint16) {
	v.span_id = C.uint16(span)
}

func (v *TunnelInitiator) AuxData() uint32 {
	return uint32(v.aux_data)
}

func (v *TunnelInitiator) SetAuxData(data uint32) {
	v.aux_data = C.uint32(data)
}

func (v *TunnelInitiator) OutLIFCountingProfile() int {
	return int(v.outlif_counting_profile)
}

func (v *TunnelInitiator) SetOutLIFCountingProfile(p int) {
	v.outlif_counting_profile = C.int(p)
}

//
// API
//
func NewTunnelInitiator(tunType TunnelType) *TunnelInitiator {
	tun := &TunnelInitiator{}
	tun.Init()
	tun.SetType(tunType)
	return tun
}

func TunnelInitiatorInit(v *TunnelInitiator) {
	C.opennsl_tunnel_initiator_t_init(v.C())
}

func (v *TunnelInitiator) Init() {
	TunnelInitiatorInit(v)
}

func TunnelInitiatorCreate(unit int, l3iface *L3Iface, v *TunnelInitiator) error {
	rc := func() C.int {
		if InfoDeviceIsDNX(unit) {
			return C.opennsl_tunnel_initiator_create(C.int(unit), l3iface.C(), v.C())
		}
		return C.opennsl_tunnel_initiator_set(C.int(unit), l3iface.C(), v.C())
	}()

	return ParseError(rc)
}

func (v *TunnelInitiator) Create(unit int, l3iface *L3Iface) error {
	return TunnelInitiatorCreate(unit, l3iface, v)
}

func (v *L3Iface) TunnelInitiatorCreate(unit int, tunnel *TunnelInitiator) error {
	return TunnelInitiatorCreate(unit, v, tunnel)
}

func TunnelInitiatorClear(unit int, l3iface *L3Iface) error {
	rc := C.opennsl_tunnel_initiator_clear(C.int(unit), l3iface.C())
	return ParseError(rc)
}

func (v *L3Iface) TunnelInitiatorClear(unit int) error {
	return TunnelInitiatorClear(unit, v)
}

func TunnelInitiatorGet(unit int, l3iface *L3Iface) (*TunnelInitiator, error) {
	tunnel := NewTunnelInitiator(TunnelTypeNone)

	rc := C.opennsl_tunnel_initiator_get(C.int(unit), l3iface.C(), tunnel.C())
	if err := ParseError(rc); err != nil {
		return nil, err
	}

	return tunnel, nil
}

func (v *L3Iface) TunnelInitiatorGet(unit int) (*TunnelInitiator, error) {
	return TunnelInitiatorGet(unit, v)
}

type TunnelInitiatorTraverseCallback func(int, *TunnelInitiator) OpenNSLError

var tunnelInitiatorTraverseCallbacks = NewCallbackMap()

//export go_opennsl_tunnel_initiator_traverse_cb
func go_opennsl_tunnel_initiator_traverse_cb(c_unit C.int, c_tunnel *C.opennsl_tunnel_initiator_t, c_data unsafe.Pointer) int {
	n := (*uint64)(c_data)
	if h, ok := tunnelInitiatorTraverseCallbacks.Get(*n); ok {
		callback := h.(TunnelInitiatorTraverseCallback)
		rc := callback(int(c_unit), (*TunnelInitiator)(c_tunnel))
		return int(rc)
	}

	return int(E_PARAM)
}

func TunnelInitiatorTraverse(unit int, callback TunnelInitiatorTraverseCallback) error {
	n := tunnelInitiatorTraverseCallbacks.Add(callback)
	defer tunnelInitiatorTraverseCallbacks.Del(n)

	rc := C.opennsl_tunnel_initiator_traverse(C.int(unit), C.opennsl_tunnel_initiator_traverse_cb(C._opennsl_tunnel_initiator_traverse_cb), unsafe.Pointer(&n))
	return ParseError(rc)
}
