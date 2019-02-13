// -*- coding: utf-8 -*-

#include <stdlib.h>
#include <string.h>
#include <opennsl/error.h>
#include <opennsl/pkt.h>
#include "libopennsl.h"
#include "logger.h"


int opennsl_tx(int unit, opennsl_pkt_t *pkt, void *cookie) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: pkt      = %p", __func__, pkt);
  LOG_DEBUG("%s: cookie   = %p", __func__, cookie);
  _opennsl_pkt_dump(__func__, pkt);

  return OPENNSL_E_NONE;
}
