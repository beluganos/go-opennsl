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
#include <opennsl/l3.h>
*/
import "C"

import (
	"fmt"
)

//
// L3Stat
//
type L3Stat C.opennsl_l3_stat_t

func (v L3Stat) C() C.opennsl_l3_stat_t {
	return C.opennsl_l3_stat_t(v)
}

const (
	L3StatNone        L3Stat = 0
	L3StatOutPackets  L3Stat = C.opennslL3StatOutPackets
	L3StatOutBytes    L3Stat = C.opennslL3StatOutBytes
	L3StatDropPackets L3Stat = C.opennslL3StatDropPackets
	L3StatDropBytes   L3Stat = C.opennslL3StatDropBytes
	L3StatInPackets   L3Stat = C.opennslL3StatInPackets
	L3StatInBytes     L3Stat = C.opennslL3StatInBytes
)

var l3Stat_names = map[L3Stat]string{
	L3StatOutPackets:  "OutPackets",
	L3StatOutBytes:    "OutBytes",
	L3StatDropPackets: "DropPackets",
	L3StatDropBytes:   "DropBytes",
	L3StatInPackets:   "InPackets",
	L3StatInBytes:     "InBytes",
}

var l3Stat_values = map[string]L3Stat{
	"OutPackets":  L3StatOutPackets,
	"OutBytes":    L3StatOutBytes,
	"DropPackets": L3StatDropPackets,
	"DropBytes":   L3StatDropBytes,
	"InPackets":   L3StatInPackets,
	"InBytes":     L3StatInBytes,
}

func (v L3Stat) String() string {
	if s, ok := l3Stat_names[v]; ok {
		return s
	}
	return fmt.Sprintf("L3Stat(%d)", v)
}

func ParseL3Stat(s string) (L3Stat, error) {
	if v, ok := l3Stat_values[s]; ok {
		return v, nil
	}
	return L3StatNone, fmt.Errorf("Invalid L3Stat. %s", s)
}

//
// API
//
func L3StatEgressCounterGet(unit int, l3eg L3EgressID, v L3Stat, counterIndexes []uint32) ([]StatValue, error) {
	size := len(counterIndexes)
	c_arr := make([]C.uint32, size)
	for index := 0; index < size; index++ {
		c_arr[index] = C.uint32(counterIndexes[index])
	}

	c_stats := make([]C.opennsl_stat_value_t, size)
	rc := C.opennsl_l3_egress_stat_counter_get(C.int(unit), l3eg.C(), v.C(), C.uint32(size), &c_arr[0], &c_stats[0])

	if err := ParseError(rc); err != nil {
		return nil, err
	}

	stats := make([]StatValue, size)
	for index, c_stat := range c_stats {
		stats[index] = StatValue(c_stat)
	}

	return stats, nil
}

func (v L3Stat) EgressCounterGet(unit int, l3eg L3EgressID, counterIndexes []uint32) ([]StatValue, error) {
	return L3StatEgressCounterGet(unit, l3eg, v, counterIndexes)
}

func (v L3EgressID) StatCounterGet(unit int, stat L3Stat, counterIndexes []uint32) ([]StatValue, error) {
	return L3StatEgressCounterGet(unit, v, stat, counterIndexes)
}

func L3EgressStatAttach(unit int, l3eg L3EgressID, counterId uint32) error {
	rc := C.opennsl_l3_egress_stat_attach(C.int(unit), l3eg.C(), C.uint32(counterId))
	return ParseError(rc)
}

func (v L3EgressID) StatAttach(unit int, counterId uint32) error {
	return L3EgressStatAttach(unit, v, counterId)
}

func L3EgressStatDetach(unit int, l3eg L3EgressID) error {
	rc := C.opennsl_l3_egress_stat_detach(C.int(unit), l3eg.C())
	return ParseError(rc)
}

func (v L3EgressID) StatDetach(unit int) error {
	return L3EgressStatDetach(unit, v)
}

func L3StatIngressCounterGet(unit int, l3ing L3IngressID, v L3Stat, counterIndexes []uint32) ([]StatValue, error) {
	size := len(counterIndexes)
	c_arr := make([]C.uint32, size)
	for index := 0; index < size; index++ {
		c_arr[index] = C.uint32(counterIndexes[index])
	}

	c_stats := make([]C.opennsl_stat_value_t, size)
	rc := C.opennsl_l3_ingress_stat_counter_get(C.int(unit), l3ing.C(), v.C(), C.uint32(size), &c_arr[0], &c_stats[0])

	if err := ParseError(rc); err != nil {
		return nil, err
	}

	stats := make([]StatValue, size)
	for index, c_stat := range c_stats {
		stats[index] = StatValue(c_stat)
	}

	return stats, nil
}

func (v L3Stat) IngressCounterGet(unit int, l3ing L3IngressID, counterIndexes []uint32) ([]StatValue, error) {
	return L3StatIngressCounterGet(unit, l3ing, v, counterIndexes)
}

func (v L3IngressID) StatCounterGet(unit int, stat L3Stat, counterIndexes []uint32) ([]StatValue, error) {
	return L3StatIngressCounterGet(unit, v, stat, counterIndexes)
}

func L3IngressStatAttach(unit int, l3ing L3IngressID, counterId uint32) error {
	rc := C.opennsl_l3_ingress_stat_attach(C.int(unit), l3ing.C(), C.uint32(counterId))
	return ParseError(rc)
}

func (v L3IngressID) StatAttach(unit int, counterId uint32) error {
	return L3IngressStatAttach(unit, v, counterId)
}

func L3IngressStatDetach(unit int, l3ing L3IngressID) error {
	rc := C.opennsl_l3_ingress_stat_detach(C.int(unit), l3ing.C())
	return ParseError(rc)
}

func (v L3IngressID) StatDetach(unit int) error {
	return L3IngressStatDetach(unit, v)
}
