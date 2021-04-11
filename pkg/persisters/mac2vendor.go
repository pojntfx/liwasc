package persisters

//go:generate sqlboiler sqlite3 -o ../db/sqlite/mac2vendor -c ../../configs/mac2vendor.toml

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	mac2vendorModels "github.com/pojntfx/liwasc/pkg/db/sqlite/mac2vendor"
)

type MAC2VendorPersister struct {
	*SQLite
	*ExternalSource
}

func NewMAC2VendorPersister(dbPath string, sourceURL string) *MAC2VendorPersister {
	return &MAC2VendorPersister{
		&SQLite{
			DBPath: dbPath,
		},
		&ExternalSource{
			SourceURL:       sourceURL,
			DestinationPath: dbPath,
		},
	}
}

func (d *MAC2VendorPersister) Open() error {
	// If database file does not exist, download & create it
	if err := d.ExternalSource.PullIfNotExists(); err != nil {
		return err
	}

	return d.SQLite.Open()
}

func (d *MAC2VendorPersister) GetVendor(mac string) (*mac2vendorModels.Vendordb, error) {
	oui, err := GetOUI(mac)
	if err != nil {
		return nil, err
	}

	vendor, err := mac2vendorModels.FindVendordb(context.Background(), d.db, int64(oui))
	if err != nil {
		return nil, err
	}

	return vendor, nil
}

func GetOUI(mac string) (uint64, error) {
	parsedMAC := strings.Split(mac, ":")
	if len(parsedMAC) < 4 {
		return 0, fmt.Errorf("invalid MAC Address: %v", mac)
	}

	res, err := strconv.ParseUint(strings.Join(parsedMAC[0:3], ""), 16, 64)
	if err != nil {
		return 0, err
	}

	return uint64(res), err
}
