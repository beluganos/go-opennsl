// -*- coding: utf-8 -*-

#include <stdio.h>
#include <opennsl/types.h>
#include "logger.h"
#include "libopennsl.h"

void strmac(char* buf, const opennsl_mac_t mac_addr) {
  sprintf(buf, "%02x:%02x:%02x:%02x:%02x:%02x",
      mac_addr[0], mac_addr[1], mac_addr[2],
      mac_addr[3], mac_addr[4], mac_addr[5]);
}

void strip4(char* buf, const opennsl_ip_t ip) {
  sprintf(buf, "%hhu.%hhu.%hhu.%hhu",
	  (((uint32)ip) >> 24) & 0xff,
	  (((uint32)ip) >> 16) & 0xff,
	  (((uint32)ip) >>  8) & 0xff,
	  ((uint32)ip) & 0xff);
}

void strip6(char* buf, const opennsl_ip6_t ip) {
  sprintf(buf, "%02x%02x:%02x%02x:%02x%02x:%02x%02x:%02x%02x:%02x%02x:%02x%02x:%02x%02x",
	  ip[0],  ip[1],  ip[2],  ip[3],
	  ip[4],  ip[5],  ip[6],  ip[7],
	  ip[8],  ip[9],  ip[10], ip[11],
	  ip[12], ip[13], ip[14], ip[15]);
}


void _opennsl_pbmp_dump(const char* name, const opennsl_pbmp_t* pbmp) {
  char buf[64];
  char* p = buf;
  int i;
  for (i = 0; i < _SHR_PBMP_WORD_MAX; i++) {
    const int write_num = sprintf(p, "%08x ", pbmp->pbits[i]);
    if ((i % 4) == 3) {
      LOG_DEBUG("%s: pbmp: %s", name, buf);
      p = buf;
    } else {
      p += write_num;
    }
  }

  if (p != buf) {
    LOG_DEBUG("%s: pbmp: %s", name, buf);
  }
}

void _opennsl_pkt_blk_dump(const char* name, const opennsl_pkt_blk_t* blk) {
  if (blk == NULL) {
    LOG_DEBUG("%s: opennsl_pkt_blk = %p", name, blk);
    return;
  }

  LOG_DEBUG("%s: opennsl_pkt_blk.data = %p", name, blk->data);
  LOG_DEBUG("%s: opennsl_pkt_blk.len  = %d", name, blk->len);

  int index;
  char buf[64];
  char* p = buf;
  for (index = 0; index < blk->len; index++) {
    p += sprintf(p, "%02x ", blk->data[index]);

    if ((index % 16) == 15) {
      LOG_DEBUG("%s: opennsl_pkt_blk.data = %s", name, buf);
      p = buf;
    }
  }

  if (p != buf)
    LOG_DEBUG("%s: opennsl_pkt_blk.data = %s", name, buf);

}

void _opennsl_pkt_dump(const char* name, const opennsl_pkt_t* pkt) {
  int index;

  if (pkt == NULL) {
    LOG_DEBUG("%s: opennsl_pkt = %p", name, pkt);
    return;
  }

  LOG_DEBUG("%s: opennsl_pkt.blk_count  = %d", name, pkt->blk_count);
  for (index = 0; index < pkt->blk_count; index++) {
    _opennsl_pkt_blk_dump(name, pkt->pkt_data + index);
  }

  _opennsl_pbmp_dump(name, &pkt->tx_pbmp);
}

void _opennsl_port_ability_dump(const char* name, const opennsl_port_ability_t* abil) {
  LOG_DEBUG("%s: port_ability.speed_half_duplex   = %u", name, abil->speed_half_duplex);
  LOG_DEBUG("%s: port_ability.speed_full_duplex   = %u", name, abil->speed_full_duplex);
  LOG_DEBUG("%s: port_ability.pause     = %u", name, abil->pause);
  LOG_DEBUG("%s: port_ability.interface = %u", name, abil->interface);
  LOG_DEBUG("%s: port_ability.medium    = %u", name, abil->medium);
  LOG_DEBUG("%s: port_ability.loopback  = %u", name, abil->loopback);
  LOG_DEBUG("%s: port_ability.flags     = %u", name, abil->flags);
  LOG_DEBUG("%s: port_ability.eee       = %u", name, abil->eee);
  LOG_DEBUG("%s: port_ability.rsvd      = %u", name, abil->rsvd);
  LOG_DEBUG("%s: port_ability.encap     = %u", name, abil->encap);
  LOG_DEBUG("%s: port_ability.fec       = %u", name, abil->fec);
  LOG_DEBUG("%s: port_ability.channel   = %u", name, abil->channel);
}

