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
#include <opennsl/stat.h>
*/
import "C"

import (
	"fmt"
)

//
// StatVal
//
type StatVal C.opennsl_stat_val_t

func (v StatVal) C() C.opennsl_stat_val_t {
	return C.opennsl_stat_val_t(v)
}

const (
	SPLSnmpIfInOctets                         StatVal = C.opennsl_spl_snmpIfInOctets
	SPLSnmpIfInUcastPkts                      StatVal = C.opennsl_spl_snmpIfInUcastPkts
	SPLSnmpIfInNUcastPkts                     StatVal = C.opennsl_spl_snmpIfInNUcastPkts
	SPLSnmpIfInDiscards                       StatVal = C.opennsl_spl_snmpIfInDiscards
	SPLSnmpIfInErrors                         StatVal = C.opennsl_spl_snmpIfInErrors
	SPLSnmpIfInUnknownProtos                  StatVal = C.opennsl_spl_snmpIfInUnknownProtos
	SPLSnmpIfOutOctets                        StatVal = C.opennsl_spl_snmpIfOutOctets
	SPLSnmpIfOutUcastPkts                     StatVal = C.opennsl_spl_snmpIfOutUcastPkts
	SPLSnmpIfOutNUcastPkts                    StatVal = C.opennsl_spl_snmpIfOutNUcastPkts
	SPLSnmpIfOutDiscards                      StatVal = C.opennsl_spl_snmpIfOutDiscards
	SPLSnmpIfOutErrors                        StatVal = C.opennsl_spl_snmpIfOutErrors
	SPLSnmpIfOutQLen                          StatVal = C.opennsl_spl_snmpIfOutQLen
	SPLSnmpIpInReceives                       StatVal = C.opennsl_spl_snmpIpInReceives
	SPLSnmpIpInHdrErrors                      StatVal = C.opennsl_spl_snmpIpInHdrErrors
	SPLSnmpIpForwDatagrams                    StatVal = C.opennsl_spl_snmpIpForwDatagrams
	SPLSnmpIpInDiscards                       StatVal = C.opennsl_spl_snmpIpInDiscards
	SPLSnmpDot1dBasePortDelayExceededDiscards StatVal = C.opennsl_spl_snmpDot1dBasePortDelayExceededDiscards
	SPLSnmpDot1dBasePortMtuExceededDiscards   StatVal = C.opennsl_spl_snmpDot1dBasePortMtuExceededDiscards
	SPLSnmpDot1dTpPortInFrames                StatVal = C.opennsl_spl_snmpDot1dTpPortInFrames
	SPLSnmpDot1dTpPortOutFrames               StatVal = C.opennsl_spl_snmpDot1dTpPortOutFrames
	SPLSnmpDot1dPortInDiscards                StatVal = C.opennsl_spl_snmpDot1dPortInDiscards
	SPLSnmpEtherStatsDropEvents               StatVal = C.opennsl_spl_snmpEtherStatsDropEvents
	SPLSnmpEtherStatsMulticastPkts            StatVal = C.opennsl_spl_snmpEtherStatsMulticastPkts
	SPLSnmpEtherStatsBroadcastPkts            StatVal = C.opennsl_spl_snmpEtherStatsBroadcastPkts
	SPLSnmpEtherStatsUndersizePkts            StatVal = C.opennsl_spl_snmpEtherStatsUndersizePkts
	SPLSnmpEtherStatsFragments                StatVal = C.opennsl_spl_snmpEtherStatsFragments
	SPLSnmpEtherStatsPkts64Octets             StatVal = C.opennsl_spl_snmpEtherStatsPkts64Octets
	SPLSnmpEtherStatsPkts65to127Octets        StatVal = C.opennsl_spl_snmpEtherStatsPkts65to127Octets
	SPLSnmpEtherStatsPkts128to255Octets       StatVal = C.opennsl_spl_snmpEtherStatsPkts128to255Octets
	SPLSnmpEtherStatsPkts256to511Octets       StatVal = C.opennsl_spl_snmpEtherStatsPkts256to511Octets
	SPLSnmpEtherStatsPkts512to1023Octets      StatVal = C.opennsl_spl_snmpEtherStatsPkts512to1023Octets
	SPLSnmpEtherStatsPkts1024to1518Octets     StatVal = C.opennsl_spl_snmpEtherStatsPkts1024to1518Octets
	SPLSnmpEtherStatsOversizePkts             StatVal = C.opennsl_spl_snmpEtherStatsOversizePkts
	SPLSnmpEtherRxOversizePkts                StatVal = C.opennsl_spl_snmpEtherRxOversizePkts
	SPLSnmpEtherTxOversizePkts                StatVal = C.opennsl_spl_snmpEtherTxOversizePkts
	SPLSnmpEtherStatsJabbers                  StatVal = C.opennsl_spl_snmpEtherStatsJabbers
	SPLSnmpEtherStatsOctets                   StatVal = C.opennsl_spl_snmpEtherStatsOctets
	SPLSnmpEtherStatsPkts                     StatVal = C.opennsl_spl_snmpEtherStatsPkts
	SPLSnmpEtherStatsCollisions               StatVal = C.opennsl_spl_snmpEtherStatsCollisions
	SPLSnmpEtherStatsCRCAlignErrors           StatVal = C.opennsl_spl_snmpEtherStatsCRCAlignErrors
	SPLSnmpEtherStatsTXNoErrors               StatVal = C.opennsl_spl_snmpEtherStatsTXNoErrors
	SPLSnmpEtherStatsRXNoErrors               StatVal = C.opennsl_spl_snmpEtherStatsRXNoErrors
	SPLSnmpDot3StatsAlignmentErrors           StatVal = C.opennsl_spl_snmpDot3StatsAlignmentErrors
	SPLSnmpDot3StatsFCSErrors                 StatVal = C.opennsl_spl_snmpDot3StatsFCSErrors
	SPLSnmpDot3StatsSingleCollisionFrames     StatVal = C.opennsl_spl_snmpDot3StatsSingleCollisionFrames
	SPLSnmpDot3StatsMultipleCollisionFrames   StatVal = C.opennsl_spl_snmpDot3StatsMultipleCollisionFrames
	SPLSnmpDot3StatsSQETTestErrors            StatVal = C.opennsl_spl_snmpDot3StatsSQETTestErrors
	SPLSnmpDot3StatsDeferredTransmissions     StatVal = C.opennsl_spl_snmpDot3StatsDeferredTransmissions
	SPLSnmpDot3StatsLateCollisions            StatVal = C.opennsl_spl_snmpDot3StatsLateCollisions
	SPLSnmpDot3StatsExcessiveCollisions       StatVal = C.opennsl_spl_snmpDot3StatsExcessiveCollisions
	SPLSnmpDot3StatsInternalMacTransmitErrors StatVal = C.opennsl_spl_snmpDot3StatsInternalMacTransmitErrors
	SPLSnmpDot3StatsCarrierSenseErrors        StatVal = C.opennsl_spl_snmpDot3StatsCarrierSenseErrors
	SPLSnmpDot3StatsFrameTooLongs             StatVal = C.opennsl_spl_snmpDot3StatsFrameTooLongs
	SPLSnmpDot3StatsInternalMacReceiveErrors  StatVal = C.opennsl_spl_snmpDot3StatsInternalMacReceiveErrors
	SPLSnmpDot3StatsSymbolErrors              StatVal = C.opennsl_spl_snmpDot3StatsSymbolErrors
	SPLSnmpDot3ControlInUnknownOpcodes        StatVal = C.opennsl_spl_snmpDot3ControlInUnknownOpcodes
	SPLSnmpDot3InPauseFrames                  StatVal = C.opennsl_spl_snmpDot3InPauseFrames
	SPLSnmpDot3OutPauseFrames                 StatVal = C.opennsl_spl_snmpDot3OutPauseFrames
	SPLSnmpIfHCInOctets                       StatVal = C.opennsl_spl_snmpIfHCInOctets
	SPLSnmpIfHCInUcastPkts                    StatVal = C.opennsl_spl_snmpIfHCInUcastPkts
	SPLSnmpIfHCInMulticastPkts                StatVal = C.opennsl_spl_snmpIfHCInMulticastPkts
	SPLSnmpIfHCInBroadcastPkts                StatVal = C.opennsl_spl_snmpIfHCInBroadcastPkts
	SPLSnmpIfHCOutOctets                      StatVal = C.opennsl_spl_snmpIfHCOutOctets
	SPLSnmpIfHCOutUcastPkts                   StatVal = C.opennsl_spl_snmpIfHCOutUcastPkts
	SPLSnmpIfHCOutMulticastPkts               StatVal = C.opennsl_spl_snmpIfHCOutMulticastPkts
	SPLSnmpIfHCOutBroadcastPckts              StatVal = C.opennsl_spl_snmpIfHCOutBroadcastPckts
	SPLSnmpIpv6IfStatsInReceives              StatVal = C.opennsl_spl_snmpIpv6IfStatsInReceives
	SPLSnmpIpv6IfStatsInHdrErrors             StatVal = C.opennsl_spl_snmpIpv6IfStatsInHdrErrors
	SPLSnmpIpv6IfStatsInAddrErrors            StatVal = C.opennsl_spl_snmpIpv6IfStatsInAddrErrors
	SPLSnmpIpv6IfStatsInDiscards              StatVal = C.opennsl_spl_snmpIpv6IfStatsInDiscards
	SPLSnmpIpv6IfStatsOutForwDatagrams        StatVal = C.opennsl_spl_snmpIpv6IfStatsOutForwDatagrams
	SPLSnmpIpv6IfStatsOutDiscards             StatVal = C.opennsl_spl_snmpIpv6IfStatsOutDiscards
	SPLSnmpIpv6IfStatsInMcastPkts             StatVal = C.opennsl_spl_snmpIpv6IfStatsInMcastPkts
	SPLSnmpIpv6IfStatsOutMcastPkts            StatVal = C.opennsl_spl_snmpIpv6IfStatsOutMcastPkts
	SPLSnmpIfInBroadcastPkts                  StatVal = C.opennsl_spl_snmpIfInBroadcastPkts
	SPLSnmpIfInMulticastPkts                  StatVal = C.opennsl_spl_snmpIfInMulticastPkts
	SPLSnmpIfOutBroadcastPkts                 StatVal = C.opennsl_spl_snmpIfOutBroadcastPkts
	SPLSnmpIfOutMulticastPkts                 StatVal = C.opennsl_spl_snmpIfOutMulticastPkts
	SPLSnmpIeee8021PfcRequests                StatVal = C.opennsl_spl_snmpIeee8021PfcRequests
	SPLSnmpIeee8021PfcIndications             StatVal = C.opennsl_spl_snmpIeee8021PfcIndications
	NSLSnmpReceivedUndersizePkts              StatVal = C.snmpOpenNSLReceivedUndersizePkts
	NSLSnmpTransmittedUndersizePkts           StatVal = C.snmpOpenNSLTransmittedUndersizePkts
	NSLSnmpIPMCBridgedPckts                   StatVal = C.snmpOpenNSLIPMCBridgedPckts
	NSLSnmpIPMCRoutedPckts                    StatVal = C.snmpOpenNSLIPMCRoutedPckts
	NSLSnmpIPMCInDroppedPckts                 StatVal = C.snmpOpenNSLIPMCInDroppedPckts
	NSLSnmpIPMCOutDroppedPckts                StatVal = C.snmpOpenNSLIPMCOutDroppedPckts
	NSLSnmpEtherStatsPkts1519to1522Octets     StatVal = C.snmpOpenNSLEtherStatsPkts1519to1522Octets
	NSLSnmpEtherStatsPkts1522to2047Octets     StatVal = C.snmpOpenNSLEtherStatsPkts1522to2047Octets
	NSLSnmpEtherStatsPkts2048to4095Octets     StatVal = C.snmpOpenNSLEtherStatsPkts2048to4095Octets
	NSLSnmpEtherStatsPkts4095to9216Octets     StatVal = C.snmpOpenNSLEtherStatsPkts4095to9216Octets
	NSLSnmpReceivedPkts64Octets               StatVal = C.snmpOpenNSLReceivedPkts64Octets
	NSLSnmpReceivedPkts65to127Octets          StatVal = C.snmpOpenNSLReceivedPkts65to127Octets
	NSLSnmpReceivedPkts128to255Octets         StatVal = C.snmpOpenNSLReceivedPkts128to255Octets
	NSLSnmpReceivedPkts256to511Octets         StatVal = C.snmpOpenNSLReceivedPkts256to511Octets
	NSLSnmpReceivedPkts512to1023Octets        StatVal = C.snmpOpenNSLReceivedPkts512to1023Octets
	NSLSnmpReceivedPkts1024to1518Octets       StatVal = C.snmpOpenNSLReceivedPkts1024to1518Octets
	NSLSnmpReceivedPkts1519to2047Octets       StatVal = C.snmpOpenNSLReceivedPkts1519to2047Octets
	NSLSnmpReceivedPkts2048to4095Octets       StatVal = C.snmpOpenNSLReceivedPkts2048to4095Octets
	NSLSnmpReceivedPkts4095to9216Octets       StatVal = C.snmpOpenNSLReceivedPkts4095to9216Octets
	NSLSnmpTransmittedPkts64Octets            StatVal = C.snmpOpenNSLTransmittedPkts64Octets
	NSLSnmpTransmittedPkts65to127Octets       StatVal = C.snmpOpenNSLTransmittedPkts65to127Octets
	NSLSnmpTransmittedPkts128to255Octets      StatVal = C.snmpOpenNSLTransmittedPkts128to255Octets
	NSLSnmpTransmittedPkts256to511Octets      StatVal = C.snmpOpenNSLTransmittedPkts256to511Octets
	NSLSnmpTransmittedPkts512to1023Octets     StatVal = C.snmpOpenNSLTransmittedPkts512to1023Octets
	NSLSnmpTransmittedPkts1024to1518Octets    StatVal = C.snmpOpenNSLTransmittedPkts1024to1518Octets
	NSLSnmpTransmittedPkts1519to2047Octets    StatVal = C.snmpOpenNSLTransmittedPkts1519to2047Octets
	NSLSnmpTransmittedPkts2048to4095Octets    StatVal = C.snmpOpenNSLTransmittedPkts2048to4095Octets
	NSLSnmpTransmittedPkts4095to9216Octets    StatVal = C.snmpOpenNSLTransmittedPkts4095to9216Octets
	NSLSnmpTxControlCells                     StatVal = C.snmpOpenNSLTxControlCells
	NSLSnmpTxDataCells                        StatVal = C.snmpOpenNSLTxDataCells
	NSLSnmpTxDataBytes                        StatVal = C.snmpOpenNSLTxDataBytes
	NSLSnmpRxCrcErrors                        StatVal = C.snmpOpenNSLRxCrcErrors
	NSLSnmpRxFecCorrectable                   StatVal = C.snmpOpenNSLRxFecCorrectable
	NSLSnmpRxBecCrcErrors                     StatVal = C.snmpOpenNSLRxBecCrcErrors
	NSLSnmpRxDisparityErrors                  StatVal = C.snmpOpenNSLRxDisparityErrors
	NSLSnmpRxControlCells                     StatVal = C.snmpOpenNSLRxControlCells
	NSLSnmpRxDataCells                        StatVal = C.snmpOpenNSLRxDataCells
	NSLSnmpRxDataBytes                        StatVal = C.snmpOpenNSLRxDataBytes
	NSLSnmpRxDroppedRetransmittedControl      StatVal = C.snmpOpenNSLRxDroppedRetransmittedControl
	NSLSnmpTxBecRetransmit                    StatVal = C.snmpOpenNSLTxBecRetransmit
	NSLSnmpRxBecRetransmit                    StatVal = C.snmpOpenNSLRxBecRetransmit
	NSLSnmpTxAsynFifoRate                     StatVal = C.snmpOpenNSLTxAsynFifoRate
	NSLSnmpRxAsynFifoRate                     StatVal = C.snmpOpenNSLRxAsynFifoRate
	NSLSnmpRxFecUncorrectable                 StatVal = C.snmpOpenNSLRxFecUncorrectable
	NSLSnmpRxBecRxFault                       StatVal = C.snmpOpenNSLRxBecRxFault
	NSLSnmpRxCodeErrors                       StatVal = C.snmpOpenNSLRxCodeErrors
	NSLSnmpRxRsFecBitError                    StatVal = C.snmpOpenNSLRxRsFecBitError
	NSLSnmpRxRsFecSymbolError                 StatVal = C.snmpOpenNSLRxRsFecSymbolError
	NSLSnmpRxLlfcPrimary                      StatVal = C.snmpOpenNSLRxLlfcPrimary
	NSLSnmpRxLlfcSecondary                    StatVal = C.snmpOpenNSLRxLlfcSecondary
	NSLSnmpCustomReceive0                     StatVal = C.snmpOpenNSLCustomReceive0
	NSLSnmpCustomReceive1                     StatVal = C.snmpOpenNSLCustomReceive1
	NSLSnmpCustomReceive2                     StatVal = C.snmpOpenNSLCustomReceive2
	NSLSnmpCustomReceive3                     StatVal = C.snmpOpenNSLCustomReceive3
	NSLSnmpCustomReceive4                     StatVal = C.snmpOpenNSLCustomReceive4
	NSLSnmpCustomReceive5                     StatVal = C.snmpOpenNSLCustomReceive5
	NSLSnmpCustomReceive6                     StatVal = C.snmpOpenNSLCustomReceive6
	NSLSnmpCustomReceive7                     StatVal = C.snmpOpenNSLCustomReceive7
	NSLSnmpCustomReceive8                     StatVal = C.snmpOpenNSLCustomReceive8
	NSLSnmpCustomTransmit0                    StatVal = C.snmpOpenNSLCustomTransmit0
	NSLSnmpCustomTransmit1                    StatVal = C.snmpOpenNSLCustomTransmit1
	NSLSnmpCustomTransmit2                    StatVal = C.snmpOpenNSLCustomTransmit2
	NSLSnmpCustomTransmit3                    StatVal = C.snmpOpenNSLCustomTransmit3
	NSLSnmpCustomTransmit4                    StatVal = C.snmpOpenNSLCustomTransmit4
	NSLSnmpCustomTransmit5                    StatVal = C.snmpOpenNSLCustomTransmit5
	NSLSnmpCustomTransmit6                    StatVal = C.snmpOpenNSLCustomTransmit6
	NSLSnmpCustomTransmit7                    StatVal = C.snmpOpenNSLCustomTransmit7
	NSLSnmpCustomTransmit8                    StatVal = C.snmpOpenNSLCustomTransmit8
	NSLSnmpCustomTransmit9                    StatVal = C.snmpOpenNSLCustomTransmit9
	NSLSnmpCustomTransmit10                   StatVal = C.snmpOpenNSLCustomTransmit10
	NSLSnmpCustomTransmit11                   StatVal = C.snmpOpenNSLCustomTransmit11
	NSLSnmpCustomTransmit12                   StatVal = C.snmpOpenNSLCustomTransmit12
	NSLSnmpCustomTransmit13                   StatVal = C.snmpOpenNSLCustomTransmit13
	NSLSnmpCustomTransmit14                   StatVal = C.snmpOpenNSLCustomTransmit14
	SPLSnmpDot3StatsInRangeLengthError        StatVal = C.opennsl_spl_snmpDot3StatsInRangeLengthError
	SPLSnmpDot3OmpEmulationCRC8Errors         StatVal = C.opennsl_spl_snmpDot3OmpEmulationCRC8Errors
	SPLSnmpDot3MpcpRxGate                     StatVal = C.opennsl_spl_snmpDot3MpcpRxGate
	SPLSnmpDot3MpcpRxRegister                 StatVal = C.opennsl_spl_snmpDot3MpcpRxRegister
	SPLSnmpDot3MpcpTxRegRequest               StatVal = C.opennsl_spl_snmpDot3MpcpTxRegRequest
	SPLSnmpDot3MpcpTxRegAck                   StatVal = C.opennsl_spl_snmpDot3MpcpTxRegAck
	SPLSnmpDot3MpcpTxReport                   StatVal = C.opennsl_spl_snmpDot3MpcpTxReport
	SPLSnmpDot3EponFecCorrectedBlocks         StatVal = C.opennsl_spl_snmpDot3EponFecCorrectedBlocks
	SPLSnmpDot3EponFecUncorrectableBlocks     StatVal = C.opennsl_spl_snmpDot3EponFecUncorrectableBlocks
	NSLSnmpPonInDroppedOctets                 StatVal = C.snmpOpenNSLPonInDroppedOctets
	NSLSnmpPonOutDroppedOctets                StatVal = C.snmpOpenNSLPonOutDroppedOctets
	NSLSnmpPonInDelayedOctets                 StatVal = C.snmpOpenNSLPonInDelayedOctets
	NSLSnmpPonOutDelayedOctets                StatVal = C.snmpOpenNSLPonOutDelayedOctets
	NSLSnmpPonInDelayedHundredUs              StatVal = C.snmpOpenNSLPonInDelayedHundredUs
	NSLSnmpPonOutDelayedHundredUs             StatVal = C.snmpOpenNSLPonOutDelayedHundredUs
	NSLSnmpPonInFrameErrors                   StatVal = C.snmpOpenNSLPonInFrameErrors
	NSLSnmpPonInOamFrames                     StatVal = C.snmpOpenNSLPonInOamFrames
	NSLSnmpPonOutOamFrames                    StatVal = C.snmpOpenNSLPonOutOamFrames
	NSLSnmpPonOutUnusedOctets                 StatVal = C.snmpOpenNSLPonOutUnusedOctets
	NSLSnmpEtherStatsPkts9217to16383Octets    StatVal = C.snmpOpenNSLEtherStatsPkts9217to16383Octets
	NSLSnmpReceivedPkts9217to16383Octets      StatVal = C.snmpOpenNSLReceivedPkts9217to16383Octets
	NSLSnmpTransmittedPkts9217to16383Octets   StatVal = C.snmpOpenNSLTransmittedPkts9217to16383Octets
	NSLSnmpRxVlanTagFrame                     StatVal = C.snmpOpenNSLRxVlanTagFrame
	NSLSnmpRxDoubleVlanTagFrame               StatVal = C.snmpOpenNSLRxDoubleVlanTagFrame
	NSLSnmpTxVlanTagFrame                     StatVal = C.snmpOpenNSLTxVlanTagFrame
	NSLSnmpTxDoubleVlanTagFrame               StatVal = C.snmpOpenNSLTxDoubleVlanTagFrame
	NSLSnmpRxPFCControlFrame                  StatVal = C.snmpOpenNSLRxPFCControlFrame
	NSLSnmpTxPFCControlFrame                  StatVal = C.snmpOpenNSLTxPFCControlFrame
	NSLSnmpRxPFCFrameXonPriority0             StatVal = C.snmpOpenNSLRxPFCFrameXonPriority0
	NSLSnmpRxPFCFrameXonPriority1             StatVal = C.snmpOpenNSLRxPFCFrameXonPriority1
	NSLSnmpRxPFCFrameXonPriority2             StatVal = C.snmpOpenNSLRxPFCFrameXonPriority2
	NSLSnmpRxPFCFrameXonPriority3             StatVal = C.snmpOpenNSLRxPFCFrameXonPriority3
	NSLSnmpRxPFCFrameXonPriority4             StatVal = C.snmpOpenNSLRxPFCFrameXonPriority4
	NSLSnmpRxPFCFrameXonPriority5             StatVal = C.snmpOpenNSLRxPFCFrameXonPriority5
	NSLSnmpRxPFCFrameXonPriority6             StatVal = C.snmpOpenNSLRxPFCFrameXonPriority6
	NSLSnmpRxPFCFrameXonPriority7             StatVal = C.snmpOpenNSLRxPFCFrameXonPriority7
	NSLSnmpRxPFCFramePriority0                StatVal = C.snmpOpenNSLRxPFCFramePriority0
	NSLSnmpRxPFCFramePriority1                StatVal = C.snmpOpenNSLRxPFCFramePriority1
	NSLSnmpRxPFCFramePriority2                StatVal = C.snmpOpenNSLRxPFCFramePriority2
	NSLSnmpRxPFCFramePriority3                StatVal = C.snmpOpenNSLRxPFCFramePriority3
	NSLSnmpRxPFCFramePriority4                StatVal = C.snmpOpenNSLRxPFCFramePriority4
	NSLSnmpRxPFCFramePriority5                StatVal = C.snmpOpenNSLRxPFCFramePriority5
	NSLSnmpRxPFCFramePriority6                StatVal = C.snmpOpenNSLRxPFCFramePriority6
	NSLSnmpRxPFCFramePriority7                StatVal = C.snmpOpenNSLRxPFCFramePriority7
	NSLSnmpTxPFCFramePriority0                StatVal = C.snmpOpenNSLTxPFCFramePriority0
	NSLSnmpTxPFCFramePriority1                StatVal = C.snmpOpenNSLTxPFCFramePriority1
	NSLSnmpTxPFCFramePriority2                StatVal = C.snmpOpenNSLTxPFCFramePriority2
	NSLSnmpTxPFCFramePriority3                StatVal = C.snmpOpenNSLTxPFCFramePriority3
	NSLSnmpTxPFCFramePriority4                StatVal = C.snmpOpenNSLTxPFCFramePriority4
	NSLSnmpTxPFCFramePriority5                StatVal = C.snmpOpenNSLTxPFCFramePriority5
	NSLSnmpTxPFCFramePriority6                StatVal = C.snmpOpenNSLTxPFCFramePriority6
	NSLSnmpTxPFCFramePriority7                StatVal = C.snmpOpenNSLTxPFCFramePriority7
	SPLSnmpFcmPortClass3RxFrames              StatVal = C.opennsl_spl_snmpFcmPortClass3RxFrames
	SPLSnmpFcmPortClass3TxFrames              StatVal = C.opennsl_spl_snmpFcmPortClass3TxFrames
	SPLSnmpFcmPortClass3Discards              StatVal = C.opennsl_spl_snmpFcmPortClass3Discards
	SPLSnmpFcmPortClass2RxFrames              StatVal = C.opennsl_spl_snmpFcmPortClass2RxFrames
	SPLSnmpFcmPortClass2TxFrames              StatVal = C.opennsl_spl_snmpFcmPortClass2TxFrames
	SPLSnmpFcmPortClass2Discards              StatVal = C.opennsl_spl_snmpFcmPortClass2Discards
	SPLSnmpFcmPortInvalidCRCs                 StatVal = C.opennsl_spl_snmpFcmPortInvalidCRCs
	SPLSnmpFcmPortDelimiterErrors             StatVal = C.opennsl_spl_snmpFcmPortDelimiterErrors
	NSLSnmpSampleIngressPkts                  StatVal = C.snmpOpenNSLSampleIngressPkts
	NSLSnmpSampleIngressSnapshotPkts          StatVal = C.snmpOpenNSLSampleIngressSnapshotPkts
	NSLSnmpSampleIngressSampledPkts           StatVal = C.snmpOpenNSLSampleIngressSampledPkts
	NSLSnmpSampleFlexPkts                     StatVal = C.snmpOpenNSLSampleFlexPkts
	NSLSnmpSampleFlexSnapshotPkts             StatVal = C.snmpOpenNSLSampleFlexSnapshotPkts
	NSLSnmpSampleFlexSampledPkts              StatVal = C.snmpOpenNSLSampleFlexSampledPkts
	NSLSnmpEgressProtectionDataDrop           StatVal = C.snmpOpenNSLEgressProtectionDataDrop
	NSLSnmpTxE2ECCControlFrames               StatVal = C.snmpOpenNSLTxE2ECCControlFrames
	NSLSnmpE2EHOLDropPkts                     StatVal = C.snmpOpenNSLE2EHOLDropPkts
	SPLSnmpEtherStatsTxCRCAlignErrors         StatVal = C.opennsl_spl_snmpEtherStatsTxCRCAlignErrors
	SPLSnmpEtherStatsTxJabbers                StatVal = C.opennsl_spl_snmpEtherStatsTxJabbers
	NSLSnmpMacMergeTxFrag                     StatVal = C.snmpOpenNSLMacMergeTxFrag
	NSLSnmpMacMergeTxVerifyFrame              StatVal = C.snmpOpenNSLMacMergeTxVerifyFrame
	NSLSnmpMacMergeTxReplyFrame               StatVal = C.snmpOpenNSLMacMergeTxReplyFrame
	NSLSnmpMacMergeRxFrameAssErrors           StatVal = C.snmpOpenNSLMacMergeRxFrameAssErrors
	NSLSnmpMacMergeRxFrameSmdErrors           StatVal = C.snmpOpenNSLMacMergeRxFrameSmdErrors
	NSLSnmpMacMergeRxFrameAss                 StatVal = C.snmpOpenNSLMacMergeRxFrameAss
	NSLSnmpMacMergeRxFrag                     StatVal = C.snmpOpenNSLMacMergeRxFrag
	NSLSnmpMacMergeRxVerifyFrame              StatVal = C.snmpOpenNSLMacMergeRxVerifyFrame
	NSLSnmpMacMergeRxReplyFrame               StatVal = C.snmpOpenNSLMacMergeRxReplyFrame
	NSLSnmpMacMergeRxFinalFragSizeError       StatVal = C.snmpOpenNSLMacMergeRxFinalFragSizeError
	NSLSnmpMacMergeRxFragSizeError            StatVal = C.snmpOpenNSLMacMergeRxFragSizeError
	NSLSnmpMacMergeRxDiscard                  StatVal = C.snmpOpenNSLMacMergeRxDiscard
	NSLSnmpMacMergeHoldCount                  StatVal = C.snmpOpenNSLMacMergeHoldCount
	NSLSnmpRxBipErrorCount                    StatVal = C.snmpOpenNSLRxBipErrorCount
)

