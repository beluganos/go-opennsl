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
#include <opennsl/mplsX.h>
*/
import "C"

import (
	"fmt"
)

//
// MplsPortMatch
//
type MplsPortMatch C.opennsl_mpls_port_match_t

func (v MplsPortMatch) C() C.opennsl_mpls_port_match_t {
	return C.opennsl_mpls_port_match_t(v)
}

const (
	MPLS_PORT_MATCH_INVALID                         MplsPortMatch = C.OPENNSL_MPLS_PORT_MATCH_INVALID
	MPLS_PORT_MATCH_NONE                            MplsPortMatch = C.OPENNSL_MPLS_PORT_MATCH_NONE
	MPLS_PORT_MATCH_PORT                            MplsPortMatch = C.OPENNSL_MPLS_PORT_MATCH_PORT
	MPLS_PORT_MATCH_PORT_VLAN                       MplsPortMatch = C.OPENNSL_MPLS_PORT_MATCH_PORT_VLAN
	MPLS_PORT_MATCH_PORT_INNER_VLAN                 MplsPortMatch = C.OPENNSL_MPLS_PORT_MATCH_PORT_INNER_VLAN
	MPLS_PORT_MATCH_PORT_VLAN_STACKED               MplsPortMatch = C.OPENNSL_MPLS_PORT_MATCH_PORT_VLAN_STACKED
	MPLS_PORT_MATCH_VLAN_PRI                        MplsPortMatch = C.OPENNSL_MPLS_PORT_MATCH_VLAN_PRI
	MPLS_PORT_MATCH_LABEL                           MplsPortMatch = C.OPENNSL_MPLS_PORT_MATCH_LABEL
	MPLS_PORT_MATCH_LABEL_PORT                      MplsPortMatch = C.OPENNSL_MPLS_PORT_MATCH_LABEL_PORT
	MPLS_PORT_MATCH_LABEL_VLAN                      MplsPortMatch = C.OPENNSL_MPLS_PORT_MATCH_LABEL_VLAN
	MPLS_PORT_MATCH_PORT_SUBPORT_PKT_VID            MplsPortMatch = C.OPENNSL_MPLS_PORT_MATCH_PORT_SUBPORT_PKT_VID
	MPLS_PORT_MATCH_PORT_SUBPORT_PKT_VID_OUTER_VLAN MplsPortMatch = C.OPENNSL_MPLS_PORT_MATCH_PORT_SUBPORT_PKT_VID_OUTER_VLAN
	MPLS_PORT_MATCH_PORT_SUBPORT_PKT_VID_INNER_VLAN MplsPortMatch = C.OPENNSL_MPLS_PORT_MATCH_PORT_SUBPORT_PKT_VID_INNER_VLAN
	MPLS_PORT_MATCH_SHARE                           MplsPortMatch = C.OPENNSL_MPLS_PORT_MATCH_SHARE
	MPLS_PORT_MATCH_PORT_VLAN_TAG                   MplsPortMatch = C.OPENNSL_MPLS_PORT_MATCH_PORT_VLAN_TAG
	MPLS_PORT_MATCH_PORT_INNER_VLAN_TAG             MplsPortMatch = C.OPENNSL_MPLS_PORT_MATCH_PORT_INNER_VLAN_TAG
)

const MPLS_PORT_MATCH_COUNT = C.OPENNSL_MPLS_PORT_MATCH_COUNT

var mplsPortMatch_names = map[MplsPortMatch]string{
	MPLS_PORT_MATCH_INVALID:                         "INVALID",
	MPLS_PORT_MATCH_NONE:                            "NONE",
	MPLS_PORT_MATCH_PORT:                            "PORT",
	MPLS_PORT_MATCH_PORT_VLAN:                       "PORT_VLAN",
	MPLS_PORT_MATCH_PORT_INNER_VLAN:                 "PORT_INNER_VLAN",
	MPLS_PORT_MATCH_PORT_VLAN_STACKED:               "PORT_VLAN_STACKED",
	MPLS_PORT_MATCH_VLAN_PRI:                        "VLAN_PRI",
	MPLS_PORT_MATCH_LABEL:                           "LABEL",
	MPLS_PORT_MATCH_LABEL_PORT:                      "LABEL_PORT",
	MPLS_PORT_MATCH_LABEL_VLAN:                      "LABEL_VLAN",
	MPLS_PORT_MATCH_PORT_SUBPORT_PKT_VID:            "PORT_SUBPORT_PKT_VID",
	MPLS_PORT_MATCH_PORT_SUBPORT_PKT_VID_OUTER_VLAN: "PORT_SUBPORT_PKT_VID_OUTER_VLAN",
	MPLS_PORT_MATCH_PORT_SUBPORT_PKT_VID_INNER_VLAN: "PORT_SUBPORT_PKT_VID_INNER_VLAN",
	MPLS_PORT_MATCH_SHARE:                           "SHARE",
	MPLS_PORT_MATCH_PORT_VLAN_TAG:                   "PORT_VLAN_TAG",
	MPLS_PORT_MATCH_PORT_INNER_VLAN_TAG:             "PORT_INNER_VLAN_TAG",
}