void _opennsl_port_info_dump(const char* name, const opennsl_port_info_t* info) {
  char pause_mac[32];
  strmac(pause_mac, info->pause_mac);

  LOG_DEBUG("%s: port_info.action_mask   = %u", name, info->action_mask);
  LOG_DEBUG("%s: port_info.action_mask2  = %u", name, info->action_mask2);
  LOG_DEBUG("%s: port_info.enable        = %d", name, info->enable);
  LOG_DEBUG("%s: port_info.linkstatus    = %d", name, info->linkstatus);
  LOG_DEBUG("%s: port_info.autoneg       = %d", name, info->autoneg);
  LOG_DEBUG("%s: port_info.speed         = %d", name, info->speed);
  LOG_DEBUG("%s: port_info.duplex        = %d", name, info->duplex);
  LOG_DEBUG("%s: port_info.linkscan      = %d", name, info->linkscan);
  LOG_DEBUG("%s: port_info.learn         = %u", name, info->learn);
  LOG_DEBUG("%s: port_info.discard       = %d", name, info->discard);
  LOG_DEBUG("%s: port_info.vlanfilter    = %u", name, info->vlanfilter);
  LOG_DEBUG("%s: port_info.untagged_pri  = %d", name, info->untagged_priority);
  LOG_DEBUG("%s: port_info.untagged_vlan = %hu", name, info->untagged_vlan);
  LOG_DEBUG("%s: port_info.stp_state     = %d", name, info->stp_state);
  LOG_DEBUG("%s: port_info.pfm           = %d", name, info->pfm);
  LOG_DEBUG("%s: port_info.loopback      = %d", name, info->loopback);
  LOG_DEBUG("%s: port_info.phy_master    = %d", name, info->phy_master);
  LOG_DEBUG("%s: port_info.interface     = %d", name, info->interface);
  LOG_DEBUG("%s: port_info.pause_tx      = %d", name, info->pause_tx);
  LOG_DEBUG("%s: port_info.pause_rx      = %d", name, info->pause_rx);
  LOG_DEBUG("%s: port_info.encap_mode    = %d", name, info->encap_mode);
  LOG_DEBUG("%s: port_info.pause_mac     = '%s'", name, pause_mac);
  LOG_DEBUG("%s: port_info.local_advert  = %u", name, info->local_advert);
  _opennsl_port_ability_dump(name, &info->local_ability);
  LOG_DEBUG("%s: port_info.remote_advert_valid = %u", name, info->remote_advert_valid);
  LOG_DEBUG("%s: port_info.remote_advert = %u", name, info->remote_advert);
  _opennsl_port_ability_dump(name, &info->remote_ability);
  LOG_DEBUG("%s: port_info.mcast_limit   = %d", name, info->mcast_limit);
  LOG_DEBUG("%s: port_info.mcast_limit_enable = %d", name, info->mcast_limit_enable);
  LOG_DEBUG("%s: port_info.bcast_limit   = %d", name, info->bcast_limit);
  LOG_DEBUG("%s: port_info.bcast_limit_enable = %d", name, info->bcast_limit_enable);
  LOG_DEBUG("%s: port_info.dlfbc_limit   = %d", name, info->dlfbc_limit);
  LOG_DEBUG("%s: port_info.dlfbc_limit_enable = %d", name, info->dlfbc_limit_enable);
  LOG_DEBUG("%s: port_info.speed_max     = %d", name, info->speed_max);
  LOG_DEBUG("%s: port_info.ability       = %u", name, info->ability);
  _opennsl_port_ability_dump(name, &info->port_ability);
  LOG_DEBUG("%s: port_info.frame_max     = %d", name, info->frame_max);
  LOG_DEBUG("%s: port_info.mdix          = %u", name, info->mdix);
  LOG_DEBUG("%s: port_info.mdix_status   = %u", name, info->mdix_status);
  LOG_DEBUG("%s: port_info.medium        = %u", name, info->medium);
  LOG_DEBUG("%s: port_info.fault         = %u", name, info->fault);
}

void _opennsl_l3_intf_qos_dump(const char* name, const opennsl_l3_intf_qos_t* qos) {
  LOG_DEBUG("%s: l3_intf_qos.flags      = %08x", name, qos->flags);
  LOG_DEBUG("%s: l3_intf_qos.qos_map_id = %d",   name, qos->qos_map_id);
  LOG_DEBUG("%s: l3_intf_qos.pri        = %hhu", name, qos->pri);
  LOG_DEBUG("%s: l3_intf_qos.cfi        = %hhu", name, qos->cfi);
  LOG_DEBUG("%s: l3_intf_qos.dscp       = %d",   name, qos->dscp);
}

void _opennsl_l3_intf_dump(const char* name, const opennsl_l3_intf_t* iface) {
  char mac[32];
  strmac(mac, iface->l3a_mac_addr);

  LOG_DEBUG("%s: l3_intf.flags    = %08x", name, iface->l3a_flags);
  LOG_DEBUG("%s: l3_intf.vrf      = %d", name, iface->l3a_vrf);
  LOG_DEBUG("%s: l3_intf.intf_id  = %d", name, iface->l3a_intf_id);
  LOG_DEBUG("%s: l3_intf.vrf      = %d", name, iface->l3a_vrf);
  LOG_DEBUG("%s: l3_intf.mac      = %s", name, mac);
  LOG_DEBUG("%s: l3_intf.vid      = %hu", name, iface->l3a_vid);
  LOG_DEBUG("%s: l3_intf.ttl      = %d", name, iface->l3a_ttl);
  LOG_DEBUG("%s: l3_intf.mtu      = %d", name, iface->l3a_mtu);
  LOG_DEBUG("%s: l3_intf.mtu_fwd  = %d", name, iface->l3a_mtu_forwarding);
  _opennsl_l3_intf_qos_dump(name, &iface->dscp_qos);
  LOG_DEBUG("%s: l3_intf.ipv4_prof= %d", name, iface->l3a_ip4_options_profile_id);
  LOG_DEBUG("%s: l3_intf.n_rt_vid = %hhu", name, iface->native_routing_vlan_tags);
}

