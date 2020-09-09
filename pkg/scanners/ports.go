package scanners

import (
	"encoding/base64"
	"fmt"
	"net"
	"sync"
	"time"
)

type ScannedPort struct {
	Target   string
	Port     int
	Protocol string
	Open     bool
}

type PortScanner struct {
	target          string
	startPort       int
	endPort         int
	timeout         time.Duration
	protocols       []string
	scannedPortChan chan *ScannedPort
}

func NewPortScanner(target string, startPort int, endPort int, timeout time.Duration, protocols []string) *PortScanner {
	return &PortScanner{target, startPort, endPort, timeout, protocols, make(chan *ScannedPort)}
}

func (s *PortScanner) Transmit() error {
	doneChan := make(chan struct{})
	fatalErrorChan := make(chan error)

	var wg sync.WaitGroup
	for port := s.startPort; port <= s.endPort; port++ {
		for _, protocol := range s.protocols {
			wg.Add(1)

			go func(innerPort int, innerProtocol string) {
				conn, err := net.DialTimeout(innerProtocol, net.JoinHostPort(s.target, fmt.Sprintf("%v", innerPort)), s.timeout)
				if err != nil {
					// TODO: Re-try if too many open files here, this depends on a unlimited ulimit atm
					s.scannedPortChan <- &ScannedPort{s.target, innerPort, innerProtocol, false}
				}

				if conn != nil {
					if innerProtocol == "udp" {
						if err := conn.SetReadDeadline(time.Now().Add(s.timeout)); err != nil {
							fatalErrorChan <- err
						}

						// TODO: Don't use this packet which is DNS-specific, use the matching one from https://pojntfx.github.io/ports2packets/
						statusPacket, err := base64.StdEncoding.DecodeString("AAAQAAAAAAAAAAAA")
						if err != nil {
							fatalErrorChan <- err
						}

						if _, err := conn.Write(statusPacket); err != nil {
							fatalErrorChan <- err
						}

						// Count every response that is at least 1 byte long as a "open port"
						buffer := make([]byte, 1)
						if _, err := conn.Read(buffer); err != nil {
							s.scannedPortChan <- &ScannedPort{s.target, innerPort, innerProtocol, false}
						} else {
							s.scannedPortChan <- &ScannedPort{s.target, innerPort, innerProtocol, true}
						}
					} else {
						conn.Close()

						s.scannedPortChan <- &ScannedPort{s.target, innerPort, innerProtocol, true}
					}
				}

				wg.Done()
			}(port, protocol)
		}
	}

	go func() {
		wg.Wait()

		close(doneChan)
	}()

	select {
	case <-doneChan:
		s.scannedPortChan <- nil

		break
	case err := <-fatalErrorChan:
		s.scannedPortChan <- nil

		return err
	}

	return nil
}

func (s *PortScanner) Read() *ScannedPort {
	port := <-s.scannedPortChan

	return port
}
