// This file is generated by SQLBoiler (https://github.com/vattle/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// DO NOT EDIT

package models

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/databrary/databrary/db/models/custom_types"
	"github.com/pkg/errors"
	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/queries"
	"github.com/vattle/sqlboiler/queries/qm"
	"github.com/vattle/sqlboiler/strmangle"
	"gopkg.in/nullbio/null.v6"
)

// Authorize is an object representing the database table.
type Authorize struct {
	Child   int                     `boil:"child" json:"authorize_child"`
	Parent  int                     `boil:"parent" json:"authorize_parent"`
	Site    custom_types.Permission `boil:"site" json:"authorize_site"`
	Member  custom_types.Permission `boil:"member" json:"authorize_member"`
	Expires null.Time               `boil:"expires" json:"authorize_expires,omitempty"`

	R *authorizeR `boil:"-" json:"-"`
	L authorizeL  `boil:"-" json:"-"`
}

// authorizeR is where relationships are stored.
type authorizeR struct {
	Parent *Party
	Child  *Party
}

// authorizeL is where Load methods for each relationship are stored.
type authorizeL struct{}

var (
	authorizeColumns               = []string{"child", "parent", "site", "member", "expires"}
	authorizeColumnsWithoutDefault = []string{"child", "parent", "expires"}
	authorizeColumnsWithDefault    = []string{"site", "member"}
	authorizePrimaryKeyColumns     = []string{"child", "parent"}
)

type (
	// AuthorizeSlice is an alias for a slice of pointers to Authorize.
	// This should generally be used opposed to []Authorize.
	AuthorizeSlice []*Authorize
	// AuthorizeHook is the signature for custom Authorize hook methods
	AuthorizeHook func(boil.Executor, *Authorize) error

	authorizeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	authorizeType                 = reflect.TypeOf(&Authorize{})
	authorizeMapping              = queries.MakeStructMapping(authorizeType)
	authorizePrimaryKeyMapping, _ = queries.BindMapping(authorizeType, authorizeMapping, authorizePrimaryKeyColumns)
	authorizeInsertCacheMut       sync.RWMutex
	authorizeInsertCache          = make(map[string]insertCache)
	authorizeUpdateCacheMut       sync.RWMutex
	authorizeUpdateCache          = make(map[string]updateCache)
	authorizeUpsertCacheMut       sync.RWMutex
	authorizeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var authorizeBeforeInsertHooks []AuthorizeHook
var authorizeBeforeUpdateHooks []AuthorizeHook
var authorizeBeforeDeleteHooks []AuthorizeHook
var authorizeBeforeUpsertHooks []AuthorizeHook

var authorizeAfterInsertHooks []AuthorizeHook
var authorizeAfterSelectHooks []AuthorizeHook
var authorizeAfterUpdateHooks []AuthorizeHook
var authorizeAfterDeleteHooks []AuthorizeHook
var authorizeAfterUpsertHooks []AuthorizeHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Authorize) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range authorizeBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Authorize) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range authorizeBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Authorize) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range authorizeBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Authorize) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range authorizeBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Authorize) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range authorizeAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Authorize) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range authorizeAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Authorize) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range authorizeAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Authorize) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range authorizeAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Authorize) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range authorizeAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAuthorizeHook registers your hook function for all future operations.
func AddAuthorizeHook(hookPoint boil.HookPoint, authorizeHook AuthorizeHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		authorizeBeforeInsertHooks = append(authorizeBeforeInsertHooks, authorizeHook)
	case boil.BeforeUpdateHook:
		authorizeBeforeUpdateHooks = append(authorizeBeforeUpdateHooks, authorizeHook)
	case boil.BeforeDeleteHook:
		authorizeBeforeDeleteHooks = append(authorizeBeforeDeleteHooks, authorizeHook)
	case boil.BeforeUpsertHook:
		authorizeBeforeUpsertHooks = append(authorizeBeforeUpsertHooks, authorizeHook)
	case boil.AfterInsertHook:
		authorizeAfterInsertHooks = append(authorizeAfterInsertHooks, authorizeHook)
	case boil.AfterSelectHook:
		authorizeAfterSelectHooks = append(authorizeAfterSelectHooks, authorizeHook)
	case boil.AfterUpdateHook:
		authorizeAfterUpdateHooks = append(authorizeAfterUpdateHooks, authorizeHook)
	case boil.AfterDeleteHook:
		authorizeAfterDeleteHooks = append(authorizeAfterDeleteHooks, authorizeHook)
	case boil.AfterUpsertHook:
		authorizeAfterUpsertHooks = append(authorizeAfterUpsertHooks, authorizeHook)
	}
}

