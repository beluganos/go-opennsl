// -*- coding: utf-8 -*-

#include <string.h>
#include <opennsl/error.h>
#include <opennsl/fieldX.h>
#include "libopennsl.h"
#include "logger.h"

typedef struct _opennsl_field_group_entry {
  opennsl_field_group_t group_id;
  opennsl_field_qset_t qset;
  opennsl_field_aset_t aset;
  opennsl_field_group_status_t status;

  int enable;
  int priority;
  int installed;
} _opennsl_field_group_entry_t;

#define _OPENNSL_FIELD_GROUP_MAX (16)
static _opennsl_field_group_entry_t s_field_group_entry_list[_OPENNSL_FIELD_GROUP_MAX] = {0};

_opennsl_field_group_entry_t* _opennsl_field_group_entry_get(opennsl_field_group_t group) {
  int index;
  for (index=0; index < _OPENNSL_FIELD_GROUP_MAX; index++) {
    _opennsl_field_group_entry_t* const entry = s_field_group_entry_list + index;
    if (entry->group_id == 0)
      continue;

    if (entry->group_id == group)
      return entry;
  }

  return NULL;
}

typedef struct _opennsl_field_q_entry {
  opennsl_port_t in_port;
  opennsl_port_t in_port_mask;
  opennsl_port_t out_port;
  opennsl_port_t out_port_mask;

  opennsl_l4_port_t l4_sport;
  opennsl_l4_port_t l4_sport_mask;
  opennsl_l4_port_t l4_dport;
  opennsl_l4_port_t l4_dport_mask;

  uint16 eth_type;
  uint16 eth_type_mask;

  uint8 ip_proto;
  uint8 ip_proto_mask;

  opennsl_ip_t ip_dst;
  opennsl_ip_t ip_dst_mask;
  opennsl_ip6_t ip6_dst;
  opennsl_ip6_t ip6_dst_mask;
  char mac_dst[64];
  char mac_dst_mask[64];

  uint32 vrf;
  uint32 vrf_mask;

  uint8 cpu_q;
  uint8 cpu_q_mask;
} _opennsl_field_q_entry_t;

typedef struct _opennsl_field_entry_entry {
  opennsl_field_entry_t entry_id;
  opennsl_field_group_t group_id;
  _opennsl_field_q_entry_t qualifier;

  int priority;
  int installed;
} _opennsl_field_entry_entry_t;

#define _OPENNSL_FIELD_ENTRY_MAX (32 *_OPENNSL_FIELD_GROUP_MAX)
static _opennsl_field_entry_entry_t s_field_entry_list[_OPENNSL_FIELD_ENTRY_MAX] = {0};

_opennsl_field_entry_entry_t* _opennsl_field_entry_entry_get(opennsl_field_entry_t entry_id) {
  int index;
  for (index = 0; index < _OPENNSL_FIELD_ENTRY_MAX; index++) {
    _opennsl_field_entry_entry_t* const entry = s_field_entry_list + index;
    if (entry->entry_id == 0)
      continue;

    if (entry->entry_id == entry_id)
      return entry;
  }

  return NULL;
}

int opennsl_field_init(int unit) {
  LOG_DEBUG("%s: unit= %d", __func__, unit);
  return OPENNSL_E_NONE;
}

int opennsl_field_detach(int unit) {
  LOG_DEBUG("%s: unit= %d", __func__, unit);
  return OPENNSL_E_NONE;
}

int opennsl_field_group_create(int unit, opennsl_field_qset_t qset, int pri, opennsl_field_group_t *group) {
  LOG_DEBUG("%s: unit  = %d", __func__, unit);
  LOG_DEBUG("%s: pri   = %d", __func__, pri);
  LOG_DEBUG("%s: group = %p", __func__, group);
  _opennsl_field_qset_dump(__func__, &qset);

  int index;
  for (index=0; index < _OPENNSL_FIELD_GROUP_MAX; index++) {
    _opennsl_field_group_entry_t* const entry = s_field_group_entry_list + index;
    if (entry->group_id != 0)
      continue;

    if (entry->group_id == index +1)
      return  OPENNSL_E_BADID;

    memset(entry, 0, sizeof(_opennsl_field_group_entry_t));
    entry->group_id = index + 1;
    entry->qset = qset;
    entry->enable = 1;

    *group = index + 1;
    return OPENNSL_E_NONE;
  }

  return OPENNSL_E_RESOURCE;
}

int opennsl_field_group_create_id(int unit, opennsl_field_qset_t qset, int pri, opennsl_field_group_t group) {
  LOG_DEBUG("%s: unit  = %d", __func__, unit);
  LOG_DEBUG("%s: pri   = %d", __func__, pri);
  LOG_DEBUG("%s: group = %d", __func__, group);
  _opennsl_field_qset_dump(__func__, &qset);

  int index;
  for (index=0; index < _OPENNSL_FIELD_GROUP_MAX; index++) {
    _opennsl_field_group_entry_t* const entry = s_field_group_entry_list + index;
    if (entry->group_id != 0)
      continue;

    if (entry->group_id == group)
      return  OPENNSL_E_BADID;

    memset(entry, 0, sizeof(_opennsl_field_group_entry_t));
    entry->group_id = group;
    entry->qset = qset;
    entry->enable = 1;

    return OPENNSL_E_NONE;
  }

  return OPENNSL_E_RESOURCE;
}

