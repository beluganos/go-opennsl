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
#include <opennsl/switch.h>
#include "helper.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

type SwitchEventCallback func(int, SwitchEvent, uint32, uint32, uint32)

var switchEventCallbacks = map[int]SwitchEventCallback{} // key:unit

func registerSwitchEventCallback(unit int, callback SwitchEventCallback) error {
	if _, ok := switchEventCallbacks[unit]; ok {
		return fmt.Errorf("unit %d already exists.", unit)
	}
	switchEventCallbacks[unit] = callback
	return nil
}

func unregisterSwitchEventCallback(unit int) {
	delete(switchEventCallbacks, unit)
}

//export go_opennsl_switch_event_cb
func go_opennsl_switch_event_cb(c_unit C.int, c_event C.opennsl_switch_event_t, c_arg1 C.uint32, c_arg2 C.uint32, c_arg3 C.uint32, userdata unsafe.Pointer) {
	unit := int(c_unit)
	if callback, ok := switchEventCallbacks[unit]; ok {
		callback(unit, SwitchEvent(c_event), uint32(c_arg1), uint32(c_arg2), uint32(c_arg3))
	}
}

func SwitchEveltRegister(unit int, callback SwitchEventCallback) error {
	if err := registerSwitchEventCallback(unit, callback); err != nil {
		return err
	}

	rc := C.opennsl_switch_event_register(C.int(unit), C.opennsl_switch_event_cb_t(C._opennsl_switch_event_cb), nil)

	if err := ParseError(rc); err != nil {
		unregisterSwitchEventCallback(unit)
		return err
	}

	return nil
}

func SwitchEveltUnregister(unit int) error {
	unregisterSwitchEventCallback(unit)

	rc := C.opennsl_switch_event_unregister(C.int(unit), C.opennsl_switch_event_cb_t(C._opennsl_switch_event_cb), nil)

	return ParseError(rc)
}
