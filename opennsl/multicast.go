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
#include <opennsl/error.h>
#include <opennsl/multicast.h>
#include "helper.h"
*/
import "C"

import (
	"unsafe"
)

//
// opennsl_multicast_t
//
type Multicast C.opennsl_multicast_t

func (v Multicast) C() C.opennsl_multicast_t {
	return C.opennsl_multicast_t(v)
}

const (
	MULTICAST Multicast = 0
)

//
// MulticastFlags
//
type MulticastFlags C.uint32

func (v MulticastFlags) C() C.uint32 {
	return C.uint32(v)
}

func NewMulticastFlags(flags ...MulticastFlags) MulticastFlags {
	v := MULTICAST_NONE
	for _, flag := range flags {
		v |= flag
	}
	return v
}

const (
	MULTICAST_NONE             MulticastFlags = 0
	MULTICAST_WITH_ID          MulticastFlags = C.OPENNSL_MULTICAST_WITH_ID
	MULTICAST_TYPE_L2          MulticastFlags = C.OPENNSL_MULTICAST_TYPE_L2
	MULTICAST_TYPE_L3          MulticastFlags = C.OPENNSL_MULTICAST_TYPE_L3
	MULTICAST_TYPE_VLAN        MulticastFlags = C.OPENNSL_MULTICAST_TYPE_VLAN
	MULTICAST_TYPE_VXLAN       MulticastFlags = C.OPENNSL_MULTICAST_TYPE_VXLAN
	MULTICAST_TYPE_PORTS_GROUP MulticastFlags = C.OPENNSL_MULTICAST_TYPE_PORTS_GROUP
	MULTICAST_TYPE_MASK        MulticastFlags = C.OPENNSL_MULTICAST_TYPE_MASK
	MULTICAST_INGRESS_GROUP    MulticastFlags = C.OPENNSL_MULTICAST_INGRESS_GROUP
	MULTICAST_EGRESS_GROUP     MulticastFlags = C.OPENNSL_MULTICAST_EGRESS_GROUP
)

//
// MulticastControl
//
type MulticastControl C.opennsl_multicast_control_t

func (v MulticastControl) C() C.opennsl_multicast_control_t {
	return C.opennsl_multicast_control_t(v)
}

const (
	MulticastControlMtu     MulticastControl = C.opennslMulticastControlMtu
	MulticastVpTrunkResolve MulticastControl = C.opennslMulticastVpTrunkResolve
	MulticastRemapGroup     MulticastControl = C.opennslMulticastRemapGroup
	MulticastControlCount   MulticastControl = C.opennslMulticastControlCount
)

//
// Multicast methods
//
func MulticastInit(unit int) error {
	rc := C.opennsl_multicast_init(C.int(unit))
	return ParseError(rc)
}

func MulticastCreate(unit int, flags MulticastFlags, mc Multicast) (Multicast, error) {
	c_mc := mc.C()

	rc := C.opennsl_multicast_create(C.int(unit), flags.C(), &c_mc)
	return Multicast(c_mc), ParseError(rc)
}

func (v Multicast) Create(unit int, flags MulticastFlags) (Multicast, error) {
	return MulticastCreate(unit, flags, v)
}

func MulticastDestroy(unit int, mc Multicast) error {
	rc := C.opennsl_multicast_destroy(C.int(unit), mc.C())
	return ParseError(rc)
}

func (v Multicast) Destroy(unit int) error {
	return MulticastDestroy(unit, v)
}

func MulticastL3EncapGet(unit int, v Multicast, gport GPort, iface L3IfaceID) (EncapID, error) {
	c_iface := C.opennsl_if_t(0)

	rc := C.opennsl_multicast_l3_encap_get(C.int(unit), v.C(), gport.C(), iface.C(), &c_iface)
	return EncapID(c_iface), ParseError(rc)
}

func (v Multicast) L3EncapGet(unit int, gport GPort, iface L3IfaceID) (EncapID, error) {
	return MulticastL3EncapGet(unit, v, gport, iface)
}

func MulticastL2EncapGet(unit int, v Multicast, gport GPort, vid Vlan) (EncapID, error) {
	c_iface := C.opennsl_if_t(0)

	rc := C.opennsl_multicast_l2_encap_get(C.int(unit), v.C(), gport.C(), vid.C(), &c_iface)
	return EncapID(c_iface), ParseError(rc)
}

func (v Multicast) L2EncapGet(unit int, gport GPort, vid Vlan) (EncapID, error) {
	return MulticastL2EncapGet(unit, v, gport, vid)
}

func MulticastVlanEncapGet(unit int, v Multicast, gport GPort, vport GPort) (EncapID, error) {
	c_iface := C.opennsl_if_t(0)

	rc := C.opennsl_multicast_vlan_encap_get(C.int(unit), v.C(), gport.C(), vport.C(), &c_iface)
	return EncapID(c_iface), ParseError(rc)
}

