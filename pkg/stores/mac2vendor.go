package stores

//go:generate sqlboiler sqlite3 -o ../db/mac2vendor -c ../../configs/mac2vendor.toml

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	mac2vendorModels "github.com/pojntfx/liwasc/pkg/db/mac2vendor"
)

type MAC2VendorDatabase struct {
	*SQLiteDatabase
	*ExternalSource
}

func NewMAC2VendorDatabase(dbPath string, sourceURL string) *MAC2VendorDatabase {
	return &MAC2VendorDatabase{
		&SQLiteDatabase{
			DBPath: dbPath,
		},
		&ExternalSource{
			SourceURL:       sourceURL,
			DestinationPath: dbPath,
		},
	}
}

func (d *MAC2VendorDatabase) Open() error {
	// If database file does not exist, download & create it
	if err := d.ExternalSource.PullIfNotExists(); err != nil {
		return err
	}

	return d.SQLiteDatabase.Open()
}

func (d *MAC2VendorDatabase) GetVendor(mac string) (*mac2vendorModels.Vendordb, error) {
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
