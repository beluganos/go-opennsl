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

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sal/driver.h>
#include <opennsl/error.h>
#include <opennsl/vlan.h>
#include <opennsl/switch.h>
#include <opennsl/l2.h>
#include <opennsl/l3.h>
#include <opennsl/tunnelX.h>
#include <examples/util.h>

#define DEFAULT_VLAN (1)

int example_l2_addr_add(int unit, opennsl_mac_t mac, int port, int vid) {

  int rv;
  opennsl_l2_addr_t l2addr;

  opennsl_l2_addr_t_init(&l2addr, mac, vid);
  l2addr.port  = port;
  l2addr.flags = (OPENNSL_L2_L3LOOKUP | OPENNSL_L2_STATIC);
  
  rv = opennsl_l2_addr_add(unit, &l2addr);
  if (rv != OPENNSL_E_NONE) {
    return rv;
  }

  return OPENNSL_E_NONE;
}

int tunnel_initiator(int unit,
		     opennsl_l3_intf_t* l3_intf,
		     opennsl_mac_t tun_dst_mac,
		     opennsl_ip_t tun_dst_ip,
		     opennsl_ip_t tun_src_ip,
		     opennsl_gport_t* tun_id) {
  opennsl_tunnel_initiator_t tunnel;
  opennsl_tunnel_initiator_t_init(&tunnel);
  tunnel.ttl = 64;
  memcpy(tunnel.dmac, tun_dst_mac, sizeof(opennsl_mac_t));
  tunnel.dip = tun_dst_ip;
  tunnel.sip = tun_src_ip;
  tunnel.vlan = DEFAULT_VLAN;
  tunnel.l3_intf_id = l3_intf->l3a_intf_id;

  int rc;

  rc = opennsl_tunnel_initiator_set(unit, l3_intf, &tunnel);
  if (rc != OPENNSL_E_NONE) {
    printf ("opennsl_tunnel_initiator_set failed. Return Code: %s \n",
        opennsl_errmsg(rc));
    return rc;
  }

  *tun_id = tunnel.tunnel_id;

  return rc;
}

int example_create_l3_intf(int unit, opennsl_gport_t port, opennsl_vlan_t vid,
			   opennsl_mac_t mac_addr, opennsl_l3_intf_t *l3_intf) {

  int rc;

  /* Create L3 interface */
  opennsl_l3_intf_t_init(l3_intf);
  memcpy(l3_intf->l3a_mac_addr, mac_addr, 6);
  l3_intf->l3a_vid = vid;
  rc = opennsl_l3_intf_create(unit, l3_intf);
  if (rc != OPENNSL_E_NONE) {
    printf("l3_setup: opennsl_l3_intf_create failed: %s\n", opennsl_errmsg(rc));
    return rc;
  }

  printf("L3 interface is created with parameters: \n  VLAN %d \n", vid);
  l2_print_mac("  MAC Address: ", mac_addr);
  printf("\n\r");
  printf("  L3 Interface ID: %d\r\n", l3_intf->l3a_intf_id);

  return rc;
}

int example_create_l3_egress(int unit, unsigned int flags, int out_port, int vlan,
			     int l3_eg_intf, opennsl_mac_t next_hop_mac_addr,
			     int *intf) {
  int rc;
  opennsl_l3_egress_t l3eg;
  opennsl_if_t l3egid;
  int mod = 0;

  opennsl_l3_egress_t_init(&l3eg);
  l3eg.intf = l3_eg_intf;
  memcpy(l3eg.mac_addr, next_hop_mac_addr, 6);

  l3eg.vlan   = vlan;
  l3eg.module = mod;
  l3eg.port   = out_port;

  l3egid = *intf;

  rc = opennsl_l3_egress_create(unit, flags, &l3eg, &l3egid);
  if (rc != OPENNSL_E_NONE) {
    return rc;
  }

  *intf = l3egid;

  printf("Created L3 egress ID %d for out_port: %d vlan: %d "
	 "L3 egress intf: %d\n",
	 *intf, out_port, vlan, l3_eg_intf);

  return rc;
}

int example_add_host(int unit, unsigned int addr, int intf) {
  int rc;
  opennsl_l3_host_t l3host;
  opennsl_l3_host_t_init(&l3host);

  l3host.l3a_flags = 0;
  l3host.l3a_ip_addr = addr;
  l3host.l3a_intf = intf;
  l3host.l3a_port_tgid = 0;

  rc = opennsl_l3_host_add(unit, &l3host);
  if (rc != OPENNSL_E_NONE) {
    printf ("opennsl_l3_host_add failed. Return Code: %s \n",
        opennsl_errmsg(rc));
    return rc;
  }

  print_ip_addr("add host ", addr);
  printf(" ---> egress-object = %d\n", intf);

  return rc;
}