void _opennsl_l3_egress_dump(const char* name, const opennsl_l3_egress_t* egr) {
  char mac[32];
  strmac(mac, egr->mac_addr);

  LOG_DEBUG("%s: l3_egress.flags    = %08x", name, egr->flags);
  LOG_DEBUG("%s: l3_egress.flags2   = %08x", name, egr->flags2);
  LOG_DEBUG("%s: l3_egress.intf     = %08x", name, egr->intf);
  LOG_DEBUG("%s: l3_egress.mac      = %s", name, mac);
  LOG_DEBUG("%s: l3_egress.vlan     = %hu", name, egr->vlan);
  LOG_DEBUG("%s: l3_egress.fmodule  = %d", name, egr->module);
  LOG_DEBUG("%s: l3_egress.port     = %d", name, egr->port);
  LOG_DEBUG("%s: l3_egress.trunk     = %d", name, egr->trunk);
}

void _opennsl_l3_egress_ecmp_dump(const char* name, const opennsl_l3_egress_ecmp_t* ecmp) {
  LOG_DEBUG("%s: l3_egress_ecmp.flags        = %08x", name, ecmp->flags);
  LOG_DEBUG("%s: l3_egress_ecmp.intf         = %d", name, ecmp->ecmp_intf);
  LOG_DEBUG("%s: l3_egress_ecmp.max_paths    = %d", name, ecmp->max_paths);
  LOG_DEBUG("%s: l3_egress_ecmp.dynamic_mode = %d", name, ecmp->dynamic_mode);
  LOG_DEBUG("%s: l3_egress_ecmp.dynamic_size = %d", name, ecmp->dynamic_size);
}

void _opennsl_l3_egress_ecmp_member_dump(const char* name, const opennsl_l3_ecmp_member_t* ecmp) {
  LOG_DEBUG("%s: l3_egress_ecmp_member.flags     = %08x", name, ecmp->flags);
  LOG_DEBUG("%s: l3_egress_ecmp_member.egress_if = %d", name, ecmp->egress_if);
  LOG_DEBUG("%s: l3_egress_ecmp.member.status    = %d", name, ecmp->status);
}

void _opennsl_l3_ingress_dump(const char* name, const opennsl_l3_ingress_t* igr) {
  LOG_DEBUG("%s: l3_ingress.flags      = %08x", name, igr->flags);
  LOG_DEBUG("%s: l3_ingress.vrf        = %d", name, igr->vrf);
  LOG_DEBUG("%s: l3_ingress.urpf_mode  = %d", name, igr->urpf_mode);
  LOG_DEBUG("%s: l3_ingress.intf_class = %d", name, igr->intf_class);
  LOG_DEBUG("%s: l3_ingress.ipmc_intf  = %d", name, igr->ipmc_intf_id);
  LOG_DEBUG("%s: l3_ingress.qos_map    = %d", name, igr->qos_map_id);
  LOG_DEBUG("%s: l3_ingress.ip4_profile= %d", name, igr->ip4_options_profile_id);
  LOG_DEBUG("%s: l3_ingress.nat_realm  = %d", name, igr->nat_realm_id);
  LOG_DEBUG("%s: l3_ingress.tun_trm_ecn= %d", name, igr->tunnel_term_ecn_map_id);
  LOG_DEBUG("%s: l3_ingress.ifclass_rt = %d", name, igr->intf_class_route_disable);
}

void _opennsl_l3_route_dump(const char* name, const opennsl_l3_route_t* route) {
  char ip4net[32];
  char ip4mask[32];
  char ip6net[64];
  char ip6mask[64];

  strip4(ip4net, route->l3a_subnet);
  strip4(ip4mask, route->l3a_ip_mask);
  strip6(ip6net, route->l3a_ip6_net);
  strip6(ip6mask, route->l3a_ip6_mask);

  LOG_DEBUG("%s: l3_route.flags    = %08x", name, route->l3a_flags);
  LOG_DEBUG("%s: l3_route.vrf      = %d", name, route->l3a_vrf);
  LOG_DEBUG("%s: l3_route.ipv4     = %s/%s", name, ip4net, ip4mask);
  LOG_DEBUG("%s: l3_route.ipv6     = %s/%s", name, ip6net, ip6mask);
  LOG_DEBUG("%s: l3_route.iface    = %d/%08x", name, route->l3a_intf, route->l3a_intf);
  LOG_DEBUG("%s: l3_route.port_tgid= %d", name, route->l3a_port_tgid);
  LOG_DEBUG("%s: l3_route.pri      = %d", name, route->l3a_pri);
}

void _opennsl_l3_host_dump(const char* name, const opennsl_l3_host_t* host) {
  char mac[32];
  strmac(mac, host->l3a_nexthop_mac);

  char ip[32];
  strip4(ip, host->l3a_ip_addr);

  char ip6[64];
  strip6(ip6, host->l3a_ip6_addr);

  LOG_DEBUG("%s: l3_host.flags    = %08x", name, host->l3a_flags);
  LOG_DEBUG("%s: l3_host.vrf      = %d", name, host->l3a_vrf);
  LOG_DEBUG("%s: l3_host.ip_addr  = %s", name, ip);
  LOG_DEBUG("%s: l3_host.ip6_addr = %s", name, ip6);
  LOG_DEBUG("%s: l3_host.pri      = %d", name, host->l3a_pri);
  LOG_DEBUG("%s: l3_host.intf     = %d/%08x", name, host->l3a_intf, host->l3a_intf);
  LOG_DEBUG("%s: l3_host.nexthop  = %s", name, mac);
  LOG_DEBUG("%s: l3_host.port_tgid= %d", name, host->l3a_port_tgid);
}

