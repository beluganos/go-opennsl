// -*- coding: utf-8 -*-

#include <stdlib.h>
#include <string.h>
#include <opennsl/types.h>
#include <opennsl/error.h>
#include <opennsl/port.h>
#include "libopennsl.h"
#include "logger.h"

#define OPENNSL_STUB_PORT_NUM_MAX  (256)
#define OPENNSL_STUB_PORT_CTRL_MAX (512)
#define OPENNSL_STUB_GPORT_MASK    (0x1000)

typedef struct _opennsl_port_flood_block {
  int id;
  opennsl_port_t igr_port;
  opennsl_port_t egr_port;
  uint32 flags;
} _opennsl_port_flood_block_t;

#define _OPENNSL_PORT_FLOOD_BLOCK_MAX (256)
static _opennsl_port_flood_block_t s_opennsl_port_flood_block[_OPENNSL_PORT_FLOOD_BLOCK_MAX];

static _opennsl_port_flood_block_t* _opennsl_port_flood_block_get(opennsl_port_t igr, opennsl_port_t egr) {
  int index;
  for (index = 0; index < _OPENNSL_PORT_FLOOD_BLOCK_MAX; index++) {
    _opennsl_port_flood_block_t* const entry = s_opennsl_port_flood_block + index;
    if (entry->igr_port == igr && entry->egr_port == egr)
      return entry;
  }

  return NULL;
}

typedef struct _opennsl_ext_port_info {
  int status;
  int linkstatus;
  opennsl_port_info_t port_info;
  opennsl_vlan_t untagged_vlan;
  int port_controls[OPENNSL_STUB_PORT_CTRL_MAX];
  int port_stat_enable;
  int port_enable;
  opennsl_mac_t pause_addr;
  opennsl_color_t cfi_color[2];
  opennsl_port_class_t port_class[16];
  uint32 learn_flags;
  int dtag_mode;
  int speed;
  int duplex;
  int l3_mtu;
  int frame_max;
  int autoneg;

  uint64 stat_val[opennsl_spl_snmpValCount];
} _opennsl_port_info_t;

static int s_opennsl_port_num = 5;
static _opennsl_port_info_t s_opennsl_port_infos[OPENNSL_STUB_PORT_NUM_MAX];

static _opennsl_port_info_t* _opennsl_port_info_get(opennsl_port_t port) {
  if (port >= OPENNSL_STUB_PORT_NUM_MAX)
    return NULL;

  return s_opennsl_port_infos + port;
}

static uint64* _opennsl_port_stat_val_get(opennsl_port_t port, opennsl_stat_val_t t) {
  _opennsl_port_info_t* const info = _opennsl_port_info_get(port);
  if (info == NULL)
    return NULL;

  if (t >= opennsl_spl_snmpValCount)
    return NULL;

  return info->stat_val + t;
}

void _opennsl_port_num_set(int n) {
  if (n < 0) {
    const char* env = getenv("OPENNSL_STUB_PORT_NUM");
    if (env != NULL) {
      n = atoi(env);
    }
  }

  if (n > 0 && n < OPENNSL_STUB_PORT_NUM_MAX) {
    s_opennsl_port_num = n;
  }

  LOG_DEBUG("%s: port_num = %d", __func__, s_opennsl_port_num);
}

int _opennsl_port_num_get(void) {
  LOG_DEBUG("%s: port_num = %d", __func__, s_opennsl_port_num);
  return s_opennsl_port_num;
}

int _opennsl_stat_set(opennsl_port_t port, opennsl_stat_val_t type, uint64 val) {
  uint64* sv = _opennsl_port_stat_val_get(port, type);
  if (sv == NULL)
    return OPENNSL_E_PARAM;

  *sv = val;
  return OPENNSL_E_NONE;
}

