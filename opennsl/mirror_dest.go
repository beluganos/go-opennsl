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
#include <opennsl/mirror.h>
*/
import "C"

import (
	"net"
)

//
// MirrorNIVFlags
//
type MirrorNIVFlags C.uint32

func (v MirrorNIVFlags) C() C.uint32 {
	return C.uint32(v)
}

func NewMirrorNIVFlags(flags ...MirrorNIVFlags) MirrorNIVFlags {
	v := MIRROR_NIV_NONE
	for _, flag := range flags {
		v |= flag
	}
	return v
}

const (
	MIRROR_NIV_NONE MirrorNIVFlags = 0
	MIRROR_NIV_LOOP MirrorNIVFlags = C.OPENNSL_MIRROR_NIV_LOOP
)

//
// MirrorDestFlags
//
type MirrorDestFlags C.uint32

func (v MirrorDestFlags) C() C.uint32 {
	return C.uint32(v)
}

func NewMirrorDestFlags(flags ...MirrorDestFlags) MirrorDestFlags {
	v := MIRROR_DEST_NONE
	for _, flag := range flags {
		v |= flag
	}
	return v
}

const (
	MIRROR_DEST_NONE                           MirrorDestFlags = 0
	MIRROR_DEST_REPLACE                        MirrorDestFlags = C.OPENNSL_MIRROR_DEST_REPLACE
	MIRROR_DEST_WITH_ID                        MirrorDestFlags = C.OPENNSL_MIRROR_DEST_WITH_ID
	MIRROR_DEST_TUNNEL_L2                      MirrorDestFlags = C.OPENNSL_MIRROR_DEST_TUNNEL_L2
	MIRROR_DEST_TUNNEL_IP_GRE                  MirrorDestFlags = C.OPENNSL_MIRROR_DEST_TUNNEL_IP_GRE
	MIRROR_DEST_PAYLOAD_UNTAGGED               MirrorDestFlags = C.OPENNSL_MIRROR_DEST_PAYLOAD_UNTAGGED
	MIRROR_DEST_TUNNEL_TRILL                   MirrorDestFlags = C.OPENNSL_MIRROR_DEST_TUNNEL_TRILL
	MIRROR_DEST_TUNNEL_NIV                     MirrorDestFlags = C.OPENNSL_MIRROR_DEST_TUNNEL_NIV
	MIRROR_DEST_UPDATE_POLICER                 MirrorDestFlags = C.OPENNSL_MIRROR_DEST_UPDATE_POLICER
	MIRROR_DEST_UPDATE_COUNTER                 MirrorDestFlags = C.OPENNSL_MIRROR_DEST_UPDATE_COUNTER
	MIRROR_DEST_DEST_MULTICAST                 MirrorDestFlags = C.OPENNSL_MIRROR_DEST_DEST_MULTICAST
	MIRROR_DEST_TUNNEL_WITH_ENCAP_ID           MirrorDestFlags = C.OPENNSL_MIRROR_DEST_TUNNEL_WITH_ENCAP_ID
	MIRROR_DEST_TUNNEL_RSPAN                   MirrorDestFlags = C.OPENNSL_MIRROR_DEST_TUNNEL_RSPAN
	MIRROR_DEST_INT_PRI_SET                    MirrorDestFlags = C.OPENNSL_MIRROR_DEST_INT_PRI_SET
	MIRROR_DEST_TUNNEL_ETAG                    MirrorDestFlags = C.OPENNSL_MIRROR_DEST_TUNNEL_ETAG
	MIRROR_DEST_TUNNEL_SFLOW                   MirrorDestFlags = C.OPENNSL_MIRROR_DEST_TUNNEL_SFLOW
	MIRROR_DEST_IS_SNOOP                       MirrorDestFlags = C.OPENNSL_MIRROR_DEST_IS_SNOOP
	MIRROR_DEST_IS_TRAP                        MirrorDestFlags = C.OPENNSL_MIRROR_DEST_IS_TRAP
	MIRROR_DEST_IS_STAT_SAMPLE                 MirrorDestFlags = C.OPENNSL_MIRROR_DEST_IS_STAT_SAMPLE
	MIRROR_DEST_EGRESS_ADD_ORIG_SYSTEM_HEADER  MirrorDestFlags = C.OPENNSL_MIRROR_DEST_EGRESS_ADD_ORIG_SYSTEM_HEADER
	MIRROR_DEST_FIELD                          MirrorDestFlags = C.OPENNSL_MIRROR_DEST_FIELD
	MIRROR_DEST_PORT                           MirrorDestFlags = C.OPENNSL_MIRROR_DEST_PORT
	MIRROR_DEST_DROP_SNIFF_IF_FWD_DROPPED      MirrorDestFlags = C.OPENNSL_MIRROR_DEST_DROP_SNIFF_IF_FWD_DROPPED
	MIRROR_DEST_UPDATE_COUNTER_1               MirrorDestFlags = C.OPENNSL_MIRROR_DEST_UPDATE_COUNTER_1
	MIRROR_DEST_UPDATE_COUNTER_2               MirrorDestFlags = C.OPENNSL_MIRROR_DEST_UPDATE_COUNTER_2
	MIRROR_DEST_DROP_FWD_IF_SNIFF_DROPPED      MirrorDestFlags = C.OPENNSL_MIRROR_DEST_DROP_FWD_IF_SNIFF_DROPPED
	MIRROR_DEST_ID_SHARE                       MirrorDestFlags = C.OPENNSL_MIRROR_DEST_ID_SHARE
	MIRROR_DEST_MTP_ADD                        MirrorDestFlags = C.OPENNSL_MIRROR_DEST_MTP_ADD
	MIRROR_DEST_MTP_DELETE                     MirrorDestFlags = C.OPENNSL_MIRROR_DEST_MTP_DELETE
	MIRROR_DEST_MTP_REPLACE                    MirrorDestFlags = C.OPENNSL_MIRROR_DEST_MTP_REPLACE
	MIRROR_DEST_EGRESS_TRAP_WITH_SYSTEM_HEADER MirrorDestFlags = C.OPENNSL_MIRROR_DEST_EGRESS_TRAP_WITH_SYSTEM_HEADER
	MIRROR_DEST_TUNNEL_PSAMP                   MirrorDestFlags = C.OPENNSL_MIRROR_DEST_TUNNEL_PSAMP
	MIRROR_DEST_OUT_MIRROR_DISABLE             MirrorDestFlags = C.OPENNSL_MIRROR_DEST_OUT_MIRROR_DISABLE
	MIRROR_DEST_TUNNEL_WITH_SEQ                MirrorDestFlags = 1 << 31 // C.OPENNSL_MIRROR_DEST_TUNNEL_WITH_SEQ
	MIRROR_DEST_UPDATE_EXT_COUNTERS            MirrorDestFlags = 1 << 31 // C.OPENNSL_MIRROR_DEST_UPDATE_EXT_COUNTERS
	MIRROR_DEST_FLAGS2_TUNNEL_VXLAN            MirrorDestFlags = C.OPENNSL_MIRROR_DEST_FLAGS2_TUNNEL_VXLAN
)

