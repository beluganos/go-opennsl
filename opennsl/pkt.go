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
#include <stdlib.h>
#include <opennsl/pkt.h>
#include "helper.h"
*/
import "C"

import (
	"fmt"
	"io"
	"strings"
	"unsafe"
)

type PktFlags uint32

func (v PktFlags) C() C.uint32 {
	return C.uint32(v)
}

func NewPktFlags(flags ...PktFlags) PktFlags {
	v := PKT_F_NONE
	for _, flag := range flags {
		v |= flag
	}
	return v
}

const (
	PKT_F_NONE    PktFlags = 0
	PKT_F_NO_VTAG PktFlags = C.OPENNSL_PKT_F_NO_VTAG
	TX_CRC_ALLOC  PktFlags = C.OPENNSL_TX_CRC_ALLOC
	TX_CRC_REGEN  PktFlags = C.OPENNSL_TX_CRC_REGEN
	TX_CRC_APPEND PktFlags = C.OPENNSL_TX_CRC_APPEND
	TX_ETHER      PktFlags = C.OPENNSL_TX_ETHER
	PKT_F_TRUNK   PktFlags = C.OPENNSL_PKT_F_TRUNK
)

var pktFlags_names = map[PktFlags]string{
	PKT_F_NO_VTAG: "PKT_F_NO_VTAG",
	TX_CRC_ALLOC:  "TX_CRC_ALLOC",
	TX_CRC_REGEN:  "TX_CRC_REGEN",
	TX_CRC_APPEND: "TX_CRC_APPEND",
	TX_ETHER:      "TX_ETHER",
	PKT_F_TRUNK:   "PKT_F_TRUNK",
}

var pktFlags_values = map[string]PktFlags{
	"PKT_F_NO_VTAG": PKT_F_NO_VTAG,
	"TX_CRC_ALLOC":  TX_CRC_ALLOC,
	"TX_CRC_REGEN":  TX_CRC_REGEN,
	"TX_CRC_APPEND": TX_CRC_APPEND,
	"TX_ETHER":      TX_ETHER,
	"PKT_F_TRUNK":   PKT_F_TRUNK,
}

func (v PktFlags) String() string {
	names := make([]string, 0)
	for val, name := range pktFlags_names {
		if val&v != 0 {
			names = append(names, name)
		}
	}

	return strings.Join(names, "|")
}

func ParsePktFlags(s string) (PktFlags, error) {
	if v, ok := pktFlags_values[s]; ok {
		return v, nil
	}
	return PKT_F_NONE, fmt.Errorf("Invalid PktFlags. %s", s)
}

//
// RxReasons
//
type RxReasons C.opennsl_rx_reasons_t

func (v *RxReasons) C() *C.opennsl_rx_reasons_t {
	return (*C.opennsl_rx_reasons_t)(v)
}

func (v *RxReasons) Bits() []uint32 {
	bits := make([]uint32, len(v.pbits))
	for index, pbit := range v.pbits {
		bits[index] = uint32(pbit)
	}
	return bits
}

func (v *RxReasons) Set(reason RxReason) {
	C._opennsl_rx_reason_set(v.C(), reason.C())
}

func (v *RxReasons) Has(reason RxReason) bool {
	return C._opennsl_rx_reason_get(v.C(), reason.C()) != 0
}

func (v *RxReasons) ForEach(callback func(reason RxReason) error) error {
	for index := C._SHR_RX_INVALID; index < C._SHR_RX_REASON_COUNT; index++ {
		if ok := v.Has(RxReason(index)); ok {
			if err := callback(RxReason(index)); err != nil {
				return err
			}
		}
	}

	return nil
}

//
// PktBlk
//
type PktBlk C.opennsl_pkt_blk_t

func (v *PktBlk) C() *C.opennsl_pkt_blk_t {
	return (*C.opennsl_pkt_blk_t)(v)
}

func (v *PktBlk) Len() int {
	return int(v.len)
}

