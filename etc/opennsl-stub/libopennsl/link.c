// -*- coding: utf-8 -*-

#include <stdlib.h>
#include <string.h>
#include <opennsl/error.h>
#include <opennsl/link.h>
#include "libopennsl.h"
#include "logger.h"

int opennsl_linkscan_detach(int unit) {
  LOG_DEBUG("%s: unit    = %d", __func__, unit);
  return OPENNSL_E_NONE;
}

int opennsl_linkscan_enable_set(int unit, int us) {
  LOG_DEBUG("%s: unit    = %d", __func__, unit);
  LOG_DEBUG("%s: us      = %d", __func__, us);
  return OPENNSL_E_NONE;
}

int opennsl_linkscan_enable_get(int unit, int *us) {
  LOG_DEBUG("%s: unit    = %d", __func__, unit);
  LOG_DEBUG("%s: us      = %p", __func__, us);

  *us = 1000000;
  return OPENNSL_E_NONE;
}

int opennsl_linkscan_mode_set(int unit, opennsl_port_t port, int mode) {
  LOG_DEBUG("%s: unit    = %d", __func__, unit);
  LOG_DEBUG("%s: port    = %d", __func__, port);
  LOG_DEBUG("%s: mode    = %d", __func__, mode);

  return OPENNSL_E_NONE;
}

int opennsl_linkscan_mode_set_pbm(int unit, opennsl_pbmp_t pbm, int mode) {
  LOG_DEBUG("%s: unit    = %d", __func__, unit);
  LOG_DEBUG("%s: mode    = %d", __func__, mode);
  _opennsl_pbmp_dump(__func__, &pbm);

  return OPENNSL_E_NONE;
}

int opennsl_linkscan_mode_get(int unit, opennsl_port_t port, int *mode) {
  LOG_DEBUG("%s: unit    = %d", __func__, unit);
  LOG_DEBUG("%s: port    = %d", __func__, port);
  LOG_DEBUG("%s: mode    = %p", __func__, mode);

  return OPENNSL_E_NONE;
}

int opennsl_linkscan_register(int unit, opennsl_linkscan_handler_t f) {
  LOG_DEBUG("%s: unit    = %d", __func__, unit);
  LOG_DEBUG("%s: handler = %p", __func__, f);
  return OPENNSL_E_NONE;
}

int opennsl_linkscan_unregister(int unit, opennsl_linkscan_handler_t f) {
  LOG_DEBUG("%s: unit    = %d", __func__, unit);
  LOG_DEBUG("%s: handler = %p", __func__, f);
  return OPENNSL_E_NONE;
}
