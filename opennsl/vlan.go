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
#include <opennsl/vlan.h>
#include "helper.h"
*/
import "C"

import (
	"fmt"
	"strconv"
	"strings"
	"unsafe"
)

//
// VlanPortFlags
//
type VlanPortFlags uint32

func (v VlanPortFlags) C() C.uint32 {
	return C.uint32(v)
}

func NewVlanPortFlag(flags ...VlanPortFlags) VlanPortFlags {
	vv := VlanPortFlags(0)
	for _, flag := range flags {
		vv |= flag
	}
	return vv
}

const (
	VLAN_PORT_NONE                VlanPortFlags = 0
	VLAN_PORT_REPLACE             VlanPortFlags = C.OPENNSL_VLAN_PORT_REPLACE
	VLAN_PORT_WITH_ID             VlanPortFlags = C.OPENNSL_VLAN_PORT_WITH_ID
	VLAN_PORT_INNER_VLAN_PRESERVE VlanPortFlags = C.OPENNSL_VLAN_PORT_INNER_VLAN_PRESERVE
	VLAN_PORT_OUTER_VLAN_PRESERVE VlanPortFlags = C.OPENNSL_VLAN_PORT_OUTER_VLAN_PRESERVE
)

var vlanPortFlags_names = map[VlanPortFlags]string{
	VLAN_PORT_REPLACE:             "REPLACE",
	VLAN_PORT_WITH_ID:             "WITH_ID",
	VLAN_PORT_INNER_VLAN_PRESERVE: "INNER_VLAN_PRESERVE",
	VLAN_PORT_OUTER_VLAN_PRESERVE: "OUTER_VLAN_PRESERVE",
}

var vlanPortFlags_values = map[string]VlanPortFlags{
	"REPLACE":             VLAN_PORT_REPLACE,
	"WITH_ID":             VLAN_PORT_WITH_ID,
	"INNER_VLAN_PRESERVE": VLAN_PORT_INNER_VLAN_PRESERVE,
	"OUTER_VLAN_PRESERVE": VLAN_PORT_OUTER_VLAN_PRESERVE,
}

func (v VlanPortFlags) String() string {
	names := []string{}
	for value, name := range vlanPortFlags_names {
		if v&value != 0 {
			names = append(names, name)
		}
	}
	return strings.Join(names, "|")
}

func ParseVlanPortFlags(s string) (VlanPortFlags, error) {
	if v, ok := vlanPortFlags_values[s]; ok {
		return v, nil
	}
	return VLAN_PORT_NONE, fmt.Errorf("Invalid VlanPortFlags. %s", s)
}

//
// VlanPortMatch
//
type VlanPortMatch C.opennsl_vlan_port_match_t

func (v VlanPortMatch) C() C.opennsl_vlan_port_match_t {
	return C.opennsl_vlan_port_match_t(v)
}

const (
	VLAN_PORT_MATCH_INVALID   VlanPortMatch = C.OPENNSL_VLAN_PORT_MATCH_INVALID
	VLAN_PORT_MATCH_NONE      VlanPortMatch = C.OPENNSL_VLAN_PORT_MATCH_NONE
	VLAN_PORT_MATCH_PORT      VlanPortMatch = C.OPENNSL_VLAN_PORT_MATCH_PORT
	VLAN_PORT_MATCH_PORT_VLAN VlanPortMatch = C.OPENNSL_VLAN_PORT_MATCH_PORT_VLAN
)

var vlanPortMatch_names = map[VlanPortMatch]string{
	VLAN_PORT_MATCH_INVALID:   "INVALID",
	VLAN_PORT_MATCH_NONE:      "NONE",
	VLAN_PORT_MATCH_PORT:      "PORT",
	VLAN_PORT_MATCH_PORT_VLAN: "PORT_VLAN",
}

var vlanPortMatch_values = map[string]VlanPortMatch{
	"INVALID":   VLAN_PORT_MATCH_INVALID,
	"NONE":      VLAN_PORT_MATCH_NONE,
	"PORT":      VLAN_PORT_MATCH_PORT,
	"PORT_VLAN": VLAN_PORT_MATCH_PORT_VLAN,
}

