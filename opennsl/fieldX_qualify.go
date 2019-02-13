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
// FieldQualify
//
type FieldQualify C.opennsl_field_qualify_t

func (v FieldQualify) C() C.opennsl_field_qualify_t {
	return C.opennsl_field_qualify_t(v)
}

const (
	FieldQualifySrcIp6                       FieldQualify = C.opennslFieldQualifySrcIp6
	FieldQualifyDstIp6                       FieldQualify = C.opennslFieldQualifyDstIp6
	FieldQualifySrcMac                       FieldQualify = C.opennslFieldQualifySrcMac
	FieldQualifyDstMac                       FieldQualify = C.opennslFieldQualifyDstMac
	FieldQualifySrcIp                        FieldQualify = C.opennslFieldQualifySrcIp
	FieldQualifyDstIp                        FieldQualify = C.opennslFieldQualifyDstIp
	FieldQualifyInPort                       FieldQualify = C.opennslFieldQualifyInPort
	FieldQualifyInPorts                      FieldQualify = C.opennslFieldQualifyInPorts
	FieldQualifyOuterVlan                    FieldQualify = C.opennslFieldQualifyOuterVlan
	FieldQualifyOuterVlanId                  FieldQualify = C.opennslFieldQualifyOuterVlanId
	FieldQualifyInnerVlanId                  FieldQualify = C.opennslFieldQualifyInnerVlanId
	FieldQualifyRangeCheck                   FieldQualify = C.opennslFieldQualifyRangeCheck
	FieldQualifyL4SrcPort                    FieldQualify = C.opennslFieldQualifyL4SrcPort
	FieldQualifyL4DstPort                    FieldQualify = C.opennslFieldQualifyL4DstPort
	FieldQualifyEtherType                    FieldQualify = C.opennslFieldQualifyEtherType
	FieldQualifyIpProtocol                   FieldQualify = C.opennslFieldQualifyIpProtocol
	FieldQualifyIp6NextHeader                FieldQualify = C.opennslFieldQualifyIp6NextHeader
	FieldQualifyDSCP                         FieldQualify = C.opennslFieldQualifyDSCP
	FieldQualifyTtl                          FieldQualify = C.opennslFieldQualifyTtl
	FieldQualifyIp6HopLimit                  FieldQualify = C.opennslFieldQualifyIp6HopLimit
	FieldQualifySrcPort                      FieldQualify = C.opennslFieldQualifySrcPort
	FieldQualifyDstPort                      FieldQualify = C.opennslFieldQualifyDstPort
	FieldQualifyDstTrunk                     FieldQualify = C.opennslFieldQualifyDstTrunk
	FieldQualifyTcpControl                   FieldQualify = C.opennslFieldQualifyTcpControl
	FieldQualifyPacketRes                    FieldQualify = C.opennslFieldQualifyPacketRes
	FieldQualifySrcClassField                FieldQualify = C.opennslFieldQualifySrcClassField
	FieldQualifyDstClassField                FieldQualify = C.opennslFieldQualifyDstClassField
	FieldQualifyIpProtocolCommon             FieldQualify = C.opennslFieldQualifyIpProtocolCommon
	FieldQualifyIpType                       FieldQualify = C.opennslFieldQualifyIpType
	FieldQualifyStage                        FieldQualify = C.opennslFieldQualifyStage
	FieldQualifyStageIngress                 FieldQualify = C.opennslFieldQualifyStageIngress
	FieldQualifyStageLookup                  FieldQualify = C.opennslFieldQualifyStageLookup
	FieldQualifyStageEgress                  FieldQualify = C.opennslFieldQualifyStageEgress
	FieldQualifyInterfaceClassPort           FieldQualify = C.opennslFieldQualifyInterfaceClassPort
	FieldQualifyL3Routable                   FieldQualify = C.opennslFieldQualifyL3Routable
	FieldQualifyIpFrag                       FieldQualify = C.opennslFieldQualifyIpFrag
	FieldQualifyL3Ingress                    FieldQualify = C.opennslFieldQualifyL3Ingress
	FieldQualifyOutPort                      FieldQualify = C.opennslFieldQualifyOutPort
	FieldQualifyIp4                          FieldQualify = C.opennslFieldQualifyIp4
	FieldQualifyIp6                          FieldQualify = C.opennslFieldQualifyIp6
	FieldQualifyIcmpTypeCode                 FieldQualify = C.opennslFieldQualifyIcmpTypeCode
	FieldQualifyDstL3Egress                  FieldQualify = C.opennslFieldQualifyDstL3Egress
	FieldQualifyColor                        FieldQualify = C.opennslFieldQualifyColor
	FieldQualifyMyStationHit                 FieldQualify = C.opennslFieldQualifyMyStationHit
	FieldQualifyDstIpLocal                   FieldQualify = C.opennslFieldQualifyDstIpLocal
	FieldQualifyCpuQueue                     FieldQualify = C.opennslFieldQualifyCpuQueue
	FieldQualifyInterfaceClassProcessingPort FieldQualify = C.opennslFieldQualifyInterfaceClassProcessingPort
	FieldQualifyIngressClassField            FieldQualify = C.opennslFieldQualifyIngressClassField
)

