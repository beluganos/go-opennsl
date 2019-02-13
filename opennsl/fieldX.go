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
#include "helper.h"
*/
import "C"

import (
	"fmt"
	"strings"
)

//
// FieldColor
//
type FieldColor C.uint8

func (v FieldColor) C() C.uint8 {
	return C.uint8(v)
}

const (
	FIELD_COLOR_NONE     FieldColor = 0
	FIELD_COLOR_PRESERVE FieldColor = C.OPENNSL_FIELD_COLOR_PRESERVE
	FIELD_COLOR_GREEN    FieldColor = C.OPENNSL_FIELD_COLOR_GREEN
	FIELD_COLOR_YELLOW   FieldColor = C.OPENNSL_FIELD_COLOR_YELLOW
	FIELD_COLOR_RED      FieldColor = C.OPENNSL_FIELD_COLOR_RED
	FIELD_COLOR_BLACK    FieldColor = C.OPENNSL_FIELD_COLOR_BLACK
)

var fieldColor_names = map[FieldColor]string{
	FIELD_COLOR_PRESERVE: "PRESERVE",
	FIELD_COLOR_GREEN:    "GREEN",
	FIELD_COLOR_YELLOW:   "YELLOW",
	FIELD_COLOR_RED:      "RED",
	FIELD_COLOR_BLACK:    "BLACK",
}

var fieldColor_values = map[string]FieldColor{
	"PRESERVE": FIELD_COLOR_PRESERVE,
	"GREEN":    FIELD_COLOR_GREEN,
	"YELLOW":   FIELD_COLOR_YELLOW,
	"RED":      FIELD_COLOR_RED,
	"BLACK":    FIELD_COLOR_BLACK,
}

func (v FieldColor) String() string {
	if s, ok := fieldColor_names[v]; ok {
		return s
	}
	return fmt.Sprintf("FieldColor(%d)", v)
}

func ParseFieldColor(s string) (FieldColor, error) {
	if v, ok := fieldColor_values[s]; ok {
		return v, nil
	}
	return FIELD_COLOR_NONE, fmt.Errorf("Invalid FieldColor. %s", s)
}

//
// FieldTcpControl
//
type FieldTcpControl C.uint8

func (v FieldTcpControl) C() C.uint8 {
	return C.uint8(v)
}

const (
	FIELD_TCPCONTROL_NONE FieldTcpControl = 0
	FIELD_TCPCONTROL_FIN  FieldTcpControl = C.OPENNSL_FIELD_TCPCONTROL_FIN
	FIELD_TCPCONTROL_SYN  FieldTcpControl = C.OPENNSL_FIELD_TCPCONTROL_SYN
	FIELD_TCPCONTROL_RST  FieldTcpControl = C.OPENNSL_FIELD_TCPCONTROL_RST
	FIELD_TCPCONTROL_PSH  FieldTcpControl = C.OPENNSL_FIELD_TCPCONTROL_PSH
	FIELD_TCPCONTROL_ACK  FieldTcpControl = C.OPENNSL_FIELD_TCPCONTROL_ACK
	FIELD_TCPCONTROL_URG  FieldTcpControl = C.OPENNSL_FIELD_TCPCONTROL_URG
	FIELD_TCPCONTROL_R40  FieldTcpControl = C.OPENNSL_FIELD_TCPCONTROL_R40
	FIELD_TCPCONTROL_R80  FieldTcpControl = C.OPENNSL_FIELD_TCPCONTROL_R80
	FIELD_TCPCONTROL_MASK FieldTcpControl = C.OPENNSL_FIELD_TCPCONTROL_MASK
)

var fieldTcpControl_names = map[FieldTcpControl]string{
	FIELD_TCPCONTROL_NONE: "NONE",
	FIELD_TCPCONTROL_FIN:  "FIN",
	FIELD_TCPCONTROL_SYN:  "SYN",
	FIELD_TCPCONTROL_RST:  "RST",
	FIELD_TCPCONTROL_PSH:  "PSH",
	FIELD_TCPCONTROL_ACK:  "ACK",
	FIELD_TCPCONTROL_URG:  "URG",
	FIELD_TCPCONTROL_R40:  "R40",
	FIELD_TCPCONTROL_R80:  "R80",
	FIELD_TCPCONTROL_MASK: "MASK",
}

