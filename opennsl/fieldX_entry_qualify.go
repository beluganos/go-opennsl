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
#include <opennsl/fieldX.h>
*/
import "C"

import (
	"net"
)

//
// FieldEntryQualify
//
type FieldEntryQualify C.opennsl_field_entry_t

func (v FieldEntryQualify) C() C.opennsl_field_entry_t {
	return C.opennsl_field_entry_t(v)
}

func (v FieldEntryQualify) Delete(unit int, qualify FieldQualify) error {
	rc := C.opennsl_field_qualifier_delete(C.int(unit), v.C(), qualify.C())
	return ParseError(rc)
}

func (v FieldEntryQualify) InPort(unit int, port Port, mask Port) error {
	rc := C.opennsl_field_qualify_InPort(C.int(unit), v.C(), port.C(), mask.C())
	return ParseError(rc)
}

func (v FieldEntryQualify) OutPort(unit int, port Port, mask Port) error {
	rc := C.opennsl_field_qualify_OutPort(C.int(unit), v.C(), port.C(), mask.C())
	return ParseError(rc)
}
func (v FieldEntryQualify) InPorts(unit int, ports *PBmp, masks *PBmp) error {
	rc := C.opennsl_field_qualify_InPorts(C.int(unit), v.C(), *ports.C(), *masks.C())
	return ParseError(rc)
}

func (v FieldEntryQualify) SrcPort(unit int, modID Module, modMask Module, port Port, mask Port) error {
	rc := C.opennsl_field_qualify_SrcPort(C.int(unit), v.C(), modID.C(), modMask.C(), port.C(), mask.C())
	return ParseError(rc)
}

func (v FieldEntryQualify) DstPort(unit int, modID Module, modMask Module, port Port, mask Port) error {
	rc := C.opennsl_field_qualify_DstPort(C.int(unit), v.C(), modID.C(), modMask.C(), port.C(), mask.C())
	return ParseError(rc)
}

func (v FieldEntryQualify) DstTrunk(unit int, trunk Trunk, mask Trunk) error {
	rc := C.opennsl_field_qualify_DstTrunk(C.int(unit), v.C(), trunk.C(), mask.C())
	return ParseError(rc)
}

func (v FieldEntryQualify) L4SrcPort(unit int, port L4Port, mask L4Port) error {
	rc := C.opennsl_field_qualify_L4SrcPort(C.int(unit), v.C(), port.C(), mask.C())
	return ParseError(rc)
}

func (v FieldEntryQualify) L4DstPort(unit int, port L4Port, mask L4Port) error {
	rc := C.opennsl_field_qualify_L4DstPort(C.int(unit), v.C(), port.C(), mask.C())
	return ParseError(rc)
}

func (v FieldEntryQualify) OuterVlan(unit int, vid Vlan, mask Vlan) error {
	rc := C.opennsl_field_qualify_OuterVlan(C.int(unit), v.C(), vid.C(), mask.C())
	return ParseError(rc)
}

func (v FieldEntryQualify) OuterVlanID(unit int, vid Vlan, mask Vlan) error {
	rc := C.opennsl_field_qualify_OuterVlanId(C.int(unit), v.C(), vid.C(), mask.C())
	return ParseError(rc)
}

func (v FieldEntryQualify) InnerVlanID(unit int, vid Vlan, mask Vlan) error {
	rc := C.opennsl_field_qualify_InnerVlanId(C.int(unit), v.C(), vid.C(), mask.C())
	return ParseError(rc)
}

func (v FieldEntryQualify) EtherType(unit int, etherType Ethertype, mask Ethertype) error {
	rc := C.opennsl_field_qualify_EtherType(C.int(unit), v.C(), etherType.C(), mask.C())
	return ParseError(rc)
}

func (v FieldEntryQualify) DestL3Egress(unit int, iface L3IfaceID) error {
	rc := C.opennsl_field_qualify_DstL3Egress(C.int(unit), v.C(), iface.C())
	return ParseError(rc)
}

func (v FieldEntryQualify) Color(unit int, color Color) error {
	rc := C.opennsl_field_qualify_Color(C.int(unit), v.C(), C.uint8(color.C()))
	return ParseError(rc)
}

func (v FieldEntryQualify) IpProtocol(unit int, proto uint8, mask uint8) error {
	rc := C.opennsl_field_qualify_IpProtocol(C.int(unit), v.C(), C.uint8(proto), C.uint8(mask))
	return ParseError(rc)
}

