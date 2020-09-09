package main

import (
	"flag"
	"log"

	"github.com/pojntfx/wascan/pkg/databases"
	"github.com/pojntfx/wascan/pkg/scanners"
)

func main() {
	// Parse flags
	deviceName := flag.String("deviceName", "eth0", "Network device name")
	macDatabasePath := flag.String("macDatabasePath", "/etc/wascan/oui-database.sqlite", "Path to the MAC database (mac2vendor flavour). Download from https://mac2vendor.com/articles/download")

	flag.Parse()

	// Create instances
	networkScanner := scanners.NewNetworkScanner(*deviceName)
	mac2VendorDatabase := databases.NewMAC2VendorDatabase(*macDatabasePath)

	// Open instances
	err, subnets := networkScanner.Open()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Scanning nodes on subnets %v via %v\n", subnets, *deviceName)

	if err := mac2VendorDatabase.Open(); err != nil {
		log.Fatal(err)
	}

	// Receive packets
	go func() {
		if err := networkScanner.Receive(); err != nil {
			log.Fatal(err)
		}
	}()

	// Transmit packets ("start a scan")
	go func() {
		if err := networkScanner.Transmit(); err != nil {
			log.Fatal(err)
		}
	}()

	// Connect instances
	for {
		node := networkScanner.Read()

		vendor, err := mac2VendorDatabase.GetVendor(node.MACAddress.String())
		if err != nil {
			log.Println(node, "could not find vendor")
		}

		log.Println(node, vendor)
	}
}
