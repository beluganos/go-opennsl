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
)

//
// FieldAction
//
type FieldAction C.opennsl_field_action_t

func (v FieldAction) C() C.opennsl_field_action_t {
	return C.opennsl_field_action_t(v)
}

const (
	FieldActionCosQNew           FieldAction = C.opennslFieldActionCosQNew
	FieldActionCosQCpuNew        FieldAction = C.opennslFieldActionCosQCpuNew
	FieldActionPrioIntNew        FieldAction = C.opennslFieldActionPrioIntNew
	FieldActionDscpNew           FieldAction = C.opennslFieldActionDscpNew
	FieldActionDscpCancel        FieldAction = C.opennslFieldActionDscpCancel
	FieldActionCopyToCpu         FieldAction = C.opennslFieldActionCopyToCpu
	FieldActionCopyToCpuCancel   FieldAction = C.opennslFieldActionCopyToCpuCancel
	FieldActionRedirectPort      FieldAction = C.opennslFieldActionRedirectPort
	FieldActionRedirectTrunk     FieldAction = C.opennslFieldActionRedirectTrunk
	FieldActionDrop              FieldAction = C.opennslFieldActionDrop
	FieldActionDropCancel        FieldAction = C.opennslFieldActionDropCancel
	FieldActionMirrorOverride    FieldAction = C.opennslFieldActionMirrorOverride
	FieldActionMirrorIngress     FieldAction = C.opennslFieldActionMirrorIngress
	FieldActionMirrorEgress      FieldAction = C.opennslFieldActionMirrorEgress
	FieldActionL3Switch          FieldAction = C.opennslFieldActionL3Switch
	FieldActionRpDrop            FieldAction = C.opennslFieldActionRpDrop
	FieldActionRpDropCancel      FieldAction = C.opennslFieldActionRpDropCancel
	FieldActionRpCopyToCpu       FieldAction = C.opennslFieldActionRpCopyToCpu
	FieldActionRpCopyToCpuCancel FieldAction = C.opennslFieldActionRpCopyToCpuCancel
	FieldActionRpDscpNew         FieldAction = C.opennslFieldActionRpDscpNew
	FieldActionRpDscpCancel      FieldAction = C.opennslFieldActionRpDscpCancel
	FieldActionRpCosQNew         FieldAction = C.opennslFieldActionRpCosQNew
	FieldActionYpDrop            FieldAction = C.opennslFieldActionYpDrop
	FieldActionYpDropCancel      FieldAction = C.opennslFieldActionYpDropCancel
	FieldActionYpCopyToCpu       FieldAction = C.opennslFieldActionYpCopyToCpu
	FieldActionYpCopyToCpuCancel FieldAction = C.opennslFieldActionYpCopyToCpuCancel
	FieldActionYpDscpNew         FieldAction = C.opennslFieldActionYpDscpNew
	FieldActionYpDscpCancel      FieldAction = C.opennslFieldActionYpDscpCancel
	FieldActionYpCosQNew         FieldAction = C.opennslFieldActionYpCosQNew
	FieldActionClassDestSet      FieldAction = C.opennslFieldActionClassDestSet
	FieldActionClassSet          FieldAction = C.opennslFieldActionClassSet
	FieldActionClassSourceSet    FieldAction = C.opennslFieldActionClassSourceSet
	FieldActionGpDrop            FieldAction = C.opennslFieldActionGpDrop
	FieldActionGpDropCancel      FieldAction = C.opennslFieldActionGpDropCancel
	FieldActionGpCopyToCpu       FieldAction = C.opennslFieldActionGpCopyToCpu
	FieldActionGpCopyToCpuCancel FieldAction = C.opennslFieldActionGpCopyToCpuCancel
	FieldActionGpDscpNew         FieldAction = C.opennslFieldActionGpDscpNew
	FieldActionGpDscpCancel      FieldAction = C.opennslFieldActionGpDscpCancel
	FieldActionGpCosQNew         FieldAction = C.opennslFieldActionGpCosQNew
	FieldActionNewClassId        FieldAction = C.opennslFieldActionNewClassId
	FieldActionRpRedirectPort    FieldAction = C.opennslFieldActionRpRedirectPort
	FieldActionRpMirrorIngress   FieldAction = C.opennslFieldActionRpMirrorIngress
	FieldActionGpRedirectPort    FieldAction = C.opennslFieldActionGpRedirectPort
	FieldActionGpMirrorIngress   FieldAction = C.opennslFieldActionGpMirrorIngress
	FieldActionEgressClassSelect FieldAction = C.opennslFieldActionEgressClassSelect
	FieldActionStat0             FieldAction = C.opennslFieldActionStat0
	FieldActionStat              FieldAction = C.opennslFieldActionStat
	FieldActionPolicerLevel0     FieldAction = C.opennslFieldActionPolicerLevel0
	FieldActionUsePolicerResult  FieldAction = C.opennslFieldActionUsePolicerResult
	FieldActionSnoop             FieldAction = C.opennslFieldActionSnoop
	FieldActionYpMirrorIngress   FieldAction = C.opennslFieldActionYpMirrorIngress
	FieldActionYpRedirectPort    FieldAction = C.opennslFieldActionYpRedirectPort
	FieldActionIngSampleEnable   FieldAction = C.opennslFieldActionIngSampleEnable
	FieldActionEgrSampleEnable   FieldAction = C.opennslFieldActionEgrSampleEnable
)

