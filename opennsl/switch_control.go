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

import (
	"fmt"
)

//
// SwitchControl
//
type SwitchControl int

const (
	SwitchCpuSamplePrio                 SwitchControl = C.opennslSwitchCpuSamplePrio
	SwitchUnknownL3DestToCpu            SwitchControl = C.opennslSwitchUnknownL3DestToCpu
	SwitchSampleIngressRandomSeed       SwitchControl = C.opennslSwitchSampleIngressRandomSeed
	SwitchSampleEgressRandomSeed        SwitchControl = C.opennslSwitchSampleEgressRandomSeed
	SwitchV6L3DstMissToCpu              SwitchControl = C.opennslSwitchV6L3DstMissToCpu
	SwitchV4L3DstMissToCpu              SwitchControl = C.opennslSwitchV4L3DstMissToCpu
	SwitchL3SlowpathToCpu               SwitchControl = C.opennslSwitchL3SlowpathToCpu
	SwitchArpReplyToCpu                 SwitchControl = C.opennslSwitchArpReplyToCpu
	SwitchArpRequestToCpu               SwitchControl = C.opennslSwitchArpRequestToCpu
	SwitchNdPktToCpu                    SwitchControl = C.opennslSwitchNdPktToCpu
	SwitchIgmpPktToCpu                  SwitchControl = C.opennslSwitchIgmpPktToCpu
	SwitchIgmpToCPU                     SwitchControl = C.opennslSwitchIgmpToCPU
	SwitchDhcpPktToCpu                  SwitchControl = C.opennslSwitchDhcpPktToCpu
	SwitchDhcpPktDrop                   SwitchControl = C.opennslSwitchDhcpPktDrop
	SwitchV4ResvdMcPktToCpu             SwitchControl = C.opennslSwitchV4ResvdMcPktToCpu
	SwitchDirectedMirroring             SwitchControl = C.opennslSwitchDirectedMirroring
	SwitchHashControl                   SwitchControl = C.opennslSwitchHashControl
	SwitchMirrorUnmarked                SwitchControl = C.opennslSwitchMirrorUnmarked
	SwitchColorSelect                   SwitchControl = C.opennslSwitchColorSelect
	SwitchHashSeed0                     SwitchControl = C.opennslSwitchHashSeed0
	SwitchHashSeed1                     SwitchControl = C.opennslSwitchHashSeed1
	SwitchHashField0PreProcessEnable    SwitchControl = C.opennslSwitchHashField0PreProcessEnable
	SwitchHashField1PreProcessEnable    SwitchControl = C.opennslSwitchHashField1PreProcessEnable
	SwitchHashField0Config              SwitchControl = C.opennslSwitchHashField0Config
	SwitchHashField0Config1             SwitchControl = C.opennslSwitchHashField0Config1
	SwitchHashField1Config              SwitchControl = C.opennslSwitchHashField1Config
	SwitchHashField1Config1             SwitchControl = C.opennslSwitchHashField1Config1
	SwitchHashSelectControl             SwitchControl = C.opennslSwitchHashSelectControl
	SwitchHashIP4Field0                 SwitchControl = C.opennslSwitchHashIP4Field0
	SwitchHashIP4Field1                 SwitchControl = C.opennslSwitchHashIP4Field1
	SwitchHashIP4TcpUdpField0           SwitchControl = C.opennslSwitchHashIP4TcpUdpField0
	SwitchHashIP4TcpUdpField1           SwitchControl = C.opennslSwitchHashIP4TcpUdpField1
	SwitchHashIP4TcpUdpPortsEqualField0 SwitchControl = C.opennslSwitchHashIP4TcpUdpPortsEqualField0
	SwitchHashIP4TcpUdpPortsEqualField1 SwitchControl = C.opennslSwitchHashIP4TcpUdpPortsEqualField1
	SwitchHashIP6Field0                 SwitchControl = C.opennslSwitchHashIP6Field0
	SwitchHashIP6Field1                 SwitchControl = C.opennslSwitchHashIP6Field1
	SwitchHashIP6TcpUdpField0           SwitchControl = C.opennslSwitchHashIP6TcpUdpField0
	SwitchHashIP6TcpUdpField1           SwitchControl = C.opennslSwitchHashIP6TcpUdpField1
	SwitchHashIP6TcpUdpPortsEqualField0 SwitchControl = C.opennslSwitchHashIP6TcpUdpPortsEqualField0
	SwitchHashIP6TcpUdpPortsEqualField1 SwitchControl = C.opennslSwitchHashIP6TcpUdpPortsEqualField1
	SwitchHashL2Field0                  SwitchControl = C.opennslSwitchHashL2Field0
	SwitchHashL2Field1                  SwitchControl = C.opennslSwitchHashL2Field1
	SwitchECMPHashSet0Offset            SwitchControl = C.opennslSwitchECMPHashSet0Offset
	SwitchECMPHashSet1Offset            SwitchControl = C.opennslSwitchECMPHashSet1Offset
	SwitchMirrorInvalidVlanDrop         SwitchControl = C.opennslSwitchMirrorInvalidVlanDrop
	SwitchMirrorPktChecksEnable         SwitchControl = C.opennslSwitchMirrorPktChecksEnable
	SwitchL3EgressMode                  SwitchControl = C.opennslSwitchL3EgressMode
	SwitchL3IngressMode                 SwitchControl = C.opennslSwitchL3IngressMode
	SwitchWarmBoot                      SwitchControl = C.opennslSwitchWarmBoot
	SwitchStableSelect                  SwitchControl = C.opennslSwitchStableSelect
	SwitchStableSize                    SwitchControl = C.opennslSwitchStableSize
	SwitchStableUsed                    SwitchControl = C.opennslSwitchStableUsed
	SwitchStableConsistent              SwitchControl = C.opennslSwitchStableConsistent
	SwitchControlSync                   SwitchControl = C.opennslSwitchControlSync
	SwitchControlAutoSync               SwitchControl = C.opennslSwitchControlAutoSync
	SwitchIpmcTtl1ToCpu                 SwitchControl = C.opennslSwitchIpmcTtl1ToCpu
	SwitchL3UcastTtl1ToCpu              SwitchControl = C.opennslSwitchL3UcastTtl1ToCpu
	SwitchL3UrpfMode                    SwitchControl = C.opennslSwitchL3UrpfMode
	SwitchBstEnable                     SwitchControl = C.opennslSwitchBstEnable
	SwitchBstTrackingMode               SwitchControl = C.opennslSwitchBstTrackingMode
	SwitchVxlanUdpDestPortSet           SwitchControl = C.opennslSwitchVxlanUdpDestPortSet
	SwitchVxlanEntropyEnable            SwitchControl = C.opennslSwitchVxlanEntropyEnable
	SwitchVxlanVnIdMissToCpu            SwitchControl = C.opennslSwitchVxlanVnIdMissToCpu
	SwitchVxlanTunnelMissToCpu          SwitchControl = C.opennslSwitchVxlanTunnelMissToCpu
	SwitchFlexibleMirrorDestinations    SwitchControl = C.opennslSwitchFlexibleMirrorDestinations
	SwitchEcmpMacroFlowHashEnable       SwitchControl = C.opennslSwitchEcmpMacroFlowHashEnable
	SwitchMcQueueSchedMode              SwitchControl = C.opennslSwitchMcQueueSchedMode
	SwitchBstSnapshotEnable             SwitchControl = C.opennslSwitchBstSnapshotEnable
	SwitchMirrorExclusive               SwitchControl = C.opennslSwitchMirrorExclusive
)

