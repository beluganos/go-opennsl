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
// FieldRange
//
type FieldRange C.opennsl_field_range_t

func (v FieldRange) C() C.opennsl_field_range_t {
	return C.opennsl_field_range_t(v)
}

func FieldRangeCreate(unit int, flags FieldRangeFlags, min, max L4Port) (FieldRange, error) {
	c_range := C.opennsl_field_range_t(0)
	rc := C.opennsl_field_range_create(C.int(unit), &c_range, flags.C(), min.C(), max.C())
	return FieldRange(c_range), ParseError(rc)
}

func FieldRangeGet(unit int, v FieldRange) (FieldRangeFlags, L4Port, L4Port, error) {
	c_flags := C.uint32(0)
	c_min := C.opennsl_l4_port_t(0)
	c_max := C.opennsl_l4_port_t(0)

	rc := C.opennsl_field_range_get(C.int(unit), v.C(), &c_flags, &c_min, &c_max)
	return FieldRangeFlags(c_flags), L4Port(c_min), L4Port(c_max), ParseError(rc)
}

func (v FieldRange) Get(unit int) (FieldRangeFlags, L4Port, L4Port, error) {
	return FieldRangeGet(unit, v)
}

func FieldRangeDestroy(unit int, v FieldRange) error {
	rc := C.opennsl_field_range_destroy(C.int(unit), v.C())
	return ParseError(rc)
}

func (v FieldRange) Destroy(unit int) error {
	return FieldRangeDestroy(unit, v)
}
