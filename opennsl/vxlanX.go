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
*/
import "C"

import (
	"fmt"
)

const (
	VXLAN_L4PORT = 4789
)

//
// VNID
//
type VNID C.uint32

func (v VNID) C() C.uint32 {
	return C.uint32(v)
}

//
// VxlanVpnFlags
//
type VxlanVpnFlags C.uint32

func (v VxlanVpnFlags) C() C.uint32 {
	return C.uint32(v)
}

func NewVxlanVpnFlags(flags ...VxlanVpnFlags) VxlanVpnFlags {
	v := VXLAN_VPN_NONE
	for _, flag := range flags {
		v |= flag
	}
	return v
}

const (
	VXLAN_VPN_NONE                VxlanVpnFlags = 0
	VXLAN_VPN_ELINE               VxlanVpnFlags = C.OPENNSL_VXLAN_VPN_ELINE
	VXLAN_VPN_ELAN                VxlanVpnFlags = C.OPENNSL_VXLAN_VPN_ELAN
	VXLAN_VPN_WITH_ID             VxlanVpnFlags = C.OPENNSL_VXLAN_VPN_WITH_ID
	VXLAN_VPN_WITH_VPNID          VxlanVpnFlags = C.OPENNSL_VXLAN_VPN_WITH_VPNID
	VXLAN_VPN_SERVICE_TAGGED      VxlanVpnFlags = C.OPENNSL_VXLAN_VPN_SERVICE_TAGGED
	VXLAN_VPN_SERVICE_VLAN_DELETE VxlanVpnFlags = C.OPENNSL_VXLAN_VPN_SERVICE_VLAN_DELETE
)

var vxlanVpnFlags_names = map[VxlanVpnFlags]string{
	VXLAN_VPN_NONE:                "VXLAN_VPN_NONE",
	VXLAN_VPN_ELINE:               "VXLAN_VPN_ELINE",
	VXLAN_VPN_ELAN:                "VXLAN_VPN_ELAN",
	VXLAN_VPN_WITH_ID:             "VXLAN_VPN_WITH_ID",
	VXLAN_VPN_WITH_VPNID:          "VXLAN_VPN_WITH_VPNID",
	VXLAN_VPN_SERVICE_TAGGED:      "VXLAN_VPN_SERVICE_TAGGED",
	VXLAN_VPN_SERVICE_VLAN_DELETE: "VXLAN_VPN_SERVICE_VLAN_DELETE",
}

var vxlanVpnFlags_values = map[string]VxlanVpnFlags{
	"VXLAN_VPN_NONE":                VXLAN_VPN_NONE,
	"VXLAN_VPN_ELINE":               VXLAN_VPN_ELINE,
	"VXLAN_VPN_ELAN":                VXLAN_VPN_ELAN,
	"VXLAN_VPN_WITH_ID":             VXLAN_VPN_WITH_ID,
	"VXLAN_VPN_WITH_VPNID":          VXLAN_VPN_WITH_VPNID,
	"VXLAN_VPN_SERVICE_TAGGED":      VXLAN_VPN_SERVICE_TAGGED,
	"VXLAN_VPN_SERVICE_VLAN_DELETE": VXLAN_VPN_SERVICE_VLAN_DELETE,
}

func (v VxlanVpnFlags) String() string {
	if s, ok := vxlanVpnFlags_names[v]; ok {
		return s
	}
	return fmt.Sprintf("VxlanVpnFlags(%d)", v)
}

func ParseVxlanVpnFlags(s string) (VxlanVpnFlags, error) {
	if v, ok := vxlanVpnFlags_values[s]; ok {
		return v, nil
	}
	return VXLAN_VPN_NONE, fmt.Errorf("Invalid VxlanVpnFlags. %s", s)
}

//
// VxlanPortFlags
//
type VxlanPortFlags C.uint32

func (v VxlanPortFlags) C() C.uint32 {
	return C.uint32(v)
}

func NewVxlanPortFlags(flags ...VxlanPortFlags) VxlanPortFlags {
	v := VXLAN_PORT_NONE
	for _, flag := range flags {
		v |= flag
	}
	return v
}

const (
	VXLAN_PORT_NONE           VxlanPortFlags = 0
	VXLAN_PORT_WITH_ID        VxlanPortFlags = C.OPENNSL_VXLAN_PORT_WITH_ID
	VXLAN_PORT_NETWORK        VxlanPortFlags = C.OPENNSL_VXLAN_PORT_NETWORK
	VXLAN_PORT_EGRESS_TUNNEL  VxlanPortFlags = C.OPENNSL_VXLAN_PORT_EGRESS_TUNNEL
	VXLAN_PORT_SERVICE_TAGGED VxlanPortFlags = C.OPENNSL_VXLAN_PORT_SERVICE_TAGGED
	VXLAN_PORT_MULTICAST      VxlanPortFlags = C.OPENNSL_VXLAN_PORT_MULTICAST
)

var vxlanPortFlags_names = map[VxlanPortFlags]string{
	VXLAN_PORT_NONE:           "VXLAN_PORT_NONE",
	VXLAN_PORT_WITH_ID:        "VXLAN_PORT_WITH_ID",
	VXLAN_PORT_NETWORK:        "VXLAN_PORT_NETWORK",
	VXLAN_PORT_EGRESS_TUNNEL:  "VXLAN_PORT_EGRESS_TUNNEL",
	VXLAN_PORT_SERVICE_TAGGED: "VXLAN_PORT_SERVICE_TAGGED",
	VXLAN_PORT_MULTICAST:      "VXLAN_PORT_MULTICAST",
}