func (v VlanPortMatch) String() string {
	if s, ok := vlanPortMatch_names[v]; ok {
		return s
	}
	return fmt.Sprintf("VlanPortMatch(%d)", v)
}

func ParseVlanPortMatch(s string) (VlanPortMatch, error) {
	if v, ok := vlanPortMatch_values[s]; ok {
		return v, nil
	}
	return VLAN_PORT_MATCH_INVALID, fmt.Errorf("Invalid VlanPortMatch. %s", s)
}

//
// VlanMcastFlood
//
type VlanMcastFlood C.opennsl_vlan_mcast_flood_t

func (v VlanMcastFlood) C() C.opennsl_vlan_mcast_flood_t {
	return C.opennsl_vlan_mcast_flood_t(v)
}

const (
	VLAN_MCAST_FLOOD_ALL     VlanMcastFlood = C.OPENNSL_VLAN_MCAST_FLOOD_ALL
	VLAN_MCAST_FLOOD_UNKNOWN VlanMcastFlood = C.OPENNSL_VLAN_MCAST_FLOOD_UNKNOWN
	VLAN_MCAST_FLOOD_NONE    VlanMcastFlood = C.OPENNSL_VLAN_MCAST_FLOOD_NONE
)

const VLAN_MCAST_FLOOD_COUNT int = C.OPENNSL_VLAN_MCAST_FLOOD_COUNT

var vlanMcastFlood_names = map[VlanMcastFlood]string{
	VLAN_MCAST_FLOOD_ALL:     "ALL",
	VLAN_MCAST_FLOOD_UNKNOWN: "UNKNOWN",
	VLAN_MCAST_FLOOD_NONE:    "NONE",
}

var vlanMcastFlood_values = map[string]VlanMcastFlood{
	"ALL":     VLAN_MCAST_FLOOD_ALL,
	"UNKNOWN": VLAN_MCAST_FLOOD_UNKNOWN,
	"NONE":    VLAN_MCAST_FLOOD_NONE,
}

func (v VlanMcastFlood) String() string {
	if s, ok := vlanMcastFlood_names[v]; ok {
		return s
	}
	return fmt.Sprintf("VlanMcastFlood(%d)", v)
}

func ParseVlanMcastFlood(s string) (VlanMcastFlood, error) {
	if v, ok := vlanMcastFlood_values[s]; ok {
		return v, nil
	}
	return VLAN_MCAST_FLOOD_UNKNOWN, fmt.Errorf("Invalid VlanMcastFlood. %s", s)
}

//
// VlanForward
//
type VlanForward C.opennsl_vlan_forward_t

func (v VlanForward) C() C.opennsl_vlan_forward_t {
	return C.opennsl_vlan_forward_t(v)
}

const (
	VlanForwardNone               VlanForward = 0
	VlanForwardBridging           VlanForward = C.opennslVlanForwardBridging
	VlanForwardSingleCrossConnect VlanForward = C.opennslVlanForwardSingleCrossConnect
	VlanForwardDoubleCrossConnect VlanForward = C.opennslVlanForwardDoubleCrossConnect
)

var vlanForward_names = map[VlanForward]string{
	VlanForwardBridging:           "Bridging",
	VlanForwardSingleCrossConnect: "SingleCrossConnect",
	VlanForwardDoubleCrossConnect: "DoubleCrossConnect",
}

var vlanForward_values = map[string]VlanForward{
	"Bridging":           VlanForwardBridging,
	"SingleCrossConnect": VlanForwardSingleCrossConnect,
	"DoubleCrossConnect": VlanForwardDoubleCrossConnect,
}

func (v VlanForward) String() string {
	if s, ok := vlanForward_names[v]; ok {
		return s
	}
	return fmt.Sprintf("VlanForward(%d)", v)
}

func ParseVlanForward(s string) (VlanForward, error) {
	if v, ok := vlanForward_values[s]; ok {
		return v, nil
	}
	return VlanForwardNone, fmt.Errorf("Invalid VlanForward. %s", s)
}

