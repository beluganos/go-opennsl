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
#include <stdlib.h>
#include <opennsl/vxlanX.h>
#include "helper.h"
*/
import "C"

import (
	"unsafe"
)

//
// VxlanVpnConfig
//
type VxlanVpnConfig C.opennsl_vxlan_vpn_config_t

func (v *VxlanVpnConfig) C() *C.opennsl_vxlan_vpn_config_t {
	return (*C.opennsl_vxlan_vpn_config_t)(v)
}

func (v *VxlanVpnConfig) BroadcastGroup() Multicast {
	return Multicast(v.broadcast_group)
}

func (v *VxlanVpnConfig) SetBroadcastGroup(group Multicast) {
	v.broadcast_group = group.C()
}

func (v *VxlanVpnConfig) DefaultVlan() Vlan {
	return Vlan(v.default_vlan)
}

func (v *VxlanVpnConfig) SetDefaultVlan(vid Vlan) {
	v.default_vlan = vid.C()
}

func (v *VxlanVpnConfig) EgressServiceTPID() uint16 {
	return uint16(v.egress_service_tpid)
}

func (v *VxlanVpnConfig) SetEgressServiceTPID(tpid uint16) {
	v.egress_service_tpid = C.uint16(tpid)
}

func (v *VxlanVpnConfig) EgressServiceVlan() Vlan {
	return Vlan(v.egress_service_vlan)
}

func (v *VxlanVpnConfig) SetEgressServiceVlan(vid Vlan) {
	v.egress_service_vlan = vid.C()
}

func (v *VxlanVpnConfig) Flags() VxlanVpnFlags {
	return VxlanVpnFlags(v.flags)
}

func (v *VxlanVpnConfig) SetFlags(flags VxlanVpnFlags) {
	v.flags = flags.C()
}

func (v *VxlanVpnConfig) MatchPortClass() GPort {
	return GPort(v.match_port_class)
}

func (v *VxlanVpnConfig) SetMatchPortClass(port GPort) {
	v.match_port_class = port.C()
}

func (v *VxlanVpnConfig) Cfi() uint8 {
	return uint8(v.pkt_cfi)
}

func (v *VxlanVpnConfig) SetCfi(cfi uint8) {
	v.pkt_cfi = C.uint8(cfi)
}

func (v *VxlanVpnConfig) Pri() uint8 {
	return uint8(v.pkt_pri)
}

func (v *VxlanVpnConfig) SetPri(pri uint8) {
	v.pkt_pri = C.uint8(pri)
}

func (v *VxlanVpnConfig) ProtocolPkt() *VlanProtocolPacketCtrl {
	return (*VlanProtocolPacketCtrl)(&v.protocol_pkt)
}

func (v *VxlanVpnConfig) UnknownMuticastGroup() Multicast {
	return Multicast(v.unknown_multicast_group)
}

func (v *VxlanVpnConfig) SetUnknownMuticastGroup(group Multicast) {
	v.unknown_multicast_group = group.C()
}

func (v *VxlanVpnConfig) UnknownUnicastGroup() Multicast {
	return Multicast(v.unknown_unicast_group)
}

func (v *VxlanVpnConfig) SetUnknownUnicastGroup(group Multicast) {
	v.unknown_unicast_group = group.C()
}

func (v *VxlanVpnConfig) Vlan() Vlan {
	return Vlan(v.vlan)
}

func (v *VxlanVpnConfig) SetVlan(vid Vlan) {
	v.vlan = vid.C()
}

func (v *VxlanVpnConfig) VNID() VNID {
	return VNID(v.vnid)
}

func (v *VxlanVpnConfig) SetVNID(vnid VNID) {
	v.vnid = vnid.C()
}

func (v *VxlanVpnConfig) Vpn() Vpn {
	return Vpn(v.vpn)
}

func (v *VxlanVpnConfig) SetVpn(vpn Vpn) {
	v.vpn = vpn.C()
}

//
// API
//
func NewVxlanVpnConfig() *VxlanVpnConfig {
	cfg := &VxlanVpnConfig{}
	cfg.Init()
	return cfg
}

func (v *VxlanVpnConfig) Init() {
	C.opennsl_vxlan_vpn_config_t_init(v.C())
}

func (v *VxlanVpnConfig) Create(unit int) error {
	rc := C.opennsl_vxlan_vpn_create(C.int(unit), v.C())
	return ParseError(rc)
}

func (v *VxlanVpnConfig) Destroy(unit int, l2vpn Vpn) error {
	rc := C.opennsl_vxlan_vpn_destroy(C.int(unit), l2vpn.C())
	return ParseError(rc)
}

func VxlanVpnConfigDestroyAll(unit int) error {
	rc := C.opennsl_vxlan_vpn_destroy_all(C.int(unit))
	return ParseError(rc)
}

func (v Vpn) VxlanConfigGet(unit int) (*VxlanVpnConfig, error) {
	cfg := NewVxlanVpnConfig()

	rc := C.opennsl_vxlan_vpn_get(C.int(unit), v.C(), cfg.C())
	if err := ParseError(rc); err != nil {
		return nil, err
	}
	return cfg, nil
}

type VxlanVpnTraverseCallback func(int, *VxlanVpnConfig) int

var vxLanVpnTraverseCallbacks = NewCallbackMap()

//export go_opennsl_vxlan_vpn_traverse_cb
func go_opennsl_vxlan_vpn_traverse_cb(unit C.int, config *C.opennsl_vxlan_vpn_config_t, data unsafe.Pointer) int {
	n := (*uint64)(data)
	if h, ok := vxLanVpnTraverseCallbacks.Get(*n); ok {
		callback := h.(VxlanVpnTraverseCallback)
		return callback(int(unit), (*VxlanVpnConfig)(config))
	}

	return int(E_PARAM)
}

func VxlanVpnTraverse(unit int, callback VxlanVpnTraverseCallback) error {
	n := vxLanVpnTraverseCallbacks.Add(callback)
	defer vxLanVpnTraverseCallbacks.Del(n)

	rc := C.opennsl_vxlan_vpn_traverse(C.int(unit), (C.opennsl_vxlan_vpn_traverse_cb)(C._opennsl_vxlan_vpn_traverse_cb), unsafe.Pointer(&n))
	return ParseError(rc)
}