int opennsl_field_group_create_mode(int unit, opennsl_field_qset_t qset, int pri, opennsl_field_group_mode_t mode, opennsl_field_group_t *group) {
  LOG_DEBUG("%s: unit  = %d", __func__, unit);
  LOG_DEBUG("%s: pri   = %d", __func__, pri);
  LOG_DEBUG("%s: mode  = %d", __func__, mode);
  LOG_DEBUG("%s: group = %p", __func__, group);
  _opennsl_field_qset_dump(__func__, &qset);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_group_create_mode_id(int unit, opennsl_field_qset_t qset, int pri, opennsl_field_group_mode_t mode, opennsl_field_group_t group) {
  LOG_DEBUG("%s: unit  = %d", __func__, unit);
  LOG_DEBUG("%s: pri   = %d", __func__, pri);
  LOG_DEBUG("%s: mode  = %d", __func__, mode);
  LOG_DEBUG("%s: group = %d", __func__, group);
  _opennsl_field_qset_dump(__func__, &qset);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_group_traverse(int unit, opennsl_field_group_traverse_cb callback, void *user_data) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: callback  = %p", __func__, callback);
  LOG_DEBUG("%s: user_data = %p", __func__, user_data);

  int index;
  for (index=0; index < _OPENNSL_FIELD_GROUP_MAX; index++) {
    _opennsl_field_group_entry_t* const entry = s_field_group_entry_list + index;
    if (entry->group_id == 0)
      continue;

    const int rc = callback(unit, entry->group_id, user_data);
    if (rc != OPENNSL_E_NONE)
      return rc;
  }

  return OPENNSL_E_NONE;
}

int opennsl_field_group_set(int unit, opennsl_field_group_t group, opennsl_field_qset_t qset) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: group     = %d", __func__, group);
  _opennsl_field_qset_dump(__func__, &qset);

  _opennsl_field_group_entry_t* const entry = _opennsl_field_group_entry_get(group);
  if (entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  entry->qset = qset;
  return OPENNSL_E_NONE;
}

int opennsl_field_group_get(int unit, opennsl_field_group_t group, opennsl_field_qset_t *qset) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: group     = %d", __func__, group);
  LOG_DEBUG("%s: qset      = %p", __func__, qset);

  _opennsl_field_group_entry_t* const entry = _opennsl_field_group_entry_get(group);
  if (entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  *qset = entry->qset;
  return OPENNSL_E_NONE;
}

int opennsl_field_group_action_set(int unit, opennsl_field_group_t group, opennsl_field_aset_t aset) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: group     = %d", __func__, group);
  _opennsl_field_aset_dump(__func__, &aset);

  _opennsl_field_group_entry_t* const entry = _opennsl_field_group_entry_get(group);
  if (entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  entry->aset = aset;
  return OPENNSL_E_NONE;
}

int opennsl_field_group_action_get(int unit, opennsl_field_group_t group, opennsl_field_aset_t *aset) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: group     = %d", __func__, group);
  LOG_DEBUG("%s: aset      = %p", __func__, aset);

  _opennsl_field_group_entry_t* const entry = _opennsl_field_group_entry_get(group);
  if (entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  *aset = entry->aset;
  return OPENNSL_E_NONE;
}

int opennsl_field_group_destroy(int unit, opennsl_field_group_t group) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: group     = %d", __func__, group);

  _opennsl_field_group_entry_t* const entry = _opennsl_field_group_entry_get(group);
  if (entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  entry->group_id = 0;
  return OPENNSL_E_NONE;
}

int opennsl_field_group_priority_set(int unit, opennsl_field_group_t group, int priority) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: group     = %d", __func__, group);
  LOG_DEBUG("%s: priority  = %d", __func__, priority);

  _opennsl_field_group_entry_t* const entry = _opennsl_field_group_entry_get(group);
  if (entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  entry->priority = priority;
  return OPENNSL_E_NONE;
}

int opennsl_field_group_priority_get(int unit, opennsl_field_group_t group, int* priority) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: group     = %d", __func__, group);
  LOG_DEBUG("%s: priority  = %p", __func__, priority);

  _opennsl_field_group_entry_t* const entry = _opennsl_field_group_entry_get(group);
  if (entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  *priority = entry->priority;
  return OPENNSL_E_NONE;
}

int opennsl_field_group_status_get(int unit, opennsl_field_group_t group, opennsl_field_group_status_t *status) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: group     = %d", __func__, group);
  LOG_DEBUG("%s: status    = %p", __func__, status);

  _opennsl_field_group_entry_t* const entry = _opennsl_field_group_entry_get(group);
  if (entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  *status = entry->status;
  return OPENNSL_E_NONE;
}

int opennsl_field_group_enable_set(int unit, opennsl_field_group_t group, int enable) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: group     = %d", __func__, group);
  LOG_DEBUG("%s: enable    = %d", __func__, enable);

  _opennsl_field_group_entry_t* const entry = _opennsl_field_group_entry_get(group);
  if (entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  entry->enable = enable;
  return OPENNSL_E_NONE;
}

int opennsl_field_group_enable_get(int unit, opennsl_field_group_t group, int* enable) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: group     = %d", __func__, group);
  LOG_DEBUG("%s: enable    = %p", __func__, enable);

  _opennsl_field_group_entry_t* const entry = _opennsl_field_group_entry_get(group);
  if (entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  *enable = entry->enable;
  return OPENNSL_E_NONE;
}

int opennsl_field_group_install(int unit, opennsl_field_group_t group) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: group     = %d", __func__, group);

  _opennsl_field_group_entry_t* const entry = _opennsl_field_group_entry_get(group);
  if (entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  if (entry->installed != 0)
    return OPENNSL_E_BUSY;

  entry->installed = 1;
  return OPENNSL_E_NONE;
}

int opennsl_field_range_create(int unit, opennsl_field_range_t *range, uint32 flags, opennsl_l4_port_t min, opennsl_l4_port_t max) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: range     = %p", __func__, range);
  LOG_DEBUG("%s: flags     = %d", __func__, flags);
  LOG_DEBUG("%s: min/max   = %d/%d", __func__, min, max);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_range_get(int unit, opennsl_field_range_t range, uint32 *flags, opennsl_l4_port_t *min, opennsl_l4_port_t *max) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: range     = %d", __func__, range);
  LOG_DEBUG("%s: flags     = %p", __func__, flags);
  LOG_DEBUG("%s: min/max   = %p/%p", __func__, min, max);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_range_destroy(int unit, opennsl_field_range_t range) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: range     = %d", __func__, range);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_entry_create(int unit, opennsl_field_group_t group_id, opennsl_field_entry_t *entry_id) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: group     = %d", __func__, group_id);
  LOG_DEBUG("%s: entry     = %p", __func__, entry_id);

  _opennsl_field_group_entry_t* const group = _opennsl_field_group_entry_get(group_id);
  if (group == NULL)
    return OPENNSL_E_NOT_FOUND;

  int index;
  for (index = 0; index < _OPENNSL_FIELD_ENTRY_MAX; index++) {
    _opennsl_field_entry_entry_t* const entry = s_field_entry_list + index;
    if (entry->entry_id != 0)
      continue;

    if (entry->entry_id == index + 1)
      return  OPENNSL_E_BADID;

    memset(entry, 0, sizeof(_opennsl_field_entry_entry_t));
    entry->entry_id = index + 1;
    entry->group_id = group_id;

    *entry_id = index + 1;
    return OPENNSL_E_NONE;
  }

  return OPENNSL_E_NONE;
}

