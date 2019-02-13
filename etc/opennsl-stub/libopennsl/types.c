// -*- coding: utf-8 -*-

#include <opennsl/types.h>
#include <opennsl/error.h>


opennsl_ip_t opennsl_ip_mask_create(int	len) {
  
  if (len > 32) {
    len = 32;
  }

  int pos;
  opennsl_ip_t ip = 0;
  for (pos = 0; pos < len; pos++) {
    ip |= (1<<pos);
  }

  return ip;
}


int opennsl_ip6_mask_create(opennsl_ip6_t ip6, int len) {
  int index = 0;
  int bytenum = len / 8;
  int bytepos = 0;

  if (len > 16 * 8) {
    return OPENNSL_E_PARAM;
  }

  do {
    int bitpos;
    for (bitpos = 0; bitpos < 8; bitpos++) {
      if (bitpos + bytepos * 8 >= len) {
	return OPENNSL_E_NONE;
      }
      ip6[bytepos] |= (uint8)(1<<bitpos);
      
      bytepos++;
    }
  } while (bytepos < bytenum);

  return OPENNSL_E_NONE;
}