void _opennsl_l2_addr_dump(const char* name, const opennsl_l2_addr_t *l2addr) {
  char mac[32];
  strmac(mac, l2addr->mac);

  LOG_DEBUG("%s: l2_addr.flags = %08x", name,  l2addr->flags);
  LOG_DEBUG("%s: l2_addr.mac   = %s", name, mac);
  LOG_DEBUG("%s: l2_addr.vid   = %d", name, l2addr->vid);
  LOG_DEBUG("%s: l2_addr.port  = %d", name, l2addr->port);
  LOG_DEBUG("%s: l2_addr.modid = %d", name, l2addr->modid);
  LOG_DEBUG("%s: l2_addr.tgid  = %d", name, l2addr->tgid);
  LOG_DEBUG("%s: l2_addr.l2mcg = %d", name, l2addr->l2mc_group);
}

void _opennsl_l2_cache_addr_dump(const char* name, const opennsl_l2_cache_addr_t* addr) {
  int i;

  char mac[32];
  strmac(mac, addr->mac);
  char mac_mask[32];
  strmac(mac_mask, addr->mac_mask);

  LOG_DEBUG("%s: l2_cache_addr.flags         = %08x", name,  addr->flags);
  LOG_DEBUG("%s: l2_cache_addr.station_flags = %08x", name,  addr->station_flags);
  LOG_DEBUG("%s: l2_cache_addr.mac           = %s/%s", name, mac, mac_mask);
  LOG_DEBUG("%s: l2_cache_addr.vlan          = %d/%d", name, addr->vlan, addr->vlan_mask);
  LOG_DEBUG("%s: l2_cache_addr.src_port      = %d/%d", name, addr->src_port, addr->src_port_mask);
  LOG_DEBUG("%s: l2_cache_addr.dest_modid    = %d", name,  addr->dest_modid);
  LOG_DEBUG("%s: l2_cache_addr.dest_port     = %d", name,  addr->dest_port);
  LOG_DEBUG("%s: l2_cache_addr.dest_trunk    = %d", name,  addr->dest_trunk);
  LOG_DEBUG("%s: l2_cache_addr.prio          = %d", name,  addr->prio);
  for (i = 0; i < _SHR_PBMP_WORD_MAX; i++) {
    LOG_DEBUG("%s: l2_cache_addr.dest_ports[%d]= %u", name, i, addr->dest_ports.pbits[i]);
  }
  LOG_DEBUG("%s: l2_cache_addr.lookup_class  = %d", name,  addr->lookup_class);
  LOG_DEBUG("%s: l2_cache_addr.subtype       = %hhu", name,  addr->subtype);
  LOG_DEBUG("%s: l2_cache_addr.encap_id      = %d", name,  addr->encap_id);
  LOG_DEBUG("%s: l2_cache_addr.group         = %d", name,  addr->group);
  LOG_DEBUG("%s: l2_cache_addr.ethertype     = %hu/%hu", name,  addr->ethertype, addr->ethertype_mask);
}

void _opennsl_field_bitscl_dump(const char* name, const char* field, const uint32* arr, int num) {
  char temp[64];
  int pos = 0;
  int index;
  for (index = 0; index < num; index++) {
    pos += sprintf(temp + pos, "%08x ", arr[index]);
    if ((index % 4) == 3) {
      LOG_DEBUG("%s: %s = %s", name, field, temp);
      pos = 0;
    }
  }

  if (index != 0) {
    LOG_DEBUG("%s: %s = %s", name, field, temp);
  }
}

void _opennsl_field_action_width_dump(const char* name, const opennsl_field_action_width_t* w) {
  LOG_DEBUG("%s: opennsl_field_action_width.action = %u", name, w->action);
  LOG_DEBUG("%s: opennsl_field_action_width.width  = %u", name, w->width);
  LOG_DEBUG("%s: opennsl_field_action_width.valid  = %hhu", name, w->valid);
  
}

void _opennsl_field_qset_dump(const char* name, const opennsl_field_qset_t* qset) {
  _opennsl_field_bitscl_dump(name, "opennsl_field_qset.w",
			     qset->w, _SHR_BITDCLSIZE(OPENNSL_FIELD_QUALIFY_MAX));
  _opennsl_field_bitscl_dump(name, "opennsl_field_qset.udf_map",
			     qset->udf_map, _SHR_BITDCLSIZE(OPENNSL_FIELD_USER_NUM_UDFS));
}

void _opennsl_field_aset_dump(const char* name, const opennsl_field_aset_t* aset) {
  _opennsl_field_bitscl_dump(name, "opennsl_field_aset.w",
			     aset->w, _SHR_BITDCLSIZE(opennslFieldActionCount));
  int index;
  const int action_num = sizeof(aset->actions_width) / sizeof(opennsl_field_action_width_t);
  for (index = 0; index < action_num; index++) {
    _opennsl_field_action_width_dump(name, aset->actions_width + index);
  }
}

void _opennsl_field_presel_set_dump(const char* name, const opennsl_field_presel_set_t* ps) {
  _opennsl_field_bitscl_dump(name, "opennsl_field_presel_set.w",
			     ps->w, _SHR_BITDCLSIZE(OPENNSL_FIELD_PRESEL_SEL_MAX));
}

