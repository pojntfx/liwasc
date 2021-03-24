package main

import (
	"context"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"golang.org/x/sync/semaphore"
)

func ScanTCPPort(targetAddress string, targetPort int, timeout time.Duration) (bool, error) {
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

func ScanUDPPort(targetAddress string, targetPort int, timeout time.Duration, packetGetter func(port int) ([]byte, error)) (bool, error) {
	// Create connection
	con, err := net.Dial("udp", net.JoinHostPort(targetAddress, strconv.Itoa(targetPort)))
	if err != nil {
		return false, err
	}

	// Set timeout
	if err := con.SetDeadline(time.Now().Add(timeout)); err != nil {
		return false, err
	}

	// Get known packet for port
	packet, err := packetGetter(targetPort)
	if err != nil {
		if strings.Contains(err.Error(), "could not find packet for port") {
			packet = []byte{} // Unknown packet for port, use empty []byte{}
		} else {
			return false, err
		}
	}

	// Write packet
	if _, err := con.Write(packet); err != nil {
		return false, err
	}

	// Count every response that is at least 1 byte long as a "open port"
	buffer := make([]byte, 1)
	if _, err := con.Read(buffer); err != nil {
		// Port is closed
		return false, nil
	} else {
		// Port is open
		return true, nil
	}
}

func main() {
	// Arguments
	targetAddress := "127.0.0.1"
	protocols := []string{"tcp", "udp"}
	timeout := time.Millisecond * 500
	ports := 65535
	jobs := 1000

	// Concurrency
	wg := sync.WaitGroup{}
	sem := semaphore.NewWeighted(int64(jobs))
	wg.Add((ports - 1) * len(protocols))

	for port := 1; port < ports; port++ {
		for _, protocol := range protocols {
			go func(targetPort int, targetProtocol string) {
				// Aquire lock
				if err := sem.Acquire(context.Background(), 1); err != nil {
					panic(err)
				}

				// Release lock
				defer sem.Release(1)
				defer wg.Done()

				// Start scan
				for {
					if targetProtocol == "tcp" {
						// Scan TCP
						open, err := ScanTCPPort(targetAddress, targetPort, timeout)
						if err != nil {
							// Re-try
							if err.(net.Error).Timeout() || strings.Contains(err.Error(), "too many open files") {
								time.Sleep(timeout)

								continue
							} else {
								panic(err)
							}
						}

						// Handle scan result
						if open {
							log.Println(targetPort, "tcp/open")
						} else {
							// log.Println(targetPort, "tcp/closed")
						}

						break
					} else if targetProtocol == "udp" {
						// Scan UDP
						open, err := ScanUDPPort(targetAddress, targetPort, timeout, func(port int) ([]byte, error) {
							return []byte("Hello, world!"), nil
						})
						if err != nil {
							// Re-try
							if err.(net.Error).Timeout() || strings.Contains(err.Error(), "too many open files") {
								time.Sleep(timeout)

								continue
							} else {
								panic(err)
							}
						}

						// Handle scan result
						if open {
							log.Println(targetPort, "udp/open")
						} else {
							// log.Println(targetPort, "udp/closed")
						}

						break
					}
				}
			}(port, protocol)
		}
	}

	// Wait till all have finished
	wg.Wait()
}