int opennsl_field_entry_create_id(int unit, opennsl_field_group_t group_id, opennsl_field_entry_t entry_id) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: group     = %d", __func__, group_id);
  LOG_DEBUG("%s: entry     = %d", __func__, entry_id);

  _opennsl_field_group_entry_t* const group = _opennsl_field_group_entry_get(group_id);
  if (group == NULL)
    return OPENNSL_E_NOT_FOUND;

  int index;
  for (index = 0; index < _OPENNSL_FIELD_ENTRY_MAX; index++) {
    _opennsl_field_entry_entry_t* const entry = s_field_entry_list + index;
    if (entry->entry_id != 0)
      continue;

    if (entry->entry_id == group_id)
      return  OPENNSL_E_BADID;

    memset(entry, 0, sizeof(_opennsl_field_entry_entry_t));
    entry->group_id = group_id;
    entry->group_id = group_id;

    return OPENNSL_E_NONE;
  }

  return OPENNSL_E_NONE;
}

int opennsl_field_entry_multi_get(int unit, opennsl_field_group_t group_id, int entry_size, opennsl_field_entry_t *entry_array, int *entry_count) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: group     = %d", __func__, group_id);
  LOG_DEBUG("%s: size      = %d", __func__, entry_size);
  LOG_DEBUG("%s: entry     = %p", __func__, entry_array);
  LOG_DEBUG("%s: count     = %p", __func__, entry_count);

  int count = 0;
  int index;
  for (index = 0; index < _OPENNSL_FIELD_ENTRY_MAX; index++) {
    _opennsl_field_entry_entry_t* const entry = s_field_entry_list + index;
    if ((entry->group_id == 0) || (entry->group_id != group_id))
      continue;

    if (entry_array != NULL) {
      if (count >= entry_size)
	break;

      entry_array[count] = entry->entry_id;
    }

    count++;
  }

  *entry_count = count;

  LOG_DEBUG("%s: count     = %d", __func__, *entry_count);

  return OPENNSL_E_NONE;
}

int opennsl_field_entry_destroy(int unit, opennsl_field_entry_t entry) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);

  _opennsl_field_entry_entry_t* const _entry = _opennsl_field_entry_entry_get(entry);
  if (_entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  _entry->entry_id = 0;
  return OPENNSL_E_NONE;
}

int opennsl_field_entry_destroy_all(int unit) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);

  int index;
  for (index = 0; index < _OPENNSL_FIELD_ENTRY_MAX; index++)
    s_field_entry_list[index].entry_id = 0;

  return OPENNSL_E_NONE;
}

int opennsl_field_entry_copy(int unit, opennsl_field_entry_t src_entry, opennsl_field_entry_t *dst_entry) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: src_entry = %d", __func__, src_entry);
  LOG_DEBUG("%s: dst_entry = %p", __func__, dst_entry);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_entry_install(int unit, opennsl_field_entry_t entry) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);

  _opennsl_field_entry_entry_t* const _entry = _opennsl_field_entry_entry_get(entry);
  if (_entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  if (_entry->installed != 0)
    return OPENNSL_E_BUSY;

  _entry->installed = 1;
  return OPENNSL_E_NONE;
}

int opennsl_field_entry_reinstall(int unit, opennsl_field_entry_t entry) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);

  _opennsl_field_entry_entry_t* const _entry = _opennsl_field_entry_entry_get(entry);
  if (_entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  if (_entry->installed == 0)
    return OPENNSL_E_NOT_FOUND;

  return OPENNSL_E_NONE;
}

int opennsl_field_entry_remove(int unit, opennsl_field_entry_t entry) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);

  _opennsl_field_entry_entry_t* const _entry = _opennsl_field_entry_entry_get(entry);
  if (_entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  _entry->entry_id = 0;

  return OPENNSL_E_NONE;
}

int opennsl_field_entry_policer_attach(int unit, opennsl_field_entry_t entry_id, int level, opennsl_policer_t policer_id) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry_id);
  LOG_DEBUG("%s: level     = %d", __func__, level);
  LOG_DEBUG("%s: policer   = %d", __func__, policer_id);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_entry_policer_detach(int unit, opennsl_field_entry_t entry_id, int level) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry_id);
  LOG_DEBUG("%s: level     = %d", __func__, level);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_entry_policer_get(int unit, opennsl_field_entry_t entry_id, int level, opennsl_policer_t* policer) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry_id);
  LOG_DEBUG("%s: level     = %d", __func__, level);
  LOG_DEBUG("%s: policer   = %p", __func__, policer);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_entry_prio_get(int unit, opennsl_field_entry_t entry, int *prio) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: prio      = %p", __func__, prio);

  _opennsl_field_entry_entry_t* const _entry = _opennsl_field_entry_entry_get(entry);
  if (_entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  *prio = _entry->priority;
  return OPENNSL_E_NONE;
}

int opennsl_field_entry_prio_set(int unit, opennsl_field_entry_t entry, int prio) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: prio      = %d", __func__, prio);

  _opennsl_field_entry_entry_t* const _entry = _opennsl_field_entry_entry_get(entry);
  if (_entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  _entry->priority = prio;
  return OPENNSL_E_NONE;
}

int opennsl_field_qualifier_delete(int unit, opennsl_field_entry_t entry, opennsl_field_qualify_t qual_id) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: qualify   = %d", __func__, qual_id);

  _opennsl_field_entry_entry_t* const _entry = _opennsl_field_entry_entry_get(entry);
  if (_entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  memset(&_entry->qualifier, 0, sizeof(_opennsl_field_q_entry_t));

  return OPENNSL_E_NONE;
}

