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
	"os"
	"time"

	flag "github.com/spf13/pflag"

	"github.com/beluganos/go-opennsl/examples/util"
	"github.com/beluganos/go-opennsl/opennsl"
	"github.com/beluganos/go-opennsl/sal"
	log "github.com/sirupsen/logrus"
)

const (
	ARGS_CONFIG_PATH_DEFAULT = "/etc/opennsl/example_l2_switch.yaml"
	ARGS_CONFIG_TYPE_DEFAULT = "yaml"
	ARGS_CONFIG_AGE_DEFAULT  = 15 * time.Second
	ARGS_CONFIG_DUMP_DEFAULT = 15 * time.Second
	ARGS_CONFIG_UNIT_DEFAULT = 0
)

type Args struct {
	configPath string
	configType string
	ageTimer   time.Duration
	dumpTimer  time.Duration
	unit       int
	verbose    bool
}

func NewArgs() *Args {
	args := Args{}
	flag.StringVarP(&args.configPath, "config-path", "c", ARGS_CONFIG_PATH_DEFAULT, "config path.")
	flag.StringVarP(&args.configType, "config-type", "", ARGS_CONFIG_TYPE_DEFAULT, "config type.")
	flag.DurationVarP(&args.ageTimer, "age-timer", "", ARGS_CONFIG_AGE_DEFAULT, "aging timer interval.")
	flag.DurationVarP(&args.dumpTimer, "dump-timer", "", ARGS_CONFIG_DUMP_DEFAULT, "dump timer interval.")
	flag.IntVarP(&args.unit, "unit", "", ARGS_CONFIG_UNIT_DEFAULT, "unit number.")
	flag.BoolVarP(&args.verbose, "verbose", "v", false, "show detail message.")
	flag.Parse()

	return &args
}

func (a *Args) AgeTimer() int {
	return int(a.ageTimer.Seconds())
}

func (a *Args) dump() {
	log.Infof("config-path: '%s'", a.configPath)
	log.Infof("config-type: '%s'", a.configType)
	log.Infof("age-timer  : %s", a.ageTimer)
}

func dumpConfig(cfg *Config) {
	for _, vlan := range cfg.Vlans {
		vid := vlan.Vlan()
		log.Infof("access: %d %v", vid, vlan.AccessPorts())
		log.Infof("trunk : %d %v", vid, vlan.TrunkPorts())
		log.Infof("nTag  : %d %v", vid, vlan.NativeTagPorts())
		log.Infof("nUntag: %d %v", vid, vlan.NativeUntagPorts())
	}
}

func main() {

	args := NewArgs()
	args.dump()

	if args.verbose {
		log.SetLevel(log.DebugLevel)
	}

	cfg, err := ReadConfig(args.configPath, args.configType)
	if err != nil {
		log.Errorf("ReadConfig error. %s", err)
		os.Exit(1)
	}

	dumpConfig(cfg)

	unit := args.unit

	// initializeinit driver.
	if err := sal.DriverInit(); err != nil {
		log.Errorf("driver_init error. %s", err)
		os.Exit(1)
	}
	defer sal.DriverExit()

	log.Infof("DriverInit ok.")

	if err := util.PortDefaultConfig(unit); err != nil {
		log.Errorf("%s", err)
		os.Exit(1)
	}

	log.Infof("PortDefaultConfig ok.")

	if err := opennsl.L2AddrAgeTimerSet(unit, args.AgeTimer()); err != nil {
		log.Errorf("L2AddrAgeTimerSet error. %s", err)
		os.Exit(1)
	}

	log.Infof("L2AddrAgeTimerSet ok. sec=%d", args.AgeTimer())

	done := make(chan struct{})
	startDumpL2Addrs(args.dumpTimer, done)
	watchL2Addrs(unit)

	tmpBmp := opennsl.NewPBmp()

	for _, vlancfg := range cfg.Vlans {
		vlan := NewVlan(vlancfg.Vlan())
		if err := vlan.Create(unit); err != nil {
			log.Errorf("VlanCreate error. vlan=%d %s", vlan.vid, err)
			continue
		}

		// access ports
		tmpBmp.Clear().Add(vlancfg.AccessPorts()...)
		log.Debugf("AddAccessPort vlan=%d ports=%s", vlan.vid, tmpBmp)

		if err := vlan.AddAccessPort(unit, tmpBmp); err != nil {
			log.Errorf("AddAccessPort error. vlan=%d ports=%s %s", vlan.vid, tmpBmp, err)
		}

		// trunk ports
		tmpBmp.Clear().Add(vlancfg.TrunkPorts()...)
		log.Debugf("AddTrunkPort vlan=%d ports=%s", vlan.vid, tmpBmp)

		if err := vlan.AddTrunkPort(unit, tmpBmp); err != nil {
			log.Errorf("AddTrunkPort error. vlan=%d ports=%s %s", vlan.vid, tmpBmp, err)
		}

		// native tag ports
		tmpBmp.Clear().Add(vlancfg.NativeTagPorts()...)
		log.Debugf("AddNativeTagPort vlan=%d port=%s", vlan.vid, tmpBmp)

		if err := vlan.AddNativeTagPort(unit, tmpBmp); err != nil {
			log.Debugf("AddNativeTagPort error. vlan=%d ports=%s %s", vlan.vid, tmpBmp, err)
		}

		// native untagged ports
		tmpBmp.Clear().Add(vlancfg.NativeUntagPorts()...)
		log.Debugf("AddNativeUntagPort vlan=%d ports=%s", vlan.vid, tmpBmp)

		if err := vlan.AddNativeUntaggedPort(unit, tmpBmp); err != nil {
			log.Debugf("AddNativeUntaggedPort error. vlan=%d ports=%s %s", vlan.vid, tmpBmp, err)
		}
	}

	<-done
}
