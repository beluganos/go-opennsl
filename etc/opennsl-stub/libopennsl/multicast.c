// -*- coding: utf-8 -*-

#include <string.h>
#include <opennsl/error.h>
#include <opennsl/multicast.h>
#include <opennsl/multicastX.h>
#include "libopennsl.h"
#include "logger.h"

int opennsl_multicast_init(int unit) {
  LOG_DEBUG("%s: unit    = %d", __func__, unit);

  return OPENNSL_E_NONE;
}

int opennsl_multicast_create(int unit, uint32 flags, opennsl_multicast_t *group) {
  LOG_DEBUG("%s: unit    = %d", __func__, unit);
  LOG_DEBUG("%s: flags   = %08x", __func__, flags);
  LOG_DEBUG("%s: group   = %p", __func__, group);

  *group = 100;
  return OPENNSL_E_NONE;
}

int opennsl_multicast_destroy(int unit, opennsl_multicast_t group) {
  LOG_DEBUG("%s: unit    = %d", __func__, unit);
  LOG_DEBUG("%s: group   = %d", __func__, group);

  return OPENNSL_E_NONE;
}

int opennsl_multicast_l3_encap_get(int unit, opennsl_multicast_t group, opennsl_gport_t port, opennsl_if_t intf, opennsl_if_t *encap_id) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: group    = %d", __func__, group);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  LOG_DEBUG("%s: intf     = %d", __func__, intf);
  LOG_DEBUG("%s: encap_id = %p", __func__, encap_id);

  *encap_id = 100 + port;

  return OPENNSL_E_NONE;
}

int opennsl_multicast_l2_encap_get(int unit, opennsl_multicast_t group, opennsl_gport_t port, opennsl_vlan_t vlan, opennsl_if_t *encap_id) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: group    = %d", __func__, group);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  LOG_DEBUG("%s: vid      = %hu", __func__, vlan);
  LOG_DEBUG("%s: encap_id = %p", __func__, encap_id);

  *encap_id = 200 + port;

  return OPENNSL_E_NONE;
}

int opennsl_multicast_vlan_encap_get(int unit, opennsl_multicast_t group, opennsl_gport_t port, opennsl_gport_t vlan_port_id, opennsl_if_t *encap_id) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: group    = %d", __func__, group);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  LOG_DEBUG("%s: vlan_port= %hu", __func__, vlan_port_id);
  LOG_DEBUG("%s: encap_id = %p", __func__, encap_id);

  *encap_id = 300 + port;

  return OPENNSL_E_NONE;
}

int opennsl_multicast_egress_object_encap_get(int unit, opennsl_multicast_t group, opennsl_if_t intf, opennsl_if_t *encap_id) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: group    = %d", __func__, group);
  LOG_DEBUG("%s: intf     = %d", __func__, intf);
  LOG_DEBUG("%s: encap_id = %p", __func__, encap_id);

  *encap_id = 400 + intf;

  return OPENNSL_E_NONE;
}

int opennsl_multicast_egress_add(int unit, opennsl_multicast_t group, opennsl_gport_t port, opennsl_if_t encap_id) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: group    = %d", __func__, group);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  LOG_DEBUG("%s: encap_id = %d", __func__, encap_id);

  return OPENNSL_E_NONE;
}

int opennsl_multicast_egress_delete(int unit, opennsl_multicast_t group, opennsl_gport_t port, opennsl_if_t encap_id) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: group    = %d", __func__, group);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  LOG_DEBUG("%s: encap_id = %d", __func__, encap_id);

  return OPENNSL_E_NONE;
}

int opennsl_multicast_egress_delete_all(int unit, opennsl_multicast_t group) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: group    = %d", __func__, group);

  return OPENNSL_E_NONE;
}

int opennsl_multicast_egress_set(int unit, opennsl_multicast_t group, int port_count, opennsl_gport_t *port_array, opennsl_if_t *encap_id_array) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: group    = %d", __func__, group);
  LOG_DEBUG("%s: port_cnt = %d", __func__, port_count);

  int index;
  for (index = 0; index < port_count; index++) {
    LOG_DEBUG("%s: port[%d] = %d", __func__, index, port_array[index]);
  }

  return OPENNSL_E_NONE;
}

int opennsl_multicast_egress_get(int unit, opennsl_multicast_t group, int port_max, opennsl_gport_t *port_array, opennsl_if_t *encap_id_array, int *port_count) {
  LOG_DEBUG("%s: unit        = %d", __func__, unit);
  LOG_DEBUG("%s: group       = %d", __func__, group);
  LOG_DEBUG("%s: port_max    = %d", __func__, port_max);
  LOG_DEBUG("%s: port_array  = %p", __func__, port_array);
  LOG_DEBUG("%s: encap_array = %p", __func__, encap_id_array);
  LOG_DEBUG("%s: port_count  = %p", __func__, port_count);

  if (port_max < 0) {
    return OPENNSL_E_PARAM;
  }

  if (port_max == 0) {
    *port_count = 3;
    return OPENNSL_E_NONE;
  }

  int index;
  for (index = 0; index < port_max; index++) {
    port_array[index] = 100 + index;
    encap_id_array[index] = 200 + index;
  }
  *port_count = port_max;

  return OPENNSL_E_NONE;
}

int opennsl_multicast_ingress_add(int unit, opennsl_multicast_t group, opennsl_gport_t port, opennsl_if_t encap_id) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: group    = %d", __func__, group);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  LOG_DEBUG("%s: encap_id = %d", __func__, encap_id);

  return OPENNSL_E_NONE;
}

