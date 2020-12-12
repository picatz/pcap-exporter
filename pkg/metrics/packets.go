package metrics

import "github.com/prometheus/client_golang/prometheus"

const SubsystemPackets = "packets"

func init() {
	prometheus.MustRegister(Packets)
	prometheus.MustRegister(PacketsSize)
	prometheus.MustRegister(PacketsLayers)
}

var Packets = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: Namespace,
		Subsystem: SubsystemPackets,
		Name:      "total",
		Help:      "Total number of packets",
	},
)

var PacketsSize = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: Namespace,
		Subsystem: SubsystemPackets,
		Name:      "size",
		Help:      "Total size in bytes for all packets",
	},
)

var PacketsLayers = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: Namespace,
		Subsystem: SubsystemPackets,
		Name:      "layers",
		Help:      "Total layers found for all packets",
	},
)