const SPLSnmpValCount = C.opennsl_spl_snmpValCount

func (v StatVal) String() string {
	if s, ok := statVal_names[v]; ok {
		return s
	}
	return fmt.Sprintf("StatVal(%d)", v)
}

var statVal_names = map[StatVal]string{
	SPLSnmpIfInOctets:                         "ifInOctets",
	SPLSnmpIfInUcastPkts:                      "ifInUcastPkts",
	SPLSnmpIfInNUcastPkts:                     "ifInNUcastPkts",
	SPLSnmpIfInDiscards:                       "ifInDiscards",
	SPLSnmpIfInErrors:                         "ifInErrors",
	SPLSnmpIfInUnknownProtos:                  "ifInUnknownProtos",
	SPLSnmpIfOutOctets:                        "ifOutOctets",
	SPLSnmpIfOutUcastPkts:                     "ifOutUcastPkts",
	SPLSnmpIfOutNUcastPkts:                    "ifOutNUcastPkts",
	SPLSnmpIfOutDiscards:                      "ifOutDiscards",
	SPLSnmpIfOutErrors:                        "ifOutErrors",
	SPLSnmpIfOutQLen:                          "ifOutQLen",
	SPLSnmpIpInReceives:                       "ipInReceives",
	SPLSnmpIpInHdrErrors:                      "ipInHdrErrors",
	SPLSnmpIpForwDatagrams:                    "ipForwDatagrams",
	SPLSnmpIpInDiscards:                       "ipInDiscards",
	SPLSnmpDot1dBasePortDelayExceededDiscards: "dot1dBasePortDelayExceededDiscards",
	SPLSnmpDot1dBasePortMtuExceededDiscards:   "dot1dBasePortMtuExceededDiscards",
	SPLSnmpDot1dTpPortInFrames:                "dot1dTpPortInFrames",
	SPLSnmpDot1dTpPortOutFrames:               "dot1dTpPortOutFrames",
	SPLSnmpDot1dPortInDiscards:                "dot1dPortInDiscards",
	SPLSnmpEtherStatsDropEvents:               "etherStatsDropEvents",
	SPLSnmpEtherStatsMulticastPkts:            "etherStatsMulticastPkts",
	SPLSnmpEtherStatsBroadcastPkts:            "etherStatsBroadcastPkts",
	SPLSnmpEtherStatsUndersizePkts:            "etherStatsUndersizePkts",
	SPLSnmpEtherStatsFragments:                "etherStatsFragments",
	SPLSnmpEtherStatsPkts64Octets:             "etherStatsPkts64Octets",
	SPLSnmpEtherStatsPkts65to127Octets:        "etherStatsPkts65to127Octets",
	SPLSnmpEtherStatsPkts128to255Octets:       "etherStatsPkts128to255Octets",
	SPLSnmpEtherStatsPkts256to511Octets:       "etherStatsPkts256to511Octets",
	SPLSnmpEtherStatsPkts512to1023Octets:      "etherStatsPkts512to1023Octets",
	SPLSnmpEtherStatsPkts1024to1518Octets:     "etherStatsPkts1024to1518Octets",
	SPLSnmpEtherStatsOversizePkts:             "etherStatsOversizePkts",
	SPLSnmpEtherRxOversizePkts:                "etherRxOversizePkts",
	SPLSnmpEtherTxOversizePkts:                "etherTxOversizePkts",
	SPLSnmpEtherStatsJabbers:                  "etherStatsJabbers",
	SPLSnmpEtherStatsOctets:                   "etherStatsOctets",
	SPLSnmpEtherStatsPkts:                     "etherStatsPkts",
	SPLSnmpEtherStatsCollisions:               "etherStatsCollisions",
	SPLSnmpEtherStatsCRCAlignErrors:           "etherStatsCRCAlignErrors",
	SPLSnmpEtherStatsTXNoErrors:               "etherStatsTXNoErrors",
	SPLSnmpEtherStatsRXNoErrors:               "etherStatsRXNoErrors",
	SPLSnmpDot3StatsAlignmentErrors:           "dot3StatsAlignmentErrors",
	SPLSnmpDot3StatsFCSErrors:                 "dot3StatsFCSErrors",
	SPLSnmpDot3StatsSingleCollisionFrames:     "dot3StatsSingleCollisionFrames",
	SPLSnmpDot3StatsMultipleCollisionFrames:   "dot3StatsMultipleCollisionFrames",
	SPLSnmpDot3StatsSQETTestErrors:            "dot3StatsSQETTestErrors",
	SPLSnmpDot3StatsDeferredTransmissions:     "dot3StatsDeferredTransmissions",
	SPLSnmpDot3StatsLateCollisions:            "dot3StatsLateCollisions",
	SPLSnmpDot3StatsExcessiveCollisions:       "dot3StatsExcessiveCollisions",
	SPLSnmpDot3StatsInternalMacTransmitErrors: "dot3StatsInternalMacTransmitErrors",
	SPLSnmpDot3StatsCarrierSenseErrors:        "dot3StatsCarrierSenseErrors",
	SPLSnmpDot3StatsFrameTooLongs:             "dot3StatsFrameTooLongs",
	SPLSnmpDot3StatsInternalMacReceiveErrors:  "dot3StatsInternalMacReceiveErrors",
	SPLSnmpDot3StatsSymbolErrors:              "dot3StatsSymbolErrors",
	SPLSnmpDot3ControlInUnknownOpcodes:        "dot3ControlInUnknownOpcodes",
	SPLSnmpDot3InPauseFrames:                  "dot3InPauseFrames",
	SPLSnmpDot3OutPauseFrames:                 "dot3OutPauseFrames",
	SPLSnmpIfHCInOctets:                       "ifHCInOctets",
	SPLSnmpIfHCInUcastPkts:                    "ifHCInUcastPkts",
	SPLSnmpIfHCInMulticastPkts:                "ifHCInMulticastPkts",
	SPLSnmpIfHCInBroadcastPkts:                "ifHCInBroadcastPkts",
	SPLSnmpIfHCOutOctets:                      "ifHCOutOctets",
	SPLSnmpIfHCOutUcastPkts:                   "ifHCOutUcastPkts",
	SPLSnmpIfHCOutMulticastPkts:               "ifHCOutMulticastPkts",
	SPLSnmpIfHCOutBroadcastPckts:              "ifHCOutBroadcastPckts",
	SPLSnmpIpv6IfStatsInReceives:              "ipv6IfStatsInReceives",
	SPLSnmpIpv6IfStatsInHdrErrors:             "ipv6IfStatsInHdrErrors",
	SPLSnmpIpv6IfStatsInAddrErrors:            "ipv6IfStatsInAddrErrors",
	SPLSnmpIpv6IfStatsInDiscards:              "ipv6IfStatsInDiscards",
	SPLSnmpIpv6IfStatsOutForwDatagrams:        "ipv6IfStatsOutForwDatagrams",
	SPLSnmpIpv6IfStatsOutDiscards:             "ipv6IfStatsOutDiscards",
	SPLSnmpIpv6IfStatsInMcastPkts:             "ipv6IfStatsInMcastPkts",
	SPLSnmpIpv6IfStatsOutMcastPkts:            "ipv6IfStatsOutMcastPkts",
	SPLSnmpIfInBroadcastPkts:                  "ifInBroadcastPkts",
	SPLSnmpIfInMulticastPkts:                  "ifInMulticastPkts",
	SPLSnmpIfOutBroadcastPkts:                 "ifOutBroadcastPkts",
	SPLSnmpIfOutMulticastPkts:                 "ifOutMulticastPkts",
	SPLSnmpIeee8021PfcRequests:                "ieee8021PfcRequests",
	SPLSnmpIeee8021PfcIndications:             "ieee8021PfcIndications",
	NSLSnmpReceivedUndersizePkts:              "NSL_receivedUndersizePkts",
	NSLSnmpTransmittedUndersizePkts:           "NSL_transmittedUndersizePkts",
	NSLSnmpIPMCBridgedPckts:                   "NSL_iPMCBridgedPckts",
	NSLSnmpIPMCRoutedPckts:                    "NSL_iPMCRoutedPckts",
	NSLSnmpIPMCInDroppedPckts:                 "NSL_iPMCInDroppedPckts",
	NSLSnmpIPMCOutDroppedPckts:                "NSL_iPMCOutDroppedPckts",
	NSLSnmpEtherStatsPkts1519to1522Octets:     "NSL_etherStatsPkts1519to1522Octets",
	NSLSnmpEtherStatsPkts1522to2047Octets:     "NSL_etherStatsPkts1522to2047Octets",
	NSLSnmpEtherStatsPkts2048to4095Octets:     "NSL_etherStatsPkts2048to4095Octets",
	NSLSnmpEtherStatsPkts4095to9216Octets:     "NSL_etherStatsPkts4095to9216Octets",
	NSLSnmpReceivedPkts64Octets:               "NSL_receivedPkts64Octets",
	NSLSnmpReceivedPkts65to127Octets:          "NSL_receivedPkts65to127Octets",
	NSLSnmpReceivedPkts128to255Octets:         "NSL_receivedPkts128to255Octets",
	NSLSnmpReceivedPkts256to511Octets:         "NSL_receivedPkts256to511Octets",
	NSLSnmpReceivedPkts512to1023Octets:        "NSL_receivedPkts512to1023Octets",
	NSLSnmpReceivedPkts1024to1518Octets:       "NSL_receivedPkts1024to1518Octets",
	NSLSnmpReceivedPkts1519to2047Octets:       "NSL_receivedPkts1519to2047Octets",
	NSLSnmpReceivedPkts2048to4095Octets:       "NSL_receivedPkts2048to4095Octets",
	NSLSnmpReceivedPkts4095to9216Octets:       "NSL_receivedPkts4095to9216Octets",
	NSLSnmpTransmittedPkts64Octets:            "NSL_transmittedPkts64Octets",
	NSLSnmpTransmittedPkts65to127Octets:       "NSL_transmittedPkts65to127Octets",
	NSLSnmpTransmittedPkts128to255Octets:      "NSL_transmittedPkts128to255Octets",
	NSLSnmpTransmittedPkts256to511Octets:      "NSL_transmittedPkts256to511Octets",
	NSLSnmpTransmittedPkts512to1023Octets:     "NSL_transmittedPkts512to1023Octets",
	NSLSnmpTransmittedPkts1024to1518Octets:    "NSL_transmittedPkts1024to1518Octets",
	NSLSnmpTransmittedPkts1519to2047Octets:    "NSL_transmittedPkts1519to2047Octets",
	NSLSnmpTransmittedPkts2048to4095Octets:    "NSL_transmittedPkts2048to4095Octets",
	NSLSnmpTransmittedPkts4095to9216Octets:    "NSL_transmittedPkts4095to9216Octets",
	NSLSnmpTxControlCells:                     "NSL_txControlCells",
	NSLSnmpTxDataCells:                        "NSL_txDataCells",
	NSLSnmpTxDataBytes:                        "NSL_txDataBytes",
	NSLSnmpRxCrcErrors:                        "NSL_rxCrcErrors",
	NSLSnmpRxFecCorrectable:                   "NSL_rxFecCorrectable",
	NSLSnmpRxBecCrcErrors:                     "NSL_rxBecCrcErrors",
	NSLSnmpRxDisparityErrors:                  "NSL_rxDisparityErrors",
	NSLSnmpRxControlCells:                     "NSL_rxControlCells",
	NSLSnmpRxDataCells:                        "NSL_rxDataCells",
	NSLSnmpRxDataBytes:                        "NSL_rxDataBytes",
	NSLSnmpRxDroppedRetransmittedControl:      "NSL_rxDroppedRetransmittedControl",
	NSLSnmpTxBecRetransmit:                    "NSL_txBecRetransmit",
	NSLSnmpRxBecRetransmit:                    "NSL_txBecRetransmit",
	NSLSnmpTxAsynFifoRate:                     "NSL_txAsynFifoRate",
	NSLSnmpRxAsynFifoRate:                     "NSL_rxAsynFifoRate",
	NSLSnmpRxFecUncorrectable:                 "NSL_rxFecUncorrectable",
	NSLSnmpRxBecRxFault:                       "NSL_rxBecRxFault",
	NSLSnmpRxCodeErrors:                       "NSL_rxCodeErrors",
	NSLSnmpRxRsFecBitError:                    "NSL_rxRsFecBitError",
	NSLSnmpRxRsFecSymbolError:                 "NSL_rxRsFecSymbolError",
	NSLSnmpRxLlfcPrimary:                      "NSL_rxLlfcPrimary",
	NSLSnmpRxLlfcSecondary:                    "NSL_rxLlfcSecondary",
	NSLSnmpCustomReceive0:                     "NSL_customReceive0",
	NSLSnmpCustomReceive1:                     "NSL_customReceive1",
	NSLSnmpCustomReceive2:                     "NSL_customReceive2",
	NSLSnmpCustomReceive3:                     "NSL_customReceive3",
	NSLSnmpCustomReceive4:                     "NSL_customReceive4",
	NSLSnmpCustomReceive5:                     "NSL_customReceive5",
	NSLSnmpCustomReceive6:                     "NSL_customReceive6",
	NSLSnmpCustomReceive7:                     "NSL_customReceive7",
	NSLSnmpCustomReceive8:                     "NSL_customReceive8",
	NSLSnmpCustomTransmit0:                    "NSL_customTransmit0",
	NSLSnmpCustomTransmit1:                    "NSL_customTransmit1",
	NSLSnmpCustomTransmit2:                    "NSL_customTransmit2",
	NSLSnmpCustomTransmit3:                    "NSL_customTransmit3",
	NSLSnmpCustomTransmit4:                    "NSL_cstomTransmit4",
	NSLSnmpCustomTransmit5:                    "NSL_customTransmit5",
	NSLSnmpCustomTransmit6:                    "NSL_customTransmit6",
	NSLSnmpCustomTransmit7:                    "NSL_customTransmit7",
	NSLSnmpCustomTransmit8:                    "NSL_customTransmit8",
	NSLSnmpCustomTransmit9:                    "NSL_customTransmit9",
	NSLSnmpCustomTransmit10:                   "NSL_customTransmit10",
	NSLSnmpCustomTransmit11:                   "NSL_customTransmit11",
	NSLSnmpCustomTransmit12:                   "NSL_customTransmit12",
	NSLSnmpCustomTransmit13:                   "NSL_customTransmit13",
	NSLSnmpCustomTransmit14:                   "NSL_customTransmit14",
	SPLSnmpDot3StatsInRangeLengthError:        "dot3StatsInRangeLengthError",
	SPLSnmpDot3OmpEmulationCRC8Errors:         "dot3OmpEmulationCRC8Errors",
	SPLSnmpDot3MpcpRxGate:                     "dot3MpcpRxGate",
	SPLSnmpDot3MpcpRxRegister:                 "dot3MpcpRxRegister",
	SPLSnmpDot3MpcpTxRegRequest:               "dot3MpcpTxRegRequest",
	SPLSnmpDot3MpcpTxRegAck:                   "dot3MpcpTxRegAck",
	SPLSnmpDot3MpcpTxReport:                   "dot3MpcpTxReport",
	SPLSnmpDot3EponFecCorrectedBlocks:         "dot3EponFecCorrectedBlocks",
	SPLSnmpDot3EponFecUncorrectableBlocks:     "dot3EponFecUncorrectableBlocks",
	NSLSnmpPonInDroppedOctets:                 "NSL_ponInDroppedOctets",
	NSLSnmpPonOutDroppedOctets:                "NSL_ponOutDroppedOctets",
	NSLSnmpPonInDelayedOctets:                 "NSL_pnInDelayedOctets",
	NSLSnmpPonOutDelayedOctets:                "NSL_ponOutDelayedOctets",
	NSLSnmpPonInDelayedHundredUs:              "NSL_ponInDelayedHundredUs",
	NSLSnmpPonOutDelayedHundredUs:             "NSL_ponOutDelayedHundredUs",
	NSLSnmpPonInFrameErrors:                   "NSL_ponInFrameErrors",
	NSLSnmpPonInOamFrames:                     "NSL_ponInOamFrames",
	NSLSnmpPonOutOamFrames:                    "NSL_ponOutOamFrames",
	NSLSnmpPonOutUnusedOctets:                 "NSL_ponOutUnusedOctets",
	NSLSnmpEtherStatsPkts9217to16383Octets:    "NSL_etherStatsPkts9217to16383Octets",
	NSLSnmpReceivedPkts9217to16383Octets:      "NSL_receivedPkts9217to16383Octets",
	NSLSnmpTransmittedPkts9217to16383Octets:   "NSL_transmittedPkts9217to16383Octets",
	NSLSnmpRxVlanTagFrame:                     "NSL_rxVlanTagFrame",
	NSLSnmpRxDoubleVlanTagFrame:               "NSL_rxDoubleVlanTagFrame",
	NSLSnmpTxVlanTagFrame:                     "NSL_txVlanTagFrame",
	NSLSnmpTxDoubleVlanTagFrame:               "NSL_TxDoubleVlanTagFrame",
	NSLSnmpRxPFCControlFrame:                  "NSL_rxPFCControlFrame",
	NSLSnmpTxPFCControlFrame:                  "NSL_txPFCControlFrame",
	NSLSnmpRxPFCFrameXonPriority0:             "NSL_rxPFCFrameXonPriority0",
	NSLSnmpRxPFCFrameXonPriority1:             "NSL_RxPFCFrameXonPriority1",
	NSLSnmpRxPFCFrameXonPriority2:             "NSL_rxPFCFrameXonPriority2",
	NSLSnmpRxPFCFrameXonPriority3:             "NSL_rxPFCFrameXonPriority3",
	NSLSnmpRxPFCFrameXonPriority4:             "NSL_rxPFCFrameXonPriority4",
	NSLSnmpRxPFCFrameXonPriority5:             "NSL_rxPFCFrameXonPriority5",
	NSLSnmpRxPFCFrameXonPriority6:             "NSL_rxPFCFrameXonPriority6",
	NSLSnmpRxPFCFrameXonPriority7:             "NSL_rxPFCFrameXonPriority7",
	NSLSnmpRxPFCFramePriority0:                "NSL_rxPFCFramePriority0",
	NSLSnmpRxPFCFramePriority1:                "NSL_rxPFCFramePriority1",
	NSLSnmpRxPFCFramePriority2:                "NSL_rxPFCFramePriority2",
	NSLSnmpRxPFCFramePriority3:                "NSL_rxPFCFramePriority3",
	NSLSnmpRxPFCFramePriority4:                "NSL_rxPFCFramePriority4",
	NSLSnmpRxPFCFramePriority5:                "NSL_rxPFCFramePriority5",
	NSLSnmpRxPFCFramePriority6:                "NSL_rxPFCFramePriority6",
	NSLSnmpRxPFCFramePriority7:                "NSL_rxPFCFramePriority7",
	NSLSnmpTxPFCFramePriority0:                "NSL_txPFCFramePriority0",
	NSLSnmpTxPFCFramePriority1:                "NSL_txPFCFramePriority1",
	NSLSnmpTxPFCFramePriority2:                "NSL_txPFCFramePriority2",
	NSLSnmpTxPFCFramePriority3:                "NSL_txPFCFramePriority3",
	NSLSnmpTxPFCFramePriority4:                "NSL_txPFCFramePriority4",
	NSLSnmpTxPFCFramePriority5:                "NSL_txPFCFramePriority5",
	NSLSnmpTxPFCFramePriority6:                "NSL_txPFCFramePriority6",
	NSLSnmpTxPFCFramePriority7:                "NSL_txPFCFramePriority7",
	SPLSnmpFcmPortClass3RxFrames:              "fcmPortClass3RxFrames",
	SPLSnmpFcmPortClass3TxFrames:              "fcmPortClass3TxFrames",
	SPLSnmpFcmPortClass3Discards:              "fcmPortClass3Discards",
	SPLSnmpFcmPortClass2RxFrames:              "fcmPortClass2RxFrames",
	SPLSnmpFcmPortClass2TxFrames:              "fcmPortClass2TxFrames",
	SPLSnmpFcmPortClass2Discards:              "fcmPortClass2Discards",
	SPLSnmpFcmPortInvalidCRCs:                 "fcmPortInvalidCRCs",
	SPLSnmpFcmPortDelimiterErrors:             "fcmPortDelimiterErrors",
	NSLSnmpSampleIngressPkts:                  "NSL_sampleIngressPkts",
	NSLSnmpSampleIngressSnapshotPkts:          "NSL_sampleIngressSnapshotPkts",
	NSLSnmpSampleIngressSampledPkts:           "NSL_sampleIngressSampledPkts",
	NSLSnmpSampleFlexPkts:                     "NSL_sampleFlexPkts",
	NSLSnmpSampleFlexSnapshotPkts:             "NSL_sampleFlexSnapshotPkts",
	NSLSnmpSampleFlexSampledPkts:              "NSL_sampleFlexSampledPkts",
	NSLSnmpEgressProtectionDataDrop:           "NSL_egressProtectionDataDrop",
	NSLSnmpTxE2ECCControlFrames:               "NSL_txE2ECCControlFrames",
	NSLSnmpE2EHOLDropPkts:                     "NSL_e2EHOLDropPkts",
	SPLSnmpEtherStatsTxCRCAlignErrors:         "etherStatsTxCRCAlignErrors",
	SPLSnmpEtherStatsTxJabbers:                "etherStatsTxJabbers",
	NSLSnmpMacMergeTxFrag:                     "NSL_macMergeTxFrag",
	NSLSnmpMacMergeTxVerifyFrame:              "NSL_macMergeTxVerifyFrame",
	NSLSnmpMacMergeTxReplyFrame:               "NSL_macMergeTxReplyFrame",
	NSLSnmpMacMergeRxFrameAssErrors:           "NSL_macMergeRxFrameAssErrors",
	NSLSnmpMacMergeRxFrameSmdErrors:           "NSL_macMergeRxFrameSmdErrors",
	NSLSnmpMacMergeRxFrameAss:                 "NSL_macMergeRxFrameAss",
	NSLSnmpMacMergeRxFrag:                     "NSL_macMergeRxFrag",
	NSLSnmpMacMergeRxVerifyFrame:              "NSL_macMergeRxVerifyFrame",
	NSLSnmpMacMergeRxReplyFrame:               "NSL_macMergeRxReplyFrame",
	NSLSnmpMacMergeRxFinalFragSizeError:       "NSL_macMergeRxFinalFragSizeError",
	NSLSnmpMacMergeRxFragSizeError:            "NSL_macMergeRxFragSizeError",
	NSLSnmpMacMergeRxDiscard:                  "NSL_macMergeRxDiscard",
	NSLSnmpMacMergeHoldCount:                  "NSL_macMergeHoldCount",
	NSLSnmpRxBipErrorCount:                    "NSL_rxBipErrorCount",
}