void _opennsl_switch_pkt_info_dump(const char* name, const opennsl_switch_pkt_info_t* info) {
  char src_mac[32];
  strmac(src_mac, info->src_mac);
  char dst_mac[32];
  strmac(dst_mac, info->dst_mac);

  char sip4[32];
  strip4(sip4, info->sip);
  char dip4[32];
  strip4(dip4, info->dip);

  char sip6[64];
  strip6(sip6, info->sip6);
  char dip6[64];
  strip6(dip6, info->dip6);

  LOG_DEBUG("%s: switch_pkt_info.flags     = %d", __func__, info->flags);
  LOG_DEBUG("%s: switch_pkt_info.src_gport = %d", __func__, info->src_gport);
  LOG_DEBUG("%s: switch_pkt_info.vid       = %hu", __func__, info->vid);
  LOG_DEBUG("%s: switch_pkt_info.ethertype = %hu", __func__, info->ethertype);
  LOG_DEBUG("%s: switch_pkt_info.dst_mac   = %s", __func__, dst_mac);
  LOG_DEBUG("%s: switch_pkt_info.src_mac   = %s", __func__, src_mac);
  LOG_DEBUG("%s: switch_pkt_info.dip       = %s", __func__, dip4);
  LOG_DEBUG("%s: switch_pkt_info.sip       = %s", __func__, sip4);
  LOG_DEBUG("%s: switch_pkt_info.dip6      = %s", __func__, dip6);
  LOG_DEBUG("%s: switch_pkt_info.sip6      = %s", __func__, sip6);
  LOG_DEBUG("%s: switch_pkt_info.protocol  = %04x", __func__, info->protocol);
  LOG_DEBUG("%s: switch_pkt_info.dst_l4port= %u", __func__, info->dst_l4_port);
  LOG_DEBUG("%s: switch_pkt_info.src_l3port= %u", __func__, info->src_l4_port);
  LOG_DEBUG("%s: switch_pkt_info.trnk_gport= %u", __func__, info->trunk_gport);
  LOG_DEBUG("%s: switch_pkt_info.mpintf    = %d", __func__, info->mpintf);
  LOG_DEBUG("%s: switch_pkt_info.fwd_reason= %d", __func__, info->fwd_reason);
}

void _opennsl_mpls_egress_label_dump(const char* name, const opennsl_mpls_egress_label_t* label) {
  LOG_DEBUG("%s: opennsl_mpls_egress_label.flags       = %08x", name, label->flags);
  LOG_DEBUG("%s: opennsl_mpls_egress_label.label       = %u", name, label->label);
  LOG_DEBUG("%s: opennsl_mpls_egress_label.qos_map_id  = %d", name, label->qos_map_id);
  LOG_DEBUG("%s: opennsl_mpls_egress_label.exp         = %d", name, label->exp);
  LOG_DEBUG("%s: opennsl_mpls_egress_label.ttl         = %d", name, label->ttl);
  LOG_DEBUG("%s: opennsl_mpls_egress_label.pkt_pri     = %hhu", name, label->pkt_pri);
  LOG_DEBUG("%s: opennsl_mpls_egress_label.pkt_cfi     = %hhu", name, label->pkt_cfi);
  LOG_DEBUG("%s: opennsl_mpls_egress_label.tunnel_id   = %d", name, label->tunnel_id);
  LOG_DEBUG("%s: opennsl_mpls_egress_label.l3_intf_id  = %d", name, label->l3_intf_id);
  LOG_DEBUG("%s: opennsl_mpls_egress_label.action      = %d", name, label->action);
  LOG_DEBUG("%s: opennsl_mpls_egress_label.egr_F.O._id = %d", name, label->egress_failover_id);
  LOG_DEBUG("%s: opennsl_mpls_egress_label.egr_F.O._if = %d", name, label->egress_failover_if_id);
  LOG_DEBUG("%s: opennsl_mpls_egress_label.outlif_cnt_prof = %d", name, label->outlif_counting_profile);
}

