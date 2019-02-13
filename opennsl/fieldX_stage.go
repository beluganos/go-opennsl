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

import (
	"fmt"
)

//
// FieldStage
//
type FieldStage C.opennsl_field_stage_t

func (v FieldStage) C() C.opennsl_field_stage_t {
	return C.opennsl_field_stage_t(v)
}

const (
	FieldStageFirst              FieldStage = C.opennslFieldStageFirst
	FieldStageIngressEarly       FieldStage = C.opennslFieldStageIngressEarly
	FieldStageIngressLate        FieldStage = C.opennslFieldStageIngressLate
	FieldStageDefault            FieldStage = C.opennslFieldStageDefault
	FieldStageLast               FieldStage = C.opennslFieldStageLast
	FieldStageIngress            FieldStage = C.opennslFieldStageIngress
	FieldStageEgress             FieldStage = C.opennslFieldStageEgress
	FieldStageExternal           FieldStage = C.opennslFieldStageExternal
	FieldStageHash               FieldStage = C.opennslFieldStageHash
	FieldStageIngressExactMatch  FieldStage = C.opennslFieldStageIngressExactMatch
	FieldStageIngressFlowtracker FieldStage = C.opennslFieldStageIngressFlowtracker
	FieldStageIngressPMF1        FieldStage = C.opennslFieldStageIngressPMF1
	FieldStageIngressPMF2        FieldStage = C.opennslFieldStageIngressPMF2
	FieldStageIngressPMF3        FieldStage = C.opennslFieldStageIngressPMF3
)

const FieldStageCount = C.opennslFieldStageCount

var fieldStage_names = map[FieldStage]string{
	FieldStageFirst:              "First",
	FieldStageIngressEarly:       "IngressEarly",
	FieldStageIngressLate:        "IngressLate",
	FieldStageDefault:            "Default",
	FieldStageLast:               "Last",
	FieldStageIngress:            "Ingress",
	FieldStageEgress:             "Egress",
	FieldStageExternal:           "External",
	FieldStageHash:               "Hash",
	FieldStageIngressExactMatch:  "IngressExactMatch",
	FieldStageIngressFlowtracker: "IngressFlowtracker",
	FieldStageIngressPMF1:        "IngressPMF1",
	FieldStageIngressPMF2:        "IngressPMF2",
	FieldStageIngressPMF3:        "IngressPMF3",
}

var fieldStage_values = map[string]FieldStage{
	"First":              FieldStageFirst,
	"IngressEarly":       FieldStageIngressEarly,
	"IngressLate":        FieldStageIngressLate,
	"Default":            FieldStageDefault,
	"Last":               FieldStageLast,
	"Ingress":            FieldStageIngress,
	"Egress":             FieldStageEgress,
	"External":           FieldStageExternal,
	"Hash":               FieldStageHash,
	"IngressExactMatch":  FieldStageIngressExactMatch,
	"IngressFlowtracker": FieldStageIngressFlowtracker,
	"IngressPMF1":        FieldStageIngressPMF1,
	"IngressPMF2":        FieldStageIngressPMF2,
	"IngressPMF3":        FieldStageIngressPMF3,
}

func (v FieldStage) String() string {
	if s, ok := fieldStage_names[v]; ok {
		return s
	}
	return fmt.Sprintf("FieldStage(%d)", v)
}

func ParseFieldStage(s string) (FieldStage, error) {
	if v, ok := fieldStage_values[s]; ok {
		return v, nil
	}
	return FieldStageFirst, fmt.Errorf("Invalid FieldStage. %s", s)
}
