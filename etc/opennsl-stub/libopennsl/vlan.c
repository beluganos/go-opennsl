// -*- coding: utf-8 -*-

#include <stdlib.h>
#include <string.h>
#include <opennsl/error.h>
#include <opennsl/vlan.h>
#include "libopennsl.h"
#include "logger.h"

typedef struct _opennsl_vlan_entry {
  opennsl_vlan_t vlan;
  opennsl_pbmp_t pbmp;
  opennsl_pbmp_t upbmp;
} _opennsl_vlan_entry_t;

static _opennsl_vlan_entry_t s_opennsl_vlan_entries[4096] = {0};
static opennsl_vlan_t s_default_vid = OPENNSL_VLAN_DEFAULT;

static int _opennsl_vlan_validate(opennsl_vlan_t vid) {
  return ((vid > 0) && (vid < 4096)) ? 0 : -1;
}
static _opennsl_vlan_entry_t* _opennsl_vlan_entry_get(opennsl_vlan_t vid) {
  if (_opennsl_vlan_validate(vid) < 0)
    return NULL;

  _opennsl_vlan_entry_t* const entry = s_opennsl_vlan_entries + vid;

  if (entry->vlan == 0)
    return NULL;
      
  return entry;
}

int opennsl_vlan_create(int unit, opennsl_vlan_t vid) {
  LOG_DEBUG("%s: unit = %d", __func__, unit);
  LOG_DEBUG("%s: vid  = %hu", __func__, vid);

  if (_opennsl_vlan_validate(vid) < 0)
    return OPENNSL_E_PARAM;

  _opennsl_vlan_entry_t* const entry = s_opennsl_vlan_entries + vid;

  if (entry->vlan != 0)
    return OPENNSL_E_EXISTS;

  entry->vlan = vid;
  OPENNSL_PBMP_CLEAR(entry->pbmp);
  OPENNSL_PBMP_CLEAR(entry->upbmp);
  
  return OPENNSL_E_NONE;
}

