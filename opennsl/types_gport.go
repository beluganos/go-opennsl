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
#include "helper.h"
*/
import "C"

//
// GPortType
//
type GPortType C.int

func (v GPortType) C() C.int {
	return C.int(v)
}

const (
	GPORT_TYPE_NONE                         GPortType = C.OPENNSL_GPORT_TYPE_NONE
	GPORT_TYPE_LOCAL                        GPortType = C.OPENNSL_GPORT_TYPE_LOCAL
	GPORT_TYPE_MODPORT                      GPortType = C.OPENNSL_GPORT_TYPE_MODPORT
	GPORT_TYPE_UCAST_QUEUE_GROUP            GPortType = C.OPENNSL_GPORT_TYPE_UCAST_QUEUE_GROUP
	GPORT_TYPE_DESTMOD_QUEUE_GROUP          GPortType = C.OPENNSL_GPORT_TYPE_DESTMOD_QUEUE_GROUP
	GPORT_TYPE_MCAST                        GPortType = C.OPENNSL_GPORT_TYPE_MCAST
	GPORT_TYPE_MCAST_QUEUE_GROUP            GPortType = C.OPENNSL_GPORT_TYPE_MCAST_QUEUE_GROUP
	GPORT_TYPE_SCHEDULER                    GPortType = C.OPENNSL_GPORT_TYPE_SCHEDULER
	GPORT_TYPE_CHILD                        GPortType = C.OPENNSL_GPORT_TYPE_CHILD
	GPORT_TYPE_EGRESS_GROUP                 GPortType = C.OPENNSL_GPORT_TYPE_EGRESS_GROUP
	GPORT_TYPE_EGRESS_CHILD                 GPortType = C.OPENNSL_GPORT_TYPE_EGRESS_CHILD
	GPORT_TYPE_EGRESS_MODPORT               GPortType = C.OPENNSL_GPORT_TYPE_EGRESS_MODPORT
	GPORT_TYPE_UCAST_SUBSCRIBER_QUEUE_GROUP GPortType = C.OPENNSL_GPORT_TYPE_UCAST_SUBSCRIBER_QUEUE_GROUP
	GPORT_TYPE_MCAST_SUBSCRIBER_QUEUE_GROUP GPortType = C.OPENNSL_GPORT_TYPE_MCAST_SUBSCRIBER_QUEUE_GROUP
	GPORT_TYPE_COSQ                         GPortType = C.OPENNSL_GPORT_TYPE_COSQ
	GPORT_TYPE_PROFILE                      GPortType = C.OPENNSL_GPORT_TYPE_PROFILE
)

//
// GPort
//
type GPort C.opennsl_gport_t

const (
	GPORT_NONE    GPort = C.OPENNSL_GPORT_TYPE_NONE
	GPORT_INVALID GPort = C.OPENNSL_GPORT_INVALID
)

func (v GPort) C() C.opennsl_gport_t {
	return C.opennsl_gport_t(v)
}

func (v GPort) ToLocal() Port {
	return Port(C._opennsl_gport_to_local(v.C()))
}

func GPortFromLocal(port Port) GPort {
	return GPort(C._opennsl_gport_from_local(port.C()))
}

func (v GPort) ToModule() (Port, Module) {
	c_gport := v.C()
	return Port(C._opennsl_gport_modport_to_port(c_gport)), Module(C._opennsl_gport_modport_to_modid(c_gport))
}

func GPortFromModule(port Port, module Module) GPort {
	return GPort(C._opennsl_gport_from_modid_and_port(port.C(), module.C()))
}
