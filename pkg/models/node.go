package models

type Node struct {
	PoweredOn        bool
	MACAddress       string
	IPAddress        string
	Vendor           string
	ServicesAndPorts []string
}
