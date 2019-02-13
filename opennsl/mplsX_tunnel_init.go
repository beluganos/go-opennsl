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

//
// opennsl_mpls_tunnel_initiator_XXX
//
func (v L3IfaceID) MplsTunInitiatorCreate(unit int, labels ...*MplsEgressLabel) error {
	labelNum := len(labels)
	c_labels := make([]C.opennsl_mpls_egress_label_t, labelNum)
	for index, label := range labels {
		c_labels[index] = *label.C()
	}

	rc := func() C.int {
		if InfoDeviceIsDNX(unit) {
			return C.opennsl_mpls_tunnel_initiator_create(C.int(unit), v.C(), C.int(labelNum), &c_labels[0])
		}
		return C.opennsl_mpls_tunnel_initiator_set(C.int(unit), v.C(), C.int(labelNum), &c_labels[0])
	}()
	return ParseError(rc)
}

func (v L3IfaceID) MplsTunnelInitiatorClear(unit int) error {
	rc := C.opennsl_mpls_tunnel_initiator_clear(C.int(unit), v.C())
	return ParseError(rc)
}

func MplsTunnelInitiatorClearAll(unit int) error {
	rc := C.opennsl_mpls_tunnel_initiator_clear_all(C.int(unit))
	return ParseError(rc)
}

func (v L3IfaceID) MplsTunnelInitiatorGet(unit int, labelMax int) ([]MplsEgressLabel, error) {
	c_count := C.int(0)
	c_labels := make([]C.opennsl_mpls_egress_label_t, labelMax)

	rc := C.opennsl_mpls_tunnel_initiator_get(C.int(unit), v.C(), C.int(labelMax), &c_labels[0], &c_count)
	if err := ParseError(rc); err != nil {
		return nil, err
	}

	labels := make([]MplsEgressLabel, int(c_count))
	for index := 0; index < int(c_count); index++ {
		labels[index] = MplsEgressLabel(c_labels[index])
	}

	return labels, nil
}