void _opennsl_mpls_tunnel_switch_dump(const char* name, const opennsl_mpls_tunnel_switch_t* info) {
  LOG_DEBUG("%s: opennsl_mpls_tunnel_switch.flags       = %08x", name, info->flags);
  LOG_DEBUG("%s: opennsl_mpls_tunnel_switch.label       = %u", name, info->label);
  LOG_DEBUG("%s: opennsl_mpls_tunnel_switch.label2      = %u", name, info->second_label);
  LOG_DEBUG("%s: opennsl_mpls_tunnel_switch.port        = %d", name, info->port);
  LOG_DEBUG("%s: opennsl_mpls_tunnel_switch.action      = %d", name, info->action);
  LOG_DEBUG("%s: opennsl_mpls_tunnel_switch.action( bos)= %d", name, info->action_if_bos);
  LOG_DEBUG("%s: opennsl_mpls_tunnel_switch.action(!bos)= %d", name, info->action_if_not_bos);
  LOG_DEBUG("%s: opennsl_mpls_tunnel_switch.mc_group    = %d", name, info->mc_group);
  LOG_DEBUG("%s: opennsl_mpls_tunnel_switch.exp_map     = %d", name, info->exp_map);
  LOG_DEBUG("%s: opennsl_mpls_tunnel_switch.int_pri     = %d", name, info->int_pri);
  LOG_DEBUG("%s: opennsl_mpls_tunnel_switch.policer_id  = %d", name, info->policer_id);
  LOG_DEBUG("%s: opennsl_mpls_tunnel_switch.vpn         = %hu", name, info->vpn);
  _opennsl_mpls_egress_label_dump(name, &info->egress_label);
  LOG_DEBUG("%s: opennsl_mpls_tunnel_switch.egress_if   = %d", name, info->egress_if);
  LOG_DEBUG("%s: opennsl_mpls_tunnel_switch.ingress_if  = %d", name, info->ingress_if);
  LOG_DEBUG("%s: opennsl_mpls_tunnel_switch.mtu         = %d", name, info->mtu);
  LOG_DEBUG("%s: opennsl_mpls_tunnel_switch.qos_map_id  = %d", name, info->qos_map_id);
  LOG_DEBUG("%s: opennsl_mpls_tunnel_switch.failover_id = %d", name, info->failover_id);
  LOG_DEBUG("%s: opennsl_mpls_tunnel_switch.tunnel_id   = %d", name, info->tunnel_id);
  LOG_DEBUG("%s: opennsl_mpls_tunnel_switch.F.O.tunnelid= %d", name, info->failover_tunnel_id);
  LOG_DEBUG("%s: opennsl_mpls_tunnel_switch.tunnel_if   = %d", name, info->tunnel_if);
  LOG_DEBUG("%s: opennsl_mpls_tunnel_switch.egress_port = %d", name, info->egress_port);
  LOG_DEBUG("%s: opennsl_mpls_tunnel_switch.oam_ctxt_id = %d", name, info->oam_global_context_id);
  LOG_DEBUG("%s: opennsl_mpls_tunnel_switch.class_id    = %d", name, info->class_id);
  LOG_DEBUG("%s: opennsl_mpls_tunnel_switch.inlif_cnt_prof = %d", name, info->inlif_counting_profile);
}

void _opennsl_mpls_vpn_config_dump(const char* name, const opennsl_mpls_vpn_config_t* info) {
  LOG_DEBUG("%s: opennsl_mpls_vpn_config.flags       = %08x", name, info->flags);
  LOG_DEBUG("%s: opennsl_mpls_vpn_config.vpn         = %hu", name, info->vpn);
  LOG_DEBUG("%s: opennsl_mpls_vpn_config.lookup_id   = %d", name, info->lookup_id);
  LOG_DEBUG("%s: opennsl_mpls_vpn_config.b.c.group   = %d", name, info->broadcast_group);
  LOG_DEBUG("%s: opennsl_mpls_vpn_config.unknown.u.c = %d", name, info->unknown_unicast_group);
  LOG_DEBUG("%s: opennsl_mpls_vpn_config.unknown.m.c = %d", name, info->unknown_multicast_group);
  LOG_DEBUG("%s: opennsl_mpls_vpn_config.policer_id  = %d", name, info->policer_id);
  // LOG_DEBUG("%s: opennsl_mpls_vpn_config.protocol_pkt= %d", name, info->protocol_pkt);
}

void _opennsl_vlan_port_dump(const char* name, const opennsl_vlan_port_t* port) {
  LOG_DEBUG("%s: vlan_port.criteria         = %d", name, port->criteria);
  LOG_DEBUG("%s: vlan_port.vsi              = %d", name, port->vsi);
  LOG_DEBUG("%s: vlan_port.match_vlan       = %hu", name, port->match_vlan);
  LOG_DEBUG("%s: vlan_port.match_inner_vlan = %hu", name, port->match_inner_vlan);
  LOG_DEBUG("%s: vlan_port.port             = %d", name, port->port);
  LOG_DEBUG("%s: vlan_port.egress_vlan      = %hu", name, port->egress_vlan);
  LOG_DEBUG("%s: vlan_port.vlan_port_id     = %d", name, port->vlan_port_id);
}

void _opennsl_init_dump(const char* name, const opennsl_init_t* init) {
  if (init == NULL) {
    LOG_DEBUG("%s: init = NULL", name);
    return;
  }

  if (init->cfg_fname != NULL) {
    LOG_DEBUG("%s: init.cfg_name = %s", name,  init->cfg_fname);
  } else {
    LOG_DEBUG("%s: init.cfg_name = NULL", name);
  }

  LOG_DEBUG("%s: init.flags    = %08x", name,  init->flags);

  if (init->wb_fname != NULL) {
    LOG_DEBUG("%s: init.wb_fname = %s", name,  init->wb_fname);
  } else {
    LOG_DEBUG("%s: init.wb_fname = NULL", name);
  }

  if (init->rmcfg_fname != NULL) {
    LOG_DEBUG("%s: init.rmcfg_fname = %s", name,  init->rmcfg_fname);
  } else {
    LOG_DEBUG("%s: init.rmcfg_fname = NULL", name);
  }

  if (init->cfg_post_fname != NULL) {
    LOG_DEBUG("%s: init.cfg_post_fname = %s", name,  init->cfg_post_fname);
  } else {
    LOG_DEBUG("%s: init.cfg_post_fname = NULL", name);
  }

  LOG_DEBUG("%s: init.opennsl_flags = %08x", name,  init->opennsl_flags);
}

void _opennsl_rx_cfg_dump(const char* name, const opennsl_rx_cfg_t* cfg) {
  // TODO:
}

void _opennsl_cosq_bst_profile_dump(const char* name, const opennsl_cosq_bst_profile_t* p) {
  LOG_DEBUG("%s: cosq_bst_profile.byte = %d (%08x)", name, p->byte, p->byte);
}

