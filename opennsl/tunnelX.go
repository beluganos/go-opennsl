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
#include <opennsl/tunnelX.h>
*/
import "C"

import (
	"fmt"
	"strings"
)

//
// TunnelID
//
type TunnelID C.opennsl_gport_t

func (v TunnelID) C() C.opennsl_gport_t {
	return C.opennsl_gport_t(v)
}

const TUNNEL TunnelID = 0

//
// TunnelFlags
//
type TunnelFlags C.uint32

func (v TunnelFlags) C() C.uint32 {
	return C.uint32(v)
}

func NewTunnelFlags(flags ...TunnelFlags) TunnelFlags {
	v := TUNNEL_NONE
	for _, flag := range flags {
		v |= flag
	}
	return v
}

const (
	TUNNEL_NONE                TunnelFlags = 0
	TUNNEL_TERM_TUNNEL_WITH_ID TunnelFlags = C.OPENNSL_TUNNEL_TERM_TUNNEL_WITH_ID
	TUNNEL_INIT_USE_INNER_DF   TunnelFlags = C.OPENNSL_TUNNEL_INIT_USE_INNER_DF
	TUNNEL_REPLACE             TunnelFlags = C.OPENNSL_TUNNEL_REPLACE
	TUNNEL_WITH_ID             TunnelFlags = C.OPENNSL_TUNNEL_WITH_ID
)

var tunnelFlags_names = map[TunnelFlags]string{
	TUNNEL_TERM_TUNNEL_WITH_ID: "TERM_TUNNEL_WITH_ID",
	TUNNEL_INIT_USE_INNER_DF:   "INIT_USE_INNER_DF",
	TUNNEL_REPLACE:             "REPLACE",
	TUNNEL_WITH_ID:             "WITH_ID",
}

var tunnelFlags_values = map[string]TunnelFlags{
	"TERM_TUNNEL_WITH_ID": TUNNEL_TERM_TUNNEL_WITH_ID,
	"INIT_USE_INNER_DF":   TUNNEL_INIT_USE_INNER_DF,
	"REPLACE":             TUNNEL_REPLACE,
	"WITH_ID":             TUNNEL_WITH_ID,
}

func (v TunnelFlags) String() string {
	names := []string{}
	for flag, name := range tunnelFlags_names {
		if v&flag != 0 {
			names = append(names, name)
		}
	}
	return strings.Join(names, "|")
}

func ParseTunnelFlags(s string) (TunnelFlags, error) {
	if v, ok := tunnelFlags_values[s]; ok {
		return v, nil
	}
	return TUNNEL_NONE, fmt.Errorf("Invalid TunnelFlags. %s", s)
}

//
// TunnelType
//
type TunnelType C.opennsl_tunnel_type_t

func (v TunnelType) C() C.opennsl_tunnel_type_t {
	return C.opennsl_tunnel_type_t(v)
}

const (
	TunnelTypeNone       TunnelType = 0
	TunnelTypeVxlan      TunnelType = C.opennslTunnelTypeVxlan
	TunnelTypeIPIP4toIP4 TunnelType = 1
	TunnelTypeIPIP4toIP6 TunnelType = 2
	TunnelTypeIPIP4encap TunnelType = 3
	TunnelTypeIPIP6toIP4 TunnelType = 4
	TunnelTypeIPIP6toIP6 TunnelType = 5
	TunnelTypeIPIP6encap TunnelType = 6
	TunnelTypeGRE4toIP4  TunnelType = 7
	TunnelTypeGRE4toIP6  TunnelType = 8
	TunnelTypeGRE4encap  TunnelType = 9
	TunnelTypeGRE6toIP4  TunnelType = 10
	TunnelTypeGRE6toIP6  TunnelType = 11
	TunnelTypeGRE6encap  TunnelType = 12
	TunnelTypePIM4       TunnelType = 18
	TunnelTypePIM4_1     TunnelType = 19
	TunnelTypePIM6       TunnelType = 20
	TunnelTypePIM6_1     TunnelType = 21
)

