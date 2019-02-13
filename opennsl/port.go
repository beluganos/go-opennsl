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
#include "helper.h"
*/
import "C"

import (
	"fmt"
	"net"
	"strings"
)

//
// PortClass
//
type PortClass C.opennsl_port_class_t

func (v PortClass) C() C.opennsl_port_class_t {
	return C.opennsl_port_class_t(v)
}

const (
	PORT_CLASS_NONE                            PortClass = 0
	PORT_CLASS_FIELD_LOOKUP                    PortClass = C.opennslPortClassFieldLookup
	PORT_CLASS_FIELD_INGRESS                   PortClass = C.opennslPortClassFieldIngress
	PORT_CLASS_FIELD_EGRESS                    PortClass = C.opennslPortClassFieldEgress
	PORT_CLASS_ID                              PortClass = C.opennslPortClassId
	PORT_CLASS_FIELD_INGRESS_PACKET_PROCESSING PortClass = C.opennslPortClassFieldIngressPacketProcessing
	PORT_CLASS_FIELD_EGRESS_PACKET_PROCESSING  PortClass = C.opennslPortClassFieldEgressPacketProcessing
	PORT_CLASS_EGRESS                          PortClass = C.opennslPortClassEgress
)

var portClass_names = map[PortClass]string{
	PORT_CLASS_FIELD_LOOKUP:                    "FIELD_LOOKUP",
	PORT_CLASS_FIELD_INGRESS:                   "FIELD_INGRESS",
	PORT_CLASS_FIELD_EGRESS:                    "FIELD_EGRESS",
	PORT_CLASS_ID:                              "ID",
	PORT_CLASS_FIELD_INGRESS_PACKET_PROCESSING: "FIELD_INGRESS_PACKET_PROCESSING",
	PORT_CLASS_FIELD_EGRESS_PACKET_PROCESSING:  "FIELD_EGRESS_PACKET_PROCESSING",
	PORT_CLASS_EGRESS:                          "EGRESS",
}

var portClass_values = map[string]PortClass{
	"FIELD_LOOKUP":                    PORT_CLASS_FIELD_LOOKUP,
	"FIELD_INGRESS":                   PORT_CLASS_FIELD_INGRESS,
	"FIELD_EGRESS":                    PORT_CLASS_FIELD_EGRESS,
	"ID":                              PORT_CLASS_ID,
	"FIELD_INGRESS_PACKET_PROCESSING": PORT_CLASS_FIELD_INGRESS_PACKET_PROCESSING,
	"FIELD_EGRESS_PACKET_PROCESSING":  PORT_CLASS_FIELD_EGRESS_PACKET_PROCESSING,
	"EGRESS":                          PORT_CLASS_EGRESS,
}

func (v PortClass) String() string {
	if s, ok := portClass_names[v]; ok {
		return s
	}
	return fmt.Sprintf("PortClass(%d)", v)
}

func ParseColorClass(s string) (PortClass, error) {
	if v, ok := portClass_values[s]; ok {
		return v, nil
	}
	return PORT_CLASS_NONE, fmt.Errorf("Invalid PortClass. %s", s)
}

//
// PortDtagMode
//
type PortDtagMode C.int

func (v PortDtagMode) C() C.int {
	return C.int(v)
}

const (
	PORT_DTAG_MODE_NONE     PortDtagMode = C.OPENNSL_PORT_DTAG_MODE_NONE
	PORT_DTAG_MODE_INTERNAL PortDtagMode = C.OPENNSL_PORT_DTAG_MODE_INTERNAL
	PORT_DTAG_MODE_EXTERNAL PortDtagMode = C.OPENNSL_PORT_DTAG_MODE_EXTERNAL
)

var portDtagMode_names = map[PortDtagMode]string{
	PORT_DTAG_MODE_NONE:     "NONE",
	PORT_DTAG_MODE_INTERNAL: "INTERNAL",
	PORT_DTAG_MODE_EXTERNAL: "EXTERNAL",
}

var portDtagMode_values = map[string]PortDtagMode{
	"NONE":     PORT_DTAG_MODE_NONE,
	"INTERNAL": PORT_DTAG_MODE_INTERNAL,
	"EXTERNAL": PORT_DTAG_MODE_EXTERNAL,
}

func (v PortDtagMode) String() string {
	if s, ok := portDtagMode_names[v]; ok {
		return s
	}
	return fmt.Sprintf("PortDtagMode(%d)", v)
}

