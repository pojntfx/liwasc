package databases

//go:generate sh -c "cd ../../ && sqlboiler sqlite3 -o pkg/sql/generated/network_and_node_scan -c pkg/sql/network_and_node_scan.toml"

import (
	"context"
	"database/sql"

	networkAndNodeScanModels "github.com/pojntfx/liwasc/pkg/sql/generated/network_and_node_scan"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type NetworkAndNodeScanDatabase struct {
	dbPath string
	db     *sql.DB
}

func NewNetworkAndNodeScanDatabase(dbPath string) *NetworkAndNodeScanDatabase {
	return &NetworkAndNodeScanDatabase{dbPath, nil}
}

func (d *NetworkAndNodeScanDatabase) Open() error {
	db, err := sql.Open("sqlite3", d.dbPath)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(1) // Prevent "database locked" errors

	d.db = db

	return nil
}

func (d *NetworkAndNodeScanDatabase) CreateNetworkScan(scan *networkAndNodeScanModels.NetworkScan) (int64, error) {
	if err := scan.Insert(context.Background(), d.db, boil.Infer()); err != nil {
		return -1, err
	}

	return scan.ID, nil
}

func (d *NetworkAndNodeScanDatabase) UpdateNetworkScan(scan *networkAndNodeScanModels.NetworkScan) (int64, error) {
	if _, err := scan.Update(context.Background(), d.db, boil.Infer()); err != nil {
		return -1, err
	}

	return scan.ID, nil
}

func (d *NetworkAndNodeScanDatabase) GetNetworkScan(id int64) (*networkAndNodeScanModels.NetworkScan, error) {
	return networkAndNodeScanModels.FindNetworkScan(context.Background(), d.db, id)
}

// GetNewestNetworkScan returns the newest network scan
func (d *NetworkAndNodeScanDatabase) GetNewestNetworkScan() (*networkAndNodeScanModels.NetworkScan, error) {
	// Get the latest scan for the node
	scan, err := networkAndNodeScanModels.NetworkScans(
		qm.OrderBy(networkAndNodeScanModels.NetworkScansNodeColumns.CreatedAt+" desc"),
		qm.Limit(1),
	).One(context.Background(), d.db)
	if err != nil {
		return nil, err
	}

	return scan, nil
}

func (d *NetworkAndNodeScanDatabase) UpsertNode(node *networkAndNodeScanModels.Node, networkScanID int64) (string, error) {
	tx, err := d.db.BeginTx(context.Background(), nil)
	if err != nil {
		return "", err
	}

	// Insert node if it doesn't exist, otherwise update (the latter is required in case i.e. IP changes for MAC address)
	exists, err := networkAndNodeScanModels.NodeExists(context.Background(), tx, node.MacAddress)
	if err != nil {
		return "", err
	}

	if exists {
		if _, err := node.Update(context.Background(), tx, boil.Infer()); err != nil {
			return "", err
		}
	} else {
		if err := node.Insert(context.Background(), tx, boil.Infer()); err != nil {
			return "", err
		}
	}

	// Create the relationship between the scan and the node so that the active nodes of a scan can be fetched later
	networkScansNode := &networkAndNodeScanModels.NetworkScansNode{
		NodeID:        node.MacAddress,
		NetworkScanID: networkScanID,
	}

	if err := networkScansNode.Insert(context.Background(), tx, boil.Infer()); err != nil {
		return "", err
	}

	if err := tx.Commit(); err != nil {
		return "", err
	}

	return node.MacAddress, nil
}

func (d *NetworkAndNodeScanDatabase) GetAllNodes() ([]*networkAndNodeScanModels.Node, error) {
	allNodes, err := networkAndNodeScanModels.Nodes().All(context.Background(), d.db)
	if err != nil {
		return nil, err
	}

	return allNodes, nil
}

// GetNewestNetworkScansForNodes returns the newest scans in descending order
func (d *NetworkAndNodeScanDatabase) GetNewestNetworkScansForNodes(nodes []*networkAndNodeScanModels.Node) (map[string][]int64, error) {
	outMap := make(map[string][]int64)
	for _, node := range nodes {
		// Get the latest scans for the node
		scans, err := networkAndNodeScanModels.NetworkScansNodes(
			qm.Where(networkAndNodeScanModels.NetworkScansNodeColumns.NodeID+"= ?", node.MacAddress),
			qm.OrderBy(networkAndNodeScanModels.NetworkScansNodeColumns.CreatedAt+" desc"),
		).All(context.Background(), d.db)
		if err != nil {
			return nil, err
		}

		for _, scan := range scans {
			outMap[node.MacAddress] = append(outMap[node.MacAddress], scan.NetworkScanID)
		}

	}

	return outMap, nil
}

// UpsertService persists a service and connects it with it's node and network scan
// If the node scan was not triggered by a network scan,
func (d *NetworkAndNodeScanDatabase) UpsertService(service *networkAndNodeScanModels.Service, nodeID string, nodeScanID int64, networkScanID int64) (int64, error) {
	tx, err := d.db.BeginTx(context.Background(), nil)
	if err != nil {
		return -1, err
	}

	// Insert service if it doesn't exist, otherwise update
	// This way each service only needs to be saved once
	exists, err := networkAndNodeScanModels.ServiceExists(context.Background(), tx, service.PortNumber)
	if err != nil {
		return -1, err
	}

	if exists {
		if _, err := service.Update(context.Background(), tx, boil.Infer()); err != nil {
			return -1, err
		}
	} else {
		if err := service.Insert(context.Background(), tx, boil.Infer()); err != nil {
			return -1, err
		}
	}

	// Create a relationship between the service and the node for later fetching
	networkScansNode := &networkAndNodeScanModels.NodeScansServicesNode{
		ServiceID:  service.PortNumber,
		NodeID:     nodeID,
		NodeScanID: nodeScanID,
	}

	if err := networkScansNode.Insert(context.Background(), tx, boil.Infer()); err != nil {
		return -1, err
	}

	if err := tx.Commit(); err != nil {
		return -1, err
	}

	return service.PortNumber, nil
}

// CreateNodeScan creates a node scan. Set nodeID to an empty string and networkScanID to -1 if there is not network scan for this node scan
func (d *NetworkAndNodeScanDatabase) CreateNodeScan(scan *networkAndNodeScanModels.NodeScan, nodeID string, networkScanID int64) (int64, error) {
	tx, err := d.db.BeginTx(context.Background(), nil)
	if err != nil {
		return -1, err
	}

	if err := scan.Insert(context.Background(), tx, boil.Infer()); err != nil {
		return -1, err
	}

	// Create a relationship between the node scan and the network scan (if there is a network scan)
	nodeNodeScanNetworkScan := &networkAndNodeScanModels.NodeNodeScansNetworkScan{
		NodeID:        nodeID,
		NetworkScanID: networkScanID,
		NodeScanID:    scan.ID,
	}

	if err := nodeNodeScanNetworkScan.Insert(context.Background(), tx, boil.Infer()); err != nil {
		return -1, err
	}

	if err := tx.Commit(); err != nil {
		return -1, err
	}

	return scan.ID, nil
}

func (d *NetworkAndNodeScanDatabase) UpdateNodeScan(scan *networkAndNodeScanModels.NodeScan) (int64, error) {
	if _, err := scan.Update(context.Background(), d.db, boil.Infer()); err != nil {
		return -1, err
	}

	return scan.ID, nil
}

// GetNodeScanIDByNetworkScanIDAndNodeID returns the node scan ID for a network scan ID and a node ID
func (d *NetworkAndNodeScanDatabase) GetNodeScanIDByNetworkScanIDAndNodeID(nodeID string, networkScanID int64) (int64, error) {
	nodeScan, err := networkAndNodeScanModels.NodeNodeScansNetworkScans(
		qm.Where(networkAndNodeScanModels.NodeNodeScansNetworkScanColumns.NodeID+"= ?", nodeID),
		qm.And(networkAndNodeScanModels.NodeNodeScansNetworkScanColumns.NetworkScanID+"= ?", networkScanID),
		qm.Limit(1),
	).One(context.Background(), d.db)
	if err != nil {
		return -1, err
	}

	return nodeScan.NodeScanID, err
}

func (d *NetworkAndNodeScanDatabase) GetNewestNodeScanIDForNodeID(nodeID string) (int64, error) {
	nodeScan, err := networkAndNodeScanModels.NodeNodeScansNetworkScans(
		qm.Where(networkAndNodeScanModels.NodeNodeScansNetworkScanColumns.NodeID+"= ?", nodeID),
		qm.OrderBy(networkAndNodeScanModels.NodeNodeScansNetworkScanColumns.CreatedAt+" desc"),
		qm.Limit(1),
	).One(context.Background(), d.db)
	if err != nil {
		return -1, err
	}

	return nodeScan.NodeScanID, err
}

func (d *NetworkAndNodeScanDatabase) GetServicesForNodeScanID(nodeScanID int64) ([]*networkAndNodeScanModels.Service, error) {
	var services []*networkAndNodeScanModels.Service
	err := networkAndNodeScanModels.NewQuery(
		qm.Select("*"),
		qm.From(networkAndNodeScanModels.TableNames.NodeScansServicesNodes),
		qm.InnerJoin(
			networkAndNodeScanModels.TableNames.Services+" on "+networkAndNodeScanModels.TableNames.NodeScansServicesNodes+"."+networkAndNodeScanModels.NodeScansServicesNodeColumns.ServiceID+" = "+networkAndNodeScanModels.TableNames.Services+"."+networkAndNodeScanModels.ServiceColumns.PortNumber,
		),
		qm.Where(networkAndNodeScanModels.TableNames.NodeScansServicesNodes+"."+networkAndNodeScanModels.NodeScansServicesNodeColumns.NodeScanID+"= ?", nodeScanID),
	).Bind(context.Background(), d.db, &services)

	return services, err
}

func (d *NetworkAndNodeScanDatabase) GetNodeScan(id int64) (*networkAndNodeScanModels.NodeScan, error) {
	return networkAndNodeScanModels.FindNodeScan(context.Background(), d.db, id)
}

func (d *NetworkAndNodeScanDatabase) GetNode(id string) (*networkAndNodeScanModels.Node, error) {
	return networkAndNodeScanModels.FindNode(context.Background(), d.db, id)
}

func (d *NetworkAndNodeScanDatabase) DeleteNode(id string) (*networkAndNodeScanModels.Node, error) {
	node, err := d.GetNode(id)
	if err != nil {
		return nil, err
	}

	tx, err := d.db.BeginTx(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	if _, err = networkAndNodeScanModels.Nodes(
		qm.Where(networkAndNodeScanModels.NodeColumns.MacAddress+"= ?", id),
	).DeleteAll(context.Background(), tx); err != nil {
		return nil, err
	}

	if _, err := networkAndNodeScanModels.NodeScansServicesNodes(
		qm.Where(networkAndNodeScanModels.NodeScansServicesNodeColumns.NodeID+"= ?", id),
	).DeleteAll(context.Background(), tx); err != nil {
		return nil, err
	}

	if _, err := networkAndNodeScanModels.NodeNodeScansNetworkScans(
		qm.Where(networkAndNodeScanModels.NodeNodeScansNetworkScanColumns.NodeID+"= ?", id),
	).DeleteAll(context.Background(), tx); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return node, nil
}

func (d *NetworkAndNodeScanDatabase) CreatePeriodicNetworkScan(scan *networkAndNodeScanModels.PeriodicNetworkScansNetworkScan) (int64, error) {
	if err := scan.Insert(context.Background(), d.db, boil.Infer()); err != nil {
		return -1, err
	}

	return scan.ID, nil
}

func (d *NetworkAndNodeScanDatabase) GetNewestPeriodicNetworkScan() (*networkAndNodeScanModels.PeriodicNetworkScansNetworkScan, error) {
	scan, err := networkAndNodeScanModels.PeriodicNetworkScansNetworkScans(
		qm.OrderBy(networkAndNodeScanModels.PeriodicNetworkScansNetworkScanColumns.CreatedAt+" desc"),
		qm.Limit(1),
	).One(context.Background(), d.db)
	if err != nil {
		return nil, err
	}

	return scan, nil
}
