// -*- coding: utf-8 -*-

#include <opennsl/error.h>
#include <opennsl/l2X.h>

void opennsl_l2_learn_limit_t_init(opennsl_l2_learn_limit_t *limit) {
  limit->flags = 1;
  limit->limit = 2;
  
}

int opennsl_l2_learn_limit_set(int unit, opennsl_l2_learn_limit_t *limit) {
  return OPENNSL_E_NONE;
}

int opennsl_l2_learn_limit_get(int unit, opennsl_l2_learn_limit_t *limit) {
  return OPENNSL_E_NONE;
}
