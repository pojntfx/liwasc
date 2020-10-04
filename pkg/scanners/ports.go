package scanners

import (
	"context"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	tcpShaker "github.com/tevino/tcp-shaker"
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
	nonFatalErrorChan := make(chan error)
	fatalErrorChan := make(chan error)

	var wg sync.WaitGroup
	for port := s.startPort; port <= s.endPort; port++ {
		for _, protocol := range s.protocols {
			wg.Add(1)

			if err := s.lock.Acquire(context.TODO(), 1); err != nil {
				return err
			}

			go func(innerPort int, innerProtocol string) {
				defer func() {
					s.lock.Release(1)
					wg.Done()
				}()

				raddr := net.JoinHostPort(s.target, fmt.Sprintf("%v", innerPort))

				if innerProtocol == "udp" {
					// Do a UDP scan using the packets from ports2packets
					conn, err := net.DialTimeout(innerProtocol, raddr, s.timeout)
					if err != nil {
						s.scannedPortChan <- &ScannedPort{s.target, innerPort, innerProtocol, false}
					}

					if conn != nil {
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
							// UDP port is closed
							s.scannedPortChan <- &ScannedPort{s.target, innerPort, innerProtocol, false}
						} else {
							// UDP port is open
							s.scannedPortChan <- &ScannedPort{s.target, innerPort, innerProtocol, true}
						}
					}
				} else {
					// Do a stealth (SYN) TCP scan
					// See https://github.com/tevino/tcp-shaker
					c := tcpShaker.NewChecker()

					ctx, stopChecker := context.WithCancel(context.Background())
					defer stopChecker()

					go func() {
						if err := c.CheckingLoop(ctx); err != nil {
							nonFatalErrorChan <- err
						}
					}()

					<-c.WaitReady()

					if err := c.CheckAddr(raddr, s.timeout); err != nil {
						// TCP port is closed
						s.scannedPortChan <- &ScannedPort{s.target, innerPort, innerProtocol, false}
					} else {
						// TCP port is open
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
	case <-nonFatalErrorChan:
		s.scannedPortChan <- nil

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