var mirrorDestFlags_names = map[MirrorDestFlags]string{
	MIRROR_DEST_NONE:                           "MIRROR_DEST_NONE",
	MIRROR_DEST_REPLACE:                        "MIRROR_DEST_REPLACE",
	MIRROR_DEST_WITH_ID:                        "MIRROR_DEST_WITH_ID",
	MIRROR_DEST_TUNNEL_L2:                      "MIRROR_DEST_TUNNEL_L2",
	MIRROR_DEST_TUNNEL_IP_GRE:                  "MIRROR_DEST_TUNNEL_IP_GRE",
	MIRROR_DEST_PAYLOAD_UNTAGGED:               "MIRROR_DEST_PAYLOAD_UNTAGGED",
	MIRROR_DEST_TUNNEL_TRILL:                   "MIRROR_DEST_TUNNEL_TRILL",
	MIRROR_DEST_TUNNEL_NIV:                     "MIRROR_DEST_TUNNEL_NIV",
	MIRROR_DEST_UPDATE_POLICER:                 "MIRROR_DEST_UPDATE_POLICER",
	MIRROR_DEST_UPDATE_COUNTER:                 "MIRROR_DEST_UPDATE_COUNTER",
	MIRROR_DEST_DEST_MULTICAST:                 "MIRROR_DEST_DEST_MULTICAST",
	MIRROR_DEST_TUNNEL_WITH_ENCAP_ID:           "MIRROR_DEST_TUNNEL_WITH_ENCAP_ID",
	MIRROR_DEST_TUNNEL_RSPAN:                   "MIRROR_DEST_TUNNEL_RSPAN",
	MIRROR_DEST_INT_PRI_SET:                    "MIRROR_DEST_INT_PRI_SET",
	MIRROR_DEST_TUNNEL_ETAG:                    "MIRROR_DEST_TUNNEL_ETAG",
	MIRROR_DEST_TUNNEL_SFLOW:                   "MIRROR_DEST_TUNNEL_SFLOW",
	MIRROR_DEST_IS_SNOOP:                       "MIRROR_DEST_IS_SNOOP",
	MIRROR_DEST_IS_TRAP:                        "MIRROR_DEST_IS_TRAP",
	MIRROR_DEST_EGRESS_ADD_ORIG_SYSTEM_HEADER:  "MIRROR_DEST_EGRESS_ADD_ORIG_SYSTEM_HEADER",
	MIRROR_DEST_FIELD:                          "MIRROR_DEST_FIELD",
	MIRROR_DEST_PORT:                           "MIRROR_DEST_PORT",
	MIRROR_DEST_DROP_SNIFF_IF_FWD_DROPPED:      "MIRROR_DEST_DROP_SNIFF_IF_FWD_DROPPED",
	MIRROR_DEST_UPDATE_COUNTER_2:               "MIRROR_DEST_UPDATE_COUNTER_2",
	MIRROR_DEST_ID_SHARE:                       "MIRROR_DEST_ID_SHARE",
	MIRROR_DEST_MTP_ADD:                        "MIRROR_DEST_MTP_ADD",
	MIRROR_DEST_MTP_DELETE:                     "MIRROR_DEST_MTP_DELETE",
	MIRROR_DEST_MTP_REPLACE:                    "MIRROR_DEST_MTP_REPLACE",
	MIRROR_DEST_EGRESS_TRAP_WITH_SYSTEM_HEADER: "MIRROR_DEST_EGRESS_TRAP_WITH_SYSTEM_HEADER",
	MIRROR_DEST_TUNNEL_PSAMP:                   "MIRROR_DEST_TUNNEL_PSAMP",
	MIRROR_DEST_OUT_MIRROR_DISABLE:             "MIRROR_DEST_OUT_MIRROR_DISABLE",
	MIRROR_DEST_TUNNEL_WITH_SEQ:                "MIRROR_DEST_TUNNEL_WITH_SEQ",
	// MIRROR_DEST_IS_STAT_SAMPLE:                 "MIRROR_DEST_IS_STAT_SAMPLE",
	// MIRROR_DEST_UPDATE_COUNTER_1:               "MIRROR_DEST_UPDATE_COUNTER_1",
	// MIRROR_DEST_DROP_FWD_IF_SNIFF_DROPPED:      "MIRROR_DEST_DROP_FWD_IF_SNIFF_DROPPED",
	// MIRROR_DEST_UPDATE_EXT_COUNTERS:            "MIRROR_DEST_UPDATE_EXT_COUNTERS",
	// MIRROR_DEST_FLAGS2_TUNNEL_VXLAN:            "MIRROR_DEST_FLAGS2_TUNNEL_VXLAN",
}

