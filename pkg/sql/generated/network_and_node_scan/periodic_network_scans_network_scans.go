// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// PeriodicNetworkScansNetworkScan is an object representing the database table.
type PeriodicNetworkScansNetworkScan struct {
	ID         int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt  time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	NodeScanID int64     `boil:"node_scan_id" json:"node_scan_id" toml:"node_scan_id" yaml:"node_scan_id"`

	R *periodicNetworkScansNetworkScanR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L periodicNetworkScansNetworkScanL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PeriodicNetworkScansNetworkScanColumns = struct {
	ID         string
	CreatedAt  string
	NodeScanID string
}{
	ID:         "id",
	CreatedAt:  "created_at",
	NodeScanID: "node_scan_id",
}

// Generated where

var PeriodicNetworkScansNetworkScanWhere = struct {
	ID         whereHelperint64
	CreatedAt  whereHelpertime_Time
	NodeScanID whereHelperint64
}{
	ID:         whereHelperint64{field: "\"periodic_network_scans_network_scans\".\"id\""},
	CreatedAt:  whereHelpertime_Time{field: "\"periodic_network_scans_network_scans\".\"created_at\""},
	NodeScanID: whereHelperint64{field: "\"periodic_network_scans_network_scans\".\"node_scan_id\""},
}

// PeriodicNetworkScansNetworkScanRels is where relationship names are stored.
var PeriodicNetworkScansNetworkScanRels = struct {
}{}

// periodicNetworkScansNetworkScanR is where relationships are stored.
type periodicNetworkScansNetworkScanR struct {
}

// NewStruct creates a new relationship struct
func (*periodicNetworkScansNetworkScanR) NewStruct() *periodicNetworkScansNetworkScanR {
	return &periodicNetworkScansNetworkScanR{}
}

// periodicNetworkScansNetworkScanL is where Load methods for each relationship are stored.
type periodicNetworkScansNetworkScanL struct{}

var (
	periodicNetworkScansNetworkScanAllColumns            = []string{"id", "created_at", "node_scan_id"}
	periodicNetworkScansNetworkScanColumnsWithoutDefault = []string{"created_at", "node_scan_id"}
	periodicNetworkScansNetworkScanColumnsWithDefault    = []string{"id"}
	periodicNetworkScansNetworkScanPrimaryKeyColumns     = []string{"id"}
)

type (
	// PeriodicNetworkScansNetworkScanSlice is an alias for a slice of pointers to PeriodicNetworkScansNetworkScan.
	// This should generally be used opposed to []PeriodicNetworkScansNetworkScan.
	PeriodicNetworkScansNetworkScanSlice []*PeriodicNetworkScansNetworkScan
	// PeriodicNetworkScansNetworkScanHook is the signature for custom PeriodicNetworkScansNetworkScan hook methods
	PeriodicNetworkScansNetworkScanHook func(context.Context, boil.ContextExecutor, *PeriodicNetworkScansNetworkScan) error

	periodicNetworkScansNetworkScanQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	periodicNetworkScansNetworkScanType                 = reflect.TypeOf(&PeriodicNetworkScansNetworkScan{})
	periodicNetworkScansNetworkScanMapping              = queries.MakeStructMapping(periodicNetworkScansNetworkScanType)
	periodicNetworkScansNetworkScanPrimaryKeyMapping, _ = queries.BindMapping(periodicNetworkScansNetworkScanType, periodicNetworkScansNetworkScanMapping, periodicNetworkScansNetworkScanPrimaryKeyColumns)
	periodicNetworkScansNetworkScanInsertCacheMut       sync.RWMutex
	periodicNetworkScansNetworkScanInsertCache          = make(map[string]insertCache)
	periodicNetworkScansNetworkScanUpdateCacheMut       sync.RWMutex
	periodicNetworkScansNetworkScanUpdateCache          = make(map[string]updateCache)
	periodicNetworkScansNetworkScanUpsertCacheMut       sync.RWMutex
	periodicNetworkScansNetworkScanUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var periodicNetworkScansNetworkScanBeforeInsertHooks []PeriodicNetworkScansNetworkScanHook
var periodicNetworkScansNetworkScanBeforeUpdateHooks []PeriodicNetworkScansNetworkScanHook
var periodicNetworkScansNetworkScanBeforeDeleteHooks []PeriodicNetworkScansNetworkScanHook
var periodicNetworkScansNetworkScanBeforeUpsertHooks []PeriodicNetworkScansNetworkScanHook

var periodicNetworkScansNetworkScanAfterInsertHooks []PeriodicNetworkScansNetworkScanHook
var periodicNetworkScansNetworkScanAfterSelectHooks []PeriodicNetworkScansNetworkScanHook
var periodicNetworkScansNetworkScanAfterUpdateHooks []PeriodicNetworkScansNetworkScanHook
var periodicNetworkScansNetworkScanAfterDeleteHooks []PeriodicNetworkScansNetworkScanHook
var periodicNetworkScansNetworkScanAfterUpsertHooks []PeriodicNetworkScansNetworkScanHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *PeriodicNetworkScansNetworkScan) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range periodicNetworkScansNetworkScanBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *PeriodicNetworkScansNetworkScan) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range periodicNetworkScansNetworkScanBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *PeriodicNetworkScansNetworkScan) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range periodicNetworkScansNetworkScanBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *PeriodicNetworkScansNetworkScan) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range periodicNetworkScansNetworkScanBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *PeriodicNetworkScansNetworkScan) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range periodicNetworkScansNetworkScanAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *PeriodicNetworkScansNetworkScan) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range periodicNetworkScansNetworkScanAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *PeriodicNetworkScansNetworkScan) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range periodicNetworkScansNetworkScanAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *PeriodicNetworkScansNetworkScan) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range periodicNetworkScansNetworkScanAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *PeriodicNetworkScansNetworkScan) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range periodicNetworkScansNetworkScanAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPeriodicNetworkScansNetworkScanHook registers your hook function for all future operations.
func AddPeriodicNetworkScansNetworkScanHook(hookPoint boil.HookPoint, periodicNetworkScansNetworkScanHook PeriodicNetworkScansNetworkScanHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		periodicNetworkScansNetworkScanBeforeInsertHooks = append(periodicNetworkScansNetworkScanBeforeInsertHooks, periodicNetworkScansNetworkScanHook)
	case boil.BeforeUpdateHook:
		periodicNetworkScansNetworkScanBeforeUpdateHooks = append(periodicNetworkScansNetworkScanBeforeUpdateHooks, periodicNetworkScansNetworkScanHook)
	case boil.BeforeDeleteHook:
		periodicNetworkScansNetworkScanBeforeDeleteHooks = append(periodicNetworkScansNetworkScanBeforeDeleteHooks, periodicNetworkScansNetworkScanHook)
	case boil.BeforeUpsertHook:
		periodicNetworkScansNetworkScanBeforeUpsertHooks = append(periodicNetworkScansNetworkScanBeforeUpsertHooks, periodicNetworkScansNetworkScanHook)
	case boil.AfterInsertHook:
		periodicNetworkScansNetworkScanAfterInsertHooks = append(periodicNetworkScansNetworkScanAfterInsertHooks, periodicNetworkScansNetworkScanHook)
	case boil.AfterSelectHook:
		periodicNetworkScansNetworkScanAfterSelectHooks = append(periodicNetworkScansNetworkScanAfterSelectHooks, periodicNetworkScansNetworkScanHook)
	case boil.AfterUpdateHook:
		periodicNetworkScansNetworkScanAfterUpdateHooks = append(periodicNetworkScansNetworkScanAfterUpdateHooks, periodicNetworkScansNetworkScanHook)
	case boil.AfterDeleteHook:
		periodicNetworkScansNetworkScanAfterDeleteHooks = append(periodicNetworkScansNetworkScanAfterDeleteHooks, periodicNetworkScansNetworkScanHook)
	case boil.AfterUpsertHook:
		periodicNetworkScansNetworkScanAfterUpsertHooks = append(periodicNetworkScansNetworkScanAfterUpsertHooks, periodicNetworkScansNetworkScanHook)
	}
}

