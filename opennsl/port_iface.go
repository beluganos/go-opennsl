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
// PortIface
//
type PortIface C.opennsl_port_if_t

func (v PortIface) C() C.opennsl_port_if_t {
	return C.opennsl_port_if_t(v)
}

const (
	PORT_IF_NOCXN      PortIface = C.OPENNSL_PORT_IF_NOCXN
	PORT_IF_NULL       PortIface = C.OPENNSL_PORT_IF_NULL
	PORT_IF_MII        PortIface = C.OPENNSL_PORT_IF_MII
	PORT_IF_GMII       PortIface = C.OPENNSL_PORT_IF_GMII
	PORT_IF_SGMII      PortIface = C.OPENNSL_PORT_IF_SGMII
	PORT_IF_TBI        PortIface = C.OPENNSL_PORT_IF_TBI
	PORT_IF_XGMII      PortIface = C.OPENNSL_PORT_IF_XGMII
	PORT_IF_RGMII      PortIface = C.OPENNSL_PORT_IF_RGMII
	PORT_IF_SFI        PortIface = C.OPENNSL_PORT_IF_SFI
	PORT_IF_XFI        PortIface = C.OPENNSL_PORT_IF_XFI
	PORT_IF_KR         PortIface = C.OPENNSL_PORT_IF_KR
	PORT_IF_KR2        PortIface = C.OPENNSL_PORT_IF_KR2
	PORT_IF_KR4        PortIface = C.OPENNSL_PORT_IF_KR4
	PORT_IF_CR         PortIface = C.OPENNSL_PORT_IF_CR
	PORT_IF_CR2        PortIface = C.OPENNSL_PORT_IF_CR2
	PORT_IF_CR4        PortIface = C.OPENNSL_PORT_IF_CR4
	PORT_IF_XLAUI      PortIface = C.OPENNSL_PORT_IF_XLAUI
	PORT_IF_XLAUI2     PortIface = C.OPENNSL_PORT_IF_XLAUI2
	PORT_IF_RXAUI      PortIface = C.OPENNSL_PORT_IF_RXAUI
	PORT_IF_XAUI       PortIface = C.OPENNSL_PORT_IF_XAUI
	PORT_IF_SPAUI      PortIface = C.OPENNSL_PORT_IF_SPAUI
	PORT_IF_QSGMII     PortIface = C.OPENNSL_PORT_IF_QSGMII
	PORT_IF_ILKN       PortIface = C.OPENNSL_PORT_IF_ILKN
	PORT_IF_RCY        PortIface = C.OPENNSL_PORT_IF_RCY
	PORT_IF_FAT_PIPE   PortIface = C.OPENNSL_PORT_IF_FAT_PIPE
	PORT_IF_SR         PortIface = C.OPENNSL_PORT_IF_SR
	PORT_IF_SR2        PortIface = C.OPENNSL_PORT_IF_SR2
	PORT_IF_CAUI       PortIface = C.OPENNSL_PORT_IF_CAUI
	PORT_IF_LR         PortIface = C.OPENNSL_PORT_IF_LR
	PORT_IF_LR4        PortIface = C.OPENNSL_PORT_IF_LR4
	PORT_IF_SR4        PortIface = C.OPENNSL_PORT_IF_SR4
	PORT_IF_KX         PortIface = C.OPENNSL_PORT_IF_KX
	PORT_IF_ZR         PortIface = C.OPENNSL_PORT_IF_ZR
	PORT_IF_SR10       PortIface = C.OPENNSL_PORT_IF_SR10
	PORT_IF_CR10       PortIface = C.OPENNSL_PORT_IF_CR10
	PORT_IF_KR10       PortIface = C.OPENNSL_PORT_IF_KR10
	PORT_IF_LR10       PortIface = C.OPENNSL_PORT_IF_LR10
	PORT_IF_OTL        PortIface = C.OPENNSL_PORT_IF_OTL
	PORT_IF_CPU        PortIface = C.OPENNSL_PORT_IF_CPU
	PORT_IF_ER         PortIface = C.OPENNSL_PORT_IF_ER
	PORT_IF_ER2        PortIface = C.OPENNSL_PORT_IF_ER2
	PORT_IF_ER4        PortIface = C.OPENNSL_PORT_IF_ER4
	PORT_IF_CX         PortIface = C.OPENNSL_PORT_IF_CX
	PORT_IF_CX2        PortIface = C.OPENNSL_PORT_IF_CX2
	PORT_IF_CX4        PortIface = C.OPENNSL_PORT_IF_CX4
	PORT_IF_CAUI_C2C   PortIface = C.OPENNSL_PORT_IF_CAUI_C2C
	PORT_IF_CAUI_C2M   PortIface = C.OPENNSL_PORT_IF_CAUI_C2M
	PORT_IF_VSR        PortIface = C.OPENNSL_PORT_IF_VSR
	PORT_IF_LR2        PortIface = C.OPENNSL_PORT_IF_LR2
	PORT_IF_LRM        PortIface = C.OPENNSL_PORT_IF_LRM
	PORT_IF_XLPPI      PortIface = C.OPENNSL_PORT_IF_XLPPI
	PORT_IF_LBG        PortIface = C.OPENNSL_PORT_IF_LBG
	PORT_IF_CAUI4      PortIface = C.OPENNSL_PORT_IF_CAUI4
	PORT_IF_OAMP       PortIface = C.OPENNSL_PORT_IF_OAMP
	PORT_IF_OLP        PortIface = C.OPENNSL_PORT_IF_OLP
	PORT_IF_ERP        PortIface = C.OPENNSL_PORT_IF_ERP
	PORT_IF_SAT        PortIface = C.OPENNSL_PORT_IF_SAT
	PORT_IF_RCY_MIRROR PortIface = C.OPENNSL_PORT_IF_RCY_MIRROR
	PORT_IF_EVENTOR    PortIface = C.OPENNSL_PORT_IF_EVENTOR
	PORT_IF_COUNT      PortIface = C.OPENNSL_PORT_IF_COUNT
	PORT_IF_10B        PortIface = C.OPENNSL_PORT_IF_10B
)
