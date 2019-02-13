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

package main

import (
	"net"

	"github.com/beluganos/go-opennsl/examples/util"
	"github.com/beluganos/go-opennsl/opennsl"
	"github.com/beluganos/go-opennsl/sal"

	log "github.com/sirupsen/logrus"
)

func setupVxlan(unit int) error {
	var err error

	vpn_id_1 := opennsl.Vpn(0x7010)
	vpn_id_2 := opennsl.Vpn(0x7020)
	vnid_1 := opennsl.VNID(0x12345)
	vnid_2 := opennsl.VNID(0x53535)

	net_vpn := opennsl.Vpn(0x7001)
	net_vnid := opennsl.VNID(0xeeee)

	acc_dummy_mac, _ := net.ParseMAC("00:00:01:00:00:01")

	acc_port_1 := opennsl.Port(52)
	acc_vid_1 := opennsl.Vlan(21)

	acc_port_2 := opennsl.Port(53)
	acc_vid_2 := opennsl.Vlan(20)

	net_port := opennsl.Port(50)
	net_vid := opennsl.Vlan(22)

	net_local_mac, _ := net.ParseMAC("00:00:00:00:22:22")
	net_remote_mac, _ := net.ParseMAC("00:00:00:00:00:02")
	tnl_local_ip := net.ParseIP("10.10.10.1")
	tnl_remote_ip := net.ParseIP("192.168.1.1")

	dlf_mac, _ := net.ParseMAC("01:00:5e:00:00:0a")
	tnl_mc_dip := net.ParseIP("224.0.0.10")

	payload_sa_1, _ := net.ParseMAC("00:00:00:00:11:aa")
	payload_da_1, _ := net.ParseMAC("00:00:00:00:11:bb")

	payload_sa_2, _ := net.ParseMAC("00:00:00:00:22:aa")
	payload_da_2, _ := net.ParseMAC("00:00:00:00:22:bb")

	if err = do_vxlan_global_setting(unit); err != nil {
		log.Errorf("do_vxlan_global_setting error. %s", err)
		return err
	}

	if err = do_vxlan_access_port_settings(unit, acc_port_1); err != nil {
		log.Errorf("do_vxlan_access_port_settings(%d) error. %s", acc_port_1, err)
		return err
	}

	if err = do_vxlan_access_port_settings(unit, acc_port_2); err != nil {
		log.Errorf("do_vxlan_access_port_settings(%d) error. %s", acc_port_2, err)
		return err
	}

	if err = do_vxlan_net_port_settings(unit, net_port); err != nil {
		log.Errorf("do_vxlan_net_port_settings(%d) error. %s", net_port, err)
		return err
	}

	acc_gport_1, _ := acc_port_1.GPortGet(unit)
	acc_gport_2, _ := acc_port_2.GPortGet(unit)
	net_gport, _ := net_port.GPortGet(unit)

	bc_gr_1, err := opennsl.MULTICAST.Create(unit, opennsl.MULTICAST_TYPE_VXLAN)
	if err != nil {
		log.Errorf("MulticastCreate(1) error. %s", err)
		return err
	}

	bc_gr_2, err := opennsl.MULTICAST.Create(unit, opennsl.MULTICAST_TYPE_VXLAN)
	if err != nil {
		log.Errorf("MulticastCreate(2) error. %s", err)
		return err
	}

	bc_gr_n, err := opennsl.MULTICAST.Create(unit, opennsl.MULTICAST_TYPE_VXLAN)
	if err != nil {
		log.Errorf("MulticastCreate(n) error. %s", err)
		return err
	}

	if err = create_vxlan_vpn(unit, vpn_id_1, vnid_1, bc_gr_1); err != nil {
		log.Errorf("create_vxlan_vpn(%d, %d, %d) error. %s", vpn_id_1, vnid_1, bc_gr_1, err)
		return err
	}

	if err = create_vxlan_vpn(unit, vpn_id_2, vnid_2, bc_gr_2); err != nil {
		log.Errorf("create_vxlan_vpn(%d, %d, %d) error. %s", vpn_id_2, vnid_2, bc_gr_2, err)
		return err
	}

	if err = create_vxlan_vpn(unit, net_vpn, net_vnid, bc_gr_n); err != nil {
		log.Errorf("create_vxlan_vpn(%d, %d, %d) error. %s", net_vpn, net_vnid, bc_gr_n, err)
		return err
	}

	//
	// Create access side VXLAN port 1 (belongs to VPN 1)
	//
	if err = vlan_create_add_port(unit, acc_vid_1, acc_port_1); err != nil {
		log.Errorf("vlan_create_add_port(%d, %d) error. %s", acc_vid_1, acc_port_1, err)
		return err
	}

	acc_intf_id_1, err := create_l3_interface(unit, acc_dummy_mac, acc_vid_1)
	if err != nil {
		log.Errorf("acc_vid_1,(%s, %d) error. %s", acc_dummy_mac, acc_vid_1, err)
		return err
	}

	acc_egr_obj_1, err := create_egr_obj(unit, acc_intf_id_1, acc_dummy_mac, acc_port_1, acc_vid_1)
	if err != nil {
		log.Errorf("create_egr_obj(%d, %s, %d, %d) error. %s", acc_intf_id_1, acc_dummy_mac, acc_port_1, acc_vid_1, err)
		return err
	}

	acc_vxlan_port_1, err := create_vxlan_acc_vp(unit, vpn_id_1, acc_gport_1, opennsl.VXLAN_PORT_MATCH_PORT, acc_egr_obj_1, acc_vid_1)
	if err != nil {
		log.Errorf("create_vxlan_acc_vp(%d, %d, %d, %d) error. %s", vpn_id_1, acc_gport_1, acc_egr_obj_1, acc_vid_1, err)
		return err
	}

	log.Debugf("acc_vxlan_port_1 = %d", acc_vxlan_port_1)

	//
	// Create access side VXLAN port 2 (belongs to VPN 2)
	//
	if err = vlan_create_add_port(unit, acc_vid_2, acc_port_2); err != nil {
		log.Errorf("vlan_create_add_port(%d, %d) error. %s", acc_vid_2, acc_port_2, err)
		return err
	}

	acc_intf_id_2, err := create_l3_interface(unit, acc_dummy_mac, acc_vid_2)
	if err != nil {
		log.Errorf("acc_dummy_mac(%s, %d) error. %s", acc_dummy_mac, acc_vid_2, err)
		return err
	}

	acc_egr_obj_2, err := create_egr_obj(unit, acc_intf_id_2, acc_dummy_mac, acc_port_2, acc_vid_2)
	if err != nil {
		log.Errorf("create_egr_obj(%d, %s, %d, %d) error. %s", acc_intf_id_2, acc_dummy_mac, acc_port_2, acc_vid_2, err)
		return err
	}

	acc_vxlan_port_2, err := create_vxlan_acc_vp(unit, vpn_id_2, acc_gport_2, opennsl.VXLAN_PORT_MATCH_PORT, acc_egr_obj_2, acc_vid_2)
	if err != nil {
		log.Errorf("create_vxlan_acc_vp(%d, %d, %d, %d) error. %s", vpn_id_2, acc_gport_2, acc_egr_obj_2, acc_vid_2, err)
		return err
	}

	log.Debugf("acc_vxlan_port_2 = %d", acc_vxlan_port_2)

	//
	// The network tunnel is shared by the two VPNs
	//
	if err = vlan_create_add_port(unit, net_vid, net_port); err != nil {
		log.Errorf("vlan_create_add_port(%d, %d) error. %s", net_vid, net_port, err)
		return err
	}

	net_intf_id, err := create_l3_interface(unit, net_local_mac, net_vid)
	if err != nil {
		log.Errorf("create_l3_interface(%s, %d) error. %s", net_local_mac, net_vid, err)
		return err
	}

	net_egr_obj, err := create_egr_obj(unit, net_intf_id, net_remote_mac, net_port, net_vid)
	if err != nil {
		log.Errorf("create_egr_obj%d, %s, %d, %d) error. %s", net_intf_id, net_remote_mac, net_port, net_vid, err)
		return err
	}

	tunnel_init_id, err := tunnel_initiator_setup(unit, tnl_local_ip, tnl_remote_ip)
	if err != nil {
		log.Errorf("tunnel_initiator_setup(%s, %s) error. %s", tnl_local_ip, tnl_remote_ip, err)
		return err
	}

	tunnel_term_id, err := tunnel_terminator_setup(unit, tnl_remote_ip, tnl_local_ip, opennsl.VLAN_ID_INVALID, tunnel_init_id)
	if err != nil {
		log.Errorf("tunnel_terminator_setup(%s, %s, %d) error. %s", tnl_remote_ip, tnl_local_ip, tunnel_init_id, err)
		return err
	}

	net_vxlan_port, err := create_vxlan_net_vp(unit, net_vpn, net_gport, opennsl.VXLAN_PORT_MATCH_VN_ID, net_egr_obj, tunnel_init_id, tunnel_term_id)
	if err != nil {
		log.Errorf("create_vxlan_net_vp(%d, %d, %d, %d, %d) error. %s", net_vpn, net_gport, net_egr_obj, tunnel_init_id, tunnel_term_id, err)
		return err
	}

	log.Debugf("net_vxlan_port = %d", net_vxlan_port)

	if _, err = add_to_l2_station(unit, net_local_mac, net_vid); err != nil {
		log.Errorf("add_to_l2_station() error. %s", err)
	}

	//
	// DLF/BC network port set up
	//
	egr_obj_mc, err := create_egr_obj(unit, net_intf_id, dlf_mac, net_port, net_vid, opennsl.L3_IPMC)
	if err != nil {
		log.Errorf("create_egr_obj(%d, %s, %d, %d) error. %s", net_intf_id, dlf_mac, net_port, net_vid, err)
		return err
	}

	tunnel_mc_init_id, err := tunnel_initiator_setup(unit, tnl_local_ip, tnl_mc_dip)
	if err != nil {
		log.Errorf("tunnel_initiator_setup(%s, %s) error. %s", tnl_local_ip, tnl_mc_dip, err)
		return err
	}

	tunnel_mc_term_id, err := tunnel_terminator_setup(unit, tnl_remote_ip, tnl_mc_dip, net_vid, tunnel_mc_init_id)
	if err != nil {
		log.Errorf("tunnel_terminator_setup(%s, %s, %d) error. %s", tnl_remote_ip, tnl_mc_dip, tunnel_mc_init_id, err)
		return err
	}

	vxlan_port_mc, err := create_vxlan_net_vp(unit, net_vpn, net_gport, opennsl.VXLAN_PORT_MATCH_NONE, egr_obj_mc, tunnel_mc_init_id, tunnel_mc_term_id, opennsl.VXLAN_PORT_MULTICAST)
	if err != nil {
		log.Errorf("create_vxlan_net_vp(%d, %d, %d, %d, %d) error. %s", net_vpn, net_gport, egr_obj_mc, tunnel_mc_init_id, tunnel_mc_term_id, err)
		return err
	}

	log.Debugf("vxlan_port_mc = %x", vxlan_port_mc)

	if _, err = add_to_l2_station(unit, dlf_mac, net_vid); err != nil {
		log.Errorf("add_to_l2_station() error. %s", err)
	}

	//
	//  MC group set up - MC group should contains all Access ports and Network non-UC port
	//
	encap_id_mc1, err := bc_gr_1.VxlanEncapGet(unit, net_gport, vxlan_port_mc)
	if err != nil {
		log.Errorf("VxlanEncapGet error. %d %d %d %s", bc_gr_1, net_gport, vxlan_port_mc, err)
		return err
	}
	if err = bc_gr_1.EgressAdd(unit, vxlan_port_mc, encap_id_mc1); err != nil {
		log.Errorf("EgressAdd error. %d %d %d %s", bc_gr_1, vxlan_port_mc, encap_id_mc1, err)
		return err
	}

	acc_encap_id_1, err := bc_gr_1.VxlanEncapGet(unit, acc_gport_1, acc_vxlan_port_1)
	if err != nil {
		log.Errorf("VxlanEncapGet error. %d %d %d %s", bc_gr_1, acc_gport_1, acc_vxlan_port_1, err)
		return err
	}
	if err = bc_gr_1.EgressAdd(unit, acc_vxlan_port_1, acc_encap_id_1); err != nil {
		log.Errorf("EgressAdd error. %d %d %d %s", bc_gr_1, acc_vxlan_port_1, acc_encap_id_1, err)
		return err
	}

	encap_id_mc2, err := bc_gr_2.VxlanEncapGet(unit, net_gport, vxlan_port_mc)
	if err != nil {
		log.Errorf("VxlanEncapGet error. %d %d %d %s", bc_gr_2, net_gport, vxlan_port_mc, err)
		return err
	}
	if err = bc_gr_2.EgressAdd(unit, vxlan_port_mc, encap_id_mc2); err != nil {
		log.Errorf("EgressAdd error. %d %d %d %s", bc_gr_2, vxlan_port_mc, encap_id_mc2, err)
		return err
	}

	acc_encap_id_2, err := bc_gr_2.VxlanEncapGet(unit, acc_gport_2, acc_vxlan_port_2)
	if err != nil {
		log.Errorf("VxlanEncapGet error. %d %d %d %s", bc_gr_2, acc_gport_2, acc_vxlan_port_2, err)
		return err
	}
	if err = bc_gr_2.EgressAdd(unit, acc_vxlan_port_2, acc_encap_id_2); err != nil {
		log.Errorf("EgressAdd error. %d %d %d %s", bc_gr_2, acc_vxlan_port_2, acc_encap_id_2, err)
		return err
	}

	//
	// Add Payload L2 address to L2 table
	//
	if err = add_to_l2_table(unit, payload_da_1, acc_vid_1, net_vxlan_port); err != nil {
		log.Errorf("add_to_l2_table() error. %s", err)
		return err
	}
	if err = add_to_l2_table(unit, payload_sa_1, acc_vid_1, acc_vxlan_port_1); err != nil {
		log.Errorf("add_to_l2_table() error. %s", err)
		return err
	}
	if err = add_to_l2_table(unit, payload_da_2, acc_vid_2, net_vxlan_port); err != nil {
		log.Errorf("add_to_l2_table() error. %s", err)
		return err
	}
	if err = add_to_l2_table(unit, payload_sa_2, acc_vid_2, acc_vxlan_port_2); err != nil {
		log.Errorf("add_to_l2_table() error. %s", err)
		return err
	}

	return nil
}

func main() {
	log.SetLevel(log.DebugLevel)

	unit := int(0)

	if err := sal.DriverInit(); err != nil {
		log.Errorf("driver_init error. %s", err)
		return
	}
	defer sal.DriverExit()

	if err := util.PortDefaultConfig(unit); err != nil {
		log.Errorf("PortDefaultConfig. %s", err)
		return
	}

	if err := setupVxlan(unit); err != nil {
		log.Errorf("setupVxlan error. %s", err)
		return
	}

	done := make(chan struct{})
	go util.WatchSignal(done)

	<-done

}
