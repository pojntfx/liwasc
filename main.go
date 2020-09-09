package main

import (
	"flag"
	"log"

	"github.com/pojntfx/wascan/pkg/scanners"
)

func main() {
	// Parse flags
	deviceName := flag.String("deviceName", "eth0", "Network device name")

	flag.Parse()

	// Create instances
	networkScanner := scanners.NewNetworkScanner(*deviceName)

	// Open instances
	err, subnets := networkScanner.Open()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Scanning nodes on subnets %v via %v\n", subnets, *deviceName)

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

		log.Println(node)
	}
}
