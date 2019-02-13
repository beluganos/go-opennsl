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
// L3Route
//
type L3Route C.opennsl_l3_route_t

func (v *L3Route) C() *C.opennsl_l3_route_t {
	return (*C.opennsl_l3_route_t)(v)
}

func (v *L3Route) Flags() L3Flags {
	return L3Flags(v.l3a_flags)
}

func (v *L3Route) SetFlags(flags L3Flags) {
	v.l3a_flags = flags.C()
}

func (v *L3Route) VRF() Vrf {
	return Vrf(v.l3a_vrf)
}

func (v *L3Route) SetVRF(vrf Vrf) {
	v.l3a_vrf = vrf.C()
}

func (v *L3Route) IP4Net() *net.IPNet {
	return &net.IPNet{
		IP:   ParseIP4(v.l3a_subnet),
		Mask: ParseIP4Mask(v.l3a_ip_mask),
	}
}

func (v *L3Route) SetIP4Net(ip4 *net.IPNet) error {
	if err := v.SetIP4(ip4.IP); err != nil {
		return err
	}

	return v.SetIP4Mask(ip4.Mask)
}

func (v *L3Route) SetIP4(subnet net.IP) error {
	c_subnet, err := NewIP4(subnet)
	if err != nil {
		return err
	}

	v.l3a_subnet = c_subnet
	return nil
}

func (v *L3Route) SetIP4Mask(mask net.IPMask) error {
	c_mask, err := NewIP4Mask(mask)
	if err != nil {
		return err
	}

	v.l3a_ip_mask = c_mask
	return nil
}

func (v *L3Route) IP6Net() *net.IPNet {
	return &net.IPNet{
		IP:   ParseIP6(v.l3a_ip6_net),
		Mask: ParseIP6Mask(v.l3a_ip6_mask),
	}
}

func (v *L3Route) SetIP6Net(ip6 *net.IPNet) error {
	if err := v.SetIP6(ip6.IP); err != nil {
		return err
	}
	return v.SetIP6Mask(ip6.Mask)
}

func (v *L3Route) SetIP6(subnet net.IP) error {
	c_subnet, err := NewIP6(subnet)
	if err != nil {
		return err
	}

	v.l3a_ip6_net = c_subnet
	return nil
}

func (v *L3Route) SetIP6Mask(mask net.IPMask) error {
	c_mask, err := NewIP6Mask(mask)
	if err != nil {
		return err
	}

	v.l3a_ip6_mask = c_mask
	return nil
}

func (v *L3Route) EgressID() L3EgressID {
	return L3EgressID(v.l3a_intf)
}

func (v *L3Route) SetEgressID(iface L3EgressID) {
	v.l3a_intf = iface.C()
}

func (v *L3Route) PortTGID() Port {
	return Port(v.l3a_port_tgid)
}

func (v *L3Route) SetPortTGID(port Port) {
	v.l3a_port_tgid = port.C()
}

func (v *L3Route) Pri() Cos {
	return Cos(v.l3a_pri)
}

func (v *L3Route) SetPri(pri Cos) {
	v.l3a_pri = pri.C()
}

//
// API
//
func NewL3Route() *L3Route {
	route := &L3Route{}
	route.Init()
	return route
}

func L3RouteInit(v *L3Route) {
	C.opennsl_l3_route_t_init(v.C())
}

func (v *L3Route) String() string {
	dst := v.IP4Net()
	if dst.IP == nil {
		dst = v.IP6Net()
	}

	return fmt.Sprintf("L3Route(%s, vrf:%d, egress:%d)", dst, v.VRF(), v.EgressID())
}

func (v *L3Route) Init() {
	L3RouteInit(v)
}

func L3RouteAdd(unit int, v *L3Route) error {
	rc := C.opennsl_l3_route_add(C.int(unit), v.C())
	return ParseError(rc)
}

func (v *L3Route) Add(unit int) error {
	return L3RouteAdd(unit, v)
}

