// -*- coding: utf-8 -*-

#include <stdlib.h>
#include <string.h>
#include <opennsl/error.h>
#include <opennsl/l3.h>
#include "libopennsl.h"
#include "logger.h"

typedef struct _opennsl_l3_intf_entry {
  opennsl_l3_intf_t l3intf;
} _opennsl_l3_intf_entry_t;

#define _OPENNSL_L3_INTERFACE_ENTRY_MAX (256)

static _opennsl_l3_intf_entry_t s_opennsl_l3_intf_entries[_OPENNSL_L3_INTERFACE_ENTRY_MAX] = {0};

static _opennsl_l3_intf_entry_t* _opennsl_l3_intf_entry_get(opennsl_if_t id) {
  int index;
  for (index = 0; index < _OPENNSL_L3_INTERFACE_ENTRY_MAX; index++) {
    _opennsl_l3_intf_entry_t* const entry = s_opennsl_l3_intf_entries + index;
    if (entry->l3intf.l3a_intf_id == id)
      return entry;
  }
  return NULL;
}

typedef struct _opennsl_l3_egress_entry {
  opennsl_if_t id;
  opennsl_l3_egress_t l3egr;
} _opennsl_l3_egress_entry_t;

#define _OPENNSL_L3_EGRESS_ENTRY_MAX (256)

static _opennsl_l3_egress_entry_t s_opennsl_l3_egress_entries[_OPENNSL_L3_EGRESS_ENTRY_MAX] = {0};

static _opennsl_l3_egress_entry_t* _opennsl_l3_egress_entry_get(opennsl_if_t id) {
  int index;
  for (index = 0; index < _OPENNSL_L3_EGRESS_ENTRY_MAX; index++) {
    _opennsl_l3_egress_entry_t* const entry = s_opennsl_l3_egress_entries + index;
    if (entry->id == id)
      return entry;
  }
  return NULL;
}

typedef struct _opennsl_l3_ingress_entry {
  opennsl_if_t id;
  opennsl_l3_ingress_t l3igr;
} _opennsl_l3_ingress_entry_t;

#define _OPENNSL_L3_INGRESS_ENTRY_MAX (256)

static _opennsl_l3_ingress_entry_t s_opennsl_l3_ingress_entries[_OPENNSL_L3_INGRESS_ENTRY_MAX] = {0};

static _opennsl_l3_ingress_entry_t* _opennsl_l3_ingress_entry_get(opennsl_if_t id) {
  int index;
  for (index = 0; index < _OPENNSL_L3_INGRESS_ENTRY_MAX; index++) {
    _opennsl_l3_ingress_entry_t* const entry = s_opennsl_l3_ingress_entries + index;
    if (entry->id == id)
      return entry;
  }

  return NULL;
}

typedef struct _opennsl_l3_host_entry {
  int id;
  opennsl_l3_host_t l3host;
} _opennsl_l3_host_entry_t;

#define _OPENNSL_L3_HOST_ENTRY_MAX (256)

static _opennsl_l3_host_entry_t s_opennsl_l3_host4_entries[_OPENNSL_L3_HOST_ENTRY_MAX] = {0};
static _opennsl_l3_host_entry_t s_opennsl_l3_host6_entries[_OPENNSL_L3_HOST_ENTRY_MAX] = {0};

static _opennsl_l3_host_entry_t* _opennsl_l3_host4_entry_get(opennsl_ip_t ip, opennsl_vrf_t vrf) {
  int index;
  for (index = 0; index < _OPENNSL_L3_HOST_ENTRY_MAX; index++) {
    _opennsl_l3_host_entry_t* const entry = s_opennsl_l3_host4_entries + index;
    if ((entry->l3host.l3a_ip_addr == ip) && (entry->l3host.l3a_vrf == vrf))
      return entry;
  }
  return NULL;
}

static _opennsl_l3_host_entry_t* _opennsl_l3_host6_entry_get(opennsl_ip6_t ip, opennsl_vrf_t vrf) {
  int index;
  for (index = 0; index < _OPENNSL_L3_HOST_ENTRY_MAX; index++) {
    _opennsl_l3_host_entry_t* const entry = s_opennsl_l3_host6_entries + index;
    if ((memcmp(entry->l3host.l3a_ip6_addr, ip, sizeof(opennsl_ip6_t)) == 0) &&
	(entry->l3host.l3a_vrf == vrf))
      return entry;
  }
  return NULL;
}

