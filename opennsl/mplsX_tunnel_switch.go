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
// MplsSwitchFlags
//
type MplsSwitchFlags C.uint32

func (v MplsSwitchFlags) C() C.uint32 {
	return C.uint32(v)
}

func NewMplsSwitchFlags(flags ...MplsSwitchFlags) MplsSwitchFlags {
	v := MplsSwitchFlags(0)
	for _, flag := range flags {
		v |= flag
	}
	return v
}

const (
	MPLS_SWITCH_NONE          MplsSwitchFlags = 0
	MPLS_SWITCH_INNER_TTL     MplsSwitchFlags = C.OPENNSL_MPLS_SWITCH_INNER_TTL
	MPLS_SWITCH_TTL_DECREMENT MplsSwitchFlags = C.OPENNSL_MPLS_SWITCH_TTL_DECREMENT
)

var mplsSwitchFlags_names = map[MplsSwitchFlags]string{
	MPLS_SWITCH_INNER_TTL:     "INNER_TTL",
	MPLS_SWITCH_TTL_DECREMENT: "TTL_DECREMENT",
}

var mplsSwitchFlags_values = map[string]MplsSwitchFlags{
	"INNER_TTL":     MPLS_SWITCH_INNER_TTL,
	"TTL_DECREMENT": MPLS_SWITCH_TTL_DECREMENT,
}

func (v MplsSwitchFlags) String() string {
	names := make([]string, 0, len(mplsSwitchFlags_names))
	for val, name := range mplsSwitchFlags_names {
		if v&val != 0 {
			names = append(names, name)
		}
	}
	return strings.Join(names, "|")
}

func ParseMplsSwitchFlags(s string) (MplsSwitchFlags, error) {
	if v, ok := mplsSwitchFlags_values[s]; ok {
		return v, nil
	}
	return MPLS_SWITCH_NONE, fmt.Errorf("invalid MplsSwitchFlags. %s", s)
}

//
// MplsSwitchAction
//
type MplsSwitchAction C.opennsl_mpls_switch_action_t

func (v MplsSwitchAction) C() C.opennsl_mpls_switch_action_t {
	return C.opennsl_mpls_switch_action_t(v)
}

const (
	MPLS_SWITCH_ACTION_SWAP        MplsSwitchAction = C.OPENNSL_MPLS_SWITCH_ACTION_SWAP
	MPLS_SWITCH_ACTION_PHP         MplsSwitchAction = C.OPENNSL_MPLS_SWITCH_ACTION_PHP
	MPLS_SWITCH_ACTION_POP         MplsSwitchAction = C.OPENNSL_MPLS_SWITCH_ACTION_POP
	MPLS_SWITCH_ACTION_POP_DIRECT  MplsSwitchAction = C.OPENNSL_MPLS_SWITCH_ACTION_POP_DIRECT
	MPLS_SWITCH_ACTION_NOP         MplsSwitchAction = C.OPENNSL_MPLS_SWITCH_ACTION_NOP
	MPLS_SWITCH_EGRESS_ACTION_PUSH MplsSwitchAction = C.OPENNSL_MPLS_SWITCH_EGRESS_ACTION_PUSH
	MPLS_SWITCH_ACTION_INVALID     MplsSwitchAction = C.OPENNSL_MPLS_SWITCH_ACTION_INVALID
)

var mplsSwitchAction_names = map[MplsSwitchAction]string{
	MPLS_SWITCH_ACTION_SWAP:        "SWAP",
	MPLS_SWITCH_ACTION_PHP:         "PHP",
	MPLS_SWITCH_ACTION_POP:         "POP",
	MPLS_SWITCH_ACTION_POP_DIRECT:  "POP_DIRECT",
	MPLS_SWITCH_ACTION_NOP:         "NOP",
	MPLS_SWITCH_EGRESS_ACTION_PUSH: "PUSH",
	MPLS_SWITCH_ACTION_INVALID:     "INVALID",
}

var mplsSwitchAction_values = map[string]MplsSwitchAction{
	"SWAP":       MPLS_SWITCH_ACTION_SWAP,
	"PHP":        MPLS_SWITCH_ACTION_PHP,
	"POP":        MPLS_SWITCH_ACTION_POP,
	"POP_DIRECT": MPLS_SWITCH_ACTION_POP_DIRECT,
	"NOP":        MPLS_SWITCH_ACTION_NOP,
	"PUSH":       MPLS_SWITCH_EGRESS_ACTION_PUSH,
	"INVALID":    MPLS_SWITCH_ACTION_INVALID,
}

