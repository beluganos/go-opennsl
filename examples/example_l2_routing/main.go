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

func main() {
	log.SetLevel(log.DebugLevel)

	unit := 0
	mac, _ := net.ParseMAC("00:00:00:00:00:01")
	vlan := opennsl.Vlan(1)
	port := opennsl.Port(53)

	if err := sal.DriverInit(); err != nil {
		log.Errorf("driver_init error. %s", err)
		return
	}
	defer sal.DriverExit()

	log.Infof("DriverInit ok.")

	if err := util.PortDefaultConfig(unit); err != nil {
		log.Errorf("%s", err)
		return
	}

	log.Infof("PortDefaultConfig ok.")

	if err := util.SwitchDefaultVlanConfig(unit); err != nil {
		log.Errorf("driver_init error. %s", err)
		return
	}

	log.Infof("SwitchDefaultVlanConfig ok")

	l2Flags := opennsl.L2_L3LOOKUP

	if _, err := util.AddL2Addr(unit, port, vlan, mac, l2Flags); err != nil {
		log.Errorf("driver_init error. %s", err)
		return
	}

	log.Infof("addL2Entry ok")

	done := make(chan struct{})
	go util.WatchSignal(done)
	<-done
}
