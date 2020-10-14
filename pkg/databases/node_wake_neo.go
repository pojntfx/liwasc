package databases

//go:generate sh -c "cd ../../ && sqlboiler sqlite3 -o pkg/sql/generated/node_wake_neo -c pkg/sql/node_wake_neo.toml"

type NodeWakeNeoDatabase struct {
	*SQLiteDatabase
}

func NewNodeWakeNeoDatabase(dbPath string) *NodeWakeNeoDatabase {
	return &NodeWakeNeoDatabase{&SQLiteDatabase{dbPath, nil}}
}
