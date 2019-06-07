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
#include <opennsl/port.h>
*/
import "C"

//
// PortPhyControl
//
type PortPhyControl C.opennsl_port_phy_control_t

func (v PortPhyControl) C() C.opennsl_port_phy_control_t {
	return C.opennsl_port_phy_control_t(v)
}

const (
	PORT_PHY_CONTROL_FORWARD_ERROR_CORRECTION           PortPhyControl = C.OPENNSL_PORT_PHY_CONTROL_FORWARD_ERROR_CORRECTION
	PORT_PHY_CONTROL_SOFTWARE_RX_LOS                    PortPhyControl = C.OPENNSL_PORT_PHY_CONTROL_SOFTWARE_RX_LOS
	PORT_PHY_CONTROL_SOFTWARE_RX_LOS_LINK_WAIT_TIMER_US PortPhyControl = C.OPENNSL_PORT_PHY_CONTROL_SOFTWARE_RX_LOS_LINK_WAIT_TIMER_US
	PORT_PHY_CONTROL_SOFTWARE_RX_LOS_RESTART_TIMER_US   PortPhyControl = C.OPENNSL_PORT_PHY_CONTROL_SOFTWARE_RX_LOS_RESTART_TIMER_US
)

const (
	// PORT_PHY_CONTROL_FORWARD_ERROR_CORRECTION
	PORT_PHY_CONTROL_FEC_OFF  = C.OPENNSL_PORT_PHY_CONTROL_FEC_OFF
	PORT_PHY_CONTROL_FEC_ON   = C.OPENNSL_PORT_PHY_CONTROL_FEC_ON
	PORT_PHY_CONTROL_FEC_AUTO = C._SHR_PORT_PHY_CONTROL_FEC_AUTO
	// PORT_PHY_CONTROL_SOFTWARE_RX_LOS
	PORT_PHY_CONTROL_RX_LOS_NONE     = C.OPENNSL_PORT_PHY_CONTROL_RX_LOS_NONE
	PORT_PHY_CONTROL_RX_LOS_SOFTWARE = C.OPENNSL_PORT_PHY_CONTROL_RX_LOS_SOFTWARE
	PORT_PHY_CONTROL_RX_LOS_FIRMWARE = C.OPENNSL_PORT_PHY_CONTROL_RX_LOS_FIRMWARE
)

func PortPhyControlSet(unit int, port Port, ctrl PortPhyControl, value uint32) error {
	rc := C.opennsl_port_phy_control_set(C.int(unit), port.C(), ctrl.C(), C.uint32(value))
	return ParseError(rc)
}

func PortPhyControlGet(unit int, port Port, ctrl PortPhyControl) (uint32, error) {
	c_val := C.uint32(0)

	rc := C.opennsl_port_phy_control_get(C.int(unit), port.C(), ctrl.C(), &c_val)
	return uint32(c_val), ParseError(rc)
}

func (v PortPhyControl) Set(unit int, port Port, value uint32) error {
	return PortPhyControlSet(unit, port, v, value)
}

//
// PortControl
//
type PortControl C.opennsl_port_control_t

func (v PortControl) C() C.opennsl_port_control_t {
	return C.opennsl_port_control_t(v)
}