func ParseStatVal(s string) (StatVal, error) {
	if v, ok := statVal_values[s]; ok {
		return v, nil
	}
	return 0, fmt.Errorf("Invalid StatVal. %s", s)
}

var statVal_values = func() map[string]StatVal {
	values := map[string]StatVal{}
	for value, name := range statVal_names {
		values[name] = value
	}
	return values
}()

//
// methods
//
func statValsToC(statVals []StatVal) []C.opennsl_stat_val_t {
	c_stats := make([]C.opennsl_stat_val_t, len(statVals))
	for index, statVal := range statVals {
		c_stats[index] = statVal.C()
	}
	return c_stats
}

func statValuesToGo32(c_values []C.uint32) []uint32 {
	values := make([]uint32, len(c_values))
	for index, c_value := range c_values {
		values[index] = uint32(c_value)
	}
	return values
}

func statValuesToGo64(c_values []C.uint64) []uint64 {
	values := make([]uint64, len(c_values))
	for index, c_value := range c_values {
		values[index] = uint64(c_value)
	}
	return values
}

func StatGet(unit int, port Port, v StatVal) (uint64, error) {
	c_val := C.uint64(0)

	rc := C.opennsl_stat_get(C.int(unit), port.C(), v.C(), &c_val)
	return uint64(c_val), ParseError(rc)
}

