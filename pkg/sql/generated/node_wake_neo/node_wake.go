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

// NodeWake is an object representing the database table.
type NodeWake struct {
	ID         int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt  time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	Done       int64     `boil:"done" json:"done" toml:"done" yaml:"done"`
	MacAddress string    `boil:"mac_address" json:"mac_address" toml:"mac_address" yaml:"mac_address"`
	PoweredOn  int64     `boil:"powered_on" json:"powered_on" toml:"powered_on" yaml:"powered_on"`

	R *nodeWakeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L nodeWakeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var NodeWakeColumns = struct {
	ID         string
	CreatedAt  string
	Done       string
	MacAddress string
	PoweredOn  string
}{
	ID:         "id",
	CreatedAt:  "created_at",
	Done:       "done",
	MacAddress: "mac_address",
	PoweredOn:  "powered_on",
}

// Generated where

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}

var NodeWakeWhere = struct {
	ID         whereHelperint64
	CreatedAt  whereHelpertime_Time
	Done       whereHelperint64
	MacAddress whereHelperstring
	PoweredOn  whereHelperint64
}{
	ID:         whereHelperint64{field: "\"node_wake\".\"id\""},
	CreatedAt:  whereHelpertime_Time{field: "\"node_wake\".\"created_at\""},
	Done:       whereHelperint64{field: "\"node_wake\".\"done\""},
	MacAddress: whereHelperstring{field: "\"node_wake\".\"mac_address\""},
	PoweredOn:  whereHelperint64{field: "\"node_wake\".\"powered_on\""},
}

// NodeWakeRels is where relationship names are stored.
var NodeWakeRels = struct {
}{}

// nodeWakeR is where relationships are stored.
type nodeWakeR struct {
}

// NewStruct creates a new relationship struct
func (*nodeWakeR) NewStruct() *nodeWakeR {
	return &nodeWakeR{}
}

// nodeWakeL is where Load methods for each relationship are stored.
type nodeWakeL struct{}

var (
	nodeWakeAllColumns            = []string{"id", "created_at", "done", "mac_address", "powered_on"}
	nodeWakeColumnsWithoutDefault = []string{"created_at", "done", "mac_address", "powered_on"}
	nodeWakeColumnsWithDefault    = []string{"id"}
	nodeWakePrimaryKeyColumns     = []string{"id"}
)

type (
	// NodeWakeSlice is an alias for a slice of pointers to NodeWake.
	// This should generally be used opposed to []NodeWake.
	NodeWakeSlice []*NodeWake
	// NodeWakeHook is the signature for custom NodeWake hook methods
	NodeWakeHook func(context.Context, boil.ContextExecutor, *NodeWake) error

	nodeWakeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	nodeWakeType                 = reflect.TypeOf(&NodeWake{})
	nodeWakeMapping              = queries.MakeStructMapping(nodeWakeType)
	nodeWakePrimaryKeyMapping, _ = queries.BindMapping(nodeWakeType, nodeWakeMapping, nodeWakePrimaryKeyColumns)
	nodeWakeInsertCacheMut       sync.RWMutex
	nodeWakeInsertCache          = make(map[string]insertCache)
	nodeWakeUpdateCacheMut       sync.RWMutex
	nodeWakeUpdateCache          = make(map[string]updateCache)
	nodeWakeUpsertCacheMut       sync.RWMutex
	nodeWakeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var nodeWakeBeforeInsertHooks []NodeWakeHook
var nodeWakeBeforeUpdateHooks []NodeWakeHook
var nodeWakeBeforeDeleteHooks []NodeWakeHook
var nodeWakeBeforeUpsertHooks []NodeWakeHook

var nodeWakeAfterInsertHooks []NodeWakeHook
var nodeWakeAfterSelectHooks []NodeWakeHook
var nodeWakeAfterUpdateHooks []NodeWakeHook
var nodeWakeAfterDeleteHooks []NodeWakeHook
var nodeWakeAfterUpsertHooks []NodeWakeHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *NodeWake) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nodeWakeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *NodeWake) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nodeWakeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *NodeWake) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nodeWakeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *NodeWake) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nodeWakeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *NodeWake) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nodeWakeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *NodeWake) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nodeWakeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *NodeWake) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nodeWakeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *NodeWake) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nodeWakeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *NodeWake) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nodeWakeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddNodeWakeHook registers your hook function for all future operations.
func AddNodeWakeHook(hookPoint boil.HookPoint, nodeWakeHook NodeWakeHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		nodeWakeBeforeInsertHooks = append(nodeWakeBeforeInsertHooks, nodeWakeHook)
	case boil.BeforeUpdateHook:
		nodeWakeBeforeUpdateHooks = append(nodeWakeBeforeUpdateHooks, nodeWakeHook)
	case boil.BeforeDeleteHook:
		nodeWakeBeforeDeleteHooks = append(nodeWakeBeforeDeleteHooks, nodeWakeHook)
	case boil.BeforeUpsertHook:
		nodeWakeBeforeUpsertHooks = append(nodeWakeBeforeUpsertHooks, nodeWakeHook)
	case boil.AfterInsertHook:
		nodeWakeAfterInsertHooks = append(nodeWakeAfterInsertHooks, nodeWakeHook)
	case boil.AfterSelectHook:
		nodeWakeAfterSelectHooks = append(nodeWakeAfterSelectHooks, nodeWakeHook)
	case boil.AfterUpdateHook:
		nodeWakeAfterUpdateHooks = append(nodeWakeAfterUpdateHooks, nodeWakeHook)
	case boil.AfterDeleteHook:
		nodeWakeAfterDeleteHooks = append(nodeWakeAfterDeleteHooks, nodeWakeHook)
	case boil.AfterUpsertHook:
		nodeWakeAfterUpsertHooks = append(nodeWakeAfterUpsertHooks, nodeWakeHook)
	}
}

