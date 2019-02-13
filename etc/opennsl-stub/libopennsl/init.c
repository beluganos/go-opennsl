// -*- coding: utf-8 -*-

#include <string.h>
#include <opennsl/error.h>
#include <opennsl/init.h>
#include "libopennsl.h"
#include "logger.h"

int opennsl_info_get(int unit, opennsl_info_t *info) {
  LOG_DEBUG("%s: unit  = %d", __func__, unit);
  LOG_DEBUG("%s: info  = %p", __func__, info);
}

void opennsl_info_t_init(opennsl_info_t *info) {
  LOG_DEBUG("%s: info  = %p", __func__, info);

  memset(info, 0, sizeof(opennsl_info_t));
}