var mirrorDestFlags_values = map[string]MirrorDestFlags{
	"MIRROR_DEST_NONE":                           MIRROR_DEST_NONE,
	"MIRROR_DEST_REPLACE":                        MIRROR_DEST_REPLACE,
	"MIRROR_DEST_WITH_ID":                        MIRROR_DEST_WITH_ID,
	"MIRROR_DEST_TUNNEL_L2":                      MIRROR_DEST_TUNNEL_L2,
	"MIRROR_DEST_TUNNEL_IP_GRE":                  MIRROR_DEST_TUNNEL_IP_GRE,
	"MIRROR_DEST_PAYLOAD_UNTAGGED":               MIRROR_DEST_PAYLOAD_UNTAGGED,
	"MIRROR_DEST_TUNNEL_TRILL":                   MIRROR_DEST_TUNNEL_TRILL,
	"MIRROR_DEST_TUNNEL_NIV":                     MIRROR_DEST_TUNNEL_NIV,
	"MIRROR_DEST_UPDATE_POLICER":                 MIRROR_DEST_UPDATE_POLICER,
	"MIRROR_DEST_UPDATE_COUNTER":                 MIRROR_DEST_UPDATE_COUNTER,
	"MIRROR_DEST_DEST_MULTICAST":                 MIRROR_DEST_DEST_MULTICAST,
	"MIRROR_DEST_TUNNEL_WITH_ENCAP_ID":           MIRROR_DEST_TUNNEL_WITH_ENCAP_ID,
	"MIRROR_DEST_TUNNEL_RSPAN":                   MIRROR_DEST_TUNNEL_RSPAN,
	"MIRROR_DEST_INT_PRI_SET":                    MIRROR_DEST_INT_PRI_SET,
	"MIRROR_DEST_TUNNEL_ETAG":                    MIRROR_DEST_TUNNEL_ETAG,
	"MIRROR_DEST_TUNNEL_SFLOW":                   MIRROR_DEST_TUNNEL_SFLOW,
	"MIRROR_DEST_IS_SNOOP":                       MIRROR_DEST_IS_SNOOP,
	"MIRROR_DEST_IS_TRAP":                        MIRROR_DEST_IS_TRAP,
	"MIRROR_DEST_IS_STAT_SAMPLE":                 MIRROR_DEST_IS_STAT_SAMPLE,
	"MIRROR_DEST_EGRESS_ADD_ORIG_SYSTEM_HEADER":  MIRROR_DEST_EGRESS_ADD_ORIG_SYSTEM_HEADER,
	"MIRROR_DEST_FIELD":                          MIRROR_DEST_FIELD,
	"MIRROR_DEST_PORT":                           MIRROR_DEST_PORT,
	"MIRROR_DEST_DROP_SNIFF_IF_FWD_DROPPED":      MIRROR_DEST_DROP_SNIFF_IF_FWD_DROPPED,
	"MIRROR_DEST_UPDATE_COUNTER_1":               MIRROR_DEST_UPDATE_COUNTER_1,
	"MIRROR_DEST_UPDATE_COUNTER_2":               MIRROR_DEST_UPDATE_COUNTER_2,
	"MIRROR_DEST_DROP_FWD_IF_SNIFF_DROPPED":      MIRROR_DEST_DROP_FWD_IF_SNIFF_DROPPED,
	"MIRROR_DEST_ID_SHARE":                       MIRROR_DEST_ID_SHARE,
	"MIRROR_DEST_MTP_ADD":                        MIRROR_DEST_MTP_ADD,
	"MIRROR_DEST_MTP_DELETE":                     MIRROR_DEST_MTP_DELETE,
	"MIRROR_DEST_MTP_REPLACE":                    MIRROR_DEST_MTP_REPLACE,
	"MIRROR_DEST_EGRESS_TRAP_WITH_SYSTEM_HEADER": MIRROR_DEST_EGRESS_TRAP_WITH_SYSTEM_HEADER,
	"MIRROR_DEST_TUNNEL_PSAMP":                   MIRROR_DEST_TUNNEL_PSAMP,
	"MIRROR_DEST_OUT_MIRROR_DISABLE":             MIRROR_DEST_OUT_MIRROR_DISABLE,
	"MIRROR_DEST_TUNNEL_WITH_SEQ":                MIRROR_DEST_TUNNEL_WITH_SEQ,
	"MIRROR_DEST_UPDATE_EXT_COUNTERS":            MIRROR_DEST_UPDATE_EXT_COUNTERS,
	"MIRROR_DEST_FLAGS2_TUNNEL_VXLAN":            MIRROR_DEST_FLAGS2_TUNNEL_VXLAN,
}

