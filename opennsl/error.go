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
#include <opennsl/error.h>
*/
import "C"

import (
	"fmt"
)

//
// OpenNSLError
//
type OpenNSLError C.opennsl_error_t

func ParseError(e C.int) error {
	return OpenNSLError(e).Error()
}

func (e OpenNSLError) C() C.opennsl_error_t {
	return C.opennsl_error_t(e)
}

func (e OpenNSLError) Error() error {
	if e != E_NONE {
		return fmt.Errorf("%s", e)
	}

	return nil
}

const (
	E_NONE      OpenNSLError = C.OPENNSL_E_NONE
	E_INTERNAL  OpenNSLError = C.OPENNSL_E_INTERNAL
	E_MEMORY    OpenNSLError = C.OPENNSL_E_MEMORY
	E_UNIT      OpenNSLError = C.OPENNSL_E_UNIT
	E_PARAM     OpenNSLError = C.OPENNSL_E_PARAM
	E_EMPTY     OpenNSLError = C.OPENNSL_E_EMPTY
	E_FULL      OpenNSLError = C.OPENNSL_E_FULL
	E_NOT_FOUND OpenNSLError = C.OPENNSL_E_NOT_FOUND
	E_EXISTS    OpenNSLError = C.OPENNSL_E_EXISTS
	E_TIMEOUT   OpenNSLError = C.OPENNSL_E_TIMEOUT
	E_BUSY      OpenNSLError = C.OPENNSL_E_BUSY
	E_FAIL      OpenNSLError = C.OPENNSL_E_FAIL
	E_DISABLED  OpenNSLError = C.OPENNSL_E_DISABLED
	E_BADID     OpenNSLError = C.OPENNSL_E_BADID
	E_RESOURCE  OpenNSLError = C.OPENNSL_E_RESOURCE
	E_CONFIG    OpenNSLError = C.OPENNSL_E_CONFIG
	E_UNAVAIL   OpenNSLError = C.OPENNSL_E_UNAVAIL
	E_INIT      OpenNSLError = C.OPENNSL_E_INIT
	E_PORT      OpenNSLError = C.OPENNSL_E_PORT
)

var errorNames = map[OpenNSLError]string{
	E_NONE:      "E_NONE",
	E_INTERNAL:  "E_INTERNAL",
	E_MEMORY:    "E_MEMORY",
	E_UNIT:      "E_UNIT",
	E_PARAM:     "E_PARAM",
	E_EMPTY:     "E_EMPTY",
	E_FULL:      "E_FULL",
	E_NOT_FOUND: "E_NOT_FOUND",
	E_EXISTS:    "E_EXISTS",
	E_TIMEOUT:   "E_TIMEOUT",
	E_BUSY:      "E_BUSY",
	E_FAIL:      "E_FAIL",
	E_DISABLED:  "E_DISABLED",
	E_BADID:     "E_BADID",
	E_RESOURCE:  "E_RESOURCE",
	E_CONFIG:    "E_CONFIG",
	E_UNAVAIL:   "E_UNAVAIL",
	E_INIT:      "E_INIT",
	E_PORT:      "E_PORT",
}

func (v OpenNSLError) String() string {
	if s, ok := errorNames[v]; ok {
		return s
	}
	return fmt.Sprintf("OpenNSLError(%d)", v)
}

//
// SwitchEvent
//
type SwitchEvent C.opennsl_switch_event_t

func (v SwitchEvent) C() C.opennsl_switch_event_t {
	return C.opennsl_switch_event_t(v)
}

const (
	SWITCH_EVENT_NONE                  SwitchEvent = 0
	SWITCH_EVENT_PARITY_ERROR          SwitchEvent = C.OPENNSL_SWITCH_EVENT_PARITY_ERROR
	SWITCH_EVENT_STABLE_FULL           SwitchEvent = C.OPENNSL_SWITCH_EVENT_STABLE_FULL
	SWITCH_EVENT_STABLE_ERROR          SwitchEvent = C.OPENNSL_SWITCH_EVENT_STABLE_ERROR
	SWITCH_EVENT_UNCONTROLLED_SHUTDOWN SwitchEvent = C.OPENNSL_SWITCH_EVENT_UNCONTROLLED_SHUTDOWN
	SWITCH_EVENT_WARM_BOOT_DOWNGRADE   SwitchEvent = C.OPENNSL_SWITCH_EVENT_WARM_BOOT_DOWNGRADE
	SWITCH_EVENT_MMU_BST_TRIGGER       SwitchEvent = C.OPENNSL_SWITCH_EVENT_MMU_BST_TRIGGER
)

var switchEvent_names = map[SwitchEvent]string{
	SWITCH_EVENT_NONE:                  "NONE",
	SWITCH_EVENT_PARITY_ERROR:          "PARITY_ERROR",
	SWITCH_EVENT_STABLE_FULL:           "STABLE_FULL",
	SWITCH_EVENT_STABLE_ERROR:          "STABLE_ERROR",
	SWITCH_EVENT_UNCONTROLLED_SHUTDOWN: "UNCONTROLLED_SHUTDOWN",
	SWITCH_EVENT_WARM_BOOT_DOWNGRADE:   "WARM_BOOT_DOWNGRADE",
	SWITCH_EVENT_MMU_BST_TRIGGER:       "MMU_BST_TRIGGER",
}

var switchEvent_values = map[string]SwitchEvent{
	"PARITY_ERROR":          SWITCH_EVENT_PARITY_ERROR,
	"STABLE_FULL":           SWITCH_EVENT_STABLE_FULL,
	"STABLE_ERROR":          SWITCH_EVENT_STABLE_ERROR,
	"UNCONTROLLED_SHUTDOWN": SWITCH_EVENT_UNCONTROLLED_SHUTDOWN,
	"WARM_BOOT_DOWNGRADE":   SWITCH_EVENT_WARM_BOOT_DOWNGRADE,
	"MMU_BST_TRIGGER":       SWITCH_EVENT_MMU_BST_TRIGGER,
}

func (v SwitchEvent) String() string {
	if s, ok := switchEvent_names[v]; ok {
		return s
	}
	return fmt.Sprintf("SwitchEvent(%d)", v)
}

func ParseSwitchEvent(s string) (SwitchEvent, error) {
	if v, ok := switchEvent_values[s]; ok {
		return v, nil
	}
	return SWITCH_EVENT_NONE, fmt.Errorf("Invalid SwitchEvent. %s", s)
}
