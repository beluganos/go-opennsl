// -*- coding: utf-8 -*-

#include <string.h>
#include <opennsl/error.h>
#include <opennsl/cosq.h>
#include <opennsl/cosqX.h>
#include "libopennsl.h"
#include "logger.h"

int opennsl_cosq_bst_profile_set(int unit, opennsl_gport_t gport, opennsl_cos_queue_t cosq, opennsl_bst_stat_id_t bid, opennsl_cosq_bst_profile_t *profile) {
  LOG_DEBUG("%s: unit         = %d", __func__, unit);
  LOG_DEBUG("%s: gport        = %d", __func__, gport);
  LOG_DEBUG("%s: cosq         = %d", __func__, cosq);
  LOG_DEBUG("%s: bst_stat_id  = %d", __func__, bid);
  LOG_DEBUG("%s: profile      = %p", __func__, profile);
  _opennsl_cosq_bst_profile_dump(__func__, profile);

  return OPENNSL_E_NONE;
}

int opennsl_cosq_bst_profile_get(int unit, opennsl_gport_t gport, opennsl_cos_queue_t cosq, opennsl_bst_stat_id_t bid, opennsl_cosq_bst_profile_t *profile) {
  LOG_DEBUG("%s: unit         = %d", __func__, unit);
  LOG_DEBUG("%s: gport        = %d", __func__, gport);
  LOG_DEBUG("%s: cosq         = %d", __func__, cosq);
  LOG_DEBUG("%s: bst_stat_id  = %d", __func__, bid);
  LOG_DEBUG("%s: profile      = %p", __func__, profile);

  profile->byte = 0x12345678;
  return OPENNSL_E_NONE;
}

int opennsl_cosq_bst_stat_sync(int unit, opennsl_bst_stat_id_t bid) {
  LOG_DEBUG("%s: unit         = %d", __func__, unit);
  LOG_DEBUG("%s: bst_stat_id  = %d", __func__, bid);

  return OPENNSL_E_NONE;
}

int opennsl_cosq_bst_stat_clear(int unit, opennsl_gport_t gport, opennsl_cos_queue_t cosq, opennsl_bst_stat_id_t bid) {
  LOG_DEBUG("%s: unit         = %d", __func__, unit);
  LOG_DEBUG("%s: gport        = %d", __func__, gport);
  LOG_DEBUG("%s: cosq         = %d", __func__, cosq);
  LOG_DEBUG("%s: bst_stat_id  = %d", __func__, bid);

  return OPENNSL_E_NONE;
}

int opennsl_cosq_bst_stat_get(int unit, opennsl_gport_t gport, opennsl_cos_queue_t cosq, opennsl_bst_stat_id_t bid, uint32 options, uint64 *value) {
  LOG_DEBUG("%s: unit         = %d", __func__, unit);
  LOG_DEBUG("%s: gport        = %d", __func__, gport);
  LOG_DEBUG("%s: cosq         = %d", __func__, cosq);
  LOG_DEBUG("%s: bst_stat_id  = %d", __func__, bid);
  LOG_DEBUG("%s: options      = %u(%08x)", __func__, options, options);
  LOG_DEBUG("%s: value         = %p", __func__, value);

  *value = 0x12345678;
  return OPENNSL_E_NONE;
}

int opennsl_cosq_bst_stat_multi_get(int unit, opennsl_gport_t gport, opennsl_cos_queue_t cosq, uint32 options, int max_values, opennsl_bst_stat_id_t *id_list, uint64 *values) {
  LOG_DEBUG("%s: unit         = %d", __func__, unit);
  LOG_DEBUG("%s: gport        = %d", __func__, gport);
  LOG_DEBUG("%s: cosq         = %d", __func__, cosq);
  LOG_DEBUG("%s: options      = %u(%08x)", __func__, options, options);
  LOG_DEBUG("%s: max_values   = %d", __func__, max_values);
  LOG_DEBUG("%s: id_list      = %p", __func__, id_list);
  LOG_DEBUG("%s: values       = %p", __func__, values);

  if (max_values < 0) {
    return OPENNSL_E_PARAM;
  }

  int index;
  for (index = 0; index < max_values; index++) {
    values[index] = id_list[index];
  }

  return OPENNSL_E_NONE;
}

int opennsl_cosq_init(int unit) {
  LOG_DEBUG("%s: unit         = %d", __func__, unit);

  return OPENNSL_E_NONE;
}

int opennsl_cosq_detach(int unit) {
  LOG_DEBUG("%s: unit         = %d", __func__, unit);

  return OPENNSL_E_NONE;
}

int opennsl_cosq_config_set(int unit, int numq) {
  LOG_DEBUG("%s: unit         = %d", __func__, unit);
  LOG_DEBUG("%s: numq         = %d", __func__, numq);

  return OPENNSL_E_NONE;
}

int opennsl_cosq_config_get(int unit, int *numq) {
  LOG_DEBUG("%s: unit         = %d", __func__, unit);
  LOG_DEBUG("%s: numq         = %p", __func__, numq);

  *numq = 5;
  return OPENNSL_E_NONE;
}

int opennsl_cosq_mapping_set(int unit, opennsl_cos_t priority, opennsl_cos_queue_t cosq) {
  LOG_DEBUG("%s: unit         = %d", __func__, unit);
  LOG_DEBUG("%s: priority     = %d", __func__, priority);
  LOG_DEBUG("%s: cosq         = %d", __func__, cosq);

  return OPENNSL_E_NONE;
}

int opennsl_cosq_mapping_get(int unit, opennsl_cos_t priority, opennsl_cos_queue_t* cosq) {
  LOG_DEBUG("%s: unit         = %d", __func__, unit);
  LOG_DEBUG("%s: priority     = %d", __func__, priority);
  LOG_DEBUG("%s: cosq         = %p", __func__, cosq);

  *cosq = priority + 1;
  return OPENNSL_E_NONE;
}

int opennsl_cosq_port_mapping_set(int unit, opennsl_port_t port, opennsl_cos_t priority, opennsl_cos_queue_t cosq) {
  LOG_DEBUG("%s: unit         = %d", __func__, unit);
  LOG_DEBUG("%s: port         = %d", __func__, port);
  LOG_DEBUG("%s: priority     = %d", __func__, priority);
  LOG_DEBUG("%s: cosq         = %d", __func__, cosq);

  return OPENNSL_E_NONE;
}

int opennsl_cosq_port_mapping_get(int unit, opennsl_port_t port, opennsl_cos_t priority, opennsl_cos_queue_t* cosq) {
  LOG_DEBUG("%s: unit         = %d", __func__, unit);
  LOG_DEBUG("%s: port         = %d", __func__, port);
  LOG_DEBUG("%s: priority     = %d", __func__, priority);
  LOG_DEBUG("%s: cosq         = %p", __func__, cosq);

  *cosq = port * 10 + priority;
  return OPENNSL_E_NONE;
}
