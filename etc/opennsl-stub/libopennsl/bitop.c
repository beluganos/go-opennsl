// -*- coding: utf-8 -*-

#include <string.h>
#include <shared/bitop.h>
#include "libopennsl.h"
#include "logger.h"

void shr_bitop_range_clear(SHR_BITDCL *a, CONST int start_pos, CONST int bit_num) {
  int pos;
  for (pos = start_pos; pos < start_pos + bit_num; pos++) {
    const SHR_BITDCL mask = 1U << (pos % 32);
    a[pos / 32] &= ~mask;
  }
}
