// -*- coding: utf-8 -*-

#include <stdlib.h>
#include <string.h>
#include <opennsl/error.h>
#include <opennsl/pkt.h>
#include "libopennsl.h"
#include "logger.h"

#define _STUB_PKT_BLK_NUM (3)

static int _opennsl_pkt_blk_init(opennsl_pkt_blk_t* blk, int size) {
  blk->data = (uint8*)malloc(size);
  if (blk->data == NULL)
    return -1;

  memset(blk->data, 0, size);
  blk->len = size;

  LOG_DEBUG("%s: opennsl_pkt_blk      = %p", __func__, blk);
  LOG_DEBUG("%s: opennsl_pkt_blk.data = %p", __func__, blk->data);
  LOG_DEBUG("%s: opennsl_pkt_blk.len  = %d", __func__, blk->len);
  return 0;
}

static void _opennsl_pkt_blk_clear(opennsl_pkt_blk_t* blk) {
  LOG_DEBUG("%s: opennsl_pkt_blk      = %p", __func__, blk);
  LOG_DEBUG("%s: opennsl_pkt_blk.data = %p", __func__, blk->data);
  LOG_DEBUG("%s: opennsl_pkt_blk.len  = %d", __func__, blk->len);
  
  if (blk->data != NULL) {
    free(blk->data);
    blk->data = NULL;
    blk->len = 0;
  }
}

int opennsl_pkt_flags_init(int unit, opennsl_pkt_t *pkt, uint32 init_flags) {
  LOG_DEBUG("%s: unit    = %d", __func__, unit);
  LOG_DEBUG("%s: pkt     = %p", __func__, pkt);
  LOG_DEBUG("%s: flags   = %d", __func__, init_flags);

  pkt->flags = init_flags;
  return OPENNSL_E_NONE;
}

int opennsl_pkt_alloc(int unit, int size, uint32 flags, opennsl_pkt_t **pkt_buf) {
  LOG_DEBUG("%s: unit    = %d", __func__, unit);
  LOG_DEBUG("%s: size    = %d", __func__, size);
  LOG_DEBUG("%s: flags   = %08x", __func__, flags);
  LOG_DEBUG("%s: pkt_buf = %p", __func__, pkt_buf);

  if (size < 0) {
    return OPENNSL_E_PARAM;
  }

  opennsl_pkt_t* pkt = (opennsl_pkt_t*)malloc(sizeof(opennsl_pkt_t));
  if (pkt == NULL)
    return OPENNSL_E_PARAM;

  opennsl_pkt_blk_t* blks = (opennsl_pkt_blk_t*)malloc(sizeof(opennsl_pkt_blk_t) * _STUB_PKT_BLK_NUM);
  if (blks == NULL) {
    free(pkt);
    return OPENNSL_E_PARAM;
  }

  int index;
  for (index = 0; index < _STUB_PKT_BLK_NUM; index++) {
    if (_opennsl_pkt_blk_init(blks + index, size) < 0) {
      free(pkt);
      free(blks);
      return OPENNSL_E_PARAM;
    }
  }

  opennsl_pkt_flags_init(unit, pkt, flags);
  pkt->pkt_data = blks;
  pkt->blk_count = _STUB_PKT_BLK_NUM;

  *pkt_buf = pkt;
  return OPENNSL_E_NONE;
}

int opennsl_pkt_free(int unit, opennsl_pkt_t *pkt) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: pkt      = %p", __func__, pkt);
  LOG_DEBUG("%s: pkt.data = %p", __func__, pkt->pkt_data);
  LOG_DEBUG("%s: pkt.cnt  = %d", __func__, pkt->blk_count);

  int index;

  opennsl_pkt_blk_t* blks = pkt->pkt_data;
  for (index = 0; index < pkt->blk_count; index++) {
    _opennsl_pkt_blk_clear(blks + index);
  }

  free(pkt);
  free(blks);

  return OPENNSL_E_NONE;
}

int opennsl_pkt_memcpy(opennsl_pkt_t *pkt, int dest_byte, uint8 *src, int len) {
  LOG_DEBUG("%s: pkt     = %p", __func__, pkt);
  LOG_DEBUG("%s: start   = %d", __func__, dest_byte);
  LOG_DEBUG("%s: data    = %p / %d", __func__, src, len);

  if (pkt->blk_count == 0 || pkt->pkt_data == NULL)
    return OPENNSL_E_PARAM;

  if ((pkt->pkt_data[0].len < (dest_byte + len)) || (pkt->pkt_data[0].data == NULL))
    return OPENNSL_E_PARAM;

  memcpy(pkt->pkt_data[0].data + dest_byte, src, len);

  _opennsl_pkt_dump(__func__, pkt);
  return OPENNSL_E_NONE;
}
