package databases

//go:generate sh -c "cd ../../ && sqlboiler sqlite3 -o pkg/sql/generated/liwasc -c pkg/sql/liwasc.toml"

import (
	"context"
	"database/sql"

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
	db, err := sql.Open("sqlite3", d.dbPath)
	if err != nil {
		return err
	}

	d.db = db

	return nil
}

func (d *LiwascDatabase) CreateNodeScan(scan *liwascModels.NodeScan) (int64, error) {
	if err := scan.Insert(context.Background(), d.db, boil.Infer()); err != nil {
		return -1, err
	}

	return scan.ID, nil
}

func (d *LiwascDatabase) UpdateNodeScan(scan *liwascModels.NodeScan) (int64, error) {
	if _, err := scan.Update(context.Background(), d.db, boil.Infer()); err != nil {
		return -1, err
	}

	return scan.ID, nil
}

func (d *LiwascDatabase) UpsertNode(node *liwascModels.Node, scanID int64) (string, error) {
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
	scansNode := &liwascModels.NodeScansNode{
		NodeID:     node.MacAddress,
		NodeScanID: scanID,
	}

	if err := scansNode.Insert(context.Background(), d.db, boil.Infer()); err != nil {
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

// GetNewestNodeScansForNodes returns the newest scans in descending order
func (d *LiwascDatabase) GetNewestNodeScansForNodes(nodes []*liwascModels.Node) (map[string][]int64, error) {
	outMap := make(map[string][]int64)
	for _, node := range nodes {
		// Get the latest scans for the node
		scans, err := models.NodeScansNodes(
			qm.Where(liwascModels.NodeScansNodeColumns.NodeID+"= ?", node.MacAddress),
			qm.OrderBy(liwascModels.NodeScansNodeColumns.CreatedAt+" desc"),
		).All(context.Background(), d.db)
		if err != nil {
			return nil, err
		}

		for _, scan := range scans {
			outMap[node.MacAddress] = append(outMap[node.MacAddress], scan.NodeScanID)
		}

	}

	return outMap, nil
}

func (d *LiwascDatabase) GetNodeScan(id int64) (*liwascModels.NodeScan, error) {
	return liwascModels.FindNodeScan(context.Background(), d.db, id)
}

// GetNewestScan returns the newest scan
func (d *LiwascDatabase) GetNewestScan() (*liwascModels.NodeScan, error) {
	// Get the latest scan for the node
	scan, err := models.NodeScans(
		qm.OrderBy(liwascModels.NodeScansNodeColumns.CreatedAt+" desc"),
		qm.Limit(1),
	).One(context.Background(), d.db)
	if err != nil {
		return nil, err
	}

	return scan, nil
}

func (d *LiwascDatabase) UpsertService(service *liwascModels.Service, nodeID string) (int64, error) {
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
	nodesService := &liwascModels.NodesService{
		NodeID:    nodeID,
		ServiceID: service.PortNumber,
	}

	if err := nodesService.Insert(context.Background(), d.db, boil.Infer()); err != nil {
		return -1, err
	}

	return service.PortNumber, nil
}
