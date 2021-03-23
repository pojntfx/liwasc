package main

import (
	"context"
	"log"
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"golang.org/x/sync/semaphore"
)

func ScanPort(targetAddress string, targetPort int, timeout time.Duration) (bool, error) {
	// Get local socket
	raddr := net.ParseIP(targetAddress).To4()
	rport := layers.TCPPort(targetPort)

	// Create connection
	con, err := net.Dial("udp", net.JoinHostPort(targetAddress, strconv.Itoa(targetPort)))
	if err != nil {
		return false, err
	}

	// Get remote socket
	laddr := con.LocalAddr().(*net.UDPAddr)
	lport := layers.TCPPort(laddr.Port)

	// Create IP packet
	outIP := &layers.IPv4{
		SrcIP:    laddr.IP,
		DstIP:    raddr,
		Protocol: layers.IPProtocolTCP,
	}

	// Create TCP segment
	outTCP := &layers.TCP{
		SrcPort: lport,
		DstPort: rport,
		Seq:     1,
		SYN:     true,
		Window:  14600,
	}

	outTCP.SetNetworkLayerForChecksum(outIP)

	// Serialize packet
	outPacket := gopacket.NewSerializeBuffer()
	if err := gopacket.SerializeLayers(
		outPacket,
		gopacket.SerializeOptions{
			ComputeChecksums: true,
			FixLengths:       true,
		},
		outTCP,
	); err != nil {
		return false, err
	}

	// Listen for incoming packets
	conn, err := net.ListenPacket("ip4:tcp", "0.0.0.0")
	if err != nil {
		return false, err
	}
	defer conn.Close()

	// Write packet
	if _, err := conn.WriteTo(
		outPacket.Bytes(),
		&net.IPAddr{
			IP: raddr,
		},
	); err != nil {
		return false, err
	}

	// Set timeout
	if err := conn.SetDeadline(time.Now().Add(timeout)); err != nil {
		return false, err
	}

	for {
		// Receive packet
		buf := make([]byte, 4096)
		n, addr, err := conn.ReadFrom(buf)
		if err != nil {
			return false, err
		}

		// If packet is not intended for our IP, skip it
		if addr.String() != raddr.String() {
			continue
		}

		// If packet is intended for our IP, process it
		inPacket := gopacket.NewPacket(buf[:n], layers.LayerTypeTCP, gopacket.Default)

		// Skip non-TCP packets
		if inTCPLayer := inPacket.Layer(layers.LayerTypeTCP); inTCPLayer != nil {
			inTCP := inTCPLayer.(*layers.TCP)

			// If segment is not intended for our port, skip it
			if inTCP.DstPort != lport {
				continue
			}

			// If SYN and ACK bits are set, the port is open
			if inTCP.SYN && inTCP.ACK {
				return true, nil
			}

			// Port is closed
			return false, nil
		}
	}
}

func main() {
	// Arguments
	targetAddress := "100.64.154.244"
	timeout := time.Millisecond * 500
	ports := 65535
	jobs := 100

	// Concurrency
	wg := sync.WaitGroup{}
	sem := semaphore.NewWeighted(int64(jobs))
	wg.Add(ports - 1)

	for port := 1; port < ports; port++ {
		go func(targetPort int) {
			// Aquire lock
			if err := sem.Acquire(context.Background(), 1); err != nil {
				panic(err)
			}

			// Release log
			defer sem.Release(1)
			defer wg.Done()

			// Start scan
			open, err := ScanPort(targetAddress, targetPort, timeout)
			if err != nil {
				panic(err)
			}

			// Handle scan result
			if open {
				log.Println(targetPort, "open")
			} else {
				// log.Println(targetPort, "closed")
			}
		}(port)
	}

	// Wait till all have finished
	wg.Wait()
}
