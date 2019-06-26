// -*- coding: utf-8 -*-

#include <stdlib.h>
#include <string.h>
#include <pthread.h>
#include <sys/time.h>
#include <opennsl/error.h>
#include <opennsl/l2.h>
#include "libopennsl.h"
#include "logger.h"

typedef struct _opennsl_l2_addr_entry {
  int id;
  opennsl_l2_addr_t addr;
} _opennsl_l2_addr_entry_t;

#define L2_ADDR_ENTRY_MAX (256)
static _opennsl_l2_addr_entry_t s_opennsl_l2_addr_entriies[L2_ADDR_ENTRY_MAX] = {0};

static _opennsl_l2_addr_entry_t* _opennsl_l2_addr_entry_get(const opennsl_mac_t mac, opennsl_vlan_t vid) {
  int index;
  for (index = 0; index < L2_ADDR_ENTRY_MAX; index++) {
    _opennsl_l2_addr_entry_t* const entry = s_opennsl_l2_addr_entriies + index;

    if (entry->id == 0)
      continue;

    if ((vid == entry->addr.vid) && (memcmp(mac, entry->addr.mac, sizeof(opennsl_mac_t)) == 0))
	return entry;
  }

  return NULL;
}

static int s_age_seconds = 0;

typedef struct _opennsl_l2_station_entry {
  int station_id;
  opennsl_l2_station_t station;
} _opennsl_l2_station_entry_t;

#define L2_STATION_ENTRY_MAX (256)
static _opennsl_l2_station_entry_t s_opennsl_l2_station_entries[L2_STATION_ENTRY_MAX] = {0};

static _opennsl_l2_station_entry_t* _opennsl_l2_station_entry_get(const opennsl_mac_t dst_mac, const opennsl_mac_t dst_mac_mask, opennsl_vlan_t vlan) {

  int index;
  for (index = 0; index < L2_STATION_ENTRY_MAX; index++) {
    _opennsl_l2_station_entry_t* const entry = s_opennsl_l2_station_entries + index;

    if (entry->station_id == 0)
      continue;

    if ((vlan == entry->station.vlan) &&
	(memcmp(dst_mac, entry->station.dst_mac, sizeof(opennsl_mac_t)) == 0) &&
	(memcmp(dst_mac_mask, entry->station.dst_mac_mask, sizeof(opennsl_mac_t)) == 0))
      return entry;
  }

  return NULL;
}

static _opennsl_l2_station_entry_t* _opennsl_l2_station_entry_get_by_id(int station_id) {
  int index;
  for (index = 0; index < L2_STATION_ENTRY_MAX; index++) {
    _opennsl_l2_station_entry_t* const entry = s_opennsl_l2_station_entries + index;

    if (entry->station_id == 0)
      continue;

    if (entry->station_id == station_id)
      return entry;
  }

  return NULL;
}

void opennsl_l2_addr_t_init(opennsl_l2_addr_t *l2addr, const opennsl_mac_t mac_addr, opennsl_vlan_t vid) {
  char mac[32];
  strmac(mac, mac_addr);
  LOG_DEBUG("%s: mac = %s", __func__, mac);
  LOG_DEBUG("%s: vid = %d", __func__, vid);

  memset(l2addr, sizeof(opennsl_l2_addr_t), 0);
  memcpy(l2addr->mac, mac_addr, 6);
  l2addr->vid = vid;
}

int opennsl_l2_addr_add(int unit, opennsl_l2_addr_t *l2addr) {
  LOG_DEBUG("%s: unit= %d", __func__, unit);
  _opennsl_l2_addr_dump(__func__, l2addr);

  int index;
  for (index = 0; index < L2_ADDR_ENTRY_MAX; index++) {
    _opennsl_l2_addr_entry_t* const entry = s_opennsl_l2_addr_entriies + index;

    if (entry->id == 0) {
      memset(entry, 0, sizeof(_opennsl_l2_addr_entry_t));
      entry->id = index + 1;
      entry->addr.vid = l2addr->vid;
      memcpy(entry->addr.mac, l2addr->mac, sizeof(opennsl_mac_t));

      return OPENNSL_E_NONE;
    }
  }

  return OPENNSL_E_RESOURCE;
}

