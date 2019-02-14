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
#include <opennsl/l3.h>
*/
import "C"

import (
	"fmt"
	"strings"
)

//
// L3Flags
//
type L3Flags uint32

func (v L3Flags) C() C.uint32 {
	return C.uint32(v)
}

func NewL3Flags(flags ...L3Flags) L3Flags {
	v := L3_NONE
	for _, flag := range flags {
		v |= flag
	}
	return v
}

const (
	L3_NONE                 L3Flags = 0
	L3_L2ONLY               L3Flags = C.OPENNSL_L3_L2ONLY
	L3_UNTAG                L3Flags = C.OPENNSL_L3_UNTAG
	L3_S_HIT                L3Flags = C.OPENNSL_L3_S_HIT
	L3_D_HIT                L3Flags = C.OPENNSL_L3_D_HIT
	L3_HIT                  L3Flags = C.OPENNSL_L3_HIT
	L3_HIT_CLEAR            L3Flags = C.OPENNSL_L3_HIT_CLEAR
	L3_ADD_TO_ARL           L3Flags = C.OPENNSL_L3_ADD_TO_ARL
	L3_WITH_ID              L3Flags = C.OPENNSL_L3_WITH_ID
	L3_REPLACE              L3Flags = C.OPENNSL_L3_REPLACE
	L3_TGID                 L3Flags = C.OPENNSL_L3_TGID
	L3_RPE                  L3Flags = C.OPENNSL_L3_RPE
	L3_IPMC                 L3Flags = C.OPENNSL_L3_IPMC
	L3_L2TOCPU              L3Flags = C.OPENNSL_L3_L2TOCPU
	L3_DEFIP_CPU            L3Flags = C.OPENNSL_L3_DEFIP_CPU
	L3_DEFIP_LOCAL          L3Flags = C.OPENNSL_L3_DEFIP_LOCAL
	L3_MULTIPATH            L3Flags = C.OPENNSL_L3_MULTIPATH
	L3_HOST_LOCAL           L3Flags = C.OPENNSL_L3_HOST_LOCAL
	L3_HOST_AS_ROUTE        L3Flags = C.OPENNSL_L3_HOST_AS_ROUTE
	L3_IP6                  L3Flags = C.OPENNSL_L3_IP6
	L3_RPF                  L3Flags = C.OPENNSL_L3_RPF
	L3_DST_DISCARD          L3Flags = C.OPENNSL_L3_DST_DISCARD
	L3_ROUTE_LABEL          L3Flags = C.OPENNSL_L3_ROUTE_LABEL
	L3_COPY_TO_CPU          L3Flags = C.OPENNSL_L3_COPY_TO_CPU
	L3_DEREFERENCED_NEXTHOP L3Flags = C.OPENNSL_L3_DEREFERENCED_NEXTHOP
	L3_ECMP_RH_REPLACE      L3Flags = C.OPENNSL_L3_ECMP_RH_REPLACE
	L3_VXLAN_ONLY           L3Flags = 1 << 31 // )C.OPENNSL_L3_VXLAN_ONLY
)

var l3flags_names = map[L3Flags]string{
	L3_NONE:                 "NONE",
	L3_L2ONLY:               "L2ONLY",
	L3_UNTAG:                "UNTAG",
	L3_S_HIT:                "S_HIT",
	L3_D_HIT:                "D_HIT",
	L3_HIT:                  "HIT",
	L3_HIT_CLEAR:            "HIT_CLEAR",
	L3_ADD_TO_ARL:           "ADD_TO_ARL",
	L3_WITH_ID:              "WITH_ID",
	L3_REPLACE:              "REPLACE",
	L3_TGID:                 "TGID",
	L3_RPE:                  "RPE",
	L3_IPMC:                 "IPMC",
	L3_L2TOCPU:              "L2TOCPU",
	L3_DEFIP_CPU:            "DEFIP_CPU",
	L3_MULTIPATH:            "MULTIPATH",
	L3_HOST_LOCAL:           "HOST_LOCAL",
	L3_HOST_AS_ROUTE:        "HOST_AS_ROUTE",
	L3_IP6:                  "IP6",
	L3_RPF:                  "RPF",
	L3_DST_DISCARD:          "DST_DISCARD",
	L3_ROUTE_LABEL:          "ROUTE_LABEL",
	L3_COPY_TO_CPU:          "COPY_TO_CPU",
	L3_DEREFERENCED_NEXTHOP: "DEREFERENCED_NEXTHOP",
	L3_VXLAN_ONLY:           "VXLAN_ONLY",
	// L3_DEFIP_LOCAL:          "DEFIP_LOCAL",
	// L3_ECMP_RH_REPLACE:      "ECMP_RH_REPLACE",
}

