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
	"os"
	"os/signal"

	"github.com/beluganos/go-opennsl/examples/util"
	"github.com/beluganos/go-opennsl/opennsl"
	"github.com/beluganos/go-opennsl/sal"

	log "github.com/sirupsen/logrus"
)

func watchSignal(done chan struct{}) {

	ch := make(chan os.Signal, 1)
	defer close(ch)

	signal.Notify(ch, os.Interrupt)
	<-ch

	log.Infof("Interrupt signal.")

	close(done)
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
		log.Errorf("PortDefaultConfig error. %s", err)
		return
	}

	if err := util.SwitchDefaultVlanConfig(unit); err != nil {
		log.Errorf("SwitchDefaultVlanConfig error. %s", err)
		return
	}

	portcfg, err := opennsl.PortConfigGet(unit)
	if err != nil {
		log.Errorf("PortConfigGet error. %s", err)
		return
	}

	pbmp, err := portcfg.PBmp(opennsl.PORT_CONFIG_E)
	if err != nil {
		log.Errorf("portcfg.PBmp(E) error. %s", err)
		return
	}

	done := make(chan struct{})
	go watchSignal(done)

	go txSetup(unit, pbmp, done)
	if err := rxSetup(unit); err != nil {
		log.Errorf("rxSetup error. %s", err)
		return
	}

	<-done
}