func ParsePortDtagMode(s string) (PortDtagMode, error) {
	if v, ok := portDtagMode_values[s]; ok {
		return v, nil
	}
	return PORT_DTAG_MODE_NONE, fmt.Errorf("Invalid PortDtagMode. %s", s)
}

//
// PortFloodBlock
//
type PortFloodBlock C.uint32

func (v PortFloodBlock) C() C.uint32 {
	return C.uint32(v)
}

func NewPortFloodBlock(flags ...PortFloodBlock) PortFloodBlock {
	v := PORT_FLOOD_BLOCK_NONE
	for _, flag := range flags {
		v |= flag
	}
	return v
}

const (
	PORT_FLOOD_BLOCK_NONE                PortFloodBlock = 0
	PORT_FLOOD_BLOCK_BCAST               PortFloodBlock = C.OPENNSL_PORT_FLOOD_BLOCK_BCAST
	PORT_FLOOD_BLOCK_UNKNOWN_UCAST       PortFloodBlock = C.OPENNSL_PORT_FLOOD_BLOCK_UNKNOWN_UCAST
	PORT_FLOOD_BLOCK_UNKNOWN_MCAST       PortFloodBlock = C.OPENNSL_PORT_FLOOD_BLOCK_UNKNOWN_MCAST
	PORT_FLOOD_BLOCK_ALL                 PortFloodBlock = C.OPENNSL_PORT_FLOOD_BLOCK_ALL
	PORT_FLOOD_BLOCK_UNKNOWN_IP_MCAST    PortFloodBlock = C.OPENNSL_PORT_FLOOD_BLOCK_UNKNOWN_IP_MCAST
	PORT_FLOOD_BLOCK_UNKNOWN_NONIP_MCAST PortFloodBlock = C.OPENNSL_PORT_FLOOD_BLOCK_UNKNOWN_NONIP_MCAST
	PORT_FLOOD_BLOCK_KNOWN_MCAST         PortFloodBlock = C.OPENNSL_PORT_FLOOD_BLOCK_KNOWN_MCAST
)

var portFloodBlock_names = map[PortFloodBlock]string{
	PORT_FLOOD_BLOCK_BCAST:               "BCAST",
	PORT_FLOOD_BLOCK_UNKNOWN_UCAST:       "UNKNOWN_UCAST",
	PORT_FLOOD_BLOCK_UNKNOWN_MCAST:       "UNKNOWN_MCAST",
	PORT_FLOOD_BLOCK_ALL:                 "ALL",
	PORT_FLOOD_BLOCK_UNKNOWN_IP_MCAST:    "IP_MCAST",
	PORT_FLOOD_BLOCK_UNKNOWN_NONIP_MCAST: "UNKNOWN_NONIP_MCAST",
	PORT_FLOOD_BLOCK_KNOWN_MCAST:         "KNOWN_MCAST",
}

var portFloodBlock_values = map[string]PortFloodBlock{
	"BCAST":               PORT_FLOOD_BLOCK_BCAST,
	"UNKNOWN_UCAST":       PORT_FLOOD_BLOCK_UNKNOWN_UCAST,
	"UNKNOWN_MCAST":       PORT_FLOOD_BLOCK_UNKNOWN_MCAST,
	"ALL":                 PORT_FLOOD_BLOCK_ALL,
	"IP_MCAST":            PORT_FLOOD_BLOCK_UNKNOWN_IP_MCAST,
	"UNKNOWN_NONIP_MCAST": PORT_FLOOD_BLOCK_UNKNOWN_NONIP_MCAST,
	"KNOWN_MCAST":         PORT_FLOOD_BLOCK_KNOWN_MCAST,
}

func (v PortFloodBlock) String() string {
	names := []string{}
	for val, name := range portFloodBlock_names {
		if v&val != 0 {
			names = append(names, name)
		}
	}
	return strings.Join(names, "|")
}

func ParsePortFloodBlock(s string) (PortFloodBlock, error) {
	if v, ok := portFloodBlock_values[s]; ok {
		return v, nil
	}
	return PORT_FLOOD_BLOCK_NONE, fmt.Errorf("Invalid PortFloodBlock. %s", s)
}

//
// PortDuplex
//
type PortDuplex C.opennsl_port_duplex_t

func (v PortDuplex) C() C.opennsl_port_duplex_t {
	return C.opennsl_port_duplex_t(v)
}

