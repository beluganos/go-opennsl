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
*/
import "C"

type SwitchTemperatureMonitor C.opennsl_switch_temperature_monitor_t

func (v *SwitchTemperatureMonitor) C() *C.opennsl_switch_temperature_monitor_t {
	return (*C.opennsl_switch_temperature_monitor_t)(v)
}

func (v *SwitchTemperatureMonitor) Curr() int {
	return int(v.curr)
}

func (v *SwitchTemperatureMonitor) SetCurr(curr int) {
	v.curr = C.int(curr)
}

func (v *SwitchTemperatureMonitor) Peak() int {
	return int(v.peak)
}

func (v *SwitchTemperatureMonitor) SetPeak(peak int) {
	v.peak = C.int(peak)
}

func SwitchTemperatureMonitorGet(unit int, temperatureMax int) ([]SwitchTemperatureMonitor, error) {
	var count C.int = 0
	var arr = make([]C.opennsl_switch_temperature_monitor_t, temperatureMax)
	rc := C.opennsl_switch_temperature_monitor_get(C.int(unit), C.int(temperatureMax), &arr[0], &count)

	temps := make([]SwitchTemperatureMonitor, count)
	for index := 0; index < int(count); index++ {
		temps[index] = SwitchTemperatureMonitor(arr[index])
	}
	return temps, ParseError(rc)
}