var l3flags_values = map[string]L3Flags{
	"NONE":                 L3_NONE,
	"L2ONLY":               L3_L2ONLY,
	"UNTAG":                L3_UNTAG,
	"S_HIT":                L3_S_HIT,
	"D_HIT":                L3_D_HIT,
	"HIT":                  L3_HIT,
	"HIT_CLEAR":            L3_HIT_CLEAR,
	"ADD_TO_ARL":           L3_ADD_TO_ARL,
	"WITH_ID":              L3_WITH_ID,
	"REPLACE":              L3_REPLACE,
	"TGID":                 L3_TGID,
	"RPE":                  L3_RPE,
	"IPMC":                 L3_IPMC,
	"L2TOCPU":              L3_L2TOCPU,
	"DEFIP_CPU":            L3_DEFIP_CPU,
	"DEFIP_LOCAL":          L3_DEFIP_LOCAL,
	"MULTIPATH":            L3_MULTIPATH,
	"HOST_LOCAL":           L3_HOST_LOCAL,
	"HOST_AS_ROUTE":        L3_HOST_AS_ROUTE,
	"IP6":                  L3_IP6,
	"RPF":                  L3_RPF,
	"DST_DISCARD":          L3_DST_DISCARD,
	"ROUTE_LABEL":          L3_ROUTE_LABEL,
	"COPY_TO_CPU":          L3_COPY_TO_CPU,
	"DEREFERENCED_NEXTHOP": L3_DEREFERENCED_NEXTHOP,
	"ECMP_RH_REPLACE":      L3_ECMP_RH_REPLACE,
	"VXLAN_ONLY":           L3_VXLAN_ONLY,
}

func (v L3Flags) String() string {
	names := make([]string, 0, len(l3flags_names))
	for value, name := range l3flags_names {
		if v&value != 0 {
			names = append(names, name)
		}
	}
	return strings.Join(names, "|")
}

func ParseL3Flags(name string) (L3Flags, error) {
	if v, ok := l3flags_values[name]; ok {
		return v, nil
	}
	return L3_NONE, fmt.Errorf("Invalid L3Flags. %s", name)
}

//
// L3IngressFlags
//
type L3IngressFlags C.uint32

func (v L3IngressFlags) C() C.uint32 {
	return C.uint32(v)
}

func NewL3IngressFlags(flags ...L3IngressFlags) L3IngressFlags {
	v := L3_INGRESS_NONE
	for _, flag := range flags {
		v |= flag
	}
	return v
}

const (
	L3_INGRESS_NONE         L3IngressFlags = 0
	L3_INGRESS_WITH_ID      L3IngressFlags = C.OPENNSL_L3_INGRESS_WITH_ID
	L3_INGRESS_REPLACE      L3IngressFlags = C.OPENNSL_L3_INGRESS_REPLACE
	L3_INGRESS_GLOBAL_ROUTE L3IngressFlags = C.OPENNSL_L3_INGRESS_GLOBAL_ROUTE
	L3_INGRESS_DSCP_TRUST   L3IngressFlags = C.OPENNSL_L3_INGRESS_DSCP_TRUST
)

var l3ingressFlags_names = map[L3IngressFlags]string{
	L3_INGRESS_WITH_ID:      "WITH_ID",
	L3_INGRESS_REPLACE:      "REPLACE",
	L3_INGRESS_GLOBAL_ROUTE: "GLOBAL_ROUTE",
	L3_INGRESS_DSCP_TRUST:   "DSCP_TRUST",
}

