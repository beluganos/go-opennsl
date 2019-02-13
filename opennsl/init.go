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
#include <opennsl/init.h>
#include <opennsl/initX.h>
*/
import "C"

const (
	INFO_DEVICE_NOP       uint32 = 0
	INFO_DEVICE_QUMRAN_MX uint32 = 0x8375
	INFO_DEVICE_DNX       uint32 = 0x8375
)

type Info C.opennsl_info_t

func (v *Info) C() *C.opennsl_info_t {
	return (*C.opennsl_info_t)(v)
}

func (v *Info) Device() uint32 {
	return uint32(v.device)
}

func (v *Info) Revision() uint32 {
	return uint32(v.revision)
}

//
// API
//
func InfoInit(info *Info) {
	C.opennsl_info_t_init(info.C())
}

func (v *Info) Init() {
	InfoInit(v)
}

func NewInfo() *Info {
	info := &Info{}
	info.Init()
	return info
}

func InfoGet(unit int) (*Info, error) {
	info := NewInfo()
	rc := C.opennsl_info_get(C.int(unit), info.C())
	if err := ParseError(rc); err != nil {
		return nil, err
	}

	return info, nil
}

//
// Extention
//
var infoDeviceCache uint32 = 0

func InfoDeviceCacheInit(unit int) {
	if info, err := InfoGet(unit); err != nil {
		InfoDeviceCacheSet(unit, info.Device())
	}
}

func InfoDeviceCacheGet(unit int) uint32 {
	return infoDeviceCache
}

func InfoDeviceCacheSet(unit int, dev uint32) {
	infoDeviceCache = dev
}

func InfoDeviceIsDNX(unit int) bool {
	return infoDeviceCache == INFO_DEVICE_DNX
}
