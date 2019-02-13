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
*/
import "C"
import "fmt"

//
// VlanAction
//
type VlanAction C.opennsl_vlan_action_t

func (v VlanAction) C() C.opennsl_vlan_action_t {
	return C.opennsl_vlan_action_t(v)
}

const (
	VlanActionNone          VlanAction = C.opennslVlanActionNone
	VlanActionAdd           VlanAction = C.opennslVlanActionAdd
	VlanActionReplace       VlanAction = C.opennslVlanActionReplace
	VlanActionDelete        VlanAction = C.opennslVlanActionDelete
	VlanActionCopy          VlanAction = C.opennslVlanActionCopy
	VlanActionCompressed    VlanAction = C.opennslVlanActionCompressed
	VlanActionMappedAdd     VlanAction = C.opennslVlanActionMappedAdd
	VlanActionMappedReplace VlanAction = C.opennslVlanActionMappedReplace
	VlanActionOuterAdd      VlanAction = C.opennslVlanActionOuterAdd
	VlanActionInnerAdd      VlanAction = C.opennslVlanActionInnerAdd
)

//
// VlanPcpAction
//
type VlanPcpAction C.opennsl_vlan_pcp_action_t

func (v VlanPcpAction) C() C.opennsl_vlan_pcp_action_t {
	return C.opennsl_vlan_pcp_action_t(v)
}

const (
	VlanPcpActionNone            VlanPcpAction = C.opennslVlanPcpActionNone
	VlanPcpActionMapped          VlanPcpAction = C.opennslVlanPcpActionMapped
	VlanPcpActionIngressInnerPcp VlanPcpAction = C.opennslVlanPcpActionIngressInnerPcp
	VlanPcpActionIngressOuterPcp VlanPcpAction = C.opennslVlanPcpActionIngressOuterPcp
	VlanPcpActionPortDefault     VlanPcpAction = C.opennslVlanPcpActionPortDefault
)

var vlanPcpAction_names = map[VlanPcpAction]string{
	VlanPcpActionNone:            "None",
	VlanPcpActionMapped:          "Mapped",
	VlanPcpActionIngressInnerPcp: "IngressInnerPcp",
	VlanPcpActionIngressOuterPcp: "IngressOuterPcp",
	VlanPcpActionPortDefault:     "Default",
}

var vlanPcpAction_values = map[string]VlanPcpAction{
	"None":            VlanPcpActionNone,
	"Mapped":          VlanPcpActionMapped,
	"IngressInnerPcp": VlanPcpActionIngressInnerPcp,
	"IngressOuterPcp": VlanPcpActionIngressOuterPcp,
	"Default":         VlanPcpActionPortDefault,
}

func (v VlanPcpAction) String() string {
	if s, ok := vlanPcpAction_names[v]; ok {
		return s
	}
	return fmt.Sprintf("VlanPcpAction(%d)", v)
}

func ParseVlanPcpAction(s string) (VlanPcpAction, error) {
	if v, ok := vlanPcpAction_values[s]; ok {
		return v, nil
	}
	return VlanPcpActionNone, fmt.Errorf("Invalid VlanPcpAction. %s", s)
}

//
// VlanTpidAction
//
type VlanTpidAction C.opennsl_vlan_tpid_action_t

func (v VlanTpidAction) C() C.opennsl_vlan_tpid_action_t {
	return C.opennsl_vlan_tpid_action_t(v)
}

const (
	VlanTpidActionNone   VlanTpidAction = C.opennslVlanTpidActionNone
	VlanTpidActionModify VlanTpidAction = C.opennslVlanTpidActionModify
	VlanTpidActionInner  VlanTpidAction = C.opennslVlanTpidActionInner
	VlanTpidActionOuter  VlanTpidAction = C.opennslVlanTpidActionOuter
)

var vlanTpidAction_names = map[VlanTpidAction]string{
	VlanTpidActionNone:   "None",
	VlanTpidActionModify: "Modify",
	VlanTpidActionInner:  "Inner",
	VlanTpidActionOuter:  "Outer",
}

var vlanTpidAction_values = map[string]VlanTpidAction{
	"None":   VlanTpidActionNone,
	"Modify": VlanTpidActionModify,
	"Inner":  VlanTpidActionInner,
	"Outer":  VlanTpidActionOuter,
}

func (v VlanTpidAction) String() string {
	if s, ok := vlanTpidAction_names[v]; ok {
		return s
	}
	return fmt.Sprintf("VlanTpidAction(%d)", v)
}

func ParseVlanTpidAction(s string) (VlanTpidAction, error) {
	if v, ok := vlanTpidAction_values[s]; ok {
		return v, nil
	}
	return VlanTpidActionNone, fmt.Errorf("Invalid VlanTpidAction. %s", s)
}