const (
	PORT_DUPLEX_NONE PortDuplex = 0
	PORT_DUPLEX_HALF PortDuplex = C.OPENNSL_PORT_DUPLEX_HALF
	PORT_DUPLEX_FULL PortDuplex = C.OPENNSL_PORT_DUPLEX_FULL
)

const PORT_DUPLEX_COUNT = C.OPENNSL_PORT_DUPLEX_COUNT

var portDuplex_names = map[PortDuplex]string{
	PORT_DUPLEX_HALF: "HALF",
	PORT_DUPLEX_FULL: "FULL",
}

var portDuplex_values = map[string]PortDuplex{
	"HALF": PORT_DUPLEX_HALF,
	"FULL": PORT_DUPLEX_FULL,
}

func (v PortDuplex) String() string {
	if s, ok := portDuplex_names[v]; ok {
		return s
	}
	return fmt.Sprintf("PortDuplex(%d)", v)
}

func ParsePortDuplex(s string) (PortDuplex, error) {
	if v, ok := portDuplex_values[s]; ok {
		return v, nil
	}
	return PORT_DUPLEX_NONE, fmt.Errorf("Invalid PortDuplex. %s", s)
}

//
// PortLearn
//
type PortLearn C.uint32

func (v PortLearn) C() C.uint32 {
	return C.uint32(v)
}

func NewPortLearn(flags ...PortLearn) PortLearn {
	v := PORT_LEARN_NONE
	for _, flag := range flags {
		v |= flag
	}
	return v
}

const (
	PORT_LEARN_NONE    PortLearn = 0
	PORT_LEARN_ARL     PortLearn = C.OPENNSL_PORT_LEARN_ARL
	PORT_LEARN_CPU     PortLearn = C.OPENNSL_PORT_LEARN_CPU
	PORT_LEARN_FWD     PortLearn = C.OPENNSL_PORT_LEARN_FWD
	PORT_LEARN_PENDING PortLearn = C.OPENNSL_PORT_LEARN_PENDING
)

var portLearn_names = map[PortLearn]string{
	PORT_LEARN_ARL:     "ARL",
	PORT_LEARN_CPU:     "CPU",
	PORT_LEARN_FWD:     "FWD",
	PORT_LEARN_PENDING: "PENDING",
}

var portLearn_values = map[string]PortLearn{
	"ARL":     PORT_LEARN_ARL,
	"CPU":     PORT_LEARN_CPU,
	"FWD":     PORT_LEARN_FWD,
	"PENDING": PORT_LEARN_PENDING,
}

func (v PortLearn) String() string {
	names := []string{}
	for val, name := range portLearn_names {
		if v&val != 0 {
			names = append(names, name)
		}
	}
	return strings.Join(names, "|")
}

func ParsePortLearn(s string) (PortLearn, error) {
	if v, ok := portLearn_values[s]; ok {
		return v, nil
	}
	return PORT_LEARN_NONE, fmt.Errorf("Invalid PortLearn. %s", s)
}

//
// PortMdix
//
type PortMdix C.opennsl_port_mdix_t

func (v PortMdix) C() C.opennsl_port_mdix_t {
	return C.opennsl_port_mdix_t(v)
}

const (
	PORT_MDIX_AUTO       PortMdix = C.OPENNSL_PORT_MDIX_AUTO
	PORT_MDIX_FORCE_AUTO PortMdix = C.OPENNSL_PORT_MDIX_FORCE_AUTO
	PORT_MDIX_NORMAL     PortMdix = C.OPENNSL_PORT_MDIX_NORMAL
	PORT_MDIX_XOVER      PortMdix = C.OPENNSL_PORT_MDIX_XOVER
)

const PORT_MDIX_COUNT = C.OPENNSL_PORT_MDIX_COUNT

var portMedix_names = map[PortMdix]string{
	PORT_MDIX_AUTO:       "AUTO",
	PORT_MDIX_FORCE_AUTO: "FORCE_AUTO",
	PORT_MDIX_NORMAL:     "NORMAL",
	PORT_MDIX_XOVER:      "XOVER",
}

var portMedix_values = map[string]PortMdix{
	"AUTO":       PORT_MDIX_AUTO,
	"FORCE_AUTO": PORT_MDIX_FORCE_AUTO,
	"NORMAL":     PORT_MDIX_NORMAL,
	"XOVER":      PORT_MDIX_XOVER,
}

