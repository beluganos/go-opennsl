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
// FieldEntry
//
type FieldEntry C.opennsl_field_entry_t

func (v FieldEntry) C() C.opennsl_field_entry_t {
	return C.opennsl_field_entry_t(v)
}

func FieldGroupEntryCreate(unit int, v FieldGroup) (FieldEntry, error) {
	c_entry := C.opennsl_field_entry_t(0)
	rc := C.opennsl_field_entry_create(C.int(unit), v.C(), &c_entry)
	return FieldEntry(c_entry), ParseError(rc)
}

func (v FieldGroup) EntryCreate(unit int) (FieldEntry, error) {
	return FieldGroupEntryCreate(unit, v)
}

func FieldEntryCreateID(unit int, group FieldGroup, v FieldEntry) error {
	rc := C.opennsl_field_entry_create_id(C.int(unit), group.C(), v.C())
	return ParseError(rc)
}

func (v FieldEntry) CreateID(unit int, group FieldGroup) error {
	return FieldEntryCreateID(unit, group, v)
}

func FieldGroupEntrySize(unit int, v FieldGroup) (int, error) {
	c_count := C.int(0)
	rc := C.opennsl_field_entry_multi_get(C.int(unit), v.C(), 0, nil, &c_count)
	return int(c_count), ParseError(rc)
}

func (v FieldGroup) EntrySize(unit int) (int, error) {
	return FieldGroupEntrySize(unit, v)
}

func FieldGroupEntryMultiGet(unit int, v FieldGroup, maxSize int) ([]FieldEntry, error) {
	if maxSize < 0 {
		size, err := v.EntrySize(unit)
		if err != nil {
			return nil, err
		}
		maxSize = size
	}

	c_count := C.int(0)
	c_entry := make([]C.opennsl_field_entry_t, maxSize)
	if maxSize > 0 {
		rc := C.opennsl_field_entry_multi_get(C.int(unit), v.C(), C.int(maxSize), &c_entry[0], &c_count)

		if err := ParseError(rc); err != nil {
			return nil, err
		}
	}

	entry := make([]FieldEntry, int(c_count))
	for index := 0; index < int(c_count); index++ {
		entry[index] = FieldEntry(c_entry[index])
	}

	return entry, nil
}

func (v FieldGroup) EntryMultiGet(unit int, maxSize int) ([]FieldEntry, error) {
	return FieldGroupEntryMultiGet(unit, v, maxSize)
}

func FieldEntryDestyroy(unit int, v FieldEntry) error {
	rc := C.opennsl_field_entry_destroy(C.int(unit), v.C())
	return ParseError(rc)
}

func (v FieldEntry) Destyroy(unit int) error {
	return FieldEntryDestyroy(unit, v)
}

func FieldEntryDestroyAll(unit int) error {
	rc := C.opennsl_field_entry_destroy_all(C.int(unit))
	return ParseError(rc)
}

func FieldEntryCopy(unit int, v FieldEntry) (FieldEntry, error) {
	c_entry := C.opennsl_field_entry_t(0)
	rc := C.opennsl_field_entry_copy(C.int(unit), v.C(), &c_entry)
	return FieldEntry(c_entry), ParseError(rc)
}

func (v FieldEntry) Copy(unit int) (FieldEntry, error) {
	return FieldEntryCopy(unit, v)
}

func FieldEntryInstall(unit int, v FieldEntry) error {
	rc := C.opennsl_field_entry_install(C.int(unit), v.C())
	return ParseError(rc)
}

func (v FieldEntry) Install(unit int) error {
	return FieldEntryInstall(unit, v)
}

func FieldEntryReinstall(unit int, v FieldEntry) error {
	rc := C.opennsl_field_entry_reinstall(C.int(unit), v.C())
	return ParseError(rc)
}

func (v FieldEntry) Reinstall(unit int) error {
	return FieldEntryReinstall(unit, v)
}

func FieldEntryRemove(unit int, v FieldEntry) error {
	rc := C.opennsl_field_entry_remove(C.int(unit), v.C())
	return ParseError(rc)
}

func (v FieldEntry) Remove(unit int) error {
	return FieldEntryRemove(unit, v)
}

func FieldEntryPolicerAttach(unit int, v FieldEntry, level int, policer Policer) error {
	rc := C.opennsl_field_entry_policer_attach(C.int(unit), v.C(), C.int(level), policer.C())
	return ParseError(rc)
}

func (v FieldEntry) PolicerAttach(unit int, level int, policer Policer) error {
	return FieldEntryPolicerAttach(unit, v, level, policer)
}

func FieldEntryPolicerDetach(unit int, v FieldEntry, level int) error {
	rc := C.opennsl_field_entry_policer_detach(C.int(unit), v.C(), C.int(level))
	return ParseError(rc)
}

func (v FieldEntry) PolicerDetach(unit int, level int) error {
	return FieldEntryPolicerDetach(unit, v, level)
}

func FieldEntryPolicerGet(unit int, v FieldEntry, level int) (Policer, error) {
	c_policer := C.opennsl_policer_t(0)
	rc := C.opennsl_field_entry_policer_get(C.int(unit), v.C(), C.int(level), &c_policer)
	return Policer(c_policer), ParseError(rc)
}

func (v FieldEntry) PolicerGet(unit int, level int) (Policer, error) {
	return FieldEntryPolicerGet(unit, v, level)
}

func FieldEntryPrioGet(unit int, v FieldEntry) (int, error) {
	c_prio := C.int(0)
	rc := C.opennsl_field_entry_prio_get(C.int(unit), v.C(), &c_prio)
	return int(c_prio), ParseError(rc)
}

func (v FieldEntry) PrioGet(unit int) (int, error) {
	return FieldEntryPrioGet(unit, v)
}

func FieldEntryPrioSet(unit int, v FieldEntry, prio int) error {
	rc := C.opennsl_field_entry_prio_set(C.int(unit), v.C(), C.int(prio))
	return ParseError(rc)
}

func (v FieldEntry) PrioSet(unit int, prio int) error {
	return FieldEntryPrioSet(unit, v, prio)
}

func (v FieldEntry) Qualify() FieldEntryQualify {
	return FieldEntryQualify(v.C())
}

func (v FieldEntry) Action() FieldEntryAction {
	return FieldEntryAction(v.C())
}

func (v FieldEntry) Stat() FieldEntryStat {
	return FieldEntryStat(v.C())
}
