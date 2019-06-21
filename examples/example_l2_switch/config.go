// -*- coding: utf-8 -*-

package main

import (
	"github.com/beluganos/go-opennsl/opennsl"
	"github.com/spf13/viper"
)

func convertPorts(src []int) []opennsl.Port {
	ports := make([]opennsl.Port, len(src))
	for index, port := range src {
		ports[index] = opennsl.Port(port)
	}
	return ports
}

type VlanConfig struct {
	Id              uint16 `mapstructure:"id"`
	AccessPort      []int  `mapstructure:"access_port"`
	TrunkPort       []int  `mapstructure:"trunk_port"`
	NativeTagPort   []int  `mapstructure:"native_tag_port"`
	NativeUntagPort []int  `mapstructure:"native_untag_port"`
}

func NewVlanConfig() *VlanConfig {
	return &VlanConfig{
		AccessPort:      []int{},
		TrunkPort:       []int{},
		NativeTagPort:   []int{},
		NativeUntagPort: []int{},
	}
}

func (c *VlanConfig) Vlan() opennsl.Vlan {
	return opennsl.Vlan(c.Id)
}

func (c *VlanConfig) AccessPorts() []opennsl.Port {
	return convertPorts(c.AccessPort)
}

func (c *VlanConfig) TrunkPorts() []opennsl.Port {
	return convertPorts(c.TrunkPort)
}

func (c *VlanConfig) NativeTagPorts() []opennsl.Port {
	return convertPorts(c.NativeTagPort)
}

func (c *VlanConfig) NativeUntagPorts() []opennsl.Port {
	return convertPorts(c.NativeUntagPort)
}

type Config struct {
	Vlans []*VlanConfig `mapstructure:"vlans"`
}

func NewConfig() *Config {
	return &Config{
		Vlans: []*VlanConfig{},
	}
}

func ReadConfig(configPath, configType string) (*Config, error) {
	v := viper.New()
	v.SetConfigType(configType)
	v.SetConfigFile(configPath)

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	c := NewConfig()
	if err := v.UnmarshalExact(c); err != nil {
		return nil, err
	}

	return c, nil
}