func StatValMultiGet(unit int, port Port, statVals ...StatVal) ([]uint64, error) {
	c_nstat := C.int(len(statVals))
	c_values := make([]C.uint64, c_nstat)

	if c_nstat > 0 {
		c_stats := statValsToC(statVals)
		rc := C.opennsl_stat_multi_get(C.int(unit), port.C(), c_nstat, &c_stats[0], &c_values[0])
		if err := ParseError(rc); err != nil {
			return nil, err
		}
	}

	return statValuesToGo64(c_values), nil
}

func StatGet32(unit int, port Port, v StatVal) (uint32, error) {
	c_val := C.uint32(0)

	rc := C.opennsl_stat_get32(C.int(unit), port.C(), v.C(), &c_val)
	return uint32(c_val), ParseError(rc)
}

func StatValMultiGet32(unit int, port Port, statVals ...StatVal) ([]uint32, error) {
	c_nstat := C.int(len(statVals))
	c_values := make([]C.uint32, c_nstat)

	if c_nstat > 0 {
		c_stats := statValsToC(statVals)
		rc := C.opennsl_stat_multi_get32(C.int(unit), port.C(), c_nstat, &c_stats[0], &c_values[0])
		if err := ParseError(rc); err != nil {
			return nil, err
		}
	}

	return statValuesToGo32(c_values), nil
}