func (v PortMdix) String() string {
	if s, ok := portMedix_names[v]; ok {
		return s
	}
	return fmt.Sprintf("PortMdix(%d)", v)
}

func ParsePortMdix(s string) (PortMdix, error) {
	if v, ok := portMedix_values[s]; ok {
		return v, nil
	}
	return PORT_MDIX_AUTO, fmt.Errorf("Invalid PortMdix. %s", s)
}

//
// PortMdixStatus
//
type PortMdixStatus C.opennsl_port_mdix_status_t

func (v PortMdixStatus) C() C.opennsl_port_mdix_status_t {
	return C.opennsl_port_mdix_status_t(v)
}

const (
	PORT_MDIX_STATUS_NORMAL PortMdixStatus = C.OPENNSL_PORT_MDIX_STATUS_NORMAL
	PORT_MDIX_STATUS_XOVER  PortMdixStatus = C.OPENNSL_PORT_MDIX_STATUS_XOVER
)

const PORT_MDIX_STATUS_COUNT = C.OPENNSL_PORT_MDIX_STATUS_COUNT

var portMdixStatus_names = map[PortMdixStatus]string{
	PORT_MDIX_STATUS_NORMAL: "NORMAL",
	PORT_MDIX_STATUS_XOVER:  "XOVER",
}

var portMdixStatus_values = map[string]PortMdixStatus{
	"NORMAL": PORT_MDIX_STATUS_NORMAL,
	"XOVER":  PORT_MDIX_STATUS_XOVER,
}

func (v PortMdixStatus) String() string {
	if s, ok := portMdixStatus_names[v]; ok {
		return s
	}
	return fmt.Sprintf("PortMdixStatus(%d)", v)
}

func ParsePortMdixStatus(s string) (PortMdixStatus, error) {
	if v, ok := portMdixStatus_values[s]; ok {
		return v, nil
	}
	return PORT_MDIX_STATUS_NORMAL, fmt.Errorf("Invalid PortMdixStatus. %s", s)
}

//
// PortMedium
//
type PortMedium C.opennsl_port_medium_t

func (v PortMedium) C() C.opennsl_port_medium_t {
	return C.opennsl_port_medium_t(v)
}

const (
	PORT_MEDIUM_NONE   PortMedium = C.OPENNSL_PORT_MEDIUM_NONE
	PORT_MEDIUM_COPPER PortMedium = C.OPENNSL_PORT_MEDIUM_COPPER
	PORT_MEDIUM_FIBER  PortMedium = C.OPENNSL_PORT_MEDIUM_FIBER
)

const PORT_MEDIUM_COUNT = C.OPENNSL_PORT_MEDIUM_COUNT

var portMedium_names = map[PortMedium]string{
	PORT_MEDIUM_NONE:   "NONE",
	PORT_MEDIUM_COPPER: "COPPER",
	PORT_MEDIUM_FIBER:  "FIBER",
}

var portMedium_values = map[string]PortMedium{
	"COPPER": PORT_MEDIUM_COPPER,
	"FIBER":  PORT_MEDIUM_FIBER,
}

func (v PortMedium) String() string {
	if s, ok := portMedium_names[v]; ok {
		return s
	}
	return fmt.Sprintf("PortMedium(%d)", v)
}

func ParsePortMedium(s string) (PortMedium, error) {
	if v, ok := portMedium_values[s]; ok {
		return v, nil
	}
	return PORT_MEDIUM_NONE, fmt.Errorf("Invalid PortMedium. %s", s)
}

//
// PortLoopback
//
type PortLoopback C.opennsl_port_loopback_t

func (v PortLoopback) C() C.opennsl_port_loopback_t {
	return C.opennsl_port_loopback_t(v)
}

const (
	OPENNSL_PORT_LOOPBACK_NONE       PortLoopback = C.OPENNSL_PORT_LOOPBACK_NONE
	OPENNSL_PORT_LOOPBACK_MAC        PortLoopback = C.OPENNSL_PORT_LOOPBACK_MAC
	OPENNSL_PORT_LOOPBACK_PHY        PortLoopback = C.OPENNSL_PORT_LOOPBACK_PHY
	OPENNSL_PORT_LOOPBACK_PHY_REMOTE PortLoopback = C.OPENNSL_PORT_LOOPBACK_PHY_REMOTE
	OPENNSL_PORT_LOOPBACK_MAC_REMOTE PortLoopback = C.OPENNSL_PORT_LOOPBACK_MAC_REMOTE
)

