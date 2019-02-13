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
#include <opennsl/types.h>
#include <opennsl/rx.h>
*/
import "C"

import (
	"fmt"
)

//
// RxReason
//
type RxReason C.opennsl_rx_reason_t

func (v RxReason) C() C.opennsl_rx_reason_t {
	return C.opennsl_rx_reason_t(v)
}

const (
	RxReasonInvalid                    RxReason = C.opennslRxReasonInvalid
	RxReasonArp                        RxReason = C.opennslRxReasonArp
	RxReasonBpdu                       RxReason = C.opennslRxReasonBpdu
	RxReasonBroadcast                  RxReason = C.opennslRxReasonBroadcast
	RxReasonClassBasedMove             RxReason = C.opennslRxReasonClassBasedMove
	RxReasonClassTagPackets            RxReason = C.opennslRxReasonClassTagPackets
	RxReasonControl                    RxReason = C.opennslRxReasonControl
	RxReasonCpuLearn                   RxReason = C.opennslRxReasonCpuLearn
	RxReasonDestLookupFail             RxReason = C.opennslRxReasonDestLookupFail
	RxReasonDhcp                       RxReason = C.opennslRxReasonDhcp
	RxReasonDosAttack                  RxReason = C.opennslRxReasonDosAttack
	RxReasonE2eHolIbp                  RxReason = C.opennslRxReasonE2eHolIbp
	RxReasonEncapHigigError            RxReason = C.opennslRxReasonEncapHigigError
	RxReasonFilterMatch                RxReason = C.opennslRxReasonFilterMatch
	RxReasonGreChecksum                RxReason = C.opennslRxReasonGreChecksum
	RxReasonGreSourceRoute             RxReason = C.opennslRxReasonGreSourceRoute
	RxReasonHigigControl               RxReason = C.opennslRxReasonHigigControl
	RxReasonHigigHdrError              RxReason = C.opennslRxReasonHigigHdrError
	RxReasonIcmpRedirect               RxReason = C.opennslRxReasonIcmpRedirect
	RxReasonIgmp                       RxReason = C.opennslRxReasonIgmp
	RxReasonIngressFilter              RxReason = C.opennslRxReasonIngressFilter
	RxReasonIp                         RxReason = C.opennslRxReasonIp
	RxReasonIpfixRateViolation         RxReason = C.opennslRxReasonIpfixRateViolation
	RxReasonIpMcastMiss                RxReason = C.opennslRxReasonIpMcastMiss
	RxReasonIpmcReserved               RxReason = C.opennslRxReasonIpmcReserved
	RxReasonIpOptionVersion            RxReason = C.opennslRxReasonIpOptionVersion
	RxReasonIpmc                       RxReason = C.opennslRxReasonIpmc
	RxReasonL2Cpu                      RxReason = C.opennslRxReasonL2Cpu
	RxReasonL2DestMiss                 RxReason = C.opennslRxReasonL2DestMiss
	RxReasonL2LearnLimit               RxReason = C.opennslRxReasonL2LearnLimit
	RxReasonL2Move                     RxReason = C.opennslRxReasonL2Move
	RxReasonL2MtuFail                  RxReason = C.opennslRxReasonL2MtuFail
	RxReasonL2NonUnicastMiss           RxReason = C.opennslRxReasonL2NonUnicastMiss
	RxReasonL2SourceMiss               RxReason = C.opennslRxReasonL2SourceMiss
	RxReasonL3AddrBindFail             RxReason = C.opennslRxReasonL3AddrBindFail
	RxReasonL3DestMiss                 RxReason = C.opennslRxReasonL3DestMiss
	RxReasonL3HeaderError              RxReason = C.opennslRxReasonL3HeaderError
	RxReasonL3MtuFail                  RxReason = C.opennslRxReasonL3MtuFail
	RxReasonL3Slowpath                 RxReason = C.opennslRxReasonL3Slowpath
	RxReasonL3SourceMiss               RxReason = C.opennslRxReasonL3SourceMiss
	RxReasonL3SourceMove               RxReason = C.opennslRxReasonL3SourceMove
	RxReasonMartianAddr                RxReason = C.opennslRxReasonMartianAddr
	RxReasonMcastIdxError              RxReason = C.opennslRxReasonMcastIdxError
	RxReasonMcastMiss                  RxReason = C.opennslRxReasonMcastMiss
	RxReasonMimServiceError            RxReason = C.opennslRxReasonMimServiceError
	RxReasonMplsCtrlWordError          RxReason = C.opennslRxReasonMplsCtrlWordError
	RxReasonMplsError                  RxReason = C.opennslRxReasonMplsError
	RxReasonMplsInvalidAction          RxReason = C.opennslRxReasonMplsInvalidAction
	RxReasonMplsInvalidPayload         RxReason = C.opennslRxReasonMplsInvalidPayload
	RxReasonMplsLabelMiss              RxReason = C.opennslRxReasonMplsLabelMiss
	RxReasonMplsSequenceNumber         RxReason = C.opennslRxReasonMplsSequenceNumber
	RxReasonMplsTtl                    RxReason = C.opennslRxReasonMplsTtl
	RxReasonMulticast                  RxReason = C.opennslRxReasonMulticast
	RxReasonNhop                       RxReason = C.opennslRxReasonNhop
	RxReasonOAMError                   RxReason = C.opennslRxReasonOAMError
	RxReasonOAMSlowpath                RxReason = C.opennslRxReasonOAMSlowpath
	RxReasonOAMLMDM                    RxReason = C.opennslRxReasonOAMLMDM
	RxReasonParityError                RxReason = C.opennslRxReasonParityError
	RxReasonProtocol                   RxReason = C.opennslRxReasonProtocol
	RxReasonSampleDest                 RxReason = C.opennslRxReasonSampleDest
	RxReasonSampleSource               RxReason = C.opennslRxReasonSampleSource
	RxReasonSharedVlanMismatch         RxReason = C.opennslRxReasonSharedVlanMismatch
	RxReasonSourceRoute                RxReason = C.opennslRxReasonSourceRoute
	RxReasonTimeStamp                  RxReason = C.opennslRxReasonTimeStamp
	RxReasonTtl                        RxReason = C.opennslRxReasonTtl
	RxReasonTtl1                       RxReason = C.opennslRxReasonTtl1
	RxReasonTunnelError                RxReason = C.opennslRxReasonTunnelError
	RxReasonUdpChecksum                RxReason = C.opennslRxReasonUdpChecksum
	RxReasonUnknownVlan                RxReason = C.opennslRxReasonUnknownVlan
	RxReasonUrpfFail                   RxReason = C.opennslRxReasonUrpfFail
	RxReasonVcLabelMiss                RxReason = C.opennslRxReasonVcLabelMiss
	RxReasonVlanFilterMatch            RxReason = C.opennslRxReasonVlanFilterMatch
	RxReasonWlanClientError            RxReason = C.opennslRxReasonWlanClientError
	RxReasonWlanSlowpath               RxReason = C.opennslRxReasonWlanSlowpath
	RxReasonWlanDot1xDrop              RxReason = C.opennslRxReasonWlanDot1xDrop
	RxReasonExceptionFlood             RxReason = C.opennslRxReasonExceptionFlood
	RxReasonTimeSync                   RxReason = C.opennslRxReasonTimeSync
	RxReasonEAVData                    RxReason = C.opennslRxReasonEAVData
	RxReasonSamePortBridge             RxReason = C.opennslRxReasonSamePortBridge
	RxReasonSplitHorizon               RxReason = C.opennslRxReasonSplitHorizon
	RxReasonL4Error                    RxReason = C.opennslRxReasonL4Error
	RxReasonStp                        RxReason = C.opennslRxReasonStp
	RxReasonEgressFilterRedirect       RxReason = C.opennslRxReasonEgressFilterRedirect
	RxReasonFilterRedirect             RxReason = C.opennslRxReasonFilterRedirect
	RxReasonLoopback                   RxReason = C.opennslRxReasonLoopback
	RxReasonVlanTranslate              RxReason = C.opennslRxReasonVlanTranslate
	RxReasonMmrp                       RxReason = C.opennslRxReasonMmrp
	RxReasonSrp                        RxReason = C.opennslRxReasonSrp
	RxReasonTunnelControl              RxReason = C.opennslRxReasonTunnelControl
	RxReasonL2Marked                   RxReason = C.opennslRxReasonL2Marked
	RxReasonWlanSlowpathKeepalive      RxReason = C.opennslRxReasonWlanSlowpathKeepalive
	RxReasonStation                    RxReason = C.opennslRxReasonStation
	RxReasonNiv                        RxReason = C.opennslRxReasonNiv
	RxReasonNivPrioDrop                RxReason = C.opennslRxReasonNivPrioDrop
	RxReasonNivInterfaceMiss           RxReason = C.opennslRxReasonNivInterfaceMiss
	RxReasonNivRpfFail                 RxReason = C.opennslRxReasonNivRpfFail
	RxReasonNivTagInvalid              RxReason = C.opennslRxReasonNivTagInvalid
	RxReasonNivTagDrop                 RxReason = C.opennslRxReasonNivTagDrop
	RxReasonNivUntagDrop               RxReason = C.opennslRxReasonNivUntagDrop
	RxReasonTrill                      RxReason = C.opennslRxReasonTrill
	RxReasonTrillInvalid               RxReason = C.opennslRxReasonTrillInvalid
	RxReasonTrillMiss                  RxReason = C.opennslRxReasonTrillMiss
	RxReasonTrillRpfFail               RxReason = C.opennslRxReasonTrillRpfFail
	RxReasonTrillSlowpath              RxReason = C.opennslRxReasonTrillSlowpath
	RxReasonTrillCoreIsIs              RxReason = C.opennslRxReasonTrillCoreIsIs
	RxReasonTrillTtl                   RxReason = C.opennslRxReasonTrillTtl
	RxReasonTrillName                  RxReason = C.opennslRxReasonTrillName
	RxReasonBfdSlowpath                RxReason = C.opennslRxReasonBfdSlowpath
	RxReasonBfd                        RxReason = C.opennslRxReasonBfd
	RxReasonMirror                     RxReason = C.opennslRxReasonMirror
	RxReasonRegexAction                RxReason = C.opennslRxReasonRegexAction
	RxReasonRegexMatch                 RxReason = C.opennslRxReasonRegexMatch
	RxReasonFailoverDrop               RxReason = C.opennslRxReasonFailoverDrop
	RxReasonWlanTunnelError            RxReason = C.opennslRxReasonWlanTunnelError
	RxReasonCongestionCnmProxy         RxReason = C.opennslRxReasonCongestionCnmProxy
	RxReasonCongestionCnmProxyError    RxReason = C.opennslRxReasonCongestionCnmProxyError
	RxReasonCongestionCnm              RxReason = C.opennslRxReasonCongestionCnm
	RxReasonMplsUnknownAch             RxReason = C.opennslRxReasonMplsUnknownAch
	RxReasonMplsLookupsExceeded        RxReason = C.opennslRxReasonMplsLookupsExceeded
	RxReasonMplsReservedEntropyLabel   RxReason = C.opennslRxReasonMplsReservedEntropyLabel
	RxReasonMplsIllegalReservedLabel   RxReason = C.opennslRxReasonMplsIllegalReservedLabel
	RxReasonMplsRouterAlertLabel       RxReason = C.opennslRxReasonMplsRouterAlertLabel
	RxReasonNivPrune                   RxReason = C.opennslRxReasonNivPrune
	RxReasonVirtualPortPrune           RxReason = C.opennslRxReasonVirtualPortPrune
	RxReasonNonUnicastDrop             RxReason = C.opennslRxReasonNonUnicastDrop
	RxReasonTrillPacketPortMismatch    RxReason = C.opennslRxReasonTrillPacketPortMismatch
	RxReasonWlanClientMove             RxReason = C.opennslRxReasonWlanClientMove
	RxReasonWlanSourcePortMiss         RxReason = C.opennslRxReasonWlanSourcePortMiss
	RxReasonWlanClientSourceMiss       RxReason = C.opennslRxReasonWlanClientSourceMiss
	RxReasonWlanClientDestMiss         RxReason = C.opennslRxReasonWlanClientDestMiss
	RxReasonWlanMtu                    RxReason = C.opennslRxReasonWlanMtu
	RxReasonL2GreSipMiss               RxReason = C.opennslRxReasonL2GreSipMiss
	RxReasonL2GreVpnIdMiss             RxReason = C.opennslRxReasonL2GreVpnIdMiss
	RxReasonTimesyncUnknownVersion     RxReason = C.opennslRxReasonTimesyncUnknownVersion
	RxReasonBfdUnknownVersion          RxReason = C.opennslRxReasonBfdUnknownVersion
	RxReasonBfdInvalidVersion          RxReason = C.opennslRxReasonBfdInvalidVersion
	RxReasonBfdLookupFailure           RxReason = C.opennslRxReasonBfdLookupFailure
	RxReasonBfdInvalidPacket           RxReason = C.opennslRxReasonBfdInvalidPacket
	RxReasonVxlanSipMiss               RxReason = C.opennslRxReasonVxlanSipMiss
	RxReasonVxlanVpnIdMiss             RxReason = C.opennslRxReasonVxlanVpnIdMiss
	RxReasonFcoeZoneCheckFail          RxReason = C.opennslRxReasonFcoeZoneCheckFail
	RxReasonIpmcInterfaceMismatch      RxReason = C.opennslRxReasonIpmcInterfaceMismatch
	RxReasonNat                        RxReason = C.opennslRxReasonNat
	RxReasonTcpUdpNatMiss              RxReason = C.opennslRxReasonTcpUdpNatMiss
	RxReasonIcmpNatMiss                RxReason = C.opennslRxReasonIcmpNatMiss
	RxReasonNatFragment                RxReason = C.opennslRxReasonNatFragment
	RxReasonNatMiss                    RxReason = C.opennslRxReasonNatMiss
	RxReasonOAMCCMSlowpath             RxReason = C.opennslRxReasonOAMCCMSlowpath
	RxReasonBHHOAM                     RxReason = C.opennslRxReasonBHHOAM
	RxReasonUnknownSubtendingPort      RxReason = C.opennslRxReasonUnknownSubtendingPort
	RxReasonReserved0                  RxReason = C.opennslRxReasonReserved0
	RxReasonOAMMplsLmDm                RxReason = C.opennslRxReasonOAMMplsLmDm
	RxReasonSat                        RxReason = C.opennslRxReasonSat
	RxReasonSampleSourceFlex           RxReason = C.opennslRxReasonSampleSourceFlex
	RxReasonFlexSflow                  RxReason = C.opennslRxReasonFlexSflow
	RxReasonVxltMiss                   RxReason = C.opennslRxReasonVxltMiss
	RxReasonTunnelDecapEcnError        RxReason = C.opennslRxReasonTunnelDecapEcnError
	RxReasonTunnelObjectValidationFail RxReason = C.opennslRxReasonTunnelObjectValidationFail
	RxReasonL3Cpu                      RxReason = C.opennslRxReasonL3Cpu
	RxReasonTunnelAdaptLookupMiss      RxReason = C.opennslRxReasonTunnelAdaptLookupMiss
	RxReasonPacketFlowSelectMiss       RxReason = C.opennslRxReasonPacketFlowSelectMiss
	RxReasonProtectionDataDrop         RxReason = C.opennslRxReasonProtectionDataDrop
	RxReasonPacketFlowSelect           RxReason = C.opennslRxReasonPacketFlowSelect
	RxReasonOtherLookupMiss            RxReason = C.opennslRxReasonOtherLookupMiss
	RxReasonInvalidTpid                RxReason = C.opennslRxReasonInvalidTpid
	RxReasonMplsControlPacket          RxReason = C.opennslRxReasonMplsControlPacket
	RxReasonTunnelTtlError             RxReason = C.opennslRxReasonTunnelTtlError
	RxReasonL2HeaderError              RxReason = C.opennslRxReasonL2HeaderError
	RxReasonOtherLookupHit             RxReason = C.opennslRxReasonOtherLookupHit
	RxReasonL2SrcLookupMiss            RxReason = C.opennslRxReasonL2SrcLookupMiss
	RxReasonL2SrcLookupHit             RxReason = C.opennslRxReasonL2SrcLookupHit
	RxReasonL2DstLookupMiss            RxReason = C.opennslRxReasonL2DstLookupMiss
	RxReasonL2DstLookupHit             RxReason = C.opennslRxReasonL2DstLookupHit
	RxReasonL3SrcRouteLookupMiss       RxReason = C.opennslRxReasonL3SrcRouteLookupMiss
	RxReasonL3SrcHostLookupMiss        RxReason = C.opennslRxReasonL3SrcHostLookupMiss
	RxReasonL3SrcRouteLookupHit        RxReason = C.opennslRxReasonL3SrcRouteLookupHit
	RxReasonL3SrcHostLookupHit         RxReason = C.opennslRxReasonL3SrcHostLookupHit
	RxReasonL3DstRouteLookupMiss       RxReason = C.opennslRxReasonL3DstRouteLookupMiss
	RxReasonL3DstHostLookupMiss        RxReason = C.opennslRxReasonL3DstHostLookupMiss
	RxReasonL3DstRouteLookupHit        RxReason = C.opennslRxReasonL3DstRouteLookupHit
	RxReasonL3DstHostLookupHit         RxReason = C.opennslRxReasonL3DstHostLookupHit
	RxReasonVlanTranslate1Lookup1Miss  RxReason = C.opennslRxReasonVlanTranslate1Lookup1Miss
	RxReasonVlanTranslate1Lookup2Miss  RxReason = C.opennslRxReasonVlanTranslate1Lookup2Miss
	RxReasonMplsLookup1Miss            RxReason = C.opennslRxReasonMplsLookup1Miss
	RxReasonMplsLookup2Miss            RxReason = C.opennslRxReasonMplsLookup2Miss
	RxReasonL3TunnelLookupMiss         RxReason = C.opennslRxReasonL3TunnelLookupMiss
	RxReasonVlanTranslate2Lookup1Miss  RxReason = C.opennslRxReasonVlanTranslate2Lookup1Miss
	RxReasonVlanTranslate2Lookup2Miss  RxReason = C.opennslRxReasonVlanTranslate2Lookup2Miss
	RxReasonL2StuFail                  RxReason = C.opennslRxReasonL2StuFail
	RxReasonSrCounterExceeded          RxReason = C.opennslRxReasonSrCounterExceeded
	RxReasonSrCopyToCpuBit0            RxReason = C.opennslRxReasonSrCopyToCpuBit0
	RxReasonSrCopyToCpuBit1            RxReason = C.opennslRxReasonSrCopyToCpuBit1
	RxReasonSrCopyToCpuBit2            RxReason = C.opennslRxReasonSrCopyToCpuBit2
	RxReasonSrCopyToCpuBit3            RxReason = C.opennslRxReasonSrCopyToCpuBit3
	RxReasonSrCopyToCpuBit4            RxReason = C.opennslRxReasonSrCopyToCpuBit4
	RxReasonSrCopyToCpuBit5            RxReason = C.opennslRxReasonSrCopyToCpuBit5
	RxReasonL3HeaderMismatch           RxReason = C.opennslRxReasonL3HeaderMismatch
)

