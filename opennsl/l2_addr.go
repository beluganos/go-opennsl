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
#include <opennsl/l2.h>
#include <helper.h>
*/
import "C"

import (
	"fmt"
	"net"
	"unsafe"
)

//
// L2Addr
//
type L2Addr C.opennsl_l2_addr_t

func (v *L2Addr) C() *C.opennsl_l2_addr_t {
	return (*C.opennsl_l2_addr_t)(v)
}

func (v *L2Addr) Flags() L2Flags {
	return L2Flags(v.flags)
}

func (v *L2Addr) SetFlags(flags L2Flags) {
	v.flags = flags.C()
}

func (v *L2Addr) MAC() net.HardwareAddr {
	return ParseMAC(v.mac)
}

func (v *L2Addr) SetMAC(hwaddr net.HardwareAddr) {
	v.mac = NewMAC(hwaddr)
}

func (v *L2Addr) VID() Vlan {
	return Vlan(v.vid)
}

func (v *L2Addr) SetVID(vid Vlan) {
	v.vid = vid.C()
}

func (v *L2Addr) Port() Port {
	return Port(v.port)
}

func (v *L2Addr) SetPort(port Port) {
	v.port = port.C()
}

func (v *L2Addr) GPort() GPort {
	return GPort(v.port)
}

func (v *L2Addr) SetGPort(gport GPort) {
	v.port = gport.C()
}

func (v *L2Addr) TrunkID() Trunk {
	return Trunk(v.tgid)
}

func (v *L2Addr) SetTrunkID(tid Trunk) {
	v.tgid = tid.C()
}

func (v *L2Addr) L2MCGroup() Multicast {
	return Multicast(v.l2mc_group)
}

func (v *L2Addr) SetL2MCGroup(l2mc_group Multicast) {
	v.l2mc_group = l2mc_group.C()
}

//
// API
//
func NewL2Addr(hwaddr net.HardwareAddr, vid Vlan) *L2Addr {
	addr := &L2Addr{}
	addr.Init(hwaddr, vid)
	return addr
}

func (v *L2Addr) String() string {
	return fmt.Sprintf("L2Addr(mac:%s, vlan:%d)", v.MAC(), v.VID())
}

func (v *L2Addr) Init(hwaddr net.HardwareAddr, vid Vlan) {
	mac := NewMAC(hwaddr)
	C.opennsl_l2_addr_t_init(v.C(), &mac[0], vid.C())
}

func (v *L2Addr) Add(unit int) error {
	rc := C.opennsl_l2_addr_add(C.int(unit), v.C())
	return ParseError(rc)
}

func L2AddrDelete(unit int, hwaddr net.HardwareAddr, vid Vlan) error {
	mac := NewMAC(hwaddr)
	rc := C.opennsl_l2_addr_delete(C.int(unit), &mac[0], vid.C())
	return ParseError(rc)
}

func (v *L2Addr) Delete(unit int) error {
	rc := C.opennsl_l2_addr_delete(C.int(unit), &v.mac[0], v.vid)
	return ParseError(rc)
}

func L2AddrDeleteByPort(unit int, mod Module, port Port, flags L2DeleteFlags) error {
	rc := C.opennsl_l2_addr_delete_by_port(C.int(unit), mod.C(), port.C(), flags.C())
	return ParseError(rc)
}

func L2AddrDeleteByMAC(unit int, hwaddr net.HardwareAddr, flags L2DeleteFlags) error {
	mac := NewMAC(hwaddr)
	rc := C.opennsl_l2_addr_delete_by_mac(C.int(unit), &mac[0], flags.C())
	return ParseError(rc)
}

func L2AddrDeleteByVID(unit int, vid Vlan, flags L2DeleteFlags) error {
	rc := C.opennsl_l2_addr_delete_by_vlan(C.int(unit), vid.C(), flags.C())
	return ParseError(rc)
}

