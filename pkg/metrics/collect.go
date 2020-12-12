package metrics

import (
	"log"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func defaultIface() string {
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}
	for _, ifc := range ifaces {
		if ifc.HardwareAddr != nil && ifc.Flags&net.FlagUp != 0 && ifc.Flags&net.FlagBroadcast != 0 {
			return ifc.Name
		}
	}
	return ""
}

func Collect(iface, filter string) {
	if iface == "" {
		iface = defaultIface()
	}

	log.Printf("capturing packets on %q", iface)

	// Find all devices
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	var (
		snapshotLen int32 = 65535
		promiscuous bool  = true

		pcapHandle *pcap.Handle
	)
	for _, device := range devices {
		if device.Name == iface {
			pcapHandle, err = pcap.OpenLive(device.Name, snapshotLen, promiscuous, pcap.BlockForever)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	if pcapHandle != nil {
		defer pcapHandle.Close()

		if filter != "" {
			err := pcapHandle.SetBPFFilter(filter)
			if err != nil {
				log.Fatal(err)
			}
		}

		packetSource := gopacket.NewPacketSource(pcapHandle, pcapHandle.LinkType())
		for packet := range packetSource.Packets() {
			Packets.Inc()
			PacketsSize.Add(float64(len(packet.Data())))
			PacketsLayers.Add(float64(len(packet.Layers())))

			for _, layer := range packet.Layers() {
				switch layer.LayerType() {
				case layers.LayerTypeEthernet:
					Ethernet.Inc()
				case layers.LayerTypeIPv4:
					IPv4.Inc()
				case layers.LayerTypeTCP:
					TCP.Inc()
				case layers.LayerTypeTLS:
					TLS.Inc()
				case layers.LayerTypeUDP:
					UDP.Inc()
				case layers.LayerTypeDNS:
					DNS.Inc()
				case layers.LayerTypeICMPv4:
					ICMP.Inc()
				case layers.LayerTypeARP:
					ARP.Inc()
				case layers.LayerTypeIPv6:
					IPv6.Inc()
				case layers.LayerTypeDHCPv4:
					DHCPv4.Inc()
				case layers.LayerTypeNTP:
					NTP.Inc()
				case layers.LayerTypeDot11:
					Dot11.Inc()
				case layers.LayerTypeLoopback:
					Loopback.Inc()
				default:
					Unknown.Inc()
				}
			}
		}
	}
}
