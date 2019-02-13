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

//
// L3IngressID
//
type L3IngressID C.opennsl_if_t

func (v L3IngressID) C() C.opennsl_if_t {
	return C.opennsl_if_t(v)
}

const (
	L3_INGRESS L3IngressID = L3IngressID(IFACE_NONE)
)

//
// L3Ingress
//
type L3Ingress C.opennsl_l3_ingress_t

func (v *L3Ingress) C() *C.opennsl_l3_ingress_t {
	return (*C.opennsl_l3_ingress_t)(v)
}

func (v *L3Ingress) Flags() L3IngressFlags {
	return L3IngressFlags(v.flags)
}

func (v *L3Ingress) SetFlags(flags L3IngressFlags) {
	v.flags = flags.C()
}

func (v *L3Ingress) VRF() Vrf {
	return Vrf(v.vrf)
}

func (v *L3Ingress) SetVRF(vrf Vrf) {
	v.vrf = vrf.C()
}

func (v *L3Ingress) UrpfMode() L3IngressUrpf {
	return L3IngressUrpf(v.urpf_mode)
}

func (v *L3Ingress) SetUrpfMode(mode L3IngressUrpf) {
	v.urpf_mode = mode.C()
}

func (v *L3Ingress) IfaceClass() int {
	return int(v.intf_class)
}

func (v *L3Ingress) SetIfaceClass(class int) {
	v.intf_class = C.int(class)
}

func (v *L3Ingress) IpmcVID() Vlan {
	return Vlan(v.ipmc_intf_id)
}

func (v *L3Ingress) SetIpmcVID(vid Vlan) {
	v.ipmc_intf_id = vid.C()
}

func (v *L3Ingress) QoSMap() int {
	return int(v.qos_map_id)
}

func (v *L3Ingress) SetQoSMap(qosMap int) {
	v.qos_map_id = C.int(qosMap)
}

func (v *L3Ingress) IPv4OptionsProfile() int {
	return int(v.ip4_options_profile_id)
}

func (v *L3Ingress) SetIPv4OptionsProfile(profile int) {
	v.ip4_options_profile_id = C.int(profile)
}

func (v *L3Ingress) NATRealm() int {
	return int(v.nat_realm_id)
}

func (v *L3Ingress) SetNATRealm(realm int) {
	v.nat_realm_id = C.int(realm)
}

func (v *L3Ingress) TuNneltermecnmap() int {
	return int(v.tunnel_term_ecn_map_id)
}

func (v *L3Ingress) SetTunnelTermEcnMap(ecnMap int) {
	v.tunnel_term_ecn_map_id = C.int(ecnMap)
}

func (v *L3Ingress) IfaceClassRouteDisable() uint32 {
	return uint32(v.intf_class_route_disable)
}

func (v *L3Ingress) SetIfaceClassRouteDisable(m uint32) {
	v.intf_class_route_disable = C.uint32(m)
}

//
// API
//
func NewL3Ingress() *L3Ingress {
	l3ing := &L3Ingress{}
	l3ing.Init()
	return l3ing
}

func L3IngressInit(v *L3Ingress) {
	C.opennsl_l3_ingress_t_init(v.C())
}

func (v *L3Ingress) Init() {
	L3IngressInit(v)
}

func L3IngressCreate(unit int, v *L3Ingress, l3ifaceID L3IfaceID) (L3IfaceID, error) {
	c_l3iface := l3ifaceID.C()
	rc := C.opennsl_l3_ingress_create(C.int(unit), v.C(), &c_l3iface)
	return L3IfaceID(c_l3iface), ParseError(rc)
}

func (v *L3Ingress) Create(unit int, l3ifaceID L3IfaceID) (L3IfaceID, error) {
	return L3IngressCreate(unit, v, l3ifaceID)
}
