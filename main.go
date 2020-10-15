package main

import (
	"flag"
	"log"

	"github.com/pojntfx/liwasc/pkg/concurrency"
	"github.com/pojntfx/liwasc/pkg/databases"
	"github.com/pojntfx/liwasc/pkg/networking"
	"github.com/pojntfx/liwasc/pkg/servers"
	"github.com/pojntfx/liwasc/pkg/services"
	"github.com/pojntfx/liwasc/pkg/validators"
	"github.com/pojntfx/liwasc/pkg/wakers"
)

func main() {
	// Parse flags
	deviceName := flag.String("deviceName", "eth0", "Network device name")
	mac2vendorDatabasePath := flag.String("mac2vendorDatabasePath", "/etc/liwasc/oui-database.sqlite", "Path to the mac2vendor database. Download from https://mac2vendor.com/articles/download")
	networkAndNodeScanDatabasePath := flag.String("networkAndNodeScanDatabasePath", "/var/liwasc/network_and_node_scan.sqlite", "Path to the persistence database for the network and node scan service.")
	nodeAndPortScanDatabasePath := flag.String("nodeAndPortScanDatabasePath", "/var/liwasc/node_and_port_scan.sqlite", "Path to the persistence database for the node and port scan service.")
	nodeWakeDatabasePath := flag.String("nodeWakeDatabasePath", "/var/liwasc/node_wake.sqlite", "Path to the persistence database for the node wake service.")
	nodeWakeNeoDatabasePath := flag.String("nodeWakeNeoDatabasePath", "/var/liwasc/node_wake_neo.sqlite", "Path to the persistence database for the neo node wake service.")
	serviceNamesPortNumbersDatabasePath := flag.String("serviceNamesPortNumbersDatabasePath", "/etc/liwasc/service-names-port-numbers.csv", "Path to the CSV input file containing the registered services. Download from https://www.iana.org/assignments/service-names-port-numbers/service-names-port-numbers.xhtml")
	ports2PacketsDatabasePath := flag.String("ports2PacketsDatabasePath", "/etc/liwasc/ports2packets.csv", "Path to the ports2packets database. Download from https://github.com/pojntfx/ports2packets/releases")
	listenAddress := flag.String("listenAddress", "0.0.0.0:15123", "Listen address.")
	webSocketListenAddress := flag.String("webSocketListenAddress", "0.0.0.0:15124", "Listen address (for the WebSocket proxy).")
	maxConcurrentPortScans := flag.Int("maxConcurrentPortScans", 100, "Maximum concurrent port scans. Be sure to set this value to something lower than the systems ulimit or increase the latter.")
	periodicScanCronExpression := flag.String("periodicScanCronExpression", "*/5 * * * *", "Cron expression for the periodic network scans & node scans. The default value will run a network & node scan every five minutes. See https://pkg.go.dev/github.com/robfig/cron for more information")
	periodicNodeScanTimeout := flag.Int("periodicNodeScanTimeout", 500, "Time in milliseconds to wait for all nodes in a network to respond in the periodic node scans.")
	periodicPortScanTimeout := flag.Int("periodicPortScanTimeout", 50, "Time in milliseconds to wait for a response per port in the periodic port scans.")
	oidcIssuer := flag.String("oidcIssuer", "https://accounts.google.com", "OIDC issuer")
	oidcClientID := flag.String("oidcClientID", "myoidcclientid", "OIDC client ID")

	flag.Parse()

	// Create instances
	mac2VendorDatabase := databases.NewMAC2VendorDatabase(*mac2vendorDatabasePath)
	networkAndNodeScanDatabase := databases.NewNetworkAndNodeScanDatabase(*networkAndNodeScanDatabasePath)
	nodeAndPortScanDatabase := databases.NewNodeAndPortScanDatabase(*nodeAndPortScanDatabasePath)
	nodeWakeDatabase := databases.NewNodeWakeDatabase(*nodeWakeDatabasePath)
	serviceNamesPortNumbersDatabase := databases.NewServiceNamesPortNumbersDatabase(*serviceNamesPortNumbersDatabasePath)
	ports2PacketsDatabase := databases.NewPorts2PacketDatabase(*ports2PacketsDatabasePath)
	oidcValidator := validators.NewOIDCValidator(*oidcIssuer, *oidcClientID)
	contextValidator := validators.NewContextValidator(services.AUTHORIZATION_METADATA_KEY, oidcValidator)
	// networkAndNodeScanService := services.NewNetworkAndNodeScanService(
	// 	*deviceName,
	// 	mac2VendorDatabase,
	// 	serviceNamesPortNumbersDatabase,
	// 	ports2PacketsDatabase,
	// 	networkAndNodeScanDatabase,
	// 	concurrency.NewGoRoutineLimiter(int32(*maxConcurrentPortScans)),
	// 	*periodicScanCronExpression,
	// 	*periodicNetworkScanTimeout,
	// 	*periodicNodeScanTimeout,
	// 	contextValidator,
	// )
	nodeAndPortScanService := services.NewNodeAndPortScanPortService(
		*deviceName,
		ports2PacketsDatabase,
		nodeAndPortScanDatabase,
		concurrency.NewGoRoutineLimiter(int32(*maxConcurrentPortScans)),
		*periodicScanCronExpression,
		*periodicNodeScanTimeout,
		*periodicPortScanTimeout,
		contextValidator,
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
		contextValidator,
	)
	interfaceInspector := networking.NewInterfaceInspector(*deviceName)
	metadataService := services.NewMetadataService(interfaceInspector, contextValidator)
	metadataNeoService := services.NewMetadataNeoService(
		interfaceInspector,
		mac2VendorDatabase,
		serviceNamesPortNumbersDatabase,
		contextValidator,
	)
	nodeWakeNeoDatabase := databases.NewNodeWakeNeoDatabase(*nodeWakeNeoDatabasePath)
	nodeWakeNeoService := services.NewNodeWakeNeoService(
		*deviceName,
		wakeOnLANWaker,
		nodeWakeNeoDatabase,
		func(macAddress string) (string, error) {
			node, err := nodeAndPortScanDatabase.GetNodeByMACAddress(macAddress)
			if err != nil {
				return "", err
			}

			return node.IPAddress, nil
		},
		contextValidator,
	)
	liwascServer := servers.NewLiwascServer(
		*listenAddress,
		*webSocketListenAddress,
		// networkAndNodeScanService,
		nodeWakeService,
		metadataService,
		nodeAndPortScanService,
		metadataNeoService,
		nodeWakeNeoService,
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

	if err := nodeAndPortScanDatabase.Open(); err != nil {
		log.Fatal("could not open networkAndNodeScanNeoDatabase", err)
	}

	if err := oidcValidator.Open(); err != nil {
		log.Fatal("could not open oidcValidator", err)
	}

	if err := nodeWakeDatabase.Open(); err != nil {
		log.Fatal("could not open nodeWakeDatabase", err)
	}

	if err := nodeWakeNeoDatabase.Open(); err != nil {
		log.Fatal("could not open nodeWakeNeoDatabase", err)
	}

	if err := wakeOnLANWaker.Open(); err != nil {
		log.Fatal("could not open wakeOnLANWaker", err)
	}

	// go func() {
	// 	if err := networkAndNodeScanService.Open(); err != nil {
	// 		log.Fatal("could not open networkAndNodeScanService", err)
	// 	}
	// }()

	go func() {
		if err := nodeAndPortScanService.Open(); err != nil {
			log.Fatal("could not open nodeAndPortScanService", err)
		}
	}()

	if err := metadataService.Open(); err != nil {
		log.Fatal("could not open metadataService", err)
	}

	if err := metadataNeoService.Open(); err != nil {
		log.Fatal("could not open metadataNeoService", err)
	}

	log.Printf("Listening on %v", *listenAddress)

	if err := liwascServer.ListenAndServe(); err != nil {
		log.Fatal("could not open liwasc server", err)
	}
}