//
// MirrorDestID
//
type MirrorDestID C.opennsl_gport_t

func (v MirrorDestID) C() C.opennsl_gport_t {
	return C.opennsl_gport_t(v)
}

//
// MirrorDest
//
type MirrorDest C.opennsl_mirror_destination_t

func (v *MirrorDest) C() *C.opennsl_mirror_destination_t {
	return (*C.opennsl_mirror_destination_t)(v)
}

func (v *MirrorDest) DF() uint8 {
	return uint8(v.df)
}

func (v *MirrorDest) SetDF(df uint8) {
	v.df = C.uint8(df)
}

func (v *MirrorDest) Dest6Addr() net.IP {
	return ParseIP6(v.dst6_addr)
}

func (v *MirrorDest) SetDest6Addr(ip net.IP) error {
	c_ip, err := NewIP6(ip)
	if err != nil {
		return err
	}

	v.dst6_addr = c_ip
	return nil
}

func (v *MirrorDest) DestAddr() net.IP {
	return ParseIP4(v.dst_addr)
}

func (v *MirrorDest) SetDestAddr(ip net.IP) error {
	c_ip, err := NewIP4(ip)
	if err != nil {
		return err
	}

	v.dst_addr = c_ip
	return nil
}

func (v *MirrorDest) DestMAC() net.HardwareAddr {
	return ParseMAC(v.dst_mac)
}

