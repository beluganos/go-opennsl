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

import (
	"fmt"
)

type RxCallbackKey struct {
	Unit int
	Pri  uint8
}

type RxCallbackMap struct {
	entries   map[RxCallbackKey]*RxCallbackKey
	callbacks map[*RxCallbackKey]RxCallback
}

func NewRxCallbackMap() *RxCallbackMap {
	return &RxCallbackMap{
		entries:   make(map[RxCallbackKey]*RxCallbackKey),
		callbacks: make(map[*RxCallbackKey]RxCallback),
	}
}

func (m *RxCallbackMap) Register(unit int, callback RxCallback, pri uint8) (*RxCallbackKey, error) {
	key := &RxCallbackKey{
		Unit: unit,
		Pri:  pri,
	}

	if _, ok := m.entries[*key]; ok {
		return nil, fmt.Errorf("callback already exist. unit=%d, pri=%d", unit, pri)
	}

	m.entries[*key] = key
	m.callbacks[key] = callback
	return key, nil
}

func (m *RxCallbackMap) Unregister(unit int, pri uint8) {
	key := RxCallbackKey{
		Unit: unit,
		Pri:  pri,
	}

	if k, ok := m.entries[key]; ok {
		delete(m.entries, key)
		delete(m.callbacks, k)
	}
}

func (m *RxCallbackMap) Get(key *RxCallbackKey) (RxCallback, bool) {
	if cb, ok := m.callbacks[key]; ok {
		return cb, true
	}
	return nil, false
}
