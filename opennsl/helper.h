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

#ifndef _GO_OPENNSL_HELPER_H
#define _GO_OPENNSL_HELPER_H

#include <opennsl/port.h>
#include <opennsl/l2.h>
#include <opennsl/l3.h>
#include <opennsl/rx.h>
#include <opennsl/switch.h>
#include <opennsl/vlan.h>
#include <opennsl/fieldX.h>
#include <opennsl/stg.h>
#include <opennsl/multicast.h>
#include <opennsl/vxlanX.h>

#ifdef __cplusplus
extern "C" {
#endif

  void _opennsl_l2_addr_cb(int unit, opennsl_l2_addr_t *l2addr, int operation,  void *userdata);

  int _opennsl_l2_traverse_cb(int unit, opennsl_l2_addr_t *info, void *user_data);

  void _opennsl_switch_event_cb(int unit, opennsl_switch_event_t event, uint32 arg1, uint32 arg2, uint32 arg3, void *userdata);

  int _opennsl_l3_egress_traverse_cb(int, opennsl_if_t, opennsl_l3_egress_t*, void*);

  int _opennsl_l3_egress_ecmp_traverse_cb(int, opennsl_l3_egress_ecmp_t*, int, opennsl_if_t*, void*);

  int _opennsl_l3_host_traverse_cb(int, int, opennsl_l3_host_t*, void*);

  int _opennsl_l3_route_traverse_cb(int, int, opennsl_l3_route_t*, void*);

  int _opennsl_field_group_traverse_cb(int, opennsl_field_group_t, void*);

  int _opennsl_mpls_vpn_traverse_cb(int, opennsl_mpls_vpn_config_t*, void*);

  int _opennsl_mpls_tunnel_switch_traverse_cb(int, opennsl_mpls_tunnel_switch_t*, void*);

  opennsl_rx_t _opennsl_rx_cb(int, opennsl_pkt_t*, void*);

  int _opennsl_multicast_group_traverse_cb(int, opennsl_multicast_t, uint32, void*);

  void _opennsl_linkscan_handler(int, opennsl_port_t, opennsl_port_info_t*);

  int _opennsl_tunnel_initiator_traverse_cb(int, opennsl_tunnel_initiator_t*, void*);

  int _opennsl_tunnel_terminator_traverse_cb(int, opennsl_tunnel_terminator_t*, void*);

  int _opennsl_vxlan_vpn_traverse_cb(int, opennsl_vxlan_vpn_config_t*, void*);

  int _opennsl_vlan_list_iter(int unit, void* p);

  int _opennsl_stg_vlan_list_iter(int unit, opennsl_stg_t stg, void* p);

  int _opennsl_stg_list_iter(int unit, void* p);

  opennsl_pkt_blk_t* _opennsl_pkt_data_get(opennsl_pkt_blk_t* pkt_data, uint8 blk_count, uint8 index);

  int* _opennsl_util_int_array_get(int* arr, uint32 count, uint32 index);

  // OPENNSL_PBMP_CLEAR
  void _opennsl_pbmp_clear(opennsl_pbmp_t* pbmp);
  // OPENNSL_PBMP_MEMBER
  int _opennsl_pbmp_member(opennsl_pbmp_t* pbmp, int port);
  // OPENNSL_PBMP_PORT_SET
  void _opennsl_pbmp_port_set(opennsl_pbmp_t* pbmp, int port);
  // OPENNSL_PBMP_PORT_ADD
  void _opennsl_pbmp_port_add(opennsl_pbmp_t* pbmp, int port);
  // OPENNSL_FIELD_QSET_INIT
  void _opennsl_field_qset_init(opennsl_field_qset_t* qset);
  // OPENNSL_FIELD_QSET_ADD
  void _opennsl_field_qset_add(opennsl_field_qset_t* qset, opennsl_field_qualify_t q);
  // OPENNSL_FIELD_QSET_REMOVE
  void _opennsl_field_qset_remove(opennsl_field_qset_t* qset, opennsl_field_qualify_t q);
  // OPENNSL_FIELD_ASET_TEST
  uint32 _opennsl_field_qset_test(const opennsl_field_qset_t* qset, opennsl_field_qualify_t q);
  // OPENNSL_FIELD_ASET_INIT
  void _opennsl_field_aset_init(opennsl_field_aset_t* aset);
  // OPENNSL_FIELD_ASET_ADD
  void _opennsl_field_aset_add(opennsl_field_aset_t* aset, opennsl_field_action_t a);
  // OPENNSL_FIELD_ASET_REMOVE
  void _opennsl_field_aset_remove(opennsl_field_aset_t* aset, opennsl_field_action_t a);
  // OPENNSL_FIELD_ASET_TEST
  uint32 _opennsl_field_aset_test(const opennsl_field_aset_t* aset, opennsl_field_action_t a);
  // OPENNSL_PORT_ABIL_SPD_MAX
  int _opennsl_port_abil_spd_max(opennsl_port_abil_t);
  // OPENNSL_PORT_ABILITY_SPEED_MAX
  int _opennsl_port_ability_speed_max(opennsl_port_abil_t);

  ////// OPENNSL_GPORT_* (type.h) //////
  // OPENNSL_GPORT_LOCAL_SET
  opennsl_gport_t _opennsl_gport_from_local(opennsl_port_t);
  // OPENNSL_GPORT_LOCAL_GET
  opennsl_port_t _opennsl_gport_to_local(opennsl_gport_t);
  // OPENNSL_GPORT_MODPORT_NODID_GET
  opennsl_module_t _opennsl_gport_modport_to_modid(opennsl_gport_t);
  // OPENNSL_GPORT_MODPORT_PORT_GET
  opennsl_port_t _opennsl_gport_modport_to_port(opennsl_gport_t);
  // OPENNSL_GPORT_MODPORT_SET
  opennsl_gport_t _opennsl_gport_from_modid_and_port(opennsl_port_t, opennsl_module_t);

  // OPENNSL_RX_REASON_CLEAR_ALL
  void _opennsl_rx_reason_clear_all(opennsl_rx_reasons_t* reasons);
  // OPENNSL_RX_REASON_GET
  int _opennsl_rx_reason_get(const opennsl_rx_reasons_t* reasons, opennsl_rx_reason_t reason);
  // OPENNSL_RX_REASON_SET(
  void _opennsl_rx_reason_set(opennsl_rx_reasons_t* reasons, opennsl_rx_reason_t reason);

#ifdef __cplusplus
}
#endif

#endif // _GO_OPENNSL_HELPER_H
