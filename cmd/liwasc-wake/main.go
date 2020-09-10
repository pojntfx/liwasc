package main

import (
	"flag"
	"github.com/pojntfx/liwasc/pkg/wakers"
	"log"
)

// This is temporary and will be integrated into `liwascd` and triggered via `liwascctl` once the server/client works

func main() {
	// Parse flags
	deviceName := flag.String("deviceName", "eth0", "Network device name")
	targetMacAddress := flag.String("targetMacAddress", "00:15:5d:89:01:02", "MAC address of the target device")

	flag.Parse()

	// Create instances
	wakeOnLANWaker := wakers.NewWakeOnLANWaker(*deviceName)

	// Open instances
	if err := wakeOnLANWaker.Open(); err != nil {
		log.Fatal(err)
	}

	// Send packet
	if err := wakeOnLANWaker.Write(*targetMacAddress); err != nil {
		log.Fatal(err)
	}

	log.Printf("Sent WOL magic packet to %v\n", *targetMacAddress)
}