func (v FieldEntryQualify) PacketRes(unit int, res uint32, mask uint32) error {
	rc := C.opennsl_field_qualify_PacketRes(C.int(unit), v.C(), C.uint32(res), C.uint32(mask))
	return ParseError(rc)
}

func (v FieldEntryQualify) SrcIp(unit int, srcIP net.IP, srcMask net.IPMask) error {
	ip, mask, err := NewIP4AndMask(srcIP, srcMask)
	if err != nil {
		return err
	}

	rc := C.opennsl_field_qualify_SrcIp(C.int(unit), v.C(), ip, mask)
	return ParseError(rc)
}

func (v FieldEntryQualify) DstIp(unit int, dstIP net.IP, dstMask net.IPMask) error {
	ip, mask, err := NewIP4AndMask(dstIP, dstMask)
	if err != nil {
		return err
	}

	rc := C.opennsl_field_qualify_DstIp(C.int(unit), v.C(), ip, mask)
	return ParseError(rc)
}

func (v FieldEntryQualify) DSCP(unit int, dscp uint8, mask uint8) error {
	rc := C.opennsl_field_qualify_DSCP(C.int(unit), v.C(), C.uint8(dscp), C.uint8(mask))
	return ParseError(rc)
}

func (v FieldEntryQualify) TcpControl(unit int, ctrl uint8, mask uint8) error {
	rc := C.opennsl_field_qualify_TcpControl(C.int(unit), v.C(), C.uint8(ctrl), C.uint8(mask))
	return ParseError(rc)
}

func (v FieldEntryQualify) TTL(unit int, ttl uint8, mask uint8) error {
	rc := C.opennsl_field_qualify_Ttl(C.int(unit), v.C(), C.uint8(ttl), C.uint8(mask))
	return ParseError(rc)
}

func (v FieldEntryQualify) RangeCheck(unit int, rng FieldRange, invert int) error {
	rc := C.opennsl_field_qualify_RangeCheck(C.int(unit), v.C(), rng.C(), C.int(invert))
	return ParseError(rc)
}

func (v FieldEntryQualify) SrcIp6(unit int, srcIP net.IP, srcMask net.IPMask) error {
	ip, mask, err := NewIP6AndMask(srcIP, srcMask)
	if err != nil {
		return err
	}

	rc := C.opennsl_field_qualify_SrcIp6(C.int(unit), v.C(), &ip[0], &mask[0])
	return ParseError(rc)
}

func (v FieldEntryQualify) DstIp6(unit int, dstIP net.IP, dstMask net.IPMask) error {
	ip, mask, err := NewIP6AndMask(dstIP, dstMask)
	if err != nil {
		return err
	}

	rc := C.opennsl_field_qualify_DstIp6(C.int(unit), v.C(), &ip[0], &mask[0])
	return ParseError(rc)
}

func (v FieldEntryQualify) Ip6NextHeader(unit int, nh uint8, mask uint8) error {
	rc := C.opennsl_field_qualify_Ip6NextHeader(C.int(unit), v.C(), C.uint8(nh), C.uint8(mask))
	return ParseError(rc)
}

func (v FieldEntryQualify) Ip6HopLimit(unit int, limit uint8, mask uint8) error {
	rc := C.opennsl_field_qualify_Ip6HopLimit(C.int(unit), v.C(), C.uint8(limit), C.uint8(mask))
	return ParseError(rc)
}

func (v FieldEntryQualify) SrcMAC(unit int, mac net.HardwareAddr, mask net.HardwareAddr) error {
	c_mac := NewMAC(mac)
	c_mask := NewMAC(mask)

	rc := C.opennsl_field_qualify_SrcMac(C.int(unit), v.C(), &c_mac[0], &c_mask[0])
	return ParseError(rc)
}

func (v FieldEntryQualify) DstMAC(unit int, mac net.HardwareAddr, mask net.HardwareAddr) error {
	c_mac := NewMAC(mac)
	c_mask := NewMAC(mask)

	rc := C.opennsl_field_qualify_DstMac(C.int(unit), v.C(), &c_mac[0], &c_mask[0])
	return ParseError(rc)
}

func (v FieldEntryQualify) IpType(unit int, ipType FieldIpType) error {
	rc := C.opennsl_field_qualify_IpType(C.int(unit), v.C(), ipType.C())
	return ParseError(rc)
}