var tunnelType_names = map[TunnelType]string{
	TunnelTypeNone:       "None",
	TunnelTypeVxlan:      "VxLAN",
	TunnelTypeIPIP4toIP4: "IPIP4toIP4",
	TunnelTypeIPIP4toIP6: "IPIP4toIP6",
	TunnelTypeIPIP4encap: "IPIP4encap",
	TunnelTypeIPIP6toIP4: "IPIP6toIP4",
	TunnelTypeIPIP6toIP6: "IPIP6toIP6",
	TunnelTypeIPIP6encap: "IPIP6encap",
	TunnelTypeGRE4toIP4:  "GRE4toIP4",
	TunnelTypeGRE4toIP6:  "GRE4toIP6",
	TunnelTypeGRE4encap:  "GRE4encap",
	TunnelTypeGRE6toIP4:  "GRE6toIP4",
	TunnelTypeGRE6toIP6:  "GRE6toIP6",
	TunnelTypeGRE6encap:  "GRE6encap",
	TunnelTypePIM4:       "PIM4",
	TunnelTypePIM4_1:     "PIM4_1",
	TunnelTypePIM6:       "PIM6",
	TunnelTypePIM6_1:     "PIM6_1",
}

var tunnelType_values = map[string]TunnelType{
	"None":       TunnelTypeNone,
	"VxLAN":      TunnelTypeVxlan,
	"IPIP4toIP4": TunnelTypeIPIP4toIP4,
	"IPIP4toIP6": TunnelTypeIPIP4toIP6,
	"IPIP4encap": TunnelTypeIPIP4encap,
	"IPIP6toIP4": TunnelTypeIPIP6toIP4,
	"IPIP6toIP6": TunnelTypeIPIP6toIP6,
	"IPIP6encap": TunnelTypeIPIP6encap,
	"GRE4toIP4":  TunnelTypeGRE4toIP4,
	"GRE4toIP6":  TunnelTypeGRE4toIP6,
	"GRE4encap":  TunnelTypeGRE4encap,
	"GRE6toIP4":  TunnelTypeGRE6toIP4,
	"GRE6toIP6":  TunnelTypeGRE6toIP6,
	"GRE6encap":  TunnelTypeGRE6encap,
	"PIM4":       TunnelTypePIM4,
	"PIM4_1":     TunnelTypePIM4_1,
	"PIM6":       TunnelTypePIM6,
	"PIM6_1":     TunnelTypePIM6_1,
}

func (v TunnelType) String() string {
	if s, ok := tunnelType_names[v]; ok {
		return s
	}
	return fmt.Sprintf("TunnelType(%d)", v)
}

func ParseTunnelType(s string) (TunnelType, error) {
	if v, ok := tunnelType_values[s]; ok {
		return v, nil
	}
	return TunnelTypeNone, fmt.Errorf("Invalid TunnelType. %s", s)
}

//
// TunnelDSCPSelect
//
type TunnelDSCPSelect C.opennsl_tunnel_dscp_select_t

func (v TunnelDSCPSelect) C() C.opennsl_tunnel_dscp_select_t {
	return C.opennsl_tunnel_dscp_select_t(v)
}

const (
	TunnelDscpNone   TunnelDSCPSelect = 0
	TunnelDscpAssign TunnelDSCPSelect = C.opennslTunnelDscpAssign
	TunnelDscpPacket TunnelDSCPSelect = C.opennslTunnelDscpPacket
	TunnelDscpMap    TunnelDSCPSelect = C.opennslTunnelDscpMap
	TunnelDscpCount  TunnelDSCPSelect = C.opennslTunnelDscpCount
)

var tunnelDSCPSelect_names = map[TunnelDSCPSelect]string{
	TunnelDscpAssign: "Assign",
	TunnelDscpPacket: "Packet",
	TunnelDscpMap:    "Map",
	TunnelDscpCount:  "Count",
}

var tunnelDSCPSelect_values = map[string]TunnelDSCPSelect{
	"Assign": TunnelDscpAssign,
	"Packet": TunnelDscpPacket,
	"Map":    TunnelDscpMap,
	"Count":  TunnelDscpCount,
}

func (v TunnelDSCPSelect) String() string {
	if s, ok := tunnelDSCPSelect_names[v]; ok {
		return s
	}
	return fmt.Sprintf("TunnelDSCPSelect(%d)", v)
}

func ParseTunnelDSCPSelect(s string) (TunnelDSCPSelect, error) {
	if s, ok := tunnelDSCPSelect_values[s]; ok {
		return s, nil
	}
	return TunnelDscpNone, fmt.Errorf("Invalid TunnelDSCPSelect. %s", s)
}
