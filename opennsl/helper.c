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

#include <string.h>
#include "helper.h"

extern void go_opennsl_l2_addr_cb(int unit, opennsl_l2_addr_t *l2addr, int operation,  void *userdata);

void _opennsl_l2_addr_cb(int unit, opennsl_l2_addr_t *l2addr, int operation,  void *userdata) {
  opennsl_l2_addr_t _l2addr = *l2addr;
  go_opennsl_l2_addr_cb(unit, &_l2addr, operation, userdata);
}

extern int go_opennsl_l2_traverse_cb(int unit, opennsl_l2_addr_t *info, void *user_data);

int _opennsl_l2_traverse_cb(int unit, opennsl_l2_addr_t *info, void *user_data) {
  opennsl_l2_addr_t _info = *info;
  return go_opennsl_l2_traverse_cb(unit, &_info, user_data);
}

extern int go_opennsl_l3_egress_traverse_cb(int, opennsl_if_t, opennsl_l3_egress_t*, void*);

int _opennsl_l3_egress_traverse_cb(int unit, opennsl_if_t intf, opennsl_l3_egress_t *info, void *user_data) {
  opennsl_l3_egress_t l3_egr = *info;
  return go_opennsl_l3_egress_traverse_cb(unit, intf, &l3_egr, user_data);
}

extern int go_opennsl_l3_egress_ecmp_traverse_cb(int, opennsl_l3_egress_ecmp_t*, int, opennsl_if_t*, void*);

int _opennsl_l3_egress_ecmp_traverse_cb(int unit, opennsl_l3_egress_ecmp_t* ecmp, int size, opennsl_if_t arr[], void* user_data) {
  return go_opennsl_l3_egress_ecmp_traverse_cb(unit, ecmp, size, arr, user_data);
}

extern int go_opennsl_l3_host_traverse_cb(int, int, opennsl_l3_host_t*, void*);

int _opennsl_l3_host_traverse_cb(int unit, int index, opennsl_l3_host_t* host, void* user_data) {
  return go_opennsl_l3_host_traverse_cb(unit, index, host, user_data);
}

extern int go_opennsl_l3_route_traverse_cb(int, int, opennsl_l3_route_t*, void*);

int _opennsl_l3_route_traverse_cb(int unit, int index, opennsl_l3_route_t* route, void* user_data) {
  return go_opennsl_l3_route_traverse_cb(unit, index, route, user_data);
}

extern int go_opennsl_field_group_traverse_cb(int, opennsl_field_group_t, void*);

int _opennsl_field_group_traverse_cb(int unit, opennsl_field_group_t group, void* user_data) {
  return go_opennsl_field_group_traverse_cb(unit, group, user_data);
}

extern int go_opennsl_mpls_vpn_traverse_cb(int, opennsl_mpls_vpn_config_t*, void*);

int _opennsl_mpls_vpn_traverse_cb(int unit, opennsl_mpls_vpn_config_t* config, void* user_data) {
  return go_opennsl_mpls_vpn_traverse_cb(unit, config, user_data);
}

extern opennsl_rx_t go_opennsl_rx_cb(int, opennsl_pkt_t*, void*);

opennsl_rx_t _opennsl_rx_cb(int unit, opennsl_pkt_t *pkt, void *cookie) {
  return go_opennsl_rx_cb(unit, pkt, cookie);
}

extern void go_opennsl_linkscan_handler(int, opennsl_port_t, opennsl_port_info_t*);

void _opennsl_linkscan_handler(int unit, opennsl_port_t port, opennsl_port_info_t* info) {
  go_opennsl_linkscan_handler(unit, port, info);
} 

void _opennsl_pbmp_clear(opennsl_pbmp_t* pbmp) {
  OPENNSL_PBMP_CLEAR((*pbmp));
}

int _opennsl_pbmp_member(opennsl_pbmp_t* pbmp, int port) {
  return OPENNSL_PBMP_MEMBER((*pbmp), port);
}

void _opennsl_pbmp_port_set(opennsl_pbmp_t* pbmp, int port) {
  OPENNSL_PBMP_PORT_SET((*pbmp), port);
}

