package persisters

import (
	"context"

	"github.com/gobuffalo/packr/v2"
	models "github.com/pojntfx/liwasc/pkg/db/sqlite/node_wake"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

//go:generate sqlboiler sqlite3 -o ../db/sqlite/node_wake -c ../../configs/sqlboiler/node_wake.toml

type NodeWakePersister struct {
	*SQLite
}

func NewNodeWakePersister(dbPath string) *NodeWakePersister {
	return &NodeWakePersister{
		&SQLite{
			DBPath: dbPath,
			Migrations: migrate.PackrMigrationSource{
				Box: packr.New("nodeWakeDatabaseMigrations", "../../db/sqlite/migrations/node_wake"),
			},
		},
	}
}

func (d *NodeWakePersister) CreateNodeWake(nodeWake *models.NodeWake) error {
	return nodeWake.Insert(context.Background(), d.db, boil.Infer())
}

func (d *NodeWakePersister) UpdateNodeWake(nodeWake *models.NodeWake) error {
	_, err := nodeWake.Update(context.Background(), d.db, boil.Infer())

	return err
}

func (d *NodeWakePersister) GetNodeWakes() (models.NodeWakeSlice, error) {
	return models.NodeWakes(qm.OrderBy(models.NodeWakeColumns.CreatedAt)).All(context.Background(), d.db)
}