// One returns a single periodicNetworkScansNetworkScan record from the query.
func (q periodicNetworkScansNetworkScanQuery) One(ctx context.Context, exec boil.ContextExecutor) (*PeriodicNetworkScansNetworkScan, error) {
	o := &PeriodicNetworkScansNetworkScan{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for periodic_network_scans_network_scans")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all PeriodicNetworkScansNetworkScan records from the query.
func (q periodicNetworkScansNetworkScanQuery) All(ctx context.Context, exec boil.ContextExecutor) (PeriodicNetworkScansNetworkScanSlice, error) {
	var o []*PeriodicNetworkScansNetworkScan

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to PeriodicNetworkScansNetworkScan slice")
	}

	if len(periodicNetworkScansNetworkScanAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all PeriodicNetworkScansNetworkScan records in the query.
func (q periodicNetworkScansNetworkScanQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count periodic_network_scans_network_scans rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q periodicNetworkScansNetworkScanQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if periodic_network_scans_network_scans exists")
	}

	return count > 0, nil
}

// PeriodicNetworkScansNetworkScans retrieves all the records using an executor.
func PeriodicNetworkScansNetworkScans(mods ...qm.QueryMod) periodicNetworkScansNetworkScanQuery {
	mods = append(mods, qm.From("\"periodic_network_scans_network_scans\""))
	return periodicNetworkScansNetworkScanQuery{NewQuery(mods...)}
}

// FindPeriodicNetworkScansNetworkScan retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPeriodicNetworkScansNetworkScan(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*PeriodicNetworkScansNetworkScan, error) {
	periodicNetworkScansNetworkScanObj := &PeriodicNetworkScansNetworkScan{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"periodic_network_scans_network_scans\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, periodicNetworkScansNetworkScanObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from periodic_network_scans_network_scans")
	}

	return periodicNetworkScansNetworkScanObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *PeriodicNetworkScansNetworkScan) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no periodic_network_scans_network_scans provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(periodicNetworkScansNetworkScanColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	periodicNetworkScansNetworkScanInsertCacheMut.RLock()
	cache, cached := periodicNetworkScansNetworkScanInsertCache[key]
	periodicNetworkScansNetworkScanInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			periodicNetworkScansNetworkScanAllColumns,
			periodicNetworkScansNetworkScanColumnsWithDefault,
			periodicNetworkScansNetworkScanColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(periodicNetworkScansNetworkScanType, periodicNetworkScansNetworkScanMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(periodicNetworkScansNetworkScanType, periodicNetworkScansNetworkScanMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"periodic_network_scans_network_scans\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"periodic_network_scans_network_scans\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"periodic_network_scans_network_scans\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, periodicNetworkScansNetworkScanPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into periodic_network_scans_network_scans")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == periodicNetworkScansNetworkScanMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for periodic_network_scans_network_scans")
	}

CacheNoHooks:
	if !cached {
		periodicNetworkScansNetworkScanInsertCacheMut.Lock()
		periodicNetworkScansNetworkScanInsertCache[key] = cache
		periodicNetworkScansNetworkScanInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the PeriodicNetworkScansNetworkScan.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *PeriodicNetworkScansNetworkScan) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	periodicNetworkScansNetworkScanUpdateCacheMut.RLock()
	cache, cached := periodicNetworkScansNetworkScanUpdateCache[key]
	periodicNetworkScansNetworkScanUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			periodicNetworkScansNetworkScanAllColumns,
			periodicNetworkScansNetworkScanPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update periodic_network_scans_network_scans, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"periodic_network_scans_network_scans\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, periodicNetworkScansNetworkScanPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(periodicNetworkScansNetworkScanType, periodicNetworkScansNetworkScanMapping, append(wl, periodicNetworkScansNetworkScanPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update periodic_network_scans_network_scans row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for periodic_network_scans_network_scans")
	}

	if !cached {
		periodicNetworkScansNetworkScanUpdateCacheMut.Lock()
		periodicNetworkScansNetworkScanUpdateCache[key] = cache
		periodicNetworkScansNetworkScanUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q periodicNetworkScansNetworkScanQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for periodic_network_scans_network_scans")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for periodic_network_scans_network_scans")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PeriodicNetworkScansNetworkScanSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), periodicNetworkScansNetworkScanPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"periodic_network_scans_network_scans\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, periodicNetworkScansNetworkScanPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in periodicNetworkScansNetworkScan slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all periodicNetworkScansNetworkScan")
	}
	return rowsAff, nil
}