//
// VlanUrpfMode
//
type VlanUrpfMode C.opennsl_vlan_urpf_mode_t

func (v VlanUrpfMode) C() C.opennsl_vlan_urpf_mode_t {
	return C.opennsl_vlan_urpf_mode_t(v)
}

const (
	VlanUrpfDisable VlanUrpfMode = C.opennslVlanUrpfDisable
	VlanUrpfLoose   VlanUrpfMode = C.opennslVlanUrpfLoose
	VlanUrpfStrict  VlanUrpfMode = C.opennslVlanUrpfStrict
)

var vlanUrpfMode_names = map[VlanUrpfMode]string{
	VlanUrpfDisable: "Disable",
	VlanUrpfLoose:   "Loose",
	VlanUrpfStrict:  "Strict",
}

var vlanUrpfMode_values = map[string]VlanUrpfMode{
	"Disable": VlanUrpfDisable,
	"Loose":   VlanUrpfLoose,
	"Strict":  VlanUrpfStrict,
}

func (v VlanUrpfMode) String() string {
	if s, ok := vlanUrpfMode_names[v]; ok {
		return s
	}
	return fmt.Sprintf("VlanUrpfMode(%d)", v)
}

func ParseVlanUrpfMode(s string) (VlanUrpfMode, error) {
	if v, ok := vlanUrpfMode_values[s]; ok {
		return v, nil
	}
	return VlanUrpfDisable, fmt.Errorf("Invaid VlanUrpfMode. %s", s)
}

//
// VlanVPMcControl
//
type VlanVPMcControl C.opennsl_vlan_vp_mc_ctrl_t

func (v VlanVPMcControl) C() C.opennsl_vlan_vp_mc_ctrl_t {
	return C.opennsl_vlan_vp_mc_ctrl_t(v)
}

const (
	VlanVPMcControlAuto    VlanVPMcControl = C.opennslVlanVPMcControlAuto
	VlanVPMcControlEnable  VlanVPMcControl = C.opennslVlanVPMcControlEnable
	VlanVPMcControlDisable VlanVPMcControl = C.opennslVlanVPMcControlDisable
)

var vlanVPMcControl_names = map[VlanVPMcControl]string{
	VlanVPMcControlAuto:    "Auto",
	VlanVPMcControlEnable:  "Enable",
	VlanVPMcControlDisable: "Disable",
}

var vlanVPMcControl_values = map[string]VlanVPMcControl{
	"Auto":    VlanVPMcControlAuto,
	"Enable":  VlanVPMcControlEnable,
	"Disable": VlanVPMcControlDisable,
}

func (v VlanVPMcControl) String() string {
	if s, ok := vlanVPMcControl_names[v]; ok {
		return s
	}
	return fmt.Sprintf("VlanVPMcControl(%d)", v)
}

func ParseVlanVPMcControl(s string) (VlanVPMcControl, error) {
	if v, ok := vlanVPMcControl_values[s]; ok {
		return v, nil
	}
	return VlanVPMcControlAuto, fmt.Errorf("Invalid VlanVPMcControl. %s", s)
}

//
// VlanProtoPkt
//
type VlanProtoPkt int

func (v VlanProtoPkt) C() C.int {
	return C.int(v)
}

func VlanProtoPktToCPU(protoPkt VlanProtoPkt) VlanProtoPkt {
	return VLAN_PROTO_PKT_TOCPU_ENABLE | protoPkt
}

const (
	VLAN_PROTO_PKT_NONE           VlanProtoPkt = 0
	VLAN_PROTO_PKT_TOCPU_ENABLE   VlanProtoPkt = C.OPENNSL_VLAN_PROTO_PKT_TOCPU_ENABLE
	VLAN_PROTO_PKT_FORWARD_ENABLE VlanProtoPkt = C.OPENNSL_VLAN_PROTO_PKT_FORWARD_ENABLE
	VLAN_PROTO_PKT_DROP_ENABLE    VlanProtoPkt = C.OPENNSL_VLAN_PROTO_PKT_DROP_ENABLE
	VLAN_PROTO_PKT_FLOOD_ENABLE   VlanProtoPkt = C.OPENNSL_VLAN_PROTO_PKT_FLOOD_ENABLE
)

