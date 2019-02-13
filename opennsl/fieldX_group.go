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

import (
	"fmt"
	"unsafe"
)

//
// FieldGroupStatus
//
type FieldGroupStatus C.opennsl_field_group_status_t

func (v *FieldGroupStatus) C() *C.opennsl_field_group_status_t {
	return (*C.opennsl_field_group_status_t)(v)
}

func (v *FieldGroupStatus) Prio() (int, int) {
	return int(v.prio_min), int(v.prio_max)
}

func (v *FieldGroupStatus) Entries() (int, int, int) {
	return int(v.entries_total), int(v.entries_free), int(v.entry_count)
}

func (v *FieldGroupStatus) Counters() (int, int, int) {
	return int(v.counters_total), int(v.counters_free), int(v.counter_count)
}

func (v *FieldGroupStatus) Meters() (int, int, int) {
	return int(v.meters_total), int(v.meters_free), int(v.meter_count)
}

//
// FieldGroupMode
//
type FieldGroupMode C.opennsl_field_group_mode_t

func (v FieldGroupMode) C() C.opennsl_field_group_mode_t {
	return C.opennsl_field_group_mode_t(v)
}

const (
	FieldGroupModeSingle FieldGroupMode = C.opennslFieldGroupModeSingle
	FieldGroupModeDouble FieldGroupMode = C.opennslFieldGroupModeDouble
	FieldGroupModeAuto   FieldGroupMode = C.opennslFieldGroupModeAuto
)

var fieldGroupMode_names = map[FieldGroupMode]string{
	FieldGroupModeSingle: "Single",
	FieldGroupModeDouble: "Double",
	FieldGroupModeAuto:   "Auto",
}

var fieldGroupMode_values = map[string]FieldGroupMode{
	"Single": FieldGroupModeSingle,
	"Double": FieldGroupModeDouble,
	"Auto":   FieldGroupModeAuto,
}

func (v FieldGroupMode) String() string {
	if s, ok := fieldGroupMode_names[v]; ok {
		return s
	}
	return fmt.Sprintf("FieldGroupMode(%d)", v)
}

func ParseFieldGroupMode(s string) (FieldGroupMode, error) {
	if v, ok := fieldGroupMode_values[s]; ok {
		return v, nil
	}
	return FieldGroupModeSingle, fmt.Errorf("Invalid FieldGroupMode. %s", s)
}

//
// FieldGroup
//
type FieldGroup C.opennsl_field_group_t

const FIELD_GROUP_PRIO_ANY = C.OPENNSL_FIELD_GROUP_PRIO_ANY

func (v FieldGroup) C() C.opennsl_field_group_t {
	return C.opennsl_field_group_t(v)
}

func FieldGroupPreselSet(unit int, v FieldGroup, ps *FieldPreselSet) error {
	rc := C.opennsl_field_group_presel_set(C.int(unit), v.C(), ps.C())
	return ParseError(rc)
}

func (v FieldGroup) PreselSet(unit int, ps *FieldPreselSet) error {
	return FieldGroupPreselSet(unit, v, ps)
}

func FieldGroupPreselGet(unit int, v FieldGroup) (*FieldPreselSet, error) {
	ps := FieldPreselSet{}
	ps.Init()

	rc := C.opennsl_field_group_presel_get(C.int(unit), v.C(), ps.C())
	if err := ParseError(rc); err != nil {
		return nil, err
	}

	return &ps, nil
}

func (v FieldGroup) PreselGet(unit int) (*FieldPreselSet, error) {
	return FieldGroupPreselGet(unit, v)
}

func FieldGroupCreate(unit int, qset *FieldQSet, pri int) (FieldGroup, error) {
	c_grp := C.opennsl_field_group_t(0)
	rc := C.opennsl_field_group_create(C.int(unit), *qset.C(), C.int(pri), &c_grp)
	return FieldGroup(c_grp), ParseError(rc)
}

func FieldGroupCreateID(unit int, qset *FieldQSet, pri int, v FieldGroup) error {
	rc := C.opennsl_field_group_create_id(C.int(unit), *qset.C(), C.int(pri), v.C())
	return ParseError(rc)
}

