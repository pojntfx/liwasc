// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testNodes(t *testing.T) {
	t.Parallel()

	query := Nodes()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testNodesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Node{}
	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Nodes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNodesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Node{}
	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Nodes().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Nodes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNodesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Node{}
	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := NodeSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Nodes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNodesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Node{}
	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := NodeExists(ctx, tx, o.MacAddress)
	if err != nil {
		t.Errorf("Unable to check if Node exists: %s", err)
	}
	if !e {
		t.Errorf("Expected NodeExists to return true, but got false.")
	}
}

func testNodesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Node{}
	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	nodeFound, err := FindNode(ctx, tx, o.MacAddress)
	if err != nil {
		t.Error(err)
	}

	if nodeFound == nil {
		t.Error("want a record, got nil")
	}
}

func testNodesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Node{}
	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Nodes().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testNodesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Node{}
	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Nodes().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testNodesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	nodeOne := &Node{}
	nodeTwo := &Node{}
	if err = randomize.Struct(seed, nodeOne, nodeDBTypes, false, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}
	if err = randomize.Struct(seed, nodeTwo, nodeDBTypes, false, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = nodeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = nodeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Nodes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testNodesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	nodeOne := &Node{}
	nodeTwo := &Node{}
	if err = randomize.Struct(seed, nodeOne, nodeDBTypes, false, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}
	if err = randomize.Struct(seed, nodeTwo, nodeDBTypes, false, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = nodeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = nodeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Nodes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func nodeBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Node) error {
	*o = Node{}
	return nil
}

func nodeAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Node) error {
	*o = Node{}
	return nil
}

func nodeAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Node) error {
	*o = Node{}
	return nil
}

func nodeBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Node) error {
	*o = Node{}
	return nil
}

func nodeAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Node) error {
	*o = Node{}
	return nil
}

func nodeBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Node) error {
	*o = Node{}
	return nil
}

func nodeAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Node) error {
	*o = Node{}
	return nil
}

func nodeBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Node) error {
	*o = Node{}
	return nil
}

func nodeAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Node) error {
	*o = Node{}
	return nil
}

func testNodesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Node{}
	o := &Node{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, nodeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Node object: %s", err)
	}

	AddNodeHook(boil.BeforeInsertHook, nodeBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	nodeBeforeInsertHooks = []NodeHook{}

	AddNodeHook(boil.AfterInsertHook, nodeAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	nodeAfterInsertHooks = []NodeHook{}

	AddNodeHook(boil.AfterSelectHook, nodeAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	nodeAfterSelectHooks = []NodeHook{}

	AddNodeHook(boil.BeforeUpdateHook, nodeBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	nodeBeforeUpdateHooks = []NodeHook{}

	AddNodeHook(boil.AfterUpdateHook, nodeAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	nodeAfterUpdateHooks = []NodeHook{}

	AddNodeHook(boil.BeforeDeleteHook, nodeBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	nodeBeforeDeleteHooks = []NodeHook{}

	AddNodeHook(boil.AfterDeleteHook, nodeAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	nodeAfterDeleteHooks = []NodeHook{}

	AddNodeHook(boil.BeforeUpsertHook, nodeBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	nodeBeforeUpsertHooks = []NodeHook{}

	AddNodeHook(boil.AfterUpsertHook, nodeAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	nodeAfterUpsertHooks = []NodeHook{}
}

func testNodesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Node{}
	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Nodes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testNodesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Node{}
	if err = randomize.Struct(seed, o, nodeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(nodeColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Nodes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testNodesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Node{}
	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testNodesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Node{}
	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := NodeSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testNodesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Node{}
	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Nodes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	nodeDBTypes = map[string]string{`MacAddress`: `TEXT`, `PoweredOn`: `INTEGER`}
	_           = bytes.MinRead
)

func testNodesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(nodePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(nodeAllColumns) == len(nodePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Node{}
	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Nodes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testNodesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(nodeAllColumns) == len(nodePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Node{}
	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Nodes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(nodeAllColumns, nodePrimaryKeyColumns) {
		fields = nodeAllColumns
	} else {
		fields = strmangle.SetComplement(
			nodeAllColumns,
			nodePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := NodeSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}
