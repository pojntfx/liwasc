package networking

import "net"

type InterfaceInspector struct {
	device string
}

func NewInterfaceInspector(device string) *InterfaceInspector {
	return &InterfaceInspector{device}
}

func (i *InterfaceInspector) GetIPv4Subnets() ([]string, error) {
	iface, err := net.InterfaceByName(i.device)
	if err != nil {
		return nil, err
	}

	addrs, err := iface.Addrs()
	if err != nil {
		return nil, err
	}

	subnets := []string{}
	for _, addr := range addrs {
		ip, _, err := net.ParseCIDR(addr.String())
		if err != nil {
			return nil, err
		}

		if ip.To4() != nil {
			subnets = append(subnets, addr.String())
		}
	}

	return subnets, nil
}

func (i *InterfaceInspector) GetDevice() string {
	return i.device
}