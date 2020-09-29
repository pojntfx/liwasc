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

// NodeScansServicesNode is an object representing the database table.
type NodeScansServicesNode struct {
	ID         int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt  time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	ServiceID  int64     `boil:"service_id" json:"service_id" toml:"service_id" yaml:"service_id"`
	NodeID     string    `boil:"node_id" json:"node_id" toml:"node_id" yaml:"node_id"`
	NodeScanID int64     `boil:"node_scan_id" json:"node_scan_id" toml:"node_scan_id" yaml:"node_scan_id"`

	R *nodeScansServicesNodeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L nodeScansServicesNodeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var NodeScansServicesNodeColumns = struct {
	ID         string
	CreatedAt  string
	ServiceID  string
	NodeID     string
	NodeScanID string
}{
	ID:         "id",
	CreatedAt:  "created_at",
	ServiceID:  "service_id",
	NodeID:     "node_id",
	NodeScanID: "node_scan_id",
}

// Generated where

var NodeScansServicesNodeWhere = struct {
	ID         whereHelperint64
	CreatedAt  whereHelpertime_Time
	ServiceID  whereHelperint64
	NodeID     whereHelperstring
	NodeScanID whereHelperint64
}{
	ID:         whereHelperint64{field: "\"node_scans_services_nodes\".\"id\""},
	CreatedAt:  whereHelpertime_Time{field: "\"node_scans_services_nodes\".\"created_at\""},
	ServiceID:  whereHelperint64{field: "\"node_scans_services_nodes\".\"service_id\""},
	NodeID:     whereHelperstring{field: "\"node_scans_services_nodes\".\"node_id\""},
	NodeScanID: whereHelperint64{field: "\"node_scans_services_nodes\".\"node_scan_id\""},
}

// NodeScansServicesNodeRels is where relationship names are stored.
var NodeScansServicesNodeRels = struct {
}{}

// nodeScansServicesNodeR is where relationships are stored.
type nodeScansServicesNodeR struct {
}

// NewStruct creates a new relationship struct
func (*nodeScansServicesNodeR) NewStruct() *nodeScansServicesNodeR {
	return &nodeScansServicesNodeR{}
}

// nodeScansServicesNodeL is where Load methods for each relationship are stored.
type nodeScansServicesNodeL struct{}

var (
	nodeScansServicesNodeAllColumns            = []string{"id", "created_at", "service_id", "node_id", "node_scan_id"}
	nodeScansServicesNodeColumnsWithoutDefault = []string{"created_at", "service_id", "node_id", "node_scan_id"}
	nodeScansServicesNodeColumnsWithDefault    = []string{"id"}
	nodeScansServicesNodePrimaryKeyColumns     = []string{"id"}
)

type (
	// NodeScansServicesNodeSlice is an alias for a slice of pointers to NodeScansServicesNode.
	// This should generally be used opposed to []NodeScansServicesNode.
	NodeScansServicesNodeSlice []*NodeScansServicesNode
	// NodeScansServicesNodeHook is the signature for custom NodeScansServicesNode hook methods
	NodeScansServicesNodeHook func(context.Context, boil.ContextExecutor, *NodeScansServicesNode) error

	nodeScansServicesNodeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	nodeScansServicesNodeType                 = reflect.TypeOf(&NodeScansServicesNode{})
	nodeScansServicesNodeMapping              = queries.MakeStructMapping(nodeScansServicesNodeType)
	nodeScansServicesNodePrimaryKeyMapping, _ = queries.BindMapping(nodeScansServicesNodeType, nodeScansServicesNodeMapping, nodeScansServicesNodePrimaryKeyColumns)
	nodeScansServicesNodeInsertCacheMut       sync.RWMutex
	nodeScansServicesNodeInsertCache          = make(map[string]insertCache)
	nodeScansServicesNodeUpdateCacheMut       sync.RWMutex
	nodeScansServicesNodeUpdateCache          = make(map[string]updateCache)
	nodeScansServicesNodeUpsertCacheMut       sync.RWMutex
	nodeScansServicesNodeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var nodeScansServicesNodeBeforeInsertHooks []NodeScansServicesNodeHook
var nodeScansServicesNodeBeforeUpdateHooks []NodeScansServicesNodeHook
var nodeScansServicesNodeBeforeDeleteHooks []NodeScansServicesNodeHook
var nodeScansServicesNodeBeforeUpsertHooks []NodeScansServicesNodeHook

var nodeScansServicesNodeAfterInsertHooks []NodeScansServicesNodeHook
var nodeScansServicesNodeAfterSelectHooks []NodeScansServicesNodeHook
var nodeScansServicesNodeAfterUpdateHooks []NodeScansServicesNodeHook
var nodeScansServicesNodeAfterDeleteHooks []NodeScansServicesNodeHook
var nodeScansServicesNodeAfterUpsertHooks []NodeScansServicesNodeHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *NodeScansServicesNode) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nodeScansServicesNodeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *NodeScansServicesNode) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nodeScansServicesNodeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *NodeScansServicesNode) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nodeScansServicesNodeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *NodeScansServicesNode) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nodeScansServicesNodeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *NodeScansServicesNode) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nodeScansServicesNodeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *NodeScansServicesNode) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nodeScansServicesNodeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *NodeScansServicesNode) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nodeScansServicesNodeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *NodeScansServicesNode) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nodeScansServicesNodeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *NodeScansServicesNode) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nodeScansServicesNodeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddNodeScansServicesNodeHook registers your hook function for all future operations.
func AddNodeScansServicesNodeHook(hookPoint boil.HookPoint, nodeScansServicesNodeHook NodeScansServicesNodeHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		nodeScansServicesNodeBeforeInsertHooks = append(nodeScansServicesNodeBeforeInsertHooks, nodeScansServicesNodeHook)
	case boil.BeforeUpdateHook:
		nodeScansServicesNodeBeforeUpdateHooks = append(nodeScansServicesNodeBeforeUpdateHooks, nodeScansServicesNodeHook)
	case boil.BeforeDeleteHook:
		nodeScansServicesNodeBeforeDeleteHooks = append(nodeScansServicesNodeBeforeDeleteHooks, nodeScansServicesNodeHook)
	case boil.BeforeUpsertHook:
		nodeScansServicesNodeBeforeUpsertHooks = append(nodeScansServicesNodeBeforeUpsertHooks, nodeScansServicesNodeHook)
	case boil.AfterInsertHook:
		nodeScansServicesNodeAfterInsertHooks = append(nodeScansServicesNodeAfterInsertHooks, nodeScansServicesNodeHook)
	case boil.AfterSelectHook:
		nodeScansServicesNodeAfterSelectHooks = append(nodeScansServicesNodeAfterSelectHooks, nodeScansServicesNodeHook)
	case boil.AfterUpdateHook:
		nodeScansServicesNodeAfterUpdateHooks = append(nodeScansServicesNodeAfterUpdateHooks, nodeScansServicesNodeHook)
	case boil.AfterDeleteHook:
		nodeScansServicesNodeAfterDeleteHooks = append(nodeScansServicesNodeAfterDeleteHooks, nodeScansServicesNodeHook)
	case boil.AfterUpsertHook:
		nodeScansServicesNodeAfterUpsertHooks = append(nodeScansServicesNodeAfterUpsertHooks, nodeScansServicesNodeHook)
	}
}

