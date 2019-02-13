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

//
// FieldStatID
//
type FieldStatID C.int

func (v FieldStatID) C() C.int {
	return C.int(v)
}

func FielsStatDestroy(unit int, statID FieldStatID) error {
	rc := C.opennsl_field_stat_destroy(C.int(unit), statID.C())
	return ParseError(rc)
}

func FieldStatSize(unit int, statID FieldStatID) (int, error) {
	c_size := C.int(0)

	rc := C.opennsl_field_stat_size(C.int(unit), statID.C(), &c_size)
	return int(c_size), ParseError(rc)
}

func FieldStatSet(unit int, v FieldStatID, stat FieldStat, value uint64) error {
	rc := C.opennsl_field_stat_set(C.int(unit), v.C(), stat.C(), C.uint64(value))
	return ParseError(rc)
}

func (v FieldStatID) Set(unit int, stat FieldStat, value uint64) error {
	return FieldStatSet(unit, v, stat, value)
}

func FieldStatAllSet(unit int, v FieldStatID, value uint64) error {
	rc := C.opennsl_field_stat_all_set(C.int(unit), v.C(), C.uint64(value))
	return ParseError(rc)
}

func (v FieldStatID) AllSet(unit int, value uint64) error {
	return FieldStatAllSet(unit, v, value)
}

func FieldStatGet(unit int, v FieldStatID, stat FieldStat) (uint64, error) {
	c_value := C.uint64(0)

	rc := C.opennsl_field_stat_get(C.int(unit), v.C(), stat.C(), &c_value)
	return uint64(c_value), ParseError(rc)
}

func (v FieldStatID) Get(unit int, stat FieldStat) (uint64, error) {
	return FieldStatGet(unit, v, stat)
}

func FieldStatDetach(unit int, v FieldStatID) error {
	rc := C.opennsl_field_stat_detach(C.int(unit), C.uint32(v.C()))
	return ParseError(rc)
}

func (v FieldStatID) Detach(unit int) error {
	return FieldStatDetach(unit, v)
}
