// -*- coding: utf-8 -*-

#include <stdlib.h>
#include <string.h>
#include <opennsl/error.h>
#include <opennsl/trunk.h>
#include "libopennsl.h"
#include "logger.h"

int opennsl_trunk_init(int unit) {
  LOG_DEBUG("%s: unit  = %d", __func__, unit);
  return OPENNSL_E_NONE;
}

int opennsl_trunk_detach(int unit) {
  LOG_DEBUG("%s: unit  = %d", __func__, unit);
  return OPENNSL_E_NONE;
}

int opennsl_trunk_create(int unit, uint32 flags, opennsl_trunk_t *tid) {
  LOG_DEBUG("%s: unit  = %d", __func__, unit);
  LOG_DEBUG("%s: flags = %08x", __func__, flags);
  LOG_DEBUG("%s: trunk = %p", __func__, tid);

  *tid = 0x123;
  return OPENNSL_E_NONE;
}

int opennsl_trunk_destroy(int unit, opennsl_trunk_t tid) {
  LOG_DEBUG("%s: unit  = %d", __func__, unit);
  LOG_DEBUG("%s: trunk = %d", __func__, tid);

  return OPENNSL_E_NONE;
}

int opennsl_trunk_psc_set(int unit, opennsl_trunk_t tid, int psc) {
  LOG_DEBUG("%s: unit  = %d", __func__, unit);
  LOG_DEBUG("%s: tid   = %d", __func__, tid);
  LOG_DEBUG("%s: psc   = %d", __func__, psc);

   return OPENNSL_E_NONE;
}

void opennsl_trunk_info_t_init(opennsl_trunk_info_t *trunk_info) {
  memset(trunk_info, 0, sizeof(opennsl_trunk_info_t));
}

int opennsl_trunk_chip_info_get(int unit, opennsl_trunk_chip_info_t *info) {
  LOG_DEBUG("%s: unit  = %d", __func__, unit);
  LOG_DEBUG("%s: info  = %p", __func__, info);

   return OPENNSL_E_NONE;
}


int opennsl_trunk_failover_get(int unit, opennsl_trunk_t tid, opennsl_gport_t failport, int* psc, uint32* flags, int array_size, opennsl_gport_t* fail_to_array, int* array_count) {
  return OPENNSL_E_NONE;
}

int opennsl_trunk_failover_set(int unit, opennsl_trunk_t tid, opennsl_gport_t failport, int psc, uint32 flags, int array_size, opennsl_gport_t* fail_to_array) {
  return OPENNSL_E_NONE;
}

int opennsl_trunk_find(int unit, opennsl_module_t modid, opennsl_gport_t gport, opennsl_trunk_t *tid) {
  return OPENNSL_E_NONE;
}

void opennsl_trunk_member_t_init (opennsl_trunk_member_t *trunk_member) {
  memset(trunk_member, 0, sizeof(opennsl_trunk_member_t));
}

int opennsl_trunk_member_add(int unit, opennsl_trunk_t tid, opennsl_trunk_member_t *member) {
  return OPENNSL_E_NONE;
}

int opennsl_trunk_member_delete(int unit, opennsl_trunk_t tid, opennsl_trunk_member_t *member) {
  return OPENNSL_E_NONE;
}

int	opennsl_trunk_set (int unit, opennsl_trunk_t tid, opennsl_trunk_info_t *trunk_info, int member_count, opennsl_trunk_member_t *member_array) {
  return OPENNSL_E_NONE;
}
