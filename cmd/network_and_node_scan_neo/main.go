package main

import (
	"flag"
	"log"

	"github.com/pojntfx/liwasc/pkg/databases"
)

func main() {
	dbPath := flag.String("dbPath", "/var/liwasc/network_and_node_scan_neo.sqlite", "Database path")

	flag.Parse()

	db := databases.NewNetworkAndNodeScanNeoDatabase(*dbPath)

	if err := db.Open(); err != nil {
		log.Fatal(err)
	}

	nodes, err := db.LookbackForNodes()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(nodes)
}
