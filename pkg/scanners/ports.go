package scanners

import (
	"context"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
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
	lock            *semaphore.Weighted
	packetGetter    func(port int) ([]byte, error)
}

func NewPortScanner(target string, startPort int, endPort int, timeout time.Duration, protocols []string, lock *semaphore.Weighted, packetGetter func(port int) ([]byte, error)) *PortScanner {
	return &PortScanner{target, startPort, endPort, timeout, protocols, make(chan *ScannedPort), lock, packetGetter}
}

func (s *PortScanner) Transmit() error {
	doneChan := make(chan struct{})
	fatalErrorChan := make(chan error)

	var wg sync.WaitGroup
	for port := s.startPort; port <= s.endPort; port++ {
		for _, protocol := range s.protocols {
			if err := s.lock.Acquire(context.TODO(), 1); err != nil {
				return err
			}
			wg.Add(1)

			go func(innerPort int, innerProtocol string) {
				defer s.lock.Release(1)
				defer wg.Done()

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

						packet, err := s.packetGetter(innerPort)
						if err != nil {
							if strings.Contains(err.Error(), "could not find packet for port") {
								packet = []byte{} // Unknown packet for port, use empty []byte{}
							} else {
								fatalErrorChan <- err
							}
						}

						if _, err := conn.Write(packet); err != nil {
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
		close(s.scannedPortChan)

		break
	case err := <-fatalErrorChan:
		s.scannedPortChan <- nil
		close(s.scannedPortChan)

		return err
	}

	return nil
}

func (s *PortScanner) Read() *ScannedPort {
	port := <-s.scannedPortChan

	return port
}
