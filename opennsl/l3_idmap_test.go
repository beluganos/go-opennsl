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
	"testing"
)

func TestL3IfaceIDMap_Reg_Unreg_1(t *testing.T) {
	m := NewL3IfaceIDMap(nil)

	port := uint32(1)
	vid := uint16(10)
	l3if := L3IfaceID(100)

	if b := m.Register(port, vid, l3if); b != true {
		t.Errorf("IDMap_L3Iface.Register error.")
	}

	if ifaceID, b := m.Get(port, vid); b != true || ifaceID != l3if {
		t.Errorf("IDMap_L3Iface.Get error or unmatch.")
	}

	m.Unregister(port, vid)

	if _, b := m.Get(port, vid); b != false {
		t.Errorf("IDMap_L3Iface.Get error or unmatch.")
	}
}

func TestL3IfaceIDMap_Reg_Unreg_2(t *testing.T) {
	m := NewL3IfaceIDMap(nil)

	port1 := uint32(1)
	port2 := uint32(2)
	vid1 := uint16(10)
	vid2 := uint16(20)
	l3if1 := L3IfaceID(100)
	l3if2 := L3IfaceID(200)

	if b := m.Register(port1, vid1, l3if1); b != true {
		t.Errorf("IDMap_L3Iface.Register error.")
	}

	if b := m.Register(port2, vid2, l3if2); b != true {
		t.Errorf("IDMap_L3Iface.Register error.")
	}

	if ifaceID, b := m.Get(port1, vid1); b != true || ifaceID != l3if1 {
		t.Errorf("IDMap_L3Iface.Get error or unmatch.")
	}

	if ifaceID, b := m.Get(port2, vid2); b != true || ifaceID != l3if2 {
		t.Errorf("IDMap_L3Iface.Get error or unmatch.")
	}

	m.Unregister(port1, vid1)

	if _, b := m.Get(port1, vid1); b != false {
		t.Errorf("IDMap_L3Iface.Get error or unmatch.")
	}

	if ifaceID, b := m.Get(port2, vid2); b != true || ifaceID != l3if2 {
		t.Errorf("IDMap_L3Iface.Get error or unmatch.")
	}

	m.Unregister(port2, vid2)

	if _, b := m.Get(port1, vid1); b != false {
		t.Errorf("IDMap_L3Iface.Get error or unmatch.")
	}

	if _, b := m.Get(port2, vid2); b != false {
		t.Errorf("IDMap_L3Iface.Get error or unmatch.")
	}
}

func TestL3IfaceIDMap_Register_dup(t *testing.T) {
	m := NewL3IfaceIDMap(nil)

	m.Register(1, 10, 100)

	if b := m.Register(1, 10, 150); b != false {
		t.Errorf("IDMap_L3Iface.Register must be error.")
	}

	if ifaceID, b := m.Get(1, 10); b != true || ifaceID != 100 {
		t.Errorf("IDMap_L3Iface.Register must not be change. ifaceID=%d", ifaceID)
	}
}

func TestL3IfaceIDMap_Get_not_found(t *testing.T) {
	m := NewL3IfaceIDMap(nil)

	m.Register(1, 10, 100)

	if _, ok := m.Get(1, 11); ok {
		t.Errorf("IDMap_L3Iface.Get must be error.")
	}

	if ifaceID, ok := m.Get(1, 10); (!ok) || ifaceID != 100 {
		t.Errorf("IDMap_L3Iface.Get error. %d", ifaceID)
	}
}

func TestL3IfaceIDMap_Traverse_0(t *testing.T) {
	m := NewL3IfaceIDMap(nil)
	var count int

	m.Traverse(func(key uint32, val L3IfaceID) bool {
		count++
		return true
	})

	if count != 0 {
		t.Errorf("IDMap_L3Iface.Traverse unmatch.")
	}
}

func TestL3IfaceIDMap_Traverse(t *testing.T) {
	m := NewL3IfaceIDMap(nil)
	var count int

	m.Register(1, 10, 100)
	m.Register(2, 20, 200)

	m.Traverse(func(key uint32, val L3IfaceID) bool {
		count++
		return true
	})

	if count != 2 {
		t.Errorf("IDMap_L3Iface.Traverse unmatch. count=%d", count)
	}
}
