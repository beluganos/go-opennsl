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
#include "helper.h"
*/
import "C"

import (
	"unsafe"
)

//
// L3EgressEcmpID
//
type L3EgressEcmpID C.opennsl_if_t

func (v L3EgressEcmpID) C() C.opennsl_if_t {
	return C.opennsl_if_t(v)
}

func (v L3EgressEcmpID) L3EgressID() L3EgressID {
	return L3EgressID(v)
}

const L3_EGRESS_ECMP = L3EgressEcmpID(IFACE_NONE)

//
// L3EgressEcmp
//
type L3EgressEcmp C.opennsl_l3_egress_ecmp_t

func (v *L3EgressEcmp) C() *C.opennsl_l3_egress_ecmp_t {
	return (*C.opennsl_l3_egress_ecmp_t)(v)
}

func (v *L3EgressEcmp) Flags() L3Flags {
	return L3Flags(v.flags)
}

func (v *L3EgressEcmp) SetFlags(flags L3Flags) {
	v.flags = flags.C()
}

func (v *L3EgressEcmp) EgressEcmp() L3EgressEcmpID {
	return L3EgressEcmpID(v.ecmp_intf)
}

func (v *L3EgressEcmp) SetEgressEcmp(l3egEcmp L3EgressEcmpID) {
	v.ecmp_intf = l3egEcmp.C()
}

func (v *L3EgressEcmp) MaxPaths() int {
	return int(v.max_paths)
}

func (v *L3EgressEcmp) SetMaxPaths(maxPaths int) {
	v.max_paths = C.int(maxPaths)
}

func (v *L3EgressEcmp) DynamicMode() L3EcmpDynamicMode {
	return L3EcmpDynamicMode(v.dynamic_mode)
}

func (v *L3EgressEcmp) SetDynamicMode(mode L3EcmpDynamicMode) {
	v.dynamic_mode = mode.C()
}

func (v *L3EgressEcmp) DynamicSize() uint32 {
	return uint32(v.dynamic_size)
}

func (v *L3EgressEcmp) SetDynamicSize(size uint32) {
	v.dynamic_size = C.uint32(size)
}

//
// L3EcmpMember
//
type L3EcmpMember C.opennsl_l3_ecmp_member_t

func (v *L3EcmpMember) C() *C.opennsl_l3_ecmp_member_t {
	return (*C.opennsl_l3_ecmp_member_t)(v)
}

func (v *L3EcmpMember) Flags() L3Flags {
	return L3Flags(v.flags)
}

func (v *L3EcmpMember) SetFlags(flags L3Flags) {
	v.flags = flags.C()
}

func (v *L3EcmpMember) Egress() L3EgressID {
	return L3EgressID(v.egress_if)
}

func (v *L3EcmpMember) SetEgress(l3eg L3EgressID) {
	v.egress_if = l3eg.C()
}

func (v *L3EcmpMember) Status() int {
	return int(v.status)
}

func (v *L3EcmpMember) SetStatus(status int) {
	v.status = C.int(status)
}

//
// API
//
func NewL3EgressEcmp() *L3EgressEcmp {
	ecmp := &L3EgressEcmp{}
	ecmp.Init()
	return ecmp
}

func L3EgressEcmpInit(v *L3EgressEcmp) {
	C.opennsl_l3_egress_ecmp_t_init(v.C())
}

func (v *L3EgressEcmp) Init() {
	L3EgressEcmpInit(v)
}

func L3EgressEcmpMembersCount(unit int, v *L3EgressEcmp) (int, error) {
	cnt := C.int(0)
	rc := C.opennsl_l3_ecmp_get(C.int(unit), v.C(), 0, nil, &cnt)
	return int(cnt), ParseError(rc)
}

func (v *L3EgressEcmp) MembersCount(unit int) (int, error) {
	return L3EgressEcmpMembersCount(unit, v)
}

func L3EgressEcmpMembers(unit int, v *L3EgressEcmp, size int) ([]L3EcmpMember, error) {
	if size <= 0 {
		cnt, err := v.MembersCount(unit)
		if err != nil {
			return nil, err
		}
		size = cnt
	}

	member_cnt := C.int(0)
	member_arr := make([]C.opennsl_l3_ecmp_member_t, size)
	rc := C.opennsl_l3_ecmp_get(C.int(unit), v.C(), C.int(size), &member_arr[0], &member_cnt)
	if err := ParseError(rc); err != nil {
		return nil, err
	}

	members := make([]L3EcmpMember, int(member_cnt))
	for index := 0; index < int(member_cnt); index++ {
		members[index] = L3EcmpMember(member_arr[index])
	}

	return members, nil
}

func (v *L3EgressEcmp) Members(unit int, size int) ([]L3EcmpMember, error) {
	return L3EgressEcmpMembers(unit, v, size)
}

func L3EgressEcmpCreate(unit int, v *L3EgressEcmp, l3egs []L3EgressID) error {
	member_cnt := len(l3egs)
	member_arr := make([]C.opennsl_if_t, member_cnt)
	for index, l3eg := range l3egs {
		member_arr[index] = l3eg.C()
	}

	rc := C.opennsl_l3_egress_ecmp_create(C.int(unit), v.C(), C.int(member_cnt), &member_arr[0])
	return ParseError(rc)
}

