// -*- coding: utf-8 -*-

#include <string.h>
#include <opennsl/error.h>
#include <opennsl/mplsX.h>
#include "libopennsl.h"
#include "logger.h"

void opennsl_mpls_port_t_init(opennsl_mpls_port_t *mpls_port) {
  LOG_DEBUG("%s: port  = %p", __func__, mpls_port);
  memset(mpls_port, 0, sizeof(opennsl_mpls_port_t));
}

void opennsl_mpls_egress_label_t_init(opennsl_mpls_egress_label_t *label) {
  LOG_DEBUG("%s: info  = %p", __func__, label);
  memset(label, 0, sizeof(opennsl_mpls_egress_label_t));
}

void opennsl_mpls_tunnel_switch_t_init(opennsl_mpls_tunnel_switch_t *info) {
  LOG_DEBUG("%s: info  = %p", __func__, info);
  memset(info, 0, sizeof(opennsl_mpls_tunnel_switch_t));
}

void opennsl_mpls_entropy_identifier_t_init(opennsl_mpls_entropy_identifier_t *info) {
  LOG_DEBUG("%s: info  = %p", __func__, info);
  memset(info, 0, sizeof(opennsl_mpls_entropy_identifier_t));
}

void opennsl_mpls_exp_map_t_init(opennsl_mpls_exp_map_t *exp_map) {
  LOG_DEBUG("%s: info  = %p", __func__, exp_map);
  memset(exp_map, 0, sizeof(opennsl_mpls_exp_map_t));
}

int opennsl_mpls_init(int unit) {
  LOG_DEBUG("%s: unit   = %d", __func__, unit);

  return OPENNSL_E_NONE;
}

int opennsl_mpls_cleanup(int unit) {
  LOG_DEBUG("%s: unit   = %d", __func__, unit);

  return OPENNSL_E_NONE;
}

void opennsl_mpls_vpn_config_t_init(opennsl_mpls_vpn_config_t *info) {
  LOG_DEBUG("%s: info  = %p", __func__, info);
  memset(info, 0, sizeof(opennsl_mpls_vpn_config_t));
}

int opennsl_mpls_vpn_id_create(int unit, opennsl_mpls_vpn_config_t *info) {
  LOG_DEBUG("%s: unit   = %d", __func__, unit);
  _opennsl_mpls_vpn_config_dump(__func__, info);

  if ((info->flags & OPENNSL_MPLS_VPN_WITH_ID) == 0) {
    info->vpn = 10000;
  }

  return OPENNSL_E_NONE;
}

int opennsl_mpls_vpn_id_destroy(int unit, opennsl_vpn_t vpn) {
  LOG_DEBUG("%s: unit   = %d", __func__, unit);
  LOG_DEBUG("%s: vpn    = %d", __func__, vpn);

  return OPENNSL_E_NONE;
}

int opennsl_mpls_vpn_id_destroy_all(int unit) {
  LOG_DEBUG("%s: unit   = %d", __func__, unit);

  return OPENNSL_E_NONE;
}

int opennsl_mpls_vpn_id_get(int unit, opennsl_vpn_t vpn, opennsl_mpls_vpn_config_t *info) {
  LOG_DEBUG("%s: unit   = %d", __func__, unit);
  LOG_DEBUG("%s: vpn    = %d", __func__, vpn);
  LOG_DEBUG("%s: config = %p", __func__, info);

  return OPENNSL_E_NONE;
}

int opennsl_mpls_vpn_traverse(int unit, opennsl_mpls_vpn_traverse_cb cb, void *user_data) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: callback  = %p", __func__, cb);
  LOG_DEBUG("%s: user_data = %p", __func__, user_data);

  return OPENNSL_E_NONE;
}

int opennsl_mpls_tunnel_initiator_set(int unit, opennsl_if_t intf, int num_labels, opennsl_mpls_egress_label_t *label_array) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: iface     = %d", __func__, intf);
  LOG_DEBUG("%s: label_num = %d", __func__, num_labels);
  int index;
  for (index = 0; index < num_labels; index++) {
    _opennsl_mpls_egress_label_dump(__func__, label_array + index);
  }

  return OPENNSL_E_NONE;
}

int opennsl_mpls_tunnel_initiator_create(int unit, opennsl_if_t intf, int num_labels, opennsl_mpls_egress_label_t *label_array) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: iface     = %d", __func__, intf);
  LOG_DEBUG("%s: label_num = %d", __func__, num_labels);
  int index;
  for (index = 0; index < num_labels; index++) {
    _opennsl_mpls_egress_label_dump(__func__, label_array + index);
  }

  return OPENNSL_E_NONE;
}

int opennsl_mpls_tunnel_initiator_clear(int unit, opennsl_if_t intf) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: iface     = %d", __func__, intf);

  return OPENNSL_E_NONE;
}

int opennsl_mpls_tunnel_initiator_clear_all(int unit) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);

  return OPENNSL_E_NONE;
}

int opennsl_mpls_tunnel_initiator_get(int unit, opennsl_if_t intf, int label_max, opennsl_mpls_egress_label_t *label_array, int *label_count) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: iface     = %d", __func__, intf);
  LOG_DEBUG("%s: max       = %d", __func__, label_max);
  LOG_DEBUG("%s: arr       = %p", __func__, label_array);
  LOG_DEBUG("%s: count     = %p", __func__, label_count);

  int index;
  for (index = 0; index < label_max; index++) {
    label_array[index].l3_intf_id = index + 1;
  }

  *label_count = label_max;
  return OPENNSL_E_NONE;
}

int opennsl_mpls_tunnel_switch_add(int unit, opennsl_mpls_tunnel_switch_t *info) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  _opennsl_mpls_tunnel_switch_dump(__func__, info);
  return OPENNSL_E_NONE;
}

int opennsl_mpls_tunnel_switch_create(int unit, opennsl_mpls_tunnel_switch_t *info) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  _opennsl_mpls_tunnel_switch_dump(__func__, info);
  return OPENNSL_E_NONE;
}

int opennsl_mpls_tunnel_switch_delete(int unit, opennsl_mpls_tunnel_switch_t *info) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  _opennsl_mpls_tunnel_switch_dump(__func__, info);
  return OPENNSL_E_NONE;
}

int opennsl_mpls_tunnel_switch_delete_all(int unit) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  return OPENNSL_E_NONE;
}

int opennsl_mpls_tunnel_switch_get(int unit, opennsl_mpls_tunnel_switch_t *info) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  _opennsl_mpls_tunnel_switch_dump(__func__, info);
  return OPENNSL_E_NONE;
}

int opennsl_mpls_tunnel_switch_traverse(int unit, opennsl_mpls_tunnel_switch_traverse_cb cb, void *user_data) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: callback  = %p", __func__, cb);
  LOG_DEBUG("%s: data      = %p", __func__, user_data);

  return OPENNSL_E_NONE;
}
