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
// TrunkMemberFlags
//
type TrunkMemberFlags C.uint32

func (v TrunkMemberFlags) C() C.uint32 {
	return C.uint32(v)
}

func NewTrunkMemberFlags(flags ...TrunkMemberFlags) TrunkMemberFlags {
	v := TRUNK_MEMBER_NONE
	for _, flag := range flags {
		v |= flag
	}
	return v
}

const (
	TRUNK_MEMBER_NONE           TrunkMemberFlags = 0
	TRUNK_MEMBER_EGRESS_DISABLE TrunkMemberFlags = C.OPENNSL_TRUNK_MEMBER_EGRESS_DISABLE
)

//
// TrunkMember
//
type TrunkMember C.opennsl_trunk_member_t

func (v *TrunkMember) C() *C.opennsl_trunk_member_t {
	return (*C.opennsl_trunk_member_t)(v)
}

func (v *TrunkMember) Flags() TrunkMemberFlags {
	return TrunkMemberFlags(v.flags)
}

func (v *TrunkMember) SetFlags(flags TrunkMemberFlags) {
	v.flags = flags.C()
}

func (v *TrunkMember) GPort() GPort {
	return GPort(v.gport)
}

func (v *TrunkMember) SetGPort(gport GPort) {
	v.gport = gport.C()
}

//
// API
//
func (v *TrunkMember) Init() {
	C.opennsl_trunk_member_t_init(v.C())
}

func NewTrunkMember() *TrunkMember {
	m := &TrunkMember{}
	m.Init()
	return m
}
