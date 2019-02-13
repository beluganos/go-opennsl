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
*/
import "C"

import (
	"fmt"
	"strings"
)

//
// MplsEgressLabelFlags
//
type MplsEgressLabelFlags C.uint32

func (v MplsEgressLabelFlags) C() C.uint32 {
	return C.uint32(v)
}

func NewMplsEgressLabelFlags(flags ...MplsEgressLabelFlags) MplsEgressLabelFlags {
	v := MplsEgressLabelFlags(0)
	for _, flag := range flags {
		v |= flag
	}
	return v
}

const (
	MPLS_EGRESS_LABEL_NONE          MplsEgressLabelFlags = 0
	MPLS_EGRESS_LABEL_TTL_SET       MplsEgressLabelFlags = C.OPENNSL_MPLS_EGRESS_LABEL_TTL_SET
	MPLS_EGRESS_LABEL_TTL_COPY      MplsEgressLabelFlags = C.OPENNSL_MPLS_EGRESS_LABEL_TTL_COPY
	MPLS_EGRESS_LABEL_TTL_DECREMENT MplsEgressLabelFlags = C.OPENNSL_MPLS_EGRESS_LABEL_TTL_DECREMENT
	MPLS_EGRESS_LABEL_ACTION_VALID  MplsEgressLabelFlags = C.OPENNSL_MPLS_EGRESS_LABEL_ACTION_VALID
	MPLS_EGRESS_LABEL_REPLACE       MplsEgressLabelFlags = C.OPENNSL_MPLS_EGRESS_LABEL_REPLACE
	MPLS_EGRESS_LABEL_WITH_ID       MplsEgressLabelFlags = C.OPENNSL_MPLS_EGRESS_LABEL_WITH_ID
)

var mplsEgressLabelFlags_names = map[MplsEgressLabelFlags]string{
	MPLS_EGRESS_LABEL_NONE:          "NONE",
	MPLS_EGRESS_LABEL_TTL_SET:       "TTL_SET",
	MPLS_EGRESS_LABEL_TTL_COPY:      "TTL_COPY",
	MPLS_EGRESS_LABEL_TTL_DECREMENT: "TTL_DECREMENT",
	MPLS_EGRESS_LABEL_ACTION_VALID:  "ACTION_VALID",
	MPLS_EGRESS_LABEL_REPLACE:       "REPLACE",
	MPLS_EGRESS_LABEL_WITH_ID:       "WITH_ID",
}

var mplsEgressLabelFlags_values = map[string]MplsEgressLabelFlags{
	"TTL_SET":       MPLS_EGRESS_LABEL_TTL_SET,
	"TTL_COPY":      MPLS_EGRESS_LABEL_TTL_COPY,
	"TTL_DECREMENT": MPLS_EGRESS_LABEL_TTL_DECREMENT,
	"ACTION_VALID":  MPLS_EGRESS_LABEL_ACTION_VALID,
	"REPLACE":       MPLS_EGRESS_LABEL_REPLACE,
	"WITH_ID":       MPLS_EGRESS_LABEL_WITH_ID,
}

func (v MplsEgressLabelFlags) String() string {
	names := make([]string, 0, len(mplsEgressLabelFlags_names))
	for val, name := range mplsEgressLabelFlags_names {
		if v&val != 0 {
			names = append(names, name)
		}
	}
	return strings.Join(names, "|")
}

func ParseMplsEgressLabelFlags(s string) (MplsEgressLabelFlags, error) {
	if v, ok := mplsEgressLabelFlags_values[s]; ok {
		return v, nil
	}
	return MPLS_EGRESS_LABEL_NONE, fmt.Errorf("Invalid MplsEgressLabelFlags. %s", s)
}

//
// MplsEgressAction
//
type MplsEgressAction C.opennsl_mpls_egress_action_t

func (v MplsEgressAction) C() C.opennsl_mpls_egress_action_t {
	return C.opennsl_mpls_egress_action_t(v)
}

const (
	MPLS_EGRESS_ACTION_SWAP         MplsEgressAction = C.OPENNSL_MPLS_EGRESS_ACTION_SWAP
	MPLS_EGRESS_ACTION_PHP          MplsEgressAction = C.OPENNSL_MPLS_EGRESS_ACTION_PHP
	MPLS_EGRESS_ACTION_PUSH         MplsEgressAction = C.OPENNSL_MPLS_EGRESS_ACTION_PUSH
	MPLS_EGRESS_ACTION_NOP          MplsEgressAction = C.OPENNSL_MPLS_EGRESS_ACTION_NOP
	MPLS_EGRESS_ACTION_SWAP_OR_PUSH MplsEgressAction = C.OPENNSL_MPLS_EGRESS_ACTION_SWAP_OR_PUSH
)

