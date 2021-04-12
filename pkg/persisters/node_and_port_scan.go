package persisters

import (
	"context"
	"fmt"

	"github.com/pojntfx/liwasc/pkg/db/sqlite/migrations/node_and_port_scan"
	models "github.com/pojntfx/liwasc/pkg/db/sqlite/node_and_port_scan"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

//go:generate sqlboiler sqlite3 -o ../db/sqlite/node_and_port_scan -c ../../configs/sqlboiler/node_and_port_scan.toml
//go:generate go-bindata -pkg node_and_port_scan -o ../db/sqlite/migrations/node_and_port_scan/migrations.go ../../db/sqlite/migrations/node_and_port_scan

type NodeAndPortScanPersister struct {
	*SQLite
}

func NewNodeAndPortScanPersister(dbPath string) *NodeAndPortScanPersister {
	return &NodeAndPortScanPersister{
		&SQLite{
			DBPath: dbPath,
			Migrations: migrate.AssetMigrationSource{
				Asset:    node_and_port_scan.Asset,
				AssetDir: node_and_port_scan.AssetDir,
				Dir:      "../../db/sqlite/migrations/node_and_port_scan",
			},
		},
	}
}

func (d *NodeAndPortScanPersister) CreateNodeScan(nodeScan *models.NodeScan) error {
	return nodeScan.Insert(context.Background(), d.db, boil.Infer())
}

func (d *NodeAndPortScanPersister) CreateNode(node *models.Node) error {
	return node.Insert(context.Background(), d.db, boil.Infer())
}

func (d *NodeAndPortScanPersister) CreatePortScan(portScan *models.PortScan) error {
	return portScan.Insert(context.Background(), d.db, boil.Infer())
}

func (d *NodeAndPortScanPersister) CreatePort(port *models.Port) error {
	return port.Insert(context.Background(), d.db, boil.Infer())
}

func (d *NodeAndPortScanPersister) GetNodeScans() (models.NodeScanSlice, error) {
	return models.NodeScans(qm.OrderBy(models.NodeScanColumns.CreatedAt+" DESC")).All(context.Background(), d.db)
}

func (d *NodeAndPortScanPersister) GetNodeScan(nodeScanID int64) (*models.NodeScan, error) {
	return models.FindNodeScan(context.Background(), d.db, nodeScanID)
}

func (d *NodeAndPortScanPersister) GetNodes(nodeScanID int64) (models.NodeSlice, error) {
	return models.Nodes(models.NodeWhere.NodeScanID.EQ(nodeScanID), qm.OrderBy(models.NodeColumns.CreatedAt+" DESC")).All(context.Background(), d.db)
}

func (d *NodeAndPortScanPersister) GetNodeByMACAddress(macAddress string) (*models.Node, error) {
	return models.Nodes(models.NodeWhere.MacAddress.EQ(macAddress)).One(context.Background(), d.db)
}

func (d *NodeAndPortScanPersister) GetLookbackNodes() (models.NodeSlice, error) {
	var uniqueNodes models.NodeSlice
	if err := queries.Raw(
		fmt.Sprintf(
			`select *, max(%v) from %v group by %v`,
			models.NodeColumns.CreatedAt,
			models.TableNames.Nodes,
			models.NodeColumns.MacAddress,
		),
	).Bind(context.Background(), d.db, &uniqueNodes); err != nil {
		return nil, err
	}

	return uniqueNodes, nil
}

func (d *NodeAndPortScanPersister) GetPortScans(nodeID int64) (models.PortScanSlice, error) {
	return models.PortScans(models.PortScanWhere.NodeID.EQ(nodeID), qm.OrderBy(models.PortScanColumns.CreatedAt+" DESC")).All(context.Background(), d.db)
}

func (d *NodeAndPortScanPersister) GetPortScan(portScanID int64) (*models.PortScan, error) {
	return models.FindPortScan(context.Background(), d.db, portScanID)
}

func (d *NodeAndPortScanPersister) GetLatestPortScanForNodeId(macAddress string) (*models.PortScan, error) {
	var latestPortScan models.PortScan
	if err := queries.Raw(
		fmt.Sprintf(
			`select * from %v where %v = 1 and %v in (select %v from %v where %v=$1 order by %v desc) order by %v desc limit 1`,
			models.TableNames.PortScans,
			models.PortScanColumns.Done,
			models.PortScanColumns.NodeID,
			models.NodeColumns.ID,
			models.TableNames.Nodes,
			models.NodeColumns.MacAddress,
			models.NodeColumns.CreatedAt,
			models.PortScanColumns.CreatedAt,
		),
		macAddress,
	).Bind(context.Background(), d.db, &latestPortScan); err != nil {
		return nil, err
	}

	return &latestPortScan, nil
}

func (d *NodeAndPortScanPersister) GetPorts(portScanID int64) (models.PortSlice, error) {
	return models.Ports(models.PortWhere.PortScanID.EQ(portScanID), qm.OrderBy(models.PortColumns.CreatedAt+" DESC")).All(context.Background(), d.db)
}

func (d *NodeAndPortScanPersister) UpdateNodeScan(nodeScan *models.NodeScan) error {
	_, err := nodeScan.Update(context.Background(), d.db, boil.Infer())

	return err
}

func (d *NodeAndPortScanPersister) UpdatePortScan(portScan *models.PortScan) error {
	_, err := portScan.Update(context.Background(), d.db, boil.Infer())

	return err
}