int opennsl_vlan_destroy(int unit, opennsl_vlan_t vid) {
  LOG_DEBUG("%s: unit = %d", __func__, unit);
  LOG_DEBUG("%s: vid  = %hu", __func__, vid);

  _opennsl_vlan_entry_t* entry = _opennsl_vlan_entry_get(vid);
  if (entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  entry->vlan = 0;
  return OPENNSL_E_NONE;
}

int opennsl_vlan_destroy_all(int unit) {
  LOG_DEBUG("%s: unit = %d", __func__, unit);

  int index;
  for (index = 0; index < 4096; index++) {
    _opennsl_vlan_entry_t* const entry = s_opennsl_vlan_entries + index;
    entry->vlan = 0;
  }

  return OPENNSL_E_NONE;
}

int opennsl_vlan_port_add(int unit, opennsl_vlan_t vid, opennsl_pbmp_t pbmp, opennsl_pbmp_t ubmp) {
  LOG_DEBUG("%s: unit = %d", __func__, unit);
  LOG_DEBUG("%s: vid  = %hu", __func__, vid);
  _opennsl_pbmp_dump(__func__, &pbmp);
  _opennsl_pbmp_dump(__func__, &ubmp);

  _opennsl_vlan_entry_t* entry = _opennsl_vlan_entry_get(vid);
  if (entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  OPENNSL_PBMP_OR(entry->pbmp, pbmp);
  OPENNSL_PBMP_OR(entry->upbmp, ubmp);

  return OPENNSL_E_NONE;
}

int opennsl_vlan_port_remove(int unit, opennsl_vlan_t vid, opennsl_pbmp_t pbmp) {
  LOG_DEBUG("%s: unit = %d", __func__, unit);
  _opennsl_pbmp_dump(__func__, &pbmp);

  _opennsl_vlan_entry_t* entry = _opennsl_vlan_entry_get(vid);
  if (entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  OPENNSL_PBMP_REMOVE(entry->pbmp, pbmp);
  OPENNSL_PBMP_REMOVE(entry->upbmp, pbmp);

  return OPENNSL_E_NONE;
}

int opennsl_vlan_port_get(int unit, opennsl_vlan_t vid, opennsl_pbmp_t *pbmp, opennsl_pbmp_t *ubmp) {
  LOG_DEBUG("%s: unit = %d", __func__, unit);
  LOG_DEBUG("%s: vid  = %hu", __func__, vid);

  _opennsl_vlan_entry_t* entry = _opennsl_vlan_entry_get(vid);
  if (entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  *pbmp = entry->pbmp;
  *ubmp = entry->upbmp;

  return OPENNSL_E_NONE;
}

int opennsl_vlan_gport_add(int unit, opennsl_vlan_t vlan, opennsl_gport_t port, int flags) {
  LOG_DEBUG("%s: unit = %d", __func__, unit);
  LOG_DEBUG("%s: vid  = %hu", __func__, vlan);
  LOG_DEBUG("%s: port = %d", __func__, port);
  LOG_DEBUG("%s: flag = %08x", __func__, flags);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_vlan_gport_delete(int unit, opennsl_vlan_t vlan, opennsl_gport_t port) {
  LOG_DEBUG("%s: unit = %d", __func__, unit);
  LOG_DEBUG("%s: vid  = %hu", __func__, vlan);
  LOG_DEBUG("%s: port = %d", __func__, port);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_vlan_gport_delete_all(int unit, opennsl_vlan_t vlan) {
  LOG_DEBUG("%s: unit = %d", __func__, unit);
  LOG_DEBUG("%s: vid  = %hu", __func__, vlan);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_vlan_gport_get(int unit, opennsl_vlan_t vlan, opennsl_gport_t port, int *flags) {
  LOG_DEBUG("%s: unit = %d", __func__, unit);
  LOG_DEBUG("%s: vid  = %hu", __func__, vlan);
  LOG_DEBUG("%s: port = %d", __func__, port);
  LOG_DEBUG("%s: flag = %p", __func__, flags);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_vlan_default_get(int unit, opennsl_vlan_t *vid_ptr) {
  LOG_DEBUG("%s: unit = %d", __func__, unit);
  LOG_DEBUG("%s: vid  = %p", __func__, vid_ptr);

  *vid_ptr = s_default_vid;
  return OPENNSL_E_NONE;

}

int opennsl_vlan_default_set(int unit, opennsl_vlan_t vid) {
  LOG_DEBUG("%s: unit = %d", __func__, unit);
  LOG_DEBUG("%s: vid  = %hu", __func__, vid);

  s_default_vid = vid;

  return OPENNSL_E_NONE;
}

int opennsl_vlan_list(int unit, opennsl_vlan_data_t **listp, int *countp) {
  LOG_DEBUG("%s: unit = %d", __func__, unit);
  LOG_DEBUG("%s: list = %p", __func__, listp);
  LOG_DEBUG("%s: count= %p", __func__, countp);

  opennsl_vlan_data_t* const list_p = (opennsl_vlan_data_t*)malloc(sizeof(opennsl_vlan_data_t) * 4096);
  int count = 0;
  int index;

  for (index = 0; index < 4096; index++) {
    _opennsl_vlan_entry_t* const entry = s_opennsl_vlan_entries + index;

    if (entry->vlan == 0)
      continue;

    list_p[count].vlan_tag = entry->vlan;
    list_p[count].port_bitmap = entry->pbmp;
    list_p[count].ut_port_bitmap = entry->upbmp;
    count++;
  }

  *listp = list_p;
  *countp = count;

  return OPENNSL_E_NONE;
}

int opennsl_vlan_list_destroy(int unit, opennsl_vlan_data_t *list, int count) {
  LOG_DEBUG("%s: unit = %d", __func__, unit);
  LOG_DEBUG("%s: list = %p", __func__, list);
  LOG_DEBUG("%s: count= %d", __func__, count);

  free(list);
  return OPENNSL_E_NONE;
}

int opennsl_vlan_control_set(int unit, opennsl_vlan_control_t type, int arg) {
  LOG_DEBUG("%s: unit = %d", __func__, unit);
  LOG_DEBUG("%s: type = %d", __func__, type);
  LOG_DEBUG("%s: arg  = %d", __func__, arg);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_vlan_control_port_set(int unit, int port, opennsl_vlan_control_port_t type, int arg) {
  LOG_DEBUG("%s: unit = %d", __func__, unit);
  LOG_DEBUG("%s: port = %d", __func__, port);
  LOG_DEBUG("%s: type = %d", __func__, type);
  LOG_DEBUG("%s: arg  = %d", __func__, arg);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_vlan_stat_get(int unit, opennsl_vlan_t vlan, opennsl_cos_t cos, opennsl_vlan_stat_t stat, uint64 *val) {
  LOG_DEBUG("%s: unit = %d", __func__, unit);
  LOG_DEBUG("%s: vid  = %hu", __func__, vlan);
  LOG_DEBUG("%s: cos  = %d", __func__, cos);
  LOG_DEBUG("%s: stat = %d", __func__, stat);
  LOG_DEBUG("%s: val  = %p", __func__, val);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_vlan_stat_set(int unit, opennsl_vlan_t vlan, opennsl_cos_t cos, opennsl_vlan_stat_t stat, uint64 val) {
  LOG_DEBUG("%s: unit = %d", __func__, unit);
  LOG_DEBUG("%s: vid  = %hu", __func__, vlan);
  LOG_DEBUG("%s: cos  = %d", __func__, cos);
  LOG_DEBUG("%s: stat = %d", __func__, stat);
  LOG_DEBUG("%s: val  = %llu", __func__, val);

  return OPENNSL_E_UNAVAIL;
}

void opennsl_vlan_port_t_init(opennsl_vlan_port_t *vlan_port) {
  memset(vlan_port, 0, sizeof(opennsl_vlan_port_t));
}

int opennsl_vlan_port_create(int unit, opennsl_vlan_port_t *vlan_port) {
  LOG_DEBUG("%s: unit = %d", __func__, unit);
  _opennsl_vlan_port_dump(__func__, vlan_port);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_vlan_port_destroy(int unit, opennsl_gport_t gport) {
  LOG_DEBUG("%s: unit = %d", __func__, unit);
  LOG_DEBUG("%s: port = %d", __func__, gport);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_vlan_port_find(int unit, opennsl_vlan_port_t *vlan_port) {
  LOG_DEBUG("%s: unit = %d", __func__, unit);
  _opennsl_vlan_port_dump(__func__, vlan_port);

  return OPENNSL_E_UNAVAIL;
}
