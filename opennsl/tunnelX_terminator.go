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
// TunnelTerminator
//
type TunnelTerminator C.opennsl_tunnel_terminator_t

func (v *TunnelTerminator) C() *C.opennsl_tunnel_terminator_t {
	return (*C.opennsl_tunnel_terminator_t)(v)
}

func (v *TunnelTerminator) Flags() TunnelFlags {
	return TunnelFlags(v.flags)
}

func (v *TunnelTerminator) SetFlags(flags TunnelFlags) {
	v.flags = flags.C()
}

func (v *TunnelTerminator) MulticastFlag() uint32 {
	return uint32(v.multicast_flag)
}

func (v *TunnelTerminator) SetMulticastFlag(flag uint32) {
	v.multicast_flag = C.uint32(flag)
}

func (v *TunnelTerminator) VRF() Vrf {
	return Vrf(v.vrf)
}

func (v *TunnelTerminator) SetVRF(vrf Vrf) {
	v.vrf = vrf.C()
}

func (v *TunnelTerminator) SrcIP4() net.IP {
	return ParseIP4(v.sip)
}

func (v *TunnelTerminator) SetSrcIP4(ip net.IP) error {
	sip, err := NewIP4(ip)
	if err != nil {
		return err
	}
	v.sip = sip
	return nil
}

func (v *TunnelTerminator) SrcIPMask4() net.IPMask {
	return ParseIP4Mask(v.sip_mask)
}

func (v *TunnelTerminator) SetSrcIPMask4(mask net.IPMask) error {
	smask, err := NewIP4Mask(mask)
	if err != nil {
		return err
	}
	v.sip_mask = smask
	return nil
}

func (v *TunnelTerminator) SrcIPNet4() *net.IPNet {
	return &net.IPNet{
		IP:   v.SrcIP4(),
		Mask: v.SrcIPMask4(),
	}
}

func (v *TunnelTerminator) SetSrcIPNet4(src *net.IPNet) error {
	if err := v.SetSrcIP4(src.IP); err != nil {
		return err
	}
	if err := v.SetSrcIPMask4(src.Mask); err != nil {
		return err
	}
	return nil
}

func (v *TunnelTerminator) DstIP4() net.IP {
	return ParseIP4(v.dip)
}

func (v *TunnelTerminator) SetDstIP4(ip net.IP) error {
	dip, err := NewIP4(ip)
	if err != nil {
		return err
	}
	v.dip = dip
	return nil
}

func (v *TunnelTerminator) DstIPMask4() net.IPMask {
	return ParseIP4Mask(v.dip_mask)
}

func (v *TunnelTerminator) SetDstIPMask4(mask net.IPMask) error {
	dmask, err := NewIP4Mask(mask)
	if err != nil {
		return err
	}
	v.dip_mask = dmask
	return nil
}

func (v *TunnelTerminator) DstIPNet4() *net.IPNet {
	return &net.IPNet{
		IP:   v.DstIP4(),
		Mask: v.DstIPMask4(),
	}
}

func (v *TunnelTerminator) SetDstIPNet4(src *net.IPNet) error {
	if err := v.SetDstIP4(src.IP); err != nil {
		return err
	}
	if err := v.SetDstIPMask4(src.Mask); err != nil {
		return err
	}
	return nil
}

func (v *TunnelTerminator) SrcIP6() net.IP {
	return ParseIP6(v.sip6)
}

func (v *TunnelTerminator) SetSrcIP6(ip net.IP) error {
	sip, err := NewIP6(ip)
	if err != nil {
		return err
	}
	v.sip6 = sip
	return nil
}

func (v *TunnelTerminator) SrcIPMask6() net.IPMask {
	return ParseIP6Mask(v.sip6_mask)
}

func (v *TunnelTerminator) SetSrcIPMask6(mask net.IPMask) error {
	smask, err := NewIP6Mask(mask)
	if err != nil {
		return err
	}
	v.sip6_mask = smask
	return nil
}

func (v *TunnelTerminator) SrcIPNet6() *net.IPNet {
	return &net.IPNet{
		IP:   v.SrcIP6(),
		Mask: v.SrcIPMask6(),
	}
}

func (v *TunnelTerminator) SetSrcIPNet6(src *net.IPNet) error {
	if err := v.SetSrcIP6(src.IP); err != nil {
		return err
	}
	if err := v.SetSrcIPMask6(src.Mask); err != nil {
		return err
	}
	return nil
}

func (v *TunnelTerminator) DstIP6() net.IP {
	return ParseIP6(v.dip6)
}

func (v *TunnelTerminator) SetDstIP6(ip net.IP) error {
	dip, err := NewIP6(ip)
	if err != nil {
		return err
	}
	v.dip6 = dip
	return nil
}

func (v *TunnelTerminator) DstIPMask6() net.IPMask {
	return ParseIP6Mask(v.dip6_mask)
}

func (v *TunnelTerminator) SetDstIPMask6(mask net.IPMask) error {
	dmask, err := NewIP6Mask(mask)
	if err != nil {
		return err
	}
	v.dip6_mask = dmask
	return nil
}

func (v *TunnelTerminator) DstIPNet6() *net.IPNet {
	return &net.IPNet{
		IP:   v.DstIP6(),
		Mask: v.DstIPMask6(),
	}
}

func (v *TunnelTerminator) SetDstIPNet6(src *net.IPNet) error {
	if err := v.SetDstIP6(src.IP); err != nil {
		return err
	}
	if err := v.SetDstIPMask6(src.Mask); err != nil {
		return err
	}
	return nil
}