var vxlanPortFlags_values = map[string]VxlanPortFlags{
	"VXLAN_PORT_NONE":           VXLAN_PORT_NONE,
	"VXLAN_PORT_WITH_ID":        VXLAN_PORT_WITH_ID,
	"VXLAN_PORT_NETWORK":        VXLAN_PORT_NETWORK,
	"VXLAN_PORT_EGRESS_TUNNEL":  VXLAN_PORT_EGRESS_TUNNEL,
	"VXLAN_PORT_SERVICE_TAGGED": VXLAN_PORT_SERVICE_TAGGED,
	"VXLAN_PORT_MULTICAST":      VXLAN_PORT_MULTICAST,
}

func (v VxlanPortFlags) String() string {
	if s, ok := vxlanPortFlags_names[v]; ok {
		return s
	}
	return fmt.Sprintf("VxlanPortFlags(%d)", v)
}

func ParseVxlanPortFlags(s string) (VxlanPortFlags, error) {
	if v, ok := vxlanPortFlags_values[s]; ok {
		return v, nil
	}
	return VXLAN_PORT_NONE, fmt.Errorf("Invalid VxlanPortFlags %s", s)
}

//
// VxlanPortMatch
//
type VxlanPortMatch C.opennsl_vxlan_port_match_t

func (v VxlanPortMatch) C() C.opennsl_vxlan_port_match_t {
	return C.opennsl_vxlan_port_match_t(v)
}

const (
	VXLAN_PORT_MATCH_INVALID   VxlanPortMatch = C.OPENNSL_VXLAN_PORT_MATCH_INVALID
	VXLAN_PORT_MATCH_NONE      VxlanPortMatch = C.OPENNSL_VXLAN_PORT_MATCH_NONE
	VXLAN_PORT_MATCH_PORT      VxlanPortMatch = C.OPENNSL_VXLAN_PORT_MATCH_PORT
	VXLAN_PORT_MATCH_PORT_VLAN VxlanPortMatch = C.OPENNSL_VXLAN_PORT_MATCH_PORT_VLAN
	VXLAN_PORT_MATCH_VN_ID     VxlanPortMatch = C.OPENNSL_VXLAN_PORT_MATCH_VN_ID
)

var vxlanPortMatch_names = map[VxlanPortMatch]string{
	VXLAN_PORT_MATCH_INVALID:   "VXLAN_PORT_MATCH_INVALID",
	VXLAN_PORT_MATCH_NONE:      "VXLAN_PORT_MATCH_NONE",
	VXLAN_PORT_MATCH_PORT:      "VXLAN_PORT_MATCH_PORT",
	VXLAN_PORT_MATCH_PORT_VLAN: "VXLAN_PORT_MATCH_PORT_VLAN",
	VXLAN_PORT_MATCH_VN_ID:     "VXLAN_PORT_MATCH_VN_ID",
}

var vxlanPortMatch_values = map[string]VxlanPortMatch{
	"VXLAN_PORT_MATCH_INVALID":   VXLAN_PORT_MATCH_INVALID,
	"VXLAN_PORT_MATCH_NONE":      VXLAN_PORT_MATCH_NONE,
	"VXLAN_PORT_MATCH_PORT":      VXLAN_PORT_MATCH_PORT,
	"VXLAN_PORT_MATCH_PORT_VLAN": VXLAN_PORT_MATCH_PORT_VLAN,
	"VXLAN_PORT_MATCH_VN_ID":     VXLAN_PORT_MATCH_VN_ID,
}

func (v VxlanPortMatch) String() string {
	if s, ok := vxlanPortMatch_names[v]; ok {
		return s
	}
	return fmt.Sprintf("VxlanPortMatch(%d)", v)
}

func ParseVxlanPortMatch(s string) (VxlanPortMatch, error) {
	if v, ok := vxlanPortMatch_values[s]; ok {
		return v, nil
	}
	return VXLAN_PORT_MATCH_INVALID, fmt.Errorf("Invalid VxlanPortMatch %s", s)
}

//
// VxlanStat
//
type VxlanStat C.opennsl_vxlan_stat_t

func (v VxlanStat) C() C.opennsl_vxlan_stat_t {
	return C.opennsl_vxlan_stat_t(v)
}

const (
	VxlanInPackets  VxlanStat = C.opennslVxlanInPackets
	VxlanOutPackets VxlanStat = C.opennslVxlanOutPackets
	VxlanInBytes    VxlanStat = C.opennslVxlanInBytes
	VxlanOutBytes   VxlanStat = C.opennslVxlanOutBytes
)

var vxlanStat_names = map[VxlanStat]string{
	VxlanInPackets:  "VxlanInPackets",
	VxlanOutPackets: "VxlanOutPackets",
	VxlanInBytes:    "VxlanInBytes",
	VxlanOutBytes:   "VxlanOutBytes",
}

var vxlanStat_values = map[string]VxlanStat{
	"VxlanInPackets":  VxlanInPackets,
	"VxlanOutPackets": VxlanOutPackets,
	"VxlanInBytes":    VxlanInBytes,
	"VxlanOutBytes":   VxlanOutBytes,
}

func (v VxlanStat) String() string {
	if s, ok := vxlanStat_names[v]; ok {
		return s
	}
	return fmt.Sprintf("VxlanStat(%d)", v)
}

func ParseVxlanStat(s string) (VxlanStat, error) {
	if v, ok := vxlanStat_values[s]; ok {
		return v, nil
	}
	return 0, fmt.Errorf("Invalid VxlanStat %s", s)
}

//
// API
//
func VxlanInit(unit int) error {
	rc := C.opennsl_vxlan_init(C.int(unit))
	return ParseError(rc)
}

func VxlanCleanup(unit int) error {
	rc := C.opennsl_vxlan_cleanup(C.int(unit))
	return ParseError(rc)
}
