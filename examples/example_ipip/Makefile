# -*- coding: utf-8 -*-

#  Copyright (C) 2018 Nippon Telegraph and Telephone Corporation.
# 
#  Licensed under the Apache License, Version 2.0 (the "License");
#  you may not use this file except in compliance with the License.
#  You may obtain a copy of the License at
# 
#     http:# www.apache.org/licenses/LICENSE-2.0
# 
#  Unless required by applicable law or agreed to in writing, software
#  distributed under the License is distributed on an "AS IS" BASIS,
#  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
#  implied.
#  See the License for the specific language governing permissions and
#  limitations under the License.

TARGET = example_ipip
OBJS = util.o main.o
CC = gcc

CFLAGS = -Wall
#CFLAGS += -DCDP_PACKAGE=1
#CFLAGS += -DOPENNSL_PRODUCT_DNX
#CFLAGS += -DOPENNSL_WARM_BOOT_SUPPORT
FEATURE_LIST = L3
CFLAGS  += $(foreach feature,$(FEATURE_LIST), -DINCLUDE_$(feature))
CFLAGS += `pkg-config --cflags libopennsl`
LDFLAGS = `pkg-config --libs-only-L libopennsl`
LIBS = `pkg-config --libs-only-l libopennsl` -lc

.SUFFIXES: .c .o

.PHONY: all
all: $(TARGET)

$(TARGET): $(OBJS)
	$(CC) -o $(TARGET) $(LDFLAGS) $^ $(LIBS)

.c.o:
	$(CC) $(CFLAGS) -c $<

.PHONY: clean
clean:
	$(RM) $(TARGET) $(OBJS)
