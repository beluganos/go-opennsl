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

// "unsafe"

//
// SwitchObject
//
type SwitchObject C.opennsl_switch_object_t

func (v SwitchObject) C() C.opennsl_switch_object_t {
	return C.opennsl_switch_object_t(v)
}

const (
	SwitchObjectL2EntryCurrent                   SwitchObject = C.opennslSwitchObjectL2EntryCurrent
	SwitchObjectVlanCurrent                      SwitchObject = C.opennslSwitchObjectVlanCurrent
	SwitchObjectL3HostCurrent                    SwitchObject = C.opennslSwitchObjectL3HostCurrent
	SwitchObjectL3RouteCurrent                   SwitchObject = C.opennslSwitchObjectL3RouteCurrent
	SwitchObjectL3EgressCurrent                  SwitchObject = C.opennslSwitchObjectL3EgressCurrent
	SwitchObjectIpmcCurrent                      SwitchObject = C.opennslSwitchObjectIpmcCurrent
	SwitchObjectEcmpCurrent                      SwitchObject = C.opennslSwitchObjectEcmpCurrent
	SwitchObjectL3RouteV4RoutesMax               SwitchObject = C.opennslSwitchObjectL3RouteV4RoutesMax
	SwitchObjectL3RouteV4RoutesFree              SwitchObject = C.opennslSwitchObjectL3RouteV4RoutesFree
	SwitchObjectL3RouteV4RoutesUsed              SwitchObject = C.opennslSwitchObjectL3RouteV4RoutesUsed
	SwitchObjectL3RouteV6Routes64bMax            SwitchObject = C.opennslSwitchObjectL3RouteV6Routes64bMax
	SwitchObjectL3RouteV6Routes64bFree           SwitchObject = C.opennslSwitchObjectL3RouteV6Routes64bFree
	SwitchObjectL3RouteV6Routes64bUsed           SwitchObject = C.opennslSwitchObjectL3RouteV6Routes64bUsed
	SwitchObjectL3RouteV6Routes128bMax           SwitchObject = C.opennslSwitchObjectL3RouteV6Routes128bMax
	SwitchObjectL3RouteV6Routes128bFree          SwitchObject = C.opennslSwitchObjectL3RouteV6Routes128bFree
	SwitchObjectL3RouteV6Routes128bUsed          SwitchObject = C.opennslSwitchObjectL3RouteV6Routes128bUsed
	SwitchObjectL3RouteTotalUsedRoutes           SwitchObject = C.opennslSwitchObjectL3RouteTotalUsedRoutes
	SwitchObjectIpmcHeadTableFree                SwitchObject = C.opennslSwitchObjectIpmcHeadTableFree
	SwitchObjectL3HostV4Used                     SwitchObject = C.opennslSwitchObjectL3HostV4Used
	SwitchObjectL3HostV6Used                     SwitchObject = C.opennslSwitchObjectL3HostV6Used
	SwitchObjectEcmpMax                          SwitchObject = C.opennslSwitchObjectEcmpMax
	SwitchObjectPFCDeadlockCosMax                SwitchObject = C.opennslSwitchObjectPFCDeadlockCosMax
	SwitchObjectL3HostV4Max                      SwitchObject = C.opennslSwitchObjectL3HostV4Max
	SwitchObjectL3HostV6Max                      SwitchObject = C.opennslSwitchObjectL3HostV6Max
	SwitchObjectL3RouteV4RoutesMinGuaranteed     SwitchObject = C.opennslSwitchObjectL3RouteV4RoutesMinGuaranteed
	SwitchObjectL3RouteV6Routes64bMinGuaranteed  SwitchObject = C.opennslSwitchObjectL3RouteV6Routes64bMinGuaranteed
	SwitchObjectL3RouteV6Routes128bMinGuaranteed SwitchObject = C.opennslSwitchObjectL3RouteV6Routes128bMinGuaranteed
	SwitchObjectL3EgressMax                      SwitchObject = C.opennslSwitchObjectL3EgressMax
	SwitchObjectIpmcV4Used                       SwitchObject = C.opennslSwitchObjectIpmcV4Used
	SwitchObjectIpmcV6Used                       SwitchObject = C.opennslSwitchObjectIpmcV6Used
	SwitchObjectIpmcV4Max                        SwitchObject = C.opennslSwitchObjectIpmcV4Max
	SwitchObjectIpmcV6Max                        SwitchObject = C.opennslSwitchObjectIpmcV6Max
	SwitchObjectL2EntryMax                       SwitchObject = C.opennslSwitchObjectL2EntryMax
	SwitchObjectCount                            SwitchObject = C.opennslSwitchObjectCount
)

func SwitchObjectCountMultiGet(unit int, objects ...SwitchObject) ([]int, error) {
	size := len(objects)
	arr := make([]C.opennsl_switch_object_t, size)
	c_entries := make([]C.int, size)

	if size > 0 {
		for index, object := range objects {
			arr[index] = C.opennsl_switch_object_t(object)
		}

		rc := C.opennsl_switch_object_count_multi_get(C.int(unit), C.int(size), &arr[0], &c_entries[0])
		if err := ParseError(rc); err != nil {
			return nil, err
		}
	}

	entries := make([]int, size)
	for index, c_entry := range c_entries {
		entries[index] = int(c_entry)
	}

	return entries, nil
}

func (v SwitchObject) Get(unit int) ([]int, error) {
	return SwitchObjectCountMultiGet(unit, v)
}