func (v FieldEntryQualify) InterfaceClassPort(unit int, port Port, mask Port) error {
	rc := C.opennsl_field_qualify_InterfaceClassPort(C.int(unit), v.C(), C.uint32(port.C()), C.uint32(mask.C()))
	return ParseError(rc)
}

func (v FieldEntryQualify) SrcClassField(unit int, field uint32, mask uint32) error {
	rc := C.opennsl_field_qualify_SrcClassField(C.int(unit), v.C(), C.uint32(field), C.uint32(mask))
	return ParseError(rc)
}

func (v FieldEntryQualify) DstClassField(unit int, field uint32, mask uint32) error {
	rc := C.opennsl_field_qualify_DstClassField(C.int(unit), v.C(), C.uint32(field), C.uint32(mask))
	return ParseError(rc)
}

func (v FieldEntryQualify) IpProtocolCommon(unit int, protocol FieldIpProtocolCommon) error {
	rc := C.opennsl_field_qualify_IpProtocolCommon(C.int(unit), v.C(), protocol.C())
	return ParseError(rc)
}

func (v FieldEntryQualify) L3Routable(unit int, table uint8, mask uint8) error {
	rc := C.opennsl_field_qualify_L3Routable(C.int(unit), v.C(), C.uint8(table), C.uint8(mask))
	return ParseError(rc)
}

func (v FieldEntryQualify) IpFrag(unit int, ipFrag FieldIpFrag) error {
	rc := C.opennsl_field_qualify_IpFrag(C.int(unit), v.C(), ipFrag.C())
	return ParseError(rc)
}

func (v FieldEntryQualify) Vrf(unit int, vrf Vrf, mask Vrf) error {
	rc := C.opennsl_field_qualify_Vrf(C.int(unit), v.C(), C.uint32(vrf.C()), C.uint32(mask.C()))
	return ParseError(rc)
}

func (v FieldEntryQualify) L3Ingress(unit int, iface L3IfaceID, mask L3IfaceIDMask) error {
	rc := C.opennsl_field_qualify_L3Ingress(C.int(unit), v.C(), C.uint32(iface.C()), C.uint32(mask.C()))
	return ParseError(rc)
}

func (v FieldEntryQualify) MyStationHit(unit int, hit uint8, mask uint8) error {
	rc := C.opennsl_field_qualify_MyStationHit(C.int(unit), v.C(), C.uint8(hit), C.uint8(mask))
	return ParseError(rc)
}

func (v FieldEntryQualify) IcmpTypeCode(unit int, code uint16, mask uint16) error {
	rc := C.opennsl_field_qualify_IcmpTypeCode(C.int(unit), v.C(), C.uint16(code), C.uint16(mask))
	return ParseError(rc)
}

func (v FieldEntryQualify) ColorGet(unit int) (Color, error) {
	c_color := C.uint8(0)
	rc := C.opennsl_field_qualify_Color_get(C.int(unit), v.C(), &c_color)
	return Color(c_color), ParseError(rc)
}

func (v FieldEntryQualify) DstL3EgressGet(unit int) (L3IfaceID, error) {
	c_iface := C.opennsl_if_t(0)
	rc := C.opennsl_field_qualify_DstL3Egress_get(C.int(unit), v.C(), &c_iface)
	return L3IfaceID(c_iface), ParseError(rc)
}

func (v FieldEntryQualify) InPortGet(unit int) (Port, Port, error) {
	c_port := C.opennsl_port_t(0)
	c_mask := C.opennsl_port_t(0)
	rc := C.opennsl_field_qualify_InPort_get(C.int(unit), v.C(), &c_port, &c_mask)
	return Port(c_port), Port(c_mask), ParseError(rc)
}

func (v FieldEntryQualify) OutPortGet(unit int) (Port, Port, error) {
	c_port := C.opennsl_port_t(0)
	c_mask := C.opennsl_port_t(0)
	rc := C.opennsl_field_qualify_OutPort_get(C.int(unit), v.C(), &c_port, &c_mask)
	return Port(c_port), Port(c_mask), ParseError(rc)
}

func (v FieldEntryQualify) InPortsGet(unit int) (*PBmp, *PBmp, error) {
	pbmp := PBmp{}
	mask := PBmp{}
	rc := C.opennsl_field_qualify_InPorts_get(C.int(unit), v.C(), pbmp.C(), mask.C())
	if err := ParseError(rc); err != nil {
		return nil, nil, err
	}

	return &pbmp, &mask, nil
}

