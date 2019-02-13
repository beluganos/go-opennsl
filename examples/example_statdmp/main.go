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
	"time"

	"github.com/beluganos/go-opennsl/examples/util"
	"github.com/beluganos/go-opennsl/opennsl"
	"github.com/beluganos/go-opennsl/sal"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)

	unit := 0
	port := opennsl.Port(50)

	if err := sal.DriverInit(); err != nil {
		log.Errorf("driver_init error. %s", err)
		return
	}
	defer sal.DriverExit()

	bootflag := sal.DriverBootFlagsGet()
	log.Infof("boot flag = %x", bootflag)

	if err := util.PortDefaultConfig(unit); err != nil {
		log.Errorf("%s", err)
		return
	}

	if err := opennsl.SwitchL3EgressMode.Set(unit, 1); err != nil {
		log.Errorf("%s", err)
		return
	}

	if err := opennsl.StatInit(unit); err != nil {
		log.Errorf("%s", err)
		return
	}

	if err := opennsl.StatSync(unit); err != nil {
		log.Errorf("%s", err)
		return
	}

	statvals := []opennsl.StatVal{
		opennsl.SPLSnmpIfInUcastPkts,
		opennsl.SPLSnmpIfInNUcastPkts,
	}

	time.Sleep(5)

	values, err := opennsl.StatValMultiGet(unit, port, statvals...)
	if err != nil {
		log.Errorf("%s", err)
		return
	}

	for index, value := range values {
		log.Infof("%s: %d", statvals[index], value)
	}
}