func (v *PktBlk) Data() []byte {
	return C.GoBytes(unsafe.Pointer(v.data), v.len)
}

func (v *PktBlk) DataN(n int) []byte {
	c_n := func() C.int {
		if cn := C.int(n); cn < v.len {
			return cn
		}
		return v.len
	}()
	return C.GoBytes(unsafe.Pointer(v.data), c_n)
}

//
// Pkt
//
type Pkt C.opennsl_pkt_t

func (v *Pkt) C() *C.opennsl_pkt_t {
	return (*C.opennsl_pkt_t)(v)
}

func (v *Pkt) BlkCount() uint8 {
	return uint8(v.blk_count)
}

func (v *Pkt) Blks() []*PktBlk {
	blks := make([]*PktBlk, v.blk_count)
	var index C.uint8
	for index = 0; index < v.blk_count; index++ {
		blk := C._opennsl_pkt_data_get(v.pkt_data, v.blk_count, index)
		blks[index] = (*PktBlk)(blk)
	}
	return blks
}

func (v *Pkt) Unit() int {
	return int(v.unit)
}

func (v *Pkt) SetUnit(unit int) {
	v.unit = C.uint8(unit)
}

func (v *Pkt) Cos() Cos {
	return Cos(v.cos)
}

func (v *Pkt) SetCos(cos Cos) {
	v.cos = C.uint8(cos.C())
}

func (v *Pkt) VID() Vlan {
	return Vlan(v.vlan)
}

func (v *Pkt) SrcPort() uint16 {
	return uint16(v.src_port)
}

func (v *Pkt) DstPort() uint16 {
	return uint16(v.dest_port)
}

func (v *Pkt) PktLen() uint16 {
	return uint16(v.pkt_len)
}

func (v *Pkt) TotLen() uint16 {
	return uint16(v.tot_len)
}

func (v *Pkt) TxPBmp() *PBmp {
	return (*PBmp)(&v.tx_pbmp)
}

func (v *Pkt) TxUntagPBmp() *PBmp {
	return (*PBmp)(&v.tx_upbmp)
}

func (v *Pkt) RxReasons() *RxReasons {
	return (*RxReasons)(&v.rx_reasons)
}

func (v *Pkt) RxPort() uint8 {
	return uint8(v.rx_port)
}

func (v *Pkt) RxUntagged() uint8 {
	return uint8(v.rx_untagged)
}

func (v *Pkt) RxMatched() uint32 {
	return uint32(v.rx_matched)
}

func (v *Pkt) Flags() PktFlags {
	return PktFlags(v.flags)
}

func PktAlloc(unit int, size int, flags PktFlags) (*Pkt, error) {
	var c_pkt *C.opennsl_pkt_t = nil

	rc := C.opennsl_pkt_alloc(C.int(unit), C.int(size), flags.C(), &c_pkt)
	if err := ParseError(rc); err != nil {
		return nil, err
	}

	return (*Pkt)(c_pkt), nil
}

func (v *Pkt) Free(unit int) error {
	rc := C.opennsl_pkt_free(C.int(unit), v.C())
	return ParseError(rc)
}

func (v *Pkt) Memcpy(start int, datas []byte) error {
	dataLen := len(datas)
	if start < 0 || start >= dataLen {
		return E_PARAM.Error()
	}

	c_datas := C.CBytes(datas)
	defer C.free(c_datas)

	cpyLen := C.opennsl_pkt_memcpy(v.C(), C.int(start), (*C.uint8)(c_datas), C.int(dataLen))
	if int(cpyLen) != dataLen {
		return E_INTERNAL.Error()
	}

	return nil
}

func (v *Pkt) WriteTo(w io.Writer) (int64, error) {
	total := int(v.TotLen())
	for _, blk := range v.Blks() {
		n, err := w.Write(blk.DataN(total))
		if err != nil {
			return 0, err
		}
		total -= n
	}

	return int64(v.TotLen()), nil
}
