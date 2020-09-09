package scanners

// Based on https://github.com/google/gopacket/blob/master/examples/arpscan/arpscan.go

import (
	"bytes"
	"encoding/binary"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/phayes/freeport"
)

type DiscoveredNode struct {
	IPAddress  net.IP
	MACAddress net.HardwareAddr
}

type NetworkScanner struct {
	deviceName         string
	handle             *pcap.Handle
	ipv4addresses      []*net.IPNet
	iface              *net.Interface
	discoveredNodeChan chan *DiscoveredNode
}

func NewNetworkScanner(device string) *NetworkScanner {
	return &NetworkScanner{device, nil, nil, nil, make(chan *DiscoveredNode)}
}

func (s *NetworkScanner) Open() (error, []*net.IPNet) {
	iface, err := net.InterfaceByName(s.deviceName)
	if err != nil {
		return err, nil
	}
	s.iface = iface

	addresses, err := iface.Addrs()
	if err != nil {
		return err, nil
	}

	for _, ipv4address := range addresses {
		ip := ipv4address.(*net.IPNet).IP.To4()

		if ip != nil {
			s.ipv4addresses = append(s.ipv4addresses, &net.IPNet{
				IP:   ip,
				Mask: ipv4address.(*net.IPNet).Mask[len(ipv4address.(*net.IPNet).Mask)-4:],
			})
		}
	}

	port, err := freeport.GetFreePort()
	if err != nil {
		return err, nil
	}

	handle, err := pcap.OpenLive(iface.Name, int32(port), true, pcap.BlockForever)
	if err != nil {
		return err, nil
	}
	s.handle = handle

	return nil, s.ipv4addresses
}

func (s *NetworkScanner) Receive() *DiscoveredNode {
	in := gopacket.NewPacketSource(s.handle, layers.LayerTypeEthernet).Packets()

	for {
		packet := <-in

		layer := packet.Layer(layers.LayerTypeARP)

		// Not an arp packet
		if layer == nil {
			continue
		}

		// Sent by us
		arp := layer.(*layers.ARP)
		if arp.Operation != layers.ARPReply || bytes.Equal([]byte(s.iface.HardwareAddr), arp.SourceHwAddress) {
			continue
		}

		s.discoveredNodeChan <- &DiscoveredNode{net.IP(arp.SourceProtAddress), net.HardwareAddr(arp.SourceHwAddress)}
	}
}

func (s *NetworkScanner) Transmit() error {
	eth := layers.Ethernet{
		SrcMAC:       s.iface.HardwareAddr,
		DstMAC:       net.HardwareAddr{0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, // Broadcast
		EthernetType: layers.EthernetTypeARP,
	}

	for _, ipv4addr := range s.ipv4addresses {
		arp := layers.ARP{
			AddrType:          layers.LinkTypeEthernet,
			Protocol:          layers.EthernetTypeIPv4,
			HwAddressSize:     6,
			ProtAddressSize:   4,
			Operation:         layers.ARPRequest,
			SourceHwAddress:   []byte(s.iface.HardwareAddr),
			SourceProtAddress: []byte(ipv4addr.IP),
			DstHwAddress:      []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, // Broadcast
		}

		buffer := gopacket.NewSerializeBuffer()
		for _, ip := range GetAddressesForNet(ipv4addr) {
			arp.DstProtAddress = []byte(ip)
			gopacket.SerializeLayers(buffer, gopacket.SerializeOptions{
				FixLengths:       true,
				ComputeChecksums: true,
			}, &eth, &arp)

			if err := s.handle.WritePacketData(buffer.Bytes()); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *NetworkScanner) Read() *DiscoveredNode {
	node := <-s.discoveredNodeChan

	return node
}

func GetAddressesForNet(n *net.IPNet) (out []net.IP) {
	num := binary.BigEndian.Uint32([]byte(n.IP))
	mask := binary.BigEndian.Uint32([]byte(n.Mask))
	num &= mask
	for mask < 0xffffffff {
		var buf [4]byte
		binary.BigEndian.PutUint32(buf[:], num)
		out = append(out, net.IP(buf[:]))
		mask++
		num++
	}
	return
}
