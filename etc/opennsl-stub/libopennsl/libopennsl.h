
#ifndef __LIBOPENNSL_H
#define __LIBOPENNSL_H

#include <opennsl/types.h>
#include <sal/driver.h>
#include <opennsl/rx.h>
#include <opennsl/pkt.h>
#include <opennsl/port.h>
#include <opennsl/l2.h>
#include <opennsl/l3.h>
#include <opennsl/field.h>
#include <opennsl/cosq.h>
#include <opennsl/vxlanX.h>

void strmac(char* buf, const opennsl_mac_t mac_addr);

void strip4(char* buf, const opennsl_ip_t ip);

void strip6(char* buf, const opennsl_ip6_t ip);

void _opennsl_pbmp_dump(const char* name, const opennsl_pbmp_t* pbmp);

void _opennsl_pkt_blk_dump(const char* name, const opennsl_pkt_blk_t* blk);

void _opennsl_pkt_dump(const char* name, const opennsl_pkt_t* pkt);

void _opennsl_rx_cfg_dump(const char* name, const opennsl_rx_cfg_t* cfg);

void _opennsl_port_ability_dump(const char* name, const opennsl_port_ability_t* abil);

void _opennsl_port_info_dump(const char* name, const opennsl_port_info_t* info);

void _opennsl_l3_intf_qos_dump(const char* name, const opennsl_l3_intf_qos_t* qos);

void _opennsl_l3_intf_dump(const char* name, const opennsl_l3_intf_t* iface);

void _opennsl_l3_egress_dump(const char* name, const opennsl_l3_egress_t* egr);

void _opennsl_l3_egress_ecmp_dump(const char* name, const opennsl_l3_egress_ecmp_t* ecmp);

void _opennsl_l3_egress_ecmp_member_dump(const char* name, const opennsl_l3_ecmp_member_t* ecmp);

void _opennsl_l3_ingress_dump(const char* name, const opennsl_l3_ingress_t* igr);

void _opennsl_l3_route_dump(const char* name, const opennsl_l3_route_t* route);

void _opennsl_l3_host_dump(const char* name, const opennsl_l3_host_t* host);

void _opennsl_l2_addr_dump(const char* name, const opennsl_l2_addr_t *l2addr);

void _opennsl_l2_cache_addr_dump(const char* name, const opennsl_l2_cache_addr_t* addr);

void _opennsl_field_bitscl_dump(const char* name, const char* field, const uint32* arr, int num);

void _opennsl_field_action_width_dump(const char* name, const opennsl_field_action_width_t* w);

void _opennsl_field_qset_dump(const char* name, const opennsl_field_qset_t* qset);

void _opennsl_field_aset_dump(const char* name, const opennsl_field_aset_t* aset);

void _opennsl_field_presel_set_dump(const char* name, const opennsl_field_presel_set_t* ps);

void _opennsl_switch_pkt_info_dump(const char* name, const opennsl_switch_pkt_info_t* info);

void _opennsl_mpls_egress_label_dump(const char* name, const opennsl_mpls_egress_label_t* label);

void _opennsl_mpls_tunnel_switch_dump(const char* name, const opennsl_mpls_tunnel_switch_t* info);

void _opennsl_mpls_vpn_config_dump(const char* name, const opennsl_mpls_vpn_config_t* info);

void _opennsl_vlan_port_dump(const char* name, const opennsl_vlan_port_t* port);

void _opennsl_init_dump(const char* name, const opennsl_init_t* init);

void _opennsl_cosq_bst_profile_dump(const char* name, const opennsl_cosq_bst_profile_t* profile);

void _opennsl_vlan_protocol_packet_ctrl_dump(const char* name, const opennsl_vlan_protocol_packet_ctrl_t* c);
void _opennsl_vxlan_vpn_config_dump(const char* name, const opennsl_vxlan_vpn_config_t* config);

void _opennsl_vxlan_port_dump(const char* name, const opennsl_vxlan_port_t* p);

void _opennsl_tunnel_initiator_dump(const char* name, const opennsl_tunnel_initiator_t* tun);

void _opennsl_tunnel_terminator_dump(const char* name, const opennsl_tunnel_terminator_t* tun);

#endif // __LIBOPENNSL_H
