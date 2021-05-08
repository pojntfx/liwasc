package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/pojntfx/liwasc/pkg/networking"
	"github.com/pojntfx/liwasc/pkg/persisters"
	"github.com/pojntfx/liwasc/pkg/servers"
	"github.com/pojntfx/liwasc/pkg/services"
	"github.com/pojntfx/liwasc/pkg/validators"
	"github.com/pojntfx/liwasc/pkg/wakers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/sync/semaphore"
)

const (
	configFileKey                          = "configFile"
	deviceNameKey                          = "deviceName"
	nodeAndPortScanDatabasePathKey         = "nodeAndPortScanDatabasePath"
	nodeWakeDatabasePathKey                = "nodeWakeDatabasePath"
	mac2vendorDatabasePathKey              = "mac2vendorDatabasePath"
	serviceNamesPortNumbersDatabasePathKey = "serviceNamesPortNumbersDatabasePath"
	ports2PacketsDatabasePathKey           = "ports2PacketsDatabasePath"
	mac2vendorDatabaseURLKey               = "mac2vendorDatabaseURL"
	serviceNamesPortNumbersDatabaseURLKey  = "serviceNamesPortNumbersDatabaseURL"
	ports2PacketsDatabaseURLKey            = "ports2PacketsDatabaseURL"
	listenAddressKey                       = "listenAddress"
	webSocketListenAddressKey              = "webSocketListenAddress"
	maxConcurrentPortScansKey              = "maxConcurrentPortScans"
	periodicScanCronExpressionKey          = "periodicScanCronExpression"
	periodicNodeScanTimeoutKey             = "periodicNodeScanTimeout"
	periodicPortScanTimeoutKey             = "periodicPortScanTimeout"
	oidcIssuerKey                          = "oidcIssuer"
	oidcClientIDKey                        = "oidcClientID"
	prepareOnlyKey                         = "prepareOnly"
)

