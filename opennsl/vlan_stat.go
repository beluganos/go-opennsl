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
#include <opennsl/vlan.h>
*/
import "C"

import (
	"fmt"
)

//
// VlanStat
//
type VlanStat C.opennsl_vlan_stat_t

func (v VlanStat) C() C.opennsl_vlan_stat_t {
	return C.opennsl_vlan_stat_t(v)
}

const (
	VlanStatPackets               VlanStat = C.opennslVlanStatPackets
	VlanStatIngressPackets        VlanStat = C.opennslVlanStatIngressPackets
	VlanStatBytes                 VlanStat = C.opennslVlanStatBytes
	VlanStatIngressBytes          VlanStat = C.opennslVlanStatIngressBytes
	VlanStatEgressPackets         VlanStat = C.opennslVlanStatEgressPackets
	VlanStatEgressBytes           VlanStat = C.opennslVlanStatEgressBytes
	VlanStatForwardedPackets      VlanStat = C.opennslVlanStatForwardedPackets
	VlanStatForwardedBytes        VlanStat = C.opennslVlanStatForwardedBytes
	VlanStatDropPackets           VlanStat = C.opennslVlanStatDropPackets
	VlanStatDropBytes             VlanStat = C.opennslVlanStatDropBytes
	VlanStatUnicastPackets        VlanStat = C.opennslVlanStatUnicastPackets
	VlanStatUnicastBytes          VlanStat = C.opennslVlanStatUnicastBytes
	VlanStatUnicastDropPackets    VlanStat = C.opennslVlanStatUnicastDropPackets
	VlanStatUnicastDropBytes      VlanStat = C.opennslVlanStatUnicastDropBytes
	VlanStatNonUnicastPackets     VlanStat = C.opennslVlanStatNonUnicastPackets
	VlanStatNonUnicastBytes       VlanStat = C.opennslVlanStatNonUnicastBytes
	VlanStatNonUnicastDropPackets VlanStat = C.opennslVlanStatNonUnicastDropPackets
	VlanStatNonUnicastDropBytes   VlanStat = C.opennslVlanStatNonUnicastDropBytes
	VlanStatL3Packets             VlanStat = C.opennslVlanStatL3Packets
	VlanStatL3Bytes               VlanStat = C.opennslVlanStatL3Bytes
	VlanStatL3DropPackets         VlanStat = C.opennslVlanStatL3DropPackets
	VlanStatL3DropBytes           VlanStat = C.opennslVlanStatL3DropBytes
	VlanStatFloodPackets          VlanStat = C.opennslVlanStatFloodPackets
	VlanStatFloodBytes            VlanStat = C.opennslVlanStatFloodBytes
	VlanStatFloodDropPackets      VlanStat = C.opennslVlanStatFloodDropPackets
	VlanStatFloodDropBytes        VlanStat = C.opennslVlanStatFloodDropBytes
	VlanStatGreenPackets          VlanStat = C.opennslVlanStatGreenPackets
	VlanStatGreenBytes            VlanStat = C.opennslVlanStatGreenBytes
	VlanStatYellowPackets         VlanStat = C.opennslVlanStatYellowPackets
	VlanStatYellowBytes           VlanStat = C.opennslVlanStatYellowBytes
	VlanStatRedPackets            VlanStat = C.opennslVlanStatRedPackets
	VlanStatRedBytes              VlanStat = C.opennslVlanStatRedBytes
	VlanStatCount                 VlanStat = C.opennslVlanStatCount
)

