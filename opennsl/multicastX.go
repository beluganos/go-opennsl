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
#include <opennsl/error.h>
#include <opennsl/multicast.h>
#include <opennsl/multicastX.h>
*/
import "C"

//
// API
//
func MulticasVxlanEncapGet(unit int, mc Multicast, port GPort, vxlanPort GPort) (EncapID, error) {
	c_encap := C.opennsl_if_t(0)

	rc := C.opennsl_multicast_vxlan_encap_get(C.int(unit), mc.C(), port.C(), vxlanPort.C(), &c_encap)
	return EncapID(c_encap), ParseError(rc)
}

func (v Multicast) VxlanEncapGet(unit int, port GPort, vxlanPort GPort) (EncapID, error) {
	return MulticasVxlanEncapGet(unit, v, port, vxlanPort)
}

func MulticastControlSet(unit int, group Multicast, mctrl MulticastControl, arg int) error {
	rc := C.opennsl_multicast_control_set(C.int(unit), group.C(), mctrl.C(), C.int(arg))
	return ParseError(rc)
}

func (v Multicast) ControlSet(unit int, mctrl MulticastControl, arg int) error {
	return MulticastControlSet(unit, v, mctrl, arg)
}

func MulticastControlGet(unit int, group Multicast, mctrl MulticastControl) (int, error) {
	c_arg := C.int(0)

	rc := C.opennsl_multicast_control_get(C.int(unit), group.C(), mctrl.C(), &c_arg)
	return int(c_arg), ParseError(rc)
}

func (v Multicast) ControlGet(unit int, mctrl MulticastControl) (int, error) {
	return MulticastControlGet(unit, v, mctrl)
}
