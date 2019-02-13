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

package sal

/*
#include <stdlib.h>
#include <sal/driver.h>
*/
import "C"

import (
	"fmt"
	"strings"
	"unsafe"

	"github.com/beluganos/go-opennsl/opennsl"
)

type InitFlags C.uint

func (v InitFlags) C() C.uint {
	return C.uint(v)
}

func NewInitFlags(flags ...InitFlags) InitFlags {
	v := INIT_NONE
	for _, flg := range flags {
		v |= flg
	}
	return v
}

const (
	INIT_NONE      InitFlags = 0
	INIT_FAST_BOOT InitFlags = C.OPENNSL_F_FAST_BOOT
)

var initFlags_names = map[InitFlags]string{
	INIT_NONE:      "NONE",
	INIT_FAST_BOOT: "FAST_BOOT",
}

var initFlags_values = map[string]InitFlags{
	"FAST_BOOT": INIT_FAST_BOOT,
}

func (v InitFlags) String() string {
	names := []string{}
	for flg, name := range initFlags_names {
		if v&flg != 0 {
			names = append(names, name)
		}
	}
	return strings.Join(names, "|")
}

func ParseInitFlags(s string) (InitFlags, error) {
	if v, ok := initFlags_values[s]; ok {
		return v, nil
	}
	return INIT_NONE, fmt.Errorf("Invalid InitFlags. %s", s)
}

type Init C.opennsl_init_t

func (v *Init) C() *C.opennsl_init_t {
	return (*C.opennsl_init_t)(v)
}

func (v *Init) Free() {
	v.FreeCfgFname()
	v.FreeWBFname()
	v.FreeRMCfgFname()
	v.FreeCfgPostFname()
}

func (v *Init) Flags() InitFlags {
	return InitFlags(v.flags)
}

func (v *Init) SetFlags(flags InitFlags) {
	v.flags = flags.C()
}

func (v *Init) OpenNSLFlags() uint32 {
	return uint32(v.opennsl_flags)
}

func (v *Init) SetOpenNSLFlags(flags uint32) {
	v.opennsl_flags = C.uint(flags)
}

func (v *Init) CfgFname() string {
	if v == nil || v.cfg_fname == nil {
		return ""
	}
	return C.GoString(v.cfg_fname)
}

func (v *Init) SetCfgFname(cfgFname string) {
	v.FreeCfgFname()
	v.cfg_fname = C.CString(cfgFname)
}

func (v *Init) FreeCfgFname() {
	if v.cfg_fname != nil {
		C.free(unsafe.Pointer(v.cfg_fname))
		v.cfg_fname = nil
	}
}

func (v *Init) WBFname() string {
	if v == nil || v.wb_fname == nil {
		return ""
	}
	return C.GoString(v.wb_fname)
}

func (v *Init) SetWBFname(wbFname string) {
	v.FreeWBFname()
	v.wb_fname = C.CString(wbFname)
}

func (v *Init) FreeWBFname() {
	if v.wb_fname != nil {
		C.free(unsafe.Pointer(v.wb_fname))
		v.wb_fname = nil
	}
}

func (v *Init) RMCfgFname() string {
	if v == nil || v.rmcfg_fname == nil {
		return ""
	}
	return C.GoString(v.rmcfg_fname)
}

func (v *Init) SetRMCfgFname(rmcfgFname string) {
	v.FreeRMCfgFname()
	v.rmcfg_fname = C.CString(rmcfgFname)
}

func (v *Init) FreeRMCfgFname() {
	if v.rmcfg_fname != nil {
		C.free(unsafe.Pointer(v.rmcfg_fname))
		v.rmcfg_fname = nil
	}
}

func (v *Init) CfgPostFname() string {
	if v == nil || v.cfg_post_fname == nil {
		return ""
	}
	return C.GoString(v.cfg_post_fname)
}

func (v *Init) SetCfgPostFname(cfgFname string) {
	v.FreeCfgPostFname()
	v.cfg_post_fname = C.CString(cfgFname)
}

func (v *Init) FreeCfgPostFname() {
	if v.cfg_post_fname != nil {
		C.free(unsafe.Pointer(v.cfg_post_fname))
		v.cfg_post_fname = nil
	}
}

func DriverInit() error {
	rc := C.opennsl_driver_init(nil)
	return opennsl.OpenNSLError(rc).Error()
}

func (v *Init) Init() error {
	if v == nil {
		return DriverInit()
	}
	rc := C.opennsl_driver_init(v.C())
	return opennsl.OpenNSLError(rc).Error()
}

func DriverExit() error {
	rc := C.opennsl_driver_exit()
	return opennsl.OpenNSLError(rc).Error()
}

func DriverBootFlagsGet() uint32 {
	return uint32(C.opennsl_driver_boot_flags_get())
}

/*
func DriverShell() error {
	rc := C.opennsl_driver_shell()
	return ParseOpenNSLError(rc).Error()
}

func DriverProcessCommand(cmd string) error {
	commandBuf := C.CString(cmd)
	defer C.free(unsafe.Pointer(commandBuf))

	rc := C.opennsl_driver_process_command(commandBuf)
	return ParseOpenNSLError(rc).Error()
}

func DeriverPlatformPhyCleanup() {
	C.platform_phy_cleanup()
}
*/