var fieldTcpControl_values = map[string]FieldTcpControl{
	"FIN":  FIELD_TCPCONTROL_FIN,
	"SYN":  FIELD_TCPCONTROL_SYN,
	"RST":  FIELD_TCPCONTROL_RST,
	"PSH":  FIELD_TCPCONTROL_PSH,
	"ACK":  FIELD_TCPCONTROL_ACK,
	"URG":  FIELD_TCPCONTROL_URG,
	"R40":  FIELD_TCPCONTROL_R40,
	"R80":  FIELD_TCPCONTROL_R80,
	"MASK": FIELD_TCPCONTROL_MASK,
}

func (v FieldTcpControl) String() string {
	names := make([]string, 0, len(fieldTcpControl_names))
	for val, name := range fieldTcpControl_names {
		if val&v != 0 {
			names = append(names, name)
		}
	}
	return strings.Join(names, "|")
}

func ParseFieldTcpControl(s string) (FieldTcpControl, error) {
	if v, ok := fieldTcpControl_values[s]; ok {
		return v, nil
	}
	return FIELD_TCPCONTROL_NONE, fmt.Errorf("Invalid FieldTcpControl. %s", s)
}

//
// FieldPktRes
//
type FieldPktRes C.uint8

func (v FieldPktRes) C() C.uint8 {
	return C.uint8(v)
}

const (
	FIELD_PKT_RES_UNKNOWN     FieldPktRes = C.OPENNSL_FIELD_PKT_RES_UNKNOWN
	FIELD_PKT_RES_CONTROL     FieldPktRes = C.OPENNSL_FIELD_PKT_RES_CONTROL
	FIELD_PKT_RES_L2BC        FieldPktRes = C.OPENNSL_FIELD_PKT_RES_L2BC
	FIELD_PKT_RES_L2UC        FieldPktRes = C.OPENNSL_FIELD_PKT_RES_L2UC
	FIELD_PKT_RES_L3MCUNKNOWN FieldPktRes = C.OPENNSL_FIELD_PKT_RES_L3MCUNKNOWN
	FIELD_PKT_RES_L3MCKNOWN   FieldPktRes = C.OPENNSL_FIELD_PKT_RES_L3MCKNOWN
	FIELD_PKT_RES_L3UCKNOWN   FieldPktRes = C.OPENNSL_FIELD_PKT_RES_L3UCKNOWN
	FIELD_PKT_RES_L3UCUNKNOWN FieldPktRes = C.OPENNSL_FIELD_PKT_RES_L3UCUNKNOWN
)

const FIELD_USER_NUM_UDFS int = C.OPENNSL_FIELD_USER_NUM_UDFS

const FIELD_EXACT_MATCH_MASK = C.OPENNSL_FIELD_EXACT_MATCH_MASK

var fieldPktRes_names = map[FieldPktRes]string{
	FIELD_PKT_RES_UNKNOWN:     "UNKNOWN",
	FIELD_PKT_RES_CONTROL:     "CONTROL",
	FIELD_PKT_RES_L2BC:        "L2BC",
	FIELD_PKT_RES_L2UC:        "L2UC",
	FIELD_PKT_RES_L3MCUNKNOWN: "L3MCUNKNOWN",
	FIELD_PKT_RES_L3MCKNOWN:   "L3MCKNOWN",
	FIELD_PKT_RES_L3UCKNOWN:   "L3UCKNOWN",
	FIELD_PKT_RES_L3UCUNKNOWN: "L3UCUNKNOWN",
}

var fieldPktRes_values = map[string]FieldPktRes{
	"UNKNOWN":     FIELD_PKT_RES_UNKNOWN,
	"CONTROL":     FIELD_PKT_RES_CONTROL,
	"L2BC":        FIELD_PKT_RES_L2BC,
	"L2UC":        FIELD_PKT_RES_L2UC,
	"L3MCUNKNOWN": FIELD_PKT_RES_L3MCUNKNOWN,
	"L3MCKNOWN":   FIELD_PKT_RES_L3MCKNOWN,
	"L3UCKNOWN":   FIELD_PKT_RES_L3UCKNOWN,
	"L3UCUNKNOWN": FIELD_PKT_RES_L3UCUNKNOWN,
}

func (v FieldPktRes) String() string {
	if s, ok := fieldPktRes_names[v]; ok {
		return s
	}
	return fmt.Sprintf("FieldPktRes(%d)", v)
}