func (v FieldGroup) CreateID(unit int, qset *FieldQSet, pri int) error {
	return FieldGroupCreateID(unit, qset, pri, v)
}

func FieldGroupCreateMode(unit int, qset *FieldQSet, pri int, mode FieldGroupMode) (FieldGroup, error) {
	c_grp := C.opennsl_field_group_t(0)
	rc := C.opennsl_field_group_create_mode(C.int(unit), *qset.C(), C.int(pri), mode.C(), &c_grp)
	return FieldGroup(c_grp), ParseError(rc)
}

func FieldGroupCreateModeID(unit int, qset *FieldQSet, pri int, mode FieldGroupMode, v FieldGroup) error {
	rc := C.opennsl_field_group_create_mode_id(C.int(unit), *qset.C(), C.int(pri), mode.C(), v.C())
	return ParseError(rc)
}

func (v FieldGroup) CreateModeID(unit int, qset *FieldQSet, pri int, mode FieldGroupMode) error {
	return FieldGroupCreateModeID(unit, qset, pri, mode, v)
}

type FieldGroupTraverseCallback func(int, *FieldGroup) int

var fieldGroupTraverseCallbacks = NewCallbackMap()

//export go_opennsl_field_group_traverse_cb
func go_opennsl_field_group_traverse_cb(c_unit C.int, c_group C.opennsl_field_group_t, c_data unsafe.Pointer) int {
	n := (*uint64)(c_data)
	if h, ok := fieldGroupTraverseCallbacks.Get(*n); ok {
		callback := h.(FieldGroupTraverseCallback)
		group := FieldGroup(c_group)
		return callback(int(c_unit), &group)
	}

	return int(E_PARAM)
}

func FieldGroupTraverse(unit int, callback FieldGroupTraverseCallback) error {
	n := fieldGroupTraverseCallbacks.Add(callback)
	defer fieldGroupTraverseCallbacks.Del(n)

	rc := C.opennsl_field_group_traverse(C.int(unit), (C.opennsl_field_group_traverse_cb)(C._opennsl_field_group_traverse_cb), unsafe.Pointer(&n))
	return ParseError(rc)
}

func FieldGroupQSetSet(unit int, v FieldGroup, qset *FieldQSet) error {
	rc := C.opennsl_field_group_set(C.int(unit), v.C(), *qset.C())
	return ParseError(rc)
}

func (v FieldGroup) QSetSet(unit int, qset *FieldQSet) error {
	return FieldGroupQSetSet(unit, v, qset)
}

func FieldGroupQSetGet(unit int, v FieldGroup) (*FieldQSet, error) {
	qset := FieldQSet{}
	rc := C.opennsl_field_group_get(C.int(unit), v.C(), qset.C())
	if err := ParseError(rc); err != nil {
		return nil, err
	}

	return &qset, nil
}

func (v FieldGroup) QSetGet(unit int) (*FieldQSet, error) {
	return FieldGroupQSetGet(unit, v)
}

func FieldGroupASetSet(unit int, v FieldGroup, aset *FieldASet) error {
	rc := C.opennsl_field_group_action_set(C.int(unit), v.C(), *aset.C())
	return ParseError(rc)
}

func (v FieldGroup) ASetSet(unit int, aset *FieldASet) error {
	return FieldGroupASetSet(unit, v, aset)
}

func FieldGroupASetGet(unit int, v FieldGroup) (*FieldASet, error) {
	aset := FieldASet{}
	rc := C.opennsl_field_group_action_get(C.int(unit), v.C(), aset.C())
	if err := ParseError(rc); err != nil {
		return nil, err
	}

	return &aset, nil
}

func (v FieldGroup) ASetGet(unit int) (*FieldASet, error) {
	return FieldGroupASetGet(unit, v)
}

func FieldGroupDestroy(unit int, v FieldGroup) error {
	rc := C.opennsl_field_group_destroy(C.int(unit), v.C())
	return ParseError(rc)
}

func (v FieldGroup) Destroy(unit int) error {
	return FieldGroupDestroy(unit, v)
}