const RxReasonCount = C.opennslRxReasonCount

func (v RxReason) String() string {
	if s, ok := rxReason_names[v]; ok {
		return s
	}
	return fmt.Sprintf("RxReason(%d)", v)
}

var rxReason_names = map[RxReason]string{
	RxReasonInvalid:                    "Invalid",
	RxReasonArp:                        "Arp",
	RxReasonBpdu:                       "Bpdu",
	RxReasonBroadcast:                  "Broadcast",
	RxReasonClassBasedMove:             "ClassBasedMove",
	RxReasonClassTagPackets:            "ClassTagPackets",
	RxReasonControl:                    "Control",
	RxReasonCpuLearn:                   "CpuLearn",
	RxReasonDestLookupFail:             "DestLookupFail",
	RxReasonDhcp:                       "Dhcp",
	RxReasonDosAttack:                  "DosAttack",
	RxReasonE2eHolIbp:                  "E2eHolIbp",
	RxReasonEncapHigigError:            "EncapHigigError",
	RxReasonFilterMatch:                "FilterMatch",
	RxReasonGreChecksum:                "GreChecksum",
	RxReasonGreSourceRoute:             "GreSourceRoute",
	RxReasonHigigControl:               "HigigControl",
	RxReasonHigigHdrError:              "HigigHdrError",
	RxReasonIcmpRedirect:               "IcmpRedirect",
	RxReasonIgmp:                       "Igmp",
	RxReasonIngressFilter:              "IngressFilter",
	RxReasonIp:                         "Ip",
	RxReasonIpfixRateViolation:         "IpfixRateViolation",
	RxReasonIpMcastMiss:                "IpMcastMiss",
	RxReasonIpmcReserved:               "IpmcReserved",
	RxReasonIpOptionVersion:            "IpOptionVersion",
	RxReasonIpmc:                       "Ipmc",
	RxReasonL2Cpu:                      "L2Cpu",
	RxReasonL2DestMiss:                 "L2DestMiss",
	RxReasonL2LearnLimit:               "L2LearnLimit",
	RxReasonL2Move:                     "L2Move",
	RxReasonL2MtuFail:                  "L2MtuFail",
	RxReasonL2NonUnicastMiss:           "L2NonUnicastMiss",
	RxReasonL2SourceMiss:               "L2SourceMiss",
	RxReasonL3AddrBindFail:             "L3AddrBindFail",
	RxReasonL3DestMiss:                 "L3DestMiss",
	RxReasonL3HeaderError:              "L3HeaderError",
	RxReasonL3MtuFail:                  "L3MtuFail",
	RxReasonL3Slowpath:                 "L3Slowpath",
	RxReasonL3SourceMiss:               "L3SourceMiss",
	RxReasonL3SourceMove:               "L3SourceMove",
	RxReasonMartianAddr:                "MartianAddr",
	RxReasonMcastIdxError:              "McastIdxError",
	RxReasonMcastMiss:                  "McastMiss",
	RxReasonMimServiceError:            "MimServiceError",
	RxReasonMplsCtrlWordError:          "MplsCtrlWordError",
	RxReasonMplsError:                  "MplsError",
	RxReasonMplsInvalidAction:          "MplsInvalidAction",
	RxReasonMplsInvalidPayload:         "MplsInvalidPayload",
	RxReasonMplsLabelMiss:              "MplsLabelMiss",
	RxReasonMplsSequenceNumber:         "MplsSequenceNumber",
	RxReasonMplsTtl:                    "MplsTtl",
	RxReasonMulticast:                  "Multicast",
	RxReasonNhop:                       "Nhop",
	RxReasonOAMError:                   "OAMError",
	RxReasonOAMSlowpath:                "OAMSlowpath",
	RxReasonOAMLMDM:                    "OAMLMDM",
	RxReasonParityError:                "ParityError",
	RxReasonProtocol:                   "Protocol",
	RxReasonSampleDest:                 "SampleDest",
	RxReasonSampleSource:               "SampleSource",
	RxReasonSharedVlanMismatch:         "SharedVlanMismatch",
	RxReasonSourceRoute:                "SourceRoute",
	RxReasonTimeStamp:                  "TimeStamp",
	RxReasonTtl:                        "Ttl",
	RxReasonTtl1:                       "Ttl1",
	RxReasonTunnelError:                "TunnelError",
	RxReasonUdpChecksum:                "UdpChecksum",
	RxReasonUnknownVlan:                "UnknownVlan",
	RxReasonUrpfFail:                   "UrpfFail",
	RxReasonVcLabelMiss:                "VcLabelMiss",
	RxReasonVlanFilterMatch:            "VlanFilterMatch",
	RxReasonWlanClientError:            "WlanClientError",
	RxReasonWlanSlowpath:               "WlanSlowpath",
	RxReasonWlanDot1xDrop:              "WlanDot1xDrop",
	RxReasonExceptionFlood:             "ExceptionFlood",
	RxReasonTimeSync:                   "TimeSync",
	RxReasonEAVData:                    "EAVData",
	RxReasonSamePortBridge:             "SamePortBridge",
	RxReasonSplitHorizon:               "SplitHorizon",
	RxReasonL4Error:                    "L4Error",
	RxReasonStp:                        "Stp",
	RxReasonEgressFilterRedirect:       "EgressFilterRedirect",
	RxReasonFilterRedirect:             "FilterRedirect",
	RxReasonLoopback:                   "Loopback",
	RxReasonVlanTranslate:              "VlanTranslate",
	RxReasonMmrp:                       "Mmrp",
	RxReasonSrp:                        "Srp",
	RxReasonTunnelControl:              "TunnelControl",
	RxReasonL2Marked:                   "L2Marked",
	RxReasonWlanSlowpathKeepalive:      "WlanSlowpathKeepalive",
	RxReasonStation:                    "Station",
	RxReasonNiv:                        "Niv",
	RxReasonNivPrioDrop:                "NivPrioDrop",
	RxReasonNivInterfaceMiss:           "NivInterfaceMiss",
	RxReasonNivRpfFail:                 "NivRpfFail",
	RxReasonNivTagInvalid:              "NivTagInvalid",
	RxReasonNivTagDrop:                 "NivTagDrop",
	RxReasonNivUntagDrop:               "NivUntagDrop",
	RxReasonTrill:                      "Trill",
	RxReasonTrillInvalid:               "TrillInvalid",
	RxReasonTrillMiss:                  "TrillMiss",
	RxReasonTrillRpfFail:               "TrillRpfFail",
	RxReasonTrillSlowpath:              "TrillSlowpath",
	RxReasonTrillCoreIsIs:              "TrillCoreIsIs",
	RxReasonTrillTtl:                   "TrillTtl",
	RxReasonTrillName:                  "TrillName",
	RxReasonBfdSlowpath:                "BfdSlowpath",
	RxReasonBfd:                        "Bfd",
	RxReasonMirror:                     "Mirror",
	RxReasonRegexAction:                "RegexAction",
	RxReasonRegexMatch:                 "RegexMatch",
	RxReasonFailoverDrop:               "FailoverDrop",
	RxReasonWlanTunnelError:            "WlanTunnelError",
	RxReasonCongestionCnmProxy:         "CongestionCnmProxy",
	RxReasonCongestionCnmProxyError:    "CongestionCnmProxyError",
	RxReasonCongestionCnm:              "CongestionCnm",
	RxReasonMplsUnknownAch:             "MplsUnknownAch",
	RxReasonMplsLookupsExceeded:        "MplsLookupsExceeded",
	RxReasonMplsReservedEntropyLabel:   "MplsReservedEntropyLabel",
	RxReasonMplsIllegalReservedLabel:   "MplsIllegalReservedLabel",
	RxReasonMplsRouterAlertLabel:       "MplsRouterAlertLabel",
	RxReasonNivPrune:                   "NivPrune",
	RxReasonVirtualPortPrune:           "VirtualPortPrune",
	RxReasonNonUnicastDrop:             "NonUnicastDrop",
	RxReasonTrillPacketPortMismatch:    "TrillPacketPortMismatch",
	RxReasonWlanClientMove:             "WlanClientMove",
	RxReasonWlanSourcePortMiss:         "WlanSourcePortMiss",
	RxReasonWlanClientSourceMiss:       "WlanClientSourceMiss",
	RxReasonWlanClientDestMiss:         "WlanClientDestMiss",
	RxReasonWlanMtu:                    "WlanMtu",
	RxReasonL2GreSipMiss:               "L2GreSipMiss",
	RxReasonL2GreVpnIdMiss:             "L2GreVpnIdMiss",
	RxReasonTimesyncUnknownVersion:     "TimesyncUnknownVersion",
	RxReasonBfdUnknownVersion:          "BfdUnknownVersion",
	RxReasonBfdInvalidVersion:          "BfdInvalidVersion",
	RxReasonBfdLookupFailure:           "BfdLookupFailure",
	RxReasonBfdInvalidPacket:           "BfdInvalidPacket",
	RxReasonVxlanSipMiss:               "VxlanSipMiss",
	RxReasonVxlanVpnIdMiss:             "VxlanVpnIdMiss",
	RxReasonFcoeZoneCheckFail:          "FcoeZoneCheckFail",
	RxReasonIpmcInterfaceMismatch:      "IpmcInterfaceMismatch",
	RxReasonNat:                        "Nat",
	RxReasonTcpUdpNatMiss:              "TcpUdpNatMiss",
	RxReasonIcmpNatMiss:                "IcmpNatMiss",
	RxReasonNatFragment:                "NatFragment",
	RxReasonNatMiss:                    "NatMiss",
	RxReasonOAMCCMSlowpath:             "OAMCCMSlowpath",
	RxReasonBHHOAM:                     "BHHOAM",
	RxReasonUnknownSubtendingPort:      "UnknownSubtendingPort",
	RxReasonReserved0:                  "Reserved0",
	RxReasonOAMMplsLmDm:                "OAMMplsLmDm",
	RxReasonSat:                        "Sat",
	RxReasonSampleSourceFlex:           "SampleSourceFlex",
	RxReasonFlexSflow:                  "FlexSflow",
	RxReasonVxltMiss:                   "VxltMiss",
	RxReasonTunnelDecapEcnError:        "TunnelDecapEcnError",
	RxReasonTunnelObjectValidationFail: "TunnelObjectValidationFail",
	RxReasonL3Cpu:                      "L3Cpu",
	RxReasonTunnelAdaptLookupMiss:      "TunnelAdaptLookupMiss",
	RxReasonPacketFlowSelectMiss:       "PacketFlowSelectMiss",
	RxReasonProtectionDataDrop:         "ProtectionDataDrop",
	RxReasonPacketFlowSelect:           "PacketFlowSelect",
	RxReasonOtherLookupMiss:            "OtherLookupMiss",
	RxReasonInvalidTpid:                "InvalidTpid",
	RxReasonMplsControlPacket:          "MplsControlPacket",
	RxReasonTunnelTtlError:             "TunnelTtlError",
	RxReasonL2HeaderError:              "L2HeaderError",
	RxReasonOtherLookupHit:             "OtherLookupHit",
	RxReasonL2SrcLookupMiss:            "L2SrcLookupMiss",
	RxReasonL2SrcLookupHit:             "L2SrcLookupHit",
	RxReasonL2DstLookupMiss:            "L2DstLookupMiss",
	RxReasonL2DstLookupHit:             "L2DstLookupHit",
	RxReasonL3SrcRouteLookupMiss:       "L3SrcRouteLookupMiss",
	RxReasonL3SrcHostLookupMiss:        "L3SrcHostLookupMiss",
	RxReasonL3SrcRouteLookupHit:        "L3SrcRouteLookupHit",
	RxReasonL3SrcHostLookupHit:         "L3SrcHostLookupHit",
	RxReasonL3DstRouteLookupMiss:       "L3DstRouteLookupMiss",
	RxReasonL3DstHostLookupMiss:        "L3DstHostLookupMiss",
	RxReasonL3DstRouteLookupHit:        "L3DstRouteLookupHit",
	RxReasonL3DstHostLookupHit:         "L3DstHostLookupHit",
	RxReasonVlanTranslate1Lookup1Miss:  "VlanTranslate1Lookup1Miss",
	RxReasonVlanTranslate1Lookup2Miss:  "VlanTranslate1Lookup2Miss",
	RxReasonMplsLookup1Miss:            "MplsLookup1Miss",
	RxReasonMplsLookup2Miss:            "MplsLookup2Miss",
	RxReasonL3TunnelLookupMiss:         "L3TunnelLookupMiss",
	RxReasonVlanTranslate2Lookup1Miss:  "VlanTranslate2Lookup1Miss",
	RxReasonVlanTranslate2Lookup2Miss:  "VlanTranslate2Lookup2Miss",
	RxReasonL2StuFail:                  "L2StuFail",
	RxReasonSrCounterExceeded:          "SrCounterExceeded",
	RxReasonSrCopyToCpuBit0:            "SrCopyToCpuBit0",
	RxReasonSrCopyToCpuBit1:            "SrCopyToCpuBit1",
	RxReasonSrCopyToCpuBit2:            "SrCopyToCpuBit2",
	RxReasonSrCopyToCpuBit3:            "SrCopyToCpuBit3",
	RxReasonSrCopyToCpuBit4:            "SrCopyToCpuBit4",
	RxReasonSrCopyToCpuBit5:            "SrCopyToCpuBit5",
	RxReasonL3HeaderMismatch:           "L3HeaderMismatch",
}