void _opennsl_pbmp_port_add(opennsl_pbmp_t* pbmp, int port) {
  OPENNSL_PBMP_PORT_ADD((*pbmp), port);
}

void _opennsl_field_qset_init(opennsl_field_qset_t* qset) {
  OPENNSL_FIELD_QSET_INIT(*qset);
}

void _opennsl_field_qset_add(opennsl_field_qset_t* qset, opennsl_field_qualify_t q) {
  OPENNSL_FIELD_QSET_ADD(*qset, q);
}

void _opennsl_field_qset_remove(opennsl_field_qset_t* qset, opennsl_field_qualify_t q) {
  OPENNSL_FIELD_QSET_REMOVE(*qset, q);
}

uint32 _opennsl_field_qset_test(const opennsl_field_qset_t* qset, opennsl_field_qualify_t q) {
  return OPENNSL_FIELD_QSET_TEST(*qset, q);
}

void _opennsl_field_aset_init(opennsl_field_aset_t* aset) {
  OPENNSL_FIELD_ASET_INIT(*aset);
}

void _opennsl_field_aset_add(opennsl_field_aset_t* aset, opennsl_field_action_t a) {
  OPENNSL_FIELD_ASET_ADD(*aset, a);
}

void _opennsl_field_aset_remove(opennsl_field_aset_t* aset, opennsl_field_action_t a) {
  OPENNSL_FIELD_ASET_REMOVE(*aset, a);
}

uint32 _opennsl_field_aset_test(const opennsl_field_aset_t* aset, opennsl_field_action_t a) {
  return OPENNSL_FIELD_ASET_TEST(*aset, a);
}

extern void go_opennsl_switch_event_cb(int unit, opennsl_switch_event_t event, uint32 arg1, uint32 arg2, uint32 arg3, void *userdata);

void _opennsl_switch_event_cb(int unit, opennsl_switch_event_t event, uint32 arg1, uint32 arg2, uint32 arg3, void *userdata) {
  go_opennsl_switch_event_cb(unit, event, arg1, arg2, arg3, userdata);
}

extern int go_opennsl_mpls_tunnel_switch_traverse_cb(int, opennsl_mpls_tunnel_switch_t*, void*);

int _opennsl_mpls_tunnel_switch_traverse_cb(int unit, opennsl_mpls_tunnel_switch_t* info, void* data) {
  return go_opennsl_mpls_tunnel_switch_traverse_cb(unit, info, data);
}

extern int go_opennsl_vlan_list_cb(int unit, opennsl_vlan_data_t* data, void* p);

int _opennsl_vlan_list_iter(int unit, void* p) {
  int rc;

  opennsl_vlan_data_t* list = NULL;
  int count = 0;
  rc = opennsl_vlan_list(unit, &list, &count);
  if (rc != OPENNSL_E_NONE) {
    return rc;
  }

  int index;
  for (index = 0; index < count; index++) {
    rc = go_opennsl_vlan_list_cb(unit, list + index, p);
    if (rc != OPENNSL_E_NONE) {
      break;
    }
  }

  if (list != NULL) {
    return opennsl_vlan_list_destroy(unit, list, count);
  }

  return OPENNSL_E_NONE;
}

extern int go_opennsl_stg_vlan_list_cb(int unit, opennsl_stg_t stg, opennsl_vlan_t vlan, void* p);

int _opennsl_stg_vlan_list_iter(int unit, opennsl_stg_t stg, void* p) {
  int rc;

  opennsl_vlan_t* list = NULL;
  int count = 0;
  rc = opennsl_stg_vlan_list(unit, stg, &list, &count);
  if (rc != OPENNSL_E_NONE) {
    return rc;
  }

  int index;
  for (index = 0; index < count; index++) {
    rc = go_opennsl_stg_vlan_list_cb(unit, stg, list[index], p);
    if (rc != OPENNSL_E_NONE) {
      break;
    }
  }

  if (list != NULL) {
    return opennsl_stg_vlan_list_destroy(unit, list, count);
  }

  return OPENNSL_E_NONE;
}

extern int go_opennsl_stg_list_cb(int unit, opennsl_stg_t stg, void* p);

