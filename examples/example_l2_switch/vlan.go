// -*- coding: utf-8 -*-

// Copyright (C) 2019 Nippon Telegraph and Telephone Corporation.
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

package main

import (
	"github.com/beluganos/go-opennsl/opennsl"
	log "github.com/sirupsen/logrus"
)

func setVlanTranslation(unit int, port opennsl.Port, enable int) error {
	log.Infof("Set vlan translation. port=%d, enable=%d", port, enable)

	ctrls := []opennsl.VlanControlPort{
		opennsl.VlanTranslateIngressEnable,
		opennsl.VlanTranslateIngressMissDrop,
		opennsl.VlanTranslateEgressEnable,
		opennsl.VlanTranslateEgressMissDrop,
	}

	// Enable VLAN translations for both ingress and egress.
	for _, ctrl := range ctrls {
		if err := ctrl.Set(unit, port, enable); err != nil {
			log.Errorf("VlanTranslate%s  error. port=%d enable=%d %s", ctrl, port, enable, err)
			return err
		}

		log.Infof("VlanTranslate%s ok. port=%d enable=%d", ctrl, port, enable)
	}

	if enable == opennsl.TRUE {
		// Set up port's double tagging mode.
		if err := port.DtagModeSet(unit, opennsl.PORT_DTAG_MODE_INTERNAL); err != nil {
			log.Errorf("PortDtagModeSet port=%d mode=%s error. %s", port, opennsl.PORT_DTAG_MODE_INTERNAL, err)
			return err
		}

		log.Infof("DtagModeSet ok port=%d, mode=%s", port, opennsl.PORT_DTAG_MODE_INTERNAL)
	}

	return nil
}

func setNativeVlan(unit int, vid opennsl.Vlan, pbmp *opennsl.PBmp, strictlyUntagged bool) error {
	log.Infof("Set native vlan. vid=%d pbmp=%s strict=%t", vid, pbmp, strictlyUntagged)

	err := pbmp.Each(func(port opennsl.Port) error {
		if err := port.UntaggedVlanSet(unit, vid); err != nil {
			log.Errorf("PortUntaggedVlanSet error. port=%d vid=%d %s", port, vid, err)
			// Don't exit.  Keep setting the other ports.
		}

		if strictlyUntagged {
			// If strictly untagged option is set, we want to enable VLAN
			// translation & set up miss-drop flags on this port (although
			// we won't need any actual old VID to new VID mapping).  This
			// is to prevent external frames with VID that happened to match
			// the internal VID to be accepted.  In other words, if a port
			// is strictly untagged, only untagged frame will be allowed.
			log.Infof("Set native vlan with strictry untagged enabld. port=%d", port)
			return setVlanTranslation(unit, port, opennsl.TRUE)
		}

		return nil
	})

	return err
}

func clearNativeVlan(unit int, pbmp *opennsl.PBmp, strictlyUntagged bool) {
	// Get the switch's default vid.
	vid, err := opennsl.VlanDefaultGet(unit)
	if err != nil {
		log.Errorf("VlanDefaultGet error. unit=%d %s", unit, err)
		return
	}

	pbmp.Each(func(port opennsl.Port) error {
		if err := port.UntaggedVlanSet(unit, vid); err != nil {
			log.Errorf("PortUntaggedVlanSet error. port=%d vid=%d %s", port, vid, err)
		}

		if strictlyUntagged {
			// Also clear translation settings if this port
			// was strictly untagged.
			log.Infof("Set native vlan with strictry untagged disabld. port=%d", port)
			setVlanTranslation(unit, port, opennsl.FALSE)
		}

		return nil
	})
}

func CreateVlan(unit int, vid opennsl.Vlan) error {
	log.Infof("Create vlan. vid=%d", vid)

	if _, err := vid.Create(unit); err != nil {
		log.Errorf("Create vlan error. vid=%d %s", vid, err)
		return err
	}

	return nil
}

func DestroyVlan(unit int, vid opennsl.Vlan) {
	log.Infof("Destroy vlan. vid=%d", vid)

	if err := vid.Destroy(unit); err != nil {
		log.Errorf("Destroy vlan error. vid=%d %s", vid, err)
	}
}

