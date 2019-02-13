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
#include <opennsl/switch.h>
*/
import "C"

type SwitchPktInfoFlags uint32

func (v SwitchPktInfoFlags) C() C.uint32 {
	return C.uint32(v)
}

func NewSwitchPktInfoFlags(flags ...SwitchPktInfoFlags) SwitchPktInfoFlags {
	var flgs SwitchPktInfoFlags = 0
	for _, flg := range flags {
		flgs |= flg
	}
	return flgs
}

const (
	SWITCH_PKT_INFO_SRC_GPORT            SwitchPktInfoFlags = C.OPENNSL_SWITCH_PKT_INFO_SRC_GPORT
	SWITCH_PKT_INFO_VLAN                 SwitchPktInfoFlags = C.OPENNSL_SWITCH_PKT_INFO_VLAN
	SWITCH_PKT_INFO_ETHERTYPE            SwitchPktInfoFlags = C.OPENNSL_SWITCH_PKT_INFO_ETHERTYPE
	SWITCH_PKT_INFO_SRC_MAC              SwitchPktInfoFlags = C.OPENNSL_SWITCH_PKT_INFO_SRC_MAC
	SWITCH_PKT_INFO_DST_MAC              SwitchPktInfoFlags = C.OPENNSL_SWITCH_PKT_INFO_DST_MAC
	SWITCH_PKT_INFO_SRC_IP               SwitchPktInfoFlags = C.OPENNSL_SWITCH_PKT_INFO_SRC_IP
	SWITCH_PKT_INFO_DST_IP               SwitchPktInfoFlags = C.OPENNSL_SWITCH_PKT_INFO_DST_IP
	SWITCH_PKT_INFO_SRC_IPV6             SwitchPktInfoFlags = C.OPENNSL_SWITCH_PKT_INFO_SRC_IPV6
	SWITCH_PKT_INFO_DST_IPV6             SwitchPktInfoFlags = C.OPENNSL_SWITCH_PKT_INFO_DST_IPV6
	SWITCH_PKT_INFO_PROTOCOL             SwitchPktInfoFlags = C.OPENNSL_SWITCH_PKT_INFO_PROTOCOL
	SWITCH_PKT_INFO_SRC_L4_PORT          SwitchPktInfoFlags = C.OPENNSL_SWITCH_PKT_INFO_SRC_L4_PORT
	SWITCH_PKT_INFO_DST_L4_PORT          SwitchPktInfoFlags = C.OPENNSL_SWITCH_PKT_INFO_DST_L4_PORT
	SWITCH_PKT_INFO_HASH_TRUNK           SwitchPktInfoFlags = C.OPENNSL_SWITCH_PKT_INFO_HASH_TRUNK
	SWITCH_PKT_INFO_HASH_MULTIPATH       SwitchPktInfoFlags = C.OPENNSL_SWITCH_PKT_INFO_HASH_MULTIPATH
	SWITCH_PKT_INFO_HASH_UDP_SOURCE_PORT SwitchPktInfoFlags = C.OPENNSL_SWITCH_PKT_INFO_HASH_UDP_SOURCE_PORT
	SWITCH_PKT_INFO_HASH_LBID            SwitchPktInfoFlags = C.OPENNSL_SWITCH_PKT_INFO_HASH_LBID
	SWITCH_PKT_INFO_HASH_OVERLAY         SwitchPktInfoFlags = C.OPENNSL_SWITCH_PKT_INFO_HASH_OVERLAY
)

//
// SwitchPktHashInfoFwdReason
//
type SwitchPktHashInfoFwdReason int

func (v SwitchPktHashInfoFwdReason) C() C.opennsl_switch_pkt_hash_info_fwd_reason_t {
	return C.opennsl_switch_pkt_hash_info_fwd_reason_t(v)
}

const (
	SwitchPktHashInfoFwdReasonUnicast SwitchPktHashInfoFwdReason = C.opennslSwitchPktHashInfoFwdReasonUnicast
	SwitchPktHashInfoFwdReasonIpmc    SwitchPktHashInfoFwdReason = C.opennslSwitchPktHashInfoFwdReasonIpmc
	SwitchPktHashInfoFwdReasonL2mc    SwitchPktHashInfoFwdReason = C.opennslSwitchPktHashInfoFwdReasonL2mc
	SwitchPktHashInfoFwdReasonBcast   SwitchPktHashInfoFwdReason = C.opennslSwitchPktHashInfoFwdReasonBcast
	SwitchPktHashInfoFwdReasonDlf     SwitchPktHashInfoFwdReason = C.opennslSwitchPktHashInfoFwdReasonDlf
)

//
// HashControl
//
type HashControls int

func (v HashControls) C() C.int {
	return C.int(v)
}

func NewHashControls(ctrls ...HashControls) HashControls {
	var vv HashControls = 0
	for _, ctrl := range ctrls {
		vv |= ctrl
	}
	return vv
}