int opennsl_l2_addr_delete(int unit, opennsl_mac_t mac_addr, opennsl_vlan_t vid) {
  char mac[32];
  strmac(mac, mac_addr);
  LOG_DEBUG("%s: unit= %d", __func__, unit);
  LOG_DEBUG("%s: mac = %s", __func__, mac);
  LOG_DEBUG("%s: vid = %d", __func__, vid);

  _opennsl_l2_addr_entry_t* const entry = _opennsl_l2_addr_entry_get(mac_addr, vid);
  if (entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  entry->id = 0;
  return OPENNSL_E_NONE;
}

int opennsl_l2_addr_delete_by_port(int unit, opennsl_module_t mod, opennsl_port_t port, uint32 flags) {
  LOG_DEBUG("%s: unit  = %d", __func__, unit);
  LOG_DEBUG("%s: module= %d", __func__, mod);
  LOG_DEBUG("%s: port  = %d", __func__, port);
  LOG_DEBUG("%s: flags = %08x", __func__, flags);
  return OPENNSL_E_UNAVAIL;
}

int opennsl_l2_addr_delete_by_mac(int unit, opennsl_mac_t mac_addr, uint32 flags) {
  char mac[32];
  strmac(mac, mac_addr);
  LOG_DEBUG("%s: unit  = %d", __func__, unit);
  LOG_DEBUG("%s: mac   = %s", __func__, mac);
  LOG_DEBUG("%s: flags = %08x", __func__, flags);

  int index;
  for (index = 0; index < L2_ADDR_ENTRY_MAX; index++) {
    _opennsl_l2_addr_entry_t* const entry = s_opennsl_l2_addr_entriies + index;

    if (entry->id == 0)
      continue;

    if (memcmp(mac_addr, entry->addr.mac, sizeof(opennsl_mac_t)) == 0)
      entry->id = 0;
  }

  return OPENNSL_E_NONE;
}

int opennsl_l2_addr_delete_by_vlan(int unit, opennsl_vlan_t vid, uint32 flags) {
  LOG_DEBUG("%s: unit  = %d", __func__, unit);
  LOG_DEBUG("%s: vid   = %d", __func__, vid);
  LOG_DEBUG("%s: flags = %08x", __func__, flags);

  int index;
  for (index = 0; index < L2_ADDR_ENTRY_MAX; index++) {
    _opennsl_l2_addr_entry_t* const entry = s_opennsl_l2_addr_entriies + index;

    if (entry->id == 0)
      continue;

    if (entry->addr.vid == vid)
      entry->id = 0;
  }

  return OPENNSL_E_UNAVAIL;
}

int opennsl_l2_addr_delete_by_trunk(int unit, opennsl_trunk_t tid, uint32 flags) {
  LOG_DEBUG("%s: unit  = %d", __func__, unit);
  LOG_DEBUG("%s: trunk = %d", __func__, tid);
  LOG_DEBUG("%s: flags = %08x", __func__, flags);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_l2_addr_delete_by_mac_port(int unit, opennsl_mac_t mac_addr, opennsl_module_t mod, opennsl_port_t port, uint32 flags) {
  char mac[32];
  strmac(mac, mac_addr);
  LOG_DEBUG("%s: unit  = %d", __func__, unit);
  LOG_DEBUG("%s: mac   = %s", __func__, mac);
  LOG_DEBUG("%s: module= %d", __func__, mod);
  LOG_DEBUG("%s: port  = %d", __func__, port);
  LOG_DEBUG("%s: flags = %08x", __func__, flags);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_l2_addr_delete_by_vlan_port(int unit, opennsl_vlan_t vid, opennsl_module_t mod, opennsl_port_t port, uint32 flags) {
  LOG_DEBUG("%s: unit  = %d", __func__, unit);
  LOG_DEBUG("%s: vid   = %d", __func__, vid);
  LOG_DEBUG("%s: module= %d", __func__, mod);
  LOG_DEBUG("%s: port  = %d", __func__, port);
  LOG_DEBUG("%s: flags = %08x", __func__, flags);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_l2_addr_delete_by_vlan_trunk(int unit, opennsl_vlan_t vid, opennsl_trunk_t tid, uint32 flags) {
  LOG_DEBUG("%s: unit  = %d", __func__, unit);
  LOG_DEBUG("%s: vid   = %d", __func__, vid);
  LOG_DEBUG("%s: trunk = %d", __func__, tid);
  LOG_DEBUG("%s: flags = %08x", __func__, flags);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_l2_addr_get(int unit, opennsl_mac_t mac_addr, opennsl_vlan_t vid, opennsl_l2_addr_t *l2addr) {
  char mac[32];
  strmac(mac, mac_addr);
  LOG_DEBUG("%s: unit  = %d", __func__, unit);
  LOG_DEBUG("%s: mac   = %s", __func__, mac);
  LOG_DEBUG("%s: vid   = %d", __func__, vid);
  LOG_DEBUG("%s: l2addr= %p", __func__, l2addr);

  _opennsl_l2_addr_entry_t* const entry = _opennsl_l2_addr_entry_get(mac_addr, vid);
  if (entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  *l2addr = entry->addr;
  return OPENNSL_E_NONE;
}

typedef struct _opennsl_l2_addr_info {
  pthread_t tid;
  int status;
  int unit;
  void* userdata;
  opennsl_l2_addr_callback_t callback;
  uint32 data_num;
  long wait_sec;

  pthread_mutex_t mtx;
  pthread_cond_t  cond;
} _opennsl_l2_addr_info_t;

static opennsl_l2_addr_t* _opennsl_l2_addr_data_create(uint32 num) {
  opennsl_l2_addr_t* datas = (opennsl_l2_addr_t*)malloc(sizeof(opennsl_l2_addr_t) * num);

  opennsl_mac_t mac;
  mac[0] = 0x00;
  mac[1] = 0x11;

  uint32 index;
  for (index = 0; index < num; index++) {
    opennsl_l2_addr_t* data = datas + index;

    mac[2] = (uint8)((index & 0xff000000) >> 24);
    mac[3] = (uint8)((index & 0x00ff0000) >> 16);
    mac[4] = (uint8)((index & 0x0000ff00) >> 8);
    mac[5] = (uint8)((index & 0x000000ff) >> 0);

    data->port = index + 1;
    data->vid = index + 11;
    memcpy(data->mac, mac, 6);
  }

  return datas;
}

static void _opennsl_l2_addr_data_destroy(opennsl_l2_addr_t* datas) {
  if (datas != NULL)
    free(datas);
}

static int _opennsl_l2_addr_oper_next(int oper) {
  switch(oper) {
  case OPENNSL_L2_CALLBACK_ADD:
    return OPENNSL_L2_CALLBACK_DELETE;

  case OPENNSL_L2_CALLBACK_DELETE:
    return OPENNSL_L2_CALLBACK_LEARN_EVENT;

  default:
    return OPENNSL_L2_CALLBACK_ADD;
  }
}

static void* _opennsl_l2_addr_main(void* arg) {

  _opennsl_l2_addr_info_t* const info = (_opennsl_l2_addr_info_t*)arg;
  opennsl_l2_addr_t* const datas = _opennsl_l2_addr_data_create(info->data_num);

  int oper = OPENNSL_L2_CALLBACK_ADD;

  pthread_mutex_lock(&info->mtx);
  while(1) {
    printf("_opennsl_l2_addr_main: slepp...%ld\n", info->wait_sec);

    struct timeval now;
    struct timespec timeout;
    gettimeofday(&now, NULL);
    timeout.tv_sec = now.tv_sec + info->wait_sec;
    timeout.tv_nsec = 0;
    pthread_cond_timedwait(&info->cond, &info->mtx, &timeout);

    if (info->status == 0) {
      printf("_opennsl_l2_addr_main: exit\n");
      break;
    }

    printf("_opennsl_l2_addr_main: callback #%d ope:%d\n", info->data_num, oper);

    int index;
    for (index = 0; index < info->data_num; index++) {
      info->callback(info->unit, datas + index, oper, info->userdata);
    }

    oper = _opennsl_l2_addr_oper_next(oper);
  }
  pthread_mutex_unlock(&info->mtx);

  _opennsl_l2_addr_data_destroy(datas);
}

static _opennsl_l2_addr_info_t s_opennsl_l2_addr_info = {0, 0, 0, NULL, NULL};

static void _opennsl_l2_addr_callback_env(uint32* data_num, long* wait_sec) {
  const char* env;

  if ((env = getenv("OPENNSL_STUB_L2ADDR_NUM")) != NULL) {
    *data_num = atoi(env);
  }

  if ((*data_num) < 1)
    *data_num = 1;

  if ((env = getenv("OPENNSL_STUB_L2ADDR_WAIT_SEC")) != NULL) {
    *wait_sec = atoi(env);
  }

  if ((*wait_sec) < 1)
    *wait_sec = 3600;
}

int opennsl_l2_addr_register(int unit, opennsl_l2_addr_callback_t callback, void *userdata) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: cakkback = %p", __func__, callback);
  LOG_DEBUG("%s: userdata = %p", __func__, userdata);

  if (s_opennsl_l2_addr_info.status != 0)
    return OPENNSL_E_INTERNAL;

  s_opennsl_l2_addr_info.status = 1;
  s_opennsl_l2_addr_info.unit = unit;
  s_opennsl_l2_addr_info.userdata = userdata;
  s_opennsl_l2_addr_info.callback = callback;
  _opennsl_l2_addr_callback_env(&s_opennsl_l2_addr_info.data_num,
				&s_opennsl_l2_addr_info.wait_sec);

  pthread_cond_init(&s_opennsl_l2_addr_info.cond, NULL);
  pthread_mutex_init(&s_opennsl_l2_addr_info.mtx, NULL);

  int rc = pthread_create(&s_opennsl_l2_addr_info.tid,
			  NULL,
			  _opennsl_l2_addr_main,
			  &s_opennsl_l2_addr_info);
  if (rc != 0)
    return OPENNSL_E_INTERNAL;

  return OPENNSL_E_NONE;
}

int opennsl_l2_addr_unregister(int unit, opennsl_l2_addr_callback_t callback, void *userdata) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: cakkback = %p", __func__, callback);
  LOG_DEBUG("%s: userdata = %p", __func__, userdata);

  if (s_opennsl_l2_addr_info.status == 0)
    return OPENNSL_E_INTERNAL;

  s_opennsl_l2_addr_info.status = 0;
  pthread_cond_signal(&s_opennsl_l2_addr_info.cond);

  s_opennsl_l2_addr_info.tid = 0;
  
  return OPENNSL_E_NONE;
}

int opennsl_l2_age_timer_set(int unit, int age_seconds) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: age      = %d", __func__, age_seconds);
  if (age_seconds < 0) {
    return OPENNSL_E_UNAVAIL;
  }
  s_age_seconds = age_seconds;
  return OPENNSL_E_NONE;
}

int opennsl_l2_age_timer_get(int unit, int *age_seconds) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: age      = %p", __func__, age_seconds);
  if (s_age_seconds < 0) {
    return OPENNSL_E_UNAVAIL;
  }
  *age_seconds = s_age_seconds;
  return OPENNSL_E_NONE;
}

int opennsl_l2_addr_freeze(int unit) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  return OPENNSL_E_UNAVAIL;
}

int opennsl_l2_addr_thaw(int unit) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  return OPENNSL_E_UNAVAIL;
}