func addPortsToVlan(unit int, allBmp *opennsl.PBmp, untagBmp *opennsl.PBmp, vid opennsl.Vlan, strictlyUntagged bool) error {

	if untagBmp.IsNotNull() {
		// Update default VLAN ID of the ports if untagged.
		if err := setNativeVlan(unit, vid, untagBmp, strictlyUntagged); err != nil {
			log.Errorf("setNativeVlan error. vid=%d untags=%s %s", vid, untagBmp, err)
			return err
		}
	}

	if allBmp.IsNotNull() {
		// Finally, add ports to VLAN.
		if err := vid.PortAdd(unit, allBmp, untagBmp); err != nil {
			log.Errorf("VlanPortAdd error. vid=%d ports=%s untags=%s", vid, allBmp, untagBmp)
			return err
		}
	}

	return nil
}

func delPortsFromVlan(unit int, allBmp *opennsl.PBmp, untagBmp *opennsl.PBmp, vid opennsl.Vlan, strictlyUntagged bool) {

	if allBmp.IsNotNull() {
		// Remove ports from VLAN.
		if _, err := vid.PortRemove(unit, allBmp); err != nil {
			log.Errorf("VlanPortRemove error. vid=%d ports=%s", vid, allBmp)
		}
	}

	if untagBmp.IsNotNull() {
		// Update default VLAN ID of the ports if untagged.
		clearNativeVlan(unit, untagBmp, strictlyUntagged)
	}
}

func CreateAccessPorts(unit int, pbmp *opennsl.PBmp, vid opennsl.Vlan) error {
	// An ACCESS port carries packets on exactly one VLAN specified
	// in the tag column.  Packets egressing on an access port have
	// no 802.1Q header.
	//
	// Any packet with an 802.1Q header with a nonzero VLAN ID
	// that ingresses on an access port is dropped, regardless of
	// whether the VLAN ID in the header is the access port's
	// VLAN ID.

	return addPortsToVlan(unit, pbmp, pbmp, vid, true)
}

func DestroyAccessPorts(unit int, pbmp *opennsl.PBmp, vid opennsl.Vlan) {
	delPortsFromVlan(unit, pbmp, pbmp, vid, true)
}

func CreateTrunkPorts(unit int, pbmp *opennsl.PBmp, vid opennsl.Vlan) error {
	// A TRUNK port carries packets on one or more specified
	// VLANs specified in the trunks column (often,  on  every
	// VLAN).  A packet that ingresses on a trunk port is in the
	// VLAN specified in its 802.1Q header, or VLAN 0 if the
	// packet has no 802.1Q header.  A packet that egresses
	// through a trunk port will have an 802.1Q header if it has
	// a nonzero VLAN ID.
	//
	// Any packet that ingresses on a trunk port tagged with a
	// VLAN that the port does not trunk is dropped.
	//
	// OpenSwitch NOTE: h/w switches does not support VLAN 0.

	nilPorts := opennsl.NewPBmp()
	return addPortsToVlan(unit, pbmp, nilPorts, vid, false)
}

func DestroyTrunkPorts(unit int, pbmp *opennsl.PBmp, vid opennsl.Vlan) {
	nilPorts := opennsl.NewPBmp()
	delPortsFromVlan(unit, pbmp, nilPorts, vid, false)
}

func CreateNativeTagPorts(unit int, pbmp *opennsl.PBmp, vid opennsl.Vlan) error {
	// A NATIVE-TAGGED port resembles a trunk port, with the
	// exception that a packet without an 802.1Q header that
	// ingresses on a native-tagged port is in the "native
	// VLAN" (specified in the tag column).

	nilPorts := opennsl.NewPBmp()
	if err := addPortsToVlan(unit, pbmp, nilPorts, vid, false); err != nil {
		return err
	}

	return setNativeVlan(unit, vid, pbmp, false)
}

func DestroyNativeTagPorts(unit int, pbmp *opennsl.PBmp, vid opennsl.Vlan) {
	nilPorts := opennsl.NewPBmp()
	delPortsFromVlan(unit, pbmp, nilPorts, vid, false)
	clearNativeVlan(unit, pbmp, false)
}

func CreateNativeUntaggedPorts(unit int, pbmp *opennsl.PBmp, vid opennsl.Vlan) error {
	// A NATIVE-UNTAGGED port resembles a native-tagged port,
	// with the exception that a packet that egresses on a
	// native-untagged port in the native VLAN will not have
	// an 802.1Q header.

	return addPortsToVlan(unit, pbmp, pbmp, vid, false)
}

func DestroyNativeUntaggedPorts(unit int, pbmp *opennsl.PBmp, vid opennsl.Vlan) {
	delPortsFromVlan(unit, pbmp, pbmp, vid, false)
}

