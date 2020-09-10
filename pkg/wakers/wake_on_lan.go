package wakers

import (
	"net"

	"github.com/mdlayher/wol"
)

type WakeOnLANWaker struct {
	deviceName string
	wolClient  *wol.RawClient
}

func NewWakeOnLANWaker(deviceName string) *WakeOnLANWaker {
	return &WakeOnLANWaker{deviceName, nil}
}

func (w *WakeOnLANWaker) Open() error {
	iface, err := net.InterfaceByName(w.deviceName)
	if err != nil {
		return err
	}

	wolClient, err := wol.NewRawClient(iface)
	if err != nil {
		return err
	}

	w.wolClient = wolClient

	return nil
}

func (w *WakeOnLANWaker) Write(targetMACAddress string) error {
	target, err := net.ParseMAC(targetMACAddress)
	if err != nil {
		return err
	}

	return w.wolClient.Wake(target)
}