int _opennsl_stat_multi_set(opennsl_port_t port, int nstat, opennsl_stat_val_t* type_arr, uint64* val_arr) {

  int index;
  for (index = 0; index < nstat; index++) {
    uint64* const sv = _opennsl_port_stat_val_get(port, type_arr[index]);
    if (sv == NULL)
      return OPENNSL_E_PARAM;

    *sv = val_arr[index];
  }

  return OPENNSL_E_NONE;
}

static void _opennsl_port_init() {

  memset(s_opennsl_port_infos, 0, sizeof(s_opennsl_port_infos));

  int index;
  for (index = 0; index < s_opennsl_port_num; index++) {
    _opennsl_port_info_t* const port_info = s_opennsl_port_infos + index;
    port_info->status = 1;
    port_info->linkstatus = 1;
    opennsl_port_info_t_init(&port_info->port_info);
    port_info->port_info.linkstatus = 1;
  }

  for (; index < OPENNSL_STUB_PORT_NUM_MAX; index++) {
    _opennsl_port_info_t* const port_info = s_opennsl_port_infos + index;
    port_info->status = 0;
    port_info->linkstatus = 1;
  }
}

//
// opennsl_port_init
//
int opennsl_port_init(int unit) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);

  _opennsl_port_num_set(-1);
  _opennsl_port_init();
  return OPENNSL_E_NONE;
}

//
// opennsl_port_clear
//
int opennsl_port_clear(int unit) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);

  _opennsl_port_init();
  return OPENNSL_E_NONE;
}

//
// opennsl_port_config_t_init
//
void opennsl_port_config_t_init(opennsl_port_config_t* pconfig) {
  LOG_DEBUG("%s: config   = %p", __func__, pconfig);
  memset(pconfig, 0, sizeof(opennsl_port_config_t));
}

//
// opennsl_port_config_get
//
int opennsl_port_config_get(int unit, opennsl_port_config_t* config) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: config   = %p", __func__, config);

  int index;
  for (index = 0; index < OPENNSL_STUB_PORT_NUM_MAX; index++) {
    if (s_opennsl_port_infos[index].status != 0) {
      OPENNSL_PBMP_PORT_ADD(config->e, index);
    }
  }

  return OPENNSL_E_NONE;
}

//
// opennsl_port_info_t_init
//
void opennsl_port_info_t_init(opennsl_port_info_t* info) {
  LOG_DEBUG("%s: info     = %p", __func__, info);
  memset(info, 0, sizeof(opennsl_port_info_t));
}

//
// opennsl_port_ability_t_init
//
void opennsl_port_ability_t_init(opennsl_port_ability_t* ability) {
  LOG_DEBUG("%s: ability   = %p", __func__, ability);
  memset(ability, 0, sizeof(opennsl_port_ability_t));
}

//
// opennsl_port_selective_get
//
int opennsl_port_selective_get(int unit, opennsl_port_t port, opennsl_port_info_t *info) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: port     = %d", __func__, port);

  _opennsl_port_info_t* const _info = _opennsl_port_info_get(port);
  if (_info == NULL)
    return OPENNSL_E_PARAM;

  *info =_info->port_info;

  _opennsl_port_info_dump(__func__, info);
  return OPENNSL_E_NONE;
}

int opennsl_port_selective_set(int unit, opennsl_port_t port, opennsl_port_info_t *info) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  _opennsl_port_info_dump(__func__, info);

  _opennsl_port_info_t* const _info = _opennsl_port_info_get(port);
  if (_info == NULL)
    return OPENNSL_E_PARAM;

  _info->port_info = *info;
  _info->port_info.linkstatus = _info->linkstatus;

  return OPENNSL_E_NONE;
}

int opennsl_port_untagged_vlan_set(int unit, opennsl_port_t port, opennsl_vlan_t vid) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  LOG_DEBUG("%s: vid      = %hu", __func__, vid);

  _opennsl_port_info_t* const _info = _opennsl_port_info_get(port);
  if (_info == NULL)
    return OPENNSL_E_PARAM;

  _info->untagged_vlan = vid;
  return OPENNSL_E_NONE;
}

