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
	"time"

	"github.com/beluganos/go-opennsl/opennsl"
	log "github.com/sirupsen/logrus"
)

func dumpL2Addrs() error {
	log.Infof("---------------------------------------------")
	return opennsl.L2Traverse(0, func(unit int, l2addr *opennsl.L2Addr) opennsl.OpenNSLError {
		log.Infof("L2Addr:%s port:%d flags:%s", l2addr, l2addr.Port(), l2addr.Flags())
		return opennsl.E_NONE
	})
}

func startDumpL2Addrs(interval time.Duration, done chan struct{}) {
	if interval == 0 {
		return
	}

	go func() {
		tick := time.NewTicker(interval)
		defer tick.Stop()

		log.Infof("dumpL2Addrs start.")

		for {
			select {
			case <-tick.C:
				if err := dumpL2Addrs(); err != nil {
					log.Errorf("dumpL2Addrs error. %s", err)
					return
				}

			case <-done:
				log.Infof("dumpL2Addrs exit.")
				return
			}
		}
	}()
}

func watchL2Addrs(unit int) {
	err := opennsl.L2AddrRegister(unit, func(unitCb int, l2addr *opennsl.L2Addr, oper opennsl.L2CallbackOper) {
		log.Infof("---------------------------------------------")
		log.Infof("L2Addr:%s port:%d flags:%s %s", l2addr, l2addr.Port(), l2addr.Flags(), oper)
	})

	if err != nil {
		log.Errorf("L2AddrRegister error. %s", err)
	}
}