const OPENNSL_PORT_LOOPBACK_COUNT = C.OPENNSL_PORT_LOOPBACK_COUNT

var portLoopback_names = map[PortLoopback]string{
	OPENNSL_PORT_LOOPBACK_NONE:       "NONE",
	OPENNSL_PORT_LOOPBACK_MAC:        "MAC",
	OPENNSL_PORT_LOOPBACK_PHY:        "PHY",
	OPENNSL_PORT_LOOPBACK_PHY_REMOTE: "PHY_REMOTE",
	OPENNSL_PORT_LOOPBACK_MAC_REMOTE: "MAC_REMOTE",
}

var portLoopback_values = map[string]PortLoopback{
	"NONE":       OPENNSL_PORT_LOOPBACK_NONE,
	"MAC":        OPENNSL_PORT_LOOPBACK_MAC,
	"PHY":        OPENNSL_PORT_LOOPBACK_PHY,
	"PHY_REMOTE": OPENNSL_PORT_LOOPBACK_PHY_REMOTE,
	"MAC_REMOTE": OPENNSL_PORT_LOOPBACK_MAC_REMOTE,
}

func (v PortLoopback) String() string {
	if s, ok := portLoopback_names[v]; ok {
		return s
	}
	return fmt.Sprintf("PortLoopback(%d)", v)
}

func ParsePortLoopback(s string) (PortLoopback, error) {
	if v, ok := portLoopback_values[s]; ok {
		return v, nil
	}
	return OPENNSL_PORT_LOOPBACK_NONE, fmt.Errorf("Invalid PortLoopback. %s", s)
}

//
// PortEnable
//
type PortEnable C.int

func (v PortEnable) C() C.int {
	return C.int(v)
}

func (v PortEnable) Flip() PortEnable {
	if v == PORT_ENABLE_FALSE {
		return PORT_ENABLE_TRUE
	}
	return PORT_ENABLE_FALSE
}

const (
	PORT_ENABLE_FALSE PortEnable = 0
	PORT_ENABLE_TRUE  PortEnable = 1
)

var portEnable_names = map[PortEnable]string{
	PORT_ENABLE_FALSE: "FALSE",
	PORT_ENABLE_TRUE:  "TRUE",
}

var portEnable_values = map[string]PortEnable{
	"FALSE": PORT_ENABLE_FALSE,
	"TRUE":  PORT_ENABLE_TRUE,
}

func (v PortEnable) String() string {
	if s, ok := portEnable_names[v]; ok {
		return s
	}
	return fmt.Sprintf("PortEnable(%d)", v)
}

func ParsePortEnable(s string) (PortEnable, error) {
	if v, ok := portEnable_values[s]; ok {
		return v, nil
	}
	return PORT_ENABLE_FALSE, fmt.Errorf("Invalid PortEnable. %s", s)
}

//
// Port IFilter mode
//
type PortIFilterMode int

func (v PortIFilterMode) C() C.int {
	return C.int(v)
}

const (
	PORT_IFILTER_OFF PortIFilterMode = 0
	PORT_IFILTER_ON  PortIFilterMode = 1
)

var portIFilterMode_names = map[PortIFilterMode]string{
	PORT_IFILTER_OFF: "OFF",
	PORT_IFILTER_ON:  "ON",
}

var portIFilterMode_values = map[string]PortIFilterMode{
	"OFF": PORT_IFILTER_OFF,
	"ON":  PORT_IFILTER_ON,
}

func (v PortIFilterMode) String() string {
	if s, ok := portIFilterMode_names[v]; ok {
		return s
	}
	return fmt.Sprintf("PortIFilterMode(%d)", v)
}

func ParsePortIFilterMode(s string) (PortIFilterMode, error) {
	if v, ok := portIFilterMode_values[s]; ok {
		return v, nil
	}
	return PORT_IFILTER_OFF, fmt.Errorf("Invalid PortIFilterMode. %s", s)
}

//
// Port methods
//
func PortInit(unit int) error {
	rc := C.opennsl_port_init(C.int(unit))
	return ParseError(rc)
}

func PortClear(unit int) error {
	rc := C.opennsl_port_clear(C.int(unit))
	return ParseError(rc)
}

