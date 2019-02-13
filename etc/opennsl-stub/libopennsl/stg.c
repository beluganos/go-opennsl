// -*- coding: utf-8 -*-                                                                            

#include <stdlib.h>
#include <string.h>
#include <opennsl/error.h>
#include <opennsl/stg.h>
#include "libopennsl.h"
#include "logger.h"

int opennsl_stg_init(int unit) {
  LOG_DEBUG("%s: unit    = %d", __func__, unit);

  return OPENNSL_E_NONE;
}

int opennsl_stg_clear(int unit) {
  LOG_DEBUG("%s: unit    = %d", __func__, unit);

  return OPENNSL_E_NONE;
}

int opennsl_stg_default_set(int unit, opennsl_stg_t stg) {
  LOG_DEBUG("%s: init    = %d", __func__, unit);
  LOG_DEBUG("%s: stg     = %d", __func__, stg);

  return OPENNSL_E_NONE;
}

int opennsl_stg_default_get(int unit, opennsl_stg_t* stg) {
  LOG_DEBUG("%s: unit    = %d", __func__, unit);
  LOG_DEBUG("%s: stg     = %p", __func__, stg);

  *stg = 123;
  return OPENNSL_E_NONE;
}

int opennsl_stg_vlan_add(int unit, opennsl_stg_t stg, opennsl_vlan_t vid) {
  LOG_DEBUG("%s: unit    = %d", __func__, unit);
  LOG_DEBUG("%s: stg     = %d", __func__, stg);
  LOG_DEBUG("%s: vid    = %hu", __func__, vid);

  return OPENNSL_E_NONE;
}

int opennsl_stg_vlan_remove(int unit, opennsl_stg_t stg, opennsl_vlan_t vid) {
  LOG_DEBUG("%s: unit    = %d", __func__, unit);
  LOG_DEBUG("%s: stg     = %d", __func__, stg);
  LOG_DEBUG("%s: vid    = %hu", __func__, vid);

  return OPENNSL_E_NONE;
}

int opennsl_stg_vlan_remove_all(int unit, opennsl_stg_t stg) {
  LOG_DEBUG("%s: unit    = %d", __func__, unit);
  LOG_DEBUG("%s: stg     = %d", __func__, stg);

  return OPENNSL_E_NONE;
}

#define _STUB_STG_VLAN_NUM (5)

int opennsl_stg_vlan_list(int unit, opennsl_stg_t stg, opennsl_vlan_t **list, int *count) {
  LOG_DEBUG("%s: unit    = %d", __func__, unit);
  LOG_DEBUG("%s: stg     = %d", __func__, stg);
  LOG_DEBUG("%s: list    = %p", __func__, list);
  LOG_DEBUG("%s: count   = %p", __func__, count);

  int index;
  *list = (opennsl_vlan_t*)malloc(sizeof(opennsl_vlan_t) * _STUB_STG_VLAN_NUM);
  for (index = 0; index < _STUB_STG_VLAN_NUM; index++) {
    (*list)[index] = 1000 + stg * 10 + index;
  }
  *count = _STUB_STG_VLAN_NUM;

  return OPENNSL_E_NONE;
}

int opennsl_stg_vlan_list_destroy(int unit, opennsl_vlan_t *list, int count) {
  LOG_DEBUG("%s: unit    = %d", __func__, unit);
  LOG_DEBUG("%s: list    = %p", __func__, list);
  LOG_DEBUG("%s: count   = %d", __func__, count);
  int index;
  for (index = 0; index < count; index++) {
    LOG_DEBUG("%s: vlan    = %d", __func__, list[count]);
  }

  free(list);

  return OPENNSL_E_NONE;
}

int opennsl_stg_create(int unit, opennsl_stg_t *stg_ptr) {
  LOG_DEBUG("%s: unit    = %d", __func__, unit);
  LOG_DEBUG("%s: stg     = %p", __func__, stg_ptr);

  *stg_ptr = 101;
  return OPENNSL_E_NONE;
}

int opennsl_stg_destroy(int unit, opennsl_stg_t stg) {
  LOG_DEBUG("%s: unit    = %d", __func__, unit);
  LOG_DEBUG("%s: stg     = %d", __func__, stg);

  return OPENNSL_E_NONE;
}

#define _STUB_STG_NUM (4)

int opennsl_stg_list(int unit, opennsl_stg_t **list, int *count) {
  LOG_DEBUG("%s: unit    = %d", __func__, unit);
  LOG_DEBUG("%s: list    = %p", __func__, list);
  LOG_DEBUG("%s: count   = %p", __func__, count);

  int index;
  *list = (opennsl_stg_t*)malloc(sizeof(opennsl_stg_t) * _STUB_STG_NUM);
  for (index = 0; index < _STUB_STG_NUM; index++) {
    (*list)[index] = 100 + index;
  }

  *count = _STUB_STG_NUM;

  return OPENNSL_E_NONE;
}

int opennsl_stg_list_destroy(int unit, opennsl_stg_t *list, int count) {
  LOG_DEBUG("%s: unit    = %d", __func__, unit);
  LOG_DEBUG("%s: list    = %p", __func__, list);
  LOG_DEBUG("%s: count   = %d", __func__, count);

  int index;
  for (index = 0; index < count; index++) {
    LOG_DEBUG("%s: stg     = %d", __func__, list[count]);
  }

  free(list);

  return OPENNSL_E_NONE;
}

int opennsl_stg_stp_set(int unit, opennsl_stg_t stg, opennsl_port_t port, int stp_state) {
  LOG_DEBUG("%s: unit    = %d", __func__, unit);
  LOG_DEBUG("%s: stg     = %d", __func__, stg);
  LOG_DEBUG("%s: port    = %d", __func__, port);
  LOG_DEBUG("%s: state   = %d", __func__, stp_state);

  return OPENNSL_E_NONE;
}

int opennsl_stg_stp_get(int unit, opennsl_stg_t stg, opennsl_port_t port, int* stp_state) {
  LOG_DEBUG("%s: unit    = %d", __func__, unit);
  LOG_DEBUG("%s: stg     = %d", __func__, stg);
  LOG_DEBUG("%s: port    = %d", __func__, port);
  LOG_DEBUG("%s: state   = %p", __func__, stp_state);

  *stp_state = 1;
  return OPENNSL_E_NONE;
}

int opennsl_stg_count_get(int unit, int *max_stg) {
  LOG_DEBUG("%s: unit    = %d", __func__, unit);
  LOG_DEBUG("%s: max_stg = %p", __func__, max_stg);

  *max_stg = 10;
  return OPENNSL_E_NONE;
}