// OneP returns a single authorize record from the query, and panics on error.
func (q authorizeQuery) OneP() *Authorize {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single authorize record from the query.
func (q authorizeQuery) One() (*Authorize, error) {
	o := &Authorize{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for authorize")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Authorize records from the query, and panics on error.
func (q authorizeQuery) AllP() AuthorizeSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Authorize records from the query.
func (q authorizeQuery) All() (AuthorizeSlice, error) {
	var o AuthorizeSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Authorize slice")
	}

	if len(authorizeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Authorize records in the query, and panics on error.
func (q authorizeQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Authorize records in the query.
func (q authorizeQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count authorize rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q authorizeQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q authorizeQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if authorize exists")
	}

	return count > 0, nil
}

// ParentG pointed to by the foreign key.
func (o *Authorize) ParentG(mods ...qm.QueryMod) partyQuery {
	return o.ParentByFk(boil.GetDB(), mods...)
}

// Parent pointed to by the foreign key.
func (o *Authorize) ParentByFk(exec boil.Executor, mods ...qm.QueryMod) partyQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.Parent),
	}

	queryMods = append(queryMods, mods...)

	query := Parties(exec, queryMods...)
	queries.SetFrom(query.Query, "\"party\"")

	return query
}

// ChildG pointed to by the foreign key.
func (o *Authorize) ChildG(mods ...qm.QueryMod) partyQuery {
	return o.ChildByFk(boil.GetDB(), mods...)
}

// Child pointed to by the foreign key.
func (o *Authorize) ChildByFk(exec boil.Executor, mods ...qm.QueryMod) partyQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.Child),
	}

	queryMods = append(queryMods, mods...)

	query := Parties(exec, queryMods...)
	queries.SetFrom(query.Query, "\"party\"")

	return query
}

// LoadParent allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (authorizeL) LoadParent(e boil.Executor, singular bool, maybeAuthorize interface{}) error {
	var slice []*Authorize
	var object *Authorize

	count := 1
	if singular {
		object = maybeAuthorize.(*Authorize)
	} else {
		slice = *maybeAuthorize.(*AuthorizeSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &authorizeR{}
		}
		args[0] = object.Parent
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &authorizeR{}
			}
			args[i] = obj.Parent
		}
	}

	query := fmt.Sprintf(
		"select * from \"party\" where \"id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Party")
	}
	defer results.Close()

	var resultSlice []*Party
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Party")
	}

	if len(authorizeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		object.R.Parent = resultSlice[0]
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.Parent == foreign.ID {
				local.R.Parent = foreign
				break
			}
		}
	}

	return nil
}

// LoadChild allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (authorizeL) LoadChild(e boil.Executor, singular bool, maybeAuthorize interface{}) error {
	var slice []*Authorize
	var object *Authorize

	count := 1
	if singular {
		object = maybeAuthorize.(*Authorize)
	} else {
		slice = *maybeAuthorize.(*AuthorizeSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &authorizeR{}
		}
		args[0] = object.Child
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &authorizeR{}
			}
			args[i] = obj.Child
		}
	}

	query := fmt.Sprintf(
		"select * from \"party\" where \"id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Party")
	}
	defer results.Close()

	var resultSlice []*Party
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Party")
	}

	if len(authorizeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		object.R.Child = resultSlice[0]
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.Child == foreign.ID {
				local.R.Child = foreign
				break
			}
		}
	}

	return nil
}

// SetParentG of the authorize to the related item.
// Sets o.R.Parent to related.
// Adds o to related.R.ParentAuthorizes.
// Uses the global database handle.
func (o *Authorize) SetParentG(insert bool, related *Party) error {
	return o.SetParent(boil.GetDB(), insert, related)
}