func StatSyncGet32(unit int, port Port, v StatVal) (uint32, error) {
	c_val := C.uint32(0)

	rc := C.opennsl_stat_sync_get32(C.int(unit), port.C(), v.C(), &c_val)
	return uint32(c_val), ParseError(rc)
}

func StatValSyncMultiGet32(unit int, port Port, statVals ...StatVal) ([]uint32, error) {
	c_nstat := C.int(len(statVals))
	c_values := make([]C.uint32, c_nstat)

	if c_nstat > 0 {
		c_stats := statValsToC(statVals)
		rc := C.opennsl_stat_sync_multi_get32(C.int(unit), port.C(), c_nstat, &c_stats[0], &c_values[0])
		if err := ParseError(rc); err != nil {
			return nil, err
		}
	}

	return statValuesToGo32(c_values), nil
}

func StatClearSingle(unit int, port Port, v StatVal) error {
	rc := C.opennsl_stat_clear_single(C.int(unit), port.C(), v.C())
	return ParseError(rc)
}

//
// StatVal methods
//
func (v StatVal) Get(unit int, port Port) (uint64, error) {
	return StatGet(unit, port, v)
}

func (v StatVal) Get32(unit int, port Port) (uint32, error) {
	return StatGet32(unit, port, v)
}

func (v StatVal) SyncGet32(unit int, port Port) (uint32, error) {
	return StatSyncGet32(unit, port, v)
}

func (v StatVal) ClearSingle(unit int, port Port) error {
	return StatClearSingle(unit, port, v)
}

//
// Port methods
//
func (v Port) StatGet(unit int, stat StatVal) (uint64, error) {
	return stat.Get(unit, v)
}

func (v Port) StatMultiGet(unit int, statVals ...StatVal) ([]uint64, error) {
	return StatValMultiGet(unit, v, statVals...)
}

func (v Port) StatGet32(unit int, stat StatVal) (uint32, error) {
	return stat.Get32(unit, v)
}

func (v Port) StatMultiGet32(unit int, statVals ...StatVal) ([]uint32, error) {
	return StatValMultiGet32(unit, v, statVals...)
}

func (v Port) StatSyncGet32(unit int, stat StatVal) (uint32, error) {
	return stat.SyncGet32(unit, v)
}

func (v Port) StatSyncMultiGet32(unit int, statVals ...StatVal) ([]uint32, error) {
	return StatValSyncMultiGet32(unit, v, statVals...)
}

func (v Port) StatClearSingle(unit int, stat StatVal) error {
	return stat.ClearSingle(unit, v)
}
