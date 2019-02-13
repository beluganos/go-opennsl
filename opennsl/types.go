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
*/
import "C"

import (
	"fmt"
	"net"
)

//
// opennsl_port_t
//
type Port C.opennsl_port_t

func (v Port) C() C.opennsl_port_t {
	return C.opennsl_port_t(v)
}

//
// opennsl_if_t(int)
//
type Iface C.opennsl_if_t

func (v Iface) C() C.opennsl_if_t {
	return C.opennsl_if_t(v)
}

const IFACE_NONE Iface = 0

//
// EncapID
//
type EncapID C.opennsl_if_t

func (v EncapID) C() C.opennsl_if_t {
	return C.opennsl_if_t(v)
}

//
// opennsl_trill_name_t(int)
//
type TrillName C.opennsl_trill_name_t

func (v TrillName) C() C.opennsl_trill_name_t {
	return C.opennsl_trill_name_t(v)
}

//
// opennsl_l4_port_t(int)
//
type L4Port C.opennsl_l4_port_t

func (v L4Port) C() C.opennsl_l4_port_t {
	return C.opennsl_l4_port_t(v)
}

//
// opennsl_if_group_t
//
type IfaceGroup C.opennsl_if_group_t

func (v IfaceGroup) C() C.opennsl_if_group_t {
	return C.opennsl_if_group_t(v)
}

//
// opennsl_vrf_t
//
type Vrf C.opennsl_vrf_t

func (v Vrf) C() C.opennsl_vrf_t {
	return C.opennsl_vrf_t(v)
}

func (v Vrf) IsValid() bool {
	return (v != VRF_NONE)
}

const VRF_NONE = Vrf(0)

//
// opennsl_mpls_label_t
//
type MplsLabel C.opennsl_mpls_label_t

func (v MplsLabel) C() C.opennsl_mpls_label_t {
	return C.opennsl_mpls_label_t(v)
}

//
// opennsl_vlan_t -> vlan_control.go
//

//
// opennsl_ethertype_t
//
type Ethertype C.opennsl_ethertype_t

func (v Ethertype) C() C.opennsl_ethertype_t {
	return C.opennsl_ethertype_t(v)
}

//
// opennsl_vpn_t
//
type Vpn C.opennsl_vpn_t

func (v Vpn) C() C.opennsl_vpn_t {
	return C.opennsl_vpn_t(v)
}

//
// opennsl_tunnel_id_t
//
/*
type TunnelID C.opennsl_tunnel_id_t

func (v TunnelID) C() C.opennsl_tunnel_id_t {
	return C.opennsl_tunnel_id_t(v)
}
*/

//
// opennsl_cos_t
//
type Cos C.opennsl_cos_t

func (v Cos) C() C.opennsl_cos_t {
	return C.opennsl_cos_t(v)
}

const (
	COS_COUNT   Cos = C.OPENNSL_COS_COUNT
	COS_DEFAULT Cos = C.OPENNSL_COS_DEFAULT
	COS_INVALID Cos = C.OPENNSL_COS_INVALID
	COS_MAX     Cos = C.OPENNSL_COS_MAX
	COS_MIN     Cos = C.OPENNSL_COS_MIN
)

//
// opennsl_cos_queue_t
//
type CosQueue C.opennsl_cos_queue_t

func (v CosQueue) C() C.opennsl_cos_queue_t {
	return C.opennsl_cos_queue_t(v)
}

//
// opennsl_module_t
//
type Module C.opennsl_module_t

func (v Module) C() C.opennsl_module_t {
	return C.opennsl_module_t(v)
}

//
// FabricDistribution
//
type FabricDistribution C.opennsl_fabric_distribution_t

func (v FabricDistribution) C() C.opennsl_fabric_distribution_t {
	return C.opennsl_fabric_distribution_t(v)
}

//
// opennsl_failover_t
//
type Failover C.opennsl_failover_t

func (v Failover) C() C.opennsl_failover_t {
	return C.opennsl_failover_t(v)
}

//
// opennsl_color_t
//
type Color C.opennsl_color_t

func (v Color) C() C.opennsl_color_t {
	return C.opennsl_color_t(v)
}

const (
	COLOR_NONE       Color = 0
	COLOR_GREEN      Color = C.opennslColorGreen
	COLOR_YELLOW     Color = C.opennslColorYellow
	COLOR_RED        Color = C.opennslColorRed
	COLOR_DROP_FIRST Color = C.opennslColorDropFirst
	COLOR_BLACK      Color = C.opennslColorBlack
	COLOR_PRESERVE   Color = C.opennslColorPreserve
)

const COLOR_COUNT = C.opennslColorCount

var color_names = map[Color]string{
	COLOR_GREEN:    "GREEN",
	COLOR_YELLOW:   "YELLOW",
	COLOR_RED:      "RED",
	COLOR_BLACK:    "BLACK",
	COLOR_PRESERVE: "PRESERVE",
	// COLOR_DROP_FIRST: "DROP_FIRST",
}

var color_values = map[string]Color{
	"GREEN":      COLOR_GREEN,
	"YELLOW":     COLOR_YELLOW,
	"RED":        COLOR_RED,
	"DROP_FIRST": COLOR_DROP_FIRST,
	"BLACK":      COLOR_BLACK,
	"PRESERVE":   COLOR_PRESERVE,
}

