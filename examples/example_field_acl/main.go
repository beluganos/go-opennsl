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
	"flag"
	"net"
	"os"
	"os/signal"

	"github.com/beluganos/go-opennsl/examples/util"
	"github.com/beluganos/go-opennsl/opennsl"
	"github.com/beluganos/go-opennsl/sal"
	"golang.org/x/sys/unix"

	log "github.com/sirupsen/logrus"
)

const (
	ETH_P_MASK     = 0xffff
	PROTO_MASK     = 0xff
	ACTION_COS_QUE = 1
	ACTION_NOPARAM = 0
	L4PORT_SIP     = 5050
)

func newIPFieldGroup(unit int) (opennsl.FieldGroup, error) {
	qset := opennsl.FieldQSet{}
	qset.Init()
	qset.Add(
		opennsl.FieldQualifyEtherType,
		opennsl.FieldQualifyDstIp,
	)

	fg, err := opennsl.FieldGroupCreate(unit, &qset, opennsl.FIELD_GROUP_PRIO_ANY)
	if err != nil {
		log.Errorf("[IPDST] FieldGroupCreate error. %s", err)
		return 0, err
	}

	return fg, nil
}

func addIPFieldEntry(unit int, group opennsl.FieldGroup, dst string) error {
	ip, ipnet, err := net.ParseCIDR(dst)
	if err != nil {
		log.Errorf("[IPDST] ParseCIDR error. %s", err)
		return err
	}

	entry, err := group.EntryCreate(unit)
	if err != nil {
		log.Errorf("[IPDST] EntryCreate error. %s", err)
		return err
	}

	entry.Qualify().EtherType(unit, unix.ETH_P_IP, ETH_P_MASK)
	entry.Qualify().DstIp(unit, ip, ipnet.Mask)
	entry.Action().AddP(unit, opennsl.NewFieldActionCosQCpuNew(ACTION_COS_QUE))
	entry.Action().AddP(unit, opennsl.NewFieldActionCopyToCpu())

	if err := entry.Install(unit); err != nil {
		log.Errorf("[IPDST] Entry install error. %s", err)
		return err
	}

	return nil
}

func addIPAcl(unit int) error {
	group, err := newIPFieldGroup(unit)
	if err != nil {
		return err
	}

	if err := addIPFieldEntry(unit, group, "20.1.1.0/24"); err != nil {
		return err
	}

	log.Debugf("[IPDST] OK.")
	return nil
}

func newIPProtoFieldGroup(unit int) (opennsl.FieldGroup, error) {
	qset := opennsl.FieldQSet{}
	qset.Init()
	qset.Add(
		opennsl.FieldQualifyEtherType,
		opennsl.FieldQualifyIpProtocol,
	)

	fg, err := opennsl.FieldGroupCreate(unit, &qset, opennsl.FIELD_GROUP_PRIO_ANY)
	if err != nil {
		log.Errorf("[PROTO] FieldGroupCreate error. %s", err)
		return 0, err
	}

	return fg, nil
}

func addIPProtoFieldEntry(unit int, group opennsl.FieldGroup, proto uint8) error {
	entry, err := group.EntryCreate(unit)
	if err != nil {
		log.Errorf("[PROTO] EntryCreate error. %s", err)
		return err
	}

	entry.Qualify().EtherType(unit, unix.ETH_P_IP, ETH_P_MASK)
	entry.Qualify().IpProtocol(unit, proto, PROTO_MASK)
	entry.Action().AddP(unit, opennsl.NewFieldActionCosQCpuNew(ACTION_COS_QUE))
	entry.Action().AddP(unit, opennsl.NewFieldActionCopyToCpu())

	if err := entry.Install(unit); err != nil {
		log.Errorf("[PROTO] Entry install error. %s", err)
		return err
	}

	return nil
}

func addIPProtoAcl(unit int) error {
	group, err := newIPProtoFieldGroup(unit)
	if err != nil {
		return err
	}

	if err := addIPProtoFieldEntry(unit, group, unix.IPPROTO_UDP); err != nil {
		return err
	}

	log.Debugf("[PROTO] OK.")
	return nil
}

