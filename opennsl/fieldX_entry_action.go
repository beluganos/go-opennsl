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
#include <opennsl/fieldX.h>
*/
import "C"

import (
	"net"
)

//
// FieldEntryAction
//
type FieldEntryAction C.opennsl_field_entry_t

func (v FieldEntryAction) C() C.opennsl_field_entry_t {
	return C.opennsl_field_entry_t(v)
}

func (v FieldEntryAction) Add(unit int, action FieldAction, param0, param1 uint32) error {
	rc := C.opennsl_field_action_add(C.int(unit), v.C(), action.C(), C.uint32(param0), C.uint32(param1))
	return ParseError(rc)
}

func (v FieldEntryAction) AddP(unit int, p FieldActionInterface) error {
	return v.Add(unit, p.Action(), p.Param0(), p.Param1())
}

func (v FieldEntryAction) Delete(unit int, action FieldAction, param0, param1 uint32) error {
	rc := C.opennsl_field_action_delete(C.int(unit), v.C(), action.C(), C.uint32(param0), C.uint32(param1))
	return ParseError(rc)
}

func (v FieldEntryAction) DeleteP(unit int, p FieldActionInterface) error {
	return v.Delete(unit, p.Action(), p.Param0(), p.Param1())
}

func (v FieldEntryAction) MACAdd(unit int, action FieldAction, mac net.HardwareAddr) error {
	c_mac := NewMAC(mac)

	rc := C.opennsl_field_action_mac_add(C.int(unit), v.C(), action.C(), &c_mac[0])
	return ParseError(rc)
}

func (v FieldEntryAction) PortsAdd(unit int, action FieldAction, ports *PBmp) error {
	rc := C.opennsl_field_action_ports_add(C.int(unit), v.C(), action.C(), *ports.C())
	return ParseError(rc)
}

func (v FieldEntryAction) Get(unit int, action FieldAction) (uint32, uint32, error) {
	c_param0 := C.uint32(0)
	c_param1 := C.uint32(0)

	rc := C.opennsl_field_action_get(C.int(unit), v.C(), action.C(), &c_param0, &c_param1)
	return uint32(c_param0), uint32(c_param1), ParseError(rc)
}

func (v FieldEntryAction) MACGet(unit int, action FieldAction) (net.HardwareAddr, error) {
	c_mac := C.opennsl_mac_t{}

	rc := C.opennsl_field_action_mac_get(C.int(unit), v.C(), action.C(), &c_mac)
	if err := ParseError(rc); err != nil {
		return nil, err
	}

	return ParseMAC(c_mac), nil
}

func (v FieldEntryAction) PortsGet(unit int, action FieldAction) (*PBmp, error) {
	pbmp := PBmp{}

	rc := C.opennsl_field_action_ports_get(C.int(unit), v.C(), action.C(), pbmp.C())
	if err := ParseError(rc); err != nil {
		return nil, err
	}

	return &pbmp, nil
}

func (v FieldEntryAction) Remove(unit int, action FieldAction) error {
	rc := C.opennsl_field_action_remove(C.int(unit), v.C(), action.C())
	return ParseError(rc)
}

func (v FieldEntryAction) RemoveAll(unit int) error {
	rc := C.opennsl_field_action_remove_all(C.int(unit), v.C())
	return ParseError(rc)
}