func (v FieldEntryQualify) SrcPortGet(unit int) (Module, Module, Port, Port, error) {
	c_mod := C.opennsl_module_t(0)
	c_mod_mask := C.opennsl_module_t(0)
	c_port := C.opennsl_port_t(0)
	c_port_mask := C.opennsl_port_t(0)

	rc := C.opennsl_field_qualify_SrcPort_get(C.int(unit), v.C(), &c_mod, &c_mod_mask, &c_port, &c_port_mask)
	return Module(c_mod), Module(c_mod_mask), Port(c_port), Port(c_port_mask), ParseError(rc)
}

func (v FieldEntryQualify) DstPortGet(unit int) (Module, Module, Port, Port, error) {
	c_mod := C.opennsl_module_t(0)
	c_mod_mask := C.opennsl_module_t(0)
	c_port := C.opennsl_port_t(0)
	c_port_mask := C.opennsl_port_t(0)

	rc := C.opennsl_field_qualify_DstPort_get(C.int(unit), v.C(), &c_mod, &c_mod_mask, &c_port, &c_port_mask)
	return Module(c_mod), Module(c_mod_mask), Port(c_port), Port(c_port_mask), ParseError(rc)
}

func (v FieldEntryQualify) DstTrunkGet(unit int) (Trunk, Trunk, error) {
	c_trunk := C.opennsl_trunk_t(0)
	c_mask := C.opennsl_trunk_t(0)

	rc := C.opennsl_field_qualify_DstTrunk_get(C.int(unit), v.C(), &c_trunk, &c_mask)
	return Trunk(c_trunk), Trunk(c_mask), ParseError(rc)
}

func (v FieldEntryQualify) L4SrcPortGet(unit int) (L4Port, L4Port, error) {
	c_port := C.opennsl_l4_port_t(0)
	c_mask := C.opennsl_l4_port_t(0)
	rc := C.opennsl_field_qualify_L4SrcPort_get(C.int(unit), v.C(), &c_port, &c_mask)
	return L4Port(c_port), L4Port(c_mask), ParseError(rc)
}

func (v FieldEntryQualify) L4DstPortGet(unit int) (L4Port, L4Port, error) {
	c_port := C.opennsl_l4_port_t(0)
	c_mask := C.opennsl_l4_port_t(0)
	rc := C.opennsl_field_qualify_L4DstPort_get(C.int(unit), v.C(), &c_port, &c_mask)
	return L4Port(c_port), L4Port(c_mask), ParseError(rc)
}

func (v FieldEntryQualify) OuterVlanGet(unit int) (Vlan, Vlan, error) {
	c_vlan := C.opennsl_vlan_t(0)
	c_mask := C.opennsl_vlan_t(0)
	rc := C.opennsl_field_qualify_OuterVlan_get(C.int(unit), v.C(), &c_vlan, &c_mask)
	return Vlan(c_vlan), Vlan(c_mask), ParseError(rc)
}

func (v FieldEntryQualify) OuterVlanIDGet(unit int) (Vlan, Vlan, error) {
	c_vlan := C.opennsl_vlan_t(0)
	c_mask := C.opennsl_vlan_t(0)
	rc := C.opennsl_field_qualify_OuterVlanId_get(C.int(unit), v.C(), &c_vlan, &c_mask)
	return Vlan(c_vlan), Vlan(c_mask), ParseError(rc)
}

func (v FieldEntryQualify) EtherTypeGet(unit int) (Ethertype, Ethertype, error) {
	c_type := C.uint16(0)
	c_mask := C.uint16(0)
	rc := C.opennsl_field_qualify_EtherType_get(C.int(unit), v.C(), &c_type, &c_mask)
	return Ethertype(c_type), Ethertype(c_mask), ParseError(rc)
}

func (v FieldEntryQualify) IpProtocolGet(unit int) (uint8, uint8, error) {
	c_proto := C.uint8(0)
	c_mask := C.uint8(0)
	rc := C.opennsl_field_qualify_IpProtocol_get(C.int(unit), v.C(), &c_proto, &c_mask)
	return uint8(c_proto), uint8(c_mask), ParseError(rc)
}