func PortAutonegGet(unit int, port Port) (bool, error) {
	c_autoneg := C.int(0)

	rc := C.opennsl_port_autoneg_get(C.int(unit), port.C(), &c_autoneg)
	if err := ParseError(rc); err != nil {
		return false, err
	}

	return (c_autoneg != 0), nil
}

func (v Port) AutonegGet(unit int) (bool, error) {
	return PortAutonegGet(unit, v)
}

func PortAutonegSet(unit int, port Port, autoneg bool) error {
	c_autoneg := func() C.int {
		if autoneg {
			return 1
		}
		return 0
	}()

	rc := C.opennsl_port_autoneg_set(C.int(unit), port.C(), c_autoneg)
	return ParseError(rc)
}

func (v Port) AutonegSet(unit int, autoneg bool) error {
	return PortAutonegSet(unit, v, autoneg)
}

func PortCfiColorGet(unit int, port Port, priority int) (Color, error) {
	c_color := COLOR_NONE.C()

	rc := C.opennsl_port_cfi_color_get(C.int(unit), port.C(), C.int(priority), &c_color)
	return Color(c_color), ParseError(rc)
}

func (v Port) CfiColorGet(unit int, priority int) (Color, error) {
	return PortCfiColorGet(unit, v, priority)
}

func PortCfiColorSet(unit int, port Port, priority int, color Color) error {
	rc := C.opennsl_port_cfi_color_set(C.int(unit), port.C(), C.int(priority), color.C())
	return ParseError(rc)
}

func (v Port) CfiColorSet(unit int, priority int, color Color) error {
	return PortCfiColorSet(unit, v, priority, color)
}

func PortClassGet(unit int, port Port, portClass PortClass) (uint32, error) {
	c_id := C.uint32(0)

	rc := C.opennsl_port_class_get(C.int(unit), port.C(), portClass.C(), &c_id)
	return uint32(c_id), ParseError(rc)
}

func (v Port) ClassGet(unit int, portClass PortClass) (uint32, error) {
	return PortClassGet(unit, v, portClass)
}

func PortClassSet(unit int, port Port, portClass PortClass, id uint32) error {
	rc := C.opennsl_port_class_set(C.int(unit), port.C(), portClass.C(), C.uint32(id))
	return ParseError(rc)
}

func (v Port) ClassSet(unit int, portClass PortClass, id uint32) error {
	return PortClassSet(unit, v, portClass, id)
}

func PortDtagModeGet(unit int, port Port) (PortDtagMode, error) {
	c_mode := PORT_DTAG_MODE_NONE.C()

	rc := C.opennsl_port_dtag_mode_get(C.int(unit), port.C(), &c_mode)
	return PortDtagMode(c_mode), ParseError(rc)
}

func (v Port) DtagModeGet(unit int) (PortDtagMode, error) {
	return PortDtagModeGet(unit, v)
}

func PortDtagModeSet(unit int, port Port, mode PortDtagMode) error {
	rc := C.opennsl_port_dtag_mode_set(C.int(unit), port.C(), mode.C())
	return ParseError(rc)
}

func (v Port) DtagModeSet(unit int, mode PortDtagMode) error {
	return PortDtagModeSet(unit, v, mode)
}

func PortDuplexGet(unit int, port Port) (PortDuplex, error) {
	c_duplex := C.int(0)

	rc := C.opennsl_port_duplex_get(C.int(unit), port.C(), &c_duplex)
	if err := ParseError(rc); err != nil {
		return PORT_DUPLEX_NONE, err
	}

	return PortDuplex(c_duplex), nil
}

func (v Port) DuplexGet(unit int) (PortDuplex, error) {
	return PortDuplexGet(unit, v)
}

func PortDuplexSet(unit int, port Port, duplex PortDuplex) error {
	c_duplex := C.int(duplex.C())

	rc := C.opennsl_port_duplex_set(C.int(unit), port.C(), c_duplex)
	return ParseError(rc)
}

func (v Port) DuplexSet(unit int, duplex PortDuplex) error {
	return PortDuplexSet(unit, v, duplex)
}

func PortEnableGet(unit int, v Port) (PortEnable, error) {
	c_enable := C.int(0)

	rc := C.opennsl_port_enable_get(C.int(unit), v.C(), &c_enable)
	return PortEnable(c_enable), ParseError(rc)
}

func (v Port) EnableGet(unit int) (PortEnable, error) {
	return PortEnableGet(unit, v)
}

