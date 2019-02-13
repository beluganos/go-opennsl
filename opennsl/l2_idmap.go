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

//
// L2StationIDConv converts id to user defined id.
//
type L2StationIDConv func(id uint32) uint32

//
// l2StationIDConvDefault returns id.
//
func l2StationIDConvDefault(id uint32) uint32 {
	return id
}

//
// L2StationIDMap has id(key) and l2-station-id(value).
//
type L2StationIDMap struct {
	sync.Map
	conv L2StationIDConv
}

//
// NewL2StationIDMap creates new instance.
//
func NewL2StationIDMap(conv L2StationIDConv) *L2StationIDMap {
	if conv == nil {
		conv = l2StationIDConvDefault
	}
	return &L2StationIDMap{
		conv: conv,
	}
}

//
// Register registers id and l3egrId.
//
func (m *L2StationIDMap) Register(id uint32, l3egrId L2StationID) bool {
	_, ok := m.Map.LoadOrStore(m.conv(id), l3egrId)
	return !ok
}

//
// Unregister removes id.
//
func (m *L2StationIDMap) Unregister(id uint32) {
	m.Map.Delete(m.conv(id))
}

//
// Get returns L2StationID by id.
//
func (m *L2StationIDMap) Get(id uint32) (L2StationID, bool) {
	if v, ok := m.Map.Load(m.conv(id)); ok {
		return v.(L2StationID), true
	}
	return 0, false
}

//
// Traverse enumerates all entries.
//
func (m *L2StationIDMap) Traverse(f func(uint32, L2StationID) bool) {
	m.Map.Range(func(key, value interface{}) bool {
		return f(key.(uint32), value.(L2StationID))
	})
}
