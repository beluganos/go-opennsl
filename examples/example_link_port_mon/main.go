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
	"flag"
	"os"
	"os/signal"
	"time"

	"github.com/beluganos/go-opennsl/examples/util"
	"github.com/beluganos/go-opennsl/opennsl"
	"github.com/beluganos/go-opennsl/sal"

	log "github.com/sirupsen/logrus"
)

const (
	LINKSCAN_REGISTER = "test"
)

func watchSignal(done chan struct{}) {

	ch := make(chan os.Signal, 1)
	defer close(ch)

	signal.Notify(ch, os.Interrupt)
	<-ch

	log.Infof("Interrupt signal.")

	close(done)
}

func flipPortEnable(unit int, port opennsl.Port, done <-chan struct{}) {
	tick := time.NewTicker(1 * time.Second)
	defer tick.Stop()

	enable := opennsl.PORT_ENABLE_TRUE

	for {
		select {
		case <-tick.C:
			log.Infof("port.EnableSet %d %s", port, enable)
			if err := port.EnableSet(unit, enable); err != nil {
				log.Errorf("port.EnableSet error. %s", err)
			}

			enable = enable.Flip()

		case <-done:
			log.Infof("flipPortEnable exit.")
			return
		}
	}
}

func main() {
	var port int
	flag.IntVar(&port, "p", 1, "monitoring port number.")
	flag.Parse()

	log.SetLevel(log.DebugLevel)

	unit := 0

	log.Infof("unit:%d port: %d", unit, port)

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

	err := opennsl.LinkscanRegister(unit, LINKSCAN_REGISTER, func(unit int, key string, port opennsl.Port, portInfo *opennsl.PortInfo) {
		log.Infof("Link Status Changed. port:%d status:%s", port, portInfo.LinkStatus())
	})

	if err != nil {
		log.Errorf("LinkscanRegister error. %s", err)
		return
	}
	defer opennsl.LinkscanUnregister(unit, LINKSCAN_REGISTER)

	done := make(chan struct{})
	go flipPortEnable(unit, opennsl.Port(port), done)
	go watchSignal(done)

	<-done
}