var mplsPortMatch_values = map[string]MplsPortMatch{
	"INVALID":                         MPLS_PORT_MATCH_INVALID,
	"NONE":                            MPLS_PORT_MATCH_NONE,
	"PORT":                            MPLS_PORT_MATCH_PORT,
	"PORT_VLAN":                       MPLS_PORT_MATCH_PORT_VLAN,
	"PORT_INNER_VLAN":                 MPLS_PORT_MATCH_PORT_INNER_VLAN,
	"PORT_VLAN_STACKED":               MPLS_PORT_MATCH_PORT_VLAN_STACKED,
	"VLAN_PRI":                        MPLS_PORT_MATCH_VLAN_PRI,
	"LABEL":                           MPLS_PORT_MATCH_LABEL,
	"LABEL_PORT":                      MPLS_PORT_MATCH_LABEL_PORT,
	"LABEL_VLAN":                      MPLS_PORT_MATCH_LABEL_VLAN,
	"PORT_SUBPORT_PKT_VID":            MPLS_PORT_MATCH_PORT_SUBPORT_PKT_VID,
	"PORT_SUBPORT_PKT_VID_OUTER_VLAN": MPLS_PORT_MATCH_PORT_SUBPORT_PKT_VID_OUTER_VLAN,
	"PORT_SUBPORT_PKT_VID_INNER_VLAN": MPLS_PORT_MATCH_PORT_SUBPORT_PKT_VID_INNER_VLAN,
	"SHARE":                           MPLS_PORT_MATCH_SHARE,
	"PORT_VLAN_TAG":                   MPLS_PORT_MATCH_PORT_VLAN_TAG,
	"PORT_INNER_VLAN_TAG":             MPLS_PORT_MATCH_PORT_INNER_VLAN_TAG,
}

func (v MplsPortMatch) String() string {
	if s, ok := mplsPortMatch_names[v]; ok {
		return s
	}
	return fmt.Sprintf("MplsPortMatch(%d)", v)
}

func ParseMplsPortMatch(s string) (MplsPortMatch, error) {
	if v, ok := mplsPortMatch_values[s]; ok {
		return v, nil
	}
	return MPLS_PORT_MATCH_INVALID, fmt.Errorf("Invalid MplsPortMatch. %s", s)
}

//
// MplsPortControlChannelType
//
type MplsPortControlChannelType C.opennsl_mpls_port_control_channel_type_t

func (v MplsPortControlChannelType) C() C.opennsl_mpls_port_control_channel_type_t {
	return C.opennsl_mpls_port_control_channel_type_t(v)
}

const (
	MplsPortControlChannelNone        MplsPortControlChannelType = C.opennslMplsPortControlChannelNone
	MplsPortControlChannelAch         MplsPortControlChannelType = C.opennslMplsPortControlChannelAch
	MplsPortControlChannelRouterAlert MplsPortControlChannelType = C.opennslMplsPortControlChannelRouterAlert
	MplsPortControlChannelTtl         MplsPortControlChannelType = C.opennslMplsPortControlChannelTtl
	MplsPortControlChannelGalUnderPw  MplsPortControlChannelType = C.opennslMplsPortControlChannelGalUnderPw
)

var MplsPortControlChannelType_names = map[MplsPortControlChannelType]string{
	MplsPortControlChannelNone:        "None",
	MplsPortControlChannelAch:         "Ach",
	MplsPortControlChannelRouterAlert: "RouterAlert",
	MplsPortControlChannelTtl:         "Ttl",
	MplsPortControlChannelGalUnderPw:  "UnderPw",
}

var MplsPortControlChannelType_values = map[string]MplsPortControlChannelType{
	"None":        MplsPortControlChannelNone,
	"Ach":         MplsPortControlChannelAch,
	"RouterAlert": MplsPortControlChannelRouterAlert,
	"Ttl":         MplsPortControlChannelTtl,
	"GalUnderPw":  MplsPortControlChannelGalUnderPw,
}

func (v MplsPortControlChannelType) String() string {
	if s, ok := MplsPortControlChannelType_names[v]; ok {
		return s
	}
	return fmt.Sprintf("MplsPortControlChannelType(%d)", v)
}

func ParseMplsPortControlChannelType(s string) (MplsPortControlChannelType, error) {
	if v, ok := MplsPortControlChannelType_values[s]; ok {
		return v, nil
	}
	return MplsPortControlChannelNone, fmt.Errorf("Invalid MplsPortControlChannelType. %s", s)
}

//
// MplsPort
//
type MplsPort C.opennsl_mpls_port_t

func (v *MplsPort) C() *C.opennsl_mpls_port_t {
	return (*C.opennsl_mpls_port_t)(v)
}

// TODO: implement MplsPort setter/getter.

func (v *MplsPort) Init() {
	C.opennsl_mpls_port_t_init(v.C())
}
