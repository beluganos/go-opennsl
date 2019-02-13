// -*- coding: utf-8 -*-

#include <opennsl/types.h>
#include <opennsl/port.h>

// set port number.
void _opennsl_port_num_set(int n);

// get port number.
int _opennsl_port_num_get(void);

// set stats
int _opennsl_stat_set(opennsl_port_t port, opennsl_stat_val_t type, uint64 val);

int _opennsl_stat_multi_set(opennsl_port_t port, int nstat, opennsl_stat_val_t* type_arr, uint64* val_arr);