// SetParentP of the authorize to the related item.
// Sets o.R.Parent to related.
// Adds o to related.R.ParentAuthorizes.
// Panics on error.
func (o *Authorize) SetParentP(exec boil.Executor, insert bool, related *Party) {
	if err := o.SetParent(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetParentGP of the authorize to the related item.
// Sets o.R.Parent to related.
// Adds o to related.R.ParentAuthorizes.
// Uses the global database handle and panics on error.
func (o *Authorize) SetParentGP(insert bool, related *Party) {
	if err := o.SetParent(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetParent of the authorize to the related item.
// Sets o.R.Parent to related.
// Adds o to related.R.ParentAuthorizes.
func (o *Authorize) SetParent(exec boil.Executor, insert bool, related *Party) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"authorize\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"parent"}),
		strmangle.WhereClause("\"", "\"", 2, authorizePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.Child, o.Parent}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.Parent = related.ID

	if o.R == nil {
		o.R = &authorizeR{
			Parent: related,
		}
	} else {
		o.R.Parent = related
	}

	if related.R == nil {
		related.R = &partyR{
			ParentAuthorizes: AuthorizeSlice{o},
		}
	} else {
		related.R.ParentAuthorizes = append(related.R.ParentAuthorizes, o)
	}

	return nil
}

// SetChildG of the authorize to the related item.
// Sets o.R.Child to related.
// Adds o to related.R.ChildAuthorizes.
// Uses the global database handle.
func (o *Authorize) SetChildG(insert bool, related *Party) error {
	return o.SetChild(boil.GetDB(), insert, related)
}

// SetChildP of the authorize to the related item.
// Sets o.R.Child to related.
// Adds o to related.R.ChildAuthorizes.
// Panics on error.
func (o *Authorize) SetChildP(exec boil.Executor, insert bool, related *Party) {
	if err := o.SetChild(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetChildGP of the authorize to the related item.
// Sets o.R.Child to related.
// Adds o to related.R.ChildAuthorizes.
// Uses the global database handle and panics on error.
func (o *Authorize) SetChildGP(insert bool, related *Party) {
	if err := o.SetChild(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetChild of the authorize to the related item.
// Sets o.R.Child to related.
// Adds o to related.R.ChildAuthorizes.
func (o *Authorize) SetChild(exec boil.Executor, insert bool, related *Party) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"authorize\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"child"}),
		strmangle.WhereClause("\"", "\"", 2, authorizePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.Child, o.Parent}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.Child = related.ID

	if o.R == nil {
		o.R = &authorizeR{
			Child: related,
		}
	} else {
		o.R.Child = related
	}

	if related.R == nil {
		related.R = &partyR{
			ChildAuthorizes: AuthorizeSlice{o},
		}
	} else {
		related.R.ChildAuthorizes = append(related.R.ChildAuthorizes, o)
	}

	return nil
}

// AuthorizesG retrieves all records.
func AuthorizesG(mods ...qm.QueryMod) authorizeQuery {
	return Authorizes(boil.GetDB(), mods...)
}

// Authorizes retrieves all the records using an executor.
func Authorizes(exec boil.Executor, mods ...qm.QueryMod) authorizeQuery {
	mods = append(mods, qm.From("\"authorize\""))
	return authorizeQuery{NewQuery(exec, mods...)}
}

// FindAuthorizeG retrieves a single record by ID.
func FindAuthorizeG(child int, parent int, selectCols ...string) (*Authorize, error) {
	return FindAuthorize(boil.GetDB(), child, parent, selectCols...)
}

// FindAuthorizeGP retrieves a single record by ID, and panics on error.
func FindAuthorizeGP(child int, parent int, selectCols ...string) *Authorize {
	retobj, err := FindAuthorize(boil.GetDB(), child, parent, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindAuthorize retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAuthorize(exec boil.Executor, child int, parent int, selectCols ...string) (*Authorize, error) {
	authorizeObj := &Authorize{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"authorize\" where \"child\"=$1 AND \"parent\"=$2", sel,
	)

	q := queries.Raw(exec, query, child, parent)

	err := q.Bind(authorizeObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from authorize")
	}

	return authorizeObj, nil
}

// FindAuthorizeP retrieves a single record by ID with an executor, and panics on error.
func FindAuthorizeP(exec boil.Executor, child int, parent int, selectCols ...string) *Authorize {
	retobj, err := FindAuthorize(exec, child, parent, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Authorize) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Authorize) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Authorize) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Authorize) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no authorize provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(authorizeColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	authorizeInsertCacheMut.RLock()
	cache, cached := authorizeInsertCache[key]
	authorizeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			authorizeColumns,
			authorizeColumnsWithDefault,
			authorizeColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(authorizeType, authorizeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(authorizeType, authorizeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"authorize\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"authorize\" DEFAULT VALUES"
		}

		if len(cache.retMapping) != 0 {
			cache.query += fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into authorize")
	}

	if !cached {
		authorizeInsertCacheMut.Lock()
		authorizeInsertCache[key] = cache
		authorizeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Authorize record. See Update for
// whitelist behavior description.
func (o *Authorize) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Authorize record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Authorize) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Authorize, and panics on error.
// See Update for whitelist behavior description.
func (o *Authorize) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Authorize.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Authorize) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	authorizeUpdateCacheMut.RLock()
	cache, cached := authorizeUpdateCache[key]
	authorizeUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(authorizeColumns, authorizePrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update authorize, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"authorize\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, authorizePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(authorizeType, authorizeMapping, append(wl, authorizePrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update authorize row")
	}

	if !cached {
		authorizeUpdateCacheMut.Lock()
		authorizeUpdateCache[key] = cache
		authorizeUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q authorizeQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q authorizeQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for authorize")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o AuthorizeSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o AuthorizeSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o AuthorizeSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AuthorizeSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authorizePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	query := fmt.Sprintf(
		"UPDATE \"authorize\" SET %s WHERE (\"child\",\"parent\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(authorizePrimaryKeyColumns), len(colNames)+1, len(authorizePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(query, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in authorize slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Authorize) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Authorize) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Authorize) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Authorize) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no authorize provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(authorizeColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs postgres problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range updateColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range whitelist {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	authorizeUpsertCacheMut.RLock()
	cache, cached := authorizeUpsertCache[key]
	authorizeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			authorizeColumns,
			authorizeColumnsWithDefault,
			authorizeColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			authorizeColumns,
			authorizePrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert authorize, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(authorizePrimaryKeyColumns))
			copy(conflict, authorizePrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"authorize\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(authorizeType, authorizeMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(authorizeType, authorizeMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert authorize")
	}

	if !cached {
		authorizeUpsertCacheMut.Lock()
		authorizeUpsertCache[key] = cache
		authorizeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Authorize record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Authorize) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Authorize record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Authorize) DeleteG() error {
	if o == nil {
		return errors.New("models: no Authorize provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Authorize record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Authorize) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Authorize record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Authorize) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Authorize provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), authorizePrimaryKeyMapping)
	query := "DELETE FROM \"authorize\" WHERE \"child\"=$1 AND \"parent\"=$2"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(query, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from authorize")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q authorizeQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q authorizeQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no authorizeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from authorize")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o AuthorizeSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o AuthorizeSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Authorize slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o AuthorizeSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AuthorizeSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Authorize slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(authorizeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authorizePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	query := fmt.Sprintf(
		"DELETE FROM \"authorize\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, authorizePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(authorizePrimaryKeyColumns), 1, len(authorizePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(query, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from authorize slice")
	}

	if len(authorizeAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Authorize) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Authorize) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Authorize) ReloadG() error {
	if o == nil {
		return errors.New("models: no Authorize provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Authorize) Reload(exec boil.Executor) error {
	ret, err := FindAuthorize(exec, o.Child, o.Parent)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *AuthorizeSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *AuthorizeSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AuthorizeSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty AuthorizeSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AuthorizeSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	authorizes := AuthorizeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authorizePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	query := fmt.Sprintf(
		"SELECT \"authorize\".* FROM \"authorize\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, authorizePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(authorizePrimaryKeyColumns), 1, len(authorizePrimaryKeyColumns)),
	)

	q := queries.Raw(exec, query, args...)

	err := q.Bind(&authorizes)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AuthorizeSlice")
	}

	*o = authorizes

	return nil
}

// AuthorizeExists checks if the Authorize row exists.
func AuthorizeExists(exec boil.Executor, child int, parent int) (bool, error) {
	var exists bool

	query := "select exists(select 1 from \"authorize\" where \"child\"=$1 AND \"parent\"=$2 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, child, parent)
	}

	row := exec.QueryRow(query, child, parent)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if authorize exists")
	}

	return exists, nil
}

// AuthorizeExistsG checks if the Authorize row exists.
func AuthorizeExistsG(child int, parent int) (bool, error) {
	return AuthorizeExists(boil.GetDB(), child, parent)
}

// AuthorizeExistsGP checks if the Authorize row exists. Panics on error.
func AuthorizeExistsGP(child int, parent int) bool {
	e, err := AuthorizeExists(boil.GetDB(), child, parent)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// AuthorizeExistsP checks if the Authorize row exists. Panics on error.
func AuthorizeExistsP(exec boil.Executor, child int, parent int) bool {
	e, err := AuthorizeExists(exec, child, parent)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