func (v SwitchControl) String() string {
	if s, ok := switchControl_names[v]; ok {
		return s
	}
	return fmt.Sprintf("SwitchControl(%d)", v)
}

var switchControl_names = map[SwitchControl]string{
	SwitchCpuSamplePrio:                 "SwitchCpuSamplePrio",
	SwitchUnknownL3DestToCpu:            "SwitchUnknownL3DestToCpu",
	SwitchSampleIngressRandomSeed:       "SwitchSampleIngressRandomSeed",
	SwitchSampleEgressRandomSeed:        "SwitchSampleEgressRandomSeed",
	SwitchV6L3DstMissToCpu:              "SwitchV6L3DstMissToCpu",
	SwitchV4L3DstMissToCpu:              "SwitchV4L3DstMissToCpu",
	SwitchL3SlowpathToCpu:               "SwitchL3SlowpathToCpu",
	SwitchArpReplyToCpu:                 "SwitchArpReplyToCpu",
	SwitchArpRequestToCpu:               "SwitchArpRequestToCpu",
	SwitchNdPktToCpu:                    "SwitchNdPktToCpu",
	SwitchIgmpPktToCpu:                  "SwitchIgmpPktToCpu",
	SwitchDhcpPktToCpu:                  "SwitchDhcpPktToCpu",
	SwitchDhcpPktDrop:                   "SwitchDhcpPktDrop",
	SwitchV4ResvdMcPktToCpu:             "SwitchV4ResvdMcPktToCpu",
	SwitchDirectedMirroring:             "SwitchDirectedMirroring",
	SwitchHashControl:                   "SwitchHashControl",
	SwitchMirrorUnmarked:                "SwitchMirrorUnmarked",
	SwitchColorSelect:                   "SwitchColorSelect",
	SwitchHashSeed0:                     "SwitchHashSeed0",
	SwitchHashSeed1:                     "SwitchHashSeed1",
	SwitchHashField0PreProcessEnable:    "SwitchHashField0PreProcessEnable",
	SwitchHashField1PreProcessEnable:    "SwitchHashField1PreProcessEnable",
	SwitchHashField0Config:              "SwitchHashField0Config",
	SwitchHashField0Config1:             "SwitchHashField0Config1",
	SwitchHashField1Config:              "SwitchHashField1Config",
	SwitchHashField1Config1:             "SwitchHashField1Config1",
	SwitchHashSelectControl:             "SwitchHashSelectControl",
	SwitchHashIP4Field0:                 "SwitchHashIP4Field0",
	SwitchHashIP4Field1:                 "SwitchHashIP4Field1",
	SwitchHashIP4TcpUdpField0:           "SwitchHashIP4TcpUdpField0",
	SwitchHashIP4TcpUdpField1:           "SwitchHashIP4TcpUdpField1",
	SwitchHashIP4TcpUdpPortsEqualField0: "SwitchHashIP4TcpUdpPortsEqualField0",
	SwitchHashIP4TcpUdpPortsEqualField1: "SwitchHashIP4TcpUdpPortsEqualField1",
	SwitchHashIP6Field0:                 "SwitchHashIP6Field0",
	SwitchHashIP6Field1:                 "SwitchHashIP6Field1",
	SwitchHashIP6TcpUdpField0:           "SwitchHashIP6TcpUdpField0",
	SwitchHashIP6TcpUdpField1:           "SwitchHashIP6TcpUdpField1",
	SwitchHashIP6TcpUdpPortsEqualField0: "SwitchHashIP6TcpUdpPortsEqualField0",
	SwitchHashIP6TcpUdpPortsEqualField1: "SwitchHashIP6TcpUdpPortsEqualField1",
	SwitchHashL2Field0:                  "SwitchHashL2Field0",
	SwitchHashL2Field1:                  "SwitchHashL2Field1",
	SwitchECMPHashSet0Offset:            "SwitchECMPHashSet0Offset",
	SwitchECMPHashSet1Offset:            "SwitchECMPHashSet1Offset",
	SwitchMirrorInvalidVlanDrop:         "SwitchMirrorInvalidVlanDrop",
	SwitchMirrorPktChecksEnable:         "SwitchMirrorPktChecksEnable",
	SwitchL3EgressMode:                  "SwitchL3EgressMode",
	SwitchL3IngressMode:                 "SwitchL3IngressMode",
	SwitchWarmBoot:                      "SwitchWarmBoot",
	SwitchStableSelect:                  "SwitchStableSelect",
	SwitchStableSize:                    "SwitchStableSize",
	SwitchStableUsed:                    "SwitchStableUsed",
	SwitchStableConsistent:              "SwitchStableConsistent",
	SwitchControlSync:                   "SwitchControlSync",
	SwitchControlAutoSync:               "SwitchControlAutoSync",
	SwitchIpmcTtl1ToCpu:                 "SwitchIpmcTtl1ToCpu",
	SwitchL3UcastTtl1ToCpu:              "SwitchL3UcastTtl1ToCpu",
	SwitchL3UrpfMode:                    "SwitchL3UrpfMode",
	SwitchBstEnable:                     "SwitchBstEnable",
	SwitchBstTrackingMode:               "SwitchBstTrackingMode",
	SwitchVxlanUdpDestPortSet:           "SwitchVxlanUdpDestPortSet",
	SwitchVxlanEntropyEnable:            "SwitchVxlanEntropyEnable",
	SwitchVxlanVnIdMissToCpu:            "SwitchVxlanVnIdMissToCpu",
	SwitchVxlanTunnelMissToCpu:          "SwitchVxlanTunnelMissToCpu",
	SwitchFlexibleMirrorDestinations:    "SwitchFlexibleMirrorDestinations",
	SwitchEcmpMacroFlowHashEnable:       "SwitchEcmpMacroFlowHashEnable",
	SwitchMcQueueSchedMode:              "SwitchMcQueueSchedMode",
	SwitchBstSnapshotEnable:             "SwitchBstSnapshotEnable",
	SwitchMirrorExclusive:               "SwitchMirrorExclusive",
}

