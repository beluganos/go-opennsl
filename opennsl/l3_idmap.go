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
// L3EgressIDConv converts id to user defined id.
//
type L3EgressIDConv func(id uint32) uint32

//
// l3EgressIDConvDefault returns id.
//
func l3EgressIDConvDefault(id uint32) uint32 {
	return id
}

//
// L3EgressIDMap has id(key) and l3-egress-id(value).
//
type L3EgressIDMap struct {
	sync.Map
	conv L3EgressIDConv
}

//
// NewL3EgressIDMap creates new instance.
//
func NewL3EgressIDMap(conv L3EgressIDConv) *L3EgressIDMap {
	if conv == nil {
		conv = l3EgressIDConvDefault
	}
	return &L3EgressIDMap{
		conv: conv,
	}
}

//
// Register registers id and l3egrId.
//
func (m *L3EgressIDMap) Register(id uint32, l3egrId L3EgressID) bool {
	_, ok := m.Map.LoadOrStore(m.conv(id), l3egrId)
	return !ok
}

//
// Unregister removes id.
//
func (m *L3EgressIDMap) Unregister(id uint32) {
	m.Map.Delete(m.conv(id))
}

//
// Get returns L3EgressID by id.
//
func (m *L3EgressIDMap) Get(id uint32) (L3EgressID, bool) {
	if v, ok := m.Map.Load(m.conv(id)); ok {
		return v.(L3EgressID), true
	}
	return 0, false
}

//
// Traverse enumerates all entries.
//
func (m *L3EgressIDMap) Traverse(f func(uint32, L3EgressID) bool) {
	m.Map.Range(func(key, value interface{}) bool {
		return f(key.(uint32), value.(L3EgressID))
	})
}

//
// L3IngressIDConv convets id to user defined id.
//
type L3IngressIDConv func(id uint32) uint32

//
// l3IngressIDConvDefault returns id.
//
func l3IngressIDConvDefault(id uint32) uint32 {
	return id
}

//
// L3IngressIDMap has id(key) and l3-ingress-id(value).
//
type L3IngressIDMap struct {
	sync.Map
	conv L3IngressIDConv
}

//
// NewL3IngressIDMap creates new instance.
//
func NewL3IngressIDMap(conv L3IngressIDConv) *L3IngressIDMap {
	if conv == nil {
		conv = l3IngressIDConvDefault
	}
	return &L3IngressIDMap{
		conv: conv,
	}
}

//
// Register registers id and l3 ingress id.
//
func (m *L3IngressIDMap) Register(id uint32, l3igrId L3IngressID) bool {
	_, ok := m.Map.LoadOrStore(m.conv(id), l3igrId)
	return !ok
}

//
// Unregister removes id.
//
func (m *L3IngressIDMap) Unregister(id uint32) {
	m.Map.Delete(m.conv(id))
}

//
// Get returns l3 ingress id by id.
//
func (m *L3IngressIDMap) Get(id uint32) (L3IngressID, bool) {
	if v, ok := m.Map.Load(m.conv(id)); ok {
		return v.(L3IngressID), true
	}
	return 0, false
}

//
// Traverse enumerates all entries.
//
func (m *L3IngressIDMap) Traverse(f func(uint32, L3IngressID) bool) {
	m.Map.Range(func(key, value interface{}) bool {
		return f(key.(uint32), value.(L3IngressID))
	})
}

//
// L3IfaceIDConv converts id to user defined id.
//
type L3IfaceIDConv func(id uint32, vid uint16) uint32

//
// l3IfaceIDConvDefault returns id.
//
func l3IfaceIDConvDefault(id uint32, vid uint16) uint32 {
	return (uint32(vid) << 16) + uint32(id)
}

//
// L3IfaceIDMap has id(key) and l3-egress-id(value).
//
type L3IfaceIDMap struct {
	sync.Map
	conv L3IfaceIDConv
}

//
// NewL3IfaceIDMap creates new instance.
//
func NewL3IfaceIDMap(conv L3IfaceIDConv) *L3IfaceIDMap {
	if conv == nil {
		conv = l3IfaceIDConvDefault
	}
	return &L3IfaceIDMap{
		conv: conv,
	}
}

//
// Register registers id and l3ifaceId.
//
func (m *L3IfaceIDMap) Register(id uint32, vid uint16, l3ifaceId L3IfaceID) bool {
	_, ok := m.Map.LoadOrStore(m.conv(id, vid), l3ifaceId)
	return !ok
}

//
// Unregister removes id.
//
func (m *L3IfaceIDMap) Unregister(id uint32, vid uint16) {
	m.Map.Delete(m.conv(id, vid))
}

//
// Get returns L3IfaceID by id.
//
func (m *L3IfaceIDMap) Get(id uint32, vid uint16) (L3IfaceID, bool) {
	if v, ok := m.Map.Load(m.conv(id, vid)); ok {
		return v.(L3IfaceID), true
	}
	return 0, false
}

//
// Traverse enumerates all entries.
//
func (m *L3IfaceIDMap) Traverse(f func(uint32, L3IfaceID) bool) {
	m.Map.Range(func(key, value interface{}) bool {
		return f(key.(uint32), value.(L3IfaceID))
	})
}
