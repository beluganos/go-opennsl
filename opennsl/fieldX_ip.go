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
	"fmt"
)

//
// FieldIpType
//
type FieldIpType C.opennsl_field_IpType_t

func (v FieldIpType) C() C.opennsl_field_IpType_t {
	return C.opennsl_field_IpType_t(v)
}

const (
	FieldIpTypeAny           FieldIpType = C.opennslFieldIpTypeAny
	FieldIpTypeNonIp         FieldIpType = C.opennslFieldIpTypeNonIp
	FieldIpTypeIpv4Not       FieldIpType = C.opennslFieldIpTypeIpv4Not
	FieldIpTypeIpv4NoOpts    FieldIpType = C.opennslFieldIpTypeIpv4NoOpts
	FieldIpTypeIpv4WithOpts  FieldIpType = C.opennslFieldIpTypeIpv4WithOpts
	FieldIpTypeIpv4Any       FieldIpType = C.opennslFieldIpTypeIpv4Any
	FieldIpTypeIpv6Not       FieldIpType = C.opennslFieldIpTypeIpv6Not
	FieldIpTypeIpv6NoExtHdr  FieldIpType = C.opennslFieldIpTypeIpv6NoExtHdr
	FieldIpTypeIpv6OneExtHdr FieldIpType = C.opennslFieldIpTypeIpv6OneExtHdr
	FieldIpTypeIpv6TwoExtHdr FieldIpType = C.opennslFieldIpTypeIpv6TwoExtHdr
	FieldIpTypeIpv6          FieldIpType = C.opennslFieldIpTypeIpv6
	FieldIpTypeIp            FieldIpType = C.opennslFieldIpTypeIp
	FieldIpTypeArp           FieldIpType = C.opennslFieldIpTypeArp
	FieldIpTypeArpRequest    FieldIpType = C.opennslFieldIpTypeArpRequest
	FieldIpTypeArpReply      FieldIpType = C.opennslFieldIpTypeArpReply
)

var fieldIpType_names = map[FieldIpType]string{
	FieldIpTypeAny:           "Any",
	FieldIpTypeNonIp:         "NonIp",
	FieldIpTypeIpv4Not:       "Ipv4Not",
	FieldIpTypeIpv4NoOpts:    "Ipv4NoOpts",
	FieldIpTypeIpv4WithOpts:  "Ipv4WithOpts",
	FieldIpTypeIpv4Any:       "Ipv4Any",
	FieldIpTypeIpv6Not:       "Ipv6Not",
	FieldIpTypeIpv6NoExtHdr:  "Ipv6NoExtHdr",
	FieldIpTypeIpv6OneExtHdr: "Ipv6OneExtHdr",
	FieldIpTypeIpv6TwoExtHdr: "Ipv6TwoExtHdr",
	FieldIpTypeIpv6:          "Ipv6",
	FieldIpTypeIp:            "Ip",
	FieldIpTypeArp:           "Arp",
	FieldIpTypeArpRequest:    "ArpRequest",
	FieldIpTypeArpReply:      "ArpReply",
}

var fieldIpType_values = map[string]FieldIpType{
	"Any":           FieldIpTypeAny,
	"NonIp":         FieldIpTypeNonIp,
	"Ipv4Not":       FieldIpTypeIpv4Not,
	"Ipv4NoOpts":    FieldIpTypeIpv4NoOpts,
	"Ipv4WithOpts":  FieldIpTypeIpv4WithOpts,
	"Ipv4Any":       FieldIpTypeIpv4Any,
	"Ipv6Not":       FieldIpTypeIpv6Not,
	"Ipv6NoExtHdr":  FieldIpTypeIpv6NoExtHdr,
	"Ipv6OneExtHdr": FieldIpTypeIpv6OneExtHdr,
	"Ipv6TwoExtHdr": FieldIpTypeIpv6TwoExtHdr,
	"Ipv6":          FieldIpTypeIpv6,
	"Ip":            FieldIpTypeIp,
	"Arp":           FieldIpTypeArp,
	"ArpRequest":    FieldIpTypeArpRequest,
	"ArpReply":      FieldIpTypeArpReply,
}

func (v FieldIpType) String() string {
	if s, ok := fieldIpType_names[v]; ok {
		return s
	}
	return fmt.Sprintf("FieldIpType(%d)", v)
}