func ParseFieldPktRes(s string) (FieldPktRes, error) {
	if v, ok := fieldPktRes_values[s]; ok {
		return v, nil
	}
	return FIELD_PKT_RES_UNKNOWN, fmt.Errorf("Invalid FieldPktRes. %s", s)
}

//
// FieldRangeFlags
//
type FieldRangeFlags C.uint32

func (v FieldRangeFlags) C() C.uint32 {
	return C.uint32(v)
}

func NewFieldRangeFlags(flags ...FieldRangeFlags) FieldRangeFlags {
	v := FieldRangeFlags(0)
	for _, flag := range flags {
		v |= flag
	}
	return v
}

const (
	FIELD_RANGE_NONE          FieldRangeFlags = 0
	FIELD_RANGE_SRCPORT       FieldRangeFlags = C.OPENNSL_FIELD_RANGE_SRCPORT
	FIELD_RANGE_DSTPORT       FieldRangeFlags = C.OPENNSL_FIELD_RANGE_DSTPORT
	FIELD_RANGE_TCP           FieldRangeFlags = C.OPENNSL_FIELD_RANGE_TCP
	FIELD_RANGE_UDP           FieldRangeFlags = C.OPENNSL_FIELD_RANGE_UDP
	FIELD_RANGE_INVERT        FieldRangeFlags = C.OPENNSL_FIELD_RANGE_INVERT
	FIELD_RANGE_OUTER_VLAN    FieldRangeFlags = C.OPENNSL_FIELD_RANGE_OUTER_VLAN
	FIELD_RANGE_INNER_VLAN    FieldRangeFlags = C.OPENNSL_FIELD_RANGE_INNER_VLAN
	FIELD_RANGE_PACKET_LENGTH FieldRangeFlags = C.OPENNSL_FIELD_RANGE_PACKET_LENGTH
	FIELD_RANGE_REPLACE       FieldRangeFlags = C.OPENNSL_FIELD_RANGE_REPLACE
)

func FieldInit(unit int) error {
	rc := C.opennsl_field_init(C.int(unit))
	return ParseError(rc)
}

func FieldDetach(unit int) error {
	rc := C.opennsl_field_detach(C.int(unit))
	return ParseError(rc)
}

var fieldRangeFlags_names = map[FieldRangeFlags]string{
	FIELD_RANGE_NONE:          "NONE",
	FIELD_RANGE_SRCPORT:       "SRCPORT",
	FIELD_RANGE_DSTPORT:       "DSTPORT",
	FIELD_RANGE_TCP:           "TCP",
	FIELD_RANGE_UDP:           "UDP",
	FIELD_RANGE_INVERT:        "INVERT",
	FIELD_RANGE_OUTER_VLAN:    "OUTER_VLAN",
	FIELD_RANGE_INNER_VLAN:    "INNER_VLAN",
	FIELD_RANGE_PACKET_LENGTH: "PACKET_LENGTH",
	FIELD_RANGE_REPLACE:       "REPLACE",
}

var fieldRangeFlags_values = map[string]FieldRangeFlags{
	"SRCPORT":       FIELD_RANGE_SRCPORT,
	"DSTPORT":       FIELD_RANGE_DSTPORT,
	"TCP":           FIELD_RANGE_TCP,
	"UDP":           FIELD_RANGE_UDP,
	"INVERT":        FIELD_RANGE_INVERT,
	"OUTER_VLAN":    FIELD_RANGE_OUTER_VLAN,
	"INNER_VLAN":    FIELD_RANGE_INNER_VLAN,
	"PACKET_LENGTH": FIELD_RANGE_PACKET_LENGTH,
	"REPLACE":       FIELD_RANGE_REPLACE,
}

func (v FieldRangeFlags) String() string {
	names := make([]string, 0, len(fieldRangeFlags_names))
	for val, name := range fieldRangeFlags_names {
		if v&val != 0 {
			names = append(names, name)
		}
	}
	return strings.Join(names, "|")
}

func ParseFieldRangeFlags(s string) (FieldRangeFlags, error) {
	if v, ok := fieldRangeFlags_values[s]; ok {
		return v, nil
	}
	return FIELD_RANGE_NONE, fmt.Errorf("Invalid FieldRangeFlags. %s", s)
}