int main(int argc, char** argv) {
  int rv;

  int unit = 0;
  int in_sysport = 50;
  int out_sysport = 51;
  opennsl_l3_intf_t l3_intf_out;
  int l3_egr_id;
  int flags;
  int host = 0x14010102; /* 20.1.1.2 */
  opennsl_mac_t my_mac = {0x00, 0x11, 0x22, 0x33, 0x99, 0x58};
  opennsl_mac_t nh_mac = {0x00, 0x00, 0x70, 0x5B, 0xC7, 0x34};
  opennsl_mac_t tun_dst_mac = {0x00, 0x11, 0x22, 0x33, 0x99, 0x60};
  int tun_dst_ip = 0x0a010102; // 10.1.1.2
  int tun_src_ip = 0x0a010101; // 10.1.1.1
  int tun_id;

  rv = opennsl_driver_init((opennsl_init_t *) NULL);
  if(rv != 0) {
    printf("\r\nFailed to initialize the system.\r\n");
    return rv;
  }

  printf("opennsl_driver_init ok. ^\n");

  rv = example_port_default_config(unit);
  if (rv != OPENNSL_E_NONE) {
    printf("\r\nFailed to apply default config on ports, rc = %d (%s).\r\n",
	   rv, opennsl_errmsg(rv));
    return EXIT_FAILURE;
  }

  printf("example_port_default_config ok \n");

  rv = example_switch_default_vlan_config(unit);
  if (rv != OPENNSL_E_NONE) {
    printf("\r\nFailed to apply default vlan config on ports, rc = %d (%s).\r\n",
           rv, opennsl_errmsg(rv));
    return EXIT_FAILURE;
  }

  printf("example_switch_default_vlan_config ok.");

  /* Set L3 Egress Mode */
  rv =  opennsl_switch_control_set(unit, opennslSwitchL3EgressMode, 1);
  if (rv != OPENNSL_E_NONE) {
    return EXIT_FAILURE;
  }
  printf("\nL3 Egress mode is set succesfully\n");

  /*** create egress router interface ***/
  rv = example_create_l3_intf(unit, out_sysport, DEFAULT_VLAN,
			      my_mac, &l3_intf_out);
  if (rv != OPENNSL_E_NONE) {
    printf("Error, create egress interface-1, out_sysport=%d. "
	   "Return code %s \n", out_sysport, opennsl_errmsg(rv));
    return EXIT_FAILURE;
  }

  // tunnel initiator
  rv = tunnel_initiator(unit, &l3_intf_out, tun_dst_mac, tun_dst_ip, tun_src_ip, &tun_id);
  if (rv != OPENNSL_E_NONE) {
    printf("Error. create tunnel initiator.\n");
    return EXIT_FAILURE;
  }

  /*** Make the address learn on a VLAN and port */
  rv = example_l2_addr_add(unit, my_mac, in_sysport, DEFAULT_VLAN);
  if (rv != OPENNSL_E_NONE) {
    printf("Failed to add L2 address. Return Code: %s\n", opennsl_errmsg(rv));
    return EXIT_FAILURE;
  }

  rv = example_l2_addr_add(unit, my_mac, out_sysport, DEFAULT_VLAN);
  if (rv != OPENNSL_E_NONE) {
    printf("Failed to add L2 address. Return Code: %s\n", opennsl_errmsg(rv));
    return EXIT_FAILURE;
  }

  /*** create egress object 1 ***/
  flags = 0;
  rv = example_create_l3_egress(unit, flags, out_sysport, DEFAULT_VLAN, l3_intf_out.l3a_intf_id,
				nh_mac, &l3_egr_id);
  if (rv != OPENNSL_E_NONE) {
    printf("Error, create egress object, out_sysport=%d. Return code %s \n",
	   out_sysport, opennsl_errmsg(rv));
    return EXIT_FAILURE;
  }

  rv = example_add_host(unit, host, l3_egr_id);
  if (rv != OPENNSL_E_NONE) {
    printf("Error, host add. Return code %s\n",
	   opennsl_errmsg(rv));
    return EXIT_FAILURE;
  }

  return EXIT_SUCCESS;
}
