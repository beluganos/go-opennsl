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
#include <opennsl/vlan.h>
*/
import "C"

import (
	"fmt"
)

//
// VlanControl
//
type VlanControl C.opennsl_vlan_control_t

func (v VlanControl) C() C.opennsl_vlan_control_t {
	return C.opennsl_vlan_control_t(v)
}

const (
	VlanDropUnknown         VlanControl = C.opennslVlanDropUnknown
	VlanShared              VlanControl = C.opennslVlanShared
	VlanSharedID            VlanControl = C.opennslVlanSharedID
	VlanTranslate           VlanControl = C.opennslVlanTranslate
	VlanIgnorePktTag        VlanControl = C.opennslVlanIgnorePktTag
	VlanMemberMismatchToCpu VlanControl = C.opennslVlanMemberMismatchToCpu
)

func (v VlanControl) Set(unit int, arg int) error {
	rc := C.opennsl_vlan_control_set(C.int(unit), v.C(), C.int(arg))
	return ParseError(rc)
}

var vlanControl_names = map[VlanControl]string{
	VlanDropUnknown:         "DropUnknown",
	VlanShared:              "Shared",
	VlanSharedID:            "SharedID",
	VlanTranslate:           "Translate",
	VlanIgnorePktTag:        "IgnorePktTag",
	VlanMemberMismatchToCpu: "MemberMismatchToCpu",
}

var vlanControl_values = map[string]VlanControl{
	"DropUnknown":         VlanDropUnknown,
	"Shared":              VlanShared,
	"SharedID":            VlanSharedID,
	"Translate":           VlanTranslate,
	"IgnorePktTag":        VlanIgnorePktTag,
	"MemberMismatchToCpu": VlanMemberMismatchToCpu,
}

func (v VlanControl) String() string {
	if s, ok := vlanControl_names[v]; ok {
		return s
	}
	return fmt.Sprintf("VlanControl(%d)", v)
}

func ParseVlanControl(s string) (VlanControl, error) {
	if v, ok := vlanControl_values[s]; ok {
		return v, nil
	}
	return VlanDropUnknown, fmt.Errorf("Invalid VlanControl. %s", s)
}

//
// VlanControlEntry
//
type VlanControlEntry struct {
	Ctrl  VlanControl
	Value int
}

func NewVlanControlEntry(ctrl VlanControl, value int) *VlanControlEntry {
	return &VlanControlEntry{
		Ctrl:  ctrl,
		Value: value,
	}
}

func (v VlanControl) Arg(value int) *VlanControlEntry {
	return NewVlanControlEntry(v, value)
}

func (e *VlanControlEntry) Set(unit int) error {
	return e.Ctrl.Set(unit, e.Value)
}

func VlanControlsSet(unit int, entries ...*VlanControlEntry) error {
	for _, entry := range entries {
		if err := entry.Set(unit); err != nil {
			return err
		}
	}
	return nil
}

//
// VlanControlPort
//
type VlanControlPort C.opennsl_vlan_control_port_t

func (v VlanControlPort) C() C.opennsl_vlan_control_port_t {
	return C.opennsl_vlan_control_port_t(v)
}

const (
	VlanTranslateNone            VlanControlPort = 0
	VlanTranslateIngressEnable   VlanControlPort = C.opennslVlanTranslateIngressEnable
	VlanTranslateIngressMissDrop VlanControlPort = C.opennslVlanTranslateIngressMissDrop
	VlanTranslateEgressEnable    VlanControlPort = C.opennslVlanTranslateEgressEnable
	VlanTranslateEgressMissDrop  VlanControlPort = C.opennslVlanTranslateEgressMissDrop
)

func (v VlanControlPort) Set(unit int, port Port, arg int) error {
	rc := C.opennsl_vlan_control_port_set(C.int(unit), port.C(), v.C(), C.int(arg))
	return ParseError(rc)
}

var vlanControlPort_names = map[VlanControlPort]string{
	VlanTranslateIngressEnable:   "IngressEnable",
	VlanTranslateIngressMissDrop: "IngressMissDrop",
	VlanTranslateEgressEnable:    "EgressEnable",
	VlanTranslateEgressMissDrop:  "EgressMissDrop",
}

var vlanControlPort_values = map[string]VlanControlPort{
	"IngressEnable":   VlanTranslateIngressEnable,
	"IngressMissDrop": VlanTranslateIngressMissDrop,
	"EgressEnable":    VlanTranslateEgressEnable,
	"EgressMissDrop":  VlanTranslateEgressMissDrop,
}

func (v VlanControlPort) String() string {
	if s, ok := vlanControlPort_names[v]; ok {
		return s
	}
	return fmt.Sprintf("VlanControlPort(%d)", v)
}

func ParseVlanControlPort(s string) (VlanControlPort, error) {
	if v, ok := vlanControlPort_values[s]; ok {
		return v, nil
	}
	return VlanTranslateNone, fmt.Errorf("Invalid VlanControlPort. %s", s)
}