typedef struct _opennsl_l3_route_entry {
  int id;
  opennsl_l3_route_t l3route;
} _opennsl_l3_route_entry_t;

#define _OPENNSL_L3_ROUTE_ENTRY_MAX (512)

static _opennsl_l3_route_entry_t s_opennsl_l3_route4_entries[_OPENNSL_L3_ROUTE_ENTRY_MAX];
static _opennsl_l3_route_entry_t s_opennsl_l3_route6_entries[_OPENNSL_L3_ROUTE_ENTRY_MAX];

static _opennsl_l3_route_entry_t* _opennsl_l3_route4_entry_get(
							       opennsl_ip_t ip,
							       opennsl_ip_t mask,
							       opennsl_vrf_t vrf
) {
  int index;
  for (index = 0; index < _OPENNSL_L3_ROUTE_ENTRY_MAX; index++) {
    _opennsl_l3_route_entry_t* const entry = s_opennsl_l3_route4_entries + index;
    if (entry->l3route.l3a_subnet == ip &&
	entry->l3route.l3a_ip_mask == mask &&
	entry->l3route.l3a_vrf == vrf)

      return entry;
  }
  return NULL;
}

static _opennsl_l3_route_entry_t* _opennsl_l3_route6_entry_get(
							       opennsl_ip6_t ip,
							       opennsl_ip6_t mask,
							       opennsl_vrf_t vrf
) {
  int index;
  for (index = 0; index < _OPENNSL_L3_ROUTE_ENTRY_MAX; index++) {
    _opennsl_l3_route_entry_t* const entry = s_opennsl_l3_route6_entries + index;
    if (memcmp(entry->l3route.l3a_ip6_net, ip, sizeof(opennsl_ip6_t)) == 0 &&
	memcmp(entry->l3route.l3a_ip6_mask, mask, sizeof(opennsl_ip6_t)) == 0 &&
	entry->l3route.l3a_vrf == vrf)

      return entry;
  }
  return NULL;
}

int opennsl_l3_init(int unit) {
  LOG_DEBUG("%s: unit = %d", __func__, unit);
  return OPENNSL_E_NONE;
}

int opennsl_l3_info(int unit, opennsl_l3_info_t *l3info) {
  LOG_DEBUG("%s: unit = %d", __func__, unit);
  LOG_DEBUG("%s: info = %p", __func__, l3info);

  return OPENNSL_E_NONE;
}

void opennsl_l3_intf_t_init(opennsl_l3_intf_t *intf) {
  memset(intf, 0, sizeof(opennsl_l3_intf_t));
}

int opennsl_l3_intf_create(int unit, opennsl_l3_intf_t *intf) {
  LOG_DEBUG("%s: unit = %d", __func__, unit);
  _opennsl_l3_intf_dump(__func__, intf);

  if (intf->l3a_flags & OPENNSL_L3_REPLACE) {
    _opennsl_l3_intf_entry_t* const entry = _opennsl_l3_intf_entry_get(intf->l3a_intf_id);
    if (entry == NULL)
      return OPENNSL_E_NOT_FOUND;

    entry->l3intf = *intf;
    return OPENNSL_E_NONE;
  }

  if ((intf->l3a_flags & OPENNSL_L3_WITH_ID) != 0) {
    _opennsl_l3_intf_entry_t* const entry = _opennsl_l3_intf_entry_get(intf->l3a_intf_id);
    if (entry != NULL)
      return OPENNSL_E_EXISTS;
  }

  int index;
  for (index = 0; index < _OPENNSL_L3_INTERFACE_ENTRY_MAX; index++) {
    _opennsl_l3_intf_entry_t* const entry = s_opennsl_l3_intf_entries + index;
    if (entry->l3intf.l3a_intf_id == 0) {

      entry->l3intf = *intf;

      if ((intf->l3a_flags & OPENNSL_L3_WITH_ID) == 0) {
	entry->l3intf.l3a_intf_id = index + 1;
	intf->l3a_intf_id = index + 1;
      }

      return OPENNSL_E_NONE;
    }
  }

  return OPENNSL_E_RESOURCE;
}

int opennsl_l3_intf_delete(int unit, opennsl_l3_intf_t *intf) {
  LOG_DEBUG("%s: unit = %d", __func__, unit);
  _opennsl_l3_intf_dump(__func__, intf);

  return OPENNSL_E_NONE;
}