int opennsl_field_qualify_InPort(int unit, opennsl_field_entry_t entry, opennsl_port_t data, opennsl_port_t mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: port      = %d/%d", __func__, data, mask);

  _opennsl_field_entry_entry_t* const _entry = _opennsl_field_entry_entry_get(entry);
  if (_entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  _entry->qualifier.in_port = data;
  _entry->qualifier.in_port_mask = mask;

  return OPENNSL_E_NONE;
}

int opennsl_field_qualify_OutPort(int unit, opennsl_field_entry_t entry, opennsl_port_t data, opennsl_port_t mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: port      = %d/%d", __func__, data, mask);

  _opennsl_field_entry_entry_t* const _entry = _opennsl_field_entry_entry_get(entry);
  if (_entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  _entry->qualifier.out_port = data;
  _entry->qualifier.out_port_mask = mask;

  return OPENNSL_E_NONE;
}

int opennsl_field_qualify_InPorts(int unit, opennsl_field_entry_t entry, opennsl_pbmp_t data, opennsl_pbmp_t mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  _opennsl_pbmp_dump(__func__, &data);
  _opennsl_pbmp_dump(__func__, &mask);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_SrcPort(int unit, opennsl_field_entry_t entry, opennsl_module_t data_modid, opennsl_module_t mask_modid, opennsl_port_t data_port, opennsl_port_t mask_port) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: module    = %d/%d", __func__, data_modid, mask_modid);
  LOG_DEBUG("%s: port      = %d/%d", __func__, data_port, mask_port);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_DstPort(int unit, opennsl_field_entry_t entry, opennsl_module_t data_modid, opennsl_module_t mask_modid, opennsl_port_t data_port, opennsl_port_t mask_port) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: module    = %d/%d", __func__, data_modid, mask_modid);
  LOG_DEBUG("%s: port      = %d/%d", __func__, data_port, mask_port);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_DstTrunk(int unit, opennsl_field_entry_t entry, opennsl_trunk_t data, opennsl_trunk_t mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: trunk     = %d/%d", __func__, data, mask);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_L4SrcPort(int unit, opennsl_field_entry_t entry, opennsl_l4_port_t data, opennsl_l4_port_t mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: port      = %d/%d", __func__, data, mask);

_opennsl_field_entry_entry_t* const _entry = _opennsl_field_entry_entry_get(entry);
  if (_entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  _entry->qualifier.l4_sport = data;
  _entry->qualifier.l4_sport_mask = mask;

  return OPENNSL_E_NONE;
}

int opennsl_field_qualify_L4DstPort(int unit, opennsl_field_entry_t entry, opennsl_l4_port_t data, opennsl_l4_port_t mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: port      = %d/%d", __func__, data, mask);

  _opennsl_field_entry_entry_t* const _entry = _opennsl_field_entry_entry_get(entry);
  if (_entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  _entry->qualifier.l4_dport = data;
  _entry->qualifier.l4_dport_mask = mask;

  return OPENNSL_E_NONE;
}

int opennsl_field_qualify_OuterVlan(int unit, opennsl_field_entry_t entry, opennsl_vlan_t data, opennsl_vlan_t mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: vlan      = %d/%d", __func__, data, mask);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_OuterVlanId(int unit, opennsl_field_entry_t entry, opennsl_vlan_t data, opennsl_vlan_t mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: vlan      = %d/%d", __func__, data, mask);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_InnerVlanId(int unit, opennsl_field_entry_t entry, opennsl_vlan_t data, opennsl_vlan_t mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: vlan      = %d/%d", __func__, data, mask);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_EtherType(int unit, opennsl_field_entry_t entry, uint16 data, uint16 mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: ethertype = %04x/%04x", __func__, data, mask);

  _opennsl_field_entry_entry_t* const _entry = _opennsl_field_entry_entry_get(entry);
  if (_entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  _entry->qualifier.eth_type = data;
  _entry->qualifier.eth_type_mask = mask;

  return OPENNSL_E_NONE;
}

int opennsl_field_qualify_DstL3Egress(int unit, opennsl_field_entry_t entry, opennsl_if_t if_id) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: iface     = %d", __func__, if_id);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_Color(int unit, opennsl_field_entry_t entry, uint8 color) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: color     = %hhu", __func__, color);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_IpProtocol(int unit, opennsl_field_entry_t entry, uint8 data, uint8 mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: ip_proto  = %hhu/%hhu", __func__, data, mask);

  _opennsl_field_entry_entry_t* const _entry = _opennsl_field_entry_entry_get(entry);
  if (_entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  _entry->qualifier.ip_proto = data;
  _entry->qualifier.ip_proto_mask = mask;

  return OPENNSL_E_NONE;
}

int opennsl_field_qualify_PacketRes(int unit, opennsl_field_entry_t entry, uint32 data, uint32 mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: ip_proto  = %u/%u", __func__, data, mask);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_SrcIp(int unit, opennsl_field_entry_t entry, opennsl_ip_t data, opennsl_ip_t mask) {
  char ip4[32];
  char mask4[32];
  strip4(ip4, data);
  strip4(mask4, mask);

  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: src_ip    = %s/%s", __func__, ip4, mask4);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_DstIp(int unit, opennsl_field_entry_t entry, opennsl_ip_t data, opennsl_ip_t mask) {
  char ip4[32];
  char mask4[32];
  strip4(ip4, data);
  strip4(mask4, mask);

  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: src_ip    = %s/%s", __func__, ip4, mask4);

  _opennsl_field_entry_entry_t* const _entry = _opennsl_field_entry_entry_get(entry);
  if (_entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  _entry->qualifier.ip_dst = data;
  _entry->qualifier.ip_dst_mask = mask;

  return OPENNSL_E_NONE;
}

int opennsl_field_qualify_DSCP(int unit, opennsl_field_entry_t entry, uint8 data, uint8 mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: dscp      = %hhu/%hhu", __func__, data, mask);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_TcpControl(int unit, opennsl_field_entry_t entry, uint8 data, uint8 mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: tcp_ctrl  = %hhu/%hhu", __func__, data, mask);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_Ttl(int unit, opennsl_field_entry_t entry, uint8 data, uint8 mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: ttl       = %hhu/%hhu", __func__, data, mask);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_RangeCheck(int unit, opennsl_field_entry_t entry, opennsl_field_range_t range, int invert) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: range     = %u", __func__, range);
  LOG_DEBUG("%s: invert    = %d", __func__, invert);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_SrcIp6(int unit, opennsl_field_entry_t entry, opennsl_ip6_t data, opennsl_ip6_t mask) {
  char ip6[64];
  char mask6[64];
  strip6(ip6, data);
  strip6(mask6, mask);

  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: src_ip    = %s/%s", __func__, ip6, mask6);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_DstIp6(int unit, opennsl_field_entry_t entry, opennsl_ip6_t data, opennsl_ip6_t mask) {
  char ip6[64];
  char mask6[64];
  strip6(ip6, data);
  strip6(mask6, mask);

  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: src_ip    = %s/%s", __func__, ip6, mask6);

  _opennsl_field_entry_entry_t* const _entry = _opennsl_field_entry_entry_get(entry);
  if (_entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  memcpy(_entry->qualifier.ip6_dst, data, sizeof(opennsl_ip6_t));
  memcpy(_entry->qualifier.ip6_dst_mask, mask, sizeof(opennsl_ip6_t));

  return OPENNSL_E_NONE;
}

int opennsl_field_qualify_Ip6NextHeader(int unit, opennsl_field_entry_t entry, uint8 data, uint8 mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: next_hdr  = %hhu/%hhu", __func__, data, mask);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_Ip6HopLimit(int unit, opennsl_field_entry_t entry, uint8 data, uint8 mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: hop_limit = %hhu/%hhu", __func__, data, mask);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_SrcMac(int unit, opennsl_field_entry_t entry, opennsl_mac_t data, opennsl_mac_t mask) {
  char mac[64];
  char msk[64];
  strmac(mac, data);
  strmac(msk, mask);

  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: MAC       = %s/%s", __func__, mac, msk);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_DstMac(int unit, opennsl_field_entry_t entry, opennsl_mac_t data, opennsl_mac_t mask) {
  char mac[64];
  char msk[64];
  strmac(mac, data);
  strmac(msk, mask);

  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: MAC       = %s/%s", __func__, mac, msk);

  _opennsl_field_entry_entry_t* const _entry = _opennsl_field_entry_entry_get(entry);
  if (_entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  strmac(_entry->qualifier.mac_dst, data);
  strmac(_entry->qualifier.mac_dst_mask, mask);

  return OPENNSL_E_NONE;
}

int opennsl_field_qualify_IpType(int unit, opennsl_field_entry_t entry, opennsl_field_IpType_t type) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: iptype    = %hu", __func__, type);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_InterfaceClassPort(int unit, opennsl_field_entry_t entry, uint32 data, uint32 mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: port      = %u/%u", __func__, data, mask);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_SrcClassField(int unit, opennsl_field_entry_t entry, uint32 data, uint32 mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: field     = %u/%u", __func__, data, mask);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_DstClassField(int unit, opennsl_field_entry_t entry, uint32 data, uint32 mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: field     = %u/%u", __func__, data, mask);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_IpProtocolCommon(int unit, opennsl_field_entry_t entry, opennsl_field_IpProtocolCommon_t protocol) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: protocol  = %d", __func__, protocol);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_L3Routable(int unit, opennsl_field_entry_t entry, uint8 data, uint8 mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: table     = %hhu/%hhu", __func__, data, mask);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_IpFrag(int unit, opennsl_field_entry_t entry, opennsl_field_IpFrag_t frag) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: frag      = %d", __func__, frag);

  return OPENNSL_E_NONE;
}

int opennsl_field_qualify_Vrf(int unit, opennsl_field_entry_t entry, uint32 data, uint32 mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: vrf       = %u/%u", __func__, data, mask);

  _opennsl_field_entry_entry_t* const _entry = _opennsl_field_entry_entry_get(entry);
  if (_entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  _entry->qualifier.vrf = data;
  _entry->qualifier.vrf_mask = mask;

  return OPENNSL_E_NONE;
}

int opennsl_field_qualify_L3Ingress(int unit, opennsl_field_entry_t entry, uint32 data, uint32 mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: l3_egress = %u/%u", __func__, data, mask);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_MyStationHit(int unit, opennsl_field_entry_t entry, uint8 data, uint8 mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: _hit      = %hhu/%hhu", __func__, data, mask);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_IcmpTypeCode(int unit, opennsl_field_entry_t entry, uint16 data, uint16 mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: type_code = %u/%u", __func__, data, mask);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_Color_get(int unit, opennsl_field_entry_t entry, uint8 *color) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: color     = %p", __func__, color);
  *color = 1;
  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_DstL3Egress_get(int unit, opennsl_field_entry_t entry, opennsl_if_t *if_id) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: iface     = %p", __func__, if_id);
  *if_id = 1;
  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_InPort_get(int unit, opennsl_field_entry_t entry, opennsl_port_t *data, opennsl_port_t *mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: port      = %p/%p", __func__, data, mask);

  _opennsl_field_entry_entry_t* const _entry = _opennsl_field_entry_entry_get(entry);
  if (_entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  *data = _entry->qualifier.in_port;
  *mask = _entry->qualifier.in_port_mask;

  return OPENNSL_E_NONE;
}

int opennsl_field_qualify_OutPort_get(int unit, opennsl_field_entry_t entry, opennsl_port_t *data, opennsl_port_t *mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: port      = %p/%p", __func__, data, mask);

  _opennsl_field_entry_entry_t* const _entry = _opennsl_field_entry_entry_get(entry);
  if (_entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  *data = _entry->qualifier.out_port;
  *mask = _entry->qualifier.out_port_mask;

  return OPENNSL_E_NONE;
}

int opennsl_field_qualify_InPorts_get(int unit, opennsl_field_entry_t entry, opennsl_pbmp_t *data, opennsl_pbmp_t *mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: ports     = %p/%p", __func__, data, mask);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_SrcPort_get(int unit, opennsl_field_entry_t entry, opennsl_module_t *data_modid, opennsl_module_t *mask_modid, opennsl_port_t *data_port, opennsl_port_t *mask_port) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: module    = %p/%p", __func__, data_modid, mask_modid);
  LOG_DEBUG("%s: module    = %p/%p", __func__, data_port, mask_port);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_DstPort_get(int unit, opennsl_field_entry_t entry, opennsl_module_t *data_modid, opennsl_module_t *mask_modid, opennsl_port_t *data_port, opennsl_port_t *mask_port) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: module    = %p/%p", __func__, data_modid, mask_modid);
  LOG_DEBUG("%s: module    = %p/%p", __func__, data_port, mask_port);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_DstTrunk_get(int unit, opennsl_field_entry_t entry, opennsl_trunk_t *data, opennsl_trunk_t *mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: trunk     = %p/%p", __func__, data, mask);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_L4SrcPort_get(int unit, opennsl_field_entry_t entry, opennsl_l4_port_t *data, opennsl_l4_port_t *mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: l4 port   = %p/%p", __func__, data, mask);

  _opennsl_field_entry_entry_t* const _entry = _opennsl_field_entry_entry_get(entry);
  if (_entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  *data = _entry->qualifier.l4_sport;
  *mask = _entry->qualifier.l4_sport_mask;

  return OPENNSL_E_NONE;
}

int opennsl_field_qualify_L4DstPort_get(int unit, opennsl_field_entry_t entry, opennsl_l4_port_t *data, opennsl_l4_port_t *mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: l4 port   = %p/%p", __func__, data, mask);

  _opennsl_field_entry_entry_t* const _entry = _opennsl_field_entry_entry_get(entry);
  if (_entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  *data	= _entry->qualifier.l4_dport;
  *mask	= _entry->qualifier.l4_dport_mask;

  return OPENNSL_E_NONE;
}

int opennsl_field_qualify_OuterVlan_get(int unit, opennsl_field_entry_t entry, opennsl_vlan_t *data, opennsl_vlan_t *mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: vlan      = %p/%p", __func__, data, mask);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_OuterVlanId_get(int unit, opennsl_field_entry_t entry, opennsl_vlan_t *data, opennsl_vlan_t *mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: vlan      = %p/%p", __func__, data, mask);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_EtherType_get(int unit, opennsl_field_entry_t entry, uint16 *data, uint16 *mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: ethertype = %p/%p", __func__, data, mask);

  _opennsl_field_entry_entry_t* const _entry = _opennsl_field_entry_entry_get(entry);
  if (_entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  *data = _entry->qualifier.eth_type;
  *mask = _entry->qualifier.eth_type_mask;

  return OPENNSL_E_NONE;
}

int opennsl_field_qualify_IpProtocol_get(int unit, opennsl_field_entry_t entry, uint8 *data, uint8 *mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: ip_proto = %p/%p", __func__, data, mask);

  _opennsl_field_entry_entry_t* const _entry = _opennsl_field_entry_entry_get(entry);
  if (_entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  *data = _entry->qualifier.ip_proto;
  *mask = _entry->qualifier.ip_proto_mask;

  return OPENNSL_E_NONE;
}

int opennsl_field_qualify_PacketRes_get(int unit, opennsl_field_entry_t entry, uint32 *data, uint32 *mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: pkt_res   = %p/%p", __func__, data, mask);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_SrcIp_get(int unit, opennsl_field_entry_t entry, opennsl_ip_t *data, opennsl_ip_t *mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: ip        = %p/%p", __func__, data, mask);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_DstIp_get(int unit, opennsl_field_entry_t entry, opennsl_ip_t *data, opennsl_ip_t *mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: ip        = %p/%p", __func__, data, mask);

  _opennsl_field_entry_entry_t* const _entry = _opennsl_field_entry_entry_get(entry);
  if (_entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  *data = _entry->qualifier.ip_dst;
  *mask = _entry->qualifier.ip_dst_mask;

  return OPENNSL_E_NONE;
}

int opennsl_field_qualify_DSCP_get(int unit, opennsl_field_entry_t entry, uint8 *data, uint8 *mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: DSCP      = %p/%p", __func__, data, mask);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_TcpControl_get(int unit, opennsl_field_entry_t entry, uint8 *data, uint8 *mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: DSCP      = %p/%p", __func__, data, mask);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_Ttl_get(int unit, opennsl_field_entry_t entry, uint8 *data, uint8 *mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: DSCP      = %p/%p", __func__, data, mask);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_RangeCheck_get(int unit, opennsl_field_entry_t entry, int max_count, opennsl_field_range_t *range, int *invert, int *count) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: max_count = %d", __func__, max_count);
  LOG_DEBUG("%s: range     = %p", __func__, range);
  LOG_DEBUG("%s: invert    = %p", __func__, invert);
  LOG_DEBUG("%s: count     = %p", __func__, count);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_SrcIp6_get(int unit, opennsl_field_entry_t entry, opennsl_ip6_t *data, opennsl_ip6_t *mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: ip        = %p/%p", __func__, data, mask);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_DstIp6_get(int unit, opennsl_field_entry_t entry, opennsl_ip6_t *data, opennsl_ip6_t *mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: ip        = %p/%p", __func__, data, mask);

  _opennsl_field_entry_entry_t* const _entry = _opennsl_field_entry_entry_get(entry);
  if (_entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  memcpy(data, _entry->qualifier.ip6_dst, sizeof(opennsl_ip6_t));
  memcpy(mask, _entry->qualifier.ip6_dst_mask, sizeof(opennsl_ip6_t));

  return OPENNSL_E_NONE;
}

int opennsl_field_qualify_Ip6NextHeader_get(int unit, opennsl_field_entry_t entry, uint8 *data, uint8 *mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: next_hdr  = %p/%p", __func__, data, mask);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_Ip6HopLimit_get(int unit, opennsl_field_entry_t entry, uint8 *data, uint8 *mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: hop_limit = %p/%p", __func__, data, mask);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_SrcMac_get(int unit, opennsl_field_entry_t entry, opennsl_mac_t *data, opennsl_mac_t *mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: mac       = %p/%p", __func__, data, mask);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_DstMac_get(int unit, opennsl_field_entry_t entry, opennsl_mac_t *data, opennsl_mac_t *mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: mac       = %p/%p", __func__, data, mask);

  _opennsl_field_entry_entry_t* const _entry = _opennsl_field_entry_entry_get(entry);
  if (_entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  memcpy(data, _entry->qualifier.mac_dst, sizeof(opennsl_mac_t));
  memcpy(mask, _entry->qualifier.mac_dst_mask, sizeof(opennsl_mac_t));

  return OPENNSL_E_NONE;
}

int opennsl_field_qualify_IpType_get(int unit, opennsl_field_entry_t entry, opennsl_field_IpType_t *type) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: ip_type   = %p", __func__, type);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_InterfaceClassPort_get(int unit, opennsl_field_entry_t entry, uint32 *data, uint32 *mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: port      = %p/%p", __func__, data, mask);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_SrcClassField_get(int unit, opennsl_field_entry_t entry, uint32 *data, uint32 *mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: field     = %p/%p", __func__, data, mask);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_DstClassField_get(int unit, opennsl_field_entry_t entry, uint32 *data, uint32 *mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: field     = %p/%p", __func__, data, mask);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_IpProtocolCommon_get(int unit, opennsl_field_entry_t entry, opennsl_field_IpProtocolCommon_t *protocol) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: protocol  = %p", __func__, protocol);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_L3Routable_get(int unit, opennsl_field_entry_t entry, uint8 *data, uint8 *mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: table     = %p/%p", __func__, data, mask);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_IpFrag_get(int unit, opennsl_field_entry_t entry, opennsl_field_IpFrag_t *frag_info) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: frag      = %p", __func__, frag_info);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_L3Ingress_get(int unit, opennsl_field_entry_t entry, uint32 *data, uint32 *mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: iface     = %p/%p", __func__, data, mask);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_MyStationHit_get(int unit, opennsl_field_entry_t entry, uint8 *data, uint8 *mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: hit       = %p/%p", __func__, data, mask);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_IcmpTypeCode_get(int unit, opennsl_field_entry_t entry, uint16 *data, uint16 *mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: code      = %p/%p", __func__, data, mask);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_DstIpLocal(int unit, opennsl_field_entry_t entry, uint8 data, uint8 mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: ip local  = %hhu/%hhu", __func__, data, mask);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_DstIpLocal_get(int unit, opennsl_field_entry_t entry, uint8 *data, uint8 *mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: ip local  = %p/%p", __func__, data, mask);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_CpuQueue(int unit, opennsl_field_entry_t entry, uint8 data, uint8 mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: cpu queue = %hhu/%hhu", __func__, data, mask);

  _opennsl_field_entry_entry_t* const _entry = _opennsl_field_entry_entry_get(entry);
  if (_entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  _entry->qualifier.cpu_q = data;
  _entry->qualifier.cpu_q_mask = mask;

  return OPENNSL_E_NONE;
}

int opennsl_field_qualify_CpuQueue_get(int unit, opennsl_field_entry_t entry, uint8* data, uint8* mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: cpu queue = %p/%p", __func__, data, mask);

  _opennsl_field_entry_entry_t* const _entry = _opennsl_field_entry_entry_get(entry);
  if (_entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  *data = _entry->qualifier.cpu_q;
  *mask = _entry->qualifier.cpu_q_mask;

  return OPENNSL_E_NONE;
}

int opennsl_field_qualify_InterfaceClassProcessingPort(int unit, opennsl_field_entry_t entry, uint64 data, uint64 mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: port      = %llu/%llu", __func__, data, mask);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_qualify_InterfaceClassProcessingPort_get(int unit, opennsl_field_entry_t entry, uint64 *data, uint64 *mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: port      = %p/%p", __func__, data, mask);

  *data	= 9;
  *mask	= 0xff;
  return OPENNSL_E_UNAVAIL;
}

int opennsl_field_action_add(int unit, opennsl_field_entry_t entry, opennsl_field_action_t action, uint32 param0, uint32 param1) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: action    = %u (%08x)", __func__, action, action);
  LOG_DEBUG("%s: param0    = %u (%08x)", __func__, param0, param0);
  LOG_DEBUG("%s: param1    = %u (%08x)", __func__, param1, param1);

  return OPENNSL_E_NONE;
}

int opennsl_field_action_delete(int unit, opennsl_field_entry_t entry, opennsl_field_action_t action, uint32 param0, uint32 param1) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: action    = %d", __func__, action);
  LOG_DEBUG("%s: param0    = %u (%08x)", __func__, param0, param0);
  LOG_DEBUG("%s: param1    = %u (%08x)", __func__, param1, param1);

  return OPENNSL_E_NONE;
}

int opennsl_field_action_mac_add(int unit, opennsl_field_entry_t entry, opennsl_field_action_t action, opennsl_mac_t mac) {
  char macstr[64];
  strmac(macstr, mac);

  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: action    = %d", __func__, action);
  LOG_DEBUG("%s: mac       = '%s'", __func__, mac);

  return OPENNSL_E_NONE;
}

int opennsl_field_action_ports_add(int unit, opennsl_field_entry_t entry, opennsl_field_action_t action, opennsl_pbmp_t pbmp) {

  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: action    = %d", __func__, action);
  _opennsl_pbmp_dump(__func__, &pbmp);

  return OPENNSL_E_NONE;
}

int opennsl_field_action_get(int unit, opennsl_field_entry_t entry, opennsl_field_action_t action, uint32 *param0, uint32 *param1) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: action    = %d", __func__, action);
  LOG_DEBUG("%s: param0    = %p", __func__, param0);
  LOG_DEBUG("%s: param1    = %p", __func__, param1);

  *param0 = 1;
  *param1 = 2;
  return OPENNSL_E_NONE;
}

int opennsl_field_action_mac_get(int unit, opennsl_field_entry_t entry, opennsl_field_action_t action, opennsl_mac_t *mac) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: action    = %d", __func__, action);
  LOG_DEBUG("%s: mac       = %p", __func__, mac);

  opennsl_mac_t m = {0,1,2,3,4,5};
  memcpy(mac, m, 6);
  return OPENNSL_E_NONE;
}

int opennsl_field_action_ports_get(int unit, opennsl_field_entry_t entry, opennsl_field_action_t action, opennsl_pbmp_t *pbmp) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: action    = %d", __func__, action);
  LOG_DEBUG("%s: pbmp      = %p", __func__, pbmp);

  return OPENNSL_E_NONE;
}

int opennsl_field_action_remove(int unit, opennsl_field_entry_t entry, opennsl_field_action_t action) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: action    = %d", __func__, action);

  return OPENNSL_E_NONE;
}

int opennsl_field_action_remove_all(int unit, opennsl_field_entry_t entry) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);

  return OPENNSL_E_NONE;
}

int opennsl_field_stat_create(int unit, opennsl_field_group_t group, int nstat, opennsl_field_stat_t *stat_arr, int *stat_id) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: group     = %d", __func__, group);
  LOG_DEBUG("%s: nstat     = %d", __func__, nstat);
  LOG_DEBUG("%s: stat_id   = %p", __func__, stat_id);
  int index;
  for (index = 0; index < nstat; index++) {
    LOG_DEBUG("%s: stat[%d]  = %d", __func__, index, stat_arr[index]);
  }

  *stat_id = 101;
  return OPENNSL_E_NONE;
}

