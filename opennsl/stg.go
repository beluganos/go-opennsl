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
#include <opennsl/stg.h>
#include "helper.h"
*/
import "C"

import (
	"unsafe"
)

//
// opennsl_stg_t
//
type Stg C.opennsl_stg_t

func (v Stg) C() C.opennsl_stg_t {
	return C.opennsl_stg_t(v)
}

//
// StgStp
//
type StgStp C.opennsl_stg_stp_t

func (v StgStp) C() C.opennsl_stg_stp_t {
	return C.opennsl_stg_stp_t(v)
}

const (
	STG_STP_DISABLE StgStp = C.OPENNSL_STG_STP_DISABLE
	STG_STP_BLOCK   StgStp = C.OPENNSL_STG_STP_BLOCK
	STG_STP_LISTEN  StgStp = C.OPENNSL_STG_STP_LISTEN
	STG_STP_LEARN   StgStp = C.OPENNSL_STG_STP_LEARN
	STG_STP_FORWARD StgStp = C.OPENNSL_STG_STP_FORWARD
	STG_STP_COUNT   StgStp = C.OPENNSL_STG_STP_COUNT
)

func StgInit(unit int) error {
	rc := C.opennsl_stg_init(C.int(unit))
	return ParseError(rc)
}

func StgClear(unit int) error {
	rc := C.opennsl_stg_clear(C.int(unit))
	return ParseError(rc)
}

func (v Stg) DefaultSet(unit int) error {
	rc := C.opennsl_stg_default_set(C.int(unit), v.C())
	return ParseError(rc)
}

func StpDefaultGet(unit int) (Stg, error) {
	c_stg := C.opennsl_stg_t(0)

	rc := C.opennsl_stg_default_get(C.int(unit), &c_stg)
	return Stg(c_stg), ParseError(rc)
}

func (v Stg) VlanAdd(unit int, vid Vlan) error {
	rc := C.opennsl_stg_vlan_add(C.int(unit), v.C(), vid.C())
	return ParseError(rc)
}

func (v Stg) VlanRemove(unit int, vid Vlan) error {
	rc := C.opennsl_stg_vlan_remove(C.int(unit), v.C(), vid.C())
	return ParseError(rc)
}

func (v Stg) VlanRemoveAll(unit int) error {
	rc := C.opennsl_stg_vlan_remove_all(C.int(unit), v.C())
	return ParseError(rc)
}

type StgVlanListCallback func(int, Stg, Vlan) int

//export go_opennsl_stg_vlan_list_cb
func go_opennsl_stg_vlan_list_cb(unit C.int, stg C.opennsl_stg_t, vlan C.opennsl_vlan_t, data unsafe.Pointer) int {
	callback := (*StgVlanListCallback)(data)
	return (*callback)(int(unit), Stg(stg), Vlan(vlan))
}

func (v Stg) VlanList(unit int, callback StgVlanListCallback) error {
	rc := C._opennsl_stg_vlan_list_iter(C.int(unit), v.C(), unsafe.Pointer(&callback))
	return ParseError(rc)
}

func StgCreate(unit int) (Stg, error) {
	c_stg := C.opennsl_stg_t(0)

	rc := C.opennsl_stg_create(C.int(unit), &c_stg)
	return Stg(c_stg), ParseError(rc)
}

func (v Stg) Destroy(unit int) error {
	rc := C.opennsl_stg_destroy(C.int(unit), v.C())
	return ParseError(rc)
}

type StgListCallback func(int, Stg) int

//export go_opennsl_stg_list_cb
func go_opennsl_stg_list_cb(unit C.int, stg C.opennsl_stg_t, data unsafe.Pointer) int {
	callback := (*StgListCallback)(data)
	return (*callback)(int(unit), Stg(stg))
}

func StgList(unit int, callback StgListCallback) error {
	rc := C._opennsl_stg_list_iter(C.int(unit), unsafe.Pointer(&callback))
	return ParseError(rc)
}

func (v Stg) StpSet(unit int, port Port, state StgStp) error {
	rc := C.opennsl_stg_stp_set(C.int(unit), v.C(), port.C(), C.int(state.C()))
	return ParseError(rc)
}

func (v Stg) StpGet(unit int, port Port) (StgStp, error) {
	c_stp := C.int(0)

	rc := C.opennsl_stg_stp_get(C.int(unit), v.C(), port.C(), &c_stp)
	return StgStp(c_stp), ParseError(rc)
}

func StgCountGet(unit int) (int, error) {
	c_count := C.int(0)

	rc := C.opennsl_stg_count_get(C.int(unit), &c_count)
	return int(c_count), ParseError(rc)
}