func (v Color) String() string {
	if s, ok := color_names[v]; ok {
		return s
	}
	return fmt.Sprintf("Color(%d)", v)
}

func ParseColor(s string) (Color, error) {
	if v, ok := color_values[s]; ok {
		return v, nil
	}
	return COLOR_NONE, fmt.Errorf("Invalid Color. &%s", s)
}

//
// opennsl_policer_t
//
type Policer C.opennsl_policer_t

func (v Policer) C() C.opennsl_policer_t {
	return C.opennsl_policer_t(v)
}

//
// opennsl_mac_t
//
func NewMAC(hwaddr net.HardwareAddr) C.opennsl_mac_t {
	l := len(hwaddr)
	if l > 6 {
		l = 6
	}

	mac := C.opennsl_mac_t{}
	for i := 0; i < l; i++ {
		mac[i] = C.uint8(hwaddr[i])
	}

	return mac
}

func ParseMAC(mac C.opennsl_mac_t) net.HardwareAddr {
	b := make([]byte, 6)
	for i := 0; i < 6; i++ {
		b[i] = byte(mac[i])
	}

	return net.HardwareAddr(b)
}

func HardwareAddrIsMC(mac net.HardwareAddr) bool {
	if len(mac) < 3 {
		return false
	}

	return (mac[0] == 0x01) && (mac[1] == 0x00) && (mac[2] == 0x5e)
}

//
// opennsl_ip_t (IPv4)
//
func ip4ToC(ip net.IP) C.opennsl_ip_t {
	ip4 := (uint32(ip[0]) << 24) + (uint32(ip[1]) << 16) + (uint32(ip[2]) << 8) + uint32(ip[3])
	return C.opennsl_ip_t(ip4)
}

func NewIP4(ip net.IP) (C.opennsl_ip_t, error) {
	ip4 := ip.To4()
	if ip4 == nil {
		return 0, fmt.Errorf("%s is not IPv4.", ip)
	}

	return ip4ToC(ip4), nil
}

func ParseIP4(ip C.opennsl_ip_t) net.IP {
	return net.IPv4(byte((ip>>24)&0xff), byte((ip>>16)&0xff), byte((ip>>8)&0xff), byte(ip&0xff))
}

func NewIP4Mask(mask net.IPMask) (C.opennsl_ip_t, error) {
	ip := net.IP(mask)
	return NewIP4(ip)
}

func CreateIP4Mask(plen int) (C.opennsl_ip_t, error) {
	return C.opennsl_ip_mask_create(C.int(plen)), nil
}

func ParseIP4Mask(mask C.opennsl_ip_t) net.IPMask {
	ip := ParseIP4(mask)
	return net.IPMask(ip)
}

func NewIP4AndMask(ip net.IP, mask net.IPMask) (c_ip C.opennsl_ip_t, c_mask C.opennsl_ip_t, err error) {
	if c_ip, err = NewIP4(ip); err != nil {
		return
	}

	if c_mask, err = NewIP4Mask(mask); err != nil {
		return
	}

	err = nil
	return
}

func ParseIP4AndMask(c_ip C.opennsl_ip_t, c_mask C.opennsl_ip_t) (net.IP, net.IPMask) {
	return ParseIP4(c_ip), ParseIP4Mask(c_mask)
}

//
// C.opennsl_ip6_t (IPv6)
//
func ip6ToC(ip net.IP) C.opennsl_ip6_t {
	ip6 := [16]C.uint8{}
	for i := 0; i < 16; i++ {
		ip6[i] = C.uint8(ip[i])
	}
	return C.opennsl_ip6_t(ip6)
}

func NewIP6(ip net.IP) (C.opennsl_ip6_t, error) {
	ip6 := ip.To16()
	if ip6 == nil {
		return C.opennsl_ip6_t{}, fmt.Errorf("%s is not IPv6.", ip)
	}

	return ip6ToC(ip6), nil
}

func ParseIP6(ip C.opennsl_ip6_t) net.IP {
	b := make([]byte, 16)
	for i := 0; i < 16; i++ {
		b[i] = byte(ip[i])
	}
	return net.IP(b)
}

func NewIP6Mask(mask net.IPMask) (C.opennsl_ip6_t, error) {
	ip := net.IP(mask)
	return NewIP6(ip)
}

func CreateIP6Mask(plen int) (C.opennsl_ip6_t, error) {
	ip6 := [16]C.uint8{}
	rc := C.opennsl_ip6_mask_create(&ip6[0], C.int(plen))
	if err := ParseError(rc); err != nil {
		return C.opennsl_ip6_t(ip6), err
	}

	return C.opennsl_ip6_t(ip6), nil
}

func ParseIP6Mask(mask C.opennsl_ip6_t) net.IPMask {
	ip := ParseIP6(mask)
	return net.IPMask(ip)
}

func NewIP6AndMask(ip net.IP, mask net.IPMask) (c_ip C.opennsl_ip6_t, c_mask C.opennsl_ip6_t, err error) {
	if c_ip, err = NewIP6(ip); err != nil {
		return
	}

	if c_mask, err = NewIP6Mask(mask); err != nil {
		return
	}

	err = nil
	return
}

func ParseIP6AndMask(c_ip C.opennsl_ip6_t, c_mask C.opennsl_ip6_t) (net.IP, net.IPMask) {
	return ParseIP6(c_ip), ParseIP6Mask(c_mask)
}
