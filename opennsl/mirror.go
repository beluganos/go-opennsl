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
#include <opennsl/mirror.h>
*/
import "C"

//
// MirrorEnable
//
type MirrorEnable C.int

func (v MirrorEnable) C() C.int {
	return C.int(v)
}

const (
	MIRROR_DISABLE MirrorEnable = C.OPENNSL_MIRROR_DISABLE
	MIRROR_L2      MirrorEnable = C.OPENNSL_MIRROR_L2
)

//
// MirrorPortFlags
//
type MirrorPortFlags C.uint32

func (v MirrorPortFlags) C() C.uint32 {
	return C.uint32(v)
}

func NewMirrorPortFlags(flags ...MirrorPortFlags) MirrorPortFlags {
	v := MIRROR_PORT_NONE
	for _, flag := range flags {
		v |= flag
	}
	return v
}

const (
	MIRROR_PORT_NONE        MirrorPortFlags = 0
	MIRROR_PORT_ENABLE      MirrorPortFlags = C.OPENNSL_MIRROR_PORT_ENABLE
	MIRROR_PORT_INGRESS     MirrorPortFlags = C.OPENNSL_MIRROR_PORT_INGRESS
	MIRROR_PORT_EGRESS      MirrorPortFlags = C.OPENNSL_MIRROR_PORT_EGRESS
	MIRROR_PORT_EGRESS_TRUE MirrorPortFlags = C.OPENNSL_MIRROR_PORT_EGRESS_TRUE
	MIRROR_PORT_DEST_TRUNK  MirrorPortFlags = C.OPENNSL_MIRROR_PORT_DEST_TRUNK
	MIRROR_PORT_EGRESS_ACL  MirrorPortFlags = C.OPENNSL_MIRROR_PORT_EGRESS_ACL
	MIRROR_PORT_SFLOW       MirrorPortFlags = C.OPENNSL_MIRROR_PORT_SFLOW
)

//
// API
//

func MirrorInit(unit int) error {
	rc := C.opennsl_mirror_init(C.int(unit))
	return ParseError(rc)
}

func MirrorToSet(unit int, port Port) error {
	rc := C.opennsl_mirror_to_set(C.int(unit), port.C())
	return ParseError(rc)
}

func MirrorToGet(unit int) (Port, error) {
	c_port := C.opennsl_port_t(0)

	rc := C.opennsl_mirror_to_get(C.int(unit), &c_port)
	return Port(c_port), ParseError(rc)
}

func MirrorIngressSet(unit int, port Port, enable MirrorEnable) error {
	rc := C.opennsl_mirror_ingress_set(C.int(unit), port.C(), enable.C())
	return ParseError(rc)
}

func MirrorIngressGet(unit int, port Port) (MirrorEnable, error) {
	c_enable := C.int(FALSE)

	rc := C.opennsl_mirror_ingress_get(C.int(unit), port.C(), &c_enable)
	return MirrorEnable(c_enable), ParseError(rc)
}

func MirrorEgressSet(unit int, port Port, enable MirrorEnable) error {
	rc := C.opennsl_mirror_egress_set(C.int(unit), port.C(), enable.C())
	return ParseError(rc)
}

func MirrorEgressGet(unit int, port Port) (MirrorEnable, error) {
	c_enable := MIRROR_DISABLE.C()

	rc := C.opennsl_mirror_egress_get(C.int(unit), port.C(), &c_enable)
	return MirrorEnable(c_enable), ParseError(rc)
}

func MirrorPortSet(unit int, port Port, module Module, dstPort Port, flags MirrorPortFlags) error {
	rc := C.opennsl_mirror_port_set(C.int(unit), port.C(), module.C(), dstPort.C(), flags.C())
	return ParseError(rc)
}

func MirrorPortGet(unit int, port Port) (Module, Port, MirrorPortFlags, error) {
	c_module := C.opennsl_module_t(0)
	c_port := C.opennsl_port_t(0)
	c_flags := C.uint32(0)

	rc := C.opennsl_mirror_port_get(C.int(unit), port.C(), &c_module, &c_port, &c_flags)
	return Module(c_module), Port(c_port), MirrorPortFlags(c_flags), ParseError(rc)
}
