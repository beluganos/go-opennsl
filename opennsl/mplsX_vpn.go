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
#include <opennsl/mplsX.h>
#include "helper.h"
*/
import "C"

import (
	"fmt"
	"strings"
	"unsafe"
)

//
// MplsVpnFlags
//
type MplsVpnFlags C.uint32

func (v MplsVpnFlags) C() C.uint32 {
	return C.uint32(v)
}

func NewMplsVpnFlags(flags ...MplsVpnFlags) MplsVpnFlags {
	v := MplsVpnFlags(0)
	for _, flag := range flags {
		v |= flag
	}
	return v
}

const (
	MPLS_VPN_NONE    MplsVpnFlags = 0
	MPLS_VPN_L3      MplsVpnFlags = C.OPENNSL_MPLS_VPN_L3
	MPLS_VPN_VPWS    MplsVpnFlags = C.OPENNSL_MPLS_VPN_VPWS
	MPLS_VPN_VPLS    MplsVpnFlags = C.OPENNSL_MPLS_VPN_VPLS
	MPLS_VPN_REPLACE MplsVpnFlags = C.OPENNSL_MPLS_VPN_REPLACE
	MPLS_VPN_WITH_ID MplsVpnFlags = C.OPENNSL_MPLS_VPN_WITH_ID
)

var mplsVpnFlags_names = map[MplsVpnFlags]string{
	MPLS_VPN_NONE:    "NONE",
	MPLS_VPN_L3:      "L3",
	MPLS_VPN_VPWS:    "VPWS",
	MPLS_VPN_VPLS:    "VPLS",
	MPLS_VPN_REPLACE: "REPLACE",
	MPLS_VPN_WITH_ID: "WITH_ID",
}

var mplsVpnFlags_values = map[string]MplsVpnFlags{
	"L3":      MPLS_VPN_L3,
	"VPWS":    MPLS_VPN_VPWS,
	"VPLS":    MPLS_VPN_VPLS,
	"REPLACE": MPLS_VPN_REPLACE,
	"WITH_ID": MPLS_VPN_WITH_ID,
}

func (v MplsVpnFlags) String() string {
	names := make([]string, 0, len(mplsVpnFlags_names))
	for val, name := range mplsVpnFlags_names {
		if v&val != 0 {
			names = append(names, name)
		}
	}
	return strings.Join(names, "|")
}

func ParseMplsVpnFlags(s string) (MplsVpnFlags, error) {
	if v, ok := mplsVpnFlags_values[s]; ok {
		return v, nil
	}
	return MPLS_VPN_NONE, fmt.Errorf("Invalid MplsVpnFlags. %s", s)
}

//
// MplsVpnConfig
//
type MplsVpnConfig C.opennsl_mpls_vpn_config_t

func (v *MplsVpnConfig) C() *C.opennsl_mpls_vpn_config_t {
	return (*C.opennsl_mpls_vpn_config_t)(v)
}

func (v *MplsVpnConfig) Flags() MplsVpnFlags {
	return MplsVpnFlags(v.flags)
}

func (v *MplsVpnConfig) SetFlags(flags MplsVpnFlags) {
	v.flags = flags.C()
}

func (v *MplsVpnConfig) Vpn() Vpn {
	return Vpn(v.vpn)
}

func (v *MplsVpnConfig) SetVpn(vpn Vpn) {
	v.vpn = vpn.C()
}

func (v *MplsVpnConfig) LookupID() int {
	return int(v.lookup_id)
}

func (v *MplsVpnConfig) SetLookupID(lookup int) {
	v.lookup_id = C.int(lookup)
}

func (v *MplsVpnConfig) BroadcastGroup() Multicast {
	return Multicast(v.broadcast_group)
}

func (v *MplsVpnConfig) SetBroadcastGroup(bc Multicast) {
	v.broadcast_group = bc.C()
}

func (v *MplsVpnConfig) UnknownUnicastGroup() Multicast {
	return Multicast(v.unknown_unicast_group)
}

func (v *MplsVpnConfig) SetUnknownUnicastGroup(uc Multicast) {
	v.unknown_unicast_group = uc.C()
}

func (v *MplsVpnConfig) UnknownMulticastGroup() Multicast {
	return Multicast(v.unknown_multicast_group)
}

func (v *MplsVpnConfig) SetUnknownMulticastGroup(mc Multicast) {
	v.unknown_multicast_group = mc.C()
}

func (v *MplsVpnConfig) Policer() Policer {
	return Policer(v.policer_id)
}

func (v *MplsVpnConfig) SetPolicer(policer Policer) {
	v.policer_id = policer.C()
}

func (v *MplsVpnConfig) ProtocolPkt() *VlanProtocolPacketCtrl {
	return (*VlanProtocolPacketCtrl)(&v.protocol_pkt)
}

//
// API
//
func NewMplsVpnConfig() *MplsVpnConfig {
	cfg := &MplsVpnConfig{}
	cfg.Init()
	return cfg
}

func (v *MplsVpnConfig) Init() {
	C.opennsl_mpls_vpn_config_t_init(v.C())
}

func (v *MplsVpnConfig) Create(unit int) error {
	rc := C.opennsl_mpls_vpn_id_create(C.int(unit), v.C())
	return ParseError(rc)
}

func (v *MplsVpnConfig) Destroy(unit int) error {
	return MplsVpnDestroy(unit, v.Vpn())
}

func MplsVpnDestroy(unit int, vpn Vpn) error {
	rc := C.opennsl_mpls_vpn_id_destroy(C.int(unit), vpn.C())
	return ParseError(rc)
}

func MplsVpnDestroyAll(unit int) error {
	rc := C.opennsl_mpls_vpn_id_destroy_all(C.int(unit))
	return ParseError(rc)
}

func MplsVpnGet(unit int, vpn Vpn) (*MplsVpnConfig, error) {
	cfg := MplsVpnConfig{}
	cfg.Init()

	rc := C.opennsl_mpls_vpn_id_get(C.int(unit), vpn.C(), cfg.C())
	if err := ParseError(rc); err != nil {
		return nil, err
	}

	return &cfg, nil
}

type MplsVpnTraverseCallback func(int, *MplsVpnConfig) int

var mplsVpnTraverseCallbacks = NewCallbackMap()

//export go_opennsl_mpls_vpn_traverse_cb
func go_opennsl_mpls_vpn_traverse_cb(unit C.int, cfg *C.opennsl_mpls_vpn_config_t, data unsafe.Pointer) int {
	n := (*uint64)(data)
	if h, ok := mplsVpnTraverseCallbacks.Get(*n); ok {
		callback := h.(MplsVpnTraverseCallback)
		return callback(int(unit), (*MplsVpnConfig)(cfg))
	}

	return int(E_PARAM)
}

func MplsVpnTraverse(unit int, callback MplsVpnTraverseCallback) error {
	n := mplsVpnTraverseCallbacks.Add(callback)
	defer mplsVpnTraverseCallbacks.Del(n)

	rc := C.opennsl_mpls_vpn_traverse(C.int(unit), C.opennsl_mpls_vpn_traverse_cb(C._opennsl_mpls_vpn_traverse_cb), unsafe.Pointer(&n))
	return ParseError(rc)
}
