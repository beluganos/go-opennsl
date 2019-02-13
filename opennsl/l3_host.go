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
// L3Host
//
type L3Host C.opennsl_l3_host_t

func (v *L3Host) C() *C.opennsl_l3_host_t {
	return (*C.opennsl_l3_host_t)(v)
}

func (v *L3Host) Flags() L3Flags {
	return L3Flags(v.l3a_flags)
}

func (v *L3Host) SetFlags(flags L3Flags) {
	v.l3a_flags = flags.C()
}

func (v *L3Host) VRF() Vrf {
	return Vrf(v.l3a_vrf)
}

func (v *L3Host) SetVRF(vrf Vrf) {
	v.l3a_vrf = vrf.C()
}

func (v *L3Host) IPAddr() net.IP {
	return ParseIP4(v.l3a_ip_addr)
}

func (v *L3Host) SetIPAddr(ip net.IP) error {
	c_ip, err := NewIP4(ip)
	if err != nil {
		return err
	}

	v.l3a_ip_addr = c_ip
	return nil
}

func (v *L3Host) IP6Addr() net.IP {
	return ParseIP6(v.l3a_ip6_addr)
}

func (v *L3Host) SetIP6Addr(ip net.IP) error {
	c_ip, err := NewIP6(ip)
	if err != nil {
		return err
	}

	v.l3a_ip6_addr = c_ip
	return nil
}

func (v *L3Host) Pri() Cos {
	return Cos(v.l3a_pri)
}

func (v *L3Host) SetPri(cos Cos) {
	v.l3a_pri = cos.C()
}

func (v *L3Host) EgressID() L3EgressID {
	return L3EgressID(v.l3a_intf)
}

func (v *L3Host) SetEgressID(l3egrID L3EgressID) {
	v.l3a_intf = l3egrID.C()
}

func (v *L3Host) NexthopMAC() net.HardwareAddr {
	return ParseMAC(v.l3a_nexthop_mac)
}

func (v *L3Host) SetNexthopMAC(mac net.HardwareAddr) {
	v.l3a_nexthop_mac = NewMAC(mac)
}

func (v *L3Host) PortTGID() Port {
	return Port(v.l3a_port_tgid)
}

func (v *L3Host) SetPortTGID(port Port) {
	v.l3a_port_tgid = port.C()
}

//
// API
//
func NewL3Host() *L3Host {
	host := &L3Host{}
	host.Init()
	return host
}

func L3HostInit(v *L3Host) {
	C.opennsl_l3_host_t_init(v.C())
}

func (v *L3Host) String() string {
	ip := v.IPAddr()
	if ip == nil {
		ip = v.IP6Addr()
	}

	return fmt.Sprintf("L3Host(%s, vrf:%d, egress:%d)", ip, v.VRF(), v.EgressID())
}

func (v *L3Host) Init() {
	L3HostInit(v)
}

func L3HostFind(unit int, ip net.IP) (*L3Host, error) {
	host := NewL3Host()

	if ipv4 := ip.To4(); ipv4 != nil {
		host.SetIPAddr(ipv4)
	} else if ipv6 := ip.To16(); ipv6 != nil {
		host.SetIP6Addr(ipv6)
	}

	rc := C.opennsl_l3_host_find(C.int(unit), host.C())
	if OpenNSLError(rc) == E_NOT_FOUND {
		return nil, nil
	}

	if err := ParseError(rc); err != nil {
		return nil, err
	}

	return host, nil
}

func L3HostAdd(unit int, v *L3Host) error {
	rc := C.opennsl_l3_host_add(C.int(unit), v.C())
	return ParseError(rc)
}

func (v *L3Host) Add(unit int) error {
	return L3HostAdd(unit, v)
}

func L3HostDelete(unit int, v *L3Host) error {
	rc := C.opennsl_l3_host_delete(C.int(unit), v.C())
	return ParseError(rc)
}

func (v *L3Host) Delete(unit int) error {
	return L3HostDelete(unit, v)
}

func L3HostDeleteByIface(unit int, v *L3Host) error {
	rc := C.opennsl_l3_host_delete_by_interface(C.int(unit), v.C())
	return ParseError(rc)
}

func (v *L3Host) DeleteByIface(unit int) error {
	return L3HostDeleteByIface(unit, v)
}

func L3HostDeleteAll(unit int, v *L3Host) error {
	rc := C.opennsl_l3_host_delete_all(C.int(unit), v.C())
	return ParseError(rc)
}

func (v *L3Host) DeleteAll(unit int) error {
	return L3HostDeleteAll(unit, v)
}

type L3HostTraverseCallback func(int, int, *L3Host) OpenNSLError

var l3HostTraverseCallbacks = NewCallbackMap()

//export go_opennsl_l3_host_traverse_cb
func go_opennsl_l3_host_traverse_cb(c_unit C.int, c_index C.int, c_host *C.opennsl_l3_host_t, c_data unsafe.Pointer) int {
	n := (*uint64)(c_data)
	if h, ok := l3HostTraverseCallbacks.Get(*n); ok {
		callback := h.(L3HostTraverseCallback)
		rc := callback(int(c_unit), int(c_index), (*L3Host)(c_host))
		return int(rc)
	}

	return int(E_PARAM)
}

func L3HostTraverse(unit int, flags uint32, start uint32, end uint32, callback L3HostTraverseCallback) error {
	n := l3HostTraverseCallbacks.Add(callback)
	defer l3HostTraverseCallbacks.Del(n)

	rc := C.opennsl_l3_host_traverse(C.int(unit), C.uint32(flags), C.uint32(start), C.uint32(end), C.opennsl_l3_host_traverse_cb(C._opennsl_l3_host_traverse_cb), unsafe.Pointer(&n))
	return ParseError(rc)
}
