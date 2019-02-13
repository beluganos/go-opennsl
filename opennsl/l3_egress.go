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
	"fmt"
	"net"
	"unsafe"
)

//
// L3EgressID
//
type L3EgressID C.opennsl_if_t

func (v L3EgressID) C() C.opennsl_if_t {
	return C.opennsl_if_t(v)
}

const (
	L3_EGRESS L3EgressID = L3EgressID(IFACE_NONE)
)

//
// L3Egress
//
type L3Egress C.opennsl_l3_egress_t

func (v *L3Egress) C() *C.opennsl_l3_egress_t {
	return (*C.opennsl_l3_egress_t)(v)
}

func (v *L3Egress) Flags() L3Flags {
	return L3Flags(v.flags)
}

func (v *L3Egress) Flags2() L3Flags {
	return L3Flags(v.flags2)
}

func (v *L3Egress) SetFlags(flags L3Flags) {
	v.flags = flags.C()
}

func (v *L3Egress) IfaceID() L3IfaceID {
	return L3IfaceID(v.intf)
}

func (v *L3Egress) SetIfaceID(iface L3IfaceID) {
	v.intf = iface.C()
}

func (v *L3Egress) MAC() net.HardwareAddr {
	return ParseMAC(v.mac_addr)
}

func (v *L3Egress) SetMAC(mac net.HardwareAddr) {
	v.mac_addr = NewMAC(mac)
}

func (v *L3Egress) VID() Vlan {
	return Vlan(v.vlan)
}

func (v *L3Egress) SetVID(vid Vlan) {
	v.vlan = vid.C()
}

func (v *L3Egress) Module() Module {
	return Module(v.module)
}

func (v *L3Egress) SetModule(module Module) {
	v.module = module.C()
}

func (v *L3Egress) Port() Port {
	return Port(v.port)
}

func (v *L3Egress) SetPort(port Port) {
	v.port = port.C()
}

func (v *L3Egress) Trunk() Trunk {
	return Trunk(v.trunk)
}

func (v *L3Egress) SetTrunk(trunk Trunk) {
	v.trunk = trunk.C()
}

//
// API
//
func NewL3Egress() *L3Egress {
	l3eg := &L3Egress{}
	l3eg.Init()
	return l3eg
}

func L3EgressInit(v *L3Egress) {
	C.opennsl_l3_egress_t_init(v.C())
}

func (v *L3Egress) String() string {
	return fmt.Sprintf("L3Egress(mac:%s, vlan:%d, port:%d, intf:%d)", v.MAC(), v.VID(), v.Port(), v.IfaceID())
}

func (v *L3Egress) Init() {
	L3EgressInit(v)
}

func L3EgressCreate(unit int, flags L3Flags, v *L3Egress, l3egID L3EgressID) (L3EgressID, error) {
	c_id := l3egID.C()
	rc := C.opennsl_l3_egress_create(C.int(unit), flags.C(), v.C(), &c_id)
	return L3EgressID(c_id), ParseError(rc)
}

func (v *L3Egress) Create(unit int, flags L3Flags, l3egID L3EgressID) (L3EgressID, error) {
	return L3EgressCreate(unit, flags, v, l3egID)
}

func (v L3EgressID) Create(unit int, flags L3Flags, l3eg *L3Egress) (L3EgressID, error) {
	return L3EgressCreate(unit, flags, l3eg, v)
}

func L3EgressDestroy(unit int, v L3EgressID) error {
	rc := C.opennsl_l3_egress_destroy(C.int(unit), v.C())
	return ParseError(rc)
}

func (v L3EgressID) Destroy(unit int) error {
	return L3EgressDestroy(unit, v)
}

func L3EgressGet(unit int, v L3EgressID) (*L3Egress, error) {
	egress := L3Egress{}
	egress.Init()

	rc := C.opennsl_l3_egress_get(C.int(unit), v.C(), egress.C())
	if OpenNSLError(rc) == E_NOT_FOUND {
		return nil, nil
	}

	if err := ParseError(rc); err != nil {
		return nil, err
	}

	return &egress, nil
}

func (v L3EgressID) Get(unit int) (*L3Egress, error) {
	return L3EgressGet(unit, v)
}

func L3EgressFind(unit int, v *L3Egress) (L3EgressID, error) {
	c_id := C.opennsl_if_t(0)
	rc := C.opennsl_l3_egress_find(C.int(unit), v.C(), &c_id)
	return L3EgressID(c_id), ParseError(rc)
}

func (v *L3Egress) Find(unit int) (L3EgressID, error) {
	return L3EgressFind(unit, v)
}

type L3EgressTraverseCallback func(int, L3EgressID, *L3Egress) OpenNSLError

var l3EgressTraverseCallbacks = NewCallbackMap()

//export go_opennsl_l3_egress_traverse_cb
func go_opennsl_l3_egress_traverse_cb(c_unit C.int, c_iface C.opennsl_if_t, c_egr *C.opennsl_l3_egress_t, c_data unsafe.Pointer) int {
	n := (*uint64)(c_data)
	if h, ok := l3EgressTraverseCallbacks.Get(*n); ok {
		callback := h.(L3EgressTraverseCallback)
		rc := callback(int(c_unit), L3EgressID(c_iface), (*L3Egress)(c_egr))
		return int(rc)
	}

	return int(E_PARAM)
}

func L3EgressTraverse(unit int, callback L3EgressTraverseCallback) error {
	n := l3EgressTraverseCallbacks.Add(callback)
	defer l3EgressTraverseCallbacks.Del(n)

	rc := C.opennsl_l3_egress_traverse(C.int(unit), C.opennsl_l3_egress_traverse_cb(C._opennsl_l3_egress_traverse_cb), unsafe.Pointer(&n))
	return ParseError(rc)
}