var fieldQualify_names = map[FieldQualify]string{
	FieldQualifySrcIp6:                       "SrcIp6",
	FieldQualifyDstIp6:                       "DstIp6",
	FieldQualifySrcMac:                       "SrcMac",
	FieldQualifyDstMac:                       "DstMac",
	FieldQualifySrcIp:                        "SrcIp",
	FieldQualifyDstIp:                        "DstIp",
	FieldQualifyInPort:                       "InPort",
	FieldQualifyInPorts:                      "InPorts",
	FieldQualifyOuterVlan:                    "OuterVlan",
	FieldQualifyOuterVlanId:                  "OuterVlanId",
	FieldQualifyInnerVlanId:                  "InnerVlanId",
	FieldQualifyRangeCheck:                   "RangeCheck",
	FieldQualifyL4SrcPort:                    "L4SrcPort",
	FieldQualifyL4DstPort:                    "L4DstPort",
	FieldQualifyEtherType:                    "EtherType",
	FieldQualifyIpProtocol:                   "IpProtocol",
	FieldQualifyDSCP:                         "DSCP",
	FieldQualifyTtl:                          "Ttl",
	FieldQualifySrcPort:                      "SrcPort",
	FieldQualifyDstPort:                      "DstPort",
	FieldQualifyDstTrunk:                     "DstTrunk",
	FieldQualifyTcpControl:                   "TcpControl",
	FieldQualifyPacketRes:                    "PacketRes",
	FieldQualifySrcClassField:                "SrcClassField",
	FieldQualifyDstClassField:                "DstClassField",
	FieldQualifyIpProtocolCommon:             "IpProtocolCommon",
	FieldQualifyIpType:                       "IpType",
	FieldQualifyStage:                        "Stage",
	FieldQualifyStageIngress:                 "StageIngress",
	FieldQualifyStageLookup:                  "StageLookup",
	FieldQualifyStageEgress:                  "StageEgress",
	FieldQualifyInterfaceClassPort:           "InterfaceClassPort",
	FieldQualifyL3Routable:                   "L3Routable",
	FieldQualifyIpFrag:                       "IpFrag",
	FieldQualifyL3Ingress:                    "L3Ingress",
	FieldQualifyOutPort:                      "OutPort",
	FieldQualifyIp4:                          "Ip4",
	FieldQualifyIp6:                          "Ip6",
	FieldQualifyIcmpTypeCode:                 "IcmpTypeCode",
	FieldQualifyDstL3Egress:                  "DstL3Egress",
	FieldQualifyColor:                        "Color",
	FieldQualifyMyStationHit:                 "MyStationHit",
	FieldQualifyDstIpLocal:                   "DstIpLocal",
	FieldQualifyCpuQueue:                     "CpuQueue",
	FieldQualifyInterfaceClassProcessingPort: "InterfaceClassProcessingPort",
	FieldQualifyIngressClassField:            "IngressClassField",
	// FieldQualifyIp6NextHeader:                "Ip6NextHeader",
	// FieldQualifyIp6HopLimit:                  "Ip6HopLimit",
}