var vlanStat_names = map[VlanStat]string{
	// VlanStatPackets:               "Packets",
	// VlanStatBytes:                 "Bytes",
	VlanStatIngressPackets:        "IngressPackets",
	VlanStatIngressBytes:          "IngressBytes",
	VlanStatEgressPackets:         "EgressPackets",
	VlanStatEgressBytes:           "EgressBytes",
	VlanStatForwardedPackets:      "ForwardedPackets",
	VlanStatForwardedBytes:        "ForwardedBytes",
	VlanStatDropPackets:           "DropPackets",
	VlanStatDropBytes:             "DropBytes",
	VlanStatUnicastPackets:        "UnicastPackets",
	VlanStatUnicastBytes:          "UnicastBytes",
	VlanStatUnicastDropPackets:    "UnicastDropPackets",
	VlanStatUnicastDropBytes:      "UnicastDropBytes",
	VlanStatNonUnicastPackets:     "NonUnicastPackets",
	VlanStatNonUnicastBytes:       "NonUnicastBytes",
	VlanStatNonUnicastDropPackets: "NonUnicastDropPackets",
	VlanStatNonUnicastDropBytes:   "NonUnicastDropBytes",
	VlanStatL3Packets:             "L3Packets",
	VlanStatL3Bytes:               "L3Bytes",
	VlanStatL3DropPackets:         "L3DropPackets",
	VlanStatL3DropBytes:           "L3DropBytes",
	VlanStatFloodPackets:          "FloodPackets",
	VlanStatFloodBytes:            "FloodBytes",
	VlanStatFloodDropPackets:      "FloodDropPackets",
	VlanStatFloodDropBytes:        "FloodDropBytes",
	VlanStatGreenPackets:          "GreenPackets",
	VlanStatGreenBytes:            "GreenBytes",
	VlanStatYellowPackets:         "YellowPackets",
	VlanStatYellowBytes:           "YellowBytes",
	VlanStatRedPackets:            "RedPackets",
	VlanStatRedBytes:              "RedBytes",
	VlanStatCount:                 "Count",
}

var vlanStat_values = map[string]VlanStat{
	"Packets":               VlanStatPackets,
	"IngressPackets":        VlanStatIngressPackets,
	"Bytes":                 VlanStatBytes,
	"IngressBytes":          VlanStatIngressBytes,
	"EgressPackets":         VlanStatEgressPackets,
	"EgressBytes":           VlanStatEgressBytes,
	"ForwardedPackets":      VlanStatForwardedPackets,
	"ForwardedBytes":        VlanStatForwardedBytes,
	"DropPackets":           VlanStatDropPackets,
	"DropBytes":             VlanStatDropBytes,
	"UnicastPackets":        VlanStatUnicastPackets,
	"UnicastBytes":          VlanStatUnicastBytes,
	"UnicastDropPackets":    VlanStatUnicastDropPackets,
	"UnicastDropBytes":      VlanStatUnicastDropBytes,
	"NonUnicastPackets":     VlanStatNonUnicastPackets,
	"NonUnicastBytes":       VlanStatNonUnicastBytes,
	"NonUnicastDropPackets": VlanStatNonUnicastDropPackets,
	"NonUnicastDropBytes":   VlanStatNonUnicastDropBytes,
	"L3Packets":             VlanStatL3Packets,
	"L3Bytes":               VlanStatL3Bytes,
	"L3DropPackets":         VlanStatL3DropPackets,
	"L3DropBytes":           VlanStatL3DropBytes,
	"FloodPackets":          VlanStatFloodPackets,
	"FloodBytes":            VlanStatFloodBytes,
	"FloodDropPackets":      VlanStatFloodDropPackets,
	"FloodDropBytes":        VlanStatFloodDropBytes,
	"GreenPackets":          VlanStatGreenPackets,
	"GreenBytes":            VlanStatGreenBytes,
	"YellowPackets":         VlanStatYellowPackets,
	"YellowBytes":           VlanStatYellowBytes,
	"RedPackets":            VlanStatRedPackets,
	"RedBytes":              VlanStatRedBytes,
	"Count":                 VlanStatCount,
}

func (v VlanStat) String() string {
	if s, ok := vlanStat_names[v]; ok {
		return s
	}
	return fmt.Sprintf("VlanStat(%d)", v)
}

func ParseVlanStat(s string) (VlanStat, error) {
	if v, ok := vlanStat_values[s]; ok {
		return v, nil
	}
	return VlanStatPackets, fmt.Errorf("Invalid VlanStat. %s", s)
}

func VlanStatGet(unit int, vid Vlan, cos Cos, stat VlanStat) (uint64, error) {
	val := C.uint64(0)
	rc := C.opennsl_vlan_stat_get(C.int(unit), vid.C(), cos.C(), stat.C(), &val)
	return uint64(val), ParseError(rc)
}

func (v VlanStat) Get(unit int, vid Vlan, cos Cos) (uint64, error) {
	return VlanStatGet(unit, vid, cos, v)
}

func VlanStatSet(unit int, vid Vlan, cos Cos, stat VlanStat, val uint64) error {
	rc := C.opennsl_vlan_stat_set(C.int(unit), vid.C(), cos.C(), stat.C(), C.uint64(val))
	return ParseError(rc)

}

func (v VlanStat) Set(unit int, vid Vlan, cos Cos, val uint64) error {
	return VlanStatSet(unit, vid, cos, v, val)
}
