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
	"encoding/hex"
	"os"
	"os/signal"

	"github.com/beluganos/go-opennsl/examples/util"
	"github.com/beluganos/go-opennsl/opennsl"
	"github.com/beluganos/go-opennsl/sal"

	log "github.com/sirupsen/logrus"
)

func rxSetup(unit int, done chan struct{}) error {

	err := opennsl.RxRegister(unit, 10, opennsl.RCO_F_ALL_COS, func(unit int, pkt *opennsl.Pkt) {
		log.Debugf("pkt  : %p len:%d tot:%d", pkt, pkt.PktLen(), pkt.TotLen())
		log.Debugf("unit : %d", pkt.Unit())
		log.Debugf("flags: %d", pkt.Flags())
		log.Debugf("cos  : %d", pkt.Cos())
		log.Debugf("vid  : %d", pkt.VID())
		log.Debugf("port : src:%d dst:%d", pkt.SrcPort(), pkt.DstPort())
		log.Debugf("rx   : port    : %d", pkt.RxPort())
		log.Debugf("rx   : untagged: %d", pkt.RxUntagged())
		log.Debugf("rx   : matched : %d", pkt.RxMatched())

		log.Debugf("rx   : reasons : %d", pkt.RxReasons())
		pkt.RxReasons().ForEach(func(r opennsl.RxReason) error {
			log.Debugf("rx   : reasons : %s", r)
			return nil
		})

		log.Debugf("blk  : #%d", pkt.BlkCount())
		for index, blk := range pkt.Blks() {
			log.Debugf("blk[%d] len=%d", index, blk.Len())
			b := blk.Data()
			log.Debugf("\n%s", hex.Dump(b[:128]))
		}
	})
	if err != nil {
		return err
	}
	defer opennsl.RxUnregister(unit, 10)

	log.Infof("RX callback registered.")

	if active := opennsl.RxActive(unit); !active {
		cfg := opennsl.NewRxCfg()
		cfg.SetPktSize(16 * 1024)
		cfg.SetPktsPerChain(16)
		cfg.SetGlobalPps(200)
		cfg.ChanCfg(1).SetChains(4)
		cfg.ChanCfg(1).SetCosBmp(0xffffffff)

		if err := opennsl.RxStart(unit, cfg); err != nil {
			return err
		}

		log.Infof("RX activated.")

		defer cfg.Stop(unit)
	}

	if done != nil {
		<-done
	}

	return nil
}

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
	vid := opennsl.Vlan(12)

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

	pcfg, err := opennsl.PortConfigGet(unit)
	if err != nil {
		log.Errorf("PortConfigGet error. %s", err)
		return
	}

	cpubmp, _ := pcfg.PBmp(opennsl.PORT_CONFIG_CPU)
	if err := opennsl.VLAN_ID_DEFAULT.PortAdd(unit, cpubmp, cpubmp); err != nil {
		log.Errorf("DEFAULT_VLAN.PortAdd error. %s", err)
		return
	}
	if _, err := vid.Create(unit); err != nil {
		log.Errorf("vid.Create error. %s", err)
		return
	}
	if err := vid.PortAdd(unit, cpubmp, cpubmp); err != nil {
		log.Errorf("vid.PortAdd error. %s", err)
		return
	}

	done := make(chan struct{})
	go watchSignal(done)

	if err := rxSetup(unit, done); err != nil {
		log.Errorf("rxSetup error. %s", err)
		return
	}
}