int opennsl_field_stat_create_id(int unit, opennsl_field_group_t group, int nstat, opennsl_field_stat_t *stat_arr, int stat_id) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: group     = %d", __func__, group);
  LOG_DEBUG("%s: nstat     = %d", __func__, nstat);
  LOG_DEBUG("%s: stat_id   = %d", __func__, stat_id);
  int index;
  for (index = 0; index < nstat; index++) {
    LOG_DEBUG("%s: stat[%d]  = %d", __func__, index, stat_arr[index]);
  }

  return OPENNSL_E_NONE;
}

int opennsl_field_stat_id_get(int unit, opennsl_field_group_t group, uint32 stat_id, uint32 *stat_counter_id) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: group     = %d", __func__, group);
  LOG_DEBUG("%s: stat_id   = %u", __func__, stat_id);
  LOG_DEBUG("%s: counter_id= %p", __func__, stat_counter_id);

  *stat_counter_id = 5;
  return OPENNSL_E_NONE;
}

int opennsl_field_stat_destroy(int unit, int stat_id) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: stat_id   = %u", __func__, stat_id);

  return OPENNSL_E_NONE;
}

int opennsl_field_stat_size(int unit, int stat_id, int *stat_size) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: stat_id   = %u", __func__, stat_id);
  LOG_DEBUG("%s: stat_size = %p", __func__, stat_size);

  *stat_size = 3;
  return OPENNSL_E_NONE;
}

