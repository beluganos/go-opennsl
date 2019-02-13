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

import (
	"net"
)

//
// PortInfo
//
type PortInfo C.opennsl_port_info_t

func (v *PortInfo) C() *C.opennsl_port_info_t {
	return (*C.opennsl_port_info_t)(v)
}

func (v *PortInfo) ActionMask() PortAttr {
	return PortAttr(v.action_mask)
}

func (v *PortInfo) SetActionMask(mask PortAttr) {
	v.action_mask = mask.C()
}

func (v *PortInfo) Enable() bool {
	return v.enable != 0
}

func (v *PortInfo) SetEnable(enable bool) {
	c_enable := func() C.int {
		if enable {
			return 1
		}
		return 0
	}()
	v.enable = c_enable
}

func (v *PortInfo) LinkStatus() LinkStatus {
	return LinkStatus(v.linkstatus)
}

func (v *PortInfo) SetLinkStatus(status LinkStatus) {
	v.linkstatus = status.C()
}

func (v *PortInfo) AutoNeg() bool {
	return v.autoneg != 0
}

func (v *PortInfo) SetAutoNeg(enable bool) {
	c_autoneg := func() C.int {
		if enable {
			return 1
		}
		return 0
	}()
	v.autoneg = c_autoneg
}

func (v *PortInfo) Speed() int {
	return int(v.speed)
}

func (v *PortInfo) SetSpeed(speed int) {
	v.speed = C.int(speed)
}

func (v *PortInfo) Duplex() PortDuplex {
	return PortDuplex(v.duplex)
}

func (v *PortInfo) SetDuplex(duplex PortDuplex) {
	v.duplex = C.int(duplex.C())
}

func (v *PortInfo) Linkscan() LinkscanMode {
	return LinkscanMode(v.linkscan)
}

func (v *PortInfo) SetLinkscan(mode LinkscanMode) {
	v.linkscan = C.int(mode.C())
}

func (v *PortInfo) Learn() uint32 {
	return uint32(v.learn)
}

func (v *PortInfo) SetLearn(learn uint32) {
	v.learn = C.uint32(learn)
}

func (v *PortInfo) Discard() int {
	return int(v.discard)
}

func (v *PortInfo) SetDiscard(discard int) {
	v.discard = C.int(discard)
}

func (v *PortInfo) VlanFilter() uint32 {
	return uint32(v.vlanfilter)
}

func (v *PortInfo) SetVlanFilter(filter uint32) {
	v.vlanfilter = C.uint32(filter)
}

func (v *PortInfo) UntaggedVlan() Vlan {
	return Vlan(v.untagged_vlan)
}

func (v *PortInfo) SetUntaggedVlan(vlan Vlan) {
	v.untagged_vlan = vlan.C()
}

func (v *PortInfo) StpState() int {
	return int(v.stp_state)
}

func (v *PortInfo) SetStpState(state int) {
	v.stp_state = C.int(state)
}

func (v *PortInfo) Loopback() PortLoopback {
	return PortLoopback(v.loopback)
}

func (v *PortInfo) SetLoopback(loopback PortLoopback) {
	v.loopback = C.int(loopback.C())
}

func (v *PortInfo) PhyMaster() int {
	return int(v.phy_master)
}

func (v *PortInfo) SetPhyMaster(master int) {
	v.phy_master = C.int(master)
}

func (v *PortInfo) Interface() PortIface {
	return PortIface(v._interface)
}

func (v *PortInfo) SetInterface(iface PortIface) {
	v._interface = iface.C()
}

func (v *PortInfo) Pfm() int {
	return int(v.pfm)
}

func (v *PortInfo) SetPfm(pfm int) {
	v.pfm = C.int(pfm)
}

func (v *PortInfo) PauseTX() PortAbilityMode {
	return PortAbilityMode(v.pause_tx)
}

func (v *PortInfo) SetPauseTX(s PortAbilityMode) {
	v.pause_tx = s.C()
}

func (v *PortInfo) PauseRX() PortAbilityMode {
	return PortAbilityMode(v.pause_rx)
}

func (v *PortInfo) SetPauseRX(s PortAbilityMode) {
	v.pause_rx = s.C()
}

func (v *PortInfo) EncapMode() int {
	return int(v.encap_mode)
}

func (v *PortInfo) SetEncapMode(mode int) {
	v.encap_mode = C.int(mode)
}

func (v *PortInfo) PauseMAC() net.HardwareAddr {
	return ParseMAC(v.pause_mac)
}

func (v *PortInfo) SetPauseMAC(mac net.HardwareAddr) {
	v.pause_mac = NewMAC(mac)
}

func (v *PortInfo) LocalAdvert() PortAbil {
	return PortAbil(v.local_advert)
}

