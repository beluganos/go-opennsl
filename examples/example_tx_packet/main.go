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

func pktSend(unit int, port opennsl.Port, data []byte) error {
	pkt, err := opennsl.PktAlloc(unit, len(data), opennsl.PKT_F_NONE)
	if err != nil {
		log.Errorf("PktAlloc error. %s", err)
		return err
	}

	defer pkt.Free(unit)

	if err := pkt.Memcpy(0, data); err != nil {
		log.Errorf("pkt.Memcpy error. %s", err)
		return err
	}

	pkt.TxPBmp().Clear()
	pkt.TxPBmp().Add(port)
	log.Debugf("pkt.TxPBMP %v", pkt.TxPBmp())

	if err := pkt.Tx(unit); err != nil {
		log.Errorf("pkt.Tx error. %s", err)
		return err
	}

	log.Debugf("pktSend len:%d, port:%d", len(data), port)
	return nil
}

func txSetup(unit int, port opennsl.Port) error {
	data := []byte{
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x82, 0x2f,
		0x2e, 0x42, 0x46, 0x74, 0x08, 0x06, 0x00, 0x01,
		0x08, 0x00, 0x06, 0x04, 0x00, 0x01, 0x82, 0x2f,
		0x2e, 0x42, 0x46, 0x74, 0x0a, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x00,
		0x00, 0x02, 0x00, 0x00, 0x00, 0x00,
	}

	for index := 0; index < 5; index++ {
		if err := pktSend(unit, port, data); err != nil {
			log.Errorf("pktSend error. %s", err)
			return err
		}
	}

	return nil
}

func main() {
	log.SetLevel(log.DebugLevel)

	unit := int(0)
	port := opennsl.Port(50)

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

	time.Sleep(1 * time.Second)
	if err := txSetup(unit, port); err != nil {
		log.Errorf("txSetup error. %s", err)
		return
	}
}
