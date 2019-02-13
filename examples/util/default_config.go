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

package util

import (
	"github.com/beluganos/go-opennsl/opennsl"
)

func PortDefaultConfig(unit int) error {
	portcfg, err := opennsl.PortConfigGet(unit)
	if err != nil {
		return err
	}

	pbmp, err := portcfg.PBmp(opennsl.PORT_CONFIG_E)
	if err != nil {
		return err
	}

	stg := opennsl.Stg(1)
	err = pbmp.Each(func(port opennsl.Port) error {
		return stg.StpSet(unit, port, opennsl.STG_STP_FORWARD)
	})
	if err != nil {
		return err
	}

	portInfo := opennsl.NewPortInfo()
	portInfo.SetSpeed(0)
	portInfo.SetDuplex(opennsl.PORT_DUPLEX_FULL)
	portInfo.SetPauseRX(opennsl.PORT_ABILITY_PAUSE_RX)
	portInfo.SetPauseTX(opennsl.PORT_ABILITY_PAUSE_TX)
	portInfo.SetLinkscan(opennsl.LINKSCAN_MODE_SW)
	portInfo.SetAutoNeg(false)
	portInfo.SetEnable(true)
	portInfo.SetActionMask(opennsl.NewPortAttr(
		opennsl.PORT_ATTR_AUTONEG_MASK,
		opennsl.PORT_ATTR_DUPLEX_MASK,
		opennsl.PORT_ATTR_PAUSE_TX_MASK,
		opennsl.PORT_ATTR_PAUSE_RX_MASK,
		opennsl.PORT_ATTR_LINKSCAN_MASK,
		opennsl.PORT_ATTR_ENABLE_MASK,
		opennsl.PORT_ATTR_SPEED_MASK,
	))

	return pbmp.Each(func(port opennsl.Port) error {
		return port.SelectiveSet(unit, portInfo)
	})
}

func SwitchDefaultVlanConfig(unit int) error {
	pcfg, err := opennsl.PortConfigGet(unit)
	if err != nil {
		return err
	}

	pbmp, _ := pcfg.PBmp(opennsl.PORT_CONFIG_E)
	if err := opennsl.VLAN_ID_DEFAULT.PortAdd(unit, pbmp, pbmp); err != nil {
		return err
	}

	return pbmp.Each(func(port opennsl.Port) error {
		return port.UntaggedVlanSet(unit, opennsl.VLAN_ID_DEFAULT)
	})
}