func (v MplsSwitchAction) String() string {
	if s, ok := mplsSwitchAction_names[v]; ok {
		return s
	}
	return fmt.Sprintf("MplsSwitchAction(%d)", v)
}

func ParseMplsSwitchAction(s string) (MplsSwitchAction, error) {
	if v, ok := mplsSwitchAction_values[s]; ok {
		return v, nil
	}
	return MPLS_SWITCH_ACTION_INVALID, fmt.Errorf("Invalid MplsSwitchAction. %s", s)
}

//
// MplsTunnelSwitch
//
type MplsTunnelSwitch C.opennsl_mpls_tunnel_switch_t

func (v *MplsTunnelSwitch) C() *C.opennsl_mpls_tunnel_switch_t {
	return (*C.opennsl_mpls_tunnel_switch_t)(v)
}

func (v *MplsTunnelSwitch) Flags() MplsSwitchFlags {
	return MplsSwitchFlags(v.flags)
}

func (v *MplsTunnelSwitch) SetFlags(flags MplsSwitchFlags) {
	v.flags = flags.C()
}

func (v *MplsTunnelSwitch) Label() MplsLabel {
	return MplsLabel(v.label)
}

func (v *MplsTunnelSwitch) SetLabel(label MplsLabel) {
	v.label = label.C()
}

func (v *MplsTunnelSwitch) Port() GPort {
	return GPort(v.port)
}

func (v *MplsTunnelSwitch) SetPort(port GPort) {
	v.port = port.C()
}

func (v *MplsTunnelSwitch) Action() MplsSwitchAction {
	return MplsSwitchAction(v.action)
}

func (v *MplsTunnelSwitch) SetAction(action MplsSwitchAction) {
	v.action = action.C()
}

func (v *MplsTunnelSwitch) ActionIfBos() MplsSwitchAction {
	return MplsSwitchAction(v.action_if_bos)
}

func (v *MplsTunnelSwitch) SetActionIfBos(action MplsSwitchAction) {
	v.action_if_bos = action.C()
}

func (v *MplsTunnelSwitch) ActionIfNotBos() MplsSwitchAction {
	return MplsSwitchAction(v.action_if_not_bos)
}

func (v *MplsTunnelSwitch) SetActionIfNotBos(action MplsSwitchAction) {
	v.action_if_not_bos = action.C()
}

func (v *MplsTunnelSwitch) MCGroup() Multicast {
	return Multicast(v.mc_group)
}

func (v *MplsTunnelSwitch) SetMCGroup(mcGroup Multicast) {
	v.mc_group = mcGroup.C()
}

func (v *MplsTunnelSwitch) ExpMap() int {
	return int(v.exp_map)
}

func (v *MplsTunnelSwitch) SetExpMap(expMap int) {
	v.exp_map = C.int(expMap)
}

func (v *MplsTunnelSwitch) InternalPri() int {
	return int(v.int_pri)
}

func (v *MplsTunnelSwitch) SetInternalPri(pri int) {
	v.int_pri = C.int(pri)
}

func (v *MplsTunnelSwitch) Policer() Policer {
	return Policer(v.policer_id)
}

func (v *MplsTunnelSwitch) SetPolicer(policer Policer) {
	v.policer_id = policer.C()
}

func (v *MplsTunnelSwitch) Vpn() Vpn {
	return Vpn(v.vpn)
}

func (v *MplsTunnelSwitch) SetVpn(vpn Vpn) {
	v.vpn = vpn.C()
}

func (v *MplsTunnelSwitch) EgressLabel() *MplsEgressLabel {
	return (*MplsEgressLabel)(&v.egress_label)
}

func (v *MplsTunnelSwitch) Egress() L3EgressID {
	return L3EgressID(v.egress_if)
}

func (v *MplsTunnelSwitch) SetEgress(l3eg L3EgressID) {
	v.egress_if = l3eg.C()
}

func (v *MplsTunnelSwitch) Ingress() L3IngressID {
	return L3IngressID(v.ingress_if)
}

func (v *MplsTunnelSwitch) SetIngress(l3ing L3IngressID) {
	v.ingress_if = l3ing.C()
}

func (v *MplsTunnelSwitch) MTU() int {
	return int(v.mtu)
}

func (v *MplsTunnelSwitch) SetMTU(mtu int) {
	v.mtu = C.int(mtu)
}

func (v *MplsTunnelSwitch) QosMap() int {
	return int(v.qos_map_id)
}

func (v *MplsTunnelSwitch) SetQosMap(qosMap int) {
	v.qos_map_id = C.int(qosMap)
}