int _opennsl_stg_list_iter(int unit, void* p) {
  int rc;

  opennsl_stg_t* list = NULL;
  int count = 0;
  rc = opennsl_stg_list(unit, &list, &count);
  if (rc != OPENNSL_E_NONE) {
    return rc;
  }

  int index;
  for (index = 0; index < count; index++) {
    rc = go_opennsl_stg_list_cb(unit, list[index], p);
    if (rc != OPENNSL_E_NONE) {
      break;
    }
  }

  if (list != NULL) {
    return opennsl_stg_list_destroy(unit, list, count);
  }

  return OPENNSL_E_NONE;
}

extern int go_opennsl_multicast_group_traverse_cb(int, opennsl_multicast_t, uint32, void*);

int _opennsl_multicast_group_traverse_cb(int unit, opennsl_multicast_t group, uint32 flags, void* data) {
  return go_opennsl_multicast_group_traverse_cb(unit, group, flags, data);
}

extern int go_opennsl_tunnel_initiator_traverse_cb(int, opennsl_tunnel_initiator_t*, void*);

int _opennsl_tunnel_initiator_traverse_cb(int unit, opennsl_tunnel_initiator_t* tun, void* data) {
  return go_opennsl_tunnel_initiator_traverse_cb(unit, tun, data);
}

extern int go_opennsl_tunnel_terminator_traverse_cb(int, opennsl_tunnel_terminator_t*, void*);

int _opennsl_tunnel_terminator_traverse_cb(int unit, opennsl_tunnel_terminator_t* tun, void* data) {
  return go_opennsl_tunnel_terminator_traverse_cb(unit, tun, data);
}
extern int go_opennsl_vxlan_vpn_traverse_cb(int, opennsl_vxlan_vpn_config_t*, void*);

int _opennsl_vxlan_vpn_traverse_cb(int unit, opennsl_vxlan_vpn_config_t* config, void* data) {
  return go_opennsl_vxlan_vpn_traverse_cb(unit, config, data);
}
opennsl_pkt_blk_t* _opennsl_pkt_data_get(opennsl_pkt_blk_t* pkt_data, uint8 blk_count, uint8 index) {
  if (index >= blk_count) {
    return NULL;
  }
  return pkt_data + index;
}

int* _opennsl_util_int_array_get(int* arr, uint32 count, uint32 index) {
  if (index >= count)
    return NULL;

  return arr + index;
}

int _opennsl_port_abil_spd_max(opennsl_port_abil_t abil) {
  return OPENNSL_PORT_ABIL_SPD_MAX(abil);
}

int _opennsl_port_ability_speed_max(opennsl_port_abil_t abil) {
  return OPENNSL_PORT_ABILITY_SPEED_MAX(abil);
}

opennsl_gport_t _opennsl_gport_from_local(opennsl_port_t port) {
  opennsl_gport_t gport;
  OPENNSL_GPORT_LOCAL_SET(gport, port);
  return gport;
}

opennsl_port_t _opennsl_gport_to_local(opennsl_gport_t gport) {
  return OPENNSL_GPORT_LOCAL_GET(gport);
}

opennsl_module_t _opennsl_gport_modport_to_modid(opennsl_gport_t gport) {
  return OPENNSL_GPORT_MODPORT_MODID_GET(gport);
}

opennsl_port_t _opennsl_gport_modport_to_port(opennsl_gport_t port) {
  return OPENNSL_GPORT_MODPORT_PORT_GET(port);
}

opennsl_gport_t _opennsl_gport_from_modid_and_port(opennsl_port_t port, opennsl_module_t module) {
  opennsl_gport_t gport;
  OPENNSL_GPORT_MODPORT_SET(gport, module, port);
  return gport;
}

void _opennsl_rx_reason_clear_all(opennsl_rx_reasons_t* reasons) {
  OPENNSL_RX_REASON_CLEAR_ALL((*reasons));
}

int _opennsl_rx_reason_get(const opennsl_rx_reasons_t* reasons, opennsl_rx_reason_t reason) {
  return OPENNSL_RX_REASON_GET((*reasons), reason);
}

void _opennsl_rx_reason_set(opennsl_rx_reasons_t* reasons, opennsl_rx_reason_t reason) {
  OPENNSL_RX_REASON_SET((*reasons), reason);
}
