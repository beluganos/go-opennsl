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
#include <opennsl/vxlanX.h>
#include "helper.h"
*/
import "C"

//
// VxlanPort
//
type VxlanPort C.opennsl_vxlan_port_t

func (v *VxlanPort) C() *C.opennsl_vxlan_port_t {
	return (*C.opennsl_vxlan_port_t)(v)
}

func (v *VxlanPort) Criteria() VxlanPortMatch {
	return VxlanPortMatch(v.criteria)
}

func (v *VxlanPort) SetCriteria(criteria VxlanPortMatch) {
	v.criteria = criteria.C()
}

func (v *VxlanPort) Egress() L3EgressID {
	return L3EgressID(v.egress_if)
}

func (v *VxlanPort) SetEgress(egressID L3EgressID) {
	v.egress_if = egressID.C()
}

func (v *VxlanPort) EgressServiceTPID() uint16 {
	return uint16(v.egress_service_tpid)
}

func (v *VxlanPort) SetEgressServiceTPID(tpid uint16) {
	v.egress_service_tpid = C.uint16(tpid)
}

func (v *VxlanPort) EgressServiceVlan() Vlan {
	return Vlan(v.egress_service_vlan)
}

func (v *VxlanPort) SetEgressServiceVlan(vid Vlan) {
	v.egress_service_vlan = vid.C()
}

func (v *VxlanPort) EgressTunnelID() TunnelID {
	return TunnelID(v.egress_tunnel_id)
}

func (v *VxlanPort) SetEgressTunnelID(tunID TunnelID) {
	v.egress_tunnel_id = tunID.C()
}

func (v *VxlanPort) Flags() VxlanPortFlags {
	return VxlanPortFlags(v.flags)
}

func (v *VxlanPort) SetFlags(flags VxlanPortFlags) {
	v.flags = flags.C()
}

func (v *VxlanPort) Pri() uint16 {
	return uint16(v.int_pri)
}

func (v *VxlanPort) SetPri(pri uint16) {
	v.int_pri = C.uint16(pri)
}

func (v *VxlanPort) MatchInnerVlan() Vlan {
	return Vlan(v.match_inner_vlan)
}

func (v *VxlanPort) SetMatchInnerVlan(vid Vlan) {
	v.match_inner_vlan = vid.C()
}

func (v *VxlanPort) MatchPort() GPort {
	return GPort(v.match_port)
}

func (v *VxlanPort) SetMatchPort(port GPort) {
	v.match_port = port.C()
}

func (v *VxlanPort) MatchTunnelID() TunnelID {
	return TunnelID(v.match_tunnel_id)
}

func (v *VxlanPort) SetMatchTunnelID(tunID TunnelID) {
	v.match_tunnel_id = tunID.C()
}

func (v *VxlanPort) MatchVlan() Vlan {
	return Vlan(v.match_vlan)
}

func (v *VxlanPort) SetMatchVlan(vid Vlan) {
	v.match_vlan = vid.C()
}

func (v *VxlanPort) MTU() uint16 {
	return uint16(v.mtu)
}

func (v *VxlanPort) SetMTU(mtu uint16) {
	v.mtu = C.uint16(mtu)
}

func (v *VxlanPort) NetworkGroupID() SwitchNetworkGroup {
	return SwitchNetworkGroup(v.network_group_id)
}

func (v *VxlanPort) SetNetworkGroupID(gid SwitchNetworkGroup) {
	v.network_group_id = gid.C()
}

func (v *VxlanPort) PktCfi() uint8 {
	return uint8(v.pkt_cfi)
}

func (v *VxlanPort) SetPktCfi(cfi uint8) {
	v.pkt_cfi = C.uint8(cfi)
}

func (v *VxlanPort) PktPri() uint8 {
	return uint8(v.pkt_pri)
}

func (v *VxlanPort) SetPktPri(pri uint8) {
	v.pkt_pri = C.uint8(pri)
}

func (v *VxlanPort) VxlanPortID() GPort {
	return GPort(v.vxlan_port_id)
}

func (v *VxlanPort) SetVxlanPortID(portID GPort) {
	v.vxlan_port_id = portID.C()
}

//
// API
//
func NewVxlanPort() *VxlanPort {
	port := &VxlanPort{}
	port.Init()
	return port
}

func (v *VxlanPort) Init() {
	C.opennsl_vxlan_port_t_init(v.C())
}

func (v *VxlanPort) Add(unit int, vpn Vpn) error {
	rc := C.opennsl_vxlan_port_add(C.int(unit), vpn.C(), v.C())
	return ParseError(rc)
}

func (v Vpn) VxlanPortAdd(unit int, port *VxlanPort) error {
	return port.Add(unit, v)
}

func VxlanPortDelete(unit int, vpn Vpn, portID GPort) error {
	rc := C.opennsl_vxlan_port_delete(C.int(unit), vpn.C(), portID.C())
	return ParseError(rc)
}

func (v Vpn) VxlanPortDelete(unit int, portID GPort) error {
	return VxlanPortDelete(unit, v, portID)
}

func VxlanPortDeleteAll(unit int, vpn Vpn) error {
	rc := C.opennsl_vxlan_port_delete_all(C.int(unit), vpn.C())
	return ParseError(rc)
}

func (v Vpn) VxlanPortDeleteAll(unit int) error {
	return VxlanPortDeleteAll(unit, v)
}

func VxlanPortGet(unit int, vpn Vpn, portID GPort) (*VxlanPort, error) {
	port := NewVxlanPort()
	port.SetVxlanPortID(portID)

	rc := C.opennsl_vxlan_port_get(C.int(unit), vpn.C(), port.C())
	if err := ParseError(rc); err != nil {
		return nil, err
	}

	return port, nil
}

func (v Vpn) VxlanPortGet(unit int, portID GPort) (*VxlanPort, error) {
	return VxlanPortGet(unit, v, portID)
}

func VxlanPortGetAll(unit int, vpn Vpn, portMax int) ([]*VxlanPort, error) {
	c_ports := make([]C.opennsl_vxlan_port_t, portMax)
	c_count := C.int(0)

	rc := C.opennsl_vxlan_port_get_all(C.int(unit), vpn.C(), C.int(portMax), &c_ports[0], &c_count)
	if err := ParseError(rc); err != nil {
		return nil, err
	}

	ports := make([]*VxlanPort, int(c_count))
	for index := 0; index < int(c_count); index++ {
		ports[index] = (*VxlanPort)(&c_ports[index])
	}

	return ports, nil
}

func (v Vpn) VxlanPortGetAll(unit int, portMax int) ([]*VxlanPort, error) {
	return VxlanPortGetAll(unit, v, portMax)
}