func (v *MirrorDest) SetDestMAC(mac net.HardwareAddr) {
	v.dst_mac = NewMAC(mac)
}

func (v *MirrorDest) EgressPacketCopySize() uint16 {
	return uint16(v.egress_packet_copy_size)
}

func (v *MirrorDest) SetEgressPacketCopySize(s uint16) {
	v.egress_packet_copy_size = C.uint16(s)
}

func (v *MirrorDest) EncapID() EncapID {
	return EncapID(v.encap_id)
}

func (v *MirrorDest) SetEncapID(id EncapID) {
	v.encap_id = id.C()
}

func (v *MirrorDest) ETagDestVID() Vlan {
	return Vlan(v.etag_dst_vid)
}

func (v *MirrorDest) SetETagDestVID(vid Vlan) {
	v.etag_dst_vid = C.uint16(vid.C())
}

func (v *MirrorDest) ETagSrcVID() Vlan {
	return Vlan(v.etag_src_vid)
}

func (v *MirrorDest) SetETagSrcVID(vid Vlan) {
	v.etag_src_vid = C.uint16(vid.C())
}

func (v *MirrorDest) Flags() MirrorDestFlags {
	return MirrorDestFlags(v.flags)
}

func (v *MirrorDest) SetFlags(flags MirrorDestFlags) {
	v.flags = flags.C()
}

func (v *MirrorDest) FlowLabel() uint32 {
	return uint32(v.flow_label)
}

func (v *MirrorDest) SetFlowLabel(label uint32) {
	v.flow_label = C.uint32(label)
}

func (v *MirrorDest) GPort() GPort {
	return GPort(v.gport)
}

func (v *MirrorDest) SetGPort(gport GPort) {
	v.gport = gport.C()
}

func (v *MirrorDest) GREProtocol() uint16 {
	return uint16(v.gre_protocol)
}

func (v *MirrorDest) SetGREProtocol(proto uint16) {
	v.gre_protocol = C.uint16(proto)
}

func (v *MirrorDest) InternalPri() uint8 {
	return uint8(v.int_pri)
}