int opennsl_port_untagged_vlan_get(int unit, opennsl_port_t port, opennsl_vlan_t *vid_ptr) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  LOG_DEBUG("%s: vid      = %p", __func__, vid_ptr);

  _opennsl_port_info_t* const _info = _opennsl_port_info_get(port);
  if (_info == NULL)
    return OPENNSL_E_PARAM;

  *vid_ptr = _info->untagged_vlan;
  return OPENNSL_E_NONE;
}

int opennsl_port_control_set(int unit, opennsl_port_t port, opennsl_port_control_t type, int value)  {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  LOG_DEBUG("%s: type     = %d", __func__, type);
  LOG_DEBUG("%s: value    = %d", __func__, value);

  uint64* const sv = _opennsl_port_stat_val_get(port, type);
  if (sv == NULL)
    return OPENNSL_E_PARAM;

  *sv = value;
  return OPENNSL_E_NONE;
}

int opennsl_port_control_get(int unit, opennsl_port_t port, opennsl_port_control_t type, int *value) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  LOG_DEBUG("%s: type     = %d", __func__, type);
  LOG_DEBUG("%s: value    = %p", __func__, value);

  uint64* const sv = _opennsl_port_stat_val_get(port, type);
  if (sv == NULL)
    return OPENNSL_E_PARAM;

  *value = *sv;
  return OPENNSL_E_NONE;

}

int opennsl_port_gport_get(int unit, opennsl_port_t port, opennsl_gport_t *gport) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  LOG_DEBUG("%s: gport    = %p", __func__, gport);

  *gport = port | OPENNSL_STUB_GPORT_MASK;
  return OPENNSL_E_NONE;
}

int opennsl_port_local_get(int unit, opennsl_gport_t gport, opennsl_port_t *local_port) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: gport    = %d", __func__, gport);
  LOG_DEBUG("%s: port     = %p", __func__, local_port);

  *local_port = gport & (~OPENNSL_STUB_GPORT_MASK);
  return OPENNSL_E_NONE;
}

int opennsl_port_stat_enable_set(int unit, opennsl_gport_t gport, int enable) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: gport    = %d", __func__, gport);
  LOG_DEBUG("%s: enable   = %d", __func__, enable);

  opennsl_port_t port = 0;
  opennsl_port_local_get(unit, gport, &port);

  _opennsl_port_info_t* const _info = _opennsl_port_info_get(port);
  if (_info == NULL)
    return OPENNSL_E_PARAM;

  _info->port_stat_enable = enable;
  return OPENNSL_E_NONE;
}

int opennsl_port_enable_get(int unit, opennsl_port_t port, int* enable) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  LOG_DEBUG("%s: enable   = %p", __func__, enable);

  _opennsl_port_info_t* const _info = _opennsl_port_info_get(port);
  if (_info == NULL)
    return OPENNSL_E_PARAM;

  *enable = _info->port_enable;
  return OPENNSL_E_NONE;
}

int opennsl_port_enable_set(int unit, opennsl_port_t port, int enable) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  LOG_DEBUG("%s: enable   = %d", __func__, enable);

  _opennsl_port_info_t* const _info = _opennsl_port_info_get(port);
  if (_info == NULL)
    return OPENNSL_E_PARAM;

  _info->port_enable = enable;
  return OPENNSL_E_NONE;
}

int opennsl_port_link_status_get(int unit, opennsl_port_t port, int *status) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  LOG_DEBUG("%s: status   = %p", __func__, status);

  _opennsl_port_info_t* const _info = _opennsl_port_info_get(port);
  if (_info == NULL)
    return OPENNSL_E_PARAM;

  *status = _info->linkstatus;
  return OPENNSL_E_NONE;
}