// Delete deletes a single PeriodicNetworkScansNetworkScan record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *PeriodicNetworkScansNetworkScan) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no PeriodicNetworkScansNetworkScan provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), periodicNetworkScansNetworkScanPrimaryKeyMapping)
	sql := "DELETE FROM \"periodic_network_scans_network_scans\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from periodic_network_scans_network_scans")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for periodic_network_scans_network_scans")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q periodicNetworkScansNetworkScanQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no periodicNetworkScansNetworkScanQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from periodic_network_scans_network_scans")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for periodic_network_scans_network_scans")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PeriodicNetworkScansNetworkScanSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(periodicNetworkScansNetworkScanBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), periodicNetworkScansNetworkScanPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"periodic_network_scans_network_scans\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, periodicNetworkScansNetworkScanPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from periodicNetworkScansNetworkScan slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for periodic_network_scans_network_scans")
	}

	if len(periodicNetworkScansNetworkScanAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *PeriodicNetworkScansNetworkScan) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPeriodicNetworkScansNetworkScan(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PeriodicNetworkScansNetworkScanSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PeriodicNetworkScansNetworkScanSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), periodicNetworkScansNetworkScanPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"periodic_network_scans_network_scans\".* FROM \"periodic_network_scans_network_scans\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, periodicNetworkScansNetworkScanPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PeriodicNetworkScansNetworkScanSlice")
	}

	*o = slice

	return nil
}

// PeriodicNetworkScansNetworkScanExists checks if the PeriodicNetworkScansNetworkScan row exists.
func PeriodicNetworkScansNetworkScanExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"periodic_network_scans_network_scans\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if periodic_network_scans_network_scans exists")
	}

	return exists, nil
}
