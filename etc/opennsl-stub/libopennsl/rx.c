// -*- coding: utf-8 -*-

#include <stdlib.h>
#include <string.h>
#include <pthread.h>
#include <opennsl/error.h>
#include <opennsl/rx.h>
#include "libopennsl.h"
#include "logger.h"

void opennsl_rx_cfg_t_init(opennsl_rx_cfg_t* cfg) {
  memset(cfg, 0, sizeof(opennsl_rx_cfg_t));
}

int opennsl_rx_cfg_init(int unit) {
  LOG_DEBUG("%s: unit   = %d", __func__, unit);

  return OPENNSL_E_NONE;
}

int opennsl_rx_start(int unit, opennsl_rx_cfg_t *cfg) {
  LOG_DEBUG("%s: unit   = %d", __func__, unit);
  _opennsl_rx_cfg_dump(__func__, cfg);

  return OPENNSL_E_NONE;
}

int opennsl_rx_stop(int unit, opennsl_rx_cfg_t *cfg) {
  LOG_DEBUG("%s: unit   = %d", __func__, unit);
  _opennsl_rx_cfg_dump(__func__, cfg);

  return OPENNSL_E_NONE;
}

int opennsl_rx_cfg_get(int unit, opennsl_rx_cfg_t *cfg) {
  LOG_DEBUG("%s: unit   = %d", __func__, unit);
  LOG_DEBUG("%s: cfg    = %p", __func__, cfg);

  return OPENNSL_E_NONE;
}

int opennsl_rx_active(int unit) {
  LOG_DEBUG("%s: unit   = %d", __func__, unit);

  return OPENNSL_E_NONE;
}

int opennsl_rx_queue_max_get(int unit, opennsl_cos_queue_t *cosq) {
  LOG_DEBUG("%s: unit   = %d", __func__, unit);
  LOG_DEBUG("%s: cosq   = %p", __func__, cosq);

  *cosq = 10;
  return OPENNSL_E_NONE;
}

static pthread_t s_rx_tid = 0;
static int s_rx_status = 0;
static pthread_mutex_t s_rx_mutex = PTHREAD_MUTEX_INITIALIZER;
static opennsl_rx_cb_f s_rx_callback = NULL;
static void* s_rx_callback_arg = NULL;

static void* proc_rx_thread(void* arg) {
  LOG_DEBUG("%s: STARTED.", __func__);

  struct timespec ts;
  ts.tv_sec = 3;
  ts.tv_nsec = 0;

  uint8 data[] = {
    0xc4, 0x02, 0x32, 0x6b, 0x00, 0x00, 0xc4, 0x01, // MAC, MAC
    0x32, 0x58, 0x00, 0x00, 0x08, 0x06, 0x00, 0x01, // MAC, eth-type(ARP), HW-Type
    0x08, 0x00, 0x06, 0x04, 0x00, 0x01, 0xc4, 0x01,
    0x32, 0x58, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x01,
    0xc4, 0x02, 0x32, 0x6b, 0x00, 0x00, 0x0a, 0x00,
    0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
    0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
    0x00, 0x00, 0x00, 00
  };
  int data_len = sizeof(data);

  while (s_rx_status != 0) {
    nanosleep(&ts, NULL);

    if (s_rx_callback != NULL) {
      opennsl_pkt_t* pkt = NULL;
      opennsl_pkt_alloc(0, 128, 0, &pkt);
      opennsl_pkt_memcpy(pkt, 0, data, sizeof(data));
      pkt->src_port = 1;
      pkt->rx_port = 1;

      s_rx_callback(0, pkt, s_rx_callback_arg);
    }
  }

  LOG_DEBUG("%s: STOPPED", __func__);
  return NULL;
}

static void start_rx_thread() {
  const char* rx_env = getenv("OPENNSL_DEBUG_RX_THREAD");
  if (rx_env == NULL)
    return;

  pthread_mutex_lock(&s_rx_mutex);

  if (s_rx_status == 0) {
    s_rx_status = 1;
    pthread_create(&s_rx_tid, NULL, proc_rx_thread, NULL);
  }

  pthread_mutex_unlock(&s_rx_mutex);
}

static void stop_rx_thread() {
  s_rx_status = 0;
}

int opennsl_rx_register(int unit, const char *name, opennsl_rx_cb_f callback, uint8 priority, void *cookie, uint32 flags) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: name     = '%s'", __func__, name);
  LOG_DEBUG("%s: callback = %p", __func__, callback);
  LOG_DEBUG("%s: priority = %hhu", __func__, priority);
  LOG_DEBUG("%s: cookie   = %p", __func__, cookie);
  LOG_DEBUG("%s: flags    = %08x", __func__, flags);

  s_rx_callback = callback;
  s_rx_callback_arg = cookie;
  start_rx_thread();

  return OPENNSL_E_NONE;
}

int opennsl_rx_unregister(int unit, opennsl_rx_cb_f callback, uint8 priority) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: callback = %p", __func__, callback);
  LOG_DEBUG("%s: priority = %hhu", __func__, priority);

  stop_rx_thread();

  return OPENNSL_E_NONE;
}