void _opennsl_vlan_protocol_packet_ctrl_dump(const char* name, const opennsl_vlan_protocol_packet_ctrl_t* c) {
  LOG_DEBUG("%s: vlan_protocol_packet_ctrl.mmrp_action        = %d", name, c->mmrp_action);
  LOG_DEBUG("%s: vlan_protocol_packet_ctrl.srp_action         = %d", name, c->srp_action);
  LOG_DEBUG("%s: vlan_protocol_packet_ctrl.arp_reply_action   = %d", name, c->arp_reply_action);
  LOG_DEBUG("%s: vlan_protocol_packet_ctrl.arp_request_action = %d", name, c->arp_request_action);
  LOG_DEBUG("%s: vlan_protocol_packet_ctrl.nd_action          = %d", name, c->nd_action);
  LOG_DEBUG("%s: vlan_protocol_packet_ctrl.dhcp_action        = %d", name, c->dhcp_action);
  LOG_DEBUG("%s: vlan_protocol_packet_ctrl.igmp_report_leave_action = %d", name, c->igmp_report_leave_action);
  LOG_DEBUG("%s: vlan_protocol_packet_ctrl.igmp_query_action        = %d", name, c->igmp_query_action);
  LOG_DEBUG("%s: vlan_protocol_packet_ctrl.igmp_unknown_msg_action  = %d", name, c->igmp_unknown_msg_action);
  LOG_DEBUG("%s: vlan_protocol_packet_ctrl.mld_report_done_action   = %d", name, c->mld_report_done_action);
  LOG_DEBUG("%s: vlan_protocol_packet_ctrl.mld_query_action         = %d", name, c->mld_query_action);
  LOG_DEBUG("%s: vlan_protocol_packet_ctrl.ip4_rsvd_mc_action = %d", name, c->ip4_rsvd_mc_action);
  LOG_DEBUG("%s: vlan_protocol_packet_ctrl.ip6_rsvd_mc_action = %d", name, c->ip6_rsvd_mc_action);
}

void _opennsl_vxlan_vpn_config_dump(const char* name, const opennsl_vxlan_vpn_config_t* c) {
  LOG_DEBUG("%s: vxlan_vpn_config.flags          = %08x", name, c->flags);
  LOG_DEBUG("%s: vxlan_vpn_config.vpn            = %hu", name, c->vpn);
  LOG_DEBUG("%s: vxlan_vpn_config.vnid           = %u", name, c->vnid);
  LOG_DEBUG("%s: vxlan_vpn_config.pkt_pri        = %hhu", name, c->pkt_pri);
  LOG_DEBUG("%s: vxlan_vpn_config.pkt_cfi        = %hhu", name, c->pkt_cfi);
  LOG_DEBUG("%s: vxlan_vpn_config.eg_svc_tpid    = %hu", name, c->egress_service_tpid);
  LOG_DEBUG("%s: vxlan_vpn_config.eg_svc_vlan    = %hu", name, c->egress_service_vlan);
  LOG_DEBUG("%s: vxlan_vpn_config.bc_group       = %d", name, c->broadcast_group);
  LOG_DEBUG("%s: vxlan_vpn_config.unknown_uc_grp = %d", name, c->unknown_unicast_group);
  LOG_DEBUG("%s: vxlan_vpn_config.unknown_mc_grp = %d", name, c->unknown_multicast_group);
  _opennsl_vlan_protocol_packet_ctrl_dump(name, &c->protocol_pkt);
  LOG_DEBUG("%s: vxlan_vpn_config.vlan           = %hu", name, c->vlan);
  LOG_DEBUG("%s: vxlan_vpn_config.match_port_cls = %d", name, c->match_port_class);
  LOG_DEBUG("%s: vxlan_vpn_config.default_vlan   = %hu", name, c->default_vlan);
}

void _opennsl_vxlan_port_dump(const char* name, const opennsl_vxlan_port_t* p) {
  LOG_DEBUG("%s: vxlan_port.vxlan_port_id    = %d", name, p->vxlan_port_id);
  LOG_DEBUG("%s: vxlan_port.flags            = %08x", name, p->flags);
  LOG_DEBUG("%s: vxlan_port.int_pri          = %hu", name, p->int_pri);
  LOG_DEBUG("%s: vxlan_port.pkt_pri          = %hhu", name, p->pkt_pri);
  LOG_DEBUG("%s: vxlan_port.pkt_cfi          = %hhu", name, p->pkt_cfi);
  LOG_DEBUG("%s: vxlan_port.eg_svc_tpid      = %hu", name, p->egress_service_tpid);
  LOG_DEBUG("%s: vxlan_port.eg_svc_vlan      = %hu", name, p->egress_service_vlan);
  LOG_DEBUG("%s: vxlan_port.mtu              = %hu", name, p->mtu);
  LOG_DEBUG("%s: vxlan_port.match_port       = %d", name, p->match_port);
  LOG_DEBUG("%s: vxlan_port.criteria         = %d", name, p->criteria);
  LOG_DEBUG("%s: vxlan_port.match_vlan       = %d", name, p->match_vlan);
  LOG_DEBUG("%s: vxlan_port.match_inner_vlan = %hu", name, p->match_inner_vlan);
  LOG_DEBUG("%s: vxlan_port.eg_tun_id        = %d", name, p->egress_tunnel_id);
  LOG_DEBUG("%s: vxlan_port.match_tun_id     = %d", name, p->match_tunnel_id);
  LOG_DEBUG("%s: vxlan_port.egress_if        = %d", name, p->egress_if);
  LOG_DEBUG("%s: vxlan_port.nw_group_id      = %d", name, p->network_group_id);
}

