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

type FieldActionInterface interface {
	Action() FieldAction
	Param0() uint32
	Param1() uint32
}

type FieldActionParam struct {
	action FieldAction
	param0 uint32
	param1 uint32
}

func (v *FieldActionParam) Action() FieldAction {
	return v.action
}

func (v *FieldActionParam) Param0() uint32 {
	return v.param0
}

func (v *FieldActionParam) Param1() uint32 {
	return v.param1
}

func NewFieldActionParam(a FieldAction, p0, p1 uint32) FieldActionInterface {
	return &FieldActionParam{
		action: a,
		param0: p0,
		param1: p1,
	}
}

func NewFieldActionCosQNew(cosq uint32) FieldActionInterface {
	return NewFieldActionParam(FieldActionCosQNew, cosq, 0)
}

func NewFieldActionCosQCpuNew(cosq uint32) FieldActionInterface {
	return NewFieldActionParam(FieldActionCosQCpuNew, cosq, 0)
}

func NewFieldActionPrioIntNew(priority uint32) FieldActionInterface {
	return NewFieldActionParam(FieldActionPrioIntNew, priority, 0)
}

func NewFieldActionDscpNew(dscp uint32) FieldActionInterface {
	return NewFieldActionParam(FieldActionDscpNew, dscp, 0)
}

func NewFieldActionDscpCancel() FieldActionInterface {
	return NewFieldActionParam(FieldActionDscpCancel, 0, 0)
}

func NewFieldActionCopyToCpu() FieldActionInterface {
	return NewFieldActionParam(FieldActionCopyToCpu, 0, 0)
}

func NewFieldActionCopyToCpuCancel() FieldActionInterface {
	return NewFieldActionParam(FieldActionCopyToCpuCancel, 0, 0)
}

func NwqFieldActionRedirectPort(modid, dstPort uint32) FieldActionInterface {
	return NewFieldActionParam(FieldActionRedirectPort, modid, dstPort)
}

func NewFieldActionRedirectTrunk(dstTrunk uint32) FieldActionInterface {
	return NewFieldActionParam(FieldActionRedirectTrunk, dstTrunk, 0)
}

func NewFieldActionDrop() FieldActionInterface {
	return NewFieldActionParam(FieldActionDrop, 0, 0)
}

func NewFieldActionDropCancel() FieldActionInterface {
	return NewFieldActionParam(FieldActionDropCancel, 0, 0)
}

func NewFieldActionMirrorOverride() FieldActionInterface {
	return NewFieldActionParam(FieldActionMirrorOverride, 0, 0)
}

func NewFieldActionMirrorIngress(dstModId, dstPort uint32) FieldActionInterface {
	return NewFieldActionParam(FieldActionMirrorIngress, dstModId, dstPort)
}

func NewFieldActionMirrorEgress(dstModId, dstPort uint32) FieldActionInterface {
	return NewFieldActionParam(FieldActionMirrorEgress, dstModId, dstPort)
}

func NewFieldActionL3Switch(ecmp uint32) FieldActionInterface {
	return NewFieldActionParam(FieldActionL3Switch, ecmp, 0)
}

func NewFieldActionRpDrop() FieldActionInterface {
	return NewFieldActionParam(FieldActionRpDrop, 0, 0)
}

func NewFieldActionRpDropCancel() FieldActionInterface {
	return NewFieldActionParam(FieldActionRpDropCancel, 0, 0)
}

func NewFieldActionRpCopyToCpu() FieldActionInterface {
	return NewFieldActionParam(FieldActionRpCopyToCpu, 0, 0)
}

func NewFieldActionRpCopyToCpuCancel() FieldActionInterface {
	return NewFieldActionParam(FieldActionRpCopyToCpuCancel, 0, 0)
}

func NewFieldActionRpDscpNew(dscp uint32) FieldActionInterface {
	return NewFieldActionParam(FieldActionRpDscpNew, dscp, 0)
}

func NewFieldActionRpDscpCancel() FieldActionInterface {
	return NewFieldActionParam(FieldActionRpDscpCancel, 0, 0)
}

func NewFieldActionRpCosQNew(cosq uint32) FieldActionInterface {
	return NewFieldActionParam(FieldActionRpCosQNew, cosq, 0)
}

func NewFieldActionYpDrop() FieldActionInterface {
	return NewFieldActionParam(FieldActionYpDrop, 0, 0)
}

func NewFieldActionYpDropCancel() FieldActionInterface {
	return NewFieldActionParam(FieldActionYpDropCancel, 0, 0)
}

func NewFieldActionYpCopyToCpu() FieldActionInterface {
	return NewFieldActionParam(FieldActionYpCopyToCpu, 0, 0)
}

