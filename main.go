package main

import (
	"flag"
	"log"

	"github.com/pojntfx/liwasc/pkg/databases"
	"github.com/pojntfx/liwasc/pkg/servers"
	"github.com/pojntfx/liwasc/pkg/services"
	"github.com/pojntfx/liwasc/pkg/wakers"
	"golang.org/x/sync/semaphore"
)

func main() {
	// Parse flags
	deviceName := flag.String("deviceName", "eth0", "Network device name")
	mac2vendorDatabasePath := flag.String("mac2vendorDatabasePath", "/etc/liwasc/oui-database.sqlite", "Path to the mac2vendor database. Download from https://mac2vendor.com/articles/download")
	networkAndNodeScanDatabasePath := flag.String("networkAndNodeScanDatabasePath", "/var/liwasc/network_and_node_scan.sqlite", "Path to the persistence database for the network and node scan service.")
	nodeWakeDatabasePath := flag.String("nodeWakeDatabasePath", "/var/liwasc/node_wake.sqlite", "Path to the persistence database for the node wake service.")
	serviceNamesPortNumbersDatabasePath := flag.String("serviceNamesPortNumbersDatabasePath", "/etc/liwasc/service-names-port-numbers.csv", "Path to the CSV input file containing the registered services. Download from https://www.iana.org/assignments/service-names-port-numbers/service-names-port-numbers.xhtml")
	ports2PacketsDatabasePath := flag.String("ports2PacketsDatabasePath", "/etc/liwasc/ports2packets.csv", "Path to the ports2packets database. Download from https://github.com/pojntfx/ports2packets/releases")
	listenAddress := flag.String("listenAddress", "0.0.0.0:15123", "Listen address.")
	webSocketListenAddress := flag.String("webSocketListenAddress", "0.0.0.0:15124", "Listen address (for the WebSocket proxy).")
	maxConcurrentPortScans := flag.Uint("maxConcurrentPortScans", 1000, "Maximum concurrent port scans. Be sure to set this value to something lower than the systems ulimit or increase the latter.")
	periodicScanCronExpression := flag.String("periodicScanCronExpression", "@every 5m", "Cron expression for the periodic network scans & node scans. The default value will run a network & node scan every five minutes. See https://pkg.go.dev/github.com/robfig/cron for more information")
	periodicNetworkScanTimeout := flag.Int("periodicNetworkScanTimeout", 10000, "Time in milliseconds to wait for node discoveries in the periodic network scans.")
	periodicNodeScanTimeout := flag.Int("periodicNodeScanTimeout", 100, "Time in milliseconds to wait for a response per port in the periodic node scans.")

	flag.Parse()

	// Create instances
	mac2VendorDatabase := databases.NewMAC2VendorDatabase(*mac2vendorDatabasePath)
	networkAndNodeScanDatabase := databases.NewNetworkAndNodeScanDatabase(*networkAndNodeScanDatabasePath)
	nodeWakeDatabase := databases.NewNodeWakeDatabase(*nodeWakeDatabasePath)
	serviceNamesPortNumbersDatabase := databases.NewServiceNamesPortNumbersDatabase(*serviceNamesPortNumbersDatabasePath)
	ports2PacketsDatabase := databases.NewPorts2PacketDatabase(*ports2PacketsDatabasePath)
	networkAndNodeScanService := services.NewNetworkAndNodeScanService(
		*deviceName,
		mac2VendorDatabase,
		serviceNamesPortNumbersDatabase,
		ports2PacketsDatabase,
		networkAndNodeScanDatabase,
		semaphore.NewWeighted(int64(*maxConcurrentPortScans)),
		*periodicScanCronExpression,
		*periodicNetworkScanTimeout,
		*periodicNodeScanTimeout,
	)
	wakeOnLANWaker := wakers.NewWakeOnLANWaker(*deviceName)
	nodeWakeService := services.NewNodeWakeService(
		*deviceName,
		nodeWakeDatabase,
		func(macAddress string) (string, error) {
			node, err := networkAndNodeScanDatabase.GetNode(macAddress)
			if err != nil {
				return "", err
			}

			return node.IPAddress, nil
		},
		wakeOnLANWaker,
	)
	liwascServer := servers.NewLiwascServer(
		*listenAddress,
		*webSocketListenAddress,
		networkAndNodeScanService,
		nodeWakeService,
	)

	// Open instances
	if err := mac2VendorDatabase.Open(); err != nil {
		log.Fatal("could not open mac2VendorDatabase", err)
	}

	if err := serviceNamesPortNumbersDatabase.Open(); err != nil {
		log.Fatal("could not open serviceNamesPortNumbersDatabase", err)
	}

	if err := ports2PacketsDatabase.Open(); err != nil {
		log.Fatal("could not open ports2PacketsDatabase", err)
	}

	if err := networkAndNodeScanDatabase.Open(); err != nil {
		log.Fatal("could not open networkAndNodeScanDatabase", err)
	}

	if err := nodeWakeDatabase.Open(); err != nil {
		log.Fatal("could not open nodeWakeDatabase", err)
	}

	if err := wakeOnLANWaker.Open(); err != nil {
		log.Fatal("could not open wakeOnLANWaker", err)
	}

	go func() {
		if err := networkAndNodeScanService.Open(); err != nil {
			log.Fatal("could not open networkAndNodeScanService", err)
		}
	}()

	log.Printf("Listening on %v", *listenAddress)

	if err := liwascServer.ListenAndServe(); err != nil {
		log.Fatal("could not open liwasc server", err)
	}
}
