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
#include <opennsl/mplsX.h>
*/
import "C"

func MplsInit(unit int) error {
	rc := C.opennsl_mpls_init(C.int(unit))
	return ParseError(rc)
}

func MplsCleanup(unit int) error {
	rc := C.opennsl_mpls_cleanup(C.int(unit))
	return ParseError(rc)
}

//
// MplsExpMap
//
type MplsExpMap C.opennsl_mpls_exp_map_t

func (v *MplsExpMap) C() *C.opennsl_mpls_exp_map_t {
	return (*C.opennsl_mpls_exp_map_t)(v)
}

func (v *MplsExpMap) Priority() int {
	return int(v.priority)
}

func (v *MplsExpMap) SetPriority(priority int) {
	v.priority = C.int(priority)
}

func (v *MplsExpMap) DSCP() uint8 {
	return uint8(v.dscp)
}

func (v *MplsExpMap) SetDSCP(dscp uint8) {
	v.dscp = C.uint8(dscp)
}

func (v *MplsExpMap) Exp() uint8 {
	return uint8(v.exp)
}

func (v *MplsExpMap) SetExp(exp uint8) {
	v.exp = C.uint8(exp)
}

func (v *MplsExpMap) PktPri() uint8 {
	return uint8(v.pkt_pri)
}

func (v *MplsExpMap) SetPktPri(pri uint8) {
	v.pkt_pri = C.uint8(pri)
}

func (v *MplsExpMap) PktCfi() uint8 {
	return uint8(v.pkt_cfi)
}

func (v *MplsExpMap) SetPktCfi(cfi uint8) {
	v.pkt_cfi = C.uint8(cfi)
}

func (v *MplsExpMap) Init() {
	C.opennsl_mpls_exp_map_t_init(v.C())
}
