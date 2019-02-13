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
// MplsEntropyIdentifier
//
type MplsEntropyIdentifier C.opennsl_mpls_entropy_identifier_t

func (v *MplsEntropyIdentifier) C() *C.opennsl_mpls_entropy_identifier_t {
	return (*C.opennsl_mpls_entropy_identifier_t)(v)
}

func (v *MplsEntropyIdentifier) Label() MplsLabel {
	return MplsLabel(v.label)
}

func (v *MplsEntropyIdentifier) SetLabel(label MplsLabel) {
	v.label = label.C()
}

func (v *MplsEntropyIdentifier) Mask() MplsLabel {
	return MplsLabel(v.mask)
}

func (v *MplsEntropyIdentifier) SetMask(mask MplsLabel) {
	v.mask = mask.C()
}

func (v *MplsEntropyIdentifier) Pri() int {
	return int(v.pri)
}

func (v *MplsEntropyIdentifier) SetPri(pri int) {
	v.pri = C.int(pri)
}

func (v *MplsEntropyIdentifier) Flags() uint32 {
	return uint32(v.flags)
}

func (v *MplsEntropyIdentifier) SetFlags(flags uint32) {
	v.flags = C.uint32(flags)
}

func (v *MplsEntropyIdentifier) Init() {
	C.opennsl_mpls_entropy_identifier_t_init(v.C())
}