func (v *MirrorDest) SetInternalPri(pri uint8) {
	v.int_pri = C.uint8(pri)
}

func (v *MirrorDest) MirrorDestID() MirrorDestID {
	return MirrorDestID(v.mirror_dest_id)
}

func (v *MirrorDest) SetMirrorDestID(id MirrorDestID) {
	v.mirror_dest_id = id.C()
}

func (v *MirrorDest) NIVDestVIface() uint16 {
	return uint16(v.niv_dst_vif)
}

func (v *MirrorDest) SetNIVDestVIface(vif uint16) {
	v.niv_dst_vif = C.uint16(vif)
}

func (v *MirrorDest) NIVSrcVIface() uint16 {
	return uint16(v.niv_src_vif)
}

func (v *MirrorDest) SetNIVSrcVIface(vif uint16) {
	v.niv_src_vif = C.uint16(vif)
}

func (v *MirrorDest) NIVFlags() MirrorNIVFlags {
	return MirrorNIVFlags(v.niv_flags)
}

func (v *MirrorDest) SetNIVFlags(flags MirrorNIVFlags) {
	v.niv_flags = flags.C()
}

func (v *MirrorDest) PacketCopySize() uint16 {
	return uint16(v.packet_copy_size)
}

func (v *MirrorDest) SetPacketCopySize(s uint16) {
	v.packet_copy_size = C.uint16(s)
}

func (v *MirrorDest) PacketPri() uint8 {
	return uint8(v.pkt_prio)
}

func (v *MirrorDest) SetPacketPri(pri uint8) {
	v.pkt_prio = C.uint8(pri)
}

func (v *MirrorDest) PolicerID() Policer {
	return Policer(v.policer_id)
}

func (v *MirrorDest) SetPolicerID(policer Policer) {
	v.policer_id = policer.C()
}

func (v *MirrorDest) RecycleContext() uint8 {
	return uint8(v.recycle_context)
}

func (v *MirrorDest) SetRecycleContext(ctxt uint8) {
	v.recycle_context = C.uint8(ctxt)
}

func (v *MirrorDest) Src6Addr() net.IP {
	return ParseIP6(v.src6_addr)
}

func (v *MirrorDest) SetSrc6Addr(ip net.IP) error {
	c_ip, err := NewIP6(ip)
	if err != nil {
		return err
	}

	v.src6_addr = c_ip
	return nil
}

func (v *MirrorDest) SrcAddr() net.IP {
	return ParseIP4(v.src_addr)
}

func (v *MirrorDest) SetSrcAddr(ip net.IP) error {
	c_ip, err := NewIP4(ip)
	if err != nil {
		return err
	}

	v.src_addr = c_ip
	return nil
}

func (v *MirrorDest) SrcMAC() net.HardwareAddr {
	return ParseMAC(v.src_mac)
}

func (v *MirrorDest) SetSrcMAC(mac net.HardwareAddr) {
	v.src_mac = NewMAC(mac)
}

func (v *MirrorDest) StatID() int {
	return int(v.stat_id)
}

func (v *MirrorDest) SetStatID(statID int) {
	v.stat_id = C.int(statID)
}

func (v *MirrorDest) Tos() uint8 {
	return uint8(v.tos)
}

func (v *MirrorDest) SetTos(tos uint8) {
	v.tos = C.uint8(tos)
}

func (v *MirrorDest) TPID() uint16 {
	return uint16(v.tpid)
}

func (v *MirrorDest) SetTPID(tpid uint16) {
	v.tpid = C.uint16(tpid)
}

func (v *MirrorDest) TTL() uint8 {
	return uint8(v.ttl)
}

func (v *MirrorDest) SetTTL(ttl uint8) {
	v.ttl = C.uint8(ttl)
}

func (v *MirrorDest) TunnelID() TunnelID {
	return TunnelID(v.tunnel_id)
}

func (v *MirrorDest) SetTunnelID(id TunnelID) {
	v.tunnel_id = id.C()
}

func (v *MirrorDest) UDPDestPort() L4Port {
	return L4Port(v.udp_dst_port)
}