func L2AddrDeleteByTrunkID(unit int, tid Trunk, flags L2DeleteFlags) error {
	rc := C.opennsl_l2_addr_delete_by_trunk(C.int(unit), tid.C(), flags.C())
	return ParseError(rc)
}

func L2AddrDeleteByMACPort(unit int, hwaddr net.HardwareAddr, mod Module, port Port, flags L2DeleteFlags) error {
	mac := NewMAC(hwaddr)
	rc := C.opennsl_l2_addr_delete_by_mac_port(C.int(unit), &mac[0], mod.C(), port.C(), flags.C())
	return ParseError(rc)
}

func L2AddrDeleteByVIDPort(unit int, vid Vlan, mod Module, port Port, flags L2DeleteFlags) error {
	rc := C.opennsl_l2_addr_delete_by_vlan_port(C.int(unit), vid.C(), mod.C(), port.C(), flags.C())
	return ParseError(rc)
}

func L2AddrDeleteByVIDTrunkID(unit int, vid Vlan, tid Trunk, flags L2DeleteFlags) error {
	rc := C.opennsl_l2_addr_delete_by_vlan_trunk(C.int(unit), vid.C(), tid.C(), flags.C())
	return ParseError(rc)
}

func L2AddrGet(unit int, hwaddr net.HardwareAddr, vid Vlan) (*L2Addr, error) {
	addr := NewL2Addr(hwaddr, vid)
	c_mac := NewMAC(hwaddr)
	rc := C.opennsl_l2_addr_get(C.int(unit), &c_mac[0], vid.C(), addr.C())
	if err := ParseError(rc); err != nil {
		return nil, err
	}

	return addr, nil
}

func L2AddrAgeTimerSet(unit int, ageSec int) error {
	rc := C.opennsl_l2_age_timer_set(C.int(unit), C.int(ageSec))
	return ParseError(rc)
}

func L2AddrAgeTimerGet(unit int) (int, error) {
	var ageSec C.int = 0
	rc := C.opennsl_l2_age_timer_get(C.int(unit), &ageSec)
	return int(ageSec), ParseError(rc)
}

func L2AddrFreeze(unit int) error {
	rc := C.opennsl_l2_addr_freeze(C.int(unit))
	return ParseError(rc)
}

func L2AddrThaw(unit int) error {
	rc := C.opennsl_l2_addr_thaw(C.int(unit))
	return ParseError(rc)
}

func (v *L2Addr) Replace(unit int, flags L2ReplaceFlags, newModule Module, newPort Port, newTrunk Trunk) error {
	rc := C.opennsl_l2_replace(C.int(unit), flags.C(), v.C(), newModule.C(), newPort.C(), newTrunk.C())
	return ParseError(rc)
}

//
// opennsl_l2_traverse
//
type L2TraverseHandler func(int, *L2Addr) OpenNSLError

var l2TraverseHandlers = NewCallbackMap()

// extern int go_opennsl_l2_traverse_cb(int unit, opennsl_l2_addr_t *info, void *user_data);
//export go_opennsl_l2_traverse_cb
func go_opennsl_l2_traverse_cb(c_unit C.int, c_l2addr *C.opennsl_l2_addr_t, c_userdata unsafe.Pointer) int {
	n := (*uint64)(c_userdata)
	if h, ok := l2TraverseHandlers.Get(*n); ok {
		callback := h.(L2TraverseHandler)
		rc := callback(int(c_unit), (*L2Addr)(c_l2addr))
		return int(rc)
	}

	return int(E_PARAM)
}

func L2Traverse(unit int, h L2TraverseHandler) error {
	n := l2TraverseHandlers.Add(h)
	defer l2TraverseHandlers.Del(n)

	rc := C.opennsl_l2_traverse(C.int(unit), C.opennsl_l2_traverse_cb(C._opennsl_l2_traverse_cb), unsafe.Pointer(&n))
	return ParseError(rc)
}