// One returns a single nodeWake record from the query.
func (q nodeWakeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*NodeWake, error) {
	o := &NodeWake{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for node_wake")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all NodeWake records from the query.
func (q nodeWakeQuery) All(ctx context.Context, exec boil.ContextExecutor) (NodeWakeSlice, error) {
	var o []*NodeWake

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to NodeWake slice")
	}

	if len(nodeWakeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all NodeWake records in the query.
func (q nodeWakeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count node_wake rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q nodeWakeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if node_wake exists")
	}

	return count > 0, nil
}

// NodeWakes retrieves all the records using an executor.
func NodeWakes(mods ...qm.QueryMod) nodeWakeQuery {
	mods = append(mods, qm.From("\"node_wake\""))
	return nodeWakeQuery{NewQuery(mods...)}
}

// FindNodeWake retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindNodeWake(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*NodeWake, error) {
	nodeWakeObj := &NodeWake{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"node_wake\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, nodeWakeObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from node_wake")
	}

	return nodeWakeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *NodeWake) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no node_wake provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(nodeWakeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	nodeWakeInsertCacheMut.RLock()
	cache, cached := nodeWakeInsertCache[key]
	nodeWakeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			nodeWakeAllColumns,
			nodeWakeColumnsWithDefault,
			nodeWakeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(nodeWakeType, nodeWakeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(nodeWakeType, nodeWakeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"node_wake\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"node_wake\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"node_wake\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, nodeWakePrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into node_wake")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == nodeWakeMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for node_wake")
	}

CacheNoHooks:
	if !cached {
		nodeWakeInsertCacheMut.Lock()
		nodeWakeInsertCache[key] = cache
		nodeWakeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the NodeWake.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *NodeWake) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	nodeWakeUpdateCacheMut.RLock()
	cache, cached := nodeWakeUpdateCache[key]
	nodeWakeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			nodeWakeAllColumns,
			nodeWakePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update node_wake, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"node_wake\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, nodeWakePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(nodeWakeType, nodeWakeMapping, append(wl, nodeWakePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update node_wake row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for node_wake")
	}

	if !cached {
		nodeWakeUpdateCacheMut.Lock()
		nodeWakeUpdateCache[key] = cache
		nodeWakeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q nodeWakeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for node_wake")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for node_wake")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o NodeWakeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), nodeWakePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"node_wake\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, nodeWakePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in nodeWake slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all nodeWake")
	}
	return rowsAff, nil
}

// Delete deletes a single NodeWake record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *NodeWake) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no NodeWake provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), nodeWakePrimaryKeyMapping)
	sql := "DELETE FROM \"node_wake\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from node_wake")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for node_wake")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q nodeWakeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no nodeWakeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from node_wake")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for node_wake")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o NodeWakeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(nodeWakeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), nodeWakePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"node_wake\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, nodeWakePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from nodeWake slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for node_wake")
	}

	if len(nodeWakeAfterDeleteHooks) != 0 {
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
func (o *NodeWake) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindNodeWake(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *NodeWakeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := NodeWakeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), nodeWakePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"node_wake\".* FROM \"node_wake\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, nodeWakePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in NodeWakeSlice")
	}

	*o = slice

	return nil
}

// NodeWakeExists checks if the NodeWake row exists.
func NodeWakeExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"node_wake\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if node_wake exists")
	}

	return exists, nil
}
