package databases

//go:generate sh -c "cd ../../ && sqlboiler sqlite3 -o pkg/sql/generated/liwasc -c pkg/sql/liwasc.toml"

import (
	"context"
	"database/sql"

	liwascModels "github.com/pojntfx/liwasc/pkg/sql/generated/liwasc"
	"github.com/volatiletech/sqlboiler/boil"
)

type LiwascDatabase struct {
	dbPath string
	db     *sql.DB
}

func NewLiwascDatabase(dbPath string) *LiwascDatabase {
	return &LiwascDatabase{dbPath, nil}
}

func (d *LiwascDatabase) Open() error {
	db, err := sql.Open("sqlite3", d.dbPath)
	if err != nil {
		return err
	}

	d.db = db

	return nil
}

func (d *LiwascDatabase) CreateScan(scan *liwascModels.Scan) (int64, error) {
	if err := scan.Insert(context.Background(), d.db, boil.Infer()); err != nil {
		return -1, err
	}

	return scan.ID, nil
}

func (d *LiwascDatabase) UpsertNode(node *liwascModels.Node, scanID int64) (string, error) {
	exists, err := liwascModels.NodeExists(context.Background(), d.db, node.MacAddress)
	if err != nil {
		return "", err
	}
	if exists {
		if _, err := node.Update(context.Background(), d.db, boil.Infer()); err != nil {
			return "", err
		}

		return node.MacAddress, nil
	}

	if err := node.Insert(context.Background(), d.db, boil.Infer()); err != nil {
		return "", err
	}

	return node.MacAddress, nil
}
