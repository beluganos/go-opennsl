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
#include <opennsl/l2.h>
*/
import "C"

import (
	"fmt"
	"net"
)

//
// L2StationID
//
type L2StationID C.int

func (v L2StationID) C() C.int {
	return C.int(v)
}

const (
	L2_STATION L2StationID = 0
)

//
// L2Station
//
type L2Station C.opennsl_l2_station_t

func (v *L2Station) C() *C.opennsl_l2_station_t {
	return (*C.opennsl_l2_station_t)(v)
}

func (v *L2Station) Flags() L2StationFlags {
	return L2StationFlags(v.flags)
}

func (v *L2Station) SetFlags(flags L2StationFlags) {
	v.flags = flags.C()
}

func (v *L2Station) DstMAC() net.HardwareAddr {
	return ParseMAC(v.dst_mac)
}

func (v *L2Station) SetDstMAC(mac net.HardwareAddr) {
	v.dst_mac = NewMAC(mac)
}

func (v *L2Station) DstMACMask() net.HardwareAddr {
	return ParseMAC(v.dst_mac_mask)
}

func (v *L2Station) SetDstMACMask(mask net.HardwareAddr) {
	v.dst_mac_mask = NewMAC(mask)
}

func (v *L2Station) VID() Vlan {
	return Vlan(v.vlan)
}

func (v *L2Station) VIDMask() Vlan {
	return Vlan(v.vlan_mask)
}

func (v *L2Station) SetVID(vid Vlan) {
	v.vlan = vid.C()
}

func (v *L2Station) SetVIDMask(mask Vlan) {
	v.vlan_mask = mask.C()
}

func (v *L2Station) SrcPort() Port {
	return Port(v.src_port)
}

func (v *L2Station) SrcPortMask() Port {
	return Port(v.src_port_mask)
}

func (v *L2Station) SetSrcPort(port Port) {
	v.src_port = port.C()
}

func (v *L2Station) SetSrcPortMask(mask Port) {
	v.src_port_mask = mask.C()
}

//
// API
//
func NewL2Station() *L2Station {
	station := &L2Station{}
	station.Init()
	return station
}

func L2StationInit(v *L2Station) {
	C.opennsl_l2_station_t_init(v.C())
}

func (v *L2Station) String() string {
	return fmt.Sprintf("L2Station(mac:%s/%s, port:%x/%x, vlan:%x/%x)", v.DstMAC(), v.DstMACMask(), v.SrcPort(), v.SrcPortMask(), v.VID(), v.VIDMask())
}

func (v *L2Station) Init() {
	L2StationInit(v)
}

func L2StationAdd(unit int, stationID L2StationID, v *L2Station) (L2StationID, error) {
	var c_stationID = stationID.C()
	rc := C.opennsl_l2_station_add(C.int(unit), &c_stationID, v.C())
	return L2StationID(c_stationID), ParseError(rc)
}

func (v *L2Station) Add(unit int, stationID L2StationID) (L2StationID, error) {
	return L2StationAdd(unit, stationID, v)
}

func (v L2StationID) Add(unit int, station *L2Station) (L2StationID, error) {
	return L2StationAdd(unit, v, station)
}

func L2StationIDDelete(unit int, v L2StationID) error {
	rc := C.opennsl_l2_station_delete(C.int(unit), v.C())
	return ParseError(rc)
}

func (v L2StationID) Delete(unit int) error {
	return L2StationIDDelete(unit, v)
}

func L2StationGet(unit int, stationID L2StationID) (*L2Station, error) {
	l2st := NewL2Station()
	rc := C.opennsl_l2_station_get(C.int(unit), stationID.C(), l2st.C())
	if err := ParseError(rc); err != nil {
		return nil, err
	}

	return l2st, nil
}

func (v L2StationID) Get(unit int) (*L2Station, error) {
	return L2StationGet(unit, v)
}