func ParseRxReason(s string) (RxReason, error) {
	if v, ok := rxReason_values[s]; ok {
		return v, nil
	}
	return 0, fmt.Errorf("Invalid RxReason. %s", s)
}

var rxReason_values = map[string]RxReason{
	"Invalid":                    RxReasonInvalid,
	"Arp":                        RxReasonArp,
	"Bpdu":                       RxReasonBpdu,
	"Broadcast":                  RxReasonBroadcast,
	"ClassBasedMove":             RxReasonClassBasedMove,
	"ClassTagPackets":            RxReasonClassTagPackets,
	"Control":                    RxReasonControl,
	"CpuLearn":                   RxReasonCpuLearn,
	"DestLookupFail":             RxReasonDestLookupFail,
	"Dhcp":                       RxReasonDhcp,
	"DosAttack":                  RxReasonDosAttack,
	"E2eHolIbp":                  RxReasonE2eHolIbp,
	"EncapHigigError":            RxReasonEncapHigigError,
	"FilterMatch":                RxReasonFilterMatch,
	"GreChecksum":                RxReasonGreChecksum,
	"GreSourceRoute":             RxReasonGreSourceRoute,
	"HigigControl":               RxReasonHigigControl,
	"HigigHdrError":              RxReasonHigigHdrError,
	"IcmpRedirect":               RxReasonIcmpRedirect,
	"Igmp":                       RxReasonIgmp,
	"IngressFilter":              RxReasonIngressFilter,
	"Ip":                         RxReasonIp,
	"IpfixRateViolation":         RxReasonIpfixRateViolation,
	"IpMcastMiss":                RxReasonIpMcastMiss,
	"IpmcReserved":               RxReasonIpmcReserved,
	"IpOptionVersion":            RxReasonIpOptionVersion,
	"Ipmc":                       RxReasonIpmc,
	"L2Cpu":                      RxReasonL2Cpu,
	"L2DestMiss":                 RxReasonL2DestMiss,
	"L2LearnLimit":               RxReasonL2LearnLimit,
	"L2Move":                     RxReasonL2Move,
	"L2MtuFail":                  RxReasonL2MtuFail,
	"L2NonUnicastMiss":           RxReasonL2NonUnicastMiss,
	"L2SourceMiss":               RxReasonL2SourceMiss,
	"L3AddrBindFail":             RxReasonL3AddrBindFail,
	"L3DestMiss":                 RxReasonL3DestMiss,
	"L3HeaderError":              RxReasonL3HeaderError,
	"L3MtuFail":                  RxReasonL3MtuFail,
	"L3Slowpath":                 RxReasonL3Slowpath,
	"L3SourceMiss":               RxReasonL3SourceMiss,
	"L3SourceMove":               RxReasonL3SourceMove,
	"MartianAddr":                RxReasonMartianAddr,
	"McastIdxError":              RxReasonMcastIdxError,
	"McastMiss":                  RxReasonMcastMiss,
	"MimServiceError":            RxReasonMimServiceError,
	"MplsCtrlWordError":          RxReasonMplsCtrlWordError,
	"MplsError":                  RxReasonMplsError,
	"MplsInvalidAction":          RxReasonMplsInvalidAction,
	"MplsInvalidPayload":         RxReasonMplsInvalidPayload,
	"MplsLabelMiss":              RxReasonMplsLabelMiss,
	"MplsSequenceNumber":         RxReasonMplsSequenceNumber,
	"MplsTtl":                    RxReasonMplsTtl,
	"Multicast":                  RxReasonMulticast,
	"Nhop":                       RxReasonNhop,
	"OAMError":                   RxReasonOAMError,
	"OAMSlowpath":                RxReasonOAMSlowpath,
	"OAMLMDM":                    RxReasonOAMLMDM,
	"ParityError":                RxReasonParityError,
	"Protocol":                   RxReasonProtocol,
	"SampleDest":                 RxReasonSampleDest,
	"SampleSource":               RxReasonSampleSource,
	"SharedVlanMismatch":         RxReasonSharedVlanMismatch,
	"SourceRoute":                RxReasonSourceRoute,
	"TimeStamp":                  RxReasonTimeStamp,
	"Ttl":                        RxReasonTtl,
	"Ttl1":                       RxReasonTtl1,
	"TunnelError":                RxReasonTunnelError,
	"UdpChecksum":                RxReasonUdpChecksum,
	"UnknownVlan":                RxReasonUnknownVlan,
	"UrpfFail":                   RxReasonUrpfFail,
	"VcLabelMiss":                RxReasonVcLabelMiss,
	"VlanFilterMatch":            RxReasonVlanFilterMatch,
	"WlanClientError":            RxReasonWlanClientError,
	"WlanSlowpath":               RxReasonWlanSlowpath,
	"WlanDot1xDrop":              RxReasonWlanDot1xDrop,
	"ExceptionFlood":             RxReasonExceptionFlood,
	"TimeSync":                   RxReasonTimeSync,
	"EAVData":                    RxReasonEAVData,
	"SamePortBridge":             RxReasonSamePortBridge,
	"SplitHorizon":               RxReasonSplitHorizon,
	"L4Error":                    RxReasonL4Error,
	"Stp":                        RxReasonStp,
	"EgressFilterRedirect":       RxReasonEgressFilterRedirect,
	"FilterRedirect":             RxReasonFilterRedirect,
	"Loopback":                   RxReasonLoopback,
	"VlanTranslate":              RxReasonVlanTranslate,
	"Mmrp":                       RxReasonMmrp,
	"Srp":                        RxReasonSrp,
	"TunnelControl":              RxReasonTunnelControl,
	"L2Marked":                   RxReasonL2Marked,
	"WlanSlowpathKeepalive":      RxReasonWlanSlowpathKeepalive,
	"Station":                    RxReasonStation,
	"Niv":                        RxReasonNiv,
	"NivPrioDrop":                RxReasonNivPrioDrop,
	"NivInterfaceMiss":           RxReasonNivInterfaceMiss,
	"NivRpfFail":                 RxReasonNivRpfFail,
	"NivTagInvalid":              RxReasonNivTagInvalid,
	"NivTagDrop":                 RxReasonNivTagDrop,
	"NivUntagDrop":               RxReasonNivUntagDrop,
	"Trill":                      RxReasonTrill,
	"TrillInvalid":               RxReasonTrillInvalid,
	"TrillMiss":                  RxReasonTrillMiss,
	"TrillRpfFail":               RxReasonTrillRpfFail,
	"TrillSlowpath":              RxReasonTrillSlowpath,
	"TrillCoreIsIs":              RxReasonTrillCoreIsIs,
	"TrillTtl":                   RxReasonTrillTtl,
	"TrillName":                  RxReasonTrillName,
	"BfdSlowpath":                RxReasonBfdSlowpath,
	"Bfd":                        RxReasonBfd,
	"Mirror":                     RxReasonMirror,
	"RegexAction":                RxReasonRegexAction,
	"RegexMatch":                 RxReasonRegexMatch,
	"FailoverDrop":               RxReasonFailoverDrop,
	"WlanTunnelError":            RxReasonWlanTunnelError,
	"CongestionCnmProxy":         RxReasonCongestionCnmProxy,
	"CongestionCnmProxyError":    RxReasonCongestionCnmProxyError,
	"CongestionCnm":              RxReasonCongestionCnm,
	"MplsUnknownAch":             RxReasonMplsUnknownAch,
	"MplsLookupsExceeded":        RxReasonMplsLookupsExceeded,
	"MplsReservedEntropyLabel":   RxReasonMplsReservedEntropyLabel,
	"MplsIllegalReservedLabel":   RxReasonMplsIllegalReservedLabel,
	"MplsRouterAlertLabel":       RxReasonMplsRouterAlertLabel,
	"NivPrune":                   RxReasonNivPrune,
	"VirtualPortPrune":           RxReasonVirtualPortPrune,
	"NonUnicastDrop":             RxReasonNonUnicastDrop,
	"TrillPacketPortMismatch":    RxReasonTrillPacketPortMismatch,
	"WlanClientMove":             RxReasonWlanClientMove,
	"WlanSourcePortMiss":         RxReasonWlanSourcePortMiss,
	"WlanClientSourceMiss":       RxReasonWlanClientSourceMiss,
	"WlanClientDestMiss":         RxReasonWlanClientDestMiss,
	"WlanMtu":                    RxReasonWlanMtu,
	"L2GreSipMiss":               RxReasonL2GreSipMiss,
	"L2GreVpnIdMiss":             RxReasonL2GreVpnIdMiss,
	"TimesyncUnknownVersion":     RxReasonTimesyncUnknownVersion,
	"BfdUnknownVersion":          RxReasonBfdUnknownVersion,
	"BfdInvalidVersion":          RxReasonBfdInvalidVersion,
	"BfdLookupFailure":           RxReasonBfdLookupFailure,
	"BfdInvalidPacket":           RxReasonBfdInvalidPacket,
	"VxlanSipMiss":               RxReasonVxlanSipMiss,
	"VxlanVpnIdMiss":             RxReasonVxlanVpnIdMiss,
	"FcoeZoneCheckFail":          RxReasonFcoeZoneCheckFail,
	"IpmcInterfaceMismatch":      RxReasonIpmcInterfaceMismatch,
	"Nat":                        RxReasonNat,
	"TcpUdpNatMiss":              RxReasonTcpUdpNatMiss,
	"IcmpNatMiss":                RxReasonIcmpNatMiss,
	"NatFragment":                RxReasonNatFragment,
	"NatMiss":                    RxReasonNatMiss,
	"OAMCCMSlowpath":             RxReasonOAMCCMSlowpath,
	"BHHOAM":                     RxReasonBHHOAM,
	"UnknownSubtendingPort":      RxReasonUnknownSubtendingPort,
	"Reserved0":                  RxReasonReserved0,
	"OAMMplsLmDm":                RxReasonOAMMplsLmDm,
	"Sat":                        RxReasonSat,
	"SampleSourceFlex":           RxReasonSampleSourceFlex,
	"FlexSflow":                  RxReasonFlexSflow,
	"VxltMiss":                   RxReasonVxltMiss,
	"TunnelDecapEcnError":        RxReasonTunnelDecapEcnError,
	"TunnelObjectValidationFail": RxReasonTunnelObjectValidationFail,
	"L3Cpu":                      RxReasonL3Cpu,
	"TunnelAdaptLookupMiss":      RxReasonTunnelAdaptLookupMiss,
	"PacketFlowSelectMiss":       RxReasonPacketFlowSelectMiss,
	"ProtectionDataDrop":         RxReasonProtectionDataDrop,
	"PacketFlowSelect":           RxReasonPacketFlowSelect,
	"OtherLookupMiss":            RxReasonOtherLookupMiss,
	"InvalidTpid":                RxReasonInvalidTpid,
	"MplsControlPacket":          RxReasonMplsControlPacket,
	"TunnelTtlError":             RxReasonTunnelTtlError,
	"L2HeaderError":              RxReasonL2HeaderError,
	"OtherLookupHit":             RxReasonOtherLookupHit,
	"L2SrcLookupMiss":            RxReasonL2SrcLookupMiss,
	"L2SrcLookupHit":             RxReasonL2SrcLookupHit,
	"L2DstLookupMiss":            RxReasonL2DstLookupMiss,
	"L2DstLookupHit":             RxReasonL2DstLookupHit,
	"L3SrcRouteLookupMiss":       RxReasonL3SrcRouteLookupMiss,
	"L3SrcHostLookupMiss":        RxReasonL3SrcHostLookupMiss,
	"L3SrcRouteLookupHit":        RxReasonL3SrcRouteLookupHit,
	"L3SrcHostLookupHit":         RxReasonL3SrcHostLookupHit,
	"L3DstRouteLookupMiss":       RxReasonL3DstRouteLookupMiss,
	"L3DstHostLookupMiss":        RxReasonL3DstHostLookupMiss,
	"L3DstRouteLookupHit":        RxReasonL3DstRouteLookupHit,
	"L3DstHostLookupHit":         RxReasonL3DstHostLookupHit,
	"VlanTranslate1Lookup1Miss":  RxReasonVlanTranslate1Lookup1Miss,
	"VlanTranslate1Lookup2Miss":  RxReasonVlanTranslate1Lookup2Miss,
	"MplsLookup1Miss":            RxReasonMplsLookup1Miss,
	"MplsLookup2Miss":            RxReasonMplsLookup2Miss,
	"L3TunnelLookupMiss":         RxReasonL3TunnelLookupMiss,
	"VlanTranslate2Lookup1Miss":  RxReasonVlanTranslate2Lookup1Miss,
	"VlanTranslate2Lookup2Miss":  RxReasonVlanTranslate2Lookup2Miss,
	"L2StuFail":                  RxReasonL2StuFail,
	"SrCounterExceeded":          RxReasonSrCounterExceeded,
	"SrCopyToCpuBit0":            RxReasonSrCopyToCpuBit0,
	"SrCopyToCpuBit1":            RxReasonSrCopyToCpuBit1,
	"SrCopyToCpuBit2":            RxReasonSrCopyToCpuBit2,
	"SrCopyToCpuBit3":            RxReasonSrCopyToCpuBit3,
	"SrCopyToCpuBit4":            RxReasonSrCopyToCpuBit4,
	"SrCopyToCpuBit5":            RxReasonSrCopyToCpuBit5,
	"L3HeaderMismatch":           RxReasonL3HeaderMismatch,
}