const FieldActionCount = C.opennslFieldActionCount

var fieldAction_names = map[FieldAction]string{
	FieldActionCosQNew:           "CosQNew",
	FieldActionCosQCpuNew:        "CosQCpuNew",
	FieldActionPrioIntNew:        "PrioIntNew",
	FieldActionDscpNew:           "DscpNew",
	FieldActionDscpCancel:        "DscpCancel",
	FieldActionCopyToCpu:         "CopyToCpu",
	FieldActionCopyToCpuCancel:   "CopyToCpuCancel",
	FieldActionRedirectPort:      "RedirectPort",
	FieldActionRedirectTrunk:     "RedirectTrunk",
	FieldActionDrop:              "Drop",
	FieldActionDropCancel:        "DropCancel",
	FieldActionMirrorOverride:    "MirrorOverride",
	FieldActionMirrorIngress:     "MirrorIngress",
	FieldActionMirrorEgress:      "MirrorEgress",
	FieldActionL3Switch:          "L3Switch",
	FieldActionRpDrop:            "RpDrop",
	FieldActionRpDropCancel:      "RpDropCancel",
	FieldActionRpCopyToCpu:       "RpCopyToCpu",
	FieldActionRpCopyToCpuCancel: "RpCopyToCpuCancel",
	FieldActionRpDscpNew:         "RpDscpNew",
	FieldActionRpDscpCancel:      "RpDscpCancel",
	FieldActionRpCosQNew:         "RpCosQNew",
	FieldActionYpDrop:            "YpDrop",
	FieldActionYpDropCancel:      "YpDropCancel",
	FieldActionYpCopyToCpu:       "YpCopyToCpu",
	FieldActionYpCopyToCpuCancel: "YpCopyToCpuCancel",
	FieldActionYpDscpNew:         "YpDscpNew",
	FieldActionYpDscpCancel:      "YpDscpCancel",
	FieldActionYpCosQNew:         "YpCosQNew",
	FieldActionClassDestSet:      "ClassDestSet",
	FieldActionClassSourceSet:    "ClassSourceSet",
	FieldActionGpDrop:            "GpDrop",
	FieldActionGpDropCancel:      "GpDropCancel",
	FieldActionGpCopyToCpu:       "GpCopyToCpu",
	FieldActionGpCopyToCpuCancel: "GpCopyToCpuCancel",
	FieldActionGpDscpNew:         "GpDscpNew",
	FieldActionGpDscpCancel:      "GpDscpCancel",
	FieldActionGpCosQNew:         "GpCosQNew",
	FieldActionNewClassId:        "NewClassId",
	FieldActionRpRedirectPort:    "RpRedirectPort",
	FieldActionRpMirrorIngress:   "RpMirrorIngress",
	FieldActionGpRedirectPort:    "GpRedirectPort",
	FieldActionGpMirrorIngress:   "GpMirrorIngress",
	FieldActionEgressClassSelect: "EgressClassSelect",
	FieldActionStat:              "Stat",
	FieldActionPolicerLevel0:     "PolicerLevel0",
	FieldActionUsePolicerResult:  "UsePolicerResult",
	FieldActionSnoop:             "Snoop",
	FieldActionYpMirrorIngress:   "YpMirrorIngress",
	FieldActionYpRedirectPort:    "YpRedirectPort",
	FieldActionIngSampleEnable:   "IngSampleEnable",
	FieldActionEgrSampleEnable:   "EgrSampleEnable",
}