var l3ingressFlags_values = map[string]L3IngressFlags{
	"WITH_ID":      L3_INGRESS_WITH_ID,
	"REPLACE":      L3_INGRESS_REPLACE,
	"GLOBAL_ROUTE": L3_INGRESS_GLOBAL_ROUTE,
	"DSCP_TRUST":   L3_INGRESS_DSCP_TRUST,
}

func (v L3IngressFlags) String() string {
	names := make([]string, 0, len(l3ingressFlags_names))
	for value, name := range l3ingressFlags_names {
		if v&value != 0 {
			names = append(names, name)
		}
	}
	return strings.Join(names, "|")
}

func ParseL3IngressFlags(name string) (L3IngressFlags, error) {
	if v, ok := l3ingressFlags_values[name]; ok {
		return v, nil
	}
	return L3_INGRESS_NONE, fmt.Errorf("Invalid L3IngressFlags. %s", name)
}

//
// L3IfaceQoSFlags
//
type L3IfaceQoSFlags uint32

func (v L3IfaceQoSFlags) C() C.uint32 {
	return C.uint32(v)
}

//
// L3IfaceQoS
//
type L3IfaceQoS C.opennsl_l3_intf_qos_t

func (v *L3IfaceQoS) C() *C.opennsl_l3_intf_qos_t {
	return (*C.opennsl_l3_intf_qos_t)(v)
}

func (v *L3IfaceQoS) Flags() L3IfaceQoSFlags {
	return L3IfaceQoSFlags(v.flags)
}

func (v *L3IfaceQoS) SetFlags(flags L3IfaceQoSFlags) {
	v.flags = flags.C()
}

func (v *L3IfaceQoS) QosMapID() int {
	return int(v.qos_map_id)
}

func (v *L3IfaceQoS) SetQosMapID(qosMapID int) {
	v.qos_map_id = C.int(qosMapID)
}

func (v *L3IfaceQoS) Pri() uint8 {
	return uint8(v.pri)
}

func (v *L3IfaceQoS) SetPri(pri uint8) {
	v.pri = C.uint8(pri)
}

func (v *L3IfaceQoS) Cfi() uint8 {
	return uint8(v.cfi)
}

func (v *L3IfaceQoS) SetCfi(cfi uint8) {
	v.cfi = C.uint8(cfi)
}

func (v *L3IfaceQoS) Dscp() int {
	return int(v.dscp)
}

func (v *L3IfaceQoS) SetDscp(dscp int) {
	v.dscp = C.int(dscp)
}

//
// OPENNSL_L3_ECMP_DYNAMIC_*
//
const (
	L3_ECMP_DYNAMIC_SCALING_FACTOR_INVALID = C.OPENNSL_L3_ECMP_DYNAMIC_SCALING_FACTOR_INVALID
	L3_ECMP_DYNAMIC_LOAD_WEIGHT_INVALID    = C.OPENNSL_L3_ECMP_DYNAMIC_LOAD_WEIGHT_INVALID
)

type L3EcmpDynamicMode uint32

func (v L3EcmpDynamicMode) C() C.uint32 {
	return C.uint32(v)
}

const (
	L3_ECMP_DYNAMIC_MODE_NONE      L3EcmpDynamicMode = 0
	L3_ECMP_DYNAMIC_MODE_NORMAL    L3EcmpDynamicMode = C.OPENNSL_L3_ECMP_DYNAMIC_MODE_NORMAL
	L3_ECMP_DYNAMIC_MODE_RESILIENT L3EcmpDynamicMode = C.OPENNSL_L3_ECMP_DYNAMIC_MODE_RESILIENT
)

//
// L3IngressUrpf
//
type L3IngressUrpf C.opennsl_l3_ingress_urpf_mode_t

func (v L3IngressUrpf) C() C.opennsl_l3_ingress_urpf_mode_t {
	return C.opennsl_l3_ingress_urpf_mode_t(v)
}

const (
	L3IngressUrpfDisable L3IngressUrpf = C.opennslL3IngressUrpfDisable
	L3IngressUrpfLoose   L3IngressUrpf = C.opennslL3IngressUrpfLoose
	L3IngressUrpfStrict  L3IngressUrpf = C.opennslL3IngressUrpfStrict
)

func L3Init(unit int) error {
	rc := C.opennsl_l3_init(C.int(unit))
	return ParseError(rc)
}
