// -*- coding: utf-8 -*-

#include <stdlib.h>
#include <string.h>
#include <opennsl/error.h>
#include <opennsl/vlan.h>
#include "libopennsl.h"
#include "logger.h"

void opennsl_vxlan_vpn_config_t_init(opennsl_vxlan_vpn_config_t *info) {
  memset(info, 0, sizeof(opennsl_vxlan_vpn_config_t));
}

int opennsl_vxlan_vpn_create(int unit, opennsl_vxlan_vpn_config_t *info) {
  LOG_DEBUG("%s: unit   = %d", __func__, unit);
  _opennsl_vxlan_vpn_config_dump(__func__, info);

  return OPENNSL_E_NONE;
}

int opennsl_vxlan_vpn_destroy(int unit, opennsl_vpn_t l2vpn) {
  LOG_DEBUG("%s: unit   = %d", __func__, unit);
  LOG_DEBUG("%s: vpn    = %hu", __func__, l2vpn);

  return OPENNSL_E_NONE;
}

int opennsl_vxlan_vpn_destroy_all(int unit) {
  LOG_DEBUG("%s: unit   = %d", __func__, unit);

  return OPENNSL_E_NONE;
}

int opennsl_vxlan_vpn_get(int unit, opennsl_vpn_t l2vpn, opennsl_vxlan_vpn_config_t *info) {
  LOG_DEBUG("%s: unit   = %d", __func__, unit);
  LOG_DEBUG("%s: vpn    = %hu", __func__, l2vpn);
  LOG_DEBUG("%s: config = %p", __func__, info);

  return OPENNSL_E_NONE;
}

int opennsl_vxlan_init(int unit) {
  LOG_DEBUG("%s: unit   = %d", __func__, unit);

  return OPENNSL_E_NONE;
}

int opennsl_vxlan_cleanup(int unit) {
  LOG_DEBUG("%s: unit   = %d", __func__, unit);

  return OPENNSL_E_NONE;
}

int opennsl_vxlan_vpn_traverse(int unit, opennsl_vxlan_vpn_traverse_cb cb, void *user_data) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: callback = %p", __func__, cb);
  LOG_DEBUG("%s: data     = %p", __func__, user_data);

  return OPENNSL_E_NONE;
}

int opennsl_vxlan_port_add(int unit, opennsl_vpn_t l2vpn, opennsl_vxlan_port_t *vxlan_port) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: vpn    = %hu", __func__, l2vpn);
  _opennsl_vxlan_port_dump(__func__, vxlan_port);

  return OPENNSL_E_NONE;
}

int opennsl_vxlan_port_delete(int unit, opennsl_vpn_t l2vpn, opennsl_gport_t vxlan_port_id) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: vpn      = %hu", __func__, l2vpn);
  LOG_DEBUG("%s: port     = %d", __func__, vxlan_port_id);

  return OPENNSL_E_NONE;
}

int opennsl_vxlan_port_delete_all(int unit, opennsl_vpn_t l2vpn) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: vpn      = %hu", __func__, l2vpn);

  return OPENNSL_E_NONE;
}

int opennsl_vxlan_port_get(int unit, opennsl_vpn_t l2vpn, opennsl_vxlan_port_t *vxlan_port) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: vpn      = %hu", __func__, l2vpn);
  LOG_DEBUG("%s: port     = %p", __func__, vxlan_port);
  _opennsl_vxlan_port_dump(__func__, vxlan_port);

  return OPENNSL_E_NONE;
}

int opennsl_vxlan_port_get_all(int unit, opennsl_vpn_t l2vpn, int port_max, opennsl_vxlan_port_t *port_array, int *port_count) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: vpn      = %hu", __func__, l2vpn);
  LOG_DEBUG("%s: port_max = %d", __func__, port_max);
  LOG_DEBUG("%s: port_arr = %p", __func__, port_array);
  LOG_DEBUG("%s: port_cnt = %p", __func__, port_count);

  if (port_max < 0) {
    return OPENNSL_E_PARAM;
  }

  int index;
  for (index = 0; index < port_max; index++) {
    port_array[index].vxlan_port_id = index + 1;
  }
  *port_count = port_max;

  return OPENNSL_E_NONE;
}

void opennsl_vxlan_port_t_init(opennsl_vxlan_port_t *vxlan_port) {
  memset(vxlan_port, 0, sizeof(opennsl_vxlan_port_t));
}

int opennsl_vxlan_tunnel_initiator_create(int unit, opennsl_tunnel_initiator_t *info) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: info     = %p", __func__, info);

  return OPENNSL_E_NONE;
}

int opennsl_vxlan_tunnel_initiator_destroy(int unit, opennsl_gport_t vxlan_tunnel_id) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: tun_id   = %d", __func__, vxlan_tunnel_id);

  return OPENNSL_E_NONE;
}

int opennsl_vxlan_tunnel_initiator_get(int unit, opennsl_tunnel_initiator_t *info) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: info     = %p", __func__, info);

  return OPENNSL_E_NONE;
}

int opennsl_vxlan_tunnel_terminator_create(int unit, opennsl_tunnel_terminator_t *info) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: info     = %p", __func__, info);

  return OPENNSL_E_NONE;
}

int opennsl_vxlan_tunnel_terminator_destroy(int unit, opennsl_gport_t vxlan_tunnel_id) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: tun_id   = %d", __func__, vxlan_tunnel_id);

  return OPENNSL_E_NONE;
}

int opennsl_vxlan_tunnel_terminator_get(int unit, opennsl_tunnel_terminator_t *info) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: info     = %p", __func__, info);

  return OPENNSL_E_NONE;
}

int opennsl_vxlan_tunnel_terminator_update(int unit, opennsl_tunnel_terminator_t *info)  {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: info     = %p", __func__, info);

  return OPENNSL_E_NONE;
}

int opennsl_vxlan_tunnel_initiator_traverse (int unit, opennsl_tunnel_initiator_traverse_cb cb, void *user_data) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: callback = %p", __func__, cb);
  LOG_DEBUG("%s: data     = %p", __func__, user_data);

  return OPENNSL_E_NONE;
}

int opennsl_vxlan_tunnel_terminator_traverse (int unit, opennsl_tunnel_terminator_traverse_cb cb, void *user_data) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: callback = %p", __func__, cb);
  LOG_DEBUG("%s: data     = %p", __func__, user_data);

  return OPENNSL_E_NONE;
}