const (
	PortControlNone                     PortControl = 0
	PortControlIP4                      PortControl = C.opennslPortControlIP4
	PortControlIP6                      PortControl = C.opennslPortControlIP6
	PortControlTrustIncomingVlan        PortControl = C.opennslPortControlTrustIncomingVlan
	PortControlDoNotCheckVlan           PortControl = C.opennslPortControlDoNotCheckVlan
	PortControlPrbsMode                 PortControl = C.opennslPortControlPrbsMode
	PortControlPrbsPolynomial           PortControl = C.opennslPortControlPrbsPolynomial
	PortControlPrbsTxInvertData         PortControl = C.opennslPortControlPrbsTxInvertData
	PortControlPrbsForceTxError         PortControl = C.opennslPortControlPrbsForceTxError
	PortControlPrbsTxEnable             PortControl = C.opennslPortControlPrbsTxEnable
	PortControlPrbsRxEnable             PortControl = C.opennslPortControlPrbsRxEnable
	PortControlPrbsRxStatus             PortControl = C.opennslPortControlPrbsRxStatus
	PortControlEgressVlanPriUsesPktPri  PortControl = C.opennslPortControlEgressVlanPriUsesPktPri
	PortControlLanes                    PortControl = C.opennslPortControlLanes
	PortControlPFCReceive               PortControl = C.opennslPortControlPFCReceive
	PortControlPFCTransmit              PortControl = C.opennslPortControlPFCTransmit
	PortControlPFCClasses               PortControl = C.opennslPortControlPFCClasses
	PortControlPFCPassFrames            PortControl = C.opennslPortControlPFCPassFrames
	PortControlL2Move                   PortControl = C.opennslPortControlL2Move
	PortControlEEETransmitWakeTime      PortControl = C.opennslPortControlEEETransmitWakeTime
	PortControlStatOversize             PortControl = C.opennslPortControlStatOversize
	PortControlEEEEnable                PortControl = C.opennslPortControlEEEEnable
	PortControlEEETransmitIdleTime      PortControl = C.opennslPortControlEEETransmitIdleTime
	PortControlVxlanEnable              PortControl = C.opennslPortControlVxlanEnable
	PortControlVxlanTunnelbasedVnId     PortControl = C.opennslPortControlVxlanTunnelbasedVnId
	PortControlVxlanDefaultTunnelEnable PortControl = C.opennslPortControlVxlanDefaultTunnelEnable
	PortControlMmuDrain                 PortControl = C.opennslPortControlMmuDrain
	PortControlMmuTrafficEnable         PortControl = C.opennslPortControlMmuTrafficEnable
	PortControlSampleIngressDest        PortControl = C.opennslPortControlSampleIngressDest
	PortControlSampleFlexDest           PortControl = C.opennslPortControlSampleFlexDest
	PortControlSampleFlexRate           PortControl = C.opennslPortControlSampleFlexRate
)

func PortControlSet(unit int, port Port, pc PortControl, val int) error {
	rc := C.opennsl_port_control_set(C.int(unit), port.C(), pc.C(), C.int(val))
	return ParseError(rc)
}

func (v PortControl) Set(unit int, port Port, val int) error {
	return PortControlSet(unit, port, v, val)
}

func PortControlGet(unit int, port Port, pc PortControl) (int, error) {
	c_val := C.int(0)

	rc := C.opennsl_port_control_get(C.int(unit), port.C(), pc.C(), &c_val)
	return int(c_val), ParseError(rc)
}

func (v PortControl) Get(unit int, port Port) (int, error) {
	return PortControlGet(unit, port, v)
}

//
// PortControlEntry
//
type PortControlEntry struct {
	Ctrl  PortControl
	Value int
}

func NewPortControlEntry(ctrl PortControl, value int) *PortControlEntry {
	return &PortControlEntry{
		Ctrl:  ctrl,
		Value: value,
	}
}

func (c PortControl) Arg(value int) *PortControlEntry {
	return NewPortControlEntry(c, value)
}

func (e *PortControlEntry) Set(unit int, port Port) error {
	return e.Ctrl.Set(unit, port, e.Value)
}

//
// Port method
//
func (v Port) PortControlsSet(unit int, entries ...*PortControlEntry) error {
	for _, entry := range entries {
		if err := entry.Set(unit, v); err != nil {
			return err
		}
	}
	return nil
}

func (v Port) PortControlGet(unit int, ctrl PortControl) (int, error) {
	return ctrl.Get(unit, v)
}

func (v Port) PhyControlSet(unit int, ctrl PortPhyControl, value uint32) error {
	return PortPhyControlSet(unit, v, ctrl, value)
}

func (v Port) PhyControlGet(unit int, ctrl PortPhyControl) (uint32, error) {
	return PortPhyControlGet(unit, v, ctrl)
}