var vlanProtoPkt_names = map[VlanProtoPkt]string{
	VLAN_PROTO_PKT_TOCPU_ENABLE:   "TOCPU_ENABLE",
	VLAN_PROTO_PKT_FORWARD_ENABLE: "FORWARD_ENABLE",
	VLAN_PROTO_PKT_DROP_ENABLE:    "DROP_ENABLE",
	VLAN_PROTO_PKT_FLOOD_ENABLE:   "FLOOD_ENABLE",
}

var vlanProtoPkt_values = map[string]VlanProtoPkt{
	"TOCPU_ENABLE":   VLAN_PROTO_PKT_TOCPU_ENABLE,
	"FORWARD_ENABLE": VLAN_PROTO_PKT_FORWARD_ENABLE,
	"DROP_ENABLE":    VLAN_PROTO_PKT_DROP_ENABLE,
	"FLOOD_ENABLE":   VLAN_PROTO_PKT_FLOOD_ENABLE,
}

func (v VlanProtoPkt) String() string {
	if s, ok := vlanProtoPkt_names[v]; ok {
		return s
	}
	return fmt.Sprintf("VlanProtoPkt(%d)", v)
}

func ParseVlanProtoPkt(s string) (VlanProtoPkt, error) {
	if v, ok := vlanProtoPkt_values[s]; ok {
		return v, nil
	}
	return VLAN_PROTO_PKT_NONE, fmt.Errorf("Invalid VlanProtoPkt. %s", s)
}

//
// Vlan learn disable
//
const VLAN_LEARN_DISABLE uint32 = C.OPENNSL_VLAN_LEARN_DISABLE

//
// Vlan
//
type Vlan C.opennsl_vlan_t

const (
	VLAN_ID_NONE    Vlan = C.OPENNSL_VLAN_NONE
	VLAN_ID_DEFAULT Vlan = C.OPENNSL_VLAN_DEFAULT
	VLAN_ID_MIN     Vlan = C.OPENNSL_VLAN_MIN
	VLAN_ID_MAX     Vlan = C.OPENNSL_VLAN_MAX
	VLAN_ID_INVALID Vlan = 0xffff
	VLAN_MASK_NONE  Vlan = 0
	VLAN_MASK_EXACT Vlan = 0x0fff
)

var vlanId_names = map[Vlan]string{
	VLAN_ID_NONE:    "NONE",
	VLAN_ID_DEFAULT: "DEFAULT",
	VLAN_ID_INVALID: "INVALID",
}

var vlanId_values = map[string]Vlan{
	"NONE":    VLAN_ID_NONE,
	"DEFAULT": VLAN_ID_DEFAULT,
	"INVALID": VLAN_ID_INVALID,
}

func (v Vlan) String() string {
	if s, ok := vlanId_names[v]; ok {
		return s
	}
	return fmt.Sprintf("%d", v)
}

func ParseVlan(s string) (Vlan, error) {
	if v, ok := vlanId_values[s]; ok {
		return v, nil
	}

	v, err := strconv.ParseUint(s, 0, 16)

	if err != nil {
		return 0, err
	}

	return Vlan(uint16(v)), nil
}

func (v Vlan) C() C.opennsl_vlan_t {
	return C.opennsl_vlan_t(v)
}

func (v Vlan) Valid() bool {
	return (v >= VLAN_ID_MIN) && (v <= VLAN_ID_MAX)
}

func (v Vlan) Create(unit int) (bool, error) {
	rc := C.opennsl_vlan_create(C.int(unit), v.C())

	if OpenNSLError(rc) == E_EXISTS {
		return false, nil
	}

	if err := ParseError(rc); err != nil {
		return false, err
	}

	return true, nil
}

func (v Vlan) Destroy(unit int) error {
	rc := C.opennsl_vlan_destroy(C.int(unit), v.C())
	return ParseError(rc)
}

func VlanDestroyAll(unit int) error {
	rc := C.opennsl_vlan_destroy_all(C.int(unit))
	return ParseError(rc)
}

