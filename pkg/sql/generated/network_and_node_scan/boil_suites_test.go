// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("NetworkScans", testNetworkScans)
	t.Run("NetworkScansNodes", testNetworkScansNodes)
	t.Run("NodeNodeScansNetworkScans", testNodeNodeScansNetworkScans)
	t.Run("NodeScans", testNodeScans)
	t.Run("NodeScansServicesNodes", testNodeScansServicesNodes)
	t.Run("Nodes", testNodes)
	t.Run("PeriodicNetworkScansNetworkScans", testPeriodicNetworkScansNetworkScans)
	t.Run("Services", testServices)
}

func TestDelete(t *testing.T) {
	t.Run("NetworkScans", testNetworkScansDelete)
	t.Run("NetworkScansNodes", testNetworkScansNodesDelete)
	t.Run("NodeNodeScansNetworkScans", testNodeNodeScansNetworkScansDelete)
	t.Run("NodeScans", testNodeScansDelete)
	t.Run("NodeScansServicesNodes", testNodeScansServicesNodesDelete)
	t.Run("Nodes", testNodesDelete)
	t.Run("PeriodicNetworkScansNetworkScans", testPeriodicNetworkScansNetworkScansDelete)
	t.Run("Services", testServicesDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("NetworkScans", testNetworkScansQueryDeleteAll)
	t.Run("NetworkScansNodes", testNetworkScansNodesQueryDeleteAll)
	t.Run("NodeNodeScansNetworkScans", testNodeNodeScansNetworkScansQueryDeleteAll)
	t.Run("NodeScans", testNodeScansQueryDeleteAll)
	t.Run("NodeScansServicesNodes", testNodeScansServicesNodesQueryDeleteAll)
	t.Run("Nodes", testNodesQueryDeleteAll)
	t.Run("PeriodicNetworkScansNetworkScans", testPeriodicNetworkScansNetworkScansQueryDeleteAll)
	t.Run("Services", testServicesQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("NetworkScans", testNetworkScansSliceDeleteAll)
	t.Run("NetworkScansNodes", testNetworkScansNodesSliceDeleteAll)
	t.Run("NodeNodeScansNetworkScans", testNodeNodeScansNetworkScansSliceDeleteAll)
	t.Run("NodeScans", testNodeScansSliceDeleteAll)
	t.Run("NodeScansServicesNodes", testNodeScansServicesNodesSliceDeleteAll)
	t.Run("Nodes", testNodesSliceDeleteAll)
	t.Run("PeriodicNetworkScansNetworkScans", testPeriodicNetworkScansNetworkScansSliceDeleteAll)
	t.Run("Services", testServicesSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("NetworkScans", testNetworkScansExists)
	t.Run("NetworkScansNodes", testNetworkScansNodesExists)
	t.Run("NodeNodeScansNetworkScans", testNodeNodeScansNetworkScansExists)
	t.Run("NodeScans", testNodeScansExists)
	t.Run("NodeScansServicesNodes", testNodeScansServicesNodesExists)
	t.Run("Nodes", testNodesExists)
	t.Run("PeriodicNetworkScansNetworkScans", testPeriodicNetworkScansNetworkScansExists)
	t.Run("Services", testServicesExists)
}

func TestFind(t *testing.T) {
	t.Run("NetworkScans", testNetworkScansFind)
	t.Run("NetworkScansNodes", testNetworkScansNodesFind)
	t.Run("NodeNodeScansNetworkScans", testNodeNodeScansNetworkScansFind)
	t.Run("NodeScans", testNodeScansFind)
	t.Run("NodeScansServicesNodes", testNodeScansServicesNodesFind)
	t.Run("Nodes", testNodesFind)
	t.Run("PeriodicNetworkScansNetworkScans", testPeriodicNetworkScansNetworkScansFind)
	t.Run("Services", testServicesFind)
}

func TestBind(t *testing.T) {
	t.Run("NetworkScans", testNetworkScansBind)
	t.Run("NetworkScansNodes", testNetworkScansNodesBind)
	t.Run("NodeNodeScansNetworkScans", testNodeNodeScansNetworkScansBind)
	t.Run("NodeScans", testNodeScansBind)
	t.Run("NodeScansServicesNodes", testNodeScansServicesNodesBind)
	t.Run("Nodes", testNodesBind)
	t.Run("PeriodicNetworkScansNetworkScans", testPeriodicNetworkScansNetworkScansBind)
	t.Run("Services", testServicesBind)
}

func TestOne(t *testing.T) {
	t.Run("NetworkScans", testNetworkScansOne)
	t.Run("NetworkScansNodes", testNetworkScansNodesOne)
	t.Run("NodeNodeScansNetworkScans", testNodeNodeScansNetworkScansOne)
	t.Run("NodeScans", testNodeScansOne)
	t.Run("NodeScansServicesNodes", testNodeScansServicesNodesOne)
	t.Run("Nodes", testNodesOne)
	t.Run("PeriodicNetworkScansNetworkScans", testPeriodicNetworkScansNetworkScansOne)
	t.Run("Services", testServicesOne)
}

func TestAll(t *testing.T) {
	t.Run("NetworkScans", testNetworkScansAll)
	t.Run("NetworkScansNodes", testNetworkScansNodesAll)
	t.Run("NodeNodeScansNetworkScans", testNodeNodeScansNetworkScansAll)
	t.Run("NodeScans", testNodeScansAll)
	t.Run("NodeScansServicesNodes", testNodeScansServicesNodesAll)
	t.Run("Nodes", testNodesAll)
	t.Run("PeriodicNetworkScansNetworkScans", testPeriodicNetworkScansNetworkScansAll)
	t.Run("Services", testServicesAll)
}

func TestCount(t *testing.T) {
	t.Run("NetworkScans", testNetworkScansCount)
	t.Run("NetworkScansNodes", testNetworkScansNodesCount)
	t.Run("NodeNodeScansNetworkScans", testNodeNodeScansNetworkScansCount)
	t.Run("NodeScans", testNodeScansCount)
	t.Run("NodeScansServicesNodes", testNodeScansServicesNodesCount)
	t.Run("Nodes", testNodesCount)
	t.Run("PeriodicNetworkScansNetworkScans", testPeriodicNetworkScansNetworkScansCount)
	t.Run("Services", testServicesCount)
}

func TestHooks(t *testing.T) {
	t.Run("NetworkScans", testNetworkScansHooks)
	t.Run("NetworkScansNodes", testNetworkScansNodesHooks)
	t.Run("NodeNodeScansNetworkScans", testNodeNodeScansNetworkScansHooks)
	t.Run("NodeScans", testNodeScansHooks)
	t.Run("NodeScansServicesNodes", testNodeScansServicesNodesHooks)
	t.Run("Nodes", testNodesHooks)
	t.Run("PeriodicNetworkScansNetworkScans", testPeriodicNetworkScansNetworkScansHooks)
	t.Run("Services", testServicesHooks)
}

func TestInsert(t *testing.T) {
	t.Run("NetworkScans", testNetworkScansInsert)
	t.Run("NetworkScans", testNetworkScansInsertWhitelist)
	t.Run("NetworkScansNodes", testNetworkScansNodesInsert)
	t.Run("NetworkScansNodes", testNetworkScansNodesInsertWhitelist)
	t.Run("NodeNodeScansNetworkScans", testNodeNodeScansNetworkScansInsert)
	t.Run("NodeNodeScansNetworkScans", testNodeNodeScansNetworkScansInsertWhitelist)
	t.Run("NodeScans", testNodeScansInsert)
	t.Run("NodeScans", testNodeScansInsertWhitelist)
	t.Run("NodeScansServicesNodes", testNodeScansServicesNodesInsert)
	t.Run("NodeScansServicesNodes", testNodeScansServicesNodesInsertWhitelist)
	t.Run("Nodes", testNodesInsert)
	t.Run("Nodes", testNodesInsertWhitelist)
	t.Run("PeriodicNetworkScansNetworkScans", testPeriodicNetworkScansNetworkScansInsert)
	t.Run("PeriodicNetworkScansNetworkScans", testPeriodicNetworkScansNetworkScansInsertWhitelist)
	t.Run("Services", testServicesInsert)
	t.Run("Services", testServicesInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("NetworkScans", testNetworkScansReload)
	t.Run("NetworkScansNodes", testNetworkScansNodesReload)
	t.Run("NodeNodeScansNetworkScans", testNodeNodeScansNetworkScansReload)
	t.Run("NodeScans", testNodeScansReload)
	t.Run("NodeScansServicesNodes", testNodeScansServicesNodesReload)
	t.Run("Nodes", testNodesReload)
	t.Run("PeriodicNetworkScansNetworkScans", testPeriodicNetworkScansNetworkScansReload)
	t.Run("Services", testServicesReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("NetworkScans", testNetworkScansReloadAll)
	t.Run("NetworkScansNodes", testNetworkScansNodesReloadAll)
	t.Run("NodeNodeScansNetworkScans", testNodeNodeScansNetworkScansReloadAll)
	t.Run("NodeScans", testNodeScansReloadAll)
	t.Run("NodeScansServicesNodes", testNodeScansServicesNodesReloadAll)
	t.Run("Nodes", testNodesReloadAll)
	t.Run("PeriodicNetworkScansNetworkScans", testPeriodicNetworkScansNetworkScansReloadAll)
	t.Run("Services", testServicesReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("NetworkScans", testNetworkScansSelect)
	t.Run("NetworkScansNodes", testNetworkScansNodesSelect)
	t.Run("NodeNodeScansNetworkScans", testNodeNodeScansNetworkScansSelect)
	t.Run("NodeScans", testNodeScansSelect)
	t.Run("NodeScansServicesNodes", testNodeScansServicesNodesSelect)
	t.Run("Nodes", testNodesSelect)
	t.Run("PeriodicNetworkScansNetworkScans", testPeriodicNetworkScansNetworkScansSelect)
	t.Run("Services", testServicesSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("NetworkScans", testNetworkScansUpdate)
	t.Run("NetworkScansNodes", testNetworkScansNodesUpdate)
	t.Run("NodeNodeScansNetworkScans", testNodeNodeScansNetworkScansUpdate)
	t.Run("NodeScans", testNodeScansUpdate)
	t.Run("NodeScansServicesNodes", testNodeScansServicesNodesUpdate)
	t.Run("Nodes", testNodesUpdate)
	t.Run("PeriodicNetworkScansNetworkScans", testPeriodicNetworkScansNetworkScansUpdate)
	t.Run("Services", testServicesUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("NetworkScans", testNetworkScansSliceUpdateAll)
	t.Run("NetworkScansNodes", testNetworkScansNodesSliceUpdateAll)
	t.Run("NodeNodeScansNetworkScans", testNodeNodeScansNetworkScansSliceUpdateAll)
	t.Run("NodeScans", testNodeScansSliceUpdateAll)
	t.Run("NodeScansServicesNodes", testNodeScansServicesNodesSliceUpdateAll)
	t.Run("Nodes", testNodesSliceUpdateAll)
	t.Run("PeriodicNetworkScansNetworkScans", testPeriodicNetworkScansNetworkScansSliceUpdateAll)
	t.Run("Services", testServicesSliceUpdateAll)
}
