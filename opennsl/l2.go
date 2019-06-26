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
#include <opennsl/l2.h>
*/
import "C"

import (
	"fmt"
	"net"
	"strings"
)

//
// Variables
//
func NewL2AddrMaskExact() net.HardwareAddr {
	mac, _ := net.ParseMAC("ff:ff:ff:ff:ff:ff")
	return mac
}

//
// L2Flags
//
type L2Flags uint32

func (v L2Flags) C() C.uint32 {
	return C.uint32(v)
}

func NewL2Flags(flags ...L2Flags) L2Flags {
	v := L2_NONE
	for _, flag := range flags {
		v |= flag
	}
	return v
}

const (
	L2_NONE         L2Flags = 0
	L2_DISCARD_SRC  L2Flags = C.OPENNSL_L2_DISCARD_SRC
	L2_DISCARD_DST  L2Flags = C.OPENNSL_L2_DISCARD_DST
	L2_COPY_TO_CPU  L2Flags = C.OPENNSL_L2_COPY_TO_CPU
	L2_L3LOOKUP     L2Flags = C.OPENNSL_L2_L3LOOKUP
	L2_STATIC       L2Flags = C.OPENNSL_L2_STATIC
	L2_HIT          L2Flags = C.OPENNSL_L2_HIT
	L2_TRUNK_MEMBER L2Flags = C.OPENNSL_L2_TRUNK_MEMBER
	L2_MCAST        L2Flags = C.OPENNSL_L2_MCAST
	L2_SRC_HIT      L2Flags = C.OPENNSL_L2_SRC_HIT
	L2_DES_HIT      L2Flags = C.OPENNSL_L2_DES_HIT
	L2_MOVE_PORT    L2Flags = C.OPENNSL_L2_MOVE_PORT
	L2_LOCAL_CPU    L2Flags = C.OPENNSL_L2_LOCAL_CPU
)

var l2flags_names = map[L2Flags]string{
	L2_NONE:         "L2_NONE",
	L2_DISCARD_SRC:  "L2_DISCARD_SRC",
	L2_DISCARD_DST:  "L2_DISCARD_DST",
	L2_COPY_TO_CPU:  "L2_COPY_TO_CPU",
	L2_L3LOOKUP:     "L2_L3LOOKUP",
	L2_STATIC:       "L2_STATIC",
	L2_TRUNK_MEMBER: "L2_TRUNK_MEMBER",
	L2_MCAST:        "L2_MCAST",
	L2_SRC_HIT:      "L2_SRC_HIT",
	L2_DES_HIT:      "L2_DES_HIT",
	L2_MOVE_PORT:    "L2_MOVE_PORT",
	L2_LOCAL_CPU:    "L2_LOCAL_CPU",
	// L2_HIT:          "L2_HIT",
}

var l2flags_values = map[string]L2Flags{
	"L2_NONE":         L2_NONE,
	"L2_DISCARD_SRC":  L2_DISCARD_SRC,
	"L2_DISCARD_DST":  L2_DISCARD_DST,
	"L2_COPY_TO_CPU":  L2_COPY_TO_CPU,
	"L2_L3LOOKUP":     L2_L3LOOKUP,
	"L2_STATIC":       L2_STATIC,
	"L2_HIT":          L2_HIT,
	"L2_TRUNK_MEMBER": L2_TRUNK_MEMBER,
	"L2_MCAST":        L2_MCAST,
	"L2_SRC_HIT":      L2_SRC_HIT,
	"L2_DES_HIT":      L2_DES_HIT,
	"L2_MOVE_PORT":    L2_MOVE_PORT,
	"L2_LOCAL_CPU":    L2_LOCAL_CPU,
}

func (v L2Flags) String() string {
	names := make([]string, 0, len(l2flags_names))
	for value, name := range l2flags_names {
		if v&value != 0 {
			names = append(names, name)
		}
	}
	return strings.Join(names, "|")
}

func ParseL2Flags(name string) (L2Flags, error) {
	if v, ok := l2flags_values[name]; ok {
		return v, nil
	}
	return L2_NONE, fmt.Errorf("Invalid L2Flags. %s", name)
}

//
// L2CallbackOper
//
type L2CallbackOper int32