void _opennsl_tunnel_initiator_dump(const char* name, const opennsl_tunnel_initiator_t* tun) {
  char dmac[32], smac[32];
  char dip[32], sip[32];
  char dip6[64], sip6[64];
  strmac(dmac, tun->dmac);
  strmac(smac, tun->smac);
  strip4(sip, tun->sip);
  strip4(dip, tun->dip);
  strip6(sip6, tun->sip6);
  strip6(dip6, tun->dip6);

  LOG_DEBUG("%s: tun_init.flags        = %08x", name, tun->flags);
  LOG_DEBUG("%s: tun_init.type         = %d", name, tun->type);
  LOG_DEBUG("%s: tun_init.l3_intf_id   = %d", name, tun->l3_intf_id);

  LOG_DEBUG("%s: tun_init.dmac         = %s", name, dmac);
  LOG_DEBUG("%s: tun_init.smac         = %s", name, smac);
  LOG_DEBUG("%s: tun_init.dip          = %s", name, dip);
  LOG_DEBUG("%s: tun_init.sip          = %s", name, sip);
  LOG_DEBUG("%s: tun_init.dip6         = %s", name, dip6);
  LOG_DEBUG("%s: tun_init.sip6         = %s", name, sip6);
  LOG_DEBUG("%s: tun_init.udp_dst_port = %hu", name, tun->udp_dst_port);
  LOG_DEBUG("%s: tun_init.udp_src_port = %hu", name, tun->udp_src_port);

  LOG_DEBUG("%s: tun_init.vlan         = %hu", name, tun->vlan);
  LOG_DEBUG("%s: tun_init.ttl          = %d", name, tun->ttl);
  LOG_DEBUG("%s: tun_init.mtu          = %d", name, tun->mtu);
  LOG_DEBUG("%s: tun_init.tunnel_id    = %d", name, tun->tunnel_id);

  LOG_DEBUG("%s: tun_init.aux_data     = %d", name, tun->aux_data);
  LOG_DEBUG("%s: tun_init.dscp         = %d", name, tun->dscp);
  LOG_DEBUG("%s: tun_init.dscp_map     = %d", name, tun->dscp_map);
  LOG_DEBUG("%s: tun_init.dscp_sel     = %d", name, tun->dscp_sel);
  LOG_DEBUG("%s: tun_init.flow_label   = %u", name, tun->flow_label);
  LOG_DEBUG("%s: tun_init.ip4_id       = %hu", name, tun->ip4_id);
  LOG_DEBUG("%s: tun_init.olif_cnt_pfl = %d", name, tun->outlif_counting_profile);
  LOG_DEBUG("%s: tun_init.pkt_cfi      = %hhu", name, tun->pkt_cfi);
  LOG_DEBUG("%s: tun_init.pkt_pri      = %hhu", name, tun->pkt_pri);
  LOG_DEBUG("%s: tun_init.span_id      = %hu", name, tun->span_id);
  LOG_DEBUG("%s: tun_init.tpid         = %hu", name, tun->tpid);
}

void _opennsl_tunnel_terminator_dump(const char* name, const opennsl_tunnel_terminator_t* tun) {
  char dip[32], dip_mask[32], sip[32], sip_mask[32];
  char dip6[64], dip6_mask[64], sip6[64], sip6_mask[64];
  strip4(sip, tun->sip);
  strip4(sip_mask, tun->sip_mask);
  strip4(dip, tun->dip);
  strip4(dip_mask, tun->dip_mask);
  strip6(sip6, tun->sip6);
  strip6(sip6_mask, tun->sip6_mask);
  strip6(dip6, tun->dip6);
  strip6(dip6_mask, tun->dip6_mask);

  LOG_DEBUG("%s: tun_init.flags        = %08x", name, tun->flags);
  LOG_DEBUG("%s: tun_init.type         = %d", name, tun->type);

  LOG_DEBUG("%s: tun_init.dip          = %s/%s", name, dip, dip_mask);
  LOG_DEBUG("%s: tun_init.sip          = %s/%s", name, sip, sip_mask);
  LOG_DEBUG("%s: tun_init.dip6         = %s/%s", name, dip6, dip6_mask);
  LOG_DEBUG("%s: tun_init.sip6         = %s/%s", name, sip6, sip6_mask);
  LOG_DEBUG("%s: tun_init.udp_dst_port = %u", name, tun->udp_dst_port);
  LOG_DEBUG("%s: tun_init.udp_src_port = %u", name, tun->udp_src_port);
  LOG_DEBUG("%s: tun_init.remote_port  = %d", name, tun->remote_port);

  LOG_DEBUG("%s: tun_init.vlan         = %hu", name, tun->vlan);
  LOG_DEBUG("%s: tun_init.vrf          = %hu", name, tun->vrf);
  LOG_DEBUG("%s: tun_init.tunnel_id    = %d", name, tun->tunnel_id);
  LOG_DEBUG("%s: tun_init.tunnel_if    = %d", name, tun->tunnel_if);

  LOG_DEBUG("%s: tun_init.ilif_cnt_pfl = %d", name, tun->inlif_counting_profile);
  LOG_DEBUG("%s: tun_init.mcact_flag   = %08x", name, tun->multicast_flag);
  LOG_DEBUG("%s: tun_init.qos_map_id   = %d", name, tun->qos_map_id);
  _opennsl_pbmp_dump(name, &tun->pbmp);
}
