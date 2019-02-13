// -*- coding: utf-8 -*-

#include <stdarg.h>
#include "logger.h"

void log_output(FILE* fp, const char* fmt, ...) {
  va_list ap;

  va_start(ap, fmt);
  vprintf(fmt, ap);
  printf("%s", "\n");
  va_end(ap);
}
