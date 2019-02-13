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
#include <opennsl/trunk.h>
*/
import "C"

//
// TrunkInfo
//
type TrunkInfo C.opennsl_trunk_info_t

func (v *TrunkInfo) C() *C.opennsl_trunk_info_t {
	return (*C.opennsl_trunk_info_t)(v)
}

func (v *TrunkInfo) Psc() int {
	return int(v.psc)
}

func (v *TrunkInfo) SetPsc(psc int) {
	v.psc = C.int(psc)
}

func (v *TrunkInfo) DLFIndex() int {
	return int(v.dlf_index)
}

func (v *TrunkInfo) SetDLFIndex(dlfIndex int) {
	v.dlf_index = C.int(dlfIndex)
}

func (v *TrunkInfo) MCIndex() int {
	return int(v.mc_index)
}

func (v *TrunkInfo) SetMCIndex(index int) {
	v.mc_index = C.int(index)
}

func (v *TrunkInfo) IPMCIndex() int {
	return int(v.ipmc_index)
}

func (v *TrunkInfo) SetIPMCIndex(index int) {
	v.ipmc_index = C.int(index)
}

//
// API
//
func (v *TrunkInfo) Init() {
	C.opennsl_trunk_info_t_init(v.C())
}

func NewTrunkInfo() *TrunkInfo {
	trunk := &TrunkInfo{}
	trunk.Init()
	return trunk
}