int opennsl_port_autoneg_get(int unit, opennsl_port_t port, int* autoneg) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  LOG_DEBUG("%s: autoneg  = %p", __func__, autoneg);

  _opennsl_port_info_t* const _info = _opennsl_port_info_get(port);
  if (_info == NULL)
    return OPENNSL_E_PARAM;

  *autoneg = _info->autoneg;

  return OPENNSL_E_NONE;
}

int opennsl_port_autoneg_set(int unit, opennsl_port_t port, int autoneg) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  LOG_DEBUG("%s: autoneg  = %d", __func__, autoneg);

  _opennsl_port_info_t* const _info = _opennsl_port_info_get(port);
  if (_info == NULL)
    return OPENNSL_E_PARAM;

  _info->autoneg = autoneg;

  return OPENNSL_E_NONE;
}

int opennsl_port_flood_block_get(int unit, opennsl_port_t ingress_port, opennsl_port_t egress_port, uint32* flags) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: igr_port = %d", __func__, ingress_port);
  LOG_DEBUG("%s: egr_port = %d", __func__, egress_port);
  LOG_DEBUG("%s: flags    = %p", __func__, flags);

  _opennsl_port_flood_block_t* const flood_blk = _opennsl_port_flood_block_get(ingress_port, egress_port);
  if (flood_blk == NULL)
    return OPENNSL_E_NOT_FOUND;

  *flags = flood_blk->flags;

  return OPENNSL_E_NONE;
}

int opennsl_port_flood_block_set(int unit, opennsl_port_t ingress_port, opennsl_port_t egress_port, uint32 flags){
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: igr_port = %d", __func__, ingress_port);
  LOG_DEBUG("%s: egr_port = %d", __func__, egress_port);
  LOG_DEBUG("%s: flags    = %d", __func__, flags);

  _opennsl_port_flood_block_t* flood_blk = _opennsl_port_flood_block_get(ingress_port, egress_port);
  if (flood_blk == NULL) {
    int index;
    for (index = 0; index < _OPENNSL_PORT_FLOOD_BLOCK_MAX; index++) {
      _opennsl_port_flood_block_t* const entry = s_opennsl_port_flood_block + index;
      if (entry->id != 0)
	continue;

      entry->id = 1;
      entry->igr_port = ingress_port;
      entry->egr_port = egress_port;
      flood_blk = entry;
      break;
    }

    if (flood_blk == NULL)
      return OPENNSL_E_RESOURCE;
  }

  flood_blk->flags = flags;
  return OPENNSL_E_NOT_FOUND;
}

int opennsl_port_frame_max_get(int unit, opennsl_port_t port, int* size) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  LOG_DEBUG("%s: size     = %p", __func__, size);

  _opennsl_port_info_t* const _info = _opennsl_port_info_get(port);
  if (_info == NULL)
    return OPENNSL_E_PARAM;
  
  *size = _info->frame_max;

  return OPENNSL_E_NONE;
}

int opennsl_port_frame_max_set(int unit, opennsl_port_t port, int size) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  LOG_DEBUG("%s: size     = %d", __func__, size);

  _opennsl_port_info_t* const _info = _opennsl_port_info_get(port);
  if (_info == NULL)
    return OPENNSL_E_PARAM;
  
  _info->frame_max = size;

  return OPENNSL_E_NONE;
}

int opennsl_port_l3_mtu_get(int unit, opennsl_port_t port, int* size) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  LOG_DEBUG("%s: size     = %p", __func__, size);

  _opennsl_port_info_t* const _info = _opennsl_port_info_get(port);
  if (_info == NULL)
    return OPENNSL_E_PARAM;
  
  *size = _info->l3_mtu;

  return OPENNSL_E_NONE;
}

int opennsl_port_l3_mtu_set(int unit, opennsl_port_t port, int size) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  LOG_DEBUG("%s: size     = %d", __func__, size);

  _opennsl_port_info_t* const _info = _opennsl_port_info_get(port);
  if (_info == NULL)
    return OPENNSL_E_PARAM;
  
  _info->l3_mtu = size;

  return OPENNSL_E_NONE;
}

