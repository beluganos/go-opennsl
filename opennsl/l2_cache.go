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
*/
import "C"

import (
	"net"
)

//
// L2CacheAddrIndex
//
type L2CacheAddrIndex C.int

func (v L2CacheAddrIndex) C() C.int {
	return C.int(v)
}

const (
	L2_CACHE_ADDR_INDEX_FIRST L2CacheAddrIndex = -1
)

//
// L2CacheAddr
//
type L2CacheAddr C.opennsl_l2_cache_addr_t

func (v *L2CacheAddr) C() *C.opennsl_l2_cache_addr_t {
	return (*C.opennsl_l2_cache_addr_t)(v)
}

func (v *L2CacheAddr) Flags() L2CacheFlags {
	return L2CacheFlags(v.flags)
}

func (v *L2CacheAddr) SetFlags(flags L2CacheFlags) {
	v.flags = flags.C()
}

func (v *L2CacheAddr) StationFlags() L2StationFlags {
	return L2StationFlags(v.station_flags)
}

func (v *L2CacheAddr) SetStationFlags(flags L2StationFlags) {
	v.station_flags = flags.C()
}

func (v *L2CacheAddr) MAC() (net.HardwareAddr, net.HardwareAddr) {
	return ParseMAC(v.mac), ParseMAC(v.mac_mask)
}

func (v *L2CacheAddr) SetMAC(mac net.HardwareAddr) {
	v.mac = NewMAC(mac)
}

func (v *L2CacheAddr) SetMACMask(mask net.HardwareAddr) {
	v.mac_mask = NewMAC(mask)
}

func (v *L2CacheAddr) SrcPort() (Port, Port) {
	return Port(v.src_port), Port(v.src_port_mask)
}

func (v *L2CacheAddr) SetSrcPort(srcPort Port) {
	v.src_port = srcPort.C()
}

func (v *L2CacheAddr) SetSrcPortMask(srcPortMask Port) {
	v.src_port_mask = srcPortMask.C()
}

func (v *L2CacheAddr) DestModule() Port {
	return Port(v.dest_modid)
}

func (v *L2CacheAddr) SetDestModule(destModule Module) {
	v.dest_modid = destModule.C()
}

func (v *L2CacheAddr) DestPort() Port {
	return Port(v.dest_port)
}

func (v *L2CacheAddr) SetDestPort(destPort Port) {
	v.dest_port = destPort.C()
}

func (v *L2CacheAddr) DestTrunk() Trunk {
	return Trunk(v.dest_trunk)
}

func (v *L2CacheAddr) SetDestTrunk(destTrunk Trunk) {
	v.dest_trunk = destTrunk.C()
}

func (v *L2CacheAddr) Prio() int {
	return int(v.prio)
}

func (v *L2CacheAddr) SetPrio(prio int) {
	v.prio = C.int(prio)
}

func (v *L2CacheAddr) DestPorts() *PBmp {
	dest_ports := v.dest_ports
	return (*PBmp)(&dest_ports)
}

func (v *L2CacheAddr) SetDestPorts(dstPorts *PBmp) {
	v.dest_ports = *dstPorts.C()
}

func (v *L2CacheAddr) LookupClass() int {
	return int(v.lookup_class)
}

func (v *L2CacheAddr) SetLookupClass(lookupClass int) {
	v.lookup_class = C.int(lookupClass)
}

func (v *L2CacheAddr) Subtype() uint8 {
	return uint8(v.subtype)
}

func (v *L2CacheAddr) SetSubtype(subtype uint8) {
	v.subtype = C.uint8(subtype)
}

func (v *L2CacheAddr) EncapID() EncapID {
	return EncapID(v.encap_id)
}

func (v *L2CacheAddr) SetEncapID(encapID EncapID) {
	v.encap_id = encapID.C()
}

func (v *L2CacheAddr) Group() Multicast {
	return Multicast(v.group)
}

func (v *L2CacheAddr) SetGroup(group Multicast) {
	v.group = group.C()
}

func (v *L2CacheAddr) EtherType() (Ethertype, Ethertype) {
	return Ethertype(v.ethertype), Ethertype(v.ethertype_mask)
}

func (v *L2CacheAddr) SetEtherType(ethertype Ethertype) {
	v.ethertype = ethertype.C()
}

func (v *L2CacheAddr) SetEtherTypeMask(mask Ethertype) {
	v.ethertype_mask = mask.C()
}

func (v *L2CacheAddr) Vlan() Vlan {
	return Vlan(v.vlan)
}

func (v *L2CacheAddr) SetVlan(vlan Vlan) {
	v.vlan = vlan.C()
}

//
// API
//
func NewL2CacheAddr() *L2CacheAddr {
	addr := &L2CacheAddr{}
	addr.Init()
	return addr
}

func (v *L2CacheAddr) Init() {
	C.opennsl_l2_cache_addr_t_init(v.C())
}

func L2CacheAddrInit(unit int) error {
	rc := C.opennsl_l2_cache_init(C.int(unit))
	return ParseError(rc)
}

func L2CacheAddrSizeGet(unit int) (int, error) {
	var size C.int = 0
	rc := C.opennsl_l2_cache_size_get(C.int(unit), &size)
	return int(size), ParseError(rc)
}

func (v *L2CacheAddr) Set(unit int, index L2CacheAddrIndex) (L2CacheAddrIndex, error) {
	var c_used C.int = 0
	rc := C.opennsl_l2_cache_set(C.int(unit), index.C(), v.C(), &c_used)
	return L2CacheAddrIndex(c_used), ParseError(rc)
}

func (v L2CacheAddrIndex) Get(unit int) (*L2CacheAddr, error) {
	cache := NewL2CacheAddr()
	rc := C.opennsl_l2_cache_get(C.int(unit), v.C(), cache.C())
	if err := ParseError(rc); err != nil {
		return nil, err
	}

	return cache, nil
}

func (v L2CacheAddrIndex) Delete(unit int) error {
	rc := C.opennsl_l2_cache_delete(C.int(unit), v.C())
	return ParseError(rc)
}

func L2CacheAddrDeleteAll(unit int) error {
	rc := C.opennsl_l2_cache_delete_all(C.int(unit))
	return ParseError(rc)
}