func PortEnableSet(unit int, v Port, enable PortEnable) error {
	rc := C.opennsl_port_enable_set(C.int(unit), v.C(), enable.C())
	return ParseError(rc)
}

func (v Port) EnableSet(unit int, enable PortEnable) error {
	return PortEnableSet(unit, v, enable)
}

func PortFloodBlockGet(unit int, ingrPort Port, egrPort Port) (PortFloodBlock, error) {
	c_flags := C.uint32(0)

	rc := C.opennsl_port_flood_block_get(C.int(unit), ingrPort.C(), egrPort.C(), &c_flags)
	if err := ParseError(rc); err != nil {
		return PORT_FLOOD_BLOCK_NONE, err
	}

	return PortFloodBlock(c_flags), nil
}

func PortFloodBlockSet(unit int, ingrPort Port, egrPort Port, flags PortFloodBlock) error {
	c_flags := C.uint32(flags.C())

	rc := C.opennsl_port_flood_block_set(C.int(unit), ingrPort.C(), egrPort.C(), c_flags)
	return ParseError(rc)
}

func PortGPortGet(unit int, v Port) (GPort, error) {
	c_gport := C.opennsl_gport_t(0)

	rc := C.opennsl_port_gport_get(C.int(unit), v.C(), &c_gport)
	return GPort(c_gport), ParseError(rc)
}

func PortFrameMaxSet(unit int, port Port, size int) error {
	rc := C.opennsl_port_frame_max_set(C.int(unit), port.C(), C.int(size))
	return ParseError(rc)
}

func (v Port) FrameMaxSet(unit int, size int) error {
	return PortFrameMaxSet(unit, v, size)
}

func PortFrameMaxGet(unit int, port Port) (int, error) {
	c_size := C.int(0)

	rc := C.opennsl_port_frame_max_get(C.int(unit), port.C(), &c_size)
	if err := ParseError(rc); err != nil {
		return 0, err
	}

	return int(c_size), nil
}

func (v Port) FrameMaxGet(unit int) (int, error) {
	return PortFrameMaxGet(unit, v)
}

func (v Port) GPortGet(unit int) (GPort, error) {
	return PortGPortGet(unit, v)
}

func PortL3MTUGet(unit int, port Port) (int, error) {
	c_mtu := C.int(0)

	rc := C.opennsl_port_l3_mtu_get(C.int(unit), port.C(), &c_mtu)
	return int(c_mtu), ParseError(rc)
}

func (v Port) L3MTUGet(unit int) (int, error) {
	return PortL3MTUGet(unit, v)
}

func PortL3MTUSet(unit int, port Port, mtu int) error {
	rc := C.opennsl_port_l3_mtu_set(C.int(unit), port.C(), C.int(mtu))
	return ParseError(rc)
}

func PortLearnGet(unit int, port Port) (PortLearn, error) {
	c_flags := PORT_LEARN_NONE.C()

	rc := C.opennsl_port_learn_get(C.int(unit), port.C(), &c_flags)
	return PortLearn(c_flags), ParseError(rc)
}

func (v Port) LearnGet(unit int) (PortLearn, error) {
	return PortLearnGet(unit, v)
}

func PortLearnSet(unit int, port Port, flags PortLearn) error {
	rc := C.opennsl_port_learn_set(C.int(unit), port.C(), flags.C())
	return ParseError(rc)
}

func (v Port) LearnSet(unit int, flags PortLearn) error {
	return PortLearnSet(unit, v, flags)
}

func PortLinkFailedClear(unit int, port Port) error {
	rc := C.opennsl_port_link_failed_clear(C.int(unit), port.C())
	return ParseError(rc)
}

func (v Port) LinkFailedClear(unit int) error {
	return PortLinkFailedClear(unit, v)
}

func PortLinkStatusGet(unit int, port Port) (int, error) {
	c_status := C.int(0)
	rc := C.opennsl_port_link_status_get(C.int(unit), port.C(), &c_status)
	return int(c_status), ParseError(rc)
}

func (v Port) LinkStatusGet(unit int) (int, error) {
	return PortLinkStatusGet(unit, v)
}

func PortLinkscanGet(unit int, port Port) (LinkscanMode, error) {
	mode := C.int(0)

	rc := C.opennsl_port_linkscan_get(C.int(unit), port.C(), &mode)
	return LinkscanMode(mode), ParseError(rc)
}

