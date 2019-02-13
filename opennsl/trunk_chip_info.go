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
#include <opennsl/trunk.h>
*/
import "C"

//
// TrunkChipInfo
//
type TrunkChipInfo C.opennsl_trunk_chip_info_t

func (v *TrunkChipInfo) C() *C.opennsl_trunk_chip_info_t {
	return (*C.opennsl_trunk_chip_info_t)(v)
}

func (v *TrunkChipInfo) GroupCount() int {
	return int(v.trunk_group_count)
}

func (v *TrunkChipInfo) SetGroupCount(count int) {
	v.trunk_group_count = C.int(count)
}

func (v *TrunkChipInfo) IDMin() int {
	return int(v.trunk_id_min)
}

func (v *TrunkChipInfo) SetIDMin(min int) {
	v.trunk_id_min = C.int(min)
}

func (v *TrunkChipInfo) IDMax() int {
	return int(v.trunk_id_max)
}

func (v *TrunkChipInfo) SetIDMax(max int) {
	v.trunk_id_max = C.int(max)
}

func (v *TrunkChipInfo) PortsMax() int {
	return int(v.trunk_ports_max)
}

func (v *TrunkChipInfo) SetPortsMax(max int) {
	v.trunk_ports_max = C.int(max)
}

func (v *TrunkChipInfo) FabricIDMin() int {
	return int(v.trunk_fabric_id_min)
}

func (v *TrunkChipInfo) SetFabricIDMin(min int) {
	v.trunk_fabric_id_min = C.int(min)
}

func (v *TrunkChipInfo) FabricIDMax() int {
	return int(v.trunk_fabric_id_max)
}

func (v *TrunkChipInfo) SetFabricIDMax(max int) {
	v.trunk_fabric_id_max = C.int(max)
}

func (v *TrunkChipInfo) FabricPortsMax() int {
	return int(v.trunk_fabric_ports_max)
}

func (v *TrunkChipInfo) SetFabricPortsMax(max int) {
	v.trunk_fabric_ports_max = C.int(max)
}

func (v *TrunkChipInfo) VPIDMin() int {
	return int(v.vp_id_min)
}

func (v *TrunkChipInfo) SetVPIDMin(min int) {
	v.vp_id_min = C.int(min)
}

func (v *TrunkChipInfo) VPIDMax() int {
	return int(v.vp_id_max)
}

func (v *TrunkChipInfo) SetVPIDMax(max int) {
	v.vp_id_max = C.int(max)
}

func (v *TrunkChipInfo) VPPortsMax() int {
	return int(v.vp_ports_max)
}

func (v *TrunkChipInfo) SetVPPortsMax(max int) {
	v.vp_ports_max = C.int(max)
}

//
// API
//
func TrunkChipInfoGet(unit int) (*TrunkChipInfo, error) {
	info := &TrunkChipInfo{}
	rc := C.opennsl_trunk_chip_info_get(C.int(unit), info.C())
	if err := ParseError(rc); err != nil {
		return nil, err
	}
	return info, nil
}
