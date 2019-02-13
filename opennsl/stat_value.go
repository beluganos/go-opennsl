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
#include <opennsl/stat.h>
*/
import "C"

//
// StatValue
//
type StatValue C.opennsl_stat_value_t

func (v *StatValue) C() *C.opennsl_stat_value_t {
	return (*C.opennsl_stat_value_t)(v)
}

func (v *StatValue) Packets() uint32 {
	return uint32(v.packets)
}

func (v *StatValue) SetPackets(packets uint32) {
	v.packets = C.uint32(packets)
}

func (v *StatValue) Bytes() uint64 {
	return uint64(v.bytes)
}

func (v *StatValue) SetBytes(bytes uint64) {
	v.bytes = C.uint64(bytes)
}

func (v *StatValue) Packets64() uint64 {
	return uint64(v.packets64)
}

func (v *StatValue) SetPackets64(packets uint64) {
	v.packets64 = C.uint64(packets)
}

//
// StatValue methods
//
func (v *StatValue) Init() {
	C.opennsl_stat_value_t_init(v.C())
}
