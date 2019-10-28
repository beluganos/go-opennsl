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

type Args struct {
	Fec     string
	Verbose bool
}

func (a *Args) Parse() {
	flag.StringVar(&a.Fec, "fec", "", "fec on/off.")
	flag.BoolVar(&a.Verbose, "v", false, "show detail message.")

	flag.Parse()
}

func NewArgs() *Args {
	args := &Args{}
	args.Parse()
	return args
}

func watchSignal(done chan struct{}) {

	ch := make(chan os.Signal, 1)
	defer close(ch)

	signal.Notify(ch, os.Interrupt)
	<-ch

	log.Infof("Interrupt signal.")

	close(done)
}

func setFec(unit int, fec string) error {

	var fecval uint32

	switch fec {
	case "on":
		log.Infof("FEC ON")
		fecval = opennsl.PORT_PHY_CONTROL_FEC_ON

	case "off":
		log.Infof("FEC OFF")
		fecval = opennsl.PORT_PHY_CONTROL_FEC_OFF

	default:
		log.Debugf("FEC not changed.")
		return nil
	}

	pcfg, err := opennsl.PortConfigGet(unit)
	if err != nil {
		return err
	}

	pbmp, _ := pcfg.PBmp(opennsl.PORT_CONFIG_E)

	return pbmp.Each(func(port opennsl.Port) error {
		err := port.PhyControlSet(
			unit,
			opennsl.PORT_PHY_CONTROL_FORWARD_ERROR_CORRECTION,
			fecval,
		)
		if err != nil {
			log.Errorf("PhyControlSet(FEC) error. port %d %s", port, err)
		} else {
			log.Infof("PhyControlSet(FEC) ok. port %d %s", port, fec)
		}

		return nil
	})
}

func main() {
	args := NewArgs()

	if args.Verbose {
		log.SetLevel(log.DebugLevel)
	}

	unit := 0

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

	setFec(unit, args.Fec)

	if err := opennsl.LinkscanRegister(unit, "test", func(unit int, key string, port opennsl.Port, portInfo *opennsl.PortInfo) {
		log.Infof("Link Status Changed. port:%d status:%s", port, portInfo.LinkStatus())
		go func() {
			enable := func() opennsl.PortEnable {
				if portInfo.LinkStatus() != opennsl.LINK_STATUS_UP {
					return opennsl.PORT_ENABLE_TRUE
				}
				return opennsl.PORT_ENABLE_FALSE
			}()

			timer := time.NewTimer(3 * time.Second)
			<-timer.C

			port.EnableSet(unit, enable)
		}()
	}); err != nil {
		log.Errorf("LinkscanRegister error. %s", err)
		return
	}
	defer opennsl.LinkscanUnregister(unit, "test")

	done := make(chan struct{})
	go watchSignal(done)

	<-done
}