func (v FieldEntryQualify) PacketResGet(unit int) (uint32, uint32, error) {
	c_res := C.uint32(0)
	c_mask := C.uint32(0)
	rc := C.opennsl_field_qualify_PacketRes_get(C.int(unit), v.C(), &c_res, &c_mask)
	return uint32(c_res), uint32(c_mask), ParseError(rc)
}

func (v FieldEntryQualify) SrcIpGet(unit int) (net.IP, net.IPMask, error) {
	c_ip := C.opennsl_ip_t(0)
	c_mask := C.opennsl_ip_t(0)
	rc := C.opennsl_field_qualify_SrcIp_get(C.int(unit), v.C(), &c_ip, &c_mask)
	if err := ParseError(rc); err != nil {
		return nil, nil, err
	}
	ip, mask := ParseIP4AndMask(c_ip, c_mask)
	return ip, mask, nil
}

func (v FieldEntryQualify) DstIpGet(unit int) (net.IP, net.IPMask, error) {
	c_ip := C.opennsl_ip_t(0)
	c_mask := C.opennsl_ip_t(0)
	rc := C.opennsl_field_qualify_DstIp_get(C.int(unit), v.C(), &c_ip, &c_mask)
	if err := ParseError(rc); err != nil {
		return nil, nil, err
	}
	ip, mask := ParseIP4AndMask(c_ip, c_mask)
	return ip, mask, nil
}

func (v FieldEntryQualify) DSCPGet(unit int) (uint8, uint8, error) {
	c_dscp := C.uint8(0)
	c_mask := C.uint8(0)
	rc := C.opennsl_field_qualify_DSCP_get(C.int(unit), v.C(), &c_dscp, &c_mask)
	return uint8(c_dscp), uint8(c_mask), ParseError(rc)
}

func (v FieldEntryQualify) TcpControlGet(unit int) (uint8, uint8, error) {
	c_ctrl := C.uint8(0)
	c_mask := C.uint8(0)
	rc := C.opennsl_field_qualify_TcpControl_get(C.int(unit), v.C(), &c_ctrl, &c_mask)
	return uint8(c_ctrl), uint8(c_mask), ParseError(rc)
}

func (v FieldEntryQualify) TTLGet(unit int) (uint8, uint8, error) {
	c_ttl := C.uint8(0)
	c_mask := C.uint8(0)
	rc := C.opennsl_field_qualify_Ttl_get(C.int(unit), v.C(), &c_ttl, &c_mask)
	return uint8(c_ttl), uint8(c_mask), ParseError(rc)
}

func (v FieldEntryQualify) RangeCheckGet(unit int, maxCount int) ([]FieldRange, []int, error) {
	c_range := make([]C.opennsl_field_range_t, maxCount)
	c_invert := make([]C.int, maxCount)
	c_count := C.int(0)
	rc := C.opennsl_field_qualify_RangeCheck_get(C.int(unit), v.C(), C.int(maxCount), &c_range[0], &c_invert[0], &c_count)
	if err := ParseError(rc); err != nil {
		return nil, nil, err
	}

	count := int(c_count)
	ranges := make([]FieldRange, count)
	invert := make([]int, count)
	for index := 0; index < count; index++ {
		ranges[index] = FieldRange(c_range[index])
		invert[index] = int(c_invert[index])
	}
	return ranges, invert, nil
}

func (v FieldEntryQualify) SrcIp6Get(unit int) (net.IP, net.IPMask, error) {
	c_ip := C.opennsl_ip6_t{}
	c_mask := C.opennsl_ip6_t{}

	rc := C.opennsl_field_qualify_SrcIp6_get(C.int(unit), v.C(), &c_ip, &c_mask)
	if err := ParseError(rc); err != nil {
		return nil, nil, err
	}
	ip, mask := ParseIP6AndMask(c_ip, c_mask)
	return ip, mask, nil
}

func (v FieldEntryQualify) DstIp6Get(unit int) (net.IP, net.IPMask, error) {
	c_ip := C.opennsl_ip6_t{}
	c_mask := C.opennsl_ip6_t{}

	rc := C.opennsl_field_qualify_DstIp6_get(C.int(unit), v.C(), &c_ip, &c_mask)
	if err := ParseError(rc); err != nil {
		return nil, nil, err
	}
	ip, mask := ParseIP6AndMask(c_ip, c_mask)
	return ip, mask, nil
}