func FieldGroupPrioritySet(unit int, v FieldGroup, priority int) error {
	rc := C.opennsl_field_group_priority_set(C.int(unit), v.C(), C.int(priority))
	return ParseError(rc)
}

func (v FieldGroup) PrioritySet(unit int, priority int) error {
	return FieldGroupPrioritySet(unit, v, priority)
}

func FieldGroupPriorityGet(unit int, v FieldGroup, priority int) (int, error) {
	c_pri := C.int(0)
	rc := C.opennsl_field_group_priority_get(C.int(unit), v.C(), &c_pri)
	return int(c_pri), ParseError(rc)
}

func (v FieldGroup) PriorityGet(unit int, priority int) (int, error) {
	return FieldGroupPriorityGet(unit, v, priority)
}

func FieldGroupStatusGet(unit int, v FieldGroup) (*FieldGroupStatus, error) {
	status := FieldGroupStatus{}
	rc := C.opennsl_field_group_status_get(C.int(unit), v.C(), status.C())
	if err := ParseError(rc); err != nil {
		return nil, err
	}

	return &status, nil
}

func (v FieldGroup) StatusGet(unit int) (*FieldGroupStatus, error) {
	return FieldGroupStatusGet(unit, v)
}

func FieldGroupEnableSet(unit int, v FieldGroup, enable int) error {
	rc := C.opennsl_field_group_enable_set(C.int(unit), v.C(), C.int(enable))
	return ParseError(rc)
}

func (v FieldGroup) EnableSet(unit int, enable int) error {
	return FieldGroupEnableSet(unit, v, enable)
}

func FieldGroupEnableGet(unit int, v FieldGroup) (int, error) {
	c_enable := C.int(0)
	rc := C.opennsl_field_group_enable_get(C.int(unit), v.C(), &c_enable)
	return int(c_enable), ParseError(rc)
}

func (v FieldGroup) EnableGet(unit int) (int, error) {
	return FieldGroupEnableGet(unit, v)
}

func FieldGroupInstall(unit int, v FieldGroup) error {
	rc := C.opennsl_field_group_install(C.int(unit), v.C())
	return ParseError(rc)
}

func (v FieldGroup) Install(unit int) error {
	return FieldGroupInstall(unit, v)
}

func FieldGroupStatCreate(unit int, v FieldGroup, stats ...FieldStat) (FieldStatID, error) {
	statLen := len(stats)
	c_stats := make([]C.opennsl_field_stat_t, statLen)
	for index, stat := range stats {
		c_stats[index] = stat.C()
	}

	c_statID := C.int(0)
	rc := C.opennsl_field_stat_create(C.int(unit), v.C(), C.int(statLen), &c_stats[0], &c_statID)
	return FieldStatID(c_statID), ParseError(rc)
}

func (v FieldGroup) StatCreate(unit int, stats ...FieldStat) (FieldStatID, error) {
	return FieldGroupStatCreate(unit, v, stats...)
}

func FieldGroupStatCreateID(unit int, v FieldGroup, statID FieldStatID, stats ...FieldStat) error {
	statLen := len(stats)
	c_stats := make([]C.opennsl_field_stat_t, statLen)
	for index, stat := range stats {
		c_stats[index] = stat.C()
	}

	rc := C.opennsl_field_stat_create_id(C.int(unit), v.C(), C.int(statLen), &c_stats[0], statID.C())
	return ParseError(rc)
}

func (v FieldGroup) StatCreateID(unit int, statID FieldStatID, stats ...FieldStat) error {
	return FieldGroupStatCreateID(unit, v, statID, stats...)
}

func FieldGroupStatIDGet(unit int, v FieldGroup, statID FieldStatID) (uint32, error) {
	c_counterID := C.uint32(0)
	rc := C.opennsl_field_stat_id_get(C.int(unit), v.C(), C.uint32(statID.C()), &c_counterID)
	return uint32(c_counterID), ParseError(rc)
}

func (v FieldGroup) StatIDGet(unit int, statID FieldStatID) (uint32, error) {
	return FieldGroupStatIDGet(unit, v, statID)
}
