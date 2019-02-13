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
#include <opennsl/types.h>
#include <opennsl/rx.h>
#include "helper.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

const (
	RX_CHANNELS       = C.OPENNSL_RX_CHANNELS
	CMICX_RX_CHANNELS = C.OPENNSL_CMICX_RX_CHANNELS
)

//
// RxCallbackFlags
//
type RxCallbackFlags C.uint32

func (v RxCallbackFlags) C() C.uint32 {
	return C.uint32(v)
}

const (
	RCO_F_NONE    RxCallbackFlags = 0
	RCO_F_ALL_COS RxCallbackFlags = C.OPENNSL_RCO_F_ALL_COS
)

func NewRxCallbackFlags(coss ...uint) RxCallbackFlags {
	flags := RCO_F_NONE
	for _, cos := range coss {
		flags |= (1 << cos)
	}
	return flags
}

func (v RxCallbackFlags) String() string {
	coss := make([]uint, 0, 32)
	for cos := uint(0); cos < 32; cos++ {
		if v&(1<<cos) != 0 {
			coss = append(coss, cos)
		}
	}
	return fmt.Sprintf("%v", coss)
}

//
// RxResult
//
type RxResult C.opennsl_rx_t

func (v RxResult) C() C.opennsl_rx_t {
	return C.opennsl_rx_t(v)
}

const (
	RX_INVALID       RxResult = C.OPENNSL_RX_INVALID
	RX_NOT_HANDLED   RxResult = C.OPENNSL_RX_NOT_HANDLED
	RX_HANDLED       RxResult = C.OPENNSL_RX_HANDLED
	RX_HANDLED_OWNED RxResult = C.OPENNSL_RX_HANDLED_OWNED
)

var rxResult_names = map[RxResult]string{
	RX_INVALID:       "INVALID",
	RX_NOT_HANDLED:   "NOT_HANDLED",
	RX_HANDLED:       "HANDLED",
	RX_HANDLED_OWNED: "HANDLED_OWNED",
}

var rxResult_values = map[string]RxResult{
	"INVALID":       RX_INVALID,
	"NOT_HANDLED":   RX_NOT_HANDLED,
	"HANDLED":       RX_HANDLED,
	"HANDLED_OWNED": RX_HANDLED_OWNED,
}

func (v RxResult) String() string {
	if s, ok := rxResult_names[v]; ok {
		return s
	}
	return fmt.Sprintf("RxResult(%d)", v)
}

func ParseRxResult(s string) (RxResult, error) {
	if v, ok := rxResult_values[s]; ok {
		return v, nil
	}
	return RX_INVALID, fmt.Errorf("Invalid RxResult. %s", s)
}

//
// API
//
func RxActive(unit int) bool {
	rc := C.opennsl_rx_active(C.int(unit))
	return (rc != 0)
}

func RxQueueMaxGet(unit int) (CosQueue, error) {
	c_cos := C.opennsl_cos_queue_t(0)

	rc := C.opennsl_rx_queue_max_get(C.int(unit), &c_cos)
	return CosQueue(c_cos), ParseError(rc)
}

type RxCallback func(int, *Pkt)

var rxCallbackMap = NewRxCallbackMap()

//export go_opennsl_rx_cb
func go_opennsl_rx_cb(c_unit C.int, c_pkt *C.opennsl_pkt_t, data unsafe.Pointer) C.opennsl_rx_t {
	key := (*RxCallbackKey)(data)
	if cb, ok := rxCallbackMap.Get(key); ok {
		cb(int(c_unit), (*Pkt)(c_pkt))
		return RX_HANDLED.C()
	}

	return RX_NOT_HANDLED.C()
}

func RxRegister(unit int, pri uint8, flags RxCallbackFlags, callback RxCallback) error {
	name := fmt.Sprintf("rx_unit_%d_pri_%d", unit, pri)
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	key, err := rxCallbackMap.Register(unit, callback, pri)
	if err != nil {
		return nil
	}

	rc := C.opennsl_rx_register(C.int(unit), c_name, (C.opennsl_rx_cb_f)(C._opennsl_rx_cb), C.uint8(pri), unsafe.Pointer(key), flags.C())
	if err := ParseError(rc); err != nil {
		rxCallbackMap.Unregister(unit, pri)
		return err
	}

	return nil
}

func RxUnregister(unit int, pri uint8) error {
	rxCallbackMap.Unregister(unit, pri)

	rc := C.opennsl_rx_unregister(C.int(unit), (C.opennsl_rx_cb_f)(C._opennsl_rx_cb), C.uint8(pri))
	return ParseError(rc)
}

func RxStart(unit int, cfg *RxCfg) error {
	return cfg.SafeStart(unit)
}
