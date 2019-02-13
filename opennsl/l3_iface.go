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
#include <opennsl/l3.h>
*/
import "C"

import (
	"fmt"
	"net"
)

//
// L3IfaceID
//
type L3IfaceID C.opennsl_if_t

func (v L3IfaceID) C() C.opennsl_if_t {
	return C.opennsl_if_t(v)
}

const L3_IFACE L3IfaceID = 0

//
// L3IfaceIDMask
//
type L3IfaceIDMask C.opennsl_if_t

func (v L3IfaceIDMask) C() C.opennsl_if_t {
	return C.opennsl_if_t(v)
}

const (
	L3_IFACE_ID_MASK_NONE  L3IfaceIDMask = 0
	L3_IFACE_ID_MASK_EXACT L3IfaceIDMask = -1
)

//
// L3Iface
//
type L3Iface C.opennsl_l3_intf_t

func (v *L3Iface) C() *C.opennsl_l3_intf_t {
	return (*C.opennsl_l3_intf_t)(v)
}

func (v *L3Iface) Flags() L3Flags {
	return L3Flags(v.l3a_flags)
}

func (v *L3Iface) SetFlags(flags L3Flags) {
	v.l3a_flags = flags.C()
}

func (v *L3Iface) VRF() Vrf {
	return Vrf(v.l3a_vrf)
}

func (v *L3Iface) SetVRF(vrf Vrf) {
	v.l3a_vrf = vrf.C()
}

func (v *L3Iface) IfaceID() L3IfaceID {
	return L3IfaceID(v.l3a_intf_id)
}

func (v *L3Iface) SetIfaceID(ifaceID L3IfaceID) {
	v.l3a_intf_id = ifaceID.C()
}

func (v *L3Iface) MAC() net.HardwareAddr {
	return ParseMAC(v.l3a_mac_addr)
}

func (v *L3Iface) SetMAC(mac net.HardwareAddr) {
	v.l3a_mac_addr = NewMAC(mac)
}

func (v *L3Iface) VID() Vlan {
	return Vlan(v.l3a_vid)
}

func (v *L3Iface) SetVID(vid Vlan) {
	v.l3a_vid = vid.C()
}

func (v *L3Iface) TTL() int {
	return int(v.l3a_ttl)
}

func (v *L3Iface) SetTTL(ttl int) {
	v.l3a_ttl = C.int(ttl)
}

func (v *L3Iface) MTU() int {
	return int(v.l3a_mtu)
}

func (v *L3Iface) SetMTU(mtu int) {
	v.l3a_mtu = C.int(mtu)
}

func (v *L3Iface) MTUFwd() int {
	return int(v.l3a_mtu_forwarding)
}

func (v *L3Iface) SetMTUFwd(mtu int) {
	v.l3a_mtu_forwarding = C.int(mtu)
}

func (v *L3Iface) DscpQoS() L3IfaceQoS {
	return L3IfaceQoS(v.dscp_qos)
}

func (v *L3Iface) SetDscpQoS(qos L3IfaceQoS) {
	v.dscp_qos = *qos.C()
}

func (v *L3Iface) IPv4OptionsProfile() int {
	return int(v.l3a_ip4_options_profile_id)
}

func (v *L3Iface) SetIPv4OptionsProfile(profile int) {
	v.l3a_ip4_options_profile_id = C.int(profile)
}

func (v *L3Iface) NativeRoutingVlanTags() uint8 {
	return uint8(v.native_routing_vlan_tags)
}

func (v *L3Iface) SetNativeRoutingVlanTags(tags uint8) {
	v.native_routing_vlan_tags = C.uint8(tags)
}

//
// API
//
func NewL3Iface() *L3Iface {
	iface := &L3Iface{}
	iface.Init()
	return iface
}

func L3IfaceInit(v *L3Iface) {
	C.opennsl_l3_intf_t_init(v.C())
}

func (v *L3Iface) String() string {
	return fmt.Sprintf("L3Iface(%d, mac:%s, vid:%d, vrf:%d)", v.IfaceID(), v.MAC(), v.VID(), v.VRF())
}

func (v *L3Iface) Init() {
	L3IfaceInit(v)
}

func L3IfaceCreate(unit int, v *L3Iface) error {
	rc := C.opennsl_l3_intf_create(C.int(unit), v.C())
	return ParseError(rc)
}

func (v *L3Iface) Create(unit int) error {
	return L3IfaceCreate(unit, v)
}

func L3IfaceDelete(unit int, v *L3Iface) error {
	rc := C.opennsl_l3_intf_delete(C.int(unit), v.C())
	return ParseError(rc)
}

func (v *L3Iface) Delete(unit int) error {
	return L3IfaceDelete(unit, v)
}

func L3IfaceFind(unit int, mac net.HardwareAddr, vid Vlan) (*L3Iface, error) {
	l3iface := NewL3Iface()
	l3iface.SetMAC(mac)
	l3iface.SetVID(vid)

	rc := C.opennsl_l3_intf_find(C.int(unit), l3iface.C())
	if rc == C.OPENNSL_E_NOT_FOUND {
		return nil, nil
	}

	if err := ParseError(rc); err != nil {
		return nil, err
	}

	return l3iface, nil
}

func L3IfaceFindVID(unit int, ifaceID L3IfaceID, vid Vlan) (*L3Iface, error) {
	l3iface := NewL3Iface()
	l3iface.SetIfaceID(ifaceID)
	l3iface.SetVID(vid)

	rc := C.opennsl_l3_intf_find_vlan(C.int(unit), l3iface.C())
	if rc == C.OPENNSL_E_NOT_FOUND {
		return nil, nil
	}

	if err := ParseError(rc); err != nil {
		return nil, err
	}

	return l3iface, nil
}

func L3IfaceGet(unit int, ifaceID L3IfaceID) (*L3Iface, error) {
	l3iface := NewL3Iface()
	l3iface.SetIfaceID(ifaceID)

	rc := C.opennsl_l3_intf_get(C.int(unit), l3iface.C())
	if err := ParseError(rc); err != nil {
		return nil, err
	}

	return l3iface, nil
}
