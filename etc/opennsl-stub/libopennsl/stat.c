// -*- coding: utf-8 -*-

#include <string.h>
#include <opennsl/error.h>
#include <opennsl/stat.h>
#include "libopennsl.h"
#include "logger.h"

int opennsl_stat_init(int unit) {
  LOG_DEBUG("%s: unit   = %d", __func__, unit);

  return OPENNSL_E_NONE;
}

int opennsl_stat_clear(int unit, opennsl_port_t port) {
  LOG_DEBUG("%s: unit   = %d", __func__, unit);
  LOG_DEBUG("%s: port   = %d", __func__, port);

  return OPENNSL_E_NONE;
}

int opennsl_stat_sync(int unit) {
  LOG_DEBUG("%s: unit   = %d", __func__, unit);

  return OPENNSL_E_NONE;
}

int opennsl_stat_get(int unit, opennsl_port_t port, opennsl_stat_val_t type, uint64 *value) {
  LOG_DEBUG("%s: unit   = %d", __func__, unit);
  LOG_DEBUG("%s: port   = %d", __func__, port);
  LOG_DEBUG("%s: type   = %d", __func__, type);
  LOG_DEBUG("%s: value  = %p", __func__, value);

  *value = 8192;
  return OPENNSL_E_NONE;
}

int opennsl_stat_get32(int unit, opennsl_port_t port, opennsl_stat_val_t type, uint32 *value) {
  LOG_DEBUG("%s: unit   = %d", __func__, unit);
  LOG_DEBUG("%s: port   = %d", __func__, port);
  LOG_DEBUG("%s: type   = %d", __func__, type);
  LOG_DEBUG("%s: value  = %p", __func__, value);

  *value = 8192;
  return OPENNSL_E_NONE;
}

int opennsl_stat_multi_get(int unit, opennsl_port_t port, int nstat, opennsl_stat_val_t *stat_arr,  uint64 *value_arr) {
  LOG_DEBUG("%s: unit   = %d", __func__, unit);
  LOG_DEBUG("%s: port   = %d", __func__, port);
  LOG_DEBUG("%s: nstat  = %d", __func__, nstat);
  LOG_DEBUG("%s: stats  = %p", __func__, stat_arr);
  LOG_DEBUG("%s: values = %p", __func__, value_arr);

  int index;
  for (index = 0; index < nstat; index++) {
    value_arr[index] = 10000 + index;
    LOG_DEBUG("%s: stats[%d]  = %llu", __func__, stat_arr[index], value_arr[index]);
  }

  return OPENNSL_E_NONE;
}

int opennsl_stat_multi_get32(int unit, opennsl_port_t port, int nstat, opennsl_stat_val_t *stat_arr,  uint32 *value_arr) {
  LOG_DEBUG("%s: unit   = %d", __func__, unit);
  LOG_DEBUG("%s: port   = %d", __func__, port);
  LOG_DEBUG("%s: nstat  = %d", __func__, nstat);
  LOG_DEBUG("%s: stats  = %p", __func__, stat_arr);
  LOG_DEBUG("%s: values = %p", __func__, value_arr);

  int index;
  for (index = 0; index < nstat; index++) {
    value_arr[index] = 10000 + index;
    LOG_DEBUG("%s: stats[%d]  = %u", __func__, stat_arr[index], value_arr[index]);
  }

  return OPENNSL_E_NONE;
}

int opennsl_stat_sync_get32(int unit, opennsl_port_t port, opennsl_stat_val_t type, uint32 *value) {
  LOG_DEBUG("%s: unit   = %d", __func__, unit);
  LOG_DEBUG("%s: port   = %d", __func__, port);
  LOG_DEBUG("%s: type   = %d", __func__, type);
  LOG_DEBUG("%s: value  = %p", __func__, value);

  *value = 8192;
  return OPENNSL_E_NONE;
}

int opennsl_stat_sync_multi_get32(int unit, opennsl_port_t port, int nstat, opennsl_stat_val_t *stat_arr,  uint32 *value_arr) {
  LOG_DEBUG("%s: unit   = %d", __func__, unit);
  LOG_DEBUG("%s: port   = %d", __func__, port);
  LOG_DEBUG("%s: nstat  = %d", __func__, nstat);
  LOG_DEBUG("%s: stats  = %p", __func__, stat_arr);
  LOG_DEBUG("%s: values = %p", __func__, value_arr);

  int index;
  for (index = 0; index < nstat; index++) {
    value_arr[index] = 10000 + index;
    LOG_DEBUG("%s: stats[%d]  = %u", __func__, stat_arr[index], value_arr[index]);
  }

  return OPENNSL_E_NONE;
}

int opennsl_stat_clear_single(int unit, opennsl_port_t port, opennsl_stat_val_t type) {
  LOG_DEBUG("%s: unit   = %d", __func__, unit);
  LOG_DEBUG("%s: port   = %d", __func__, port);
  LOG_DEBUG("%s: type   = %d", __func__, type);

  return OPENNSL_E_NONE;
}

void opennsl_stat_value_t_init(opennsl_stat_value_t *stat_value) {
  memset(stat_value, 0, sizeof(opennsl_stat_value_t));
}
