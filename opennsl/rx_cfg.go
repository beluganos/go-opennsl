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

package opennsl

/*
#include <opennsl/types.h>
#include <opennsl/rx.h>
#include "helper.h"
*/
import "C"

//
// RxChanCfg
//
type RxChanCfg C.opennsl_rx_chan_cfg_t

func (v *RxChanCfg) C() *C.opennsl_rx_chan_cfg_t {
	return (*C.opennsl_rx_chan_cfg_t)(v)
}

func (v *RxChanCfg) Chains() int {
	return int(v.chains)
}

func (v *RxChanCfg) SetChains(chains int) {
	v.chains = C.int(chains)
}

/* DEPRECATED
func (v *RxChanCfg) RatePps() int {
	return int(v.rate_pps)
}

func (v *RxChanCfg) SetRatePps(pps int) {
	v.rate_pps = C.int(pps)
}
*/

func (v *RxChanCfg) Flags() int {
	return int(v.flags)
}

func (v *RxChanCfg) SetFlags(flags int) {
	v.flags = C.int(flags)
}
func (v *RxChanCfg) CosBmp() uint32 {
	return uint32(v.cos_bmp)
}

func (v *RxChanCfg) SetCosBmp(bmp uint32) {
	v.cos_bmp = C.uint32(bmp)
}

//
// RxCfg
//
type RxCfg C.opennsl_rx_cfg_t

func (v *RxCfg) C() *C.opennsl_rx_cfg_t {
	return (*C.opennsl_rx_cfg_t)(v)
}

func (v *RxCfg) PktSize() int {
	return int(v.pkt_size)
}

func (v *RxCfg) SetPktSize(pktSize int) {
	v.pkt_size = C.int(pktSize)
}

func (v *RxCfg) PktsPerChain() int {
	return int(v.pkts_per_chain)
}

func (v *RxCfg) SetPktsPerChain(pktsPerChain int) {
	v.pkts_per_chain = C.int(pktsPerChain)
}

func (v *RxCfg) GlobalPps() int {
	return int(v.global_pps)
}

func (v *RxCfg) SetGlobalPps(pps int) {
	v.global_pps = C.int(pps)
}

func (v *RxCfg) MaxBurst() int {
	return int(v.max_burst)
}

func (v *RxCfg) SetMaxBurst(maxBurst int) {
	v.max_burst = C.int(maxBurst)
}

func (v *RxCfg) ChanCfgs() []*RxChanCfg {
	cfgs := make([]*RxChanCfg, len(v.chan_cfg))
	for index := 0; index < len(v.chan_cfg); index++ {
		cfgs[index] = v.ChanCfg(index)
	}
	return cfgs
}

func (v *RxCfg) ChanCfg(index int) *RxChanCfg {
	if index < 0 || index >= len(v.chan_cfg) {
		return nil
	}
	return (*RxChanCfg)(&v.chan_cfg[index])
}

func (v *RxCfg) CPUAddrs() []int {
	num := C.uint32(v.num_of_cpu_addresses)
	var index C.uint32

	addrs := make([]int, num)
	for index = 0; index < num; index++ {
		addr := C._opennsl_util_int_array_get(v.cpu_address, num, index)
		addrs[index] = int(*addr)
	}
	return addrs
}

//
// API
//
func NewRxCfg() *RxCfg {
	cfg := &RxCfg{}
	cfg.Init()
	return cfg
}

func (v *RxCfg) Init() {
	C.opennsl_rx_cfg_t_init(v.C())
}

func RxCfgInit(unit int) error {
	rc := C.opennsl_rx_cfg_init(C.int(unit))
	return ParseError(rc)
}

func (v *RxCfg) SafeStart(unit int) error {
	if active := RxActive(unit); !active {
		return v.Start(unit)
	}
	return nil
}

func RxCfgStart(unit int, v *RxCfg) error {
	rc := C.opennsl_rx_start(C.int(unit), v.C())
	return ParseError(rc)
}

func (v *RxCfg) Start(unit int) error {
	return RxCfgStart(unit, v)
}

func RxCfgStop(unit int, v *RxCfg) error {
	rc := C.opennsl_rx_stop(C.int(unit), v.C())
	return ParseError(rc)
}

func (v *RxCfg) Stop(unit int) error {
	return RxCfgStop(unit, v)
}

func RxCfgGet(unit int) (*RxCfg, error) {
	cfg := RxCfg{}
	cfg.Init()

	rc := C.opennsl_rx_cfg_get(C.int(unit), cfg.C())
	if err := ParseError(rc); err != nil {
		return nil, err
	}

	return &cfg, nil
}