func NewL2CallbackOper(v C.int) L2CallbackOper {
	return L2CallbackOper(v)
}

func (v L2CallbackOper) C() C.int {
	return C.int(v)
}

const (
	L2_CALLBACK_NONE        L2CallbackOper = 0
	L2_CALLBACK_DELETE      L2CallbackOper = C.OPENNSL_L2_CALLBACK_DELETE
	L2_CALLBACK_ADD         L2CallbackOper = C.OPENNSL_L2_CALLBACK_ADD
	L2_CALLBACK_REPORT      L2CallbackOper = C.OPENNSL_L2_CALLBACK_REPORT
	L2_CALLBACK_LEARN_EVENT L2CallbackOper = C.OPENNSL_L2_CALLBACK_LEARN_EVENT
	L2_CALLBACK_AGE_EVENT   L2CallbackOper = C.OPENNSL_L2_CALLBACK_AGE_EVENT
	L2_CALLBACK_MOVE_EVENT  L2CallbackOper = C.OPENNSL_L2_CALLBACK_MOVE_EVENT
)

var l2CallbackOper_names = map[L2CallbackOper]string{
	L2_CALLBACK_DELETE:      "DELETE",
	L2_CALLBACK_ADD:         "ADD",
	L2_CALLBACK_REPORT:      "REPORT",
	L2_CALLBACK_LEARN_EVENT: "LEARN_EVENT",
	L2_CALLBACK_AGE_EVENT:   "AGE_EVENT",
	L2_CALLBACK_MOVE_EVENT:  "MOVE_EVENT",
}

var l2CallbackOper_values = map[string]L2CallbackOper{
	"DELETE":      L2_CALLBACK_DELETE,
	"ADD":         L2_CALLBACK_ADD,
	"REPORT":      L2_CALLBACK_REPORT,
	"LEARN_EVENT": L2_CALLBACK_LEARN_EVENT,
	"AGE_EVENT":   L2_CALLBACK_AGE_EVENT,
	"MOVE_EVENT":  L2_CALLBACK_MOVE_EVENT,
}

func (v L2CallbackOper) String() string {
	if s, ok := l2CallbackOper_names[v]; ok {
		return s
	}
	return fmt.Sprintf("L2CallbackOper(%d)", v)
}

func ParseL2CallbackOper(s string) (L2CallbackOper, error) {
	if v, ok := l2CallbackOper_values[s]; ok {
		return v, nil
	}
	return L2_CALLBACK_NONE, fmt.Errorf("Invalid L2CallbackOper. %s", s)
}

//
// L2CacheFlags
//
type L2CacheFlags uint32

func (v L2CacheFlags) C() C.uint32 {
	return C.uint32(v)
}

func NewL2CacheFlags(flags ...L2CacheFlags) L2CacheFlags {
	v := L2_CACHE_NONE
	for _, flag := range flags {
		v |= flag
	}
	return v
}

const (
	L2_CACHE_NONE    L2CacheFlags = 0
	L2_CACHE_CPU     L2CacheFlags = C.OPENNSL_L2_CACHE_CPU
	L2_CACHE_DISCARD L2CacheFlags = C.OPENNSL_L2_CACHE_DISCARD
	L2_CACHE_BPDU    L2CacheFlags = C.OPENNSL_L2_CACHE_BPDU
)

var l2cacheFlags_names = map[L2CacheFlags]string{
	L2_CACHE_CPU:     "CPU",
	L2_CACHE_DISCARD: "DISCARD",
	L2_CACHE_BPDU:    "BPDU",
}

var l2cacheFlags_values = map[string]L2CacheFlags{
	"CPU":     L2_CACHE_CPU,
	"DISCARD": L2_CACHE_DISCARD,
	"BPDU":    L2_CACHE_BPDU,
}

func (v L2CacheFlags) String() string {
	names := make([]string, 0, len(l2cacheFlags_names))
	for value, name := range l2cacheFlags_names {
		if v&value != 0 {
			names = append(names, name)
		}
	}
	return strings.Join(names, "|")
}

func ParseL2CacheFlags(name string) (L2CacheFlags, error) {
	if v, ok := l2cacheFlags_values[name]; ok {
		return v, nil
	}
	return L2_CACHE_NONE, fmt.Errorf("Invalid L2CacheFlags. %s", name)
}

