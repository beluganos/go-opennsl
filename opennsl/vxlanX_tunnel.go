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
#include <stdlib.h>
#include <opennsl/vxlanX.h>
#include "helper.h"
*/
import "C"

import (
	"unsafe"
)

//
// VxlanTunnelInitiator
//
type VxlanTunnelInitiator struct {
	*TunnelInitiator
}

func NewVxlanTunnelInitiator() *VxlanTunnelInitiator {
	return &VxlanTunnelInitiator{
		TunnelInitiator: NewTunnelInitiator(TunnelTypeVxlan),
	}
}

func VxlanTunnelInitiatorCreate(unit int, info *TunnelInitiator) error {
	rc := C.opennsl_vxlan_tunnel_initiator_create(C.int(unit), info.C())
	return ParseError(rc)
}

func (v *VxlanTunnelInitiator) Create(unit int) error {
	return VxlanTunnelInitiatorCreate(unit, v.TunnelInitiator)
}

func VxlanTunnelInitiatorDestroy(unit int, tunID TunnelID) error {
	rc := C.opennsl_vxlan_tunnel_initiator_destroy(C.int(unit), tunID.C())
	return ParseError(rc)
}

func (v *VxlanTunnelInitiator) Destroy(unit int) error {
	return VxlanTunnelInitiatorDestroy(unit, v.TunnelID())
}

func VxlanTunnelInitiatorGet(unit int, tunID TunnelID) (*VxlanTunnelInitiator, error) {
	info := NewVxlanTunnelInitiator()
	info.SetTunnelID(tunID)

	rc := C.opennsl_vxlan_tunnel_initiator_get(C.int(unit), info.C())
	if err := ParseError(rc); err != nil {
		return nil, err
	}

	return info, nil
}

func VxlanTunnelInitiatorTraverse(unit int, callback TunnelInitiatorTraverseCallback) error {
	n := tunnelTerminatorTraverseCallbacks.Add(callback)
	defer tunnelTerminatorTraverseCallbacks.Del(n)

	rc := C.opennsl_vxlan_tunnel_initiator_traverse(C.int(unit), C.opennsl_tunnel_initiator_traverse_cb(C._opennsl_tunnel_initiator_traverse_cb), unsafe.Pointer(&n))
	return ParseError(rc)
}

//
// VxlanTunnelTerminator
//
type VxlanTunnelTerminator struct {
	*TunnelTerminator
}

func NewVxlanTunnelTerminator() *VxlanTunnelTerminator {
	return &VxlanTunnelTerminator{
		TunnelTerminator: NewTunnelTerminator(TunnelTypeVxlan),
	}
}

func VxlanTunnelTerminatorCreate(unit int, info *TunnelTerminator) error {
	rc := C.opennsl_vxlan_tunnel_terminator_create(C.int(unit), info.C())
	return ParseError(rc)
}

func (v *VxlanTunnelTerminator) Create(unit int) error {
	return VxlanTunnelTerminatorCreate(unit, v.TunnelTerminator)
}

func VxlanTunnelTerminatorUpdate(unit int, info *TunnelTerminator) error {
	rc := C.opennsl_vxlan_tunnel_terminator_update(C.int(unit), info.C())
	return ParseError(rc)
}

func (v *VxlanTunnelTerminator) Update(unit int) error {
	return VxlanTunnelTerminatorUpdate(unit, v.TunnelTerminator)
}

func VxlanTunnelTerminatorDestroy(unit int, tunID TunnelID) error {
	rc := C.opennsl_vxlan_tunnel_terminator_destroy(C.int(unit), tunID.C())
	return ParseError(rc)
}

func (v *VxlanTunnelTerminator) Destroy(unit int) error {
	return VxlanTunnelTerminatorDestroy(unit, v.TunnelID())
}

func VxlanTunnelTerminatorGet(unit int, tunID TunnelID) (*VxlanTunnelTerminator, error) {
	info := NewVxlanTunnelTerminator()
	info.SetTunnelID(tunID)

	rc := C.opennsl_vxlan_tunnel_terminator_get(C.int(unit), info.C())
	if err := ParseError(rc); err != nil {
		return nil, err
	}

	return info, nil
}

func VxlanTunnelTerminatorTraverse(unit int, callback TunnelTerminatorTraverseCallback) error {
	n := tunnelTerminatorTraverseCallbacks.Add(callback)
	defer tunnelTerminatorTraverseCallbacks.Del(n)

	rc := C.opennsl_vxlan_tunnel_terminator_traverse(C.int(unit), C.opennsl_tunnel_terminator_traverse_cb(C._opennsl_tunnel_terminator_traverse_cb), unsafe.Pointer(&n))
	return ParseError(rc)
}
