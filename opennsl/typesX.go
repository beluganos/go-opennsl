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
#include <opennsl/typesX.h>
*/
import "C"

//
// FieldStat
//
type FieldStat C.opennsl_field_stat_t

func (v FieldStat) C() C.opennsl_field_stat_t {
	return C.opennsl_field_stat_t(v)
}

const (
	FieldStatBytes                    FieldStat = C.opennslFieldStatBytes
	FieldStatPackets                  FieldStat = C.opennslFieldStatPackets
	FieldStatDefault                  FieldStat = C.opennslFieldStatDefault
	FieldStatGreenBytes               FieldStat = C.opennslFieldStatGreenBytes
	FieldStatGreenPackets             FieldStat = C.opennslFieldStatGreenPackets
	FieldStatYellowBytes              FieldStat = C.opennslFieldStatYellowBytes
	FieldStatYellowPackets            FieldStat = C.opennslFieldStatYellowPackets
	FieldStatRedBytes                 FieldStat = C.opennslFieldStatRedBytes
	FieldStatRedPackets               FieldStat = C.opennslFieldStatRedPackets
	FieldStatNotGreenBytes            FieldStat = C.opennslFieldStatNotGreenBytes
	FieldStatNotGreenPackets          FieldStat = C.opennslFieldStatNotGreenPackets
	FieldStatNotYellowBytes           FieldStat = C.opennslFieldStatNotYellowBytes
	FieldStatNotYellowPackets         FieldStat = C.opennslFieldStatNotYellowPackets
	FieldStatNotRedBytes              FieldStat = C.opennslFieldStatNotRedBytes
	FieldStatNotRedPacketsc           FieldStat = C.opennslFieldStatNotRedPackets
	FieldStatAcceptedBytes            FieldStat = C.opennslFieldStatAcceptedBytes
	FieldStatAcceptedPackets          FieldStat = C.opennslFieldStatAcceptedPackets
	FieldStatAcceptedGreenBytes       FieldStat = C.opennslFieldStatAcceptedGreenBytes
	FieldStatAcceptedGreenPackets     FieldStat = C.opennslFieldStatAcceptedGreenPackets
	FieldStatAcceptedYellowBytes      FieldStat = C.opennslFieldStatAcceptedYellowBytes
	FieldStatAcceptedYellowPackets    FieldStat = C.opennslFieldStatAcceptedYellowPackets
	FieldStatAcceptedRedBytes         FieldStat = C.opennslFieldStatAcceptedRedBytes
	FieldStatAcceptedRedPackets       FieldStat = C.opennslFieldStatAcceptedRedPackets
	FieldStatAcceptedNotGreenBytes    FieldStat = C.opennslFieldStatAcceptedNotGreenBytes
	FieldStatAcceptedNotGreenPackets  FieldStat = C.opennslFieldStatAcceptedNotGreenPackets
	FieldStatAcceptedNotYellowBytes   FieldStat = C.opennslFieldStatAcceptedNotYellowBytes
	FieldStatAcceptedNotYellowPackets FieldStat = C.opennslFieldStatAcceptedNotYellowPackets
	FieldStatAcceptedNotRedBytes      FieldStat = C.opennslFieldStatAcceptedNotRedBytes
	FieldStatAcceptedNotRedPackets    FieldStat = C.opennslFieldStatAcceptedNotRedPackets
	FieldStatDroppedBytes             FieldStat = C.opennslFieldStatDroppedBytes
	FieldStatDroppedPackets           FieldStat = C.opennslFieldStatDroppedPackets
	FieldStatDroppedGreenBytes        FieldStat = C.opennslFieldStatDroppedGreenBytes
	FieldStatDroppedGreenPackets      FieldStat = C.opennslFieldStatDroppedGreenPackets
	FieldStatDroppedYellowBytes       FieldStat = C.opennslFieldStatDroppedYellowBytes
	FieldStatDroppedYellowPackets     FieldStat = C.opennslFieldStatDroppedYellowPackets
	FieldStatDroppedRedBytes          FieldStat = C.opennslFieldStatDroppedRedBytes
	FieldStatDroppedRedPackets        FieldStat = C.opennslFieldStatDroppedRedPackets
	FieldStatDroppedNotGreenBytes     FieldStat = C.opennslFieldStatDroppedNotGreenBytes
	FieldStatDroppedNotGreenPackets   FieldStat = C.opennslFieldStatDroppedNotGreenPackets
	FieldStatDroppedNotYellowBytes    FieldStat = C.opennslFieldStatDroppedNotYellowBytes
	FieldStatDroppedNotYellowPackets  FieldStat = C.opennslFieldStatDroppedNotYellowPackets
	FieldStatDroppedNotRedBytes       FieldStat = C.opennslFieldStatDroppedNotRedBytes
	FieldStatDroppedNotRedPackets     FieldStat = C.opennslFieldStatDroppedNotRedPackets
	FieldStatOffset0Bytes             FieldStat = C.opennslFieldStatOffset0Bytes
	FieldStatOffset0Packets           FieldStat = C.opennslFieldStatOffset0Packets
	FieldStatOffset1Bytes             FieldStat = C.opennslFieldStatOffset1Bytes
	FieldStatOffset1Packets           FieldStat = C.opennslFieldStatOffset1Packets
	FieldStatOffset2Bytes             FieldStat = C.opennslFieldStatOffset2Bytes
	FieldStatOffset2Packets           FieldStat = C.opennslFieldStatOffset2Packets
	FieldStatOffset3Bytes             FieldStat = C.opennslFieldStatOffset3Bytes
	FieldStatOffset3Packets           FieldStat = C.opennslFieldStatOffset3Packets
	FieldStatOffset4Bytes             FieldStat = C.opennslFieldStatOffset4Bytes
	FieldStatOffset4Packets           FieldStat = C.opennslFieldStatOffset4Packets
	FieldStatOffset5Bytes             FieldStat = C.opennslFieldStatOffset5Bytes
	FieldStatOffset5Packets           FieldStat = C.opennslFieldStatOffset5Packets
	FieldStatOffset6Bytes             FieldStat = C.opennslFieldStatOffset6Bytes
	FieldStatOffset6Packets           FieldStat = C.opennslFieldStatOffset6Packets
	FieldStatOffset7Bytes             FieldStat = C.opennslFieldStatOffset7Bytes
	FieldStatOffset7Packets           FieldStat = C.opennslFieldStatOffset7Packets
	FieldStatCount                    FieldStat = C.opennslFieldStatCount
)
