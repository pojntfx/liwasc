package databases

import (
	"context"
	"fmt"

	models "github.com/pojntfx/liwasc/pkg/sql/generated/node_and_port_scan"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

//go:generate sh -c "cd ../../ && sqlboiler sqlite3 -o pkg/sql/generated/node_and_port_scan -c pkg/sql/node_and_port_scan.toml"

type NodeAndPortScanDatabase struct {
	*SQLiteDatabase
}

func NewNodeAndPortScanDatabase(dbPath string) *NodeAndPortScanDatabase {
	return &NodeAndPortScanDatabase{&SQLiteDatabase{dbPath, nil}}
}

func (d *NodeAndPortScanDatabase) CreateNodeScan(nodeScan *models.NodeScan) error {
	return nodeScan.Insert(context.Background(), d.db, boil.Infer())
}

func (d *NodeAndPortScanDatabase) CreateNode(node *models.Node) error {
	return node.Insert(context.Background(), d.db, boil.Infer())
}

func (d *NodeAndPortScanDatabase) CreatePortScan(portScan *models.PortScan) error {
	return portScan.Insert(context.Background(), d.db, boil.Infer())
}

func (d *NodeAndPortScanDatabase) CreatePort(port *models.Port) error {
	return port.Insert(context.Background(), d.db, boil.Infer())
}

func (d *NodeAndPortScanDatabase) GetNodeScans() (models.NodeScanSlice, error) {
	return models.NodeScans(qm.OrderBy(models.NodeScanColumns.CreatedAt+" DESC")).All(context.Background(), d.db)
}

func (d *NodeAndPortScanDatabase) GetNodeScan(nodeScanID int64) (*models.NodeScan, error) {
	return models.FindNodeScan(context.Background(), d.db, nodeScanID)
}

func (d *NodeAndPortScanDatabase) GetNodes(nodeScanID int64) (models.NodeSlice, error) {
	return models.Nodes(models.NodeWhere.NodeScanID.EQ(nodeScanID), qm.OrderBy(models.NodeColumns.CreatedAt+" DESC")).All(context.Background(), d.db)
}

func (d *NodeAndPortScanDatabase) GetNodeByMACAddress(macAddress string) (*models.Node, error) {
	return models.Nodes(models.NodeWhere.MacAddress.EQ(macAddress)).One(context.Background(), d.db)
}

func (d *NodeAndPortScanDatabase) GetLookbackNodes() (models.NodeSlice, error) {
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

func (d *NodeAndPortScanDatabase) GetPortScans(nodeID int64) (models.PortScanSlice, error) {
	return models.PortScans(models.PortScanWhere.NodeID.EQ(nodeID), qm.OrderBy(models.PortScanColumns.CreatedAt+" DESC")).All(context.Background(), d.db)
}

func (d *NodeAndPortScanDatabase) GetPortScan(portScanID int64) (*models.PortScan, error) {
	return models.FindPortScan(context.Background(), d.db, portScanID)
}

func (d *NodeAndPortScanDatabase) GetLatestPortScanForNodeId(macAddress string) (*models.PortScan, error) {
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

func (d *NodeAndPortScanDatabase) GetPorts(portScanID int64) (models.PortSlice, error) {
	return models.Ports(models.PortWhere.PortScanID.EQ(portScanID), qm.OrderBy(models.PortColumns.CreatedAt+" DESC")).All(context.Background(), d.db)
}

func (d *NodeAndPortScanDatabase) UpdateNodeScan(nodeScan *models.NodeScan) error {
	_, err := nodeScan.Update(context.Background(), d.db, boil.Infer())

	return err
}

func (d *NodeAndPortScanDatabase) UpdatePortScan(portScan *models.PortScan) error {
	_, err := portScan.Update(context.Background(), d.db, boil.Infer())

	return err
}
