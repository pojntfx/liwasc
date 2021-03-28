package databases

import (
	"context"

	"github.com/gobuffalo/packr/v2"
	models "github.com/pojntfx/liwasc/pkg/databases/generated/node_wake"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

//go:generate sqlboiler sqlite3 -o generated/node_wake -c node_wake.toml

type NodeWakeDatabase struct {
	*SQLiteDatabase
}

func NewNodeWakeDatabase(dbPath string) *NodeWakeDatabase {
	return &NodeWakeDatabase{
		&SQLiteDatabase{
			DBPath: dbPath,
			Migrations: migrate.PackrMigrationSource{
				Box: packr.New("nodeWakeDatabaseMigrations", "../migrations/node_wake"),
			},
		},
	}
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