func newL4PortFieldGroup(unit int) (opennsl.FieldGroup, error) {
	qset := opennsl.FieldQSet{}
	qset.Init()
	qset.Add(
		opennsl.FieldQualifyEtherType,
		// opennsl.FieldQualifyIpProtocol,
		opennsl.FieldQualifyL4DstPort,
	)

	fg, err := opennsl.FieldGroupCreate(unit, &qset, opennsl.FIELD_GROUP_PRIO_ANY)
	if err != nil {
		log.Errorf("[L4PORT] FieldGroupCreate error. %s", err)
		return 0, err
	}

	return fg, nil
}

func addL4PortFieldEntry(unit int, group opennsl.FieldGroup, proto uint8, port opennsl.L4Port) error {
	entry, err := group.EntryCreate(unit)
	if err != nil {
		log.Errorf("[L4PORT] EntryCreate error. %s", err)
		return err
	}

	entry.Qualify().EtherType(unit, unix.ETH_P_IP, ETH_P_MASK)
	// entry.Qualify().IpProtocol(unit, proto, 0xff)
	entry.Qualify().L4DstPort(unit, port, PROTO_MASK)
	entry.Action().AddP(unit, opennsl.NewFieldActionCosQCpuNew(ACTION_COS_QUE))
	entry.Action().AddP(unit, opennsl.NewFieldActionCopyToCpu())

	if err := entry.Install(unit); err != nil {
		log.Errorf("[L4PORT] Entry install error. %s", err)
		return err
	}

	return nil
}

func addL4PortAcl(unit int) error {
	group, err := newL4PortFieldGroup(unit)
	if err != nil {
		return err
	}

	if err := addL4PortFieldEntry(unit, group, unix.IPPROTO_UDP, L4PORT_SIP); err != nil {
		return err
	}

	log.Debugf("[L4PORT] OK.")
	return nil
}

func aclSetup(unit int, names []string) error {

	for _, name := range names {
		switch name {
		case "ip":
			if err := addIPAcl(unit); err != nil {
				log.Errorf("addIPAcl. %s", err)
				return err
			}

		case "ipproto":
			if err := addIPProtoAcl(unit); err != nil {
				log.Errorf("addIPProtoAcl. %s", err)
				return err
			}

		case "l4port":
			if err := addL4PortAcl(unit); err != nil {
				log.Errorf("addL4PortAcl. %s", err)
				return err
			}

		default:
			log.Warnf("unknown name. %s", name)
		}
	}

	return nil
}

func rxSetup(unit int, done chan struct{}) error {
	flg := opennsl.NewRxCallbackFlags(1)
	err := opennsl.RxRegister(unit, 100, flg, func(unit int, pkt *opennsl.Pkt) {
		if ok := pkt.RxReasons().Has(opennsl.RxReasonFilterMatch); !ok {
			log.Debugf("rx skip")
			return
		}

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
	defer opennsl.RxUnregister(unit, 100)

	log.Infof("RX callback registered.")

	if active := opennsl.RxActive(unit); !active {
		cfg := opennsl.NewRxCfg()
		cfg.SetPktSize(16 * 1024)
		cfg.SetPktsPerChain(16)
		cfg.SetGlobalPps(30000)
		cfg.ChanCfg(1).SetChains(4)
		//cfg.ChanCfg(1).SetCosBmp(0xffffffff)

		if err := opennsl.RxStart(unit, cfg); err != nil {
			return err
		}

		log.Infof("RX activated.")

		defer cfg.Stop(unit)
	}

	<-done

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

func getopts() []string {
	flag.Parse()
	return flag.Args()
}

func main() {
	names := getopts()

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

	pcfg, err := opennsl.PortConfigGet(unit)
	if err != nil {
		log.Errorf("PortConfigGet error. %s", err)
		return
	}

	cpubmp, _ := pcfg.PBmp(opennsl.PORT_CONFIG_CPU)
	if err := opennsl.VLAN_ID_DEFAULT.PortAdd(unit, cpubmp, cpubmp); err != nil {
		log.Errorf("DEFAULT_VLAN.PortAdd. %s", err)
		return
	}

	if err := aclSetup(unit, names); err != nil {
		log.Errorf("aclSetup error. %s", err)
		return
	}

	done := make(chan struct{})
	go watchSignal(done)

	if err := rxSetup(unit, done); err != nil {
		log.Errorf("rxSetup error %s", err)
		return
	}
}