int opennsl_l2_cache_init(int unit) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  return OPENNSL_E_UNAVAIL;
}

void opennsl_l2_cache_addr_t_init(opennsl_l2_cache_addr_t *addr) {
  LOG_DEBUG("%s: addr     = %p", __func__, addr);
}

int opennsl_l2_cache_size_get(int unit, int *size) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: size     = %p", __func__, size);
  return OPENNSL_E_UNAVAIL;
}

int opennsl_l2_cache_set(int unit, int index, opennsl_l2_cache_addr_t *addr, int *index_used) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: index    = %d", __func__, index);
  _opennsl_l2_cache_addr_dump(__func__, addr);
  LOG_DEBUG("%s: used     = %p", __func__, index_used);
  return OPENNSL_E_UNAVAIL;
}

int opennsl_l2_cache_get(int unit, int index, opennsl_l2_cache_addr_t *addr) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: index    = %d", __func__, index);
  LOG_DEBUG("%s: addr     = %p", __func__, addr);
  return OPENNSL_E_UNAVAIL;
}

int opennsl_l2_cache_delete(int unit, int index) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: index    = %d", __func__, index);
  return OPENNSL_E_UNAVAIL;
}

int opennsl_l2_cache_delete_all(int unit) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  return OPENNSL_E_UNAVAIL;
}