func (v *PortInfo) SetLocalAdvert(abil PortAbil) {
	v.local_advert = abil.C()
}

func (v *PortInfo) LocalAbility() *PortAbility {
	return (*PortAbility)(&v.local_ability)
}

func (v *PortInfo) RemoteAdvertValid() int {
	return int(v.remote_advert_valid)
}

func (v *PortInfo) SetRemoteAdvertValid(ra int) {
	v.remote_advert_valid = C.int(ra)
}

func (v *PortInfo) RemoteAdvert() PortAbil {
	return PortAbil(v.remote_advert)
}

func (v *PortInfo) SetRemoteAdvert(abil PortAbil) {
	v.remote_advert = abil.C()
}

func (v *PortInfo) RemoteAbility() *PortAbility {
	return (*PortAbility)(&v.remote_ability)
}

func (v *PortInfo) McastLimit() (int, int) {
	return int(v.mcast_limit), int(v.mcast_limit_enable)
}

func (v *PortInfo) SetMcastLimit(limit int) {
	v.mcast_limit = C.int(limit)
}

func (v *PortInfo) SetMcastLimitEnable(enable int) {
	v.mcast_limit_enable = C.int(enable)
}

func (v *PortInfo) BcastLimit() (int, int) {
	return int(v.bcast_limit), int(v.bcast_limit_enable)
}

func (v *PortInfo) SetBcastLimit(limit int) {
	v.bcast_limit = C.int(limit)
}

func (v *PortInfo) SetBcastLimitEnable(enable int) {
	v.bcast_limit_enable = C.int(enable)
}

func (v *PortInfo) DlfbcLimit() (int, int) {
	return int(v.dlfbc_limit), int(v.dlfbc_limit_enable)
}

func (v *PortInfo) SetDlfbcLimit(limit int) {
	v.dlfbc_limit = C.int(limit)
}

func (v *PortInfo) SetDlfbcLimitEnable(enable int) {
	v.dlfbc_limit_enable = C.int(enable)
}

func (v *PortInfo) SpeedMax() int {
	return int(v.speed_max)
}

func (v *PortInfo) SetSpeedMax(speed int) {
	v.speed_max = C.int(speed)
}

func (v *PortInfo) Ability() PortAbil {
	return PortAbil(v.ability)
}

func (v *PortInfo) SetAbility(abil PortAbil) {
	v.ability = abil.C()
}

func (v *PortInfo) PortAbility() *PortAbility {
	return (*PortAbility)(&v.port_ability)
}

func (v *PortInfo) FrameMax() int {
	return int(v.frame_max)
}

func (v *PortInfo) SetFrameMax(frame int) {
	v.frame_max = C.int(frame)
}

func (v *PortInfo) Mdix() PortMdix {
	return PortMdix(v.mdix)
}

func (v *PortInfo) SetMdix(mdix PortMdix) {
	v.mdix = mdix.C()
}

func (v *PortInfo) MdixStatus() PortMdixStatus {
	return PortMdixStatus(v.mdix_status)
}

func (v *PortInfo) SetMdixStatus(status PortMdixStatus) {
	v.mdix_status = status.C()
}

func (v *PortInfo) Medium() PortMedium {
	return PortMedium(v.medium)
}

func (v *PortInfo) SetMedium(medium PortMedium) {
	v.medium = medium.C()
}

func (v *PortInfo) Fault() uint32 {
	return uint32(v.fault)
}

func (v *PortInfo) SetFault(fault uint32) {
	v.fault = C.uint32(fault)
}

//
// API
//
func NewPortInfo() *PortInfo {
	info := &PortInfo{}
	info.Init()
	return info
}

func PortInfoInit(v *PortInfo) {
	C.opennsl_port_info_t_init(v.C())
}

func (v *PortInfo) Init() {
	PortInfoInit(v)
}

func PortInfoPortSelectiveGet(unit int, port Port, v *PortInfo) error {
	rc := C.opennsl_port_selective_get(C.int(unit), port.C(), v.C())
	return ParseError(rc)
}

func (v *PortInfo) PortSelectiveGet(unit int, port Port) error {
	return PortInfoPortSelectiveGet(unit, port, v)
}

func PortInfoPortSelectiveSet(unit int, port Port, v *PortInfo) error {
	rc := C.opennsl_port_selective_set(C.int(unit), port.C(), v.C())
	return ParseError(rc)
}

func (v *PortInfo) PortSelectiveSet(unit int, port Port) error {
	return PortInfoPortSelectiveSet(unit, port, v)
}

//
// Port methods
//
func (v Port) SelectiveGet(unit int, portInfo *PortInfo) error {
	return portInfo.PortSelectiveGet(unit, v)
}

func (v Port) SelectiveSet(unit int, portInfo *PortInfo) error {
	return portInfo.PortSelectiveSet(unit, v)
}