var fieldQualify_values = map[string]FieldQualify{
	"SrcIp6":                       FieldQualifySrcIp6,
	"DstIp6":                       FieldQualifyDstIp6,
	"SrcMac":                       FieldQualifySrcMac,
	"DstMac":                       FieldQualifyDstMac,
	"SrcIp":                        FieldQualifySrcIp,
	"DstIp":                        FieldQualifyDstIp,
	"InPort":                       FieldQualifyInPort,
	"InPorts":                      FieldQualifyInPorts,
	"OuterVlan":                    FieldQualifyOuterVlan,
	"OuterVlanId":                  FieldQualifyOuterVlanId,
	"InnerVlanId":                  FieldQualifyInnerVlanId,
	"RangeCheck":                   FieldQualifyRangeCheck,
	"L4SrcPort":                    FieldQualifyL4SrcPort,
	"L4DstPort":                    FieldQualifyL4DstPort,
	"EtherType":                    FieldQualifyEtherType,
	"IpProtocol":                   FieldQualifyIpProtocol,
	"Ip6NextHeader":                FieldQualifyIp6NextHeader,
	"DSCP":                         FieldQualifyDSCP,
	"Ttl":                          FieldQualifyTtl,
	"Ip6HopLimit":                  FieldQualifyIp6HopLimit,
	"SrcPort":                      FieldQualifySrcPort,
	"DstPort":                      FieldQualifyDstPort,
	"DstTrunk":                     FieldQualifyDstTrunk,
	"TcpControl":                   FieldQualifyTcpControl,
	"PacketRes":                    FieldQualifyPacketRes,
	"SrcClassField":                FieldQualifySrcClassField,
	"DstClassField":                FieldQualifyDstClassField,
	"IpProtocolCommon":             FieldQualifyIpProtocolCommon,
	"IpType":                       FieldQualifyIpType,
	"Stage":                        FieldQualifyStage,
	"StageIngress":                 FieldQualifyStageIngress,
	"StageLookup":                  FieldQualifyStageLookup,
	"StageEgress":                  FieldQualifyStageEgress,
	"InterfaceClassPort":           FieldQualifyInterfaceClassPort,
	"L3Routable":                   FieldQualifyL3Routable,
	"IpFrag":                       FieldQualifyIpFrag,
	"L3Ingress":                    FieldQualifyL3Ingress,
	"OutPort":                      FieldQualifyOutPort,
	"Ip4":                          FieldQualifyIp4,
	"Ip6":                          FieldQualifyIp6,
	"IcmpTypeCode":                 FieldQualifyIcmpTypeCode,
	"DstL3Egress":                  FieldQualifyDstL3Egress,
	"Color":                        FieldQualifyColor,
	"MyStationHit":                 FieldQualifyMyStationHit,
	"DstIpLocal":                   FieldQualifyDstIpLocal,
	"CpuQueue":                     FieldQualifyCpuQueue,
	"InterfaceClassProcessingPort": FieldQualifyInterfaceClassProcessingPort,
	"IngressClassField":            FieldQualifyIngressClassField,
}

func (v FieldQualify) String() string {
	if s, ok := fieldQualify_names[v]; ok {
		return s
	}
	return fmt.Sprintf("FieldQualify(%d)", v)
}

func ParseFieldQualify(s string) (FieldQualify, error) {
	if v, ok := fieldQualify_values[s]; ok {
		return v, nil
	}
	return 0, fmt.Errorf("Invalid FieldQualify. %s", s)
}

const FieldQualifyCount = C.opennslFieldQualifyCount
const FIELD_QUALIFY_MAX = C.OPENNSL_FIELD_QUALIFY_MAX
const FIELD_QUALIFY_PRESEL = C.OPENNSL_FIELD_QUALIFY_PRESEL

//
// FieldQSet
//
type FieldQSet C.opennsl_field_qset_t

func (v *FieldQSet) C() *C.opennsl_field_qset_t {
	return (*C.opennsl_field_qset_t)(v)
}

func NewFieldQSet() *FieldQSet {
	qset := &FieldQSet{}
	qset.Init()
	return qset
}

func FieldQSetInit(v *FieldQSet) {
	C._opennsl_field_qset_init(v.C())
}

func (v *FieldQSet) Init() {
	FieldQSetInit(v)
}

func FieldQSetAdd(v *FieldQSet, qs ...FieldQualify) {
	for _, q := range qs {
		C._opennsl_field_qset_add(v.C(), q.C())
	}
}

func (v *FieldQSet) Add(qs ...FieldQualify) {
	FieldQSetAdd(v, qs...)
}

func FieldQSetRemove(v *FieldQSet, qs ...FieldQualify) {
	for _, q := range qs {
		C._opennsl_field_qset_remove(v.C(), q.C())
	}
}

func (v *FieldQSet) Remove(qs ...FieldQualify) {
	FieldQSetRemove(v, qs...)
}

func FieldQSetTest(v *FieldQSet, q FieldQualify) uint32 {
	return uint32(C._opennsl_field_qset_test(v.C(), q.C()))
}

func (v *FieldQSet) Test(q FieldQualify) uint32 {
	return FieldQSetTest(v, q)
}