func L3RouteDelete(unit int, v *L3Route) error {
	rc := C.opennsl_l3_route_delete(C.int(unit), v.C())
	return ParseError(rc)
}

func (v *L3Route) Delete(unit int) error {
	return L3RouteDelete(unit, v)
}

func L3RouteDeleteByIface(unit int, v *L3Route) error {
	rc := C.opennsl_l3_route_delete_by_interface(C.int(unit), v.C())
	return ParseError(rc)
}

func (v *L3Route) DeleteByIface(unit int) error {
	return L3RouteDeleteByIface(unit, v)
}

func (v *L3Route) DeleteAll(unit int) error {
	rc := C.opennsl_l3_route_delete_all(C.int(unit), v.C())
	return ParseError(rc)
}

func L3RouteGet(unit int, ipnet *net.IPNet) (*L3Route, error) {
	route := NewL3Route()

	if ipv4 := ipnet.IP.To4(); ipv4 != nil {
		route.SetIP4(ipv4)
		route.SetIP4Mask(ipnet.Mask)
	}

	if ipv6 := ipnet.IP.To16(); ipv6 != nil {
		route.SetIP6(ipv6)
		route.SetIP6Mask(ipnet.Mask)
	}

	rc := C.opennsl_l3_route_get(C.int(unit), route.C())
	if OpenNSLError(rc) == E_NOT_FOUND {
		return nil, nil
	}

	if err := ParseError(rc); err != nil {
		return nil, err
	}

	return route, nil
}

func L3RouteMultipathGet(unit int, v *L3Route, maxPath int) ([]L3Route, error) {
	c_count := C.int(0)
	c_routes := make([]C.opennsl_l3_route_t, maxPath)

	rc := C.opennsl_l3_route_multipath_get(C.int(unit), v.C(), &c_routes[0], C.int(maxPath), &c_count)
	if err := ParseError(rc); err != nil {
		return nil, err
	}

	routes := make([]L3Route, int(c_count))
	for index := 0; index < int(c_count); index++ {
		routes[index] = L3Route(c_routes[index])
	}

	return routes, nil
}

func (v *L3Route) MultipathGet(unit int, maxPath int) ([]L3Route, error) {
	return L3RouteMultipathGet(unit, v, maxPath)
}

type L3RouteTraverseCallback func(int, int, *L3Route) OpenNSLError

var l3RouteTraverseCallbacks = NewCallbackMap()

//export go_opennsl_l3_route_traverse_cb
func go_opennsl_l3_route_traverse_cb(c_unit C.int, c_index C.int, c_route *C.opennsl_l3_route_t, c_data unsafe.Pointer) int {
	n := (*uint64)(c_data)
	if h, ok := l3RouteTraverseCallbacks.Get(*n); ok {
		callback := h.(L3RouteTraverseCallback)
		rc := callback(int(c_unit), int(c_index), (*L3Route)(c_route))
		return int(rc)
	}

	return int(E_PARAM)
}

func L3RouteTraverse(unit int, flags uint32, start uint32, end uint32, callback L3RouteTraverseCallback) error {
	n := l3RouteTraverseCallbacks.Add(callback)
	defer l3RouteTraverseCallbacks.Del(n)

	rc := C.opennsl_l3_route_traverse(C.int(unit), C.uint32(flags), C.uint32(start), C.uint32(end), C.opennsl_l3_route_traverse_cb(C._opennsl_l3_route_traverse_cb), unsafe.Pointer(&n))
	return ParseError(rc)
}

func L3RouteMaxECMPSet(unit int, max int) error {
	rc := C.opennsl_l3_route_max_ecmp_set(C.int(unit), C.int(max))
	return ParseError(rc)
}

func L3RouteMaxECMPGet(unit int) (int, error) {
	c_max := C.int(0)
	rc := C.opennsl_l3_route_max_ecmp_get(C.int(unit), &c_max)
	return int(c_max), ParseError(rc)
}