int opennsl_port_duplex_get(int unit, opennsl_port_t port, int* duplex) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  LOG_DEBUG("%s: duplex   = %p", __func__, duplex);

  _opennsl_port_info_t* const _info = _opennsl_port_info_get(port);
  if (_info == NULL)
    return OPENNSL_E_PARAM;
  
  *duplex = _info->duplex;

  return OPENNSL_E_NONE;
}

int opennsl_port_duplex_set(int unit, opennsl_port_t port, int duplex) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  LOG_DEBUG("%s: duplex   = %d", __func__, duplex);

  _opennsl_port_info_t* const _info = _opennsl_port_info_get(port);
  if (_info == NULL)
    return OPENNSL_E_PARAM;
  
  _info->duplex = duplex;

  return OPENNSL_E_NONE;
}

int opennsl_port_speed_max(int unit, opennsl_port_t port, int *speed) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  LOG_DEBUG("%s: speed    = %p", __func__, speed);

  _opennsl_port_info_t* const _info = _opennsl_port_info_get(port);
  if (_info == NULL)
    return OPENNSL_E_PARAM;
  
  *speed = 1024 * 1024;

  return OPENNSL_E_NONE;
}

int opennsl_port_speed_set(int unit, opennsl_port_t port, int speed) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  LOG_DEBUG("%s: speed    = %d", __func__, speed);

  _opennsl_port_info_t* const _info = _opennsl_port_info_get(port);
  if (_info == NULL)
    return OPENNSL_E_PARAM;
  
  _info->speed = speed;

  return OPENNSL_E_NONE;
}

int opennsl_port_speed_get(int unit, opennsl_port_t port, int *speed) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  LOG_DEBUG("%s: speed    = %p", __func__, speed);

  _opennsl_port_info_t* const _info = _opennsl_port_info_get(port);
  if (_info == NULL)
    return OPENNSL_E_PARAM;
  
  *speed = _info->speed;

  return OPENNSL_E_NONE;
}

int opennsl_port_learn_set(int unit, opennsl_port_t port, uint32 flags) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  LOG_DEBUG("%s: flags    = %08x", __func__, flags);

  _opennsl_port_info_t* const _info = _opennsl_port_info_get(port);
  if (_info == NULL)
    return OPENNSL_E_PARAM;
  
  _info->learn_flags = flags;

  return OPENNSL_E_NONE;
}

int opennsl_port_learn_get(int unit, opennsl_port_t port, uint32 *flags) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  LOG_DEBUG("%s: flags    = %p", __func__, flags);

  _opennsl_port_info_t* const _info = _opennsl_port_info_get(port);
  if (_info == NULL)
    return OPENNSL_E_PARAM;
  
  *flags = _info->learn_flags;

  return OPENNSL_E_NONE;
}

int opennsl_port_dtag_mode_get(int unit, opennsl_port_t port, int *mode) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  LOG_DEBUG("%s: mode     = %p", __func__, mode);

  _opennsl_port_info_t* const _info = _opennsl_port_info_get(port);
  if (_info == NULL)
    return OPENNSL_E_PARAM;
  
  *mode = _info->dtag_mode;

  return OPENNSL_E_NONE;
}

int opennsl_port_dtag_mode_set(int unit, opennsl_port_t port, int mode) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  LOG_DEBUG("%s: mode     = %d", __func__, mode);

  _opennsl_port_info_t* const _info = _opennsl_port_info_get(port);
  if (_info == NULL)
    return OPENNSL_E_PARAM;
  
  _info->dtag_mode = mode;

  return OPENNSL_E_NONE;
}

