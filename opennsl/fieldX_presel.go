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

//
// FieldPresel
//
type FieldPresel C.opennsl_field_presel_t

func (v FieldPresel) C() C.opennsl_field_presel_t {
	return C.opennsl_field_presel_t(v)
}

func FieldPreselCreate(unit int) (FieldPresel, error) {
	c_presel := C.opennsl_field_presel_t(0)

	rc := C.opennsl_field_presel_create(C.int(unit), &c_presel)
	return FieldPresel(c_presel), ParseError(rc)
}

func FieldPreselCreateID(unit int, presel FieldPresel) error {
	rc := C.opennsl_field_presel_create_id(C.int(unit), presel.C())
	return ParseError(rc)
}

func (v FieldPresel) Destroy(unit int) error {
	rc := C.opennsl_field_presel_destroy(C.int(unit), v.C())
	return ParseError(rc)
}

//
// FieldPreselSet
//
type FieldPreselSet C.opennsl_field_presel_set_t

func (v *FieldPreselSet) C() *C.opennsl_field_presel_set_t {
	return (*C.opennsl_field_presel_set_t)(v)
}

func (v *FieldPreselSet) Init() {
	// TODO:
}

func (v *FieldPreselSet) Add(preselId FieldPresel) {
	// TODO
}

func (v *FieldPreselSet) Remove(preselId FieldPresel) {
	// TODO
}

func (v *FieldPreselSet) Test(preselId FieldPresel) uint32 {
	// TODO
	return 0
}