func (v *MplsTunnelSwitch) Failover() Failover {
	return Failover(v.failover_id)
}

func (v *MplsTunnelSwitch) SetFailover(failover Failover) {
	v.failover_id = failover.C()
}

func (v *MplsTunnelSwitch) TunnelID() GPort {
	return GPort(v.tunnel_id)
}
func (v *MplsTunnelSwitch) SetTunnelID(tunnel GPort) {
	v.tunnel_id = tunnel.C()
}

func (v *MplsTunnelSwitch) FailoerTunnel() GPort {
	return GPort(v.failover_tunnel_id)
}
func (v *MplsTunnelSwitch) SetFailoverTunnel(tunnel GPort) {
	v.failover_tunnel_id = tunnel.C()
}

func (v *MplsTunnelSwitch) Tunnel() Iface {
	return Iface(v.tunnel_if)
}

func (v *MplsTunnelSwitch) SetTunnel(iface Iface) {
	v.tunnel_if = iface.C()
}

func (v *MplsTunnelSwitch) EgressPort() GPort {
	return GPort(v.egress_port)
}

func (v *MplsTunnelSwitch) SetEgressPort(port GPort) {
	v.egress_port = port.C()
}

func (v *MplsTunnelSwitch) OAMGlobalContext() uint16 {
	return uint16(v.oam_global_context_id)
}

func (v *MplsTunnelSwitch) SetOAMGlobalContext(ctxt uint16) {
	v.oam_global_context_id = C.uint16(ctxt)
}

func (v *MplsTunnelSwitch) Class() uint32 {
	return uint32(v.class_id)
}

func (v *MplsTunnelSwitch) SetClass(cls uint32) {
	v.class_id = C.uint32(cls)
}

func (v *MplsTunnelSwitch) InLIFCountingProfile() int {
	return int(v.inlif_counting_profile)
}

func (v *MplsTunnelSwitch) SetInLIFCountingProfile(profile int) {
	v.inlif_counting_profile = C.int(profile)
}

//
// API
//

func NewMplsTunnelSwitch() *MplsTunnelSwitch {
	tunsw := &MplsTunnelSwitch{}
	tunsw.Init()
	return tunsw
}

func (v *MplsTunnelSwitch) Init() {
	C.opennsl_mpls_tunnel_switch_t_init(v.C())
}

func (v *MplsTunnelSwitch) Create(unit int) error {
	rc := func() C.int {
		if InfoDeviceIsDNX(unit) {
			return C.opennsl_mpls_tunnel_switch_add(C.int(unit), v.C())
		}
		return C.opennsl_mpls_tunnel_switch_create(C.int(unit), v.C())
	}()

	return ParseError(rc)
}

func (v *MplsTunnelSwitch) Delete(unit int) error {
	rc := C.opennsl_mpls_tunnel_switch_delete(C.int(unit), v.C())
	return ParseError(rc)
}

func MplsTunnelSwitchDeleteAll(unit int) error {
	rc := C.opennsl_mpls_tunnel_switch_delete_all(C.int(unit))
	return ParseError(rc)
}

func (v *MplsTunnelSwitch) Get(unit int) error {
	rc := C.opennsl_mpls_tunnel_switch_get(C.int(unit), v.C())
	return ParseError(rc)
}

type MplsTunnelSwitchTraverseCallback func(int, *MplsTunnelSwitch) int

var mplsTunnelSwitchTraverseCallbacks = NewCallbackMap()

//export go_opennsl_mpls_tunnel_switch_traverse_cb
func go_opennsl_mpls_tunnel_switch_traverse_cb(unit C.int, info *C.opennsl_mpls_tunnel_switch_t, data unsafe.Pointer) int {
	n := (*uint64)(data)
	if h, ok := mplsTunnelSwitchTraverseCallbacks.Get(*n); ok {
		callback := h.(MplsTunnelSwitchTraverseCallback)
		return callback(int(unit), (*MplsTunnelSwitch)(info))
	}

	return int(E_PARAM)
}

func MplsTunnelSwitchTraverse(unit int, callback MplsTunnelSwitchTraverseCallback) error {
	n := mplsTunnelSwitchTraverseCallbacks.Add(callback)
	defer mplsTunnelSwitchTraverseCallbacks.Del(n)

	rc := C.opennsl_mpls_tunnel_switch_traverse(C.int(unit), C.opennsl_mpls_tunnel_switch_traverse_cb(C._opennsl_mpls_tunnel_switch_traverse_cb), unsafe.Pointer(&n))
	return ParseError(rc)
}
