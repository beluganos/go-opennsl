// -*- coding: utf-8 -*-

#include <opennsl/error.h>
#include <sal/driver.h>
#include "libopennsl.h"
#include "logger.h"

int opennsl_driver_init(opennsl_init_t *init) {
  _opennsl_init_dump(__func__, init);
}

int opennsl_driver_exit() {
  LOG_DEBUG("%s:", __func__);
}

#ifndef OPENNSL_BOOT_F_WARM_BOOT
#define OPENNSL_BOOT_F_WARM_BOOT (0x200000)
#endif

unsigned int opennsl_driver_boot_flags_get(void) {
  LOG_DEBUG("%s:", __func__);
  return OPENNSL_BOOT_F_WARM_BOOT;
}

int opennsl_driver_shell() {
  LOG_DEBUG("%s:", __func__);
  return OPENNSL_E_NONE;
}

int opennsl_driver_process_command(char *commandBuf) {
  LOG_DEBUG("%s: cmd = '%s'", __func__, commandBuf);
  return OPENNSL_E_NONE;
}

/*
char *readline(const char *prompt) {
  return NULL
}
*/

void platform_phy_cleanup() {
  LOG_DEBUG("%s:", __func__);
}
