package main

import (
	"flag"
	"log"
	"time"

	"github.com/pojntfx/liwasc/pkg/scanners"
)

// This is temporary and will be integrated into `liwascd` and triggered via `liwascctl` once the server/client works

func main() {
	// Parse flags
	deviceName := flag.String("deviceName", "eth0", "Network device name")
	targetIPAddress := flag.String("targetIPAddress", "10.0.0.1", "IP address of the target device")
	timeout := flag.Int("timeout", 10000, "Timeout for ping in milliseconds")

	flag.Parse()

	// Create instances
	wakeScanner := scanners.NewWakeScanner(*targetIPAddress, *deviceName, time.Millisecond*time.Duration(*timeout))

	// Open instances
	if err := wakeScanner.Open(); err != nil {
		log.Fatal(err)
	}

	// Send packet
	go func() {
		if err := wakeScanner.Transmit(); err != nil {
			log.Fatal(err)
		}
	}()

	// Receive packet
	for {
		node := wakeScanner.Read()

		// Wake scan scan is done
		if node == nil {
			break
		}

		log.Printf("Host %v is %v\n", node.MacAddress, func() string {
			if node.Awake {
				return "awake"
			}

			return "not awake"
		}())
	}
}