func ParseSwitchControl(s string) (SwitchControl, error) {
	if v, ok := switchControl_values[s]; ok {
		return v, nil
	}
	return 0, fmt.Errorf("Invalid SwitchControl(%s)", s)
}

var switchControl_values = map[string]SwitchControl{
	"SwitchCpuSamplePrio":                 SwitchCpuSamplePrio,
	"SwitchUnknownL3DestToCpu":            SwitchUnknownL3DestToCpu,
	"SwitchSampleIngressRandomSeed":       SwitchSampleIngressRandomSeed,
	"SwitchSampleEgressRandomSeed":        SwitchSampleEgressRandomSeed,
	"SwitchV6L3DstMissToCpu":              SwitchV6L3DstMissToCpu,
	"SwitchV4L3DstMissToCpu":              SwitchV4L3DstMissToCpu,
	"SwitchL3SlowpathToCpu":               SwitchL3SlowpathToCpu,
	"SwitchArpReplyToCpu":                 SwitchArpReplyToCpu,
	"SwitchArpRequestToCpu":               SwitchArpRequestToCpu,
	"SwitchNdPktToCpu":                    SwitchNdPktToCpu,
	"SwitchIgmpPktToCpu":                  SwitchIgmpPktToCpu,
	"SwitchIgmpToCPU":                     SwitchIgmpToCPU,
	"SwitchDhcpPktToCpu":                  SwitchDhcpPktToCpu,
	"SwitchDhcpPktDrop":                   SwitchDhcpPktDrop,
	"SwitchV4ResvdMcPktToCpu":             SwitchV4ResvdMcPktToCpu,
	"SwitchDirectedMirroring":             SwitchDirectedMirroring,
	"SwitchHashControl":                   SwitchHashControl,
	"SwitchMirrorUnmarked":                SwitchMirrorUnmarked,
	"SwitchColorSelect":                   SwitchColorSelect,
	"SwitchHashSeed0":                     SwitchHashSeed0,
	"SwitchHashSeed1":                     SwitchHashSeed1,
	"SwitchHashField0PreProcessEnable":    SwitchHashField0PreProcessEnable,
	"SwitchHashField1PreProcessEnable":    SwitchHashField1PreProcessEnable,
	"SwitchHashField0Config":              SwitchHashField0Config,
	"SwitchHashField0Config1":             SwitchHashField0Config1,
	"SwitchHashField1Config":              SwitchHashField1Config,
	"SwitchHashField1Config1":             SwitchHashField1Config1,
	"SwitchHashSelectControl":             SwitchHashSelectControl,
	"SwitchHashIP4Field0":                 SwitchHashIP4Field0,
	"SwitchHashIP4Field1":                 SwitchHashIP4Field1,
	"SwitchHashIP4TcpUdpField0":           SwitchHashIP4TcpUdpField0,
	"SwitchHashIP4TcpUdpField1":           SwitchHashIP4TcpUdpField1,
	"SwitchHashIP4TcpUdpPortsEqualField0": SwitchHashIP4TcpUdpPortsEqualField0,
	"SwitchHashIP4TcpUdpPortsEqualField1": SwitchHashIP4TcpUdpPortsEqualField1,
	"SwitchHashIP6Field0":                 SwitchHashIP6Field0,
	"SwitchHashIP6Field1":                 SwitchHashIP6Field1,
	"SwitchHashIP6TcpUdpField0":           SwitchHashIP6TcpUdpField0,
	"SwitchHashIP6TcpUdpField1":           SwitchHashIP6TcpUdpField1,
	"SwitchHashIP6TcpUdpPortsEqualField0": SwitchHashIP6TcpUdpPortsEqualField0,
	"SwitchHashIP6TcpUdpPortsEqualField1": SwitchHashIP6TcpUdpPortsEqualField1,
	"SwitchHashL2Field0":                  SwitchHashL2Field0,
	"SwitchHashL2Field1":                  SwitchHashL2Field1,
	"SwitchECMPHashSet0Offset":            SwitchECMPHashSet0Offset,
	"SwitchECMPHashSet1Offset":            SwitchECMPHashSet1Offset,
	"SwitchMirrorInvalidVlanDrop":         SwitchMirrorInvalidVlanDrop,
	"SwitchMirrorPktChecksEnable":         SwitchMirrorPktChecksEnable,
	"SwitchL3EgressMode":                  SwitchL3EgressMode,
	"SwitchL3IngressMode":                 SwitchL3IngressMode,
	"SwitchWarmBoot":                      SwitchWarmBoot,
	"SwitchStableSelect":                  SwitchStableSelect,
	"SwitchStableSize":                    SwitchStableSize,
	"SwitchStableUsed":                    SwitchStableUsed,
	"SwitchStableConsistent":              SwitchStableConsistent,
	"SwitchControlSync":                   SwitchControlSync,
	"SwitchControlAutoSync":               SwitchControlAutoSync,
	"SwitchIpmcTtl1ToCpu":                 SwitchIpmcTtl1ToCpu,
	"SwitchL3UcastTtl1ToCpu":              SwitchL3UcastTtl1ToCpu,
	"SwitchL3UrpfMode":                    SwitchL3UrpfMode,
	"SwitchBstEnable":                     SwitchBstEnable,
	"SwitchBstTrackingMode":               SwitchBstTrackingMode,
	"SwitchVxlanUdpDestPortSet":           SwitchVxlanUdpDestPortSet,
	"SwitchVxlanEntropyEnable":            SwitchVxlanEntropyEnable,
	"SwitchVxlanVnIdMissToCpu":            SwitchVxlanVnIdMissToCpu,
	"SwitchVxlanTunnelMissToCpu":          SwitchVxlanTunnelMissToCpu,
	"SwitchFlexibleMirrorDestinations":    SwitchFlexibleMirrorDestinations,
	"SwitchEcmpMacroFlowHashEnable":       SwitchEcmpMacroFlowHashEnable,
	"SwitchMcQueueSchedMode":              SwitchMcQueueSchedMode,
	"SwitchBstSnapshotEnable":             SwitchBstSnapshotEnable,
	"SwitchMirrorExclusive":               SwitchMirrorExclusive,
}