var mplsEgressAction_names = map[MplsEgressAction]string{
	MPLS_EGRESS_ACTION_SWAP:         "SWAP",
	MPLS_EGRESS_ACTION_PHP:          "PHP",
	MPLS_EGRESS_ACTION_PUSH:         "PUSH",
	MPLS_EGRESS_ACTION_NOP:          "NOP",
	MPLS_EGRESS_ACTION_SWAP_OR_PUSH: "SWAP_OR_PUSH",
}

var mplsEgressAction_values = map[string]MplsEgressAction{
	"SWAP":         MPLS_EGRESS_ACTION_SWAP,
	"PHP":          MPLS_EGRESS_ACTION_PHP,
	"PUSH":         MPLS_EGRESS_ACTION_PUSH,
	"NOP":          MPLS_EGRESS_ACTION_NOP,
	"SWAP_OR_PUSH": MPLS_EGRESS_ACTION_SWAP_OR_PUSH,
}

func (v MplsEgressAction) String() string {
	if s, ok := mplsEgressAction_names[v]; ok {
		return s
	}
	return fmt.Sprintf("MplsEgressAction(%d)", v)
}

func ParseMplsEgressAction(s string) (MplsEgressAction, error) {
	if v, ok := mplsEgressAction_values[s]; ok {
		return v, nil
	}
	return MPLS_EGRESS_ACTION_NOP, fmt.Errorf("Invalid MplsEgressAction. %s", s)
}

//
// MplsEgressLabel
//
type MplsEgressLabel C.opennsl_mpls_egress_label_t

func (v *MplsEgressLabel) C() *C.opennsl_mpls_egress_label_t {
	return (*C.opennsl_mpls_egress_label_t)(v)
}

func (v *MplsEgressLabel) Flags() MplsEgressLabelFlags {
	return MplsEgressLabelFlags(v.flags)
}

func (v *MplsEgressLabel) SetFlags(flags MplsEgressLabelFlags) {
	v.flags = flags.C()
}

func (v *MplsEgressLabel) Label() MplsLabel {
	return MplsLabel(v.label)
}

func (v *MplsEgressLabel) SetLabel(label MplsLabel) {
	v.label = label.C()
}

func (v *MplsEgressLabel) QosMap() int {
	return int(v.qos_map_id)
}

func (v *MplsEgressLabel) SetQosMap(qosMap int) {
	v.qos_map_id = C.int(qosMap)
}

func (v *MplsEgressLabel) Exp() uint8 {
	return uint8(v.exp)
}

func (v *MplsEgressLabel) SetExp(exp uint8) {
	v.exp = C.uint8(exp)
}

func (v *MplsEgressLabel) TTL() uint8 {
	return uint8(v.ttl)
}

func (v *MplsEgressLabel) SetTTL(ttl uint8) {
	v.ttl = C.uint8(ttl)
}

func (v *MplsEgressLabel) PktPri() uint8 {
	return uint8(v.pkt_pri)
}

func (v *MplsEgressLabel) SetPktPri(pri uint8) {
	v.pkt_pri = C.uint8(pri)
}

func (v *MplsEgressLabel) PktCfi() uint8 {
	return uint8(v.pkt_cfi)
}

func (v *MplsEgressLabel) SetPktCfi(cfi uint8) {
	v.pkt_cfi = C.uint8(cfi)
}

func (v *MplsEgressLabel) Tunnel() Iface {
	return Iface(v.tunnel_id)
}

func (v *MplsEgressLabel) SetTunnel(iface Iface) {
	v.tunnel_id = iface.C()
}

func (v *MplsEgressLabel) L3Iface() L3IfaceID {
	return L3IfaceID(v.l3_intf_id)
}

func (v *MplsEgressLabel) SetL3Iface(iface L3IfaceID) {
	v.l3_intf_id = iface.C()
}

func (v *MplsEgressLabel) Action() MplsEgressAction {
	return MplsEgressAction(v.action)
}

func (v *MplsEgressLabel) SetAction(action MplsEgressAction) {
	v.action = action.C()
}

func (v *MplsEgressLabel) EgressFailover() Failover {
	return Failover(v.egress_failover_id)
}

func (v *MplsEgressLabel) SetEgressFailover(failover Failover) {
	v.egress_failover_id = failover.C()
}

func (v *MplsEgressLabel) EgressFailoverIface() Iface {
	return Iface(v.egress_failover_if_id)
}

func (v *MplsEgressLabel) SetEgressFailoverIFace(failover Iface) {
	v.egress_failover_if_id = failover.C()
}

func (v *MplsEgressLabel) OutLIFCountingProfile() int {
	return int(v.outlif_counting_profile)
}

func (v *MplsEgressLabel) SetOutLIFCountingProfile(profile int) {
	v.outlif_counting_profile = C.int(profile)
}

//
// API
//

func NewMplsEgressLabel() *MplsEgressLabel {
	label := &MplsEgressLabel{}
	label.Init()
	return label
}

func (v *MplsEgressLabel) Init() {
	C.opennsl_mpls_egress_label_t_init(v.C())
}