int opennsl_field_stat_set(int unit, int stat_id, opennsl_field_stat_t stat, uint64 value) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: stat_id   = %u", __func__, stat_id);
  LOG_DEBUG("%s: stat      = %d:%llu", __func__, stat, value);

  return OPENNSL_E_NONE;
}

int opennsl_field_stat_all_set(int unit, int stat_id, uint64 value) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: stat_id   = %u", __func__, stat_id);
  LOG_DEBUG("%s: stat      = all:%llu", __func__, value);

  return OPENNSL_E_NONE;
}

int opennsl_field_stat_get(int unit, int stat_id, opennsl_field_stat_t stat, uint64 *value) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: stat_id   = %u", __func__, stat_id);
  LOG_DEBUG("%s: stat      = %d:%p", __func__, stat, value);

  *value = 10000;
  return OPENNSL_E_NONE;
}

int opennsl_field_entry_stat_attach(int unit, opennsl_field_entry_t entry, int stat_id) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: stat_id   = %u", __func__, stat_id);

  return OPENNSL_E_NONE;
}

int opennsl_field_entry_stat_detach(int unit, opennsl_field_entry_t entry, int stat_id) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: stat_id   = %u", __func__, stat_id);

  return OPENNSL_E_NONE;
}

int opennsl_field_entry_stat_get(int unit, opennsl_field_entry_t entry, int *stat_id) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: stat_id   = %p", __func__, stat_id);

  *stat_id = 5;
  return OPENNSL_E_NONE;
}

