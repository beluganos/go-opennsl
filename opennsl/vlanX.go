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
#include <opennsl/vlanX.h>
*/
import "C"

//
// VlanProtocolPacketCtrl
//
type VlanProtocolPacketCtrl C.opennsl_vlan_protocol_packet_ctrl_t

func (v *VlanProtocolPacketCtrl) C() *C.opennsl_vlan_protocol_packet_ctrl_t {
	return (*C.opennsl_vlan_protocol_packet_ctrl_t)(v)
}

//
// VlanControlVlan
//
type VlanControlVlan C.opennsl_vlan_control_vlan_t

func (v *VlanControlVlan) C() *C.opennsl_vlan_control_vlan_t {
	return (*C.opennsl_vlan_control_vlan_t)(v)
}
