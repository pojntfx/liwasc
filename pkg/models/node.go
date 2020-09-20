package models

type Node struct {
	PoweredOn        bool
	MACAddress       string
	IPAddress        string
	Vendor           string
	ServicesAndPorts []string
	Registry         string
	Organization     string
	Address          string
	Visible          bool
}
