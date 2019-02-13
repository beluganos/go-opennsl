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
#include <opennsl/stat.h>
*/
import "C"

//
// StatObject
//
type StatObject C.opennsl_stat_object_t

func (v StatObject) C() C.opennsl_stat_object_t {
	return C.opennsl_stat_object_t(v)
}

const (
	StatObjectIngL3Intf StatObject = C.opennslStatObjectIngL3Intf
	StatObjectEgrL3Intf StatObject = C.opennslStatObjectEgrL3Intf
)

//
// StatGroupModeAttrPktType
//
type StatGroupModeAttrPktType C.opennsl_stat_group_mode_attr_pkt_type_t

func (v StatGroupModeAttrPktType) C() C.opennsl_stat_group_mode_attr_pkt_type_t {
	return C.opennsl_stat_group_mode_attr_pkt_type_t(v)
}

const (
	StatGroupModeAttrPktTypeKnownL3UC   StatGroupModeAttrPktType = C.opennslStatGroupModeAttrPktTypeKnownL3UC
	StatGroupModeAttrPktTypeUnknownL3UC StatGroupModeAttrPktType = C.opennslStatGroupModeAttrPktTypeUnknownL3UC
	StatGroupModeAttrPktTypeKnownIPMC   StatGroupModeAttrPktType = C.opennslStatGroupModeAttrPktTypeKnownIPMC
	StatGroupModeAttrPktTypeUnknownIPMC StatGroupModeAttrPktType = C.opennslStatGroupModeAttrPktTypeUnknownIPMC
)

//
// StatGroupMode
//
type StatGroupMode C.opennsl_stat_group_mode_t

func (v StatGroupMode) C() C.opennsl_stat_group_mode_t {
	return C.opennsl_stat_group_mode_t(v)
}

const (
	StatGroupModeSingle      StatGroupMode = C.opennslStatGroupModeSingle
	StatGroupModeTrafficType StatGroupMode = C.opennslStatGroupModeTrafficType
)
