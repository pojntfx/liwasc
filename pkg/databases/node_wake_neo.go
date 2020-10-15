package databases

import (
	"context"

	models "github.com/pojntfx/liwasc/pkg/sql/generated/node_wake_neo"
	"github.com/volatiletech/sqlboiler/boil"
)

//go:generate sh -c "cd ../../ && sqlboiler sqlite3 -o pkg/sql/generated/node_wake_neo -c pkg/sql/node_wake_neo.toml"

type NodeWakeNeoDatabase struct {
	*SQLiteDatabase
}

func NewNodeWakeNeoDatabase(dbPath string) *NodeWakeNeoDatabase {
	return &NodeWakeNeoDatabase{&SQLiteDatabase{dbPath, nil}}
}

func (d *NodeWakeNeoDatabase) CreateNodeWake(nodeWake *models.NodeWakesNeo) error {
	return nodeWake.Insert(context.Background(), d.db, boil.Infer())
}

func (d *NodeWakeNeoDatabase) UpdateNodeWake(nodeWake *models.NodeWakesNeo) error {
	_, err := nodeWake.Update(context.Background(), d.db, boil.Infer())

	return err
}
