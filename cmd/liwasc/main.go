package main

import (
	"flag"
	"log"
	"math"
	"time"

	"github.com/pojntfx/liwasc/pkg/databases"
	"github.com/pojntfx/liwasc/pkg/scanners"
)

func main() {
	// Parse flags
	deviceName := flag.String("deviceName", "eth0", "Network device name")
	portScanningTimeout := flag.Int("portScanningTimeout", 15000, "Port scanning timeout (in milliseconds)")

	macDatabasePath := flag.String("macDatabasePath", "/etc/liwasc/oui-database.sqlite", "Path to the MAC database (mac2vendor flavour). Download from https://mac2vendor.com/articles/download")
	serviceNamesPortNumbersDatabasePath := flag.String("serviceNamesPortNumbersDatabasePath", "/etc/liwasc/service-names-port-numbers.csv", "Path to the CSV input file containing the registered services. Download from https://www.iana.org/assignments/service-names-port-numbers/service-names-port-numbers.xhtml")
	ports2PacketsDatabasePath := flag.String("ports2PacketsDatabasePath", "/etc/liwasc/ports2packets.csv", "Path to the ports2packets database. Download from https://github.com/pojntfx/ports2packets/releases")

	flag.Parse()

	// Create instances
	networkScanner := scanners.NewNetworkScanner(*deviceName)
	mac2VendorDatabase := databases.NewMAC2VendorDatabase(*macDatabasePath)
	serviceNamesPortNumbersDatabase := databases.NewServiceNamesPortNumbersDatabase(*serviceNamesPortNumbersDatabasePath)
	ports2PacketsDatabase := databases.NewPorts2PacketDatabase(*ports2PacketsDatabasePath)

	// Open instances
	err, subnets := networkScanner.Open()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Scanning nodes on subnets %v via %v\n", subnets, *deviceName)

	if err := mac2VendorDatabase.Open(); err != nil {
		log.Fatal(err)
	}

	if err := serviceNamesPortNumbersDatabase.Open(); err != nil {
		log.Fatal(err)
	}

	if err := ports2PacketsDatabase.Open(); err != nil {
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

		// future NODE_UPDATE message
		log.Println(node, "starting node scan")

		// Lookup vendor information for node
		go func() {
			vendor, err := mac2VendorDatabase.GetVendor(node.MACAddress.String())
			if err != nil {
				// future NODE_VENDOR_UPDATE message
				log.Println(node, "could not find vendor")

				return
			}

			// future NODE_VENDOR_UPDATE message
			log.Println(node, vendor)
		}()

		// Scan for open ports for node
		portScanner := scanners.NewPortScanner(node.IPAddress.String(), 0, math.MaxUint16, time.Millisecond*time.Duration(*portScanningTimeout), []string{"tcp", "udp"}, func(port int) ([]byte, error) {
			packet, err := ports2PacketsDatabase.GetPacket(port)

			if err != nil {
				return nil, err
			}

			return packet.Packet, err
		})

		// Dial and/or transmit packets ("start a scan")
		go func() {
			if err := portScanner.Transmit(); err != nil {
				log.Fatal(err)
			}
		}()

		go func() {
			for {
				port := portScanner.Read()

				if port == nil {
					// All ports have been scanned
					// future NODE_PORT_SCAN_DONE message
					log.Println(node, "port scan done")

					return
				}

				// Only send status info for open ports right now, but in the future closed port info should also be send to the frontend for status info ("how many ports have been scanned?")
				if port.Open {
					// future NODE_PORT_UPDATE message
					log.Println(node, port)
				}

				// Note above does not apply here, there is no point in transmitting this info if the ports are closed
				if port.Open {
					go func() {
						service, err := serviceNamesPortNumbersDatabase.GetService(port.Port, port.Protocol)
						if err != nil {
							// future NODE_PORT_SERVICE_INFO message
							log.Println(node, err)

							return
						}

						// future NODE_PORT_SERVICE_INFO message
						log.Println(port, service)
					}()
				}
			}
		}()
	}
}
