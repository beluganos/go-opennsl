// -*- coding: utf-8 -*-

// Copyright (C) 2018 Nippon Telegraph and Telephone Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package opennsl

/*
#include <opennsl/types.h>
#include <opennsl/port.h>
*/
import "C"

import (
	"fmt"
)

//
// PortConfigType
//
type PortConfigType int

const (
	PORT_CONFIG_NONE PortConfigType = iota
	PORT_CONFIG_FE
	PORT_CONFIG_GE
	PORT_CONFIG_XE
	PORT_CONFIG_CE
	PORT_CONFIG_E
	PORT_CONFIG_HG
	PORT_CONFIG_SCI
	PORT_CONFIG_SFI
	PORT_CONFIG_SPI
	PORT_CONFIG_SPI_SUB
	PORT_CONFIG_PORT
	PORT_CONFIG_CPU
	PORT_CONFIG_ALL
	PORT_CONFIG_STACK_EXT
	PORT_CONFIG_TDM
	PORT_CONFIG_PON
	PORT_CONFIG_LLID
	PORT_CONFIG_IL
	PORT_CONFIG_XL
	PORT_CONFIG_RCY
	PORT_CONFIG_NIF
	PORT_CONFIG_CONTROL
)

//
// PortConfig
//
type PortConfig C.opennsl_port_config_t

func (v *PortConfig) C() *C.opennsl_port_config_t {
	return (*C.opennsl_port_config_t)(v)
}

func (v *PortConfig) PBmp(pcType PortConfigType) (*PBmp, error) {
	pbmp := func() *C.opennsl_pbmp_t {
		switch pcType {
		case PORT_CONFIG_FE:
			return &v.fe
		case PORT_CONFIG_GE:
			return &v.ge
		case PORT_CONFIG_XE:
			return &v.xe
		case PORT_CONFIG_CE:
			return &v.ce
		case PORT_CONFIG_E:
			return &v.e
		case PORT_CONFIG_HG:
			return &v.hg
		case PORT_CONFIG_SCI:
			return &v.sci
		case PORT_CONFIG_SFI:
			return &v.sfi
		case PORT_CONFIG_SPI:
			return &v.spi
		case PORT_CONFIG_SPI_SUB:
			return &v.spi_subport
		case PORT_CONFIG_PORT:
			return &v.port
		case PORT_CONFIG_CPU:
			return &v.cpu
		case PORT_CONFIG_ALL:
			return &v.all
		case PORT_CONFIG_STACK_EXT:
			return &v.stack_ext
		case PORT_CONFIG_TDM:
			return &v.tdm
		case PORT_CONFIG_PON:
			return &v.pon
		case PORT_CONFIG_LLID:
			return &v.llid
		case PORT_CONFIG_IL:
			return &v.il
		case PORT_CONFIG_XL:
			return &v.xl
		case PORT_CONFIG_RCY:
			return &v.rcy
		case PORT_CONFIG_NIF:
			return &v.nif
		case PORT_CONFIG_CONTROL:
			return &v.control
		default:
			return nil
		}
	}()
	if pbmp == nil {
		return nil, fmt.Errorf("Invalid PortConfigtype. %d", pcType)
	}

	return (*PBmp)(pbmp), nil
}

func NewPortConfig() *PortConfig {
	config := &PortConfig{}
	config.Init()
	return config
}

func (v *PortConfig) Init() {
	C.opennsl_port_config_t_init(v.C())
}

func PortConfigGet(unit int) (*PortConfig, error) {
	cfg := PortConfig{}
	cfg.Init()

	rc := C.opennsl_port_config_get(C.int(unit), cfg.C())
	if err := ParseError(rc); err != nil {
		return nil, err
	}

	return &cfg, nil
}