int opennsl_l2_tunnel_add(int unit, opennsl_mac_t mac_addr, opennsl_vlan_t vid) {
  char mac[32];
  strmac(mac, mac_addr);

  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: mac      = %s", __func__, mac);
  LOG_DEBUG("%s: vid      = %d", __func__, vid);

  return OPENNSL_E_UNAVAIL;
}

int opennsl_l2_traverse(int unit, opennsl_l2_traverse_cb callback, void *user_data) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: callback = %p", __func__, callback);
  LOG_DEBUG("%s: user_data= %p", __func__, user_data);

  int index;
  for (index = 0; index < L2_ADDR_ENTRY_MAX; index++) {
    _opennsl_l2_addr_entry_t* const entry = s_opennsl_l2_addr_entriies + index;

    if (entry->id == 0)
      continue;

    callback(unit, &entry->addr, user_data);
  }

  return OPENNSL_E_NONE;
}

int opennsl_l2_replace(int unit, uint32 flags, opennsl_l2_addr_t *match_addr, opennsl_module_t new_module, opennsl_port_t new_port, opennsl_trunk_t new_trunk) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: flags    = %d", __func__, flags);
  _opennsl_l2_addr_dump(__func__, match_addr);
  LOG_DEBUG("%s: module   = %d", __func__, new_module);
  LOG_DEBUG("%s: port     = %d", __func__, new_port);
  LOG_DEBUG("%s: trunk    = %d", __func__, new_trunk);
  return OPENNSL_E_UNAVAIL;
}

