# pcap-exporter

 🦈 Prometheus exporter for pcap metrics

## Usage

```console
$ go run main.go
2020/12/11 21:38:59 starting server
2020/12/11 21:38:59 capturing packets on "en0"
...
```

```console
$ curl localhost:9249/metrics | grep "pcap" | grep -v "#"
pcap_headers_arp 2
pcap_headers_dhcp 0
pcap_headers_dns 0
pcap_headers_dot11 0
pcap_headers_eth 23
pcap_headers_icmp 0
pcap_headers_ipv4 21
pcap_headers_ipv6 0
pcap_headers_loopback 0
pcap_headers_ntp 0
pcap_headers_tcp 12
pcap_headers_tls 0
pcap_headers_udp 9
pcap_headers_unknown 9
pcap_packets_layers 76
pcap_packets_size 2888
pcap_packets_total 23
```
