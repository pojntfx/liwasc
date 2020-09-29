package databases

//go:generate sh -c "cd ../../ && sqlboiler sqlite3 -o pkg/sql/generated/liwasc -c pkg/sql/liwasc.toml"

import (
	"context"
	"database/sql"
	"fmt"

	liwascModels "github.com/pojntfx/liwasc/pkg/sql/generated/liwasc"
	models "github.com/pojntfx/liwasc/pkg/sql/generated/liwasc"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type LiwascDatabase struct {
	dbPath string
	db     *sql.DB
}

func NewLiwascDatabase(dbPath string) *LiwascDatabase {
	return &LiwascDatabase{dbPath, nil}
}

func (d *LiwascDatabase) Open() error {
	db, err := sql.Open("sqlite3", fmt.Sprintf("%v?cache=shared", d.dbPath)) // Prevent "database locked" errors
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(1) // Prevent "database locked" errors

	d.db = db

	return nil
}

func (d *LiwascDatabase) CreateNetworkScan(scan *liwascModels.NetworkScan) (int64, error) {
	if err := scan.Insert(context.Background(), d.db, boil.Infer()); err != nil {
		return -1, err
	}

	return scan.ID, nil
}

func (d *LiwascDatabase) UpdateNetworkScan(scan *liwascModels.NetworkScan) (int64, error) {
	if _, err := scan.Update(context.Background(), d.db, boil.Infer()); err != nil {
		return -1, err
	}

	return scan.ID, nil
}

func (d *LiwascDatabase) GetNetworkScan(id int64) (*liwascModels.NetworkScan, error) {
	return liwascModels.FindNetworkScan(context.Background(), d.db, id)
}

// GetNewestNetworkScan returns the newest network scan
func (d *LiwascDatabase) GetNewestNetworkScan() (*liwascModels.NetworkScan, error) {
	// Get the latest scan for the node
	scan, err := models.NetworkScans(
		qm.OrderBy(liwascModels.NetworkScansNodeColumns.CreatedAt+" desc"),
		qm.Limit(1),
	).One(context.Background(), d.db)
	if err != nil {
		return nil, err
	}

	return scan, nil
}

func (d *LiwascDatabase) UpsertNode(node *liwascModels.Node, networkScanID int64) (string, error) {
	// Insert node if it doesn't exist, otherwise update (the latter is required in case i.e. IP changes for MAC address)
	exists, err := liwascModels.NodeExists(context.Background(), d.db, node.MacAddress)
	if err != nil {
		return "", err
	}

	if exists {
		if _, err := node.Update(context.Background(), d.db, boil.Infer()); err != nil {
			return "", err
		}
	} else {
		if err := node.Insert(context.Background(), d.db, boil.Infer()); err != nil {
			return "", err
		}
	}

	// Create the relationship between the scan and the node so that the active nodes of a scan can be fetched later
	networkScansNode := &liwascModels.NetworkScansNode{
		NodeID:        node.MacAddress,
		NetworkScanID: networkScanID,
	}

	if err := networkScansNode.Insert(context.Background(), d.db, boil.Infer()); err != nil {
		return "", err
	}

	return node.MacAddress, nil
}

func (d *LiwascDatabase) GetAllNodes() ([]*liwascModels.Node, error) {
	allNodes, err := liwascModels.Nodes().All(context.Background(), d.db)
	if err != nil {
		return nil, err
	}

	return allNodes, nil
}

// GetNewestNetworkScansForNodes returns the newest scans in descending order
func (d *LiwascDatabase) GetNewestNetworkScansForNodes(nodes []*liwascModels.Node) (map[string][]int64, error) {
	outMap := make(map[string][]int64)
	for _, node := range nodes {
		// Get the latest scans for the node
		scans, err := models.NetworkScansNodes(
			qm.Where(liwascModels.NetworkScansNodeColumns.NodeID+"= ?", node.MacAddress),
			qm.OrderBy(liwascModels.NetworkScansNodeColumns.CreatedAt+" desc"),
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
func (d *LiwascDatabase) UpsertService(service *liwascModels.Service, nodeID string, nodeScanID int64, networkScanID int64) (int64, error) {
	// Insert service if it doesn't exist, otherwise update
	// This way each service only needs to be saved once
	exists, err := liwascModels.ServiceExists(context.Background(), d.db, service.PortNumber)
	if err != nil {
		return -1, err
	}

	if exists {
		if _, err := service.Update(context.Background(), d.db, boil.Infer()); err != nil {
			return -1, err
		}
	} else {
		if err := service.Insert(context.Background(), d.db, boil.Infer()); err != nil {
			return -1, err
		}
	}

	// Create a relationship between the service and the node for later fetching
	networkScansNode := &liwascModels.NodeScansServicesNode{
		ServiceID: service.PortNumber,
		NodeID:    nodeID,
	}

	if err := networkScansNode.Insert(context.Background(), d.db, boil.Infer()); err != nil {
		return -1, err
	}

	return service.PortNumber, nil
}

// CreateNodeScan creates a node scan. Set nodeID to an empty string and networkScanID to -1 if there is not network scan for this node scan
func (d *LiwascDatabase) CreateNodeScan(scan *liwascModels.NodeScan, nodeID string, networkScanID int64) (int64, error) {
	if err := scan.Insert(context.Background(), d.db, boil.Infer()); err != nil {
		return -1, err
	}

	// Create a relationship between the node scan and the network scan (if there is a network scan)
	if networkScanID != -1 && nodeID != "" {
		nodeNodeScanNetworkScan := &liwascModels.NodeNodeScansNetworkScan{
			NodeID:        nodeID,
			NetworkScanID: networkScanID,
			NodeScanID:    scan.ID,
		}

		if err := nodeNodeScanNetworkScan.Insert(context.Background(), d.db, boil.Infer()); err != nil {
			return -1, err
		}
	}

	return scan.ID, nil
}

func (d *LiwascDatabase) UpdateNodeScan(scan *liwascModels.NodeScan) (int64, error) {
	if _, err := scan.Update(context.Background(), d.db, boil.Infer()); err != nil {
		return -1, err
	}

	return scan.ID, nil
}

// GetNodeScanIDByNetworkScanIDAndNodeID returns the node scan ID for a network scan ID and a node ID
func (d *LiwascDatabase) GetNodeScanIDByNetworkScanIDAndNodeID(nodeID string, networkScanID int64) (int64, error) {
	nodeScan, err := models.NodeNodeScansNetworkScans(
		qm.Where(liwascModels.NodeNodeScansNetworkScanColumns.NodeID+"= ?", nodeID),
		qm.And(liwascModels.NodeNodeScansNetworkScanColumns.NetworkScanID+"= ?", networkScanID),
		qm.Limit(1),
	).One(context.Background(), d.db)
	if err != nil {
		return -1, err
	}

	return nodeScan.NodeScanID, err
}
