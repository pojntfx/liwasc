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

func testScans(t *testing.T) {
	t.Parallel()

	query := Scans()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testScansDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Scan{}
	if err = randomize.Struct(seed, o, scanDBTypes, true, scanColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Scan struct: %s", err)
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

	count, err := Scans().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testScansQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Scan{}
	if err = randomize.Struct(seed, o, scanDBTypes, true, scanColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Scan struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Scans().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Scans().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testScansSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Scan{}
	if err = randomize.Struct(seed, o, scanDBTypes, true, scanColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Scan struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ScanSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Scans().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testScansExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Scan{}
	if err = randomize.Struct(seed, o, scanDBTypes, true, scanColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Scan struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ScanExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Scan exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ScanExists to return true, but got false.")
	}
}

func testScansFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Scan{}
	if err = randomize.Struct(seed, o, scanDBTypes, true, scanColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Scan struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	scanFound, err := FindScan(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if scanFound == nil {
		t.Error("want a record, got nil")
	}
}

func testScansBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Scan{}
	if err = randomize.Struct(seed, o, scanDBTypes, true, scanColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Scan struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Scans().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testScansOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Scan{}
	if err = randomize.Struct(seed, o, scanDBTypes, true, scanColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Scan struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Scans().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testScansAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	scanOne := &Scan{}
	scanTwo := &Scan{}
	if err = randomize.Struct(seed, scanOne, scanDBTypes, false, scanColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Scan struct: %s", err)
	}
	if err = randomize.Struct(seed, scanTwo, scanDBTypes, false, scanColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Scan struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = scanOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = scanTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Scans().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testScansCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	scanOne := &Scan{}
	scanTwo := &Scan{}
	if err = randomize.Struct(seed, scanOne, scanDBTypes, false, scanColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Scan struct: %s", err)
	}
	if err = randomize.Struct(seed, scanTwo, scanDBTypes, false, scanColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Scan struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = scanOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = scanTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Scans().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func scanBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Scan) error {
	*o = Scan{}
	return nil
}

func scanAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Scan) error {
	*o = Scan{}
	return nil
}

func scanAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Scan) error {
	*o = Scan{}
	return nil
}

func scanBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Scan) error {
	*o = Scan{}
	return nil
}

func scanAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Scan) error {
	*o = Scan{}
	return nil
}

func scanBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Scan) error {
	*o = Scan{}
	return nil
}

func scanAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Scan) error {
	*o = Scan{}
	return nil
}

func scanBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Scan) error {
	*o = Scan{}
	return nil
}

func scanAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Scan) error {
	*o = Scan{}
	return nil
}

func testScansHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Scan{}
	o := &Scan{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, scanDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Scan object: %s", err)
	}

	AddScanHook(boil.BeforeInsertHook, scanBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	scanBeforeInsertHooks = []ScanHook{}

	AddScanHook(boil.AfterInsertHook, scanAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	scanAfterInsertHooks = []ScanHook{}

	AddScanHook(boil.AfterSelectHook, scanAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	scanAfterSelectHooks = []ScanHook{}

	AddScanHook(boil.BeforeUpdateHook, scanBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	scanBeforeUpdateHooks = []ScanHook{}

	AddScanHook(boil.AfterUpdateHook, scanAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	scanAfterUpdateHooks = []ScanHook{}

	AddScanHook(boil.BeforeDeleteHook, scanBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	scanBeforeDeleteHooks = []ScanHook{}

	AddScanHook(boil.AfterDeleteHook, scanAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	scanAfterDeleteHooks = []ScanHook{}

	AddScanHook(boil.BeforeUpsertHook, scanBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	scanBeforeUpsertHooks = []ScanHook{}

	AddScanHook(boil.AfterUpsertHook, scanAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	scanAfterUpsertHooks = []ScanHook{}
}

func testScansInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Scan{}
	if err = randomize.Struct(seed, o, scanDBTypes, true, scanColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Scan struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Scans().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testScansInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Scan{}
	if err = randomize.Struct(seed, o, scanDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Scan struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(scanColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Scans().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testScanToManyNodes(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Scan
	var b, c Node

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, scanDBTypes, true, scanColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Scan struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, nodeDBTypes, false, nodeColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, nodeDBTypes, false, nodeColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.ScanID = a.ID
	c.ScanID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Nodes().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ScanID == b.ScanID {
			bFound = true
		}
		if v.ScanID == c.ScanID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ScanSlice{&a}
	if err = a.L.LoadNodes(ctx, tx, false, (*[]*Scan)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Nodes); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Nodes = nil
	if err = a.L.LoadNodes(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Nodes); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testScanToManyAddOpNodes(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Scan
	var b, c, d, e Node

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, scanDBTypes, false, strmangle.SetComplement(scanPrimaryKeyColumns, scanColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Node{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, nodeDBTypes, false, strmangle.SetComplement(nodePrimaryKeyColumns, nodeColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Node{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddNodes(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.ScanID {
			t.Error("foreign key was wrong value", a.ID, first.ScanID)
		}
		if a.ID != second.ScanID {
			t.Error("foreign key was wrong value", a.ID, second.ScanID)
		}

		if first.R.Scan != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Scan != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Nodes[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Nodes[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Nodes().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testScansReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Scan{}
	if err = randomize.Struct(seed, o, scanDBTypes, true, scanColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Scan struct: %s", err)
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

func testScansReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Scan{}
	if err = randomize.Struct(seed, o, scanDBTypes, true, scanColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Scan struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ScanSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testScansSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Scan{}
	if err = randomize.Struct(seed, o, scanDBTypes, true, scanColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Scan struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Scans().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	scanDBTypes = map[string]string{`ID`: `INTEGER`}
	_           = bytes.MinRead
)

func testScansUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(scanPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(scanAllColumns) == len(scanPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Scan{}
	if err = randomize.Struct(seed, o, scanDBTypes, true, scanColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Scan struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Scans().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, scanDBTypes, true, scanPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Scan struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testScansSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(scanAllColumns) == len(scanPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Scan{}
	if err = randomize.Struct(seed, o, scanDBTypes, true, scanColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Scan struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Scans().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, scanDBTypes, true, scanPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Scan struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(scanAllColumns, scanPrimaryKeyColumns) {
		fields = scanAllColumns
	} else {
		fields = strmangle.SetComplement(
			scanAllColumns,
			scanPrimaryKeyColumns,
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

	slice := ScanSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}