type Vlan struct {
	vid         opennsl.Vlan
	maskPorts   *opennsl.PBmp
	accessPorts *opennsl.PBmp
	trunkPorts  *opennsl.PBmp
	nTagPorts   *opennsl.PBmp
	nUntagPorts *opennsl.PBmp
}

func NewVlan(vid opennsl.Vlan) *Vlan {
	return &Vlan{
		vid:         vid,
		maskPorts:   opennsl.NewPBmp(),
		accessPorts: opennsl.NewPBmp(),
		trunkPorts:  opennsl.NewPBmp(),
		nTagPorts:   opennsl.NewPBmp(),
		nUntagPorts: opennsl.NewPBmp(),
	}
}

func (v *Vlan) filterPorts(pbmp *opennsl.PBmp) *opennsl.PBmp {
	newBmp := opennsl.NewPBmp().OR(pbmp)
	if v.maskPorts.IsNotNull() {
		newBmp.AND(v.maskPorts)
	}

	return newBmp
}

func (v *Vlan) SetMaskPorts(pbmp *opennsl.PBmp) {
	v.maskPorts = opennsl.NewPBmp().OR(pbmp)
}

func (v *Vlan) Create(unit int) error {
	return CreateVlan(unit, v.vid)
}

func (v *Vlan) Destroy(unit int) {
	DestroyVlan(unit, v.vid)
}

func (v *Vlan) AddAccessPort(unit int, pbmp *opennsl.PBmp) error {
	tmpBmp := v.filterPorts(pbmp)
	if err := CreateAccessPorts(unit, tmpBmp, v.vid); err != nil {
		log.Errorf("CreateAccessPorts error. vid=%d port=%s %s", v.vid, tmpBmp, err)
		return err
	}

	v.accessPorts.OR(pbmp)
	return nil
}

func (v *Vlan) DeleteAccessPort(unit int, pbmp *opennsl.PBmp) {
	tmpBmp := v.filterPorts(pbmp)
	DestroyAccessPorts(unit, tmpBmp, v.vid)

	v.accessPorts.RemovePorts(pbmp)
}

func (v *Vlan) AddTrunkPort(unit int, pbmp *opennsl.PBmp) error {
	tmpBmp := v.filterPorts(pbmp)
	if err := CreateTrunkPorts(unit, tmpBmp, v.vid); err != nil {
		log.Errorf("CreateTrunkPorts error. vid=%d port=%s %s", v.vid, tmpBmp, err)
		return err
	}

	v.trunkPorts.OR(pbmp)
	return nil
}

func (v *Vlan) DeleteTrunkPort(unit int, pbmp *opennsl.PBmp) {
	tmpBmp := v.filterPorts(pbmp)
	DestroyTrunkPorts(unit, tmpBmp, v.vid)

	v.trunkPorts.RemovePorts(pbmp)
}

func (v *Vlan) AddNativeTagPort(unit int, pbmp *opennsl.PBmp) error {
	tmpBmp := v.filterPorts(pbmp)
	if err := CreateNativeTagPorts(unit, tmpBmp, v.vid); err != nil {
		log.Errorf("CreateNativeTagPorts error. vid=%d port=%s %s", v.vid, tmpBmp, err)
		return err
	}

	v.nTagPorts.OR(pbmp)
	return nil
}

func (v *Vlan) DeleteNativeTagPort(unit int, pbmp *opennsl.PBmp) {
	tmpBmp := v.filterPorts(pbmp)
	DestroyNativeTagPorts(unit, tmpBmp, v.vid)

	v.nTagPorts.RemovePorts(pbmp)
}

func (v *Vlan) AddNativeUntaggedPort(unit int, pbmp *opennsl.PBmp) error {
	tmpBmp := v.filterPorts(pbmp)
	if err := CreateNativeUntaggedPorts(unit, tmpBmp, v.vid); err != nil {
		log.Errorf("CreateNativeUntaggedPorts error. vid=%d port=%s %s", v.vid, tmpBmp, err)
		return err
	}

	v.nUntagPorts.OR(pbmp)
	return nil
}

func (v *Vlan) DeleteNativeUntaggedPort(unit int, pbmp *opennsl.PBmp) {
	tmpBmp := v.filterPorts(pbmp)
	DestroyNativeUntaggedPorts(unit, tmpBmp, v.vid)

	v.nUntagPorts.RemovePorts(pbmp)
}