int opennsl_multicast_ingress_delete(int unit, opennsl_multicast_t group, opennsl_gport_t port, opennsl_if_t encap_id) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: group    = %d", __func__, group);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  LOG_DEBUG("%s: encap_id = %d", __func__, encap_id);

  return OPENNSL_E_NONE;
}

int opennsl_multicast_ingress_delete_all(int unit, opennsl_multicast_t group) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: group    = %d", __func__, group);

  return OPENNSL_E_NONE;
}

int opennsl_multicast_ingress_set(int unit, opennsl_multicast_t group, int port_count, opennsl_gport_t *port_array, opennsl_if_t *encap_id_array) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: group    = %d", __func__, group);
  LOG_DEBUG("%s: port_cnt = %d", __func__, port_count);

  int index;
  for (index = 0; index < port_count; index++) {
    LOG_DEBUG("%s: port[%d] = %d", __func__, index, port_array[index]);
  }

  return OPENNSL_E_NONE;
}

int opennsl_multicast_ingress_get(int unit, opennsl_multicast_t group, int port_max, opennsl_gport_t *port_array, opennsl_if_t *encap_id_array, int *port_count) {
  LOG_DEBUG("%s: unit        = %d", __func__, unit);
  LOG_DEBUG("%s: group       = %d", __func__, group);
  LOG_DEBUG("%s: port_max    = %d", __func__, port_max);
  LOG_DEBUG("%s: port_array  = %p", __func__, port_array);
  LOG_DEBUG("%s: encap_array = %p", __func__, encap_id_array);
  LOG_DEBUG("%s: port_count  = %p", __func__, port_count);

  if (port_max < 0) {
    return OPENNSL_E_PARAM;
  }

  if (port_max == 0) {
    *port_count = 3;
    return OPENNSL_E_NONE;
  }

  int index;
  for (index = 0; index < port_max; index++) {
    port_array[index] = 100 + index;
    encap_id_array[index] = 200 + index;
  }
  *port_count = port_max;

  return OPENNSL_E_NONE;
}

int opennsl_multicast_group_get(int unit, opennsl_multicast_t group, uint32 *flags) {
  LOG_DEBUG("%s: unit        = %d", __func__, unit);
  LOG_DEBUG("%s: group       = %d", __func__, group);
  LOG_DEBUG("%s: flags       = %p", __func__, flags);

  *flags = 0x11;
  return OPENNSL_E_NONE;
}

int opennsl_multicast_group_is_free(int unit, opennsl_multicast_t group) {
  LOG_DEBUG("%s: unit        = %d", __func__, unit);
  LOG_DEBUG("%s: group       = %d", __func__, group);

  if (group == 0)
    return OPENNSL_E_PARAM;

  if (group == 1)
    return OPENNSL_E_UNAVAIL;

  if (group == 2)
    return OPENNSL_E_EXISTS;

  return OPENNSL_E_NONE;
}

int opennsl_multicast_group_free_range_get(int unit, uint32 type_flag, opennsl_multicast_t *group_min, opennsl_multicast_t *group_max) {
  LOG_DEBUG("%s: unit        = %d", __func__, unit);
  LOG_DEBUG("%s: flags       = %08x", __func__, type_flag);
  LOG_DEBUG("%s: min/max     = %p/%p", __func__, group_min, group_max);

  *group_min = 10;
  *group_max = 20;
  return OPENNSL_E_NONE;
}

int opennsl_multicast_group_traverse(int unit, opennsl_multicast_group_traverse_cb_t trav_fn, uint32 flags, void *user_data) {
  LOG_DEBUG("%s: unit        = %d", __func__, unit);
  LOG_DEBUG("%s: callback    = %p", __func__, trav_fn);
  LOG_DEBUG("%s: flags       = %08x", __func__, flags);
  LOG_DEBUG("%s: data        = %p", __func__, user_data);

  return OPENNSL_E_NONE;
}

int opennsl_multicast_vxlan_encap_get(int unit, opennsl_multicast_t group, opennsl_gport_t port, opennsl_gport_t vxlan_port_id, opennsl_if_t *encap_id) {
  LOG_DEBUG("%s: unit        = %d", __func__, unit);
  LOG_DEBUG("%s: group       = %d", __func__, group);
  LOG_DEBUG("%s: port        = %d", __func__, port);
  LOG_DEBUG("%s: vxlan_port  = %d", __func__, vxlan_port_id);
  LOG_DEBUG("%s: encap_id    = %p", __func__, encap_id);

  *encap_id = 10;
  return OPENNSL_E_NONE;
}

int opennsl_multicast_control_set(int unit, opennsl_multicast_t group, opennsl_multicast_control_t type, int arg)  {
  LOG_DEBUG("%s: unit        = %d", __func__, unit);
  LOG_DEBUG("%s: group       = %d", __func__, group);
  LOG_DEBUG("%s: type        = %d", __func__, type);
  LOG_DEBUG("%s: arg         = %d", __func__, arg);

  return OPENNSL_E_NONE;
}

int opennsl_multicast_control_get(int unit, opennsl_multicast_t group, opennsl_multicast_control_t type, int* arg)  {
  LOG_DEBUG("%s: unit        = %d", __func__, unit);
  LOG_DEBUG("%s: group       = %d", __func__, group);
  LOG_DEBUG("%s: type        = %d", __func__, type);
  LOG_DEBUG("%s: arg         = %p", __func__, arg);

  *arg = 1;
  return OPENNSL_E_NONE;
}
