package main

import (
	"flag"
	"log"

	"github.com/pojntfx/liwasc/pkg/databases"
	"github.com/pojntfx/liwasc/pkg/servers"
	"github.com/pojntfx/liwasc/pkg/services"
)

func main() {
	// Parse flags
	deviceName := flag.String("deviceName", "eth0", "Network device name")
	mac2vendorDatabasePath := flag.String("mac2vendorDatabasePath", "/etc/liwasc/oui-database.sqlite", "Path to the mac2vendor database. Download from https://mac2vendor.com/articles/download")
	liwascDatabasePath := flag.String("liwascDatabasePath", "/var/liwasc/liwasc.sqlite", "Path to the persistence database.")
	listenAddress := flag.String("listenAddress", "0.0.0.0:15123", "Listen address.")

	flag.Parse()

	// Create instances
	mac2VendorDatabase := databases.NewMAC2VendorDatabase(*mac2vendorDatabasePath)
	liwascDatabase := databases.NewLiwascDatabase(*liwascDatabasePath)
	networkAndNodeScanService := services.NewNetworkAndNodeScanService(*deviceName, mac2VendorDatabase, liwascDatabase)
	liwascServer := servers.NewLiwascServer(*listenAddress, networkAndNodeScanService)

	// Open instances
	if err := mac2VendorDatabase.Open(); err != nil {
		log.Fatal("could not open mac2VendorDatabase", err)
	}

	if err := liwascDatabase.Open(); err != nil {
		log.Fatal("could not open liwasc", err)
	}

	log.Printf("Listening on %v", *listenAddress)

	if err := liwascServer.Open(); err != nil {
		log.Fatal("could not open liwasc server", err)
	}
}