//
// L2StationFlags uint32
//
type L2StationFlags uint32

func (v L2StationFlags) C() C.uint32 {
	return C.uint32(v)
}

func NewL2StationFlags(flags ...L2StationFlags) L2StationFlags {
	var flgs L2StationFlags = 0
	for _, flg := range flags {
		flgs |= flg
	}
	return flgs
}

const (
	L2_STATION_NONE     L2StationFlags = 0
	L2_STATION_WITH_ID  L2StationFlags = C.OPENNSL_L2_STATION_WITH_ID
	L2_STATION_IPV4     L2StationFlags = C.OPENNSL_L2_STATION_IPV4
	L2_STATION_IPV6     L2StationFlags = C.OPENNSL_L2_STATION_IPV6
	L2_STATION_ARP_RARP L2StationFlags = C.OPENNSL_L2_STATION_ARP_RARP
)

var l2stationFlags_names = map[L2StationFlags]string{
	L2_STATION_WITH_ID:  "WITH_ID",
	L2_STATION_IPV4:     "IPV4",
	L2_STATION_IPV6:     "IPV6",
	L2_STATION_ARP_RARP: "ARP_RARP",
}

var l2stationFlags_values = map[string]L2StationFlags{
	"WITH_ID":  L2_STATION_WITH_ID,
	"IPV4":     L2_STATION_IPV4,
	"IPV6":     L2_STATION_IPV6,
	"ARP_RARP": L2_STATION_ARP_RARP,
}

func (v L2StationFlags) String() string {
	names := make([]string, 0, len(l2stationFlags_names))
	for value, name := range l2stationFlags_names {
		if v&value != 0 {
			names = append(names, name)
		}
	}
	return strings.Join(names, "|")
}

func ParseL2StationFlags(name string) (L2StationFlags, error) {
	if v, ok := l2stationFlags_values[name]; ok {
		return v, nil
	}
	return L2_STATION_NONE, fmt.Errorf("Invalid L2StationFlags. %s", name)
}

//
// L2LearnLimitFlags
//
type L2LearnLimitFlags uint32

func (v L2LearnLimitFlags) C() C.uint32 {
	return C.uint32(v)
}

const (
	L2_LEARN_NONE         L2LearnLimitFlags = 0
	L2_LEARN_LIMIT_SYSTEM L2LearnLimitFlags = C.OPENNSL_L2_LEARN_LIMIT_SYSTEM
)

var l2LearnLimitFlags_names = map[L2LearnLimitFlags]string{
	L2_LEARN_LIMIT_SYSTEM: "LIMIT_SYSTEM",
}

var l2LearnLimitFlags_values = map[string]L2LearnLimitFlags{
	"LIMIT_SYSTEM": L2_LEARN_LIMIT_SYSTEM,
}

func (v L2LearnLimitFlags) String() string {
	if s, ok := l2LearnLimitFlags_names[v]; ok {
		return s
	}
	return fmt.Sprintf("L2LearnLimitFlags(%d)", v)
}

func ParseL2LearnLimitFlags(s string) (L2LearnLimitFlags, error) {
	if v, ok := l2LearnLimitFlags_values[s]; ok {
		return v, nil
	}
	return L2_LEARN_NONE, fmt.Errorf("Invalid L2LearnLimitFlags. %s", s)
}

//
// L2DeleteFlags
//
type L2DeleteFlags uint32

func (v L2DeleteFlags) C() C.uint32 {
	return C.uint32(v)
}

func NewL2DeleteFlags(flags ...L2DeleteFlags) L2DeleteFlags {
	var flgs L2DeleteFlags = 0
	for _, flg := range flags {
		flgs |= flg
	}
	return flgs
}

const (
	L2_DELETE_NONE         L2DeleteFlags = 0
	L2_DELETE_STATIC       L2DeleteFlags = C.OPENNSL_L2_DELETE_STATIC
	L2_DELETE_PENDING      L2DeleteFlags = C.OPENNSL_L2_DELETE_PENDING
	L2_DELETE_NO_CALLBACKS L2DeleteFlags = C.OPENNSL_L2_DELETE_NO_CALLBACKS
)