var fieldAction_values = map[string]FieldAction{
	"CosQNew":           FieldActionCosQNew,
	"CosQCpuNew":        FieldActionCosQCpuNew,
	"PrioIntNew":        FieldActionPrioIntNew,
	"DscpNew":           FieldActionDscpNew,
	"DscpCancel":        FieldActionDscpCancel,
	"CopyToCpu":         FieldActionCopyToCpu,
	"CopyToCpuCancel":   FieldActionCopyToCpuCancel,
	"RedirectPort":      FieldActionRedirectPort,
	"RedirectTrunk":     FieldActionRedirectTrunk,
	"Drop":              FieldActionDrop,
	"DropCancel":        FieldActionDropCancel,
	"MirrorOverride":    FieldActionMirrorOverride,
	"MirrorIngress":     FieldActionMirrorIngress,
	"MirrorEgress":      FieldActionMirrorEgress,
	"L3Switch":          FieldActionL3Switch,
	"RpDrop":            FieldActionRpDrop,
	"RpDropCancel":      FieldActionRpDropCancel,
	"RpCopyToCpu":       FieldActionRpCopyToCpu,
	"RpCopyToCpuCancel": FieldActionRpCopyToCpuCancel,
	"RpDscpNew":         FieldActionRpDscpNew,
	"RpDscpCancel":      FieldActionRpDscpCancel,
	"RpCosQNew":         FieldActionRpCosQNew,
	"YpDrop":            FieldActionYpDrop,
	"YpDropCancel":      FieldActionYpDropCancel,
	"YpCopyToCpu":       FieldActionYpCopyToCpu,
	"YpCopyToCpuCancel": FieldActionYpCopyToCpuCancel,
	"YpDscpNew":         FieldActionYpDscpNew,
	"YpDscpCancel":      FieldActionYpDscpCancel,
	"YpCosQNew":         FieldActionYpCosQNew,
	"ClassDestSet":      FieldActionClassDestSet,
	"ClassSet":          FieldActionClassSet,
	"ClassSourceSet":    FieldActionClassSourceSet,
	"GpDrop":            FieldActionGpDrop,
	"GpDropCancel":      FieldActionGpDropCancel,
	"GpCopyToCpu":       FieldActionGpCopyToCpu,
	"GpCopyToCpuCancel": FieldActionGpCopyToCpuCancel,
	"GpDscpNew":         FieldActionGpDscpNew,
	"GpDscpCancel":      FieldActionGpDscpCancel,
	"GpCosQNew":         FieldActionGpCosQNew,
	"NewClassId":        FieldActionNewClassId,
	"RpRedirectPort":    FieldActionRpRedirectPort,
	"RpMirrorIngress":   FieldActionRpMirrorIngress,
	"GpRedirectPort":    FieldActionGpRedirectPort,
	"GpMirrorIngress":   FieldActionGpMirrorIngress,
	"EgressClassSelect": FieldActionEgressClassSelect,
	"Stat":              FieldActionStat,
	"Stat0":             FieldActionStat0,
	"PolicerLevel0":     FieldActionPolicerLevel0,
	"UsePolicerResult":  FieldActionUsePolicerResult,
	"Snoop":             FieldActionSnoop,
	"YpMirrorIngress":   FieldActionYpMirrorIngress,
	"YpRedirectPort":    FieldActionYpRedirectPort,
	"IngSampleEnable":   FieldActionIngSampleEnable,
	"EgrSampleEnable":   FieldActionEgrSampleEnable,
}

func (v FieldAction) String() string {
	if s, ok := fieldAction_names[v]; ok {
		return s
	}
	return fmt.Sprintf("FieldAction(%d)", v)
}

func ParseFieldAction(s string) (FieldAction, error) {
	if v, ok := fieldAction_values[s]; ok {
		return v, nil
	}
	return 0, fmt.Errorf("Invalid FieldAction. %s", s)
}

//
// FieldActionWidth
//
type FieldActionWidth C.opennsl_field_action_width_t

func (v *FieldActionWidth) C() *C.opennsl_field_action_width_t {
	return (*C.opennsl_field_action_width_t)(v)
}

func (v *FieldActionWidth) Action() uint32 {
	return uint32(v.action)
}

func (v *FieldActionWidth) SetAction(action uint32) {
	v.action = C.uint32(action)
}

func (v *FieldActionWidth) Width() uint32 {
	return uint32(v.width)
}

func (v *FieldActionWidth) SetWidth(width uint32) {
	v.width = C.uint32(width)
}

func (v *FieldActionWidth) Valid() uint8 {
	return uint8(v.valid)
}

func (v *FieldActionWidth) SetValid(valid uint8) {
	v.valid = C.uint8(valid)
}

//
// FieldASet
//
type FieldASet C.opennsl_field_aset_t

func (v *FieldASet) C() *C.opennsl_field_aset_t {
	return (*C.opennsl_field_aset_t)(v)
}

func NewFieldASet() *FieldASet {
	aset := &FieldASet{}
	aset.Init()
	return aset
}

func (v *FieldASet) Init() {
	C._opennsl_field_aset_init(v.C())
}

func (v *FieldASet) Add(a FieldAction) {
	C._opennsl_field_aset_add(v.C(), a.C())
}

func (v *FieldASet) Remove(a FieldAction) {
	C._opennsl_field_aset_remove(v.C(), a.C())
}

func (v *FieldASet) Test(a FieldAction) {
	C._opennsl_field_aset_test(v.C(), a.C())
}
