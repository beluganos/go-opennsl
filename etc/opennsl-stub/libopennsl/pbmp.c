// -*- coding: utf-8 -*-

#include <opennsl/error.h>
#include <shared/pbmp.h>
#include "logger.h"

int _shr_pbmp_bmeq (_shr_pbmp_t* b1, _shr_pbmp_t* b2) {
  int index;
  for (index = 0; index < _SHR_PBMP_WORD_MAX; index ++) {
    if (b1->pbits[index] != b2->pbits[index]) {
      return 0;
    }
  }
  return 1;
}
 
int _shr_pbmp_bmnull (_shr_pbmp_t* b) {
  int index;
  for (index = 0; index < _SHR_PBMP_WORD_MAX; index ++) {
    if (b->pbits[index] != 0) {
      return 0;
    }
  }
  return 1;
}
 
char* _shr_pbmp_format (_shr_pbmp_t b, char* buf) {
  return buf;
}
 
