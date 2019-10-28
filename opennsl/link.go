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
#include <opennsl/link.h>
#include "helper.h"
*/
import "C"

import (
	"fmt"
	"time"
)

//
// LinkStatus
//
type LinkStatus C.int

func (v LinkStatus) C() C.int {
	return C.int(v)
}

const (
	LINK_STATUS_UP LinkStatus = C.OPENNSL_PORT_LINK_STATUS_UP
)

func (v LinkStatus) IsUp() bool {
	return (v == LINK_STATUS_UP)
}

func (v LinkStatus) String() string {
	if v == LINK_STATUS_UP {
		return "UP"
	} else {
		return "DOWN"
	}
}

//
// LinkscanMode
//
type LinkscanMode C.opennsl_linkscan_mode_t

func (v LinkscanMode) C() C.opennsl_linkscan_mode_t {
	return C.opennsl_linkscan_mode_t(v)
}

const (
	LINKSCAN_MODE_NONE  LinkscanMode = C.OPENNSL_LINKSCAN_MODE_NONE
	LINKSCAN_MODE_SW    LinkscanMode = C.OPENNSL_LINKSCAN_MODE_SW
	LINKSCAN_MODE_HW    LinkscanMode = C.OPENNSL_LINKSCAN_MODE_HW
	LINKSCAN_MODE_COUNT LinkscanMode = C.OPENNSL_LINKSCAN_MODE_COUNT
)

var linkscanMode_names = map[LinkscanMode]string{
	LINKSCAN_MODE_NONE:  "NONE",
	LINKSCAN_MODE_SW:    "SW",
	LINKSCAN_MODE_HW:    "HW",
	LINKSCAN_MODE_COUNT: "COUNT",
}

var linkscanMode_values = map[string]LinkscanMode{
	"SW":    LINKSCAN_MODE_SW,
	"HW":    LINKSCAN_MODE_HW,
	"COUNT": LINKSCAN_MODE_COUNT,
}

func (v LinkscanMode) String() string {
	if s, ok := linkscanMode_names[v]; ok {
		return s
	}
	return fmt.Sprintf("LinkscanMode(%d)", v)
}

func ParseLinkscanMode(s string) (LinkscanMode, error) {
	if v, ok := linkscanMode_values[s]; ok {
		return v, nil
	}
	return LINKSCAN_MODE_NONE, fmt.Errorf("Invalid LinkscanMode. %s", s)
}

//
// API
//
func LinkscanDetach(unit int) error {
	rc := C.opennsl_linkscan_detach(C.int(unit))
	return ParseError(rc)
}

func LinkscanEnableSet(unit int, duration time.Duration) error {
	us := duration.Nanoseconds() / 1000
	rc := C.opennsl_linkscan_enable_set(C.int(unit), C.int(us))
	return ParseError(rc)
}

func LinkscanEnableGet(unit int) (time.Duration, error) {
	c_us := C.int(0)

	rc := C.opennsl_linkscan_enable_get(C.int(unit), &c_us)
	return time.Duration(int64(c_us) * 1000), ParseError(rc)
}

func LinkscanModeSet(unit int, port Port, mode LinkscanMode) error {
	rc := C.opennsl_linkscan_mode_set(C.int(unit), port.C(), C.int(mode.C()))
	return ParseError(rc)
}

func LinkscanModeSetPbm(unit int, pbmp *PBmp, mode LinkscanMode) error {
	rc := C.opennsl_linkscan_mode_set_pbm(C.int(unit), *pbmp.C(), C.int(mode.C()))
	return ParseError(rc)
}

func LinkscanModeGet(unit int, port Port) (LinkscanMode, error) {
	c_mode := C.int(0)
	rc := C.opennsl_linkscan_mode_get(C.int(unit), port.C(), &c_mode)
	return LinkscanMode(c_mode), ParseError(rc)
}

type LinkscanCallback func(int, string, Port, *PortInfo)

var linkscanCallbacks = map[string]LinkscanCallback{}

func linkscanCallbacksRegister(unit int, key string, callback LinkscanCallback) error {
	if len(linkscanCallbacks) == 0 {
		rc := C.opennsl_linkscan_register(C.int(unit), C.opennsl_linkscan_handler_t(C._opennsl_linkscan_handler))
		if err := ParseError(rc); err != nil {
			return err
		}
	}

	linkscanCallbacks[key] = callback
	return nil
}

func linkscanCallbacksUnregister(unit int, key string) error {
	if _, ok := linkscanCallbacks[key]; !ok {
		return nil
	}

	delete(linkscanCallbacks, key)

	if len(linkscanCallbacks) == 0 {
		rc := C.opennsl_linkscan_unregister(C.int(unit), C.opennsl_linkscan_handler_t(C._opennsl_linkscan_handler))
		return ParseError(rc)
	}

	return nil
}

//export go_opennsl_linkscan_handler
func go_opennsl_linkscan_handler(unit C.int, port C.opennsl_port_t, info *C.opennsl_port_info_t) {
	for key, callback := range linkscanCallbacks {
		callback(int(unit), key, Port(port), (*PortInfo)(info))
	}
}

func LinkscanRegister(unit int, key string, callback LinkscanCallback) error {
	return linkscanCallbacksRegister(unit, key, callback)
}

func LinkscanUnregister(unit int, key string) error {
	return linkscanCallbacksUnregister(unit, key)
}

//
// Port method
//
func (v Port) LinkscanMode(unit int) (LinkscanMode, error) {
	return LinkscanModeGet(unit, v)
}

func (v Port) LinkscanModeSet(unit int, mode LinkscanMode) error {
	return LinkscanModeSet(unit, v, mode)
}