// One returns a single nodeScansServicesNode record from the query.
func (q nodeScansServicesNodeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*NodeScansServicesNode, error) {
	o := &NodeScansServicesNode{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for node_scans_services_nodes")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all NodeScansServicesNode records from the query.
func (q nodeScansServicesNodeQuery) All(ctx context.Context, exec boil.ContextExecutor) (NodeScansServicesNodeSlice, error) {
	var o []*NodeScansServicesNode

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to NodeScansServicesNode slice")
	}

	if len(nodeScansServicesNodeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all NodeScansServicesNode records in the query.
func (q nodeScansServicesNodeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count node_scans_services_nodes rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q nodeScansServicesNodeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if node_scans_services_nodes exists")
	}

	return count > 0, nil
}

// NodeScansServicesNodes retrieves all the records using an executor.
func NodeScansServicesNodes(mods ...qm.QueryMod) nodeScansServicesNodeQuery {
	mods = append(mods, qm.From("\"node_scans_services_nodes\""))
	return nodeScansServicesNodeQuery{NewQuery(mods...)}
}

// FindNodeScansServicesNode retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindNodeScansServicesNode(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*NodeScansServicesNode, error) {
	nodeScansServicesNodeObj := &NodeScansServicesNode{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"node_scans_services_nodes\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, nodeScansServicesNodeObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from node_scans_services_nodes")
	}

	return nodeScansServicesNodeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *NodeScansServicesNode) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no node_scans_services_nodes provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(nodeScansServicesNodeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	nodeScansServicesNodeInsertCacheMut.RLock()
	cache, cached := nodeScansServicesNodeInsertCache[key]
	nodeScansServicesNodeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			nodeScansServicesNodeAllColumns,
			nodeScansServicesNodeColumnsWithDefault,
			nodeScansServicesNodeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(nodeScansServicesNodeType, nodeScansServicesNodeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(nodeScansServicesNodeType, nodeScansServicesNodeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"node_scans_services_nodes\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"node_scans_services_nodes\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"node_scans_services_nodes\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, nodeScansServicesNodePrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into node_scans_services_nodes")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == nodeScansServicesNodeMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for node_scans_services_nodes")
	}

CacheNoHooks:
	if !cached {
		nodeScansServicesNodeInsertCacheMut.Lock()
		nodeScansServicesNodeInsertCache[key] = cache
		nodeScansServicesNodeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the NodeScansServicesNode.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *NodeScansServicesNode) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	nodeScansServicesNodeUpdateCacheMut.RLock()
	cache, cached := nodeScansServicesNodeUpdateCache[key]
	nodeScansServicesNodeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			nodeScansServicesNodeAllColumns,
			nodeScansServicesNodePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update node_scans_services_nodes, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"node_scans_services_nodes\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, nodeScansServicesNodePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(nodeScansServicesNodeType, nodeScansServicesNodeMapping, append(wl, nodeScansServicesNodePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update node_scans_services_nodes row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for node_scans_services_nodes")
	}

	if !cached {
		nodeScansServicesNodeUpdateCacheMut.Lock()
		nodeScansServicesNodeUpdateCache[key] = cache
		nodeScansServicesNodeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q nodeScansServicesNodeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for node_scans_services_nodes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for node_scans_services_nodes")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o NodeScansServicesNodeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), nodeScansServicesNodePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"node_scans_services_nodes\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, nodeScansServicesNodePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in nodeScansServicesNode slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all nodeScansServicesNode")
	}
	return rowsAff, nil
}