func (v Vlan) PortAdd(unit int, pbmp *PBmp, ubmp *PBmp) error {
	rc := C.opennsl_vlan_port_add(C.int(unit), v.C(), *(pbmp.C()), *(ubmp.C()))
	return ParseError(rc)
}

func (v Vlan) PortRemove(unit int, pbmp *PBmp) (bool, error) {
	rc := C.opennsl_vlan_port_remove(C.int(unit), v.C(), *(pbmp.C()))
	if rc == C.int(E_NOT_FOUND) {
		return false, nil
	}
	return (rc == C.int(E_NONE)), ParseError(rc)
}

func (v Vlan) PortGet(unit int) (*PBmp, *PBmp, error) {
	pbmp := PBmp{}
	ubmp := PBmp{}
	pbmp.Clear()
	ubmp.Clear()
	rc := C.opennsl_vlan_port_get(C.int(unit), v.C(), pbmp.C(), ubmp.C())
	if err := ParseError(rc); err != nil {
		return nil, nil, err
	}

	return &pbmp, &ubmp, nil
}

func (v Vlan) GPortAdd(unit int, gport GPort, flags VlanPortFlags) error {
	rc := C.opennsl_vlan_gport_add(C.int(unit), v.C(), gport.C(), C.int(flags.C()))
	return ParseError(rc)
}

func (v Vlan) GPortDelete(unit int, gport GPort) (bool, error) {
	rc := C.opennsl_vlan_gport_delete(C.int(unit), v.C(), gport.C())
	if OpenNSLError(rc) == E_NOT_FOUND {
		return false, nil
	}

	if err := ParseError(rc); err != nil {
		return false, err
	}

	return true, nil
}

func (v Vlan) GPortDeleteAll(unit int) error {
	rc := C.opennsl_vlan_gport_delete_all(C.int(unit), v.C())
	return ParseError(rc)
}

func (v Vlan) Get(unit int, gport GPort) (VlanPortFlags, error) {
	flags := C.int(0)
	rc := C.opennsl_vlan_gport_get(C.int(unit), v.C(), gport.C(), &flags)
	if OpenNSLError(rc) == E_NOT_FOUND {
		return VLAN_PORT_NONE, nil
	}

	if err := ParseError(rc); err != nil {
		return VLAN_PORT_NONE, err
	}

	return VlanPortFlags(flags), nil
}

func VlanDefaultGet(unit int) (Vlan, error) {
	vid := C.opennsl_vlan_t(0)
	rc := C.opennsl_vlan_default_get(C.int(unit), &vid)
	return Vlan(vid), ParseError(rc)
}

func VlanDefaultMustGet(unit int) Vlan {
	if vlan, err := VlanDefaultGet(unit); err == nil {
		return vlan
	}
	return VLAN_ID_DEFAULT
}

func (v Vlan) DefaultSet(unit int) error {
	rc := C.opennsl_vlan_default_set(C.int(unit), v.C())
	return ParseError(rc)
}

//
// VlanData
//
type VlanTraverseHandler func(int, Vlan, *PBmp, *PBmp) OpenNSLError

var vlanTraverseHandlers = NewCallbackMap()

//export go_opennsl_vlan_list_cb
func go_opennsl_vlan_list_cb(unit int, data *C.opennsl_vlan_data_t, c_data unsafe.Pointer) int {
	n := (*uint64)(c_data)
	if h, ok := vlanTraverseHandlers.Get(*n); ok {
		callback := h.(VlanTraverseHandler)
		rc := callback(
			unit,
			Vlan(data.vlan_tag),
			(*PBmp)(&data.port_bitmap),
			(*PBmp)(&data.ut_port_bitmap),
		)
		return int(rc)
	}

	return int(E_PARAM)
}

func VlanTraverse(unit int, handler VlanTraverseHandler) error {
	n := vlanTraverseHandlers.Add(handler)
	defer vlanTraverseHandlers.Del(n)

	rc := C._opennsl_vlan_list_iter(C.int(unit), unsafe.Pointer(&n))
	return ParseError(rc)
}