func (v FieldEntryQualify) Ip6NextHeaderGet(unit int) (uint8, uint8, error) {
	c_nh := C.uint8(0)
	c_mask := C.uint8(0)
	rc := C.opennsl_field_qualify_Ip6NextHeader_get(C.int(unit), v.C(), &c_nh, &c_mask)
	return uint8(c_nh), uint8(c_mask), ParseError(rc)
}

func (v FieldEntryQualify) Ip6HopLimitGet(unit int) (uint8, uint8, error) {
	c_limit := C.uint8(0)
	c_mask := C.uint8(0)
	rc := C.opennsl_field_qualify_Ip6HopLimit_get(C.int(unit), v.C(), &c_limit, &c_mask)
	return uint8(c_limit), uint8(c_mask), ParseError(rc)
}

func (v FieldEntryQualify) SrcMACGet(unit int) (net.HardwareAddr, net.HardwareAddr, error) {
	c_mac := C.opennsl_mac_t{}
	c_mask := C.opennsl_mac_t{}

	rc := C.opennsl_field_qualify_SrcMac_get(C.int(unit), v.C(), &c_mac, &c_mask)
	if err := ParseError(rc); err != nil {
		return nil, nil, err
	}

	return ParseMAC(c_mac), ParseMAC(c_mask), nil
}

func (v FieldEntryQualify) DstMACGet(unit int) (net.HardwareAddr, net.HardwareAddr, error) {
	c_mac := C.opennsl_mac_t{}
	c_mask := C.opennsl_mac_t{}

	rc := C.opennsl_field_qualify_DstMac_get(C.int(unit), v.C(), &c_mac, &c_mask)
	if err := ParseError(rc); err != nil {
		return nil, nil, err
	}

	return ParseMAC(c_mac), ParseMAC(c_mask), nil
}

func (v FieldEntryQualify) IpTypeGet(unit int) (FieldIpType, error) {
	c_iptype := C.opennsl_field_IpType_t(0)
	rc := C.opennsl_field_qualify_IpType_get(C.int(unit), v.C(), &c_iptype)
	return FieldIpType(c_iptype), ParseError(rc)
}

func (v FieldEntryQualify) IInterfaceClassPort(unit int) (Port, Port, error) {
	c_port := C.uint32(0)
	c_mask := C.uint32(0)

	rc := C.opennsl_field_qualify_InterfaceClassPort_get(C.int(unit), v.C(), &c_port, &c_mask)
	return Port(c_port), Port(c_mask), ParseError(rc)
}

func (v FieldEntryQualify) SrcClassFieldGet(unit int) (uint32, uint32, error) {
	c_field := C.uint32(0)
	c_mask := C.uint32(0)

	rc := C.opennsl_field_qualify_SrcClassField_get(C.int(unit), v.C(), &c_field, &c_mask)
	return uint32(c_field), uint32(c_mask), ParseError(rc)
}

func (v FieldEntryQualify) DstClassFieldGet(unit int) (uint32, uint32, error) {
	c_field := C.uint32(0)
	c_mask := C.uint32(0)

	rc := C.opennsl_field_qualify_DstClassField_get(C.int(unit), v.C(), &c_field, &c_mask)
	return uint32(c_field), uint32(c_mask), ParseError(rc)
}

func (v FieldEntryQualify) IpProtocolCommonGet(unit int) (FieldIpProtocolCommon, error) {
	c_proto := C.opennsl_field_IpProtocolCommon_t(0)
	rc := C.opennsl_field_qualify_IpProtocolCommon_get(C.int(unit), v.C(), &c_proto)
	return FieldIpProtocolCommon(c_proto), ParseError(rc)
}

func (v FieldEntryQualify) L3RoutableGet(unit int) (uint8, uint8, error) {
	c_table := C.uint8(0)
	c_mask := C.uint8(0)

	rc := C.opennsl_field_qualify_L3Routable_get(C.int(unit), v.C(), &c_table, &c_mask)
	return uint8(c_table), uint8(c_mask), ParseError(rc)
}

func (v FieldEntryQualify) IpFragGet(unit int) (FieldIpFrag, error) {
	c_frag := C.opennsl_field_IpFrag_t(0)

	rc := C.opennsl_field_qualify_IpFrag_get(C.int(unit), v.C(), &c_frag)
	return FieldIpFrag(c_frag), ParseError(rc)
}

