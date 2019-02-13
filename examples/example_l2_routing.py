#! /usr/bin/env python
# -*- coding: utf-8 -*-

from scapy.all import *

# DST_MAC = "00:11:22:33:44:55"
DST_MAC = "00:16:3e:0c:53:93"
SRC_MAC = "00:00:70:5b:c7:34"

DST_IPV4 = "55.0.0.1"
SRC_IPV4 = "10.0.0.1"
DST_IPV6 = "2001:db8:100::1"
SRC_IPV6 = "2001:db8:100::2"
DST_TUN6 = "2010:2020::1"
SRC_TUN6 = "2010:2010::4"

VLAN_VID = 10
IFACES = ["enp131s0f0", "enp131s0f1", "enp132s0f0", "enp132s0f1"]

class Counter:
    CNT = 0
    @classmethod
    def get(cls):
        cls.CNT += 1
        return cls.CNT

def make_pkt(dst_ip, data):
    udp = UDP(sport=5051, dport=5050) / data
    pkt = Ether(dst=DST_MAC, src=SRC_MAC)/IP(dst=dst_ip, src=SRC_IPV4)/udp
    pkt.show2()
    return pkt


def make_pkt6(dst_ip, data):
    udp = UDP(sport=5051, dport=5050) / data
    pkt = Ether(dst=DST_MAC, src=SRC_MAC)/IPv6(dst=dst_ip, src=SRC_IPV6)/udp
    pkt.show2()
    return pkt

def make_tun6(dst_ip, data):
    udp = UDP(sport=5051, dport=5050) / data
    ip4 = IP(dst=dst_ip, src=SRC_IPV4)/udp
    pkt =  Ether(dst=DST_MAC, src=SRC_MAC)/IPv6(dst=DST_TUN6, src=SRC_TUN6)/ip4
    pkt.show2()
    return pkt

def send_ipv4(dst_ip):
    for iface in IFACES:
        pkt = make_pkt(dst_ip, "{0}/{1}".format(iface, Counter.get()))
        sendp(pkt, iface=iface)


def send_ipv6(dst_ip):
    for iface in IFACES:
        pkt = make_pkt6(dst_ip, "{0}/{1}".format(iface, Counter.get()))
        sendp(pkt, iface=iface)


def send_tun6(dst_ip):
    for iface in IFACES:
        pkt = make_tun6(dst_ip, "{0}/{1}".format(iface, Counter.get()))
        sendp(pkt, iface=iface)


def _main():
    send_ipv4("20.1.1.2")
    #send_ipv4("100.100.100.10")
    #send_ipv4("200.200.200.20")
    send_tun6("100.100.100.10")
    send_tun6("200.200.200.20")
    # send_ipv6()

if __name__ == "__main__":
    _main()