var l2deleteFlags_names = map[L2DeleteFlags]string{
	L2_DELETE_STATIC:       "STATIC",
	L2_DELETE_PENDING:      "PENDING",
	L2_DELETE_NO_CALLBACKS: "NO_CALLBACKS",
}

var l2deleteFlags_values = map[string]L2DeleteFlags{
	"STATIC":       L2_DELETE_STATIC,
	"PENDING":      L2_DELETE_PENDING,
	"NO_CALLBACKS": L2_DELETE_NO_CALLBACKS,
}

func (v L2DeleteFlags) String() string {
	names := make([]string, 0, len(l2deleteFlags_names))
	for value, name := range l2deleteFlags_names {
		if v&value != 0 {
			names = append(names, name)
		}
	}
	return strings.Join(names, "|")
}

func ParseL2DeleteFlags(name string) (L2DeleteFlags, error) {
	if v, ok := l2deleteFlags_values[name]; ok {
		return v, nil
	}
	return L2_DELETE_NONE, fmt.Errorf("Invalid L2DeleteFlags. %s", name)
}

//
// L2ReplaceFlags
//
type L2ReplaceFlags uint32

func (v L2ReplaceFlags) C() C.uint32 {
	return C.uint32(v)
}

func NewL2ReplaceFlags(flags ...L2ReplaceFlags) L2ReplaceFlags {
	var flgs L2ReplaceFlags = 0
	for _, flg := range flags {
		flgs |= flg
	}
	return flgs
}

const (
	L2_REPLACE_NONE         L2ReplaceFlags = 0
	L2_REPLACE_MATCH_STATIC L2ReplaceFlags = C.OPENNSL_L2_REPLACE_MATCH_STATIC
	L2_REPLACE_MATCH_MAC    L2ReplaceFlags = C.OPENNSL_L2_REPLACE_MATCH_MAC
	L2_REPLACE_MATCH_VLAN   L2ReplaceFlags = C.OPENNSL_L2_REPLACE_MATCH_VLAN
	L2_REPLACE_MATCH_DEST   L2ReplaceFlags = C.OPENNSL_L2_REPLACE_MATCH_DEST
	L2_REPLACE_DELETE       L2ReplaceFlags = C.OPENNSL_L2_REPLACE_DELETE
	L2_REPLACE_NO_CALLBACKS L2ReplaceFlags = C.OPENNSL_L2_REPLACE_NO_CALLBACKS
)

var l2replaceFlags_names = map[L2ReplaceFlags]string{
	L2_REPLACE_MATCH_STATIC: "MATCH_STATIC",
	L2_REPLACE_MATCH_MAC:    "MATCH_MAC",
	L2_REPLACE_MATCH_VLAN:   "MATCH_VLAN",
	L2_REPLACE_MATCH_DEST:   "_MATCH_DEST",
	L2_REPLACE_DELETE:       "DELETE",
	L2_REPLACE_NO_CALLBACKS: "NO_CALLBACKS",
}

var l2replaceFlags_values = map[string]L2ReplaceFlags{
	"MATCH_STATIC": L2_REPLACE_MATCH_STATIC,
	"MATCH_MAC":    L2_REPLACE_MATCH_MAC,
	"MATCH_VLAN":   L2_REPLACE_MATCH_VLAN,
	"MATCH_DEST":   L2_REPLACE_MATCH_DEST,
	"DELETE":       L2_REPLACE_DELETE,
	"NO_CALLBACKS": L2_REPLACE_NO_CALLBACKS,
}

func (v L2ReplaceFlags) String() string {
	names := make([]string, 0, len(l2replaceFlags_names))
	for value, name := range l2replaceFlags_names {
		if v&value != 0 {
			names = append(names, name)
		}
	}
	return strings.Join(names, "|")
}

func ParseL2ReplaceFlags(name string) (L2ReplaceFlags, error) {
	if v, ok := l2replaceFlags_values[name]; ok {
		return v, nil
	}
	return L2_REPLACE_NONE, fmt.Errorf("Invalid L2ReplaceFlags. %s", name)
}

//
// API
//
func L2TunnelAdd(unit int, hwaddr net.HardwareAddr, vid Vlan) error {
	mac := NewMAC(hwaddr)
	rc := C.opennsl_l2_tunnel_add(C.int(unit), &mac[0], vid.C())
	return ParseError(rc)
}
