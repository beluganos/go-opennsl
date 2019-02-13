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
#include <opennsl/cosqX.h>
*/
import "C"

func CosqInit(unit int) error {
	rc := C.opennsl_cosq_init(C.int(unit))
	return ParseError(rc)
}

func CosqDetach(unit int) error {
	rc := C.opennsl_cosq_detach(C.int(unit))
	return ParseError(rc)
}

func CosqConfigSet(unit int, numq int) error {
	rc := C.opennsl_cosq_config_set(C.int(unit), C.int(numq))
	return ParseError(rc)
}

func CosqConfigGet(unit int) (int, error) {
	c_numq := C.int(0)

	rc := C.opennsl_cosq_config_get(C.int(unit), &c_numq)
	return int(c_numq), ParseError(rc)
}

func CosMappingSet(unit int, v Cos, cosq CosQueue) error {
	rc := C.opennsl_cosq_mapping_set(C.int(unit), v.C(), cosq.C())
	return ParseError(rc)
}

func (v Cos) MappingSet(unit int, cosq CosQueue) error {
	return CosMappingSet(unit, v, cosq)
}

func CosMappingGet(unit int, v Cos) (CosQueue, error) {
	c_cosq := C.opennsl_cos_queue_t(0)

	rc := C.opennsl_cosq_mapping_get(C.int(unit), v.C(), &c_cosq)
	return CosQueue(c_cosq), ParseError(rc)
}

func (v Cos) MappingGet(unit int) (CosQueue, error) {
	return CosMappingGet(unit, v)
}

func CosPortMappingSet(unit int, port Port, v Cos, cosq CosQueue) error {
	rc := C.opennsl_cosq_port_mapping_set(C.int(unit), port.C(), v.C(), cosq.C())
	return ParseError(rc)
}

func (v Cos) PortMappingSet(unit int, port Port, cosq CosQueue) error {
	return CosPortMappingSet(unit, port, v, cosq)
}

func CosPortMappingGet(unit int, port Port, v Cos) (CosQueue, error) {
	c_cosq := C.opennsl_cos_queue_t(0)

	rc := C.opennsl_cosq_port_mapping_get(C.int(unit), port.C(), v.C(), &c_cosq)
	return CosQueue(c_cosq), ParseError(rc)
}

func (v Cos) PortMappingGet(unit int, port Port) (CosQueue, error) {
	return CosPortMappingGet(unit, port, v)
}