func (v *L3EgressEcmp) Create(unit int, l3egs []L3EgressID) error {
	return L3EgressEcmpCreate(unit, v, l3egs)
}

func L3EgressEcmpDestroy(unit int, v *L3EgressEcmp) error {
	rc := C.opennsl_l3_egress_ecmp_destroy(C.int(unit), v.C())
	return ParseError(rc)
}

func (v *L3EgressEcmp) Destroy(unit int) error {
	return L3EgressEcmpDestroy(unit, v)
}

func L3EgressEcmpCount(unit int, v *L3EgressEcmp) (int, error) {
	cnt := C.int(0)
	rc := C.opennsl_l3_egress_ecmp_get(C.int(unit), v.C(), 0, nil, &cnt)
	return int(cnt), ParseError(rc)
}

func (v *L3EgressEcmp) Count(unit int) (int, error) {
	return L3EgressEcmpCount(unit, v)
}

func L3EgressEcmpGet(unit int, v *L3EgressEcmp, size int) ([]L3EgressID, error) {
	if size == 0 {
		cnt, err := v.Count(unit)
		if err != nil {
			return nil, err
		}
		size = cnt
	}

	iface_cnt := C.int(0)
	iface_arr := make([]C.opennsl_if_t, size)
	rc := C.opennsl_l3_egress_ecmp_get(C.int(unit), v.C(), C.int(size), &iface_arr[0], &iface_cnt)
	if err := ParseError(rc); err != nil {
		return nil, err
	}

	l3egs := make([]L3EgressID, iface_cnt)
	for index := 0; index < int(iface_cnt); index++ {
		l3egs[index] = L3EgressID(iface_arr[index])
	}

	return l3egs, nil
}

func (v *L3EgressEcmp) Get(unit int, size int) ([]L3EgressID, error) {
	return L3EgressEcmpGet(unit, v, size)
}

func L3EgressEcmpAdd(unit int, v *L3EgressEcmp, l3eg L3EgressID) error {
	rc := C.opennsl_l3_egress_ecmp_add(C.int(unit), v.C(), l3eg.C())
	return ParseError(rc)
}

func (v *L3EgressEcmp) Add(unit int, l3eg L3EgressID) error {
	return L3EgressEcmpAdd(unit, v, l3eg)
}

func L3EgressEcmpDelete(unit int, v *L3EgressEcmp, l3eg L3EgressID) error {
	rc := C.opennsl_l3_egress_ecmp_delete(C.int(unit), v.C(), l3eg.C())
	return ParseError(rc)
}

func (v *L3EgressEcmp) Delete(unit int, l3eg L3EgressID) error {
	return L3EgressEcmpDelete(unit, v, l3eg)
}

func L3EgressEcmpFind(unit int, v *L3EgressEcmp, l3egs []L3EgressID) error {
	l3egs_cnt := len(l3egs)
	l3egs_arr := make([]C.opennsl_if_t, l3egs_cnt)

	for index, l3eg := range l3egs {
		l3egs_arr[index] = l3eg.C()
	}

	rc := C.opennsl_l3_egress_ecmp_find(C.int(unit), C.int(l3egs_cnt), &l3egs_arr[0], v.C())
	return ParseError(rc)
}

func (v *L3EgressEcmp) Find(unit int, l3egs []L3EgressID) error {
	return L3EgressEcmpFind(unit, v, l3egs)
}

type L3EgressEcmpTraverseCallback func(int, L3EgressEcmp, []L3EgressID) int

var l3EgressEcmpTraverseCallbacks = NewCallbackMap()

//export go_opennsl_l3_egress_ecmp_traverse_cb
func go_opennsl_l3_egress_ecmp_traverse_cb(unit int, c_ecmp *C.opennsl_l3_egress_ecmp_t, c_count int, c_ifaces []C.opennsl_if_t, c_data unsafe.Pointer) int {
	count := int(c_count)
	l3egs := make([]L3EgressID, count)
	for index := 0; index < count; index++ {
		l3egs[index] = L3EgressID(c_ifaces[index])
	}

	n := (*uint64)(c_data)
	if h, ok := l3EgressEcmpTraverseCallbacks.Get(*n); ok {
		callback := h.(L3EgressEcmpTraverseCallback)
		return callback(int(unit), L3EgressEcmp(*c_ecmp), l3egs)
	}

	return int(E_PARAM)
}

func L3EgressEcmpTraverse(unit int, callback L3EgressEcmpTraverseCallback) error {
	n := l3EgressEcmpTraverseCallbacks.Add(callback)
	defer l3EgressEcmpTraverseCallbacks.Del(n)

	rc := C.opennsl_l3_egress_ecmp_traverse(C.int(unit), (C.opennsl_l3_egress_ecmp_traverse_cb)(C._opennsl_l3_egress_ecmp_traverse_cb), unsafe.Pointer(&n))
	return ParseError(rc)
}
