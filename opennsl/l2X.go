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
#include <opennsl/l2X.h>
*/
import "C"

//
// L2LearnLimit
//
type L2LearnLimit C.opennsl_l2_learn_limit_t

func NewL2LearnLimit() *L2LearnLimit {
	limit := &L2LearnLimit{}
	limit.Init()
	return limit
}

func (v *L2LearnLimit) C() *C.opennsl_l2_learn_limit_t {
	return (*C.opennsl_l2_learn_limit_t)(v)
}

func (v *L2LearnLimit) SetFlags(flags L2LearnLimitFlags) {
	v.flags = flags.C()
}

func (v *L2LearnLimit) Flags() L2LearnLimitFlags {
	return L2LearnLimitFlags(v.flags)
}

func (v *L2LearnLimit) SetLimit(limit int) {
	v.limit = C.int(limit)
}

func (v *L2LearnLimit) Limit() int {
	return int(v.limit)
}

func (v *L2LearnLimit) Init() {
	C.opennsl_l2_learn_limit_t_init(v.C())
}

func (v *L2LearnLimit) Set(unit int) error {
	rc := C.opennsl_l2_learn_limit_set(C.int(unit), v.C())
	return ParseError(rc)
}

func (v *L2LearnLimit) Get(unit int) error {
	rc := C.opennsl_l2_learn_limit_get(C.int(unit), v.C())
	return ParseError(rc)
}