func (v Multicast) VlanEncapGet(unit int, gport GPort, vport GPort) (EncapID, error) {
	return MulticastVlanEncapGet(unit, v, gport, vport)
}

func MulticastEgressObjectEncapGet(unit int, v Multicast, l3egrID L3EgressID) (EncapID, error) {
	c_iface := C.opennsl_if_t(0)

	rc := C.opennsl_multicast_egress_object_encap_get(C.int(unit), v.C(), l3egrID.C(), &c_iface)
	return EncapID(c_iface), ParseError(rc)
}

func (v Multicast) EgressObjectEncapGet(unit int, l3egrID L3EgressID) (EncapID, error) {
	return MulticastEgressObjectEncapGet(unit, v, l3egrID)
}

func MulticastEgressAdd(unit int, v Multicast, gport GPort, encapID EncapID) error {
	rc := C.opennsl_multicast_egress_add(C.int(unit), v.C(), gport.C(), encapID.C())
	return ParseError(rc)
}

func (v Multicast) EgressAdd(unit int, gport GPort, encapID EncapID) error {
	return MulticastEgressAdd(unit, v, gport, encapID)
}

func MulticastEgressDelete(unit int, v Multicast, gport GPort, encapId EncapID) error {
	rc := C.opennsl_multicast_egress_delete(C.int(unit), v.C(), gport.C(), encapId.C())
	return ParseError(rc)
}

func (v Multicast) EgressDelete(unit int, gport GPort, encapId EncapID) error {
	return MulticastEgressDelete(unit, v, gport, encapId)
}

func (v Multicast) EgressDeleteAll(unit int) error {
	rc := C.opennsl_multicast_egress_delete_all(C.int(unit), v.C())
	return ParseError(rc)
}

func MulticastEgressSet(unit int, v Multicast, gports []GPort, encapIDs []EncapID) error {
	num := len(gports)
	if num != len(encapIDs) {
		return E_PARAM.Error()
	}

	c_gports := make([]C.opennsl_gport_t, num)
	c_encaps := make([]C.opennsl_if_t, num)
	for index, gport := range gports {
		c_gports[index] = gport.C()
	}
	for index, encapID := range encapIDs {
		c_encaps[index] = encapID.C()
	}

	rc := C.opennsl_multicast_egress_set(C.int(unit), v.C(), C.int(num), &c_gports[0], &c_encaps[0])
	return ParseError(rc)
}

func (v Multicast) EgressSet(unit int, gports []GPort, encapIDs []EncapID) error {
	return MulticastEgressSet(unit, v, gports, encapIDs)
}

func MulticastEgressCount(unit int, v Multicast) (int, error) {
	c_count := C.int(0)

	rc := C.opennsl_multicast_egress_get(C.int(unit), v.C(), 0, nil, nil, &c_count)
	return int(c_count), ParseError(rc)
}

func (v Multicast) EgressCount(unit int) (int, error) {
	return MulticastEgressCount(unit, v)
}

func MulticastEgressGet(unit int, v Multicast, portMax int) ([]GPort, []EncapID, error) {
	if portMax <= 0 {
		cnt, err := v.EgressCount(unit)
		if err != nil {
			return nil, nil, err
		}
		portMax = cnt
	}

	c_gports := make([]C.opennsl_gport_t, portMax)
	c_encaps := make([]C.opennsl_if_t, portMax)
	c_count := C.int(0)

	rc := C.opennsl_multicast_egress_get(C.int(unit), v.C(), C.int(portMax), &c_gports[0], &c_encaps[0], &c_count)
	if err := ParseError(rc); err != nil {
		return nil, nil, err
	}

	gports := make([]GPort, c_count)
	encaps := make([]EncapID, c_count)

	for index := 0; index < int(c_count); index++ {
		gports[index] = GPort(c_gports[index])
		encaps[index] = EncapID(c_encaps[index])
	}

	return gports, encaps, nil
}

func (v Multicast) EgressGet(unit int, portMax int) ([]GPort, []EncapID, error) {
	return MulticastEgressGet(unit, v, portMax)
}

func MulticastIngressAdd(unit int, v Multicast, gport GPort, encapId EncapID) error {
	rc := C.opennsl_multicast_ingress_add(C.int(unit), v.C(), gport.C(), encapId.C())
	return ParseError(rc)
}

func (v Multicast) IngressAdd(unit int, gport GPort, encapId EncapID) error {
	return MulticastIngressAdd(unit, v, gport, encapId)
}

func MulticastIngressDelete(unit int, v Multicast, gport GPort, encapId EncapID) error {
	rc := C.opennsl_multicast_ingress_delete(C.int(unit), v.C(), gport.C(), encapId.C())
	return ParseError(rc)
}

func (v Multicast) IngressDelete(unit int, gport GPort, encapId EncapID) error {
	return MulticastIngressDelete(unit, v, gport, encapId)
}

func MulticastIngressDeleteAll(unit int, v Multicast) error {
	rc := C.opennsl_multicast_ingress_delete_all(C.int(unit), v.C())
	return ParseError(rc)
}

