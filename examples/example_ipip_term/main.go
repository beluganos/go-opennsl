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
	"net"

	"github.com/beluganos/go-opennsl/examples/util"
	"github.com/beluganos/go-opennsl/opennsl"
	"github.com/beluganos/go-opennsl/sal"

	log "github.com/sirupsen/logrus"
)

func rxPacket(unit int, done <-chan struct{}) error {
	pcfg, err := opennsl.PortConfigGet(unit)
	if err != nil {
		log.Errorf("PortConfigGet error. %s", err)
		return err
	}

	cpubmp, _ := pcfg.PBmp(opennsl.PORT_CONFIG_CPU)
	if err := opennsl.VLAN_ID_DEFAULT.PortAdd(unit, cpubmp, cpubmp); err != nil {
		log.Errorf("DEFAULT_VLAN.PortAdd error. %s", err)
		return err
	}

	flags := opennsl.RCO_F_ALL_COS
	err = opennsl.RxRegister(unit, 10, flags, func(unitRx int, pkt *opennsl.Pkt) {
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
		log.Errorf("RxRegister error. %s", err)
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

	log.Infof("rxStart ok")

	<-done
	return nil
}

func tunTerminator4(unit int, dst, src *net.IPNet) error {

	log.Infof("tunTerminator4 dst:%s, src:%s", dst, src)

	// IPv4 over IPv4
	tun4 := opennsl.NewTunnelTerminator(opennsl.TunnelTypeIPIP4toIP4)
	// tun4.SetVID(opennsl.VLAN_ID_DEFAULT)
	tun4.SetDstIPNet4(dst)
	tun4.SetSrcIPNet4(src)
	tun4.PBmp().Add(opennsl.Port(54), opennsl.Port(55))
	if err := tun4.Create(unit); err != nil {
		log.Errorf("Tun.Create error. %s", err)
		return err
	}

	// IPv6 over IPv4
	tun6 := opennsl.NewTunnelTerminator(opennsl.TunnelTypeIPIP4toIP6)
	// tun6.SetVID(opennsl.VLAN_ID_DEFAULT)
	tun6.SetDstIPNet4(dst)
	tun6.SetSrcIPNet4(src)
	tun6.PBmp().Add(opennsl.Port(56), opennsl.Port(57))
	if err := tun6.Create(unit); err != nil {
		log.Errorf("Tun.Create error. %s", err)
		return err
	}

	return nil
}

func tunTerminator6(unit int, dst, src *net.IPNet) error {

	log.Infof("tunTerminator6 dst:%s, src:%s", dst, src)

	// IPv4 over IPv6
	tun4 := opennsl.NewTunnelTerminator(opennsl.TunnelTypeIPIP6toIP4)
	// tun4.SetVID(opennsl.VLAN_ID_DEFAULT)
	tun4.SetDstIPNet6(dst)
	tun4.SetSrcIPNet6(src)
	tun4.PBmp().Add(opennsl.Port(54), opennsl.Port(55))
	if err := tun4.Create(unit); err != nil {
		log.Errorf("Tun.Create error. %s", err)
		return err
	}

	// IPv6 over IPv6
	tun6 := opennsl.NewTunnelTerminator(opennsl.TunnelTypeIPIP6toIP6)
	// tun6.SetVID(opennsl.VLAN_ID_DEFAULT)
	tun6.SetDstIPNet6(dst)
	tun6.SetSrcIPNet6(src)
	tun6.PBmp().Add(opennsl.Port(56), opennsl.Port(57))
	if err := tun6.Create(unit); err != nil {
		log.Errorf("Tun.Create error. %s", err)
		return err
	}

	return nil
}

func initDriver(unit int) error {
	if err := sal.DriverInit(); err != nil {
		log.Errorf("driver_init error. %s", err)
		return err
	}

	if err := util.PortDefaultConfig(unit); err != nil {
		log.Errorf("%s", err)
		return err
	}

	//if err := util.SwitchDefaultVlanConfig(unit); err != nil {
	//	log.Errorf("%s", err)
	//	return err
	//}

	if err := opennsl.SwitchL3EgressMode.Set(unit, 1); err != nil {
		log.Errorf("SwitchL3EgressMode.Set error. %s", err)
		return err
	}

	return nil
}

func main() {
	log.SetLevel(log.DebugLevel)

	unit := 0

	vid := opennsl.VLAN_ID_DEFAULT
	vid0 := opennsl.Vlan(11)
	vid1 := opennsl.Vlan(12)
	vid2 := opennsl.Vlan(13)
	vid3 := opennsl.Vlan(14)

	port0 := opennsl.Port(54)
	port1 := opennsl.Port(55)
	port2 := opennsl.Port(56)
	port3 := opennsl.Port(57)

	my_mac0, _ := net.ParseMAC("00:11:22:33:44:00")
	my_mac1, _ := net.ParseMAC("00:11:22:33:44:01")
	my_mac2, _ := net.ParseMAC("00:11:22:33:44:02")
	my_mac3, _ := net.ParseMAC("00:11:22:33:44:03")

	nh_mac0, _ := net.ParseMAC("00:66:77:88:99:00")
	nh_mac1, _ := net.ParseMAC("00:66:77:88:99:01")
	nh_mac2, _ := net.ParseMAC("00:66:77:88:99:02")
	nh_mac3, _ := net.ParseMAC("00:66:77:88:99:03")

	_, tun_local4, _ := net.ParseCIDR("100.0.1.1/32")
	_, tun_remot4, _ := net.ParseCIDR("100.0.1.2/32")
	_, tun_local6, _ := net.ParseCIDR("2010:2020::1/128")
	_, tun_remot6, _ := net.ParseCIDR("2010:2010::4/128")

	host4 := net.ParseIP("10.0.1.1")
	_, route4, _ := net.ParseCIDR("55.0.0.0/8")
	host6 := net.ParseIP("2010:2010::3")
	_, route6, _ := net.ParseCIDR("2010:2011::/64")

	log.Infof("DEVAULT_VLAN=%d", vid)

	//
	// Driver Initialize
	//
	if err := initDriver(unit); err != nil {
		log.Errorf("initDriver error. %s", err)
	}
	defer sal.DriverExit()

	//
	// L3 Interface
	//
	iface0, err := util.NewL3IfaceObjUntagged(unit, port0, vid0, my_mac0)
	if err != nil {
		log.Errorf("%s", err)
		return
	}

	iface1, err := util.NewL3IfaceObjUntagged(unit, port1, vid1, my_mac1)
	if err != nil {
		log.Errorf("%s", err)
		return
	}

	iface2, err := util.NewL3IfaceObjUntagged(unit, port2, vid2, my_mac2)
	if err != nil {
		log.Errorf("%s", err)
		return
	}

	iface3, err := util.NewL3IfaceObjUntagged(unit, port3, vid3, my_mac3)
	if err != nil {
		log.Errorf("%s", err)
		return
	}

	//
	// L3 Egress
	//
	if _, err := util.NewL3Egress(unit, port0, vid, iface0.IfaceID(), nh_mac0); err != nil {
		log.Errorf("%s", err)
		return
	}

	if _, err := util.NewL3Egress(unit, port1, vid, iface1.IfaceID(), nh_mac1); err != nil {
		log.Errorf("%s", err)
		return
	}

	l3eg2, err := util.NewL3Egress(unit, port2, vid, iface2.IfaceID(), nh_mac2)
	if err != nil {
		log.Errorf("%s", err)
		return
	}

	l3eg3, err := util.NewL3Egress(unit, port3, vid, iface3.IfaceID(), nh_mac3)
	if err != nil {
		log.Errorf("%s", err)
		return
	}

	//
	// L2 Address
	//
	if _, err := util.AddL2Addr(unit, port0, vid, my_mac0, opennsl.L2_L3LOOKUP); err != nil {
		log.Errorf("%s", err)
		return
	}
	if _, err := util.AddL2Addr(unit, port1, vid, my_mac1, opennsl.L2_L3LOOKUP); err != nil {
		log.Errorf("%s", err)
		return
	}
	if _, err := util.AddL2Addr(unit, port2, vid, my_mac2, opennsl.L2_L3LOOKUP); err != nil {
		log.Errorf("%s", err)
		return
	}
	if _, err := util.AddL2Addr(unit, port3, vid, my_mac3, opennsl.L2_L3LOOKUP); err != nil {
		log.Errorf("%s", err)
		return
	}

	//
	// Tunnel Terminator
	//
	if err := tunTerminator4(unit, tun_local4, tun_remot4); err != nil {
		log.Errorf("%s", err)
		return
	}

	if err := tunTerminator6(unit, tun_local6, tun_remot6); err != nil {
		log.Errorf("%s", err)
		return
	}

	//
	// Host
	//
	if _, err := util.AddHost(unit, host4, l3eg2, opennsl.VRF_NONE); err != nil {
		log.Errorf("%s", err)
		return
	}

	if _, err := util.AddHost6(unit, host6, l3eg2, opennsl.VRF_NONE); err != nil {
		log.Errorf("%s", err)
		return
	}

	//
	// Route
	//
	if _, err := util.AddRoute(unit, route4, l3eg3, opennsl.VRF_NONE); err != nil {
		log.Errorf("%s", err)
		return
	}

	if _, err := util.AddRoute6(unit, route6, l3eg3, opennsl.VRF_NONE); err != nil {
		log.Errorf("%s", err)
		return
	}

	//
	// Wait signal
	//
	done := make(chan struct{})
	go util.WatchSignal(done)

	//
	// Rx
	//
	if err := rxPacket(unit, done); err != nil {
		log.Errorf("%s", err)
		return
	}
}
