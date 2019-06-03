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
#include <opennsl/l2.h>
#include <helper.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

type L2AddrCallback func(int, *L2Addr, L2CallbackOper)

var l2addrCallbacks = map[int]L2AddrCallback{}

func registerL2AddrCallback(unit int, callback L2AddrCallback) error {
	if _, ok := l2addrCallbacks[unit]; ok {
		return fmt.Errorf("unit %d already exists.", unit)
	}

	l2addrCallbacks[unit] = callback
	return nil
}

func unregisterL2AddrCallback(unit int) {
	delete(l2addrCallbacks, unit)
}

//export go_opennsl_l2_addr_cb
func go_opennsl_l2_addr_cb(c_unit C.int, c_l2addr *C.opennsl_l2_addr_t, c_operation C.int, c_userdata unsafe.Pointer) {
	unit := int(c_unit)
	if callback, ok := l2addrCallbacks[unit]; ok {
		callback(unit, (*L2Addr)(c_l2addr), L2CallbackOper(c_operation))
	}
}

func L2AddrRegister(unit int, callback L2AddrCallback) error {
	if err := registerL2AddrCallback(unit, callback); err != nil {
		return err
	}

	rc := C.opennsl_l2_addr_register(C.int(unit), C.opennsl_l2_addr_callback_t(C._opennsl_l2_addr_cb), nil)

	if err := ParseError(rc); err != nil {
		unregisterL2AddrCallback(unit)
		return err
	}

	return nil
}

func L2AddrUnregister(unit int) error {
	unregisterL2AddrCallback(unit)

	rc := C.opennsl_l2_addr_unregister(C.int(unit), C.opennsl_l2_addr_callback_t(C._opennsl_l2_addr_cb), nil)

	return ParseError(rc)
}