func ParseFieldIpType(s string) (FieldIpType, error) {
	if v, ok := fieldIpType_values[s]; ok {
		return v, nil
	}
	return FieldIpTypeAny, fmt.Errorf("Invalid FieldIpType. %s", s)
}

//
// FieldIpProtocolCommon
//
type FieldIpProtocolCommon C.opennsl_field_IpProtocolCommon_t

func (v FieldIpProtocolCommon) C() C.opennsl_field_IpProtocolCommon_t {
	return C.opennsl_field_IpProtocolCommon_t(v)
}

const (
	FieldIpProtocolCommonTcp  FieldIpProtocolCommon = C.opennslFieldIpProtocolCommonTcp
	FieldIpProtocolCommonUdp  FieldIpProtocolCommon = C.opennslFieldIpProtocolCommonUdp
	FieldIpProtocolCommonIgmp FieldIpProtocolCommon = C.opennslFieldIpProtocolCommonIgmp
	FieldIpProtocolCommonIcmp FieldIpProtocolCommon = C.opennslFieldIpProtocolCommonIcmp
)

var fieldIpProtocolCommon_names = map[FieldIpProtocolCommon]string{
	FieldIpProtocolCommonTcp:  "Tcp",
	FieldIpProtocolCommonUdp:  "Udp",
	FieldIpProtocolCommonIgmp: "Igmp",
	FieldIpProtocolCommonIcmp: "Icmp",
}

var fieldIpProtocolCommon_values = map[string]FieldIpProtocolCommon{
	"Tcp":  FieldIpProtocolCommonTcp,
	"Udp":  FieldIpProtocolCommonUdp,
	"Igmp": FieldIpProtocolCommonIgmp,
	"Icmp": FieldIpProtocolCommonIcmp,
}

func (v FieldIpProtocolCommon) String() string {
	if s, ok := fieldIpProtocolCommon_names[v]; ok {
		return s
	}
	return fmt.Sprintf("FieldIpProtocolCommon(%d)", v)
}

func ParseFieldIpProtocolCommon(s string) (FieldIpProtocolCommon, error) {
	if v, ok := fieldIpProtocolCommon_values[s]; ok {
		return v, nil
	}
	return FieldIpProtocolCommonTcp, fmt.Errorf("Invalid FieldIpProtocolCommon. %s", s)
}

//
// FieldIpFrag
//
type FieldIpFrag C.opennsl_field_IpFrag_t

func (v FieldIpFrag) C() C.opennsl_field_IpFrag_t {
	return C.opennsl_field_IpFrag_t(v)
}

const (
	FieldIpFragNon        FieldIpFrag = C.opennslFieldIpFragNon
	FieldIpFragFirst      FieldIpFrag = C.opennslFieldIpFragFirst
	FieldIpFragNonOrFirst FieldIpFrag = C.opennslFieldIpFragNonOrFirst
	FieldIpFragNotFirst   FieldIpFrag = C.opennslFieldIpFragNotFirst
	FieldIpFragAny        FieldIpFrag = C.opennslFieldIpFragAny
)

const FieldIpFragCount = C.opennslFieldIpFragCount

var fieldIpFrag_names = map[FieldIpFrag]string{
	FieldIpFragNon:        "Non",
	FieldIpFragFirst:      "First",
	FieldIpFragNonOrFirst: "NonOrFirst",
	FieldIpFragNotFirst:   "NotFirst",
	FieldIpFragAny:        "Any",
}

var fieldIpFrag_values = map[string]FieldIpFrag{
	"Non":        FieldIpFragNon,
	"First":      FieldIpFragFirst,
	"NonOrFirst": FieldIpFragNonOrFirst,
	"NotFirst":   FieldIpFragNotFirst,
	"Any":        FieldIpFragAny,
}

func (v FieldIpFrag) String() string {
	if s, ok := fieldIpFrag_names[v]; ok {
		return s
	}
	return fmt.Sprintf("FieldIpFrag(%d)", v)
}

func ParseFieldIpFrag(s string) (FieldIpFrag, error) {
	if v, ok := fieldIpFrag_values[s]; ok {
		return v, nil
	}
	return FieldIpFragNon, fmt.Errorf("Invalid FieldIpFrag. %s", s)
}