func (v *MirrorDest) SetUDPDestPort(port L4Port) {
	v.udp_dst_port = C.uint16(port.C())
}

func (v *MirrorDest) UDPSrcPort() L4Port {
	return L4Port(v.udp_src_port)
}

func (v *MirrorDest) SetUDPSrcPort(port L4Port) {
	v.udp_src_port = C.uint16(port.C())
}

func (v *MirrorDest) IPVer() uint8 {
	return uint8(v.version)
}

func (v *MirrorDest) SetIPVer(ipver uint8) {
	v.version = C.uint8(ipver)
}

func (v *MirrorDest) VID() Vlan {
	return Vlan(v.vlan_id)
}

func (v *MirrorDest) SetVID(vid Vlan) {
	v.vlan_id = vid.C()
}

//
// API
//
func MirrorDestInit(v *MirrorDest) {
	C.opennsl_mirror_destination_t_init(v.C())
}

func (v *MirrorDest) Init() {
	MirrorDestInit(v)
}

func NewMirrorDest() *MirrorDest {
	v := &MirrorDest{}
	v.Init()
	return v
}

func MirrorDestCreate(unit int, v *MirrorDest) (bool, error) {
	rc := C.opennsl_mirror_destination_create(C.int(unit), v.C())
	if OpenNSLError(rc) == E_EXISTS {
		return false, nil
	}
	if err := ParseError(rc); err != nil {
		return false, err
	}

	return true, ParseError(rc)
}

func (v *MirrorDest) Create(unit int) (bool, error) {
	return MirrorDestCreate(unit, v)
}

//
// API
//
func MirrorPortDestAdd(unit int, port Port, flags MirrorDestFlags, mirrorDest MirrorDestID) error {
	rc := C.opennsl_mirror_port_dest_add(C.int(unit), port.C(), flags.C(), mirrorDest.C())
	return ParseError(rc)
}

func (v MirrorDestID) PortAdd(unit int, port Port, flags MirrorDestFlags) error {
	return MirrorPortDestAdd(unit, port, flags, v)
}

func MirrorPortDestDelete(unit int, port Port, flags MirrorDestFlags, mirrorDest MirrorDestID) error {
	rc := C.opennsl_mirror_port_dest_delete(C.int(unit), port.C(), flags.C(), mirrorDest.C())
	return ParseError(rc)
}

func (v MirrorDestID) PortDelete(unit int, port Port, flags MirrorDestFlags) error {
	return MirrorPortDestDelete(unit, port, flags, v)
}

func MirrorPortDestDeleteAll(unit int, port Port, flags MirrorDestFlags) error {
	rc := C.opennsl_mirror_port_dest_delete_all(C.int(unit), port.C(), flags.C())
	return ParseError(rc)
}

func MirrorPortDestCount(unit int, port Port, flags MirrorDestFlags) (int, error) {
	c_count := C.int(0)

	rc := C.opennsl_mirror_port_dest_get(C.int(unit), port.C(), flags.C(), 0, nil, &c_count)
	return int(c_count), ParseError(rc)
}

func MirrorPortDestGet(unit int, port Port, flags MirrorDestFlags, maxSize int) ([]MirrorDestID, error) {
	if maxSize < 0 {
		return nil, E_PARAM.Error()
	}

	if maxSize == 0 {
		count, err := MirrorPortDestCount(unit, port, flags)
		if err != nil {
			return nil, err
		}
		maxSize = count
	}

	c_gports := make([]C.opennsl_gport_t, maxSize)
	c_count := C.int(0)

	rc := C.opennsl_mirror_port_dest_get(C.int(unit), port.C(), flags.C(), C.int(maxSize), &c_gports[0], &c_count)

	if err := ParseError(rc); err != nil {
		return nil, err
	}

	gports := make([]MirrorDestID, int(c_count))
	for index := 0; index < int(c_count); index++ {
		gports[index] = MirrorDestID(c_gports[index])
	}

	return gports, nil
}
