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
#include <opennsl/tx.h>
*/
import "C"

//
// opennsl_tx
//
func Tx(unit int, pkt *Pkt) error {
	rc := C.opennsl_tx(C.int(unit), pkt.C(), nil)
	return ParseError(rc)
}

func (v *Pkt) Tx(unit int) error {
	return Tx(unit, v)
}
