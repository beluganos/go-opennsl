// -*- coding: utf-8 -*-

#include <string.h>
#include <opennsl/error.h>
#include <opennsl/l3.h>
#include <opennsl/tunnelX.h>
#include "libopennsl.h"
#include "logger.h"

//
// opennsl_tunnel_initiator_t
//
typedef struct _opennsl_tunnel_initiator_entry {
  opennsl_tunnel_initiator_t tuninit;
} _opennsl_tunnel_initiator_entry_t;

#define _OPENNSL_TUNNEL_INITIATOR_ENTRY_MAX (256)

static _opennsl_tunnel_initiator_entry_t s_opennsl_tunnel_initiator_entries[_OPENNSL_TUNNEL_INITIATOR_ENTRY_MAX] = {0};

static _opennsl_tunnel_initiator_entry_t* _opennsl_tunnel_initiator_entry_get(opennsl_gport_t id) {
  int index;
  for (index = 0; index < _OPENNSL_TUNNEL_INITIATOR_ENTRY_MAX; index++) {
    _opennsl_tunnel_initiator_entry_t* const entry = s_opennsl_tunnel_initiator_entries + index;
    if (entry->tuninit.tunnel_id == id)
      return entry;
  }
  return NULL;
}

static _opennsl_tunnel_initiator_entry_t* _opennsl_tunnel_initiator_entry_get_by_key(opennsl_if_t id) {
  int index;
  for (index = 0; index < _OPENNSL_TUNNEL_INITIATOR_ENTRY_MAX; index++) {
    _opennsl_tunnel_initiator_entry_t* const entry = s_opennsl_tunnel_initiator_entries + index;
    if (entry->tuninit.l3_intf_id == id)
      return entry;
  }
  return NULL;

}

//
// opennsl_tunnel_terminator_t
//
typedef struct _opennsl_tunnel_terminator_entry {
  opennsl_tunnel_terminator_t tunterm;
} _opennsl_tunnel_terminator_entry_t;

#define _OPENNSL_TUNNEL_TERMINATOR_ENTRY_MAX (256)

static _opennsl_tunnel_terminator_entry_t s_opennsl_tunnel_terminator_entries[_OPENNSL_TUNNEL_TERMINATOR_ENTRY_MAX] = {0};

static _opennsl_tunnel_terminator_entry_t* _opennsl_tunnel_terminator_entry_get(opennsl_gport_t id) {
  int index;
  for (index = 0; index < _OPENNSL_TUNNEL_TERMINATOR_ENTRY_MAX; index++) {
    _opennsl_tunnel_terminator_entry_t* const entry = s_opennsl_tunnel_terminator_entries + index;
    if (entry->tunterm.tunnel_id == id)
      return entry;
  }
  return NULL;
}

static _opennsl_tunnel_terminator_entry_t* _opennsl_tunnel_terminator_entry_get_by_key(const opennsl_tunnel_terminator_t* info) {

  int index;
  for (index = 0; index < _OPENNSL_TUNNEL_TERMINATOR_ENTRY_MAX; index++) {
    _opennsl_tunnel_terminator_entry_t* const entry = s_opennsl_tunnel_terminator_entries + index;
    if ((info->dip != entry->tunterm.dip) || (info->sip != entry->tunterm.sip))
      continue;
    if ((info->dip_mask != entry->tunterm.dip_mask) || (info->sip_mask != entry->tunterm.sip_mask))
      continue;
    if ((_opennsl_ip6_cmp(info->dip6, entry->tunterm.dip6) != 0) ||
	(_opennsl_ip6_cmp(info->sip6, entry->tunterm.sip6) != 0))
      continue;
    if ((_opennsl_ip6_cmp(info->dip6_mask, entry->tunterm.dip6_mask) != 0) ||
	(_opennsl_ip6_cmp(info->sip6_mask, entry->tunterm.sip6_mask) != 0))
      continue;
    if (info->vlan != entry->tunterm.vlan)
      continue;
    if (info->vrf != entry->tunterm.vrf)
      continue;

    return entry;
  }

  return NULL;
}