int opennsl_port_class_set(int unit, opennsl_port_t port, opennsl_port_class_t pclass, uint32 class_id) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  LOG_DEBUG("%s: class:%d = %d", __func__, pclass, class_id);

  _opennsl_port_info_t* const _info = _opennsl_port_info_get(port);
  if (_info == NULL)
    return OPENNSL_E_PARAM;
  
  if (pclass < 0 || pclass > 15)
    pclass = 15;

  _info->port_class[pclass] = class_id;

  return OPENNSL_E_NONE;
}

int opennsl_port_class_get(int unit, opennsl_port_t port, opennsl_port_class_t pclass, uint32 *class_id) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  LOG_DEBUG("%s: class:%d = %p", __func__, pclass, class_id);

  _opennsl_port_info_t* const _info = _opennsl_port_info_get(port);
  if (_info == NULL)
    return OPENNSL_E_PARAM;
  
  if (pclass < 0 || pclass > 15)
    pclass = 15;

  *class_id = _info->port_class[pclass];

  return OPENNSL_E_NONE;
}

int opennsl_port_cfi_color_set(int unit, opennsl_port_t port, int cfi, opennsl_color_t color) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  LOG_DEBUG("%s: cfi      = %d", __func__, cfi);
  LOG_DEBUG("%s: color    = %d", __func__, color);

  _opennsl_port_info_t* const _info = _opennsl_port_info_get(port);
  if (_info == NULL)
    return OPENNSL_E_PARAM;

  if (cfi != 0)
    cfi = 1;

  _info->cfi_color[cfi] = color;

  return OPENNSL_E_NONE;
}

int opennsl_port_cfi_color_get(int unit, opennsl_port_t port, int cfi, opennsl_color_t *color) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  LOG_DEBUG("%s: cfi      = %d", __func__, cfi);
  LOG_DEBUG("%s: color    = %p", __func__, color);

  _opennsl_port_info_t* const _info = _opennsl_port_info_get(port);
  if (_info == NULL)
    return OPENNSL_E_PARAM;

  if (cfi != 0)
    cfi = 1;

  *color = _info->cfi_color[cfi];

  return OPENNSL_E_NONE;
}

int opennsl_port_pause_addr_set(int unit, opennsl_port_t port, opennsl_mac_t mac) {
  char mac_str[64];
  strmac(mac_str, mac);

  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  LOG_DEBUG("%s: mac      = %s", __func__, mac_str);

  _opennsl_port_info_t* const _info = _opennsl_port_info_get(port);
  if (_info == NULL)
    return OPENNSL_E_PARAM;

  memcpy(_info->pause_addr, mac, sizeof(opennsl_mac_t));
  return OPENNSL_E_NONE;
}

int opennsl_port_pause_addr_get(int unit, opennsl_port_t port, opennsl_mac_t mac) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  LOG_DEBUG("%s: mac      = %p", __func__, mac);

  _opennsl_port_info_t* const _info = _opennsl_port_info_get(port);
  if (_info == NULL)
    return OPENNSL_E_PARAM;

  memcpy(mac, _info->pause_addr, sizeof(opennsl_mac_t));
  return OPENNSL_E_NONE;
}

int opennsl_port_ifilter_set(int unit, opennsl_port_t port, int mode) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  LOG_DEBUG("%s: mode     = %d", __func__, mode);

  return OPENNSL_E_NONE;
}

int opennsl_port_link_failed_clear(int unit, opennsl_port_t port) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: port     = %d", __func__, port);

  return OPENNSL_E_NONE;
}

int opennsl_port_linkscan_get(int unit, opennsl_port_t port, int *linkscan) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  LOG_DEBUG("%s: linkscan = %p", __func__, linkscan);

  *linkscan = 1;

  return OPENNSL_E_NONE;
}

int opennsl_port_linkscan_set(int unit, opennsl_port_t port, int linkscan) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  LOG_DEBUG("%s: linkscan = %d", __func__, linkscan);

  return OPENNSL_E_NONE;
}