func (v Multicast) IngressDeleteAll(unit int) error {
	return MulticastIngressDeleteAll(unit, v)
}

func MulticastIngressSet(unit int, v Multicast, gports []GPort, encapIDs []EncapID) error {
	num := len(gports)
	if num != len(encapIDs) {
		return E_PARAM.Error()
	}

	c_gports := make([]C.opennsl_gport_t, num)
	c_encaps := make([]C.opennsl_if_t, num)
	for index, gport := range gports {
		c_gports[index] = gport.C()
	}
	for index, encapID := range encapIDs {
		c_encaps[index] = encapID.C()
	}

	rc := C.opennsl_multicast_ingress_set(C.int(unit), v.C(), C.int(num), &c_gports[0], &c_encaps[0])
	return ParseError(rc)
}

func (v Multicast) IngressSet(unit int, gports []GPort, encapIDs []EncapID) error {
	return MulticastIngressSet(unit, v, gports, encapIDs)
}

func MulticastIngressCount(unit int, v Multicast) (int, error) {
	c_count := C.int(0)

	rc := C.opennsl_multicast_ingress_get(C.int(unit), v.C(), 0, nil, nil, &c_count)
	return int(c_count), ParseError(rc)
}

func (v Multicast) IngressCount(unit int) (int, error) {
	return MulticastIngressCount(unit, v)
}

func MulticastIngressGet(unit int, v Multicast, portMax int) ([]GPort, []EncapID, error) {
	if portMax <= 0 {
		cnt, err := v.IngressCount(unit)
		if err != nil {
			return nil, nil, err
		}
		portMax = cnt
	}

	c_gports := make([]C.opennsl_gport_t, portMax)
	c_encaps := make([]C.opennsl_if_t, portMax)
	c_count := C.int(0)

	rc := C.opennsl_multicast_ingress_get(C.int(unit), v.C(), C.int(portMax), &c_gports[0], &c_encaps[0], &c_count)
	if err := ParseError(rc); err != nil {
		return nil, nil, err
	}

	gports := make([]GPort, c_count)
	encaps := make([]EncapID, c_count)

	for index := 0; index < int(c_count); index++ {
		gports[index] = GPort(c_gports[index])
		encaps[index] = EncapID(c_encaps[index])
	}

	return gports, encaps, nil
}

func (v Multicast) IngressGet(unit int, portMax int) ([]GPort, []EncapID, error) {
	return MulticastIngressGet(unit, v, portMax)
}

func MulticastGroupGet(unit int, v Multicast) (MulticastFlags, error) {
	c_flags := C.uint32(0)

	rc := C.opennsl_multicast_group_get(C.int(unit), v.C(), &c_flags)
	return MulticastFlags(c_flags), ParseError(rc)
}

func (v Multicast) GroupGet(unit int) (MulticastFlags, error) {
	return MulticastGroupGet(unit, v)
}

func MulticastIsFree(unit int, v Multicast) (bool, error) {
	rc := C.opennsl_multicast_group_is_free(C.int(unit), v.C())
	switch rc {
	case C.OPENNSL_E_NONE:
		return true, nil
	case C.OPENNSL_E_EXISTS:
		return false, nil
	default:
		return false, ParseError(rc)
	}
}

func (v Multicast) IsFree(unit int) (bool, error) {
	return MulticastIsFree(unit, v)
}

func MulticastGroupFreeRange(unit int, flags MulticastFlags) (Multicast, Multicast, error) {
	c_min := C.opennsl_multicast_t(0)
	c_max := C.opennsl_multicast_t(0)

	rc := C.opennsl_multicast_group_free_range_get(C.int(unit), flags.C(), &c_min, &c_max)
	return Multicast(c_min), Multicast(c_max), ParseError(rc)
}

type MulticastGroupTraverseCallback func(int, Multicast, MulticastFlags) int

var multicastGroupTraverseCallbacks = NewCallbackMap()

//export go_opennsl_multicast_group_traverse_cb
func go_opennsl_multicast_group_traverse_cb(unit C.int, group C.opennsl_multicast_t, flags C.uint32, data unsafe.Pointer) int {
	n := (*uint64)(data)
	if h, ok := multicastGroupTraverseCallbacks.Get(*n); ok {
		callback := h.(MulticastGroupTraverseCallback)
		return callback(int(unit), Multicast(group), MulticastFlags(flags))
	}

	return int(E_PARAM)
}

func MuticastGroupTraverse(unit int, callback MulticastGroupTraverseCallback, flags MulticastFlags) error {
	n := multicastGroupTraverseCallbacks.Add(callback)
	defer multicastGroupTraverseCallbacks.Del(n)

	rc := C.opennsl_multicast_group_traverse(C.int(unit), (C.opennsl_multicast_group_traverse_cb_t)(C._opennsl_multicast_group_traverse_cb), flags.C(), unsafe.Pointer(&n))
	return ParseError(rc)
}