func (v Port) LinkscanGet(unit int) (LinkscanMode, error) {
	return PortLinkscanGet(unit, v)
}

func PortLinkscanSet(unit int, port Port, mode LinkscanMode) error {
	rc := C.opennsl_port_linkscan_set(C.int(unit), port.C(), C.int(mode))
	return ParseError(rc)
}

func (v Port) LinkscanSet(unit int, mode LinkscanMode) error {
	return PortLinkscanSet(unit, v, mode)
}

func PortPauseAddrGet(unit int, port Port) (net.HardwareAddr, error) {
	var c_mac C.opennsl_mac_t

	rc := C.opennsl_port_pause_addr_get(C.int(unit), port.C(), &c_mac[0])
	if err := ParseError(rc); err != nil {
		return net.HardwareAddr{}, err
	}

	return ParseMAC(c_mac), nil
}

func (v Port) PauseAddrGet(unit int) (net.HardwareAddr, error) {
	return PortPauseAddrGet(unit, v)
}

func PortPauseAddrSet(unit int, port Port, hwaddr net.HardwareAddr) error {
	c_mac := NewMAC(hwaddr)

	rc := C.opennsl_port_pause_addr_set(C.int(unit), port.C(), &c_mac[0])
	return ParseError(rc)
}

func (v Port) PauseAddrSet(unit int, hwaddr net.HardwareAddr) error {
	return PortPauseAddrSet(unit, v, hwaddr)
}

func PortModPortGet(unit int, v Port, module Module) GPort {
	c_gport := C._opennsl_gport_from_modid_and_port(v.C(), module.C())
	return GPort(c_gport)
}

func (v Port) ModPortGet(unit int, module Module) GPort {
	return PortModPortGet(unit, v, module)
}

func PortSpeedGet(unit int, port Port) (int, error) {
	c_speed := C.int(0)

	rc := C.opennsl_port_speed_get(C.int(unit), port.C(), &c_speed)
	return int(c_speed), ParseError(rc)
}

func (v Port) SpeedGet(unit int) (int, error) {
	return PortSpeedGet(unit, v)
}

func PortSpeedMax(unit int, port Port) (int, error) {
	c_speed := C.int(0)

	rc := C.opennsl_port_speed_max(C.int(unit), port.C(), &c_speed)
	return int(c_speed), ParseError(rc)
}

func (v Port) SpeedMax(unit int) (int, error) {
	return PortSpeedMax(unit, v)
}

func PortSpeedSet(unit int, port Port, speed int) error {
	rc := C.opennsl_port_speed_set(C.int(unit), port.C(), C.int(speed))
	return ParseError(rc)
}

func (v Port) SpeedSet(unit int, speed int) error {
	return PortSpeedSet(unit, v, speed)
}

func PortUntaggedVlanSet(unit int, port Port, vid Vlan) error {
	rc := C.opennsl_port_untagged_vlan_set(C.int(unit), port.C(), vid.C())
	return ParseError(rc)
}

func (v Port) UntaggedVlanSet(unit int, vid Vlan) error {
	return PortUntaggedVlanSet(unit, v, vid)
}

func PortUntaggedVlanGet(unit int, port Port) (Vlan, error) {
	c_vid := C.opennsl_vlan_t(0)

	rc := C.opennsl_port_untagged_vlan_get(C.int(unit), port.C(), &c_vid)
	return Vlan(c_vid), ParseError(rc)
}

func (v Port) UntaggedVlanGet(unit int) (Vlan, error) {
	return PortUntaggedVlanGet(unit, v)
}

func PortIFilterSet(unit int, port Port, mode PortIFilterMode) error {
	rc := C.opennsl_port_ifilter_set(C.int(unit), port.C(), mode.C())
	return ParseError(rc)
}

func (v Port) IFilterSet(unit int, mode PortIFilterMode) error {
	return PortIFilterSet(unit, v, mode)
}

//
// GPort methods
//
func (v GPort) LocalPortGet(unit int) (Port, error) {
	c_port := C.opennsl_port_t(0)

	rc := C.opennsl_port_local_get(C.int(unit), v.C(), &c_port)
	return Port(c_port), ParseError(rc)
}

func (v GPort) StatEnableSet(unit int, enable int) error {
	rc := C.opennsl_port_stat_enable_set(C.int(unit), v.C(), C.int(enable))
	return ParseError(rc)
}