func (v *TunnelTerminator) UdpDstPort() uint16 {
	return uint16(v.udp_dst_port)
}

func (v *TunnelTerminator) SetUdpDstPort(port uint16) {
	v.udp_dst_port = C.uint32(port)
}

func (v *TunnelTerminator) UdpSrcPort() uint16 {
	return uint16(v.udp_src_port)
}

func (v *TunnelTerminator) SetUdpSrcPort(port uint16) {
	v.udp_src_port = C.uint32(port)
}

func (v *TunnelTerminator) Type() TunnelType {
	return TunnelType(v._type)
}

func (v *TunnelTerminator) SetType(t TunnelType) {
	v._type = t.C()
}

func (v *TunnelTerminator) PBmp() *PBmp {
	return (*PBmp)(&v.pbmp)
}

func (v *TunnelTerminator) VID() Vlan {
	return Vlan(v.vlan)
}

func (v *TunnelTerminator) SetVID(vlan Vlan) {
	v.vlan = vlan.C()
}

func (v *TunnelTerminator) RemotePort() GPort {
	return GPort(v.remote_port)
}

func (v *TunnelTerminator) SetRemotePort(gport GPort) {
	v.remote_port = gport.C()
}

func (v *TunnelTerminator) TunnelID() TunnelID {
	return TunnelID(v.tunnel_id)
}

func (v *TunnelTerminator) SetTunnelID(tunID TunnelID) {
	v.tunnel_id = tunID.C()
}

func (v *TunnelTerminator) TunnelIface() Iface {
	return Iface(v.tunnel_if)
}

func (v *TunnelTerminator) SetTunnelIface(tunIface Iface) {
	v.tunnel_if = tunIface.C()
}

func (v *TunnelTerminator) QosMapID() int {
	return int(v.qos_map_id)
}

func (v *TunnelTerminator) SetQosMapID(mapid int) {
	v.qos_map_id = C.int(mapid)
}

func (v *TunnelTerminator) InLIFCountingProfile() int {
	return int(v.inlif_counting_profile)
}

func (v *TunnelTerminator) SetInLIFCountingProfile(p int) {
	v.inlif_counting_profile = C.int(p)
}

//
// API
//
func NewTunnelTerminator(tunType TunnelType) *TunnelTerminator {
	tun := &TunnelTerminator{}
	tun.Init()
	tun.SetType(tunType)
	return tun
}

func TunnelTerminatorInit(v *TunnelTerminator) {
	C.opennsl_tunnel_terminator_t_init(v.C())
}

func (v *TunnelTerminator) Init() {
	TunnelTerminatorInit(v)
}

func TunnelTerminatorCreate(unit int, v *TunnelTerminator) error {
	rc := func() C.int {
		if InfoDeviceIsDNX(unit) {
			return C.opennsl_tunnel_terminator_create(C.int(unit), v.C())
		}
		return C.opennsl_tunnel_terminator_add(C.int(unit), v.C())
	}()
	return ParseError(rc)
}

func (v *TunnelTerminator) Create(unit int) error {
	return TunnelTerminatorCreate(unit, v)
}

func TunnelTerminatorDelete(unit int, v *TunnelTerminator) error {
	rc := C.opennsl_tunnel_terminator_delete(C.int(unit), v.C())
	return ParseError(rc)
}

func (v *TunnelTerminator) Delete(unit int) error {
	return TunnelTerminatorDelete(unit, v)
}

func TunnelTerminatorUpdate(unit int, v *TunnelTerminator) error {
	rc := C.opennsl_tunnel_terminator_update(C.int(unit), v.C())
	return ParseError(rc)
}

func (v *TunnelTerminator) Update(unit int) error {
	return TunnelTerminatorUpdate(unit, v)
}

func TunnelTerminatorGet(unit int) (*TunnelTerminator, error) {
	tunnel := NewTunnelTerminator(TunnelTypeNone)

	rc := C.opennsl_tunnel_terminator_get(C.int(unit), tunnel.C())
	if err := ParseError(rc); err != nil {
		return nil, err
	}

	return tunnel, nil
}

type TunnelTerminatorTraverseCallback func(int, *TunnelTerminator) OpenNSLError

var tunnelTerminatorTraverseCallbacks = NewCallbackMap()

//export go_opennsl_tunnel_terminator_traverse_cb
func go_opennsl_tunnel_terminator_traverse_cb(unit C.int, tunnel *C.opennsl_tunnel_terminator_t, data unsafe.Pointer) int {
	n := (*uint64)(data)
	if h, ok := tunnelTerminatorTraverseCallbacks.Get(*n); ok {
		callback := h.(TunnelTerminatorTraverseCallback)
		rc := callback(int(unit), (*TunnelTerminator)(tunnel))
		return int(rc)
	}

	return int(E_PARAM)
}

func TunnelTerminatorTraverse(unit int, callback TunnelTerminatorTraverseCallback) error {
	return E_UNAVAIL.Error()

	//
	// bug:
	// opennsl_tunnel_terminator_traverse calls callback forever.
	//
	//n := tunnelTerminatorTraverseCallbacks.Add(callback)
	//defer tunnelTerminatorTraverseCallbacks.Del(n)

	//rc := C.opennsl_tunnel_terminator_traverse(C.int(unit), C.opennsl_tunnel_terminator_traverse_cb(C._opennsl_tunnel_terminator_traverse_cb), unsafe.Pointer(&n))
	//return ParseError(rc)
}
