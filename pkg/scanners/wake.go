package scanners

import (
	"errors"
	"net"
	"time"

	"github.com/j-keck/arping"
)

type ScannedNode struct {
	MacAddress string
	Awake      bool
}

type WakeScanner struct {
	ip         string
	deviceName string
	raddr      net.IP
	timeout    time.Duration
	statusChan chan *ScannedNode
}

func NewWakeScanner(ip string, deviceName string, timeout time.Duration) *WakeScanner {
	return &WakeScanner{
		ip,
		deviceName,
		nil,
		timeout,
		make(chan *ScannedNode),
	}
}

func (w *WakeScanner) Open() error {
	raddr := net.ParseIP(w.ip)
	if raddr == nil {
		return errors.New("could not parse IP")
	}

	w.raddr = raddr

	arping.SetTimeout(w.timeout)

	return nil
}

func (w *WakeScanner) Transmit() error {
	if _, _, err := arping.PingOverIfaceByName(w.raddr, w.deviceName); err != nil {
		if err == arping.ErrTimeout {
			w.statusChan <- &ScannedNode{w.ip, false}
			w.statusChan <- nil

			return nil
		}

		w.statusChan <- nil

		return err
	}

	w.statusChan <- &ScannedNode{w.ip, true}
	w.statusChan <- nil

	return nil
}

func (w *WakeScanner) Read() *ScannedNode {
	status := <-w.statusChan

	return status
}
