// -*- coding: utf-8 -*-

#include <string.h>
#include <opennsl/error.h>
#include <opennsl/switch.h>
#include "libopennsl.h"
#include "logger.h"

int opennsl_switch_control_get(int unit, opennsl_switch_control_t type, int *arg) {
  LOG_DEBUG("%s: unit = %d", __func__, unit);
  LOG_DEBUG("%s: type = %d", __func__, type);
  LOG_DEBUG("%s: arg  = %p", __func__, arg);
  *arg = 1;
  return OPENNSL_E_NONE;
}

int opennsl_switch_control_set(int unit, opennsl_switch_control_t type, int arg) {
  LOG_DEBUG("%s: unit = %d", __func__, unit);
  LOG_DEBUG("%s: type = %d", __func__, type);
  LOG_DEBUG("%s: arg  = %d", __func__, arg);
  return OPENNSL_E_NONE;
}

int opennsl_switch_control_port_get(int unit, opennsl_port_t port, opennsl_switch_control_t type, int *arg) {
  LOG_DEBUG("%s: unit = %d", __func__, unit);
  LOG_DEBUG("%s: port = %d", __func__, port);
  LOG_DEBUG("%s: type = %d", __func__, type);
  LOG_DEBUG("%s: arg  = %p", __func__, arg);
  *arg = 1;
  return OPENNSL_E_NONE;
}

int opennsl_switch_control_port_set(int unit, opennsl_port_t port, opennsl_switch_control_t type, int arg) {
  LOG_DEBUG("%s: unit = %d", __func__, unit);
  LOG_DEBUG("%s: port = %d", __func__, port);
  LOG_DEBUG("%s: type = %d", __func__, type);
  LOG_DEBUG("%s: arg  = %d", __func__, arg);
  return OPENNSL_E_NONE;
}

int opennsl_switch_temperature_monitor_get(int unit, int temperature_max, opennsl_switch_temperature_monitor_t *temperature_array, int *temperature_count) {
  LOG_DEBUG("%s: unit = %d", __func__, unit);
  LOG_DEBUG("%s: max  = %d", __func__, temperature_max);
  LOG_DEBUG("%s: array= %p", __func__, temperature_array);
  LOG_DEBUG("%s: count= %p", __func__, temperature_count);

  if (temperature_max < 1) {
    return OPENNSL_E_PARAM;
  }

  int i;
  for (i = 0; i < temperature_max - 1; i++) {
    temperature_array[i].curr = 10 + i;
    temperature_array[i].peak = 100 + i;
  }
  *temperature_count = temperature_max - 1;

  return OPENNSL_E_NONE;
}

void opennsl_switch_pkt_info_t_init(opennsl_switch_pkt_info_t *pkt_info) {
  LOG_DEBUG("%s: info = %p", __func__, pkt_info);
  memset(pkt_info, 0, sizeof(opennsl_switch_pkt_info_t));
}

int opennsl_switch_pkt_info_hash_get(int unit, opennsl_switch_pkt_info_t *pkt_info, opennsl_gport_t *dst_gport, opennsl_if_t *dst_intf) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  _opennsl_switch_pkt_info_dump(__func__, pkt_info);
  LOG_DEBUG("%s: dst_gport = %p", __func__, dst_gport);
  LOG_DEBUG("%s: dst_iface = %p", __func__, dst_intf);
  return OPENNSL_E_NONE;
}

int opennsl_switch_event_register(int unit, opennsl_switch_event_cb_t cb, void *userdata) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: callback  = %p", __func__, cb);
  LOG_DEBUG("%s: userdata  = %p", __func__, userdata);
  return OPENNSL_E_NONE;
}

int opennsl_switch_event_unregister(int unit, opennsl_switch_event_cb_t cb, void *userdata) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: callback  = %p", __func__, cb);
  LOG_DEBUG("%s: userdata  = %p", __func__, userdata);
  return OPENNSL_E_NONE;
}

int opennsl_switch_object_count_multi_get(int unit, int object_size, opennsl_switch_object_t *object_array, int *entries) {
  LOG_DEBUG("%s: unit             = %d", __func__, unit);
  LOG_DEBUG("%s: object_size      = %d", __func__, object_size);
  int index;
  for (index = 0; index < object_size; index++) {
    LOG_DEBUG("%s: object_array[%d] = %d", __func__, index, object_array[index]);
  }
  LOG_DEBUG("%s: entries         = %p", __func__, entries);

  if (object_size < 1) {
    return OPENNSL_E_PARAM;
  }

  *entries = object_size - 1;
  return OPENNSL_E_NONE;
}
