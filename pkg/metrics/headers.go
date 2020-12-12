package metrics

import "github.com/prometheus/client_golang/prometheus"

const SubsystemHeaders = "headers"

func init() {
	prometheus.MustRegister(ICMP)
	prometheus.MustRegister(ARP)
	prometheus.MustRegister(Ethernet)
	prometheus.MustRegister(IPv4)
	prometheus.MustRegister(IPv6)
	prometheus.MustRegister(TCP)
	prometheus.MustRegister(TLS)
	prometheus.MustRegister(UDP)
	prometheus.MustRegister(DNS)
	prometheus.MustRegister(DHCPv4)
	prometheus.MustRegister(Loopback)
	prometheus.MustRegister(NTP)
	prometheus.MustRegister(Dot11)
	prometheus.MustRegister(Unknown)
}

var ICMP = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: Namespace,
		Subsystem: SubsystemHeaders,
		Name:      "icmp",
		Help:      "Total number of ICMP headers",
	},
)

var ARP = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: Namespace,
		Subsystem: SubsystemHeaders,
		Name:      "arp",
		Help:      "Total number of ARP headers",
	},
)

var Ethernet = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: Namespace,
		Subsystem: SubsystemHeaders,
		Name:      "eth",
		Help:      "Total number of ethernet headers",
	},
)

var IPv4 = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: Namespace,
		Subsystem: SubsystemHeaders,
		Name:      "ipv4",
		Help:      "Total number of IPv4 headers",
	},
)

var IPv6 = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: Namespace,
		Subsystem: SubsystemHeaders,
		Name:      "ipv6",
		Help:      "Total number of IPv4 headers",
	},
)

var TCP = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: Namespace,
		Subsystem: SubsystemHeaders,
		Name:      "tcp",
		Help:      "Total number of TCP headers",
	},
)

var TLS = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: Namespace,
		Subsystem: SubsystemHeaders,
		Name:      "tls",
		Help:      "Total number of TLS headers",
	},
)

var UDP = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: Namespace,
		Subsystem: SubsystemHeaders,
		Name:      "udp",
		Help:      "Total number of UDP headers",
	},
)

var DNS = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: Namespace,
		Subsystem: SubsystemHeaders,
		Name:      "dns",
		Help:      "Total number of DNS headers",
	},
)

var DHCPv4 = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: Namespace,
		Subsystem: SubsystemHeaders,
		Name:      "dhcp",
		Help:      "Total number of DHCPv4 headers",
	},
)

var NTP = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: Namespace,
		Subsystem: SubsystemHeaders,
		Name:      "ntp",
		Help:      "Total number of DHCPv4 headers",
	},
)

var Dot11 = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: Namespace,
		Subsystem: SubsystemHeaders,
		Name:      "dot11",
		Help:      "Total number of Dot11 (WiFi) headers",
	},
)

var Loopback = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: Namespace,
		Subsystem: SubsystemHeaders,
		Name:      "loopback",
		Help:      "Total number of loopback headers",
	},
)

var Unknown = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: Namespace,
		Subsystem: SubsystemHeaders,
		Name:      "unknown",
		Help:      "Total number of unknown (possibly unimplemented) headers",
	},
)
