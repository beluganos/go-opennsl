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

package main

import (
	"net"

	"github.com/beluganos/go-opennsl/examples/util"
	"github.com/beluganos/go-opennsl/opennsl"
	"github.com/beluganos/go-opennsl/sal"

	log "github.com/sirupsen/logrus"
)

func newL3FieldGroup(unit int) (opennsl.FieldGroup, error) {
	qset := opennsl.FieldQSet{}
	qset.Init()
	qset.Add(
		opennsl.FieldQualifyEtherType,
		opennsl.FieldQualifySrcIp,
	)

	return opennsl.FieldGroupCreate(unit, &qset, opennsl.FIELD_GROUP_PRIO_ANY)
}

func addDropFieldEntry(unit int, group opennsl.FieldGroup, dst string) error {
	ip, ipnet, err := net.ParseCIDR(dst)
	if err != nil {
		return err
	}

	entry, err := group.EntryCreate(unit)
	if err != nil {
		return err
	}

	entry.Qualify().EtherType(unit, 0x0800, 0xffff)
	entry.Qualify().SrcIp(unit, ip, ipnet.Mask)
	entry.Action().Add(unit, opennsl.FieldActionDropCancel, 0, 0)

	return entry.Install(unit)
}

func ip_block(unit int) error {
	group, err := newL3FieldGroup(unit)
	if err != nil {
		return err
	}

	if err := addDropFieldEntry(unit, group, "192.168.199.36/30"); err != nil {
		return err
	}

	if err := addDropFieldEntry(unit, group, "192.168.199.40/32"); err != nil {
		return err
	}

	return nil
}

func main() {
	log.SetLevel(log.DebugLevel)

	unit := int(0)

	if err := sal.DriverInit(); err != nil {
		log.Errorf("driver_init error. %s", err)
		return
	}
	defer sal.DriverExit()

	if err := util.PortDefaultConfig(unit); err != nil {
		log.Errorf("driver_init error. %s", err)
		return
	}

	if err := util.SwitchDefaultVlanConfig(unit); err != nil {
		log.Errorf("driver_init error. %s", err)
		return
	}

	if err := ip_block(unit); err != nil {
		log.Errorf("driver_init error. %s", err)
		return
	}
}
