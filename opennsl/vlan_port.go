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
#include <opennsl/vlan.h>
*/
import "C"

//
// VlanPort
//
type VlanPort C.opennsl_vlan_port_t

func (v *VlanPort) C() *C.opennsl_vlan_port_t {
	return (*C.opennsl_vlan_port_t)(v)
}

func (v *VlanPort) Criteria() VlanPortMatch {
	return VlanPortMatch(v.criteria)
}

func (v *VlanPort) SetCriteria(criteria VlanPortMatch) {
	v.criteria = criteria.C()
}

func (v *VlanPort) Flags() VlanPortFlags {
	return VlanPortFlags(v.flags)
}

func (v *VlanPort) SetFlags(flags VlanPortFlags) {
	v.flags = flags.C()
}

func (v *VlanPort) VSI() Vlan {
	return Vlan(v.vsi)
}

func (v *VlanPort) SetVSI(vsi Vlan) {
	v.vsi = vsi.C()
}

func (v *VlanPort) MatchVID() Vlan {
	return Vlan(v.match_vlan)
}

func (v *VlanPort) SetMatchVID(matchVid Vlan) {
	v.match_vlan = matchVid.C()
}

func (v *VlanPort) MatchInnerVID() Vlan {
	return Vlan(v.match_inner_vlan)
}

func (v *VlanPort) SetMatchInnerVID(matchInnerVid Vlan) {
	v.match_inner_vlan = matchInnerVid.C()
}

func (v *VlanPort) Port() GPort {
	return GPort(v.port)
}

func (v *VlanPort) SetPort(port GPort) {
	v.port = port.C()
}

func (v *VlanPort) EgressVID() Vlan {
	return Vlan(v.egress_vlan)
}

func (v *VlanPort) SetEgressVID(vid Vlan) {
	v.egress_vlan = vid.C()
}

func (v *VlanPort) VlanPort() GPort {
	return GPort(v.vlan_port_id)
}

func (v *VlanPort) SetVlanPort(port GPort) {
	v.vlan_port_id = port.C()
}

func (v *VlanPort) Init() {
	C.opennsl_vlan_port_t_init(v.C())
}

func (v *VlanPort) Create(unit int) error {
	rc := C.opennsl_vlan_port_create(C.int(unit), v.C())
	return ParseError(rc)
}

func (v *VlanPort) Destroy(unit int) error {
	rc := C.opennsl_vlan_port_destroy(C.int(unit), v.vlan_port_id)
	return ParseError(rc)
}

func VlanPortFindVSI(unit int, port GPort, vid Vlan, vsi Vlan, criteria VlanPortMatch) (*VlanPort, error) {
	v := VlanPort{}
	v.Init()
	v.SetPort(port)
	v.SetMatchVID(vid)
	v.SetVSI(vsi)
	v.SetCriteria(criteria)

	rc := C.opennsl_vlan_port_find(C.int(unit), v.C())
	if err := ParseError(rc); err != nil {
		return nil, err
	}

	return &v, nil
}

func VlanPortFind(unit int, port GPort, vid Vlan, criteria VlanPortMatch) (*VlanPort, error) {
	return VlanPortFindVSI(unit, port, vid, 0, criteria)
}
