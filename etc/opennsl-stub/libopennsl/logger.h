// -*- coding: utf-8 -*-

#ifndef _LOGGER_H
#define _LOGGER_H

#include <stdio.h>

#define LOG_ATTR(a,b) __attribute__ ((format (printf, a, b)))
#define LOG_INFO(fmt, ...)  log_output(stdout, "[INFO ] "fmt, ## __VA_ARGS__)
#define LOG_WARN(fmt, ...)  log_output(stdout, "[WARN ] "fmt, ## __VA_ARGS__)
#define LOG_DEBUG(fmt, ...) log_output(stdout, "[DEBUG] "fmt, ## __VA_ARGS__)

#ifdef __cplusplus
extern "C" {
#endif

  void log_output(FILE* fp, const char* fmt, ...) LOG_ATTR(2,3);

#ifdef __cplusplus
}
#endif

#endif // _LOGGER_H
