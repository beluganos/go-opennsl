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

import (
	"fmt"
	"strings"
)

const TRUNK_MAX_PORTCNT = C.OPENNSL_TRUNK_MAX_PORTCNT
const TRUNK_UNSPEC_INDEX C.int = C.OPENNSL_TRUNK_UNSPEC_INDEX

//
// TrunkFlags
//
type TrunkFlags C.uint32

func (v TrunkFlags) C() C.uint32 {
	return C.uint32(v)
}

func NewTrunkFlags(flags ...TrunkFlags) TrunkFlags {
	v := TRUNK_FLAG_NONE
	for _, flag := range flags {
		v |= flag
	}
	return v
}

const (
	TRUNK_FLAG_NONE                TrunkFlags = 0
	TRUNK_FLAG_FAILOVER_NEXT       TrunkFlags = C.OPENNSL_TRUNK_FLAG_FAILOVER_NEXT
	TRUNK_FLAG_FAILOVER_NEXT_LOCAL TrunkFlags = C.OPENNSL_TRUNK_FLAG_FAILOVER_NEXT_LOCAL
	TRUNK_FLAG_FAILOVER            TrunkFlags = C.OPENNSL_TRUNK_FLAG_FAILOVER
	TRUNK_FLAG_WITH_ID             TrunkFlags = C.OPENNSL_TRUNK_FLAG_WITH_ID
	TRUNK_FLAG_IPMC_CLEAVE         TrunkFlags = C.OPENNSL_TRUNK_FLAG_IPMC_CLEAVE
)

var trunkFlags_names = map[TrunkFlags]string{
	TRUNK_FLAG_FAILOVER_NEXT:       "OPENNSL_TRUNK_FLAG_FAILOVER_NEXT",
	TRUNK_FLAG_FAILOVER_NEXT_LOCAL: "OPENNSL_TRUNK_FLAG_FAILOVER_NEXT_LOCAL",
	TRUNK_FLAG_WITH_ID:             "OPENNSL_TRUNK_FLAG_WITH_ID",
	TRUNK_FLAG_IPMC_CLEAVE:         "OPENNSL_TRUNK_FLAG_IPMC_CLEAVE",
	// TRUNK_FLAG_FAILOVER:            "OPENNSL_TRUNK_FLAG_FAILOVER",
}

func (v TrunkFlags) String() string {
	names := []string{}
	for flag, name := range trunkFlags_names {
		if flag&v != 0 {
			names = append(names, name)
		}
	}
	return strings.Join(names, "|")
}

var trunkFlags_values = map[string]TrunkFlags{
	"OPENNSL_TRUNK_FLAG_FAILOVER_NEXT":       TRUNK_FLAG_FAILOVER_NEXT,
	"OPENNSL_TRUNK_FLAG_FAILOVER_NEXT_LOCAL": TRUNK_FLAG_FAILOVER_NEXT_LOCAL,
	"OPENNSL_TRUNK_FLAG_FAILOVER":            TRUNK_FLAG_FAILOVER,
	"OPENNSL_TRUNK_FLAG_WITH_ID":             TRUNK_FLAG_WITH_ID,
	"OPENNSL_TRUNK_FLAG_IPMC_CLEAVE":         TRUNK_FLAG_IPMC_CLEAVE,
}

func ParseTrunkFlags(names ...string) (TrunkFlags, error) {
	flags := TRUNK_FLAG_NONE
	for _, name := range names {
		flag, ok := trunkFlags_values[name]
		if !ok {
			return TRUNK_FLAG_NONE, fmt.Errorf("Invalid TrunkFlags. %s", name)
		}
		flags |= flag
	}
	return flags, nil
}

//
// TrunkPsc
//
type TrunkPsc C.int

func (v TrunkPsc) C() C.int {
	return C.int(v)
}

//
// Trunk
//
type Trunk C.opennsl_trunk_t

func (v Trunk) C() C.opennsl_trunk_t {
	return C.opennsl_trunk_t(v)
}

//
// API
//
func TrunkInit(unit int) error {
	rc := C.opennsl_trunk_init(C.int(unit))
	return ParseError(rc)
}

func TrunkDetach(unit int) error {
	rc := C.opennsl_trunk_detach(C.int(unit))
	return ParseError(rc)
}

func TrunkCreate(unit int, flags TrunkFlags) (Trunk, error) {
	tid := C.opennsl_trunk_t(0)
	rc := C.opennsl_trunk_create(C.int(unit), flags.C(), &tid)
	return Trunk(tid), ParseError(rc)
}

func TrunkFind(unit int, mod Module, gport GPort) (Trunk, error) {
	tid := C.opennsl_trunk_t(0)
	rc := C.opennsl_trunk_find(C.int(unit), mod.C(), gport.C(), &tid)
	return Trunk(tid), ParseError(rc)
}

func (v Trunk) Destroy(unit int) error {
	rc := C.opennsl_trunk_destroy(C.int(unit), v.C())
	return ParseError(rc)
}

func (v Trunk) PscSet(unit int, psc TrunkPsc) error {
	rc := C.opennsl_trunk_psc_set(C.int(unit), v.C(), psc.C())
	return ParseError(rc)
}

func (v Trunk) FailoverGet(unit int, failPort GPort, arraySize int) (TrunkPsc, TrunkFlags, []GPort, error) {
	c_psc := C.int(0)
	c_flags := C.uint32(0)
	c_ports := make([]C.opennsl_gport_t, arraySize)
	c_count := C.int(0)

	rc := C.opennsl_trunk_failover_get(C.int(unit), v.C(), failPort.C(), &c_psc, &c_flags, C.int(arraySize), &c_ports[0], &c_count)

	if err := ParseError(rc); err != nil {
		return 0, 0, nil, err
	}

	ports := make([]GPort, int(c_count))
	for index := 0; index < int(c_count); index++ {
		ports[index] = GPort(c_ports[index])
	}

	return TrunkPsc(c_psc), TrunkFlags(c_flags), ports, nil
}

func (v Trunk) FailoverSet(unit int, failPort GPort, psc TrunkPsc, flags TrunkFlags, ports []GPort) error {
	count := len(ports)
	c_ports := make([]C.opennsl_gport_t, count)

	for index, port := range ports {
		c_ports[index] = port.C()
	}

	rc := C.opennsl_trunk_failover_set(C.int(unit), v.C(), failPort.C(), psc.C(), flags.C(), C.int(count), &c_ports[0])
	return ParseError(rc)
}

func (v Trunk) MemberAdd(unit int, member *TrunkMember) error {
	rc := C.opennsl_trunk_member_add(C.int(unit), v.C(), member.C())
	return ParseError(rc)
}

func (v Trunk) MemberDelete(unit int, member *TrunkMember) error {
	rc := C.opennsl_trunk_member_delete(C.int(unit), v.C(), member.C())
	return ParseError(rc)
}

func (v Trunk) MemberSet(unit int, info *TrunkInfo, members []TrunkMember) error {
	count := len(members)
	c_members := make([]C.opennsl_trunk_member_t, count)

	for index, member := range members {
		c_members[index] = *member.C()
	}

	rc := C.opennsl_trunk_set(C.int(unit), v.C(), info.C(), C.int(count), &c_members[0])
	return ParseError(rc)
}
