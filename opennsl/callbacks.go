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
	"sync"
)

type CallbackMap struct {
	Counter   uint64
	Callbacks map[uint64]interface{}
	Mutex     sync.Mutex
}

func NewCallbackMap() *CallbackMap {
	return &CallbackMap{
		Counter:   0,
		Callbacks: make(map[uint64]interface{}),
	}
}

func (m *CallbackMap) Add(cb interface{}) uint64 {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	m.Counter++
	m.Callbacks[m.Counter] = cb
	return m.Counter
}

func (m *CallbackMap) Del(n uint64) {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	delete(m.Callbacks, n)
}

func (m *CallbackMap) Get(n uint64) (interface{}, bool) {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	cb, ok := m.Callbacks[n]
	return cb, ok
}
