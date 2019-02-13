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
#include <opennsl/l3.h>
*/
import "C"

//
// L3Info
//
type L3Info C.opennsl_l3_info_t

func (v *L3Info) C() *C.opennsl_l3_info_t {
	return (*C.opennsl_l3_info_t)(v)
}

func (v *L3Info) MaxIface() int {
	return int(v.l3info_max_intf)
}

func (v *L3Info) SetMaxIface(maxIface int) {
	v.l3info_max_intf = C.int(maxIface)
}

func (v *L3Info) MaxHost() int {
	return int(v.l3info_max_host)
}

func (v *L3Info) SetMaxHost(maxHost int) {
	v.l3info_max_host = C.int(maxHost)
}

func (v *L3Info) MaxRoute() int {
	return int(v.l3info_max_route)
}

func (v *L3Info) SetMaxRoute(maxRoute int) {
	v.l3info_max_route = C.int(maxRoute)
}

func (v *L3Info) UsedIface() int {
	return int(v.l3info_used_intf)
}

func (v *L3Info) SetUsedIface(usedIface int) {
	v.l3info_used_intf = C.int(usedIface)
}

//
// API
//
func NewL3Info() *L3Info {
	info := &L3Info{}
	info.Init()
	return info
}

func L3InfoInit(v *L3Info) {
	C.opennsl_l3_info_t_init(v.C())
}

func (v *L3Info) Init() {
	L3InfoInit(v)
}

func L3InfoGet(unit int) (*L3Info, error) {
	l3info := NewL3Info()
	rc := C.opennsl_l3_info(C.int(unit), l3info.C())
	if err := ParseError(rc); err != nil {
		return nil, err
	}

	return l3info, nil
}
