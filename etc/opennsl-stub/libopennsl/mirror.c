// -*- coding: utf-8 -*-

#include <stdlib.h>
#include <string.h>
#include <opennsl/error.h>
#include <opennsl/mirror.h>
#include "libopennsl.h"
#include "logger.h"

int opennsl_mirror_init(int unit) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);

  return OPENNSL_E_NONE;
}

int opennsl_mirror_to_set(int unit, opennsl_port_t port) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: port     = %d", __func__, port);

  return OPENNSL_E_NONE;
}

int opennsl_mirror_to_get(int unit, opennsl_port_t *port) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: port     = %p", __func__, port);

  *port = 1;
  return OPENNSL_E_NONE;
}

int opennsl_mirror_ingress_set(int unit, opennsl_port_t port, int val) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  LOG_DEBUG("%s: val      = %d", __func__, val);

  return OPENNSL_E_NONE;
}

int opennsl_mirror_ingress_get(int unit, opennsl_port_t port, int* val) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  LOG_DEBUG("%s: val      = %p", __func__, val);

  *val = 1;
  return OPENNSL_E_NONE;
}

int opennsl_mirror_egress_set(int unit, opennsl_port_t port, int val) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  LOG_DEBUG("%s: val      = %d", __func__, val);

  return OPENNSL_E_NONE;
}

int opennsl_mirror_egress_get(int unit, opennsl_port_t port, int* val) {
  LOG_DEBUG("%s: unit     = %d", __func__, unit);
  LOG_DEBUG("%s: port     = %d", __func__, port);
  LOG_DEBUG("%s: val      = %p", __func__, val);

  *val = 1;
  return OPENNSL_E_NONE;
}

int opennsl_mirror_port_set(int unit, opennsl_port_t port, opennsl_module_t dest_mod, opennsl_port_t dest_port,  uint32 flags) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: port      = %d", __func__, port);
  LOG_DEBUG("%s: dest_mod  = %d", __func__, dest_mod);
  LOG_DEBUG("%s: dest_port = %d", __func__, dest_port);
  LOG_DEBUG("%s: flags     = %08x", __func__, flags);

  return OPENNSL_E_NONE;
}

int opennsl_mirror_port_get(int unit, opennsl_port_t port, opennsl_module_t* dest_mod, opennsl_port_t* dest_port,  uint32* flags) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: port      = %d", __func__, port);
  LOG_DEBUG("%s: dest_mod  = %p", __func__, dest_mod);
  LOG_DEBUG("%s: dest_port = %p", __func__, dest_port);
  LOG_DEBUG("%s: flags     = %p", __func__, flags);

  *dest_mod = 1;
  *dest_port = 2;
  *flags = 3;

  return OPENNSL_E_NONE;
}

int opennsl_mirror_port_dest_add(int unit, opennsl_port_t port, uint32 flags,  opennsl_gport_t mirror_dest_id) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: port      = %d", __func__, port);
  LOG_DEBUG("%s: flags     = %08x", __func__, flags);
  LOG_DEBUG("%s: dest_id   = %d", __func__, mirror_dest_id);

  return OPENNSL_E_NONE;
}

int opennsl_mirror_port_dest_delete(int unit, opennsl_port_t port, uint32 flags, opennsl_gport_t mirror_dest_id) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: port      = %d", __func__, port);
  LOG_DEBUG("%s: flags     = %08x", __func__, flags);
  LOG_DEBUG("%s: dest_id   = %d", __func__, mirror_dest_id);

  return OPENNSL_E_NONE;
}

int opennsl_mirror_port_dest_delete_all(int unit, opennsl_port_t port, uint32 flags) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: port      = %d", __func__, port);
  LOG_DEBUG("%s: flags     = %08x", __func__, flags);

  return OPENNSL_E_NONE;
}

int opennsl_mirror_port_dest_get(int unit, opennsl_port_t port, uint32 flags, int mirror_dest_size, opennsl_gport_t *mirror_dest, int *mirror_dest_count) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: port      = %d", __func__, port);
  LOG_DEBUG("%s: flags     = %08x", __func__, flags);
  LOG_DEBUG("%s: size      = %d", __func__, mirror_dest_size);
  LOG_DEBUG("%s: dest      = %p", __func__, mirror_dest);
  LOG_DEBUG("%s: count     = %p", __func__, mirror_dest_count);

  int index;
  for (index = 0; index < mirror_dest_size; index++) {
    mirror_dest[index] = index + 1;
  }

  *mirror_dest_count = mirror_dest_size;
  return OPENNSL_E_NONE;
}

void opennsl_mirror_destination_t_init(opennsl_mirror_destination_t *mirror_dest) {
  memset(mirror_dest, 0, sizeof(opennsl_mirror_destination_t));
}

int opennsl_mirror_destination_create(int unit, opennsl_mirror_destination_t *mirror_dest) {
  LOG_DEBUG("%s: unit      = %d", __func__, unit);
  LOG_DEBUG("%s: dest      = %p", __func__, mirror_dest);

  return OPENNSL_E_NONE;
}
