package main

import (
	"flag"
	"log"

	"github.com/pojntfx/liwasc/pkg/databases"
	"github.com/pojntfx/liwasc/pkg/networking"
	"github.com/pojntfx/liwasc/pkg/servers"
	"github.com/pojntfx/liwasc/pkg/services"
	"github.com/pojntfx/liwasc/pkg/validators"
	"github.com/pojntfx/liwasc/pkg/wakers"
	"golang.org/x/sync/semaphore"
)

func main() {
	// Parse flags
	deviceName := flag.String("deviceName", "eth0", "Network device name")

	nodeAndPortScanDatabasePath := flag.String("nodeAndPortScanDatabasePath", "/var/lib/liwasc/node_and_port_scan.sqlite", "Path to the node and port scan database")
	nodeWakeDatabasePath := flag.String("nodeWakeDatabasePath", "/var/lib/liwasc/node_wake.sqlite", "Path to the node wake database")

	mac2vendorDatabasePath := flag.String("mac2vendorDatabasePath", "/etc/liwasc/oui-database.sqlite", "Path to the mac2vendor database")
	serviceNamesPortNumbersDatabasePath := flag.String("serviceNamesPortNumbersDatabasePath", "/etc/liwasc/service-names-port-numbers.csv", "Path to the CSV input file containing the registered services")
	ports2PacketsDatabasePath := flag.String("ports2PacketsDatabasePath", "/etc/liwasc/ports2packets.csv", "Path to the ports2packets database")

	mac2vendorDatabaseURL := flag.String("mac2vendorDatabaseURL", "https://mac2vendor.com/download/oui-database.sqlite", "URL to the mac2vendor database; will be downloaded on the first run if it doesn't exist")
	serviceNamesPortNumbersDatabaseURL := flag.String("serviceNamesPortNumbersDatabaseURL", "https://www.iana.org/assignments/service-names-port-numbers/service-names-port-numbers.csv", "URL to the CSV input file containing the registered services; will be downloaded on the first run if it doesn't exist")
	ports2PacketsDatabaseURL := flag.String("ports2PacketsDatabaseURL", "https://github.com/pojntfx/ports2packets/releases/download/weekly-csv/ports2packets.csv", "URL to the ports2packets database; will be downloaded on the first run if it doesn't exist")

	listenAddress := flag.String("listenAddress", "localhost:15123", "Listen address")
	webSocketListenAddress := flag.String("webSocketListenAddress", "localhost:15124", "Listen address (for the WebSocket proxy)")
	maxConcurrentPortScans := flag.Int64("maxConcurrentPortScans", 100, "Maximum concurrent port scans. Be sure to set this value to something lower than the systems ulimit or increase the latter")

	periodicScanCronExpression := flag.String("periodicScanCronExpression", "*/10 * * * *", "Cron expression for the periodic network scans & node scans. The default value will run a network & node scan every ten minutes. See https://pkg.go.dev/github.com/robfig/cron for more information")
	periodicNodeScanTimeout := flag.Int("periodicNodeScanTimeout", 500, "Time in milliseconds to wait for all nodes in a network to respond in the periodic node scans")
	periodicPortScanTimeout := flag.Int("periodicPortScanTimeout", 10, "Time in milliseconds to wait for a response per port in the periodic port scans")

	oidcIssuer := flag.String("oidcIssuer", "https://accounts.google.com", "OIDC issuer")
	oidcClientID := flag.String("oidcClientID", "myoidcclientid", "OIDC client ID")

	flag.Parse()

	// Create databases
	mac2VendorDatabase := databases.NewMAC2VendorDatabase(*mac2vendorDatabasePath, *mac2vendorDatabaseURL)
	serviceNamesPortNumbersDatabase := databases.NewServiceNamesPortNumbersDatabase(*serviceNamesPortNumbersDatabasePath, *serviceNamesPortNumbersDatabaseURL)
	ports2PacketsDatabase := databases.NewPorts2PacketDatabase(*ports2PacketsDatabasePath, *ports2PacketsDatabaseURL)
	nodeAndPortScanDatabase := databases.NewNodeAndPortScanDatabase(*nodeAndPortScanDatabasePath)
	nodeWakeDatabase := databases.NewNodeWakeDatabase(*nodeWakeDatabasePath)

	// Create generic utilities
	wakeOnLANWaker := wakers.NewWakeOnLANWaker(*deviceName)
	interfaceInspector := networking.NewInterfaceInspector(*deviceName)

	// Create auth utilities
	oidcValidator := validators.NewOIDCValidator(*oidcIssuer, *oidcClientID)
	contextValidator := validators.NewContextValidator(services.AUTHORIZATION_METADATA_KEY, oidcValidator)

	// Create services
	nodeAndPortScanService := services.NewNodeAndPortScanPortService(
		*deviceName,
		ports2PacketsDatabase,
		nodeAndPortScanDatabase,
		semaphore.NewWeighted(*maxConcurrentPortScans),
		*periodicScanCronExpression,
		*periodicNodeScanTimeout,
		*periodicPortScanTimeout,
		contextValidator,
	)
	metadataService := services.NewMetadataService(
		interfaceInspector,
		mac2VendorDatabase,
		serviceNamesPortNumbersDatabase,
		contextValidator,
	)
	nodeWakeService := services.NewNodeWakeService(
		*deviceName,
		wakeOnLANWaker,
		nodeWakeDatabase,
		func(macAddress string) (string, error) {
			node, err := nodeAndPortScanDatabase.GetNodeByMACAddress(macAddress)
			if err != nil {
				return "", err
			}

			return node.IPAddress, nil
		},
		contextValidator,
	)

	// Create server
	liwascServer := servers.NewLiwascServer(
		*listenAddress,
		*webSocketListenAddress,

		nodeAndPortScanService,
		metadataService,
		nodeWakeService,
	)

	// Open databases
	if err := mac2VendorDatabase.Open(); err != nil {
		log.Fatal("could not open mac2VendorDatabase", err)
	}
	if err := serviceNamesPortNumbersDatabase.Open(); err != nil {
		log.Fatal("could not open serviceNamesPortNumbersDatabase", err)
	}
	if err := ports2PacketsDatabase.Open(); err != nil {
		log.Fatal("could not open ports2PacketsDatabase", err)
	}
	if err := nodeAndPortScanDatabase.Open(); err != nil {
		log.Fatal("could not open networkAndNodeScanDatabase", err)
	}
	if err := nodeWakeDatabase.Open(); err != nil {
		log.Fatal("could not open nodeWakeDatabase", err)
	}

	// Open utilities
	if err := wakeOnLANWaker.Open(); err != nil {
		log.Fatal("could not open wakeOnLANWaker", err)
	}
	if err := oidcValidator.Open(); err != nil {
		log.Fatal("could not open oidcValidator", err)
	}

	// Open services
	if err := metadataService.Open(); err != nil {
		log.Fatal("could not open metadataService", err)
	}
	go func() {
		if err := nodeAndPortScanService.Open(); err != nil {
			log.Fatal("could not open nodeAndPortScanService", err)
		}
	}()

	// Start server
	log.Printf("liwasc backend listening on %v (gRPC) and %v (gRPC-Web)\n", *listenAddress, *webSocketListenAddress)

	if err := liwascServer.ListenAndServe(); err != nil {
		log.Fatalf("could not open liwasc backend: %v\n", err)
	}
}