int opennsl_l3_intf_find(int unit, opennsl_l3_intf_t *intf) {
  LOG_DEBUG("%s: unit = %d", __func__, unit);
  _opennsl_l3_intf_dump(__func__, intf);

  return OPENNSL_E_NONE;
}

int opennsl_l3_intf_find_vlan(int unit, opennsl_l3_intf_t *intf) {
  LOG_DEBUG("%s: unit = %d", __func__, unit);
  _opennsl_l3_intf_dump(__func__, intf);

  return OPENNSL_E_NONE;
}

int opennsl_l3_intf_get(int unit, opennsl_l3_intf_t *intf) {
  LOG_DEBUG("%s: unit = %d", __func__, unit);
  _opennsl_l3_intf_dump(__func__, intf);

  _opennsl_l3_intf_entry_t* const entry = _opennsl_l3_intf_entry_get(intf->l3a_intf_id);
  if (entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  *intf = entry->l3intf;
  return OPENNSL_E_NONE;
}

void opennsl_l3_egress_t_init(opennsl_l3_egress_t *egr) {
  memset(egr, 0, sizeof(opennsl_l3_egress_t));
}

int opennsl_l3_egress_create(int unit, uint32 flags, opennsl_l3_egress_t *egr, opennsl_if_t *if_id) {
  LOG_DEBUG("%s: unit  = %d", __func__, unit);
  LOG_DEBUG("%s: iface = %d", __func__, *if_id);
  _opennsl_l3_egress_dump(__func__, egr);

  if (flags & OPENNSL_L3_REPLACE) {
    _opennsl_l3_egress_entry_t* const entry = _opennsl_l3_egress_entry_get(*if_id);
    if (entry == NULL)
      return OPENNSL_E_NOT_FOUND;

    entry->l3egr = *egr;
    return OPENNSL_E_NONE;
  }

  if ((flags & OPENNSL_L3_WITH_ID) != 0) {
    if ((*if_id) < 1)
      return OPENNSL_E_PARAM;

    _opennsl_l3_egress_entry_t* const entry = _opennsl_l3_egress_entry_get(*if_id);
    if (entry != NULL)
      return OPENNSL_E_EXISTS;
  }

  int index;
  for (index = 0; index < _OPENNSL_L3_EGRESS_ENTRY_MAX; index++) {
    _opennsl_l3_egress_entry_t*	const entry = s_opennsl_l3_egress_entries + index;
    if (entry->id == 0) {
      if ((flags & OPENNSL_L3_WITH_ID) != 0) {
	entry->id = *if_id;
      } else {
	entry->id = index + 1;
	*if_id    = index + 1;
      }

      entry->l3egr = *egr;

      return OPENNSL_E_NONE;
    }
  }

  return OPENNSL_E_RESOURCE;
}

int opennsl_l3_egress_destroy(int unit, opennsl_if_t intf) {
  LOG_DEBUG("%s: unit  = %d", __func__, unit);
  LOG_DEBUG("%s: iface = %d", __func__, intf);

  _opennsl_l3_egress_entry_t* const entry = _opennsl_l3_egress_entry_get(intf);
  if (entry == NULL)
    return OPENNSL_E_BADID;

  entry->id = 0;

  return OPENNSL_E_NONE;
}

int opennsl_l3_egress_get(int unit, opennsl_if_t intf, opennsl_l3_egress_t *egr) {
  LOG_DEBUG("%s: unit  = %d", __func__, unit);
  LOG_DEBUG("%s: iface = %d", __func__, intf);
  LOG_DEBUG("%s: engre = %p", __func__, egr);

  _opennsl_l3_egress_entry_t* const entry = _opennsl_l3_egress_entry_get(intf);
  if (entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  *egr = entry->l3egr;

  return OPENNSL_E_NONE;
}

int opennsl_l3_egress_find(int unit, opennsl_l3_egress_t *egr, opennsl_if_t *intf) {
  LOG_DEBUG("%s: unit  = %d", __func__, unit);
  LOG_DEBUG("%s: iface = %d", __func__, *intf);
  _opennsl_l3_egress_dump(__func__, egr);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_l3_egress_traverse(int unit, opennsl_l3_egress_traverse_cb trav_fn, void *user_data) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: callback = %p", __func__, trav_fn);
  LOG_DEBUG("%s: user_data= %p", __func__, user_data);

  int index;
  for (index = 0; index < _OPENNSL_L3_EGRESS_ENTRY_MAX; index++) {
    _opennsl_l3_egress_entry_t* const entry = s_opennsl_l3_egress_entries + index;
    if (entry->id == 0)
      continue;

    const int rc = trav_fn(unit, entry->id, &entry->l3egr, user_data);
    if (rc != OPENNSL_E_NONE)
      return rc;
  }

  return OPENNSL_E_NONE;
}

void opennsl_l3_egress_ecmp_t_init(opennsl_l3_egress_ecmp_t *ecmp) {
  memset(ecmp, 0, sizeof(opennsl_l3_egress_ecmp_t));
}

int opennsl_l3_ecmp_get(int unit, opennsl_l3_egress_ecmp_t *ecmp_info, int ecmp_member_size, opennsl_l3_ecmp_member_t *ecmp_member_array, int *ecmp_member_count) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: size     = %d", __func__, ecmp_member_size);
  LOG_DEBUG("%s: array    = %p", __func__, ecmp_member_array);
  LOG_DEBUG("%s: count    = %p", __func__, ecmp_member_count);
  _opennsl_l3_egress_ecmp_dump(__func__, ecmp_info);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_l3_egress_ecmp_create(int unit, opennsl_l3_egress_ecmp_t *ecmp, int intf_count, opennsl_if_t *intf_array) {
  int index;

  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  _opennsl_l3_egress_ecmp_dump(__func__, ecmp);
  LOG_DEBUG("%s: count    = %d", __func__, intf_count);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_l3_egress_ecmp_destroy(int unit, opennsl_l3_egress_ecmp_t *ecmp) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  _opennsl_l3_egress_ecmp_dump(__func__, ecmp);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_l3_egress_ecmp_get(int unit, opennsl_l3_egress_ecmp_t *ecmp, int intf_size, opennsl_if_t *intf_array, int *intf_count) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: size     = %d", __func__, intf_size);
  LOG_DEBUG("%s: array    = %p", __func__, intf_array);
  LOG_DEBUG("%s: count    = %p", __func__, intf_count);
  _opennsl_l3_egress_ecmp_dump(__func__, ecmp);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_l3_egress_ecmp_add(int unit, opennsl_l3_egress_ecmp_t *ecmp, opennsl_if_t intf) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: intf     = %d", __func__, intf);
  _opennsl_l3_egress_ecmp_dump(__func__, ecmp);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_l3_egress_ecmp_delete(int unit, opennsl_l3_egress_ecmp_t *ecmp, opennsl_if_t intf) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: intf     = %d", __func__, intf);
  _opennsl_l3_egress_ecmp_dump(__func__, ecmp);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_l3_egress_ecmp_find(int unit, int intf_count, opennsl_if_t *intf_array, opennsl_l3_egress_ecmp_t *ecmp) {
  int index;

  LOG_DEBUG("%s: unit       = %d", __func__, unit);
  LOG_DEBUG("%s: intf_count = %d", __func__, intf_count);
  LOG_DEBUG("%s: ecmp       = %p", __func__, ecmp);
  for (index = 0; index < intf_count; index++) {
    LOG_DEBUG("%s: iface[%d]  = %d", __func__, index, intf_array[index]);
  }

  return OPENNSL_E_UNAVAIL;
}

void opennsl_l3_ingress_t_init(opennsl_l3_ingress_t *ing_intf) {
  memset(ing_intf, 0, sizeof(opennsl_l3_ingress_t));
}

int opennsl_l3_ingress_create(int unit, opennsl_l3_ingress_t *ing_intf, opennsl_if_t *intf_id) {
  LOG_DEBUG("%s: unit       = %d", __func__, unit);
  LOG_DEBUG("%s: iface      = %d", __func__, *intf_id);
  _opennsl_l3_ingress_dump(__func__, ing_intf);

  if (ing_intf->flags & OPENNSL_L3_REPLACE) {
    _opennsl_l3_ingress_entry_t* const entry = _opennsl_l3_ingress_entry_get(*intf_id);
    if (entry == NULL)
      return OPENNSL_E_NOT_FOUND;

    entry->l3igr = *ing_intf;
    return OPENNSL_E_NONE;
  }

  if ((ing_intf->flags & OPENNSL_L3_WITH_ID) != 0) {
    _opennsl_l3_ingress_entry_t* const entry = _opennsl_l3_ingress_entry_get(*intf_id);
    if (entry != NULL)
      return OPENNSL_E_EXISTS;
  }

  int index;
  for (index = 0; index < _OPENNSL_L3_INGRESS_ENTRY_MAX; index++) {
    _opennsl_l3_ingress_entry_t* const entry = s_opennsl_l3_ingress_entries + index;
    if (entry->id == 0) {
      if ((ing_intf->flags & OPENNSL_L3_WITH_ID) != 0) {
	entry->id = *intf_id;
      } else {
	entry->id = index + 1;
	*intf_id  = index + 1;
      }

      entry->l3igr = *ing_intf;

      return OPENNSL_E_NONE;
    }
  }

  return OPENNSL_E_RESOURCE;
}

void opennsl_l3_host_t_init(opennsl_l3_host_t *host) {
  memset(host, 0, sizeof(opennsl_l3_host_t));
}

int opennsl_l3_host_find(int unit, opennsl_l3_host_t *info) {
  LOG_DEBUG("%s: unit       = %d", __func__, unit);
  _opennsl_l3_host_dump(__func__, info);

  if ((info->l3a_flags & OPENNSL_L3_IP6) == 0) {
    _opennsl_l3_host_entry_t* const entry = 
      _opennsl_l3_host4_entry_get(info->l3a_ip_addr, info->l3a_vrf);

    if (entry != NULL) {
      *info = entry->l3host;
      return OPENNSL_E_NONE;
    }

  } else {
    _opennsl_l3_host_entry_t* const entry =
      _opennsl_l3_host6_entry_get(info->l3a_ip6_addr, info->l3a_vrf);

    if (entry != NULL) {
      *info = entry->l3host;
      return OPENNSL_E_NONE;
    }
  }

  return OPENNSL_E_NOT_FOUND;
}

int opennsl_l3_host_add(int unit, opennsl_l3_host_t *info) {
  LOG_DEBUG("%s: unit       = %d", __func__, unit);
  _opennsl_l3_host_dump(__func__, info);

  if ((info->l3a_flags & OPENNSL_L3_IP6) == 0) {
    _opennsl_l3_host_entry_t* const entry =
      _opennsl_l3_host4_entry_get(info->l3a_ip_addr, info->l3a_vrf);

    if (entry != NULL)
      return OPENNSL_E_EXISTS;

    int index;
    for (index = 0; index < _OPENNSL_L3_HOST_ENTRY_MAX; index++) {
      _opennsl_l3_host_entry_t* const entry = s_opennsl_l3_host4_entries + index;

      if (entry->id == 0) {
	entry->id = (index + 1) * 2;
	entry->l3host = *info;

	return OPENNSL_E_NONE;
      }
    }
  } else {
    _opennsl_l3_host_entry_t* const entry =
      _opennsl_l3_host6_entry_get(info->l3a_ip6_addr, info->l3a_vrf);

    if (entry != NULL)
      return OPENNSL_E_EXISTS;

    int index;
    for (index = 0; index < _OPENNSL_L3_HOST_ENTRY_MAX; index++) {
      _opennsl_l3_host_entry_t* const entry = s_opennsl_l3_host6_entries + index;

      if (entry->id == 0) {
        entry->id = (index + 1) * 2 + 1;
        entry->l3host = *info;

        return OPENNSL_E_NONE;
      }
    }
  }

  return OPENNSL_E_RESOURCE;
}

int opennsl_l3_host_delete(int unit, opennsl_l3_host_t *info) {
  LOG_DEBUG("%s: unit       = %d", __func__, unit);
  _opennsl_l3_host_dump(__func__, info);

  if ((info->l3a_flags & OPENNSL_L3_IP6) == 0) {
    _opennsl_l3_host_entry_t* const entry =
      _opennsl_l3_host4_entry_get(info->l3a_ip_addr, info->l3a_vrf);

    if (entry == NULL)
      return OPENNSL_E_NOT_FOUND;

    entry->id = 0;

  } else {
    _opennsl_l3_host_entry_t* const entry =
      _opennsl_l3_host6_entry_get(info->l3a_ip6_addr, info->l3a_vrf);

    if (entry == NULL)
      return OPENNSL_E_NOT_FOUND;

    entry->id = 0;

  }

  return OPENNSL_E_NONE;
}

int opennsl_l3_host_delete_by_interface(int unit, opennsl_l3_host_t *info) {
  LOG_DEBUG("%s: unit       = %d", __func__, unit);
  _opennsl_l3_host_dump(__func__, info);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_l3_host_delete_all(int unit, opennsl_l3_host_t *info) {
  LOG_DEBUG("%s: unit       = %d", __func__, unit);
  _opennsl_l3_host_dump(__func__, info);

  int index;
  for (index = 0; index < _OPENNSL_L3_HOST_ENTRY_MAX; index++) {
    s_opennsl_l3_host4_entries[index].id = 0;
    s_opennsl_l3_host6_entries[index].id = 0;
  }

  return OPENNSL_E_NONE;
}

int opennsl_l3_host_traverse(int unit, uint32 flags, uint32 start, uint32 end, opennsl_l3_host_traverse_cb cb, void *user_data) {
  LOG_DEBUG("%s: unit       = %d", __func__, unit);
  LOG_DEBUG("%s: flags      = %08x", __func__, flags);
  LOG_DEBUG("%s: start      = %u", __func__, start);
  LOG_DEBUG("%s: end        = %u", __func__, end);
  LOG_DEBUG("%s: callback   = %p", __func__, cb);
  LOG_DEBUG("%s: user_data  = %p", __func__, user_data);

  int index;
  int count = 0;

  if ((flags & OPENNSL_L3_IP6) == 0) {
    for (index = 0; index < _OPENNSL_L3_HOST_ENTRY_MAX; index++) {
      _opennsl_l3_host_entry_t* const entry = s_opennsl_l3_host4_entries + index;
      if (entry->id == 0)
	continue;

      count ++;

      if ((count <= start) || (count > end))
	continue;

      const int rc = cb(unit, count-1, &entry->l3host, user_data);
      if (rc != OPENNSL_E_NONE)
	return rc;
    }
  } else {
    for (index = 0; index < _OPENNSL_L3_HOST_ENTRY_MAX; index++) {
      _opennsl_l3_host_entry_t* const entry = s_opennsl_l3_host6_entries + index;
      if (entry->id == 0)
	continue;

      count ++;

      if ((count <= start) || (count > end))
	continue;

      const int rc = cb(unit, count-1, &entry->l3host, user_data);
      if (rc != OPENNSL_E_NONE)
	return rc;
    }
  }

  return OPENNSL_E_NONE;
}

void opennsl_l3_route_t_init(opennsl_l3_route_t *route) {
  memset(route, 0, sizeof(opennsl_l3_route_t));
}

int opennsl_l3_route_add(int unit, opennsl_l3_route_t *info) {
  LOG_DEBUG("%s: unit       = %d", __func__, unit);
  _opennsl_l3_route_dump(__func__, info);

  if ((info->l3a_flags & OPENNSL_L3_IP6) == 0) {
    _opennsl_l3_route_entry_t* const entry =
      _opennsl_l3_route4_entry_get(info->l3a_subnet, info->l3a_ip_mask, info->l3a_vrf);

    if (entry != NULL)
      return OPENNSL_E_EXISTS;

    int index;
    for (index = 0; index < _OPENNSL_L3_ROUTE_ENTRY_MAX; index++) {
      _opennsl_l3_route_entry_t* const entry = s_opennsl_l3_route4_entries + index;

      if (entry->id == 0) {
        entry->id = (index + 1) * 2;
        entry->l3route = *info;

	return OPENNSL_E_NONE;
      }
    }
  } else {
    _opennsl_l3_route_entry_t* const entry =
      _opennsl_l3_route6_entry_get(info->l3a_ip6_net, info->l3a_ip6_mask, info->l3a_vrf);

    if (entry != NULL)
      return OPENNSL_E_EXISTS;

    int index;
    for (index = 0; index < _OPENNSL_L3_ROUTE_ENTRY_MAX; index++) {
      _opennsl_l3_route_entry_t* const entry = s_opennsl_l3_route6_entries + index;

      if (entry->id == 0) {
        entry->id = (index + 1) * 2 + 1;
        entry->l3route = *info;

	return OPENNSL_E_NONE;
      }
    }
  }

  return OPENNSL_E_RESOURCE;
}

int opennsl_l3_route_delete(int unit, opennsl_l3_route_t *info) {
  LOG_DEBUG("%s: unit       = %d", __func__, unit);
  _opennsl_l3_route_dump(__func__, info);

  return OPENNSL_E_NONE;
}

int opennsl_l3_route_delete_by_interface(int unit, opennsl_l3_route_t *info) {
  LOG_DEBUG("%s: unit       = %d", __func__, unit);
  _opennsl_l3_route_dump(__func__, info);

  return OPENNSL_E_NONE;
}

int opennsl_l3_route_delete_all(int unit, opennsl_l3_route_t *info) {
  LOG_DEBUG("%s: unit       = %d", __func__, unit);
  _opennsl_l3_route_dump(__func__, info);

  if ((info->l3a_flags & OPENNSL_L3_IP6) == 0) {
    _opennsl_l3_route_entry_t* const entry =
      _opennsl_l3_route4_entry_get(info->l3a_subnet, info->l3a_ip_mask, info->l3a_vrf);

    if (entry == NULL)
      return OPENNSL_E_NOT_FOUND;

    entry->id = 0;

  } else {
    _opennsl_l3_route_entry_t* const entry =
      _opennsl_l3_route6_entry_get(info->l3a_ip6_net, info->l3a_ip6_mask, info->l3a_vrf);

    if (entry == NULL)
      return OPENNSL_E_NOT_FOUND;

    entry->id = 0;
  }

  return OPENNSL_E_NONE;
}

int opennsl_l3_route_get(int unit, opennsl_l3_route_t *info) {
  LOG_DEBUG("%s: unit       = %d", __func__, unit);
  _opennsl_l3_route_dump(__func__, info);

  if ((info->l3a_flags & OPENNSL_L3_IP6) == 0) {
    _opennsl_l3_route_entry_t* const entry = 
      _opennsl_l3_route4_entry_get(info->l3a_subnet, info->l3a_ip_mask, info->l3a_vrf);

    if (entry != NULL) {
      *info = entry->l3route;
      return OPENNSL_E_NONE;
    }

  } else {
    _opennsl_l3_route_entry_t* const entry =
      _opennsl_l3_route6_entry_get(info->l3a_ip6_net, info->l3a_ip6_mask, info->l3a_vrf);

    if (entry != NULL) {
      *info = entry->l3route;
      return OPENNSL_E_NONE;
    }
  }

  return OPENNSL_E_NOT_FOUND;
}

int opennsl_l3_route_multipath_get(int unit, opennsl_l3_route_t *the_route, opennsl_l3_route_t *path_array, int max_path, int *path_count) {
  LOG_DEBUG("%s: unit       = %d", __func__, unit);
  LOG_DEBUG("%s: max_path   = %d", __func__, max_path);
  LOG_DEBUG("%s: path_count = %p", __func__, path_count);
  _opennsl_l3_route_dump(__func__, the_route);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_l3_route_traverse(int unit, uint32 flags, uint32 start, uint32 end, opennsl_l3_route_traverse_cb cb, void *user_data) {
  LOG_DEBUG("%s: unit       = %d", __func__, unit);
  LOG_DEBUG("%s: flags      = %08x", __func__, flags);
  LOG_DEBUG("%s: start      = %u", __func__, start);
  LOG_DEBUG("%s: end        = %u", __func__, end);
  LOG_DEBUG("%s: callback   = %p", __func__, cb);
  LOG_DEBUG("%s: user_data  = %p", __func__, user_data);

  int index;
  int count = 0;

  if ((flags & OPENNSL_L3_IP6) == 0) {
    for (index = 0; index < _OPENNSL_L3_ROUTE_ENTRY_MAX; index++) {
      _opennsl_l3_route_entry_t* const entry = s_opennsl_l3_route4_entries + index;
      if (entry->id == 0)
	continue;

      count ++;

      if ((count <= start) || (count > end))
	continue;

      const int rc = cb(unit, count-1, &entry->l3route, user_data);
      if (rc != OPENNSL_E_NONE)
	return rc;
    }
  } else {
    for (index = 0; index < _OPENNSL_L3_ROUTE_ENTRY_MAX; index++) {
      _opennsl_l3_route_entry_t* const entry = s_opennsl_l3_route6_entries + index;
      if (entry->id == 0)
	continue;

      count ++;

      if ((count <= start) || (count > end))
	continue;

      const int rc = cb(unit, count-1, &entry->l3route, user_data);
      if (rc != OPENNSL_E_NONE)
	return rc;
    }
  }

  return OPENNSL_E_NONE;
}

int opennsl_l3_route_max_ecmp_set(int unit, int max) {
  LOG_DEBUG("%s: unit       = %d", __func__, unit);
  LOG_DEBUG("%s: max        = %d", __func__, max);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_l3_route_max_ecmp_get(int unit, int *max) {
  LOG_DEBUG("%s: unit       = %d", __func__, unit);
  LOG_DEBUG("%s: max        = %p", __func__, max);

  *max = 10;
  return OPENNSL_E_UNAVAIL;
}

void opennsl_l3_info_t_init(opennsl_l3_info_t *info) {
  LOG_DEBUG("%s", __func__);

  memset(info, 0, sizeof(opennsl_l3_info_t));
}

int opennsl_l3_vrrp_add(int unit, opennsl_vlan_t vlan, uint32 vrid) {
  LOG_DEBUG("%s: unit       = %d", __func__, unit);
  LOG_DEBUG("%s: vid        = %hu", __func__, vlan);
  LOG_DEBUG("%s: vrid       = %u", __func__, vrid);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_l3_vrrp_delete(int unit, opennsl_vlan_t vlan, uint32 vrid) {
  LOG_DEBUG("%s: unit       = %d", __func__, unit);
  LOG_DEBUG("%s: vid        = %hu", __func__, vlan);
  LOG_DEBUG("%s: vrid       = %u", __func__, vrid);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_l3_vrrp_delete_all(int unit, opennsl_vlan_t vlan) {
  LOG_DEBUG("%s: unit       = %d", __func__, unit);
  LOG_DEBUG("%s: vid        = %hu", __func__, vlan);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_l3_vrrp_get(int unit, opennsl_vlan_t vlan, int alloc_size, int *vrid_array, int *count) {
  LOG_DEBUG("%s: unit       = %d", __func__, unit);
  LOG_DEBUG("%s: vid        = %hu", __func__, vlan);
  LOG_DEBUG("%s: size       = %d", __func__, alloc_size);
  LOG_DEBUG("%s: array      = %p", __func__, vrid_array);
  LOG_DEBUG("%s: count      = %p", __func__, count);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_l3_egress_ecmp_traverse(int unit, opennsl_l3_egress_ecmp_traverse_cb trav_fn, void* user_data) {
  LOG_DEBUG("%s: unit       = %d", __func__, unit);
  LOG_DEBUG("%s: callback   = %p", __func__, trav_fn);
  LOG_DEBUG("%s: data       = %p", __func__, user_data);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_l3_egress_stat_counter_get(int unit, opennsl_if_t intf_id, opennsl_l3_stat_t stat, uint32 num_entries, uint32 *counter_indexes, opennsl_stat_value_t *counter_values) {
  LOG_DEBUG("%s: unit       = %d", __func__, unit);
  LOG_DEBUG("%s: iface      = %d", __func__, intf_id);
  LOG_DEBUG("%s: start      = %u", __func__, stat);
  LOG_DEBUG("%s: size       = %d", __func__, num_entries);
  LOG_DEBUG("%s: values     = %p", __func__, counter_values);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_l3_ingress_stat_counter_get(int unit, opennsl_if_t intf_id, opennsl_l3_stat_t stat, uint32 num_entries, uint32 *counter_indexes, opennsl_stat_value_t *counter_values) {
  LOG_DEBUG("%s: unit       = %d", __func__, unit);
  LOG_DEBUG("%s: iface      = %d", __func__, intf_id);
  LOG_DEBUG("%s: start      = %u", __func__, stat);
  LOG_DEBUG("%s: size       = %d", __func__, num_entries);
  LOG_DEBUG("%s: values     = %p", __func__, counter_values);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_l3_egress_stat_attach(int unit, opennsl_if_t intf_id, uint32 stat_counter_id) {
  LOG_DEBUG("%s: unit       = %d", __func__, unit);
  LOG_DEBUG("%s: iface      = %d", __func__, intf_id);
  LOG_DEBUG("%s: counter_id = %u", __func__, stat_counter_id);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_l3_egress_stat_detach(int unit, opennsl_if_t intf_id) {
  LOG_DEBUG("%s: unit       = %d", __func__, unit);
  LOG_DEBUG("%s: iface      = %d", __func__, intf_id);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_l3_ingress_stat_attach(int unit, opennsl_if_t intf_id, uint32 stat_counter_id) {
  LOG_DEBUG("%s: unit       = %d", __func__, unit);
  LOG_DEBUG("%s: iface      = %d", __func__, intf_id);
  LOG_DEBUG("%s: counter_id = %u", __func__, stat_counter_id);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_l3_ingress_stat_detach(int unit, opennsl_if_t intf_id) {
  LOG_DEBUG("%s: unit       = %d", __func__, unit);
  LOG_DEBUG("%s: iface      = %d", __func__, intf_id);

  return OPENNSL_E_UNAVAIL;
}
