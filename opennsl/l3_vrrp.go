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
#include "helper.h"
*/
import "C"

//
// L3VrrpID
//
type L3VrrpID C.uint32

func (v L3VrrpID) C() C.uint32 {
	return C.uint32(v)
}

func L3VrrpAdd(unit int, vid Vlan, v L3VrrpID) error {
	rc := C.opennsl_l3_vrrp_add(C.int(unit), vid.C(), v.C())
	return ParseError(rc)
}

func (v L3VrrpID) Add(unit int, vid Vlan) error {
	return L3VrrpAdd(unit, vid, v)
}

func L3VrrpDelete(unit int, vid Vlan, v L3VrrpID) error {
	rc := C.opennsl_l3_vrrp_delete(C.int(unit), vid.C(), v.C())
	return ParseError(rc)
}

func (v L3VrrpID) Delete(unit int, vid Vlan) error {
	return L3VrrpDelete(unit, vid, v)
}

func L3VrrpDeleteAll(unit int, vid Vlan) error {
	rc := C.opennsl_l3_vrrp_delete_all(C.int(unit), vid.C())
	return ParseError(rc)
}

func L3VrrpGet(unit int, vid Vlan, maxSize int) ([]L3VrrpID, error) {
	c_count := C.int(0)
	c_arr := make([]C.int, maxSize)

	rc := C.opennsl_l3_vrrp_get(C.int(unit), vid.C(), C.int(maxSize), &c_arr[0], &c_count)
	if err := ParseError(rc); err != nil {
		return nil, err
	}

	vrids := make([]L3VrrpID, int(c_count))
	for index := 0; index < int(c_count); index++ {
		vrids[index] = L3VrrpID(int(c_arr[index]))
	}

	return vrids, nil
}