func main() {
	// Create command
	cmd := &cobra.Command{
		Use:   "liwasc-backend",
		Short: "List, wake and scan nodes in a network.",
		Long: `liwasc is a high-performance network and port scanner. It can quickly give you a overview of the nodes in your network, the services that run on them and manage their power status.

For more information, please visit https://github.com/pojntfx/liwasc.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Bind config file
			if !(viper.GetString(configFileKey) == "") {
				viper.SetConfigFile(viper.GetString(configFileKey))

				if err := viper.ReadInConfig(); err != nil {
					return err
				}
			}

			// Create persisters
			mac2VendorPersister := persisters.NewMAC2VendorPersister(viper.GetString(mac2vendorDatabasePathKey), viper.GetString(mac2vendorDatabaseURLKey))
			serviceNamesPortNumbersPersister := persisters.NewServiceNamesPortNumbersPersister(viper.GetString(serviceNamesPortNumbersDatabasePathKey), viper.GetString(serviceNamesPortNumbersDatabaseURLKey))
			ports2PacketsPersister := persisters.NewPorts2PacketPersister(viper.GetString(ports2PacketsDatabasePathKey), viper.GetString(ports2PacketsDatabaseURLKey))
			nodeAndPortScanPersister := persisters.NewNodeAndPortScanPersister(viper.GetString(nodeAndPortScanDatabasePathKey))
			nodeWakePersister := persisters.NewNodeWakePersister(viper.GetString(nodeWakeDatabasePathKey))

			// Create generic utilities
			wakeOnLANWaker := wakers.NewWakeOnLANWaker(viper.GetString(deviceNameKey))
			interfaceInspector := networking.NewInterfaceInspector(viper.GetString(deviceNameKey))

			// Create auth utilities
			oidcValidator := validators.NewOIDCValidator(viper.GetString(oidcIssuerKey), viper.GetString(oidcClientIDKey))
			contextValidator := validators.NewContextValidator(services.AUTHORIZATION_METADATA_KEY, oidcValidator)

			// Create services
			nodeAndPortScanService := services.NewNodeAndPortScanPortService(
				viper.GetString(deviceNameKey),
				ports2PacketsPersister,
				nodeAndPortScanPersister,
				semaphore.NewWeighted(viper.GetInt64(maxConcurrentPortScansKey)),
				viper.GetString(periodicScanCronExpressionKey),
				viper.GetInt(periodicNodeScanTimeoutKey),
				viper.GetInt(periodicPortScanTimeoutKey),
				contextValidator,
			)
			metadataService := services.NewMetadataService(
				interfaceInspector,
				mac2VendorPersister,
				serviceNamesPortNumbersPersister,
				contextValidator,
			)
			nodeWakeService := services.NewNodeWakeService(
				viper.GetString(deviceNameKey),
				wakeOnLANWaker,
				nodeWakePersister,
				func(macAddress string) (string, error) {
					node, err := nodeAndPortScanPersister.GetNodeByMACAddress(macAddress)
					if err != nil {
						return "", err
					}

					return node.IPAddress, nil
				},
				contextValidator,
			)

			// Create server
			liwascServer := servers.NewLiwascServer(
				viper.GetString(listenAddressKey),
				viper.GetString(webSocketListenAddressKey),

				nodeAndPortScanService,
				metadataService,
				nodeWakeService,
			)

			// Open persisters
			if err := mac2VendorPersister.Open(); err != nil {
				log.Fatal("could not open mac2VendorDatabase", err)
			}
			if err := serviceNamesPortNumbersPersister.Open(); err != nil {
				log.Fatal("could not open serviceNamesPortNumbersDatabase", err)
			}
			if err := ports2PacketsPersister.Open(); err != nil {
				log.Fatal("could not open ports2PacketsDatabase", err)
			}
			if err := nodeAndPortScanPersister.Open(); err != nil {
				log.Fatal("could not open networkAndNodeScanDatabase", err)
			}
			if err := nodeWakePersister.Open(); err != nil {
				log.Fatal("could not open nodeWakeDatabase", err)
			}

			// Init is done, exit
			if viper.GetBool(prepareOnlyKey) {
				os.Exit(0)
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
			log.Printf("liwasc backend listening on %v (gRPC) and %v (gRPC-Web)\n", viper.GetString(listenAddressKey), viper.GetString(webSocketListenAddressKey))

			return liwascServer.ListenAndServe()
		},
	}

	// Get prefix
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("could not get home directory", err)
	}
	prefix := filepath.Join(home, ".local", "share", "liwasc")

	// Bind flags
	cmd.PersistentFlags().StringP(configFileKey, "c", "", "Config file to use")
	cmd.PersistentFlags().StringP(deviceNameKey, "d", "eth0", "Network device name")

	cmd.PersistentFlags().String(nodeAndPortScanDatabasePathKey, filepath.Join(prefix, "var", "lib", "liwasc", "node_and_port_scan.sqlite"), "Path to the node and port scan database")
	cmd.PersistentFlags().String(nodeWakeDatabasePathKey, filepath.Join(prefix, "var", "lib", "liwasc", "node_wake.sqlite"), "Path to the node wake database")

	cmd.PersistentFlags().String(mac2vendorDatabasePathKey, filepath.Join(prefix, "etc", "liwasc", "oui-database.sqlite"), "Path to the mac2vendor database")
	cmd.PersistentFlags().String(serviceNamesPortNumbersDatabasePathKey, filepath.Join(prefix, "etc", "liwasc", "service-names-port-numbers.csv"), "Path to the CSV input file containing the registered services")
	cmd.PersistentFlags().String(ports2PacketsDatabasePathKey, filepath.Join(prefix, "etc", "liwasc", "ports2packets.csv"), "Path to the ports2packets database")

	cmd.PersistentFlags().String(mac2vendorDatabaseURLKey, "https://mac2vendor.com/download/oui-database.sqlite", "URL to the mac2vendor database; will be downloaded on the first run if it doesn't exist")
	cmd.PersistentFlags().String(serviceNamesPortNumbersDatabaseURLKey, "https://www.iana.org/assignments/service-names-port-numbers/service-names-port-numbers.csv", "URL to the CSV input file containing the registered services; will be downloaded on the first run if it doesn't exist")
	cmd.PersistentFlags().String(ports2PacketsDatabaseURLKey, "https://github.com/pojntfx/ports2packets/releases/download/weekly-csv/ports2packets.csv", "URL to the ports2packets database; will be downloaded on the first run if it doesn't exist")

	cmd.PersistentFlags().StringP(listenAddressKey, "l", "localhost:15123", "Listen address")
	cmd.PersistentFlags().StringP(webSocketListenAddressKey, "w", "localhost:15124", "Listen address (for the WebSocket proxy)")
	cmd.PersistentFlags().Int64P(maxConcurrentPortScansKey, "u", 100, "Maximum concurrent port scans. Be sure to set this value to something lower than the systems ulimit or increase the latter")

	cmd.PersistentFlags().StringP(periodicScanCronExpressionKey, "e", "*/60 * * * *", "Cron expression for the periodic network scans & node scans. The default value will run a network & node scan every 60 minutes. See https://pkg.go.dev/github.com/robfig/cron for more information")
	cmd.PersistentFlags().IntP(periodicNodeScanTimeoutKey, "n", 500, "Time in milliseconds to wait for all nodes in a network to respond in the periodic node scans")
	cmd.PersistentFlags().IntP(periodicPortScanTimeoutKey, "p", 10, "Time in milliseconds to wait for a response per port in the periodic port scans")

	cmd.PersistentFlags().StringP(oidcIssuerKey, "i", "https://pojntfx.eu.auth0.com/", "OIDC issuer")
	cmd.PersistentFlags().StringP(oidcClientIDKey, "t", "myoidcclientid", "OIDC client ID")

	cmd.PersistentFlags().BoolP(prepareOnlyKey, "o", false, "Only download external databases & prepare them, then exit")

	// Bind env variables
	if err := viper.BindPFlags(cmd.PersistentFlags()); err != nil {
		log.Fatal(err)
	}
	viper.SetEnvPrefix("liwasc_backend")
	viper.AutomaticEnv()

	// Run command
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
