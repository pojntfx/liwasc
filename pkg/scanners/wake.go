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
	macAddress   string
	deviceName   string
	raddr        net.IP
	timeout      time.Duration
	statusChan   chan *ScannedNode
	getIPAddress func(string) (string, error)
}

func NewWakeScanner(macAddress string, deviceName string, timeout time.Duration, getIPAddress func(string) (string, error)) *WakeScanner {
	return &WakeScanner{
		macAddress,
		deviceName,
		nil,
		timeout,
		make(chan *ScannedNode),
		getIPAddress,
	}
}

func (w *WakeScanner) Open() error {
	ip, err := w.getIPAddress(w.macAddress)
	if err != nil {
		return err
	}

	raddr := net.ParseIP(ip)
	if raddr == nil {
		return errors.New("could not parse IP")
	}

	w.raddr = raddr

	arping.SetTimeout(w.timeout)

	return nil
}

func (w *WakeScanner) Transmit() error {
	macAddress, _, err := arping.PingOverIfaceByName(w.raddr, w.deviceName)
	if err != nil {
		if err == arping.ErrTimeout {
			w.statusChan <- &ScannedNode{macAddress.String(), false}
			w.statusChan <- nil

			return nil
		}

		w.statusChan <- nil

		return err
	}

	w.statusChan <- &ScannedNode{macAddress.String(), true}
	w.statusChan <- nil

	return nil
}

func (w *WakeScanner) Read() *ScannedNode {
	status := <-w.statusChan

	return status
}