int opennsl_tunnel_initiator_set(int unit, opennsl_l3_intf_t *intf, opennsl_tunnel_initiator_t *tunnel) {
  LOG_DEBUG("%s: unit   = %d", __func__, unit);
  _opennsl_l3_intf_dump(__func__, intf);
  _opennsl_tunnel_initiator_dump(__func__, tunnel);

  int rc;
  opennsl_l3_intf_t l3_intf;
  l3_intf = *intf;
  rc = opennsl_l3_intf_get(unit, &l3_intf);
  if (rc != OPENNSL_E_NONE)
    return rc;

  if (tunnel->flags & OPENNSL_TUNNEL_REPLACE) {
    _opennsl_tunnel_initiator_entry_t* const entry = _opennsl_tunnel_initiator_entry_get_by_key(tunnel->l3_intf_id);
    if (entry == NULL)
      return OPENNSL_E_NOT_FOUND;

    entry->tuninit = *tunnel;
    return OPENNSL_E_NONE;
  }

  if ((tunnel->flags & OPENNSL_TUNNEL_WITH_ID) != 0) {
    _opennsl_tunnel_initiator_entry_t* const entry = _opennsl_tunnel_initiator_entry_get(tunnel->tunnel_id);
      if (entry != NULL)
	return OPENNSL_E_EXISTS;
  }

  int index;
  for (index = 0; index < _OPENNSL_TUNNEL_INITIATOR_ENTRY_MAX; index++) {
    _opennsl_tunnel_initiator_entry_t* const entry = s_opennsl_tunnel_initiator_entries + index;
    if (entry->tuninit.tunnel_id == 0) {
      entry->tuninit = *tunnel;

      if ((tunnel->flags & OPENNSL_TUNNEL_WITH_ID) == 0) {
	entry->tuninit.tunnel_id = index + 1;
      }

      return OPENNSL_E_NONE;
    }
  }

  return OPENNSL_E_RESOURCE;
}

