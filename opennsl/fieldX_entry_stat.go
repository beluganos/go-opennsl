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
// FieldEntryStat
//
type FieldEntryStat C.opennsl_field_entry_t

func (v FieldEntryStat) C() C.opennsl_field_entry_t {
	return C.opennsl_field_entry_t(v)
}

func (v FieldEntryStat) Attach(unit int, statID FieldStatID) error {
	rc := C.opennsl_field_entry_stat_attach(C.int(unit), v.C(), statID.C())
	return ParseError(rc)
}

func (v FieldEntryStat) Detach(unit int, statID FieldStatID) error {
	rc := C.opennsl_field_entry_stat_detach(C.int(unit), v.C(), statID.C())
	return ParseError(rc)
}

func (v FieldEntryStat) Get(unit int) (FieldStatID, error) {
	c_statid := C.int(0)

	rc := C.opennsl_field_entry_stat_get(C.int(unit), v.C(), &c_statid)
	return FieldStatID(c_statid), ParseError(rc)
}