func NewFieldActionYpCopyToCpuCancel() FieldActionInterface {
	return NewFieldActionParam(FieldActionYpCopyToCpuCancel, 0, 0)
}

func NewFieldActionYpDscpNew(dscp uint32) FieldActionInterface {
	return NewFieldActionParam(FieldActionYpDscpNew, dscp, 0)
}

func NewFieldActionYpDscpCancel() FieldActionInterface {
	return NewFieldActionParam(FieldActionYpDscpCancel, 0, 0)
}

func NewFieldActionYpCosQNew(cosq uint32) FieldActionInterface {
	return NewFieldActionParam(FieldActionYpCosQNew, cosq, 0)
}

func NewFieldActionGpDrop() FieldActionInterface {
	return NewFieldActionParam(FieldActionGpDrop, 0, 0)
}

func NewFieldActionGpDropCancel() FieldActionInterface {
	return NewFieldActionParam(FieldActionGpDropCancel, 0, 0)
}

func NewFieldActionGpCopyToCpu() FieldActionInterface {
	return NewFieldActionParam(FieldActionGpCopyToCpu, 0, 0)
}

func NewFieldActionGpCopyToCpuCancel() FieldActionInterface {
	return NewFieldActionParam(FieldActionGpCopyToCpuCancel, 0, 0)
}

func NewFieldActionGpDscpNew(dscp uint32) FieldActionInterface {
	return NewFieldActionParam(FieldActionGpDscpNew, dscp, 0)
}

func NewFieldActionGpDscpCancel() FieldActionInterface {
	return NewFieldActionParam(FieldActionGpDscpCancel, 0, 0)
}

func NewFieldActionGpCosQNew(cosq uint32) FieldActionInterface {
	return NewFieldActionParam(FieldActionGpCosQNew, cosq, 0)
}

func NewFieldActionClassDestSet() FieldActionInterface {
	return NewFieldActionParam(FieldActionClassDestSet, 0, 0)
}

func NewFieldActionClassSet() FieldActionInterface {
	return NewFieldActionParam(FieldActionClassSet, 0, 0)
}

func NewFieldActionClassSourceSet() FieldActionInterface {
	return NewFieldActionParam(FieldActionClassSourceSet, 0, 0)
}

func NewFieldActionNewClassId() FieldActionInterface {
	return NewFieldActionParam(FieldActionNewClassId, 0, 0)
}

func NewFieldActionRpRedirectPort(modid, dstPort uint32) FieldActionInterface {
	return NewFieldActionParam(FieldActionRpRedirectPort, modid, dstPort)
}

func NewFieldActionRpMirrorIngress(modid, dstPort uint32) FieldActionInterface {
	return NewFieldActionParam(FieldActionRpMirrorIngress, modid, dstPort)
}

func NewFieldActionGpRedirectPort(modid, dstPort uint32) FieldActionInterface {
	return NewFieldActionParam(FieldActionGpRedirectPort, modid, dstPort)
}

func NewFieldActionGpMirrorIngress(modid, dstPort uint32) FieldActionInterface {
	return NewFieldActionParam(FieldActionGpMirrorIngress, modid, dstPort)
}

func NewFieldActionEgressClassSelect() FieldActionInterface {
	return NewFieldActionParam(FieldActionEgressClassSelect, 0, 0)
}

func NewFieldActionStat0(statId uint32) FieldActionInterface {
	return NewFieldActionParam(FieldActionStat0, statId, 0)
}

func NewFieldActionStat(statId uint32) FieldActionInterface {
	return NewFieldActionStat0(statId)
}

func NewFieldActionPolicerLevel0(policierId uint32) FieldActionInterface {
	return NewFieldActionParam(FieldActionPolicerLevel0, policierId, 0)
}

func NewFieldActionUsePolicerResult(flags uint32) FieldActionInterface {
	return NewFieldActionParam(FieldActionUsePolicerResult, flags, 0)
}

func NewFieldActionSnoop() FieldActionInterface {
	return NewFieldActionParam(FieldActionSnoop, 0, 0)
}

func NewFieldActionYpMirrorIngress(modid, dstPort uint32) FieldActionInterface {
	return NewFieldActionParam(FieldActionYpMirrorIngress, modid, dstPort)
}

func NewFieldActionYpRedirectPort(modid, dstPort uint32) FieldActionInterface {
	return NewFieldActionParam(FieldActionYpRedirectPort, modid, dstPort)
}

func NewFieldActionIngSampleEnable() FieldActionInterface {
	return NewFieldActionParam(FieldActionIngSampleEnable, 0, 0)
}

func NewFieldActionEgrSampleEnable() FieldActionInterface {
	return NewFieldActionParam(FieldActionEgrSampleEnable, 0, 0)
}