func (v FieldEntryQualify) L3IngressGet(unit int) (L3IfaceID, L3IfaceIDMask, error) {
	c_iface := C.uint32(0)
	c_mask := C.uint32(0)

	rc := C.opennsl_field_qualify_L3Ingress_get(C.int(unit), v.C(), &c_iface, &c_mask)
	return L3IfaceID(c_iface), L3IfaceIDMask(c_mask), ParseError(rc)
}

func (v FieldEntryQualify) MyStationHitGet(unit int) (uint8, uint8, error) {
	c_hit := C.uint8(0)
	c_mask := C.uint8(0)

	rc := C.opennsl_field_qualify_MyStationHit_get(C.int(unit), v.C(), &c_hit, &c_mask)
	return uint8(c_hit), uint8(c_mask), ParseError(rc)
}

func (v FieldEntryQualify) IcmpTypeCodeGet(unit int) (uint16, uint16, error) {
	c_code := C.uint16(0)
	c_mask := C.uint16(0)

	rc := C.opennsl_field_qualify_IcmpTypeCode_get(C.int(unit), v.C(), &c_code, &c_mask)
	return uint16(c_code), uint16(c_mask), ParseError(rc)
}

func (v FieldEntryQualify) DstIpLocal(unit int, iplocal uint8, mask uint8) error {
	rc := C.opennsl_field_qualify_DstIpLocal(C.int(unit), v.C(), C.uint8(iplocal), C.uint8(mask))
	return ParseError(rc)
}

func (v FieldEntryQualify) DstIpLocalGet(unit int) (uint8, uint8, error) {
	c_iplocal := C.uint8(0)
	c_mask := C.uint8(0)

	rc := C.opennsl_field_qualify_DstIpLocal_get(C.int(unit), v.C(), &c_iplocal, &c_mask)
	return uint8(c_iplocal), uint8(c_mask), ParseError(rc)
}

func (v FieldEntryQualify) CpuQueue(unit int, queue uint8, mask uint8) error {
	rc := C.opennsl_field_qualify_CpuQueue(C.int(unit), v.C(), C.uint8(queue), C.uint8(mask))
	return ParseError(rc)
}

func (v FieldEntryQualify) CpuQueueGet(unit int) (uint8, uint8, error) {
	c_queue := C.uint8(0)
	c_mask := C.uint8(0)

	rc := C.opennsl_field_qualify_CpuQueue_get(C.int(unit), v.C(), &c_queue, &c_mask)
	return uint8(c_queue), uint8(c_mask), ParseError(rc)
}

func (v FieldEntryQualify) InterfaceClassProcessingPort(unit int, port uint64, mask uint64) error {
	rc := C.opennsl_field_qualify_InterfaceClassProcessingPort(C.int(unit), v.C(), C.uint64(port), C.uint64(mask))
	return ParseError(rc)
}

func (v FieldEntryQualify) InterfaceClassProcessingPortGet(unit int) (uint64, uint64, error) {
	c_port := C.uint64(0)
	c_mask := C.uint64(0)

	rc := C.opennsl_field_qualify_InterfaceClassProcessingPort_get(C.int(unit), v.C(), &c_port, &c_mask)
	return uint64(c_port), uint64(c_mask), ParseError(rc)
}

func (v FieldEntryQualify) IngressClassField(unit int, field uint32, mask uint32) error {
	rc := C.opennsl_field_qualify_IngressClassField(C.int(unit), v.C(), C.uint32(field), C.uint32(mask))
	return ParseError(rc)
}

func (v FieldEntryQualify) IngressClassFieldGet(unit int) (uint32, uint32, error) {
	c_field := C.uint32(0)
	c_mask := C.uint32(0)

	rc := C.opennsl_field_qualify_IngressClassField_get(C.int(unit), v.C(), &c_field, &c_mask)
	return uint32(c_field), uint32(c_mask), ParseError(rc)
}

func (v FieldEntryQualify) Stage(unit int, stage FieldStage) error {
	rc := C.opennsl_field_qualify_Stage(C.int(unit), v.C(), stage.C())
	return ParseError(rc)
}

func (v FieldEntryQualify) StageGet(unit int) (FieldStage, error) {
	c_stage := C.opennsl_field_stage_t(0)

	rc := C.opennsl_field_qualify_Stage_get(C.int(unit), v.C(), &c_stage)
	return FieldStage(c_stage), ParseError(rc)
}