int opennsl_field_stat_detach(int unit, uint32 stat_id) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: stat_id   = %u", __func__, stat_id);

  return OPENNSL_E_NONE;
}

int opennsl_field_presel_create(int unit, opennsl_field_presel_t *presel_id) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: presel_id = %p", __func__, presel_id);

  *presel_id = 21;
  return OPENNSL_E_NONE;
}

int opennsl_field_presel_create_id(int unit, opennsl_field_presel_t presel_id) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: presel_id = %d", __func__, presel_id);

  return OPENNSL_E_NONE;
}

int opennsl_field_presel_destroy(int unit, opennsl_field_presel_t presel_id) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: presel_id = %d", __func__, presel_id);

  return OPENNSL_E_NONE;
}

int opennsl_field_group_presel_set(int unit, opennsl_field_group_t group, opennsl_field_presel_set_t *presel) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: group     = %d", __func__, group);
  _opennsl_field_presel_set_dump(__func__, presel);

  return OPENNSL_E_NONE;
}

int opennsl_field_group_presel_get(int unit, opennsl_field_group_t group, opennsl_field_presel_set_t *presel) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: group     = %d", __func__, group);
  LOG_DEBUG("%s: presel    = %p", __func__, presel);

  return OPENNSL_E_NONE;
}

int opennsl_field_qualify_IngressClassField(int unit, opennsl_field_entry_t entry, uint32 data, uint32 mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: field     = %d/%d", __func__, data, mask);

  return OPENNSL_E_NONE;
}

int opennsl_field_qualify_IngressClassField_get(int unit, opennsl_field_entry_t entry, uint32 *data, uint32 *mask) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: field     = %p/%p", __func__, data, mask);

  *data = 123;
  *mask = 0xfff;
  return OPENNSL_E_NONE;
}

int opennsl_field_qualify_Stage(int unit, opennsl_field_entry_t entry, opennsl_field_stage_t data) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: stage     = %d", __func__, data);

  return OPENNSL_E_NONE;
}

int opennsl_field_qualify_Stage_get(int unit, opennsl_field_entry_t entry, opennsl_field_stage_t *data) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: entry     = %d", __func__, entry);
  LOG_DEBUG("%s: stage     = %p", __func__, data);

  *data = 10;
  return OPENNSL_E_NONE;
}