void _opennsl_l2_station_dump(const char* name, const opennsl_l2_station_t* s) {
  char mac[32];
  strmac(mac, s->dst_mac);
  char mac_mask[32];
  strmac(mac_mask, s->dst_mac_mask);
  
  LOG_DEBUG("%s: l2_station.flags   = %08x", name, s->flags);
  LOG_DEBUG("%s: l2_station.dest_mac = %s/%s", name, mac, mac_mask);
  LOG_DEBUG("%s: l2_station.vid      = %d/%d", name, s->vlan, s->vlan_mask);
  LOG_DEBUG("%s: l2_station.src_port = %d/%d", name, s->src_port, s->src_port_mask);
}

void opennsl_l2_station_t_init(opennsl_l2_station_t *addr) {
  memset(addr, 0, sizeof(opennsl_l2_station_t));
}

int opennsl_l2_station_add(int unit, int *station_id, opennsl_l2_station_t *station) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: stationId= %p (%d)", __func__, station_id, *station_id);
  _opennsl_l2_station_dump(__func__, station);

  int index;
  for (index = 0; index < L2_STATION_ENTRY_MAX; index++) {
    _opennsl_l2_station_entry_t* const entry = s_opennsl_l2_station_entries + index;

    if (entry->station_id == 0) {
      if ((station->flags & OPENNSL_L2_STATION_WITH_ID) != 0) {
	if (*station_id == 0) {
	  return OPENNSL_E_PARAM;
	}
      } else {
	*station_id = index + 1;
      }

      entry->station_id = *station_id;

    } else {
      if ((station->flags & OPENNSL_L2_STATION_WITH_ID) != 0)
	if (entry->station_id == (*station_id))
	  return OPENNSL_E_EXISTS;
    }
  }
  
  return OPENNSL_E_NONE;
}

int opennsl_l2_station_delete(int unit, int station_id) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: stationId= %d", __func__, station_id);

  _opennsl_l2_station_entry_t* const entry = _opennsl_l2_station_entry_get_by_id(station_id);
  if (entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  entry->station_id = 0;

  return OPENNSL_E_NONE;
}

int opennsl_l2_station_get(int unit, int station_id, opennsl_l2_station_t *station) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: stationID= %d", __func__, station_id);
  LOG_DEBUG("%s: station  = %p", __func__, station);

  _opennsl_l2_station_entry_t* const entry = _opennsl_l2_station_entry_get_by_id(station_id);
  if (entry == NULL)
    return OPENNSL_E_NOT_FOUND;

  *station = entry->station;

  return OPENNSL_E_NONE;
}
