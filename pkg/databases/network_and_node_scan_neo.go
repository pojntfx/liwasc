package databases

import (
	"context"
	"database/sql"
	"fmt"

	models "github.com/pojntfx/liwasc/pkg/sql/generated/network_and_node_scan_neo"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
)

//go:generate sh -c "cd ../../ && sqlboiler sqlite3 -o pkg/sql/generated/network_and_node_scan_neo -c pkg/sql/network_and_node_scan_neo.toml"

type NetworkAndNodeScanNeoDatabase struct {
	dbPath string
	db     *sql.DB
}

func NewNetworkAndNodeScanNeoDatabase(dbPath string) *NetworkAndNodeScanNeoDatabase {
	return &NetworkAndNodeScanNeoDatabase{dbPath, nil}
}

func (d *NetworkAndNodeScanNeoDatabase) Open() error {
	db, err := sql.Open("sqlite3", d.dbPath)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(1) // Prevent "database locked" errors

	d.db = db

	return nil
}

func (d *NetworkAndNodeScanNeoDatabase) CreateNetworkScan(networkScan *models.NetworkScan) error {
	return networkScan.Insert(context.Background(), d.db, boil.Infer())
}

func (d *NetworkAndNodeScanNeoDatabase) CreateNode(node *models.Node) error {
	return node.Insert(context.Background(), d.db, boil.Infer())
}

func (d *NetworkAndNodeScanNeoDatabase) CreateNodeScan(nodeScan *models.NodeScan) error {
	return nodeScan.Insert(context.Background(), d.db, boil.Infer())
}

func (d *NetworkAndNodeScanNeoDatabase) CreateService(service *models.Service) error {
	return service.Insert(context.Background(), d.db, boil.Infer())
}

func (d *NetworkAndNodeScanNeoDatabase) GetNetworkScan(id int64) (*models.NetworkScan, error) {
	return models.FindNetworkScan(context.Background(), d.db, id)
}

func (d *NetworkAndNodeScanNeoDatabase) LookbackForNodes() (models.NodeSlice, error) {
	var uniqueNodes models.NodeSlice
	if err := queries.Raw(
		fmt.Sprintf(
			`select * from %v where %v in (select max(%v) from %v group by %v)`,
			models.TableNames.Nodes,
			models.NodeColumns.ID,
			models.NodeColumns.CreatedAt,
			models.TableNames.Nodes,
			models.NodeColumns.MacAddress,
		),
	).Bind(context.Background(), d.db, &uniqueNodes); err != nil {
		return nil, err
	}

	return uniqueNodes, nil
}