const (
	HASH_CONTROL_MULTIPATH_L4PORTS HashControls = C.OPENNSL_HASH_CONTROL_MULTIPATH_L4PORTS
	HASH_CONTROL_MULTIPATH_DIP     HashControls = C.OPENNSL_HASH_CONTROL_MULTIPATH_DIP
	HASH_CONTROL_TRUNK_UC_SRCPORT  HashControls = C.OPENNSL_HASH_CONTROL_TRUNK_UC_SRCPORT
	HASH_CONTROL_TRUNK_NUC_DST     HashControls = C.OPENNSL_HASH_CONTROL_TRUNK_NUC_DST
	HASH_CONTROL_TRUNK_NUC_SRC     HashControls = C.OPENNSL_HASH_CONTROL_TRUNK_NUC_SRC
	HASH_CONTROL_ECMP_ENHANCE      HashControls = C.OPENNSL_HASH_CONTROL_ECMP_ENHANCE
	HASH_CONTROL_TRUNK_NUC_ENHANCE HashControls = C.OPENNSL_HASH_CONTROL_TRUNK_NUC_ENHANCE
)

//
// HashFields
//
type HashFields uint32

func (v HashFields) C() C.uint32 {
	return C.uint32(v)
}

func NewHashFields(fields ...HashFields) HashFields {
	var vv HashFields = 0
	for _, field := range fields {
		vv |= field
	}
	return vv
}

const (
	HASH_FIELD_CONFIG_CRC16XOR8  HashFields = C.OPENNSL_HASH_FIELD_CONFIG_CRC16XOR8
	HASH_FIELD_CONFIG_XOR16      HashFields = C.OPENNSL_HASH_FIELD_CONFIG_XOR16
	HASH_FIELD_CONFIG_CRC16CCITT HashFields = C.OPENNSL_HASH_FIELD_CONFIG_CRC16CCITT
	HASH_FIELD_CONFIG_CRC32LO    HashFields = C.OPENNSL_HASH_FIELD_CONFIG_CRC32LO
	HASH_FIELD_CONFIG_CRC32HI    HashFields = C.OPENNSL_HASH_FIELD_CONFIG_CRC32HI
	HASH_FIELD_DSTL4             HashFields = C.OPENNSL_HASH_FIELD_DSTL4
	HASH_FIELD_SRCL4             HashFields = C.OPENNSL_HASH_FIELD_SRCL4
	HASH_FIELD_VLAN              HashFields = C.OPENNSL_HASH_FIELD_VLAN
	HASH_FIELD_IP4DST_LO         HashFields = C.OPENNSL_HASH_FIELD_IP4DST_LO
	HASH_FIELD_IP4DST_HI         HashFields = C.OPENNSL_HASH_FIELD_IP4DST_HI
	HASH_FIELD_IP4SRC_LO         HashFields = C.OPENNSL_HASH_FIELD_IP4SRC_LO
	HASH_FIELD_IP4SRC_HI         HashFields = C.OPENNSL_HASH_FIELD_IP4SRC_HI
	HASH_FIELD_IP6DST_LO         HashFields = C.OPENNSL_HASH_FIELD_IP6DST_LO
	HASH_FIELD_IP6DST_HI         HashFields = C.OPENNSL_HASH_FIELD_IP6DST_HI
	HASH_FIELD_IP6SRC_LO         HashFields = C.OPENNSL_HASH_FIELD_IP6SRC_LO
	HASH_FIELD_IP6SRC_HI         HashFields = C.OPENNSL_HASH_FIELD_IP6SRC_HI
	HASH_FIELD_MACDA_LO          HashFields = C.OPENNSL_HASH_FIELD_MACDA_LO
	HASH_FIELD_MACDA_MI          HashFields = C.OPENNSL_HASH_FIELD_MACDA_MI
	HASH_FIELD_MACDA_HI          HashFields = C.OPENNSL_HASH_FIELD_MACDA_HI
	HASH_FIELD_MACSA_LO          HashFields = C.OPENNSL_HASH_FIELD_MACSA_LO
	HASH_FIELD_MACSA_MI          HashFields = C.OPENNSL_HASH_FIELD_MACSA_MI
	HASH_FIELD_MACSA_HI          HashFields = C.OPENNSL_HASH_FIELD_MACSA_HI
)

//
// Color
//
const (
	COLOR_PRIORITY  = C.OPENNSL_COLOR_PRIORITY
	COLOR_OUTER_CFI = C.OPENNSL_COLOR_OUTER_CFI
)

//
// SwitchNetworkGroup
//
type SwitchNetworkGroup C.opennsl_switch_network_group_t

func (v SwitchNetworkGroup) C() C.opennsl_switch_network_group_t {
	return C.opennsl_switch_network_group_t(v)
}
