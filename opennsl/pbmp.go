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
#include <shared/pbmp.h>
#include "helper.h"
*/
import "C"

import (
	"fmt"
)

const (
	OPENNSL_VLAN_NONE    uint16 = C.OPENNSL_VLAN_NONE
	OPENNSL_VLAN_DEFAULT uint16 = C.OPENNSL_VLAN_DEFAULT
)

const (
	SHR_PBMP_WORD_MAX     = C._SHR_PBMP_WORD_MAX
	OPENNSL_PBMP_PORT_MAX = C.OPENNSL_PBMP_PORT_MAX
)

type PBmp C.opennsl_pbmp_t

func (v *PBmp) C() *C.opennsl_pbmp_t {
	return (*C.opennsl_pbmp_t)(v)
}

func (v *PBmp) String() string {
	ports := make([]uint32, 0)
	for index, bits := range v.PBits() {
		for bit := uint32(0); bit < 32; bit++ {
			if bits&(1<<bit) != 0 {
				ports = append(ports, uint32(index)*32+bit)
			}
		}
	}
	return fmt.Sprintf("%v", ports)
}

//
// API
//
func NewPBmp() *PBmp {
	pbmp := &PBmp{}
	pbmp.Init()
	return pbmp
}

func (v *PBmp) Init() *PBmp {
	C._opennsl_pbmp_clear(v.C())
	return v
}

func (v *PBmp) Has(port Port) bool {
	return C._opennsl_pbmp_member(v.C(), port.C()) != 0
}

func (v *PBmp) Count() int {
	cnt := C._opennsl_pbmp_count(v.C())
	return int(cnt)
}

func (v *PBmp) IsNull() bool {
	rv := C._opennsl_pbmp_is_null(v.C())
	return (rv != 0)
}

func (v *PBmp) IsNotNull() bool {
	rv := C._opennsl_pbmp_is_not_null(v.C())
	return (rv != 0)
}

func (v *PBmp) Equal(pbmp *PBmp) bool {
	rv := C._opennsl_pbmp_eq(v.C(), pbmp.C())
	return (rv != 0)
}

func (v *PBmp) NotEqual(pbmp *PBmp) bool {
	rv := C._opennsl_pbmp_neq(v.C(), pbmp.C())
	return (rv != 0)
}

func (v *PBmp) Assign(src *PBmp) *PBmp {
	C._opennsl_pbmp_assign(v.C(), src.C())
	return v
}

func (v *PBmp) AND(pbmp *PBmp) *PBmp {
	C._opennsl_pbmp_and(v.C(), pbmp.C())
	return v
}

func (v *PBmp) OR(pbmp *PBmp) *PBmp {
	C._opennsl_pbmp_or(v.C(), pbmp.C())
	return v
}

func (v *PBmp) XOR(pbmp *PBmp) *PBmp {
	C._opennsl_pbmp_xor(v.C(), pbmp.C())
	return v
}

func (v *PBmp) RemovePorts(pbmp *PBmp) *PBmp {
	C._opennsl_pbmp_remove(v.C(), pbmp.C())
	return v
}

func (v *PBmp) Negate(pbmp *PBmp) *PBmp {
	C._opennsl_pbmp_negate(v.C(), pbmp.C())
	return v
}

func (v *PBmp) Set(port Port) {
	C._opennsl_pbmp_port_set(v.C(), port.C())
}

func (v *PBmp) Add(ports ...Port) *PBmp {
	for _, port := range ports {
		C._opennsl_pbmp_port_add(v.C(), port.C())
	}
	return v
}

func (v *PBmp) Remove(ports ...Port) *PBmp {
	for _, port := range ports {
		C._opennsl_pbmp_port_remove(v.C(), port.C())
	}
	return v
}

func (v *PBmp) Clear() *PBmp {
	C._opennsl_pbmp_clear(v.C())
	return v
}

func (v *PBmp) Each(f func(Port) error) error {
	for index := 0; index < OPENNSL_PBMP_PORT_MAX; index++ {
		port := Port(index)
		if v.Has(port) {
			if err := f(port); err != nil {
				return err
			}
		}
	}
	return nil
}

func (v *PBmp) PortList() []Port {
	ports := []Port{}
	v.Each(func(port Port) error {
		ports = append(ports, port)
		return nil
	})
	return ports
}

func (v *PBmp) Reverse(f func(Port) error) error {
	for index := (OPENNSL_PBMP_PORT_MAX - 1); index >= 0; index-- {
		port := Port(index)
		if v.Has(port) {
			if err := f(port); err != nil {
				return err
			}
		}
	}
	return nil
}

func (v *PBmp) PBits() []uint32 {
	bits := make([]uint32, len(v.pbits))
	for i, bit := range v.pbits {
		bits[i] = uint32(bit)
	}
	return bits
}

func (v *PBmp) SetPBits(bits []uint32) {
	l := len(bits)
	if l > SHR_PBMP_WORD_MAX {
		l = SHR_PBMP_WORD_MAX
	}
	for i := 0; i < l; i++ {
		v.pbits[i] = C.uint32(bits[i])
	}
}