// Delete deletes a single NodeScansServicesNode record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *NodeScansServicesNode) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no NodeScansServicesNode provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), nodeScansServicesNodePrimaryKeyMapping)
	sql := "DELETE FROM \"node_scans_services_nodes\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from node_scans_services_nodes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for node_scans_services_nodes")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q nodeScansServicesNodeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no nodeScansServicesNodeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from node_scans_services_nodes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for node_scans_services_nodes")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o NodeScansServicesNodeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(nodeScansServicesNodeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), nodeScansServicesNodePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"node_scans_services_nodes\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, nodeScansServicesNodePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from nodeScansServicesNode slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for node_scans_services_nodes")
	}

	if len(nodeScansServicesNodeAfterDeleteHooks) != 0 {
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
func (o *NodeScansServicesNode) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindNodeScansServicesNode(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *NodeScansServicesNodeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := NodeScansServicesNodeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), nodeScansServicesNodePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"node_scans_services_nodes\".* FROM \"node_scans_services_nodes\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, nodeScansServicesNodePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in NodeScansServicesNodeSlice")
	}

	*o = slice

	return nil
}

// NodeScansServicesNodeExists checks if the NodeScansServicesNode row exists.
func NodeScansServicesNodeExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"node_scans_services_nodes\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if node_scans_services_nodes exists")
	}

	return exists, nil
}