func (v SwitchControl) C() C.opennsl_switch_control_t {
	return (C.opennsl_switch_control_t)(v)
}

func (v SwitchControl) Get(unit int) (int, error) {
	var arg C.int = 0
	rc := C.opennsl_switch_control_get(C.int(unit), v.C(), &arg)
	return int(arg), ParseError(rc)
}

func (v SwitchControl) Set(unit int, arg int) error {
	rc := C.opennsl_switch_control_set(C.int(unit), v.C(), C.int(arg))
	return ParseError(rc)
}

func (v SwitchControl) PortGet(unit int, port Port) (int, error) {
	var arg C.int = 0
	rc := C.opennsl_switch_control_port_get(C.int(unit), port.C(), v.C(), &arg)
	return int(arg), ParseError(rc)
}

func (v SwitchControl) PortSet(unit int, port Port, arg int) error {
	rc := C.opennsl_switch_control_port_set(C.int(unit), port.C(), v.C(), C.int(arg))
	return ParseError(rc)
}

//
// SwitchControlEntry
//
type SwitchControlEntry struct {
	Ctrl  SwitchControl
	Value int
}

func NewSwitchControlEntry(ctrl SwitchControl, value int) *SwitchControlEntry {
	return &SwitchControlEntry{
		Ctrl:  ctrl,
		Value: value,
	}
}

func (v SwitchControl) Arg(value int) *SwitchControlEntry {
	return NewSwitchControlEntry(v, value)
}

func (e *SwitchControlEntry) Set(unit int) error {
	return e.Ctrl.Set(unit, e.Value)
}

func SwitchControlsSet(unit int, entries ...*SwitchControlEntry) error {
	for _, entry := range entries {
		if err := entry.Set(unit); err != nil {
			return err
		}
	}
	return nil
}
