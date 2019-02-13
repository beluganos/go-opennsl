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
#include <opennsl/cosq.h>
*/
import "C"

import (
	"fmt"
)

//
// BstStatId
//
type BstStatId C.opennsl_bst_stat_id_t

func (v BstStatId) C() C.opennsl_bst_stat_id_t {
	return C.opennsl_bst_stat_id_t(v)
}

const (
	BstStatIdInvalid                BstStatId = C.opennslBstStatIdInvalid
	BstStatIdDevice                 BstStatId = C.opennslBstStatIdDevice
	BstStatIdEgrPool                BstStatId = C.opennslBstStatIdEgrPool
	BstStatIdEgrMCastPool           BstStatId = C.opennslBstStatIdEgrMCastPool
	BstStatIdIngPool                BstStatId = C.opennslBstStatIdIngPool
	BstStatIdPortPool               BstStatId = C.opennslBstStatIdPortPool
	BstStatIdPriGroupShared         BstStatId = C.opennslBstStatIdPriGroupShared
	BstStatIdPriGroupHeadroom       BstStatId = C.opennslBstStatIdPriGroupHeadroom
	BstStatIdUcast                  BstStatId = C.opennslBstStatIdUcast
	BstStatIdMcast                  BstStatId = C.opennslBstStatIdMcast
	BstStatIdHeadroomPool           BstStatId = C.opennslBstStatIdHeadroomPool
	BstStatIdEgrPortPoolSharedUcast BstStatId = C.opennslBstStatIdEgrPortPoolSharedUcast
	BstStatIdEgrPortPoolSharedMcast BstStatId = C.opennslBstStatIdEgrPortPoolSharedMcast
)

const BstStatIdMaxCount = C.opennslBstStatIdMaxCount

var bstStatId_names = map[BstStatId]string{
	BstStatIdInvalid:                "Invalid",
	BstStatIdDevice:                 "Device",
	BstStatIdEgrPool:                "EgrPool",
	BstStatIdEgrMCastPool:           "MCastPool",
	BstStatIdIngPool:                "IngPool",
	BstStatIdPortPool:               "PortPool",
	BstStatIdPriGroupShared:         "PriGroupShared",
	BstStatIdPriGroupHeadroom:       "PriGroupHeadroom",
	BstStatIdUcast:                  "Ucast",
	BstStatIdMcast:                  "Mcast",
	BstStatIdHeadroomPool:           "HeadroomPool",
	BstStatIdEgrPortPoolSharedUcast: "EgrPortPoolSharedUcast",
	BstStatIdEgrPortPoolSharedMcast: "EgrPortPoolSharedMcast",
}

var bstStatI_values = map[string]BstStatId{
	"Invalid":                BstStatIdInvalid,
	"Device":                 BstStatIdDevice,
	"EgrPool":                BstStatIdEgrPool,
	"EgrMCastPool":           BstStatIdEgrMCastPool,
	"IngPool":                BstStatIdIngPool,
	"PortPool":               BstStatIdPortPool,
	"PriGroupShared":         BstStatIdPriGroupShared,
	"PriGroupHeadroom":       BstStatIdPriGroupHeadroom,
	"Ucast":                  BstStatIdUcast,
	"Mcast":                  BstStatIdMcast,
	"HeadroomPool":           BstStatIdHeadroomPool,
	"EgrPortPoolSharedUcast": BstStatIdEgrPortPoolSharedUcast,
	"EgrPortPoolSharedMcast": BstStatIdEgrPortPoolSharedMcast,
}

func (v BstStatId) String() string {
	if s, ok := bstStatId_names[v]; ok {
		return s
	}
	return fmt.Sprintf("BstStatId(%d)", v)
}

func ParseBstStatId(s string) (BstStatId, error) {
	if v, ok := bstStatI_values[s]; ok {
		return v, nil
	}
	return BstStatIdInvalid, fmt.Errorf("Invalid BstStatId. %s", s)
}

//
// CosqBstProfile
//
type CosqBstProfile C.opennsl_cosq_bst_profile_t

func (v *CosqBstProfile) C() *C.opennsl_cosq_bst_profile_t {
	return (*C.opennsl_cosq_bst_profile_t)(v)
}

func (v *CosqBstProfile) Byte() uint32 {
	return uint32(v.byte)
}

func (v *CosqBstProfile) SetByte(b uint32) {
	v.byte = C.uint32(b)
}

func NewCosqBstProfile() *CosqBstProfile {
	p := &CosqBstProfile{}
	p.Init()
	return p
}
func (v *CosqBstProfile) Init() {
	v.byte = 0
}

//
// API
//
func (v *CosqBstProfile) Set(unit int, gport GPort, cosq CosQueue, bid BstStatId) error {
	rc := C.opennsl_cosq_bst_profile_set(C.int(unit), gport.C(), cosq.C(), bid.C(), v.C())
	return ParseError(rc)
}

func CosqBstProfileGet(unit int, gport GPort, cosq CosQueue, bid BstStatId) (*CosqBstProfile, error) {
	profile := NewCosqBstProfile()

	rc := C.opennsl_cosq_bst_profile_get(C.int(unit), gport.C(), cosq.C(), bid.C(), profile.C())
	if err := ParseError(rc); err != nil {
		return nil, err
	}

	return profile, nil
}

func (v BstStatId) Sync(unit int) error {
	rc := C.opennsl_cosq_bst_stat_sync(C.int(unit), v.C())
	return ParseError(rc)
}

func (v BstStatId) Clear(unit int, gport GPort, cosq CosQueue) error {
	rc := C.opennsl_cosq_bst_stat_clear(C.int(unit), gport.C(), cosq.C(), v.C())
	return ParseError(rc)
}

func (v BstStatId) Get(unit int, gport GPort, cosq CosQueue, opt uint32) (uint64, error) {
	c_value := C.uint64(0)
	rc := C.opennsl_cosq_bst_stat_get(C.int(unit), gport.C(), cosq.C(), v.C(), C.uint32(opt), &c_value)
	return uint64(c_value), ParseError(rc)
}

func BstStatMultiGet(unit int, gport GPort, cosq CosQueue, opt uint32, bids ...BstStatId) ([]uint64, error) {
	num := len(bids)
	c_vals := make([]C.uint64, num)

	if num > 0 {
		c_bids := make([]C.opennsl_bst_stat_id_t, num)
		for index, bid := range bids {
			c_bids[index] = bid.C()
		}

		rc := C.opennsl_cosq_bst_stat_multi_get(C.int(unit), gport.C(), cosq.C(), C.uint32(opt), C.int(num), &c_bids[0], &c_vals[0])
		if err := ParseError(rc); err != nil {
			return nil, err
		}
	}

	values := make([]uint64, num)
	for index, c_val := range c_vals {
		values[index] = uint64(c_val)
	}

	return values, nil
}
