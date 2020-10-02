package main

import (
	"flag"
	"log"
	"time"

	arp "github.com/ItsJimi/go-arp"
	"github.com/pojntfx/liwasc/pkg/scanners"
)

// This is temporary and will be integrated into `liwascd` and triggered via `liwascctl` once the server/client works

func main() {
	// Parse flags
	deviceName := flag.String("deviceName", "eth0", "Network device name")
	targetMacAddress := flag.String("targetMacAddress", "00:15:5d:89:01:02", "MAC address of the target device")
	timeout := flag.Int("timeout", 10000, "Timeout for ping in milliseconds")

	flag.Parse()

	// Create instances
	wakeScanner := scanners.NewWakeScanner(*targetMacAddress, *deviceName, time.Millisecond*time.Duration(*timeout), func(macAddress string) (string, error) {
		ip, err := arp.GetEntryFromMAC(macAddress)

		return ip.IPAddress, err
	})

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
