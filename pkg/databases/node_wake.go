package databases

import (
	"context"
	"database/sql"

	nodeWakeModels "github.com/pojntfx/liwasc/pkg/sql/generated/node_wake"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

//go:generate sh -c "cd ../../ && sqlboiler sqlite3 -o pkg/sql/generated/node_wake -c pkg/sql/node_wake.toml"

type NodeWakeDatabase struct {
	dbPath string
	db     *sql.DB
}

func NewNodeWakeDatabase(dbPath string) *NodeWakeDatabase {
	return &NodeWakeDatabase{dbPath, nil}
}

func (d *NodeWakeDatabase) Open() error {
	db, err := sql.Open("sqlite3", d.dbPath)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(1) // Prevent "database locked" errors

	d.db = db

	return nil
}

func (d *NodeWakeDatabase) CreateNodeWake(nodeWake *nodeWakeModels.NodeWake) (int64, error) {
	if err := nodeWake.Insert(context.Background(), d.db, boil.Infer()); err != nil {
		return -1, err
	}

	return nodeWake.ID, nil
}

func (d *NodeWakeDatabase) UpsertNodeAndRelations(node *nodeWakeModels.Node, nodeWakeID int64) (string, error) {
	tx, err := d.db.BeginTx(context.Background(), nil)
	if err != nil {
		return "", err
	}

	// Insert node if it doesn't exist, otherwise update
	exists, err := nodeWakeModels.NodeExists(context.Background(), tx, node.MacAddress)
	if err != nil {
		return "", err
	}

	if !exists {
		if err := node.Insert(context.Background(), tx, boil.Infer()); err != nil {
			return "", err
		}
	}

	// Create the relationship between the wake and the node so that the nodes of a wake can be fetched later
	networkScansNode := &nodeWakeModels.NodeWakesNode{
		NodeID:      node.MacAddress,
		NodeWakesID: nodeWakeID,
	}

	if err := networkScansNode.Insert(context.Background(), tx, boil.Infer()); err != nil {
		return "", err
	}

	if err := tx.Commit(); err != nil {
		return "", err
	}

	return node.MacAddress, nil
}

func (d *NodeWakeDatabase) UpdateNodeWakeScan(scan *nodeWakeModels.NodeWake) (int64, error) {
	if _, err := scan.Update(context.Background(), d.db, boil.Infer()); err != nil {
		return -1, err
	}

	return scan.ID, nil
}

func (d *NodeWakeDatabase) GetNodeWake(id int64) (*nodeWakeModels.NodeWake, error) {
	return nodeWakeModels.FindNodeWake(context.Background(), d.db, id)
}

func (d *NodeWakeDatabase) GetNewestNodeWakeIDForNodeID(nodeID string) (int64, error) {
	nodeWake, err := nodeWakeModels.NodeWakesNodes(
		qm.Where(nodeWakeModels.NodeWakesNodeColumns.NodeID+"= ?", nodeID),
		qm.OrderBy(nodeWakeModels.NodeWakesNodeColumns.CreatedAt+" desc"),
		qm.Limit(1),
	).One(context.Background(), d.db)
	if err != nil {
		return -1, err
	}

	return nodeWake.NodeWakesID, err
}

func (d *NodeWakeDatabase) GetNodeForNodeWakeID(nodeWakeID int64) (*nodeWakeModels.Node, error) {
	var nodes []*nodeWakeModels.Node
	err := nodeWakeModels.NewQuery(
		qm.Select("*"),
		qm.From(nodeWakeModels.TableNames.NodeWakesNodes),
		qm.InnerJoin(
			nodeWakeModels.TableNames.Nodes+" on "+
				nodeWakeModels.TableNames.NodeWakesNodes+"."+nodeWakeModels.NodeWakesNodeColumns.NodeID+" = "+
				nodeWakeModels.TableNames.Nodes+"."+nodeWakeModels.NodeColumns.MacAddress,
		),
		qm.Where(nodeWakeModels.TableNames.NodeWakesNodes+"."+nodeWakeModels.NodeWakesNodeColumns.NodeWakesID+"= ?", nodeWakeID),
	).Bind(context.Background(), d.db, &nodes)

	return nodes[0], err
}