int opennsl_tunnel_initiator_create(int unit, opennsl_l3_intf_t *intf, opennsl_tunnel_initiator_t *tunnel) {
  LOG_DEBUG("%s: unit   = %d", __func__, unit);
  _opennsl_l3_intf_dump(__func__, intf);
  _opennsl_tunnel_initiator_dump(__func__, tunnel);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_tunnel_initiator_clear(int unit,  opennsl_l3_intf_t *intf) {
  LOG_DEBUG("%s: unit   = %d", __func__, unit);
  _opennsl_l3_intf_dump(__func__, intf);

  int rc;
  opennsl_l3_intf_t l3_intf;
  l3_intf = *intf;
  rc = opennsl_l3_intf_get(unit, &l3_intf);
  if (rc != OPENNSL_E_NONE)
    return rc;

  _opennsl_tunnel_initiator_entry_t* const entry = _opennsl_tunnel_initiator_entry_get_by_key(intf->l3a_intf_id);
  if (entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  memset(entry, 0, sizeof(_opennsl_tunnel_initiator_entry_t));
  return OPENNSL_E_NONE;
}

int opennsl_tunnel_initiator_get(int unit, opennsl_l3_intf_t *intf, opennsl_tunnel_initiator_t *tunnel) {
  LOG_DEBUG("%s: unit   = %d", __func__, unit);
  _opennsl_l3_intf_dump(__func__, intf);
  LOG_DEBUG("%s: tunnel = %p", __func__, tunnel);

  int rc;
  opennsl_l3_intf_t l3_intf;
  l3_intf = *intf;
  rc = opennsl_l3_intf_get(unit, &l3_intf);
  if (rc != OPENNSL_E_NONE)
    return rc;

  _opennsl_tunnel_initiator_entry_t* const entry = _opennsl_tunnel_initiator_entry_get_by_key(intf->l3a_intf_id);
  if (entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  *tunnel = entry->tuninit;
  return OPENNSL_E_NONE;
}

void opennsl_tunnel_initiator_t_init(opennsl_tunnel_initiator_t *tunnel_init) {
  memset(tunnel_init, 0, sizeof(opennsl_tunnel_initiator_t));
}

void opennsl_tunnel_terminator_t_init(opennsl_tunnel_terminator_t *tunnel_term) {
  memset(tunnel_term, 0, sizeof(opennsl_tunnel_terminator_t));
}

int opennsl_tunnel_terminator_add(int unit, opennsl_tunnel_terminator_t *info) {
  LOG_DEBUG("%s: unit   = %d", __func__, unit);
  _opennsl_tunnel_terminator_dump(__func__, info);
  if (info->flags & OPENNSL_TUNNEL_REPLACE) {
    _opennsl_tunnel_terminator_entry_t* const entry = _opennsl_tunnel_terminator_entry_get_by_key(info);
    if (entry == NULL)
      return OPENNSL_E_NOT_FOUND;

    entry->tunterm = *info;
    return OPENNSL_E_NONE;
  }

  if ((info->flags & OPENNSL_TUNNEL_TERM_TUNNEL_WITH_ID) != 0) {
    _opennsl_tunnel_terminator_entry_t* const entry = _opennsl_tunnel_terminator_entry_get(info->tunnel_id);
    if (entry != NULL)
      return OPENNSL_E_EXISTS;
  }

  int index;
  for (index = 0; index < _OPENNSL_TUNNEL_TERMINATOR_ENTRY_MAX; index++) {
    _opennsl_tunnel_terminator_entry_t* const entry = s_opennsl_tunnel_terminator_entries + index;
    if (entry->tunterm.tunnel_id == 0) {
      entry->tunterm = *info;

      if ((info->flags & OPENNSL_TUNNEL_TERM_TUNNEL_WITH_ID) != 0) {
	entry->tunterm.tunnel_id = info->tunnel_id;
      } else {
	entry->tunterm.tunnel_id = index + 1;
      }

      return OPENNSL_E_NONE;
    }
  }

  return OPENNSL_E_RESOURCE;
}

int opennsl_tunnel_terminator_create(int unit, opennsl_tunnel_terminator_t *info) {
  LOG_DEBUG("%s: unit   = %d", __func__, unit);
  _opennsl_tunnel_terminator_dump(__func__, info);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_tunnel_terminator_delete(int unit, opennsl_tunnel_terminator_t *info) {
  LOG_DEBUG("%s: unit   = %d", __func__, unit);
  _opennsl_tunnel_terminator_dump(__func__, info);

  _opennsl_tunnel_terminator_entry_t* const entry = _opennsl_tunnel_terminator_entry_get_by_key(info);
  if (entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  memset(entry, 0, sizeof(_opennsl_tunnel_terminator_entry_t));
  return OPENNSL_E_NONE;
}

int opennsl_tunnel_terminator_update(int unit, opennsl_tunnel_terminator_t *info) {
  LOG_DEBUG("%s: unit   = %d", __func__, unit);
  _opennsl_tunnel_terminator_dump(__func__, info);

  _opennsl_tunnel_terminator_entry_t* const entry = _opennsl_tunnel_terminator_entry_get_by_key(info);
  if (entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  entry->tunterm = *info;;
  return OPENNSL_E_NONE;
}

int opennsl_tunnel_terminator_get(int unit, opennsl_tunnel_terminator_t *info) {
  LOG_DEBUG("%s: unit   = %d", __func__, unit);
  _opennsl_tunnel_terminator_dump(__func__, info);

  _opennsl_tunnel_terminator_entry_t* const entry = _opennsl_tunnel_terminator_entry_get_by_key(info);
  if (entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  *info = entry->tunterm;
  return OPENNSL_E_NONE;
}

int opennsl_tunnel_initiator_traverse (int unit, opennsl_tunnel_initiator_traverse_cb cb, void *user_data) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: callback = %p", __func__, cb);
  LOG_DEBUG("%s: data     = %p", __func__, user_data);

  int index;
  for (index = 0; index < _OPENNSL_TUNNEL_INITIATOR_ENTRY_MAX; index++) {
    _opennsl_tunnel_initiator_entry_t* const entry = s_opennsl_tunnel_initiator_entries + index;
    if (entry->tuninit.tunnel_id == 0)
      continue;

    int rc = cb(unit, &entry->tuninit, user_data);
    if (rc != OPENNSL_E_NONE)
      return rc;
  }

  return OPENNSL_E_NONE;
}

int opennsl_tunnel_terminator_traverse (int unit, opennsl_tunnel_terminator_traverse_cb cb, void *user_data) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: callback = %p", __func__, cb);
  LOG_DEBUG("%s: data     = %p", __func__, user_data);

  int index;
  for (index = 0; index < _OPENNSL_TUNNEL_TERMINATOR_ENTRY_MAX; index++) {
    _opennsl_tunnel_terminator_entry_t* const entry = s_opennsl_tunnel_terminator_entries + index;
    if (entry->tunterm.tunnel_id == 0)
      continue;

    int rc = cb(unit, &entry->tunterm, user_data);
    if (rc != OPENNSL_E_NONE)
      return rc;
  }

  return OPENNSL_E_NONE;
}
