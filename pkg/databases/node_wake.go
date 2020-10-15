package databases

import (
	"context"

	models "github.com/pojntfx/liwasc/pkg/sql/generated/node_wake"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

//go:generate sh -c "cd ../../ && sqlboiler sqlite3 -o pkg/sql/generated/node_wake -c pkg/sql/node_wake.toml"

type NodeWakeDatabase struct {
	*SQLiteDatabase
}

func NewNodeWakeDatabase(dbPath string) *NodeWakeDatabase {
	return &NodeWakeDatabase{&SQLiteDatabase{dbPath, nil}}
}

func (d *NodeWakeDatabase) CreateNodeWake(nodeWake *models.NodeWake) error {
	return nodeWake.Insert(context.Background(), d.db, boil.Infer())
}

func (d *NodeWakeDatabase) UpdateNodeWake(nodeWake *models.NodeWake) error {
	_, err := nodeWake.Update(context.Background(), d.db, boil.Infer())

	return err
}

func (d *NodeWakeDatabase) GetNodeWakes() (models.NodeWakeSlice, error) {
	return models.NodeWakes(qm.OrderBy(models.NodeWakeColumns.CreatedAt)).All(context.Background(), d.db)
}
