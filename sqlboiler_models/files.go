// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// File is an object representing the database table.
type File struct {
	ID              int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	FileHash        string      `boil:"file_hash" json:"file_hash" toml:"file_hash" yaml:"file_hash"`
	Signature       string      `boil:"signature" json:"signature" toml:"signature" yaml:"signature"`
	TransactionHash null.String `boil:"transaction_hash" json:"transaction_hash,omitempty" toml:"transaction_hash" yaml:"transaction_hash,omitempty"`
	Name            null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	Description     null.String `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`
	BlockID         null.Int64  `boil:"block_id" json:"block_id,omitempty" toml:"block_id" yaml:"block_id,omitempty"`
	Address         null.String `boil:"address" json:"address,omitempty" toml:"address" yaml:"address,omitempty"`
	Confirmed       null.Bool   `boil:"confirmed" json:"confirmed,omitempty" toml:"confirmed" yaml:"confirmed,omitempty"`
	UUID            null.String `boil:"uuid" json:"uuid,omitempty" toml:"uuid" yaml:"uuid,omitempty"`
	Email           null.String `boil:"email" json:"email,omitempty" toml:"email" yaml:"email,omitempty"`
	FileIndex       int         `boil:"file_index" json:"file_index" toml:"file_index" yaml:"file_index"`
	CreateTime      time.Time   `boil:"create_time" json:"create_time" toml:"create_time" yaml:"create_time"`
	UpdateTime      time.Time   `boil:"update_time" json:"update_time" toml:"update_time" yaml:"update_time"`

	R *fileR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L fileL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var FileColumns = struct {
	ID              string
	FileHash        string
	Signature       string
	TransactionHash string
	Name            string
	Description     string
	BlockID         string
	Address         string
	Confirmed       string
	UUID            string
	Email           string
	FileIndex       string
	CreateTime      string
	UpdateTime      string
}{
	ID:              "id",
	FileHash:        "file_hash",
	Signature:       "signature",
	TransactionHash: "transaction_hash",
	Name:            "name",
	Description:     "description",
	BlockID:         "block_id",
	Address:         "address",
	Confirmed:       "confirmed",
	UUID:            "uuid",
	Email:           "email",
	FileIndex:       "file_index",
	CreateTime:      "create_time",
	UpdateTime:      "update_time",
}

// FileRels is where relationship names are stored.
var FileRels = struct {
	FilesLabels   string
	FilesStorages string
	Notifications string
}{
	FilesLabels:   "FilesLabels",
	FilesStorages: "FilesStorages",
	Notifications: "Notifications",
}

// fileR is where relationships are stored.
type fileR struct {
	FilesLabels   FilesLabelSlice
	FilesStorages FilesStorageSlice
	Notifications NotificationSlice
}

// NewStruct creates a new relationship struct
func (*fileR) NewStruct() *fileR {
	return &fileR{}
}

// fileL is where Load methods for each relationship are stored.
type fileL struct{}

var (
	fileColumns               = []string{"id", "file_hash", "signature", "transaction_hash", "name", "description", "block_id", "address", "confirmed", "uuid", "email", "file_index", "create_time", "update_time"}
	fileColumnsWithoutDefault = []string{"file_hash", "signature", "transaction_hash", "name", "description", "block_id", "address", "uuid", "email", "file_index"}
	fileColumnsWithDefault    = []string{"id", "confirmed", "create_time", "update_time"}
	filePrimaryKeyColumns     = []string{"id"}
)

type (
	// FileSlice is an alias for a slice of pointers to File.
	// This should generally be used opposed to []File.
	FileSlice []*File
	// FileHook is the signature for custom File hook methods
	FileHook func(boil.Executor, *File) error

	fileQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	fileType                 = reflect.TypeOf(&File{})
	fileMapping              = queries.MakeStructMapping(fileType)
	filePrimaryKeyMapping, _ = queries.BindMapping(fileType, fileMapping, filePrimaryKeyColumns)
	fileInsertCacheMut       sync.RWMutex
	fileInsertCache          = make(map[string]insertCache)
	fileUpdateCacheMut       sync.RWMutex
	fileUpdateCache          = make(map[string]updateCache)
	fileUpsertCacheMut       sync.RWMutex
	fileUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var fileBeforeInsertHooks []FileHook
var fileBeforeUpdateHooks []FileHook
var fileBeforeDeleteHooks []FileHook
var fileBeforeUpsertHooks []FileHook

var fileAfterInsertHooks []FileHook
var fileAfterSelectHooks []FileHook
var fileAfterUpdateHooks []FileHook
var fileAfterDeleteHooks []FileHook
var fileAfterUpsertHooks []FileHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *File) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range fileBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *File) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range fileBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *File) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range fileBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *File) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range fileBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *File) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range fileAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *File) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range fileAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *File) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range fileAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *File) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range fileAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *File) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range fileAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFileHook registers your hook function for all future operations.
func AddFileHook(hookPoint boil.HookPoint, fileHook FileHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		fileBeforeInsertHooks = append(fileBeforeInsertHooks, fileHook)
	case boil.BeforeUpdateHook:
		fileBeforeUpdateHooks = append(fileBeforeUpdateHooks, fileHook)
	case boil.BeforeDeleteHook:
		fileBeforeDeleteHooks = append(fileBeforeDeleteHooks, fileHook)
	case boil.BeforeUpsertHook:
		fileBeforeUpsertHooks = append(fileBeforeUpsertHooks, fileHook)
	case boil.AfterInsertHook:
		fileAfterInsertHooks = append(fileAfterInsertHooks, fileHook)
	case boil.AfterSelectHook:
		fileAfterSelectHooks = append(fileAfterSelectHooks, fileHook)
	case boil.AfterUpdateHook:
		fileAfterUpdateHooks = append(fileAfterUpdateHooks, fileHook)
	case boil.AfterDeleteHook:
		fileAfterDeleteHooks = append(fileAfterDeleteHooks, fileHook)
	case boil.AfterUpsertHook:
		fileAfterUpsertHooks = append(fileAfterUpsertHooks, fileHook)
	}
}

// One returns a single file record from the query.
func (q fileQuery) One(exec boil.Executor) (*File, error) {
	o := &File{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for files")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all File records from the query.
func (q fileQuery) All(exec boil.Executor) (FileSlice, error) {
	var o []*File

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to File slice")
	}

	if len(fileAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all File records in the query.
func (q fileQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count files rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q fileQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if files exists")
	}

	return count > 0, nil
}

// FilesLabels retrieves all the files_label's FilesLabels with an executor.
func (o *File) FilesLabels(mods ...qm.QueryMod) filesLabelQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`files_labels`.`file_id`=?", o.ID),
	)

	query := FilesLabels(queryMods...)
	queries.SetFrom(query.Query, "`files_labels`")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"`files_labels`.*"})
	}

	return query
}

// FilesStorages retrieves all the files_storage's FilesStorages with an executor.
func (o *File) FilesStorages(mods ...qm.QueryMod) filesStorageQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`files_storages`.`file_id`=?", o.ID),
	)

	query := FilesStorages(queryMods...)
	queries.SetFrom(query.Query, "`files_storages`")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"`files_storages`.*"})
	}

	return query
}

// Notifications retrieves all the notification's Notifications with an executor.
func (o *File) Notifications(mods ...qm.QueryMod) notificationQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`notifications`.`file_id`=?", o.ID),
	)

	query := Notifications(queryMods...)
	queries.SetFrom(query.Query, "`notifications`")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"`notifications`.*"})
	}

	return query
}

// LoadFilesLabels allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (fileL) LoadFilesLabels(e boil.Executor, singular bool, maybeFile interface{}, mods queries.Applicator) error {
	var slice []*File
	var object *File

	if singular {
		object = maybeFile.(*File)
	} else {
		slice = *maybeFile.(*[]*File)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &fileR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &fileR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	query := NewQuery(qm.From(`files_labels`), qm.WhereIn(`file_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load files_labels")
	}

	var resultSlice []*FilesLabel
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice files_labels")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on files_labels")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for files_labels")
	}

	if len(filesLabelAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.FilesLabels = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &filesLabelR{}
			}
			foreign.R.File = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.FileID {
				local.R.FilesLabels = append(local.R.FilesLabels, foreign)
				if foreign.R == nil {
					foreign.R = &filesLabelR{}
				}
				foreign.R.File = local
				break
			}
		}
	}

	return nil
}

// LoadFilesStorages allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (fileL) LoadFilesStorages(e boil.Executor, singular bool, maybeFile interface{}, mods queries.Applicator) error {
	var slice []*File
	var object *File

	if singular {
		object = maybeFile.(*File)
	} else {
		slice = *maybeFile.(*[]*File)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &fileR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &fileR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	query := NewQuery(qm.From(`files_storages`), qm.WhereIn(`file_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load files_storages")
	}

	var resultSlice []*FilesStorage
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice files_storages")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on files_storages")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for files_storages")
	}

	if len(filesStorageAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.FilesStorages = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &filesStorageR{}
			}
			foreign.R.File = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.FileID {
				local.R.FilesStorages = append(local.R.FilesStorages, foreign)
				if foreign.R == nil {
					foreign.R = &filesStorageR{}
				}
				foreign.R.File = local
				break
			}
		}
	}

	return nil
}

// LoadNotifications allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (fileL) LoadNotifications(e boil.Executor, singular bool, maybeFile interface{}, mods queries.Applicator) error {
	var slice []*File
	var object *File

	if singular {
		object = maybeFile.(*File)
	} else {
		slice = *maybeFile.(*[]*File)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &fileR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &fileR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	query := NewQuery(qm.From(`notifications`), qm.WhereIn(`file_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load notifications")
	}

	var resultSlice []*Notification
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice notifications")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on notifications")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for notifications")
	}

	if len(notificationAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Notifications = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &notificationR{}
			}
			foreign.R.File = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.FileID {
				local.R.Notifications = append(local.R.Notifications, foreign)
				if foreign.R == nil {
					foreign.R = &notificationR{}
				}
				foreign.R.File = local
				break
			}
		}
	}

	return nil
}

// AddFilesLabels adds the given related objects to the existing relationships
// of the file, optionally inserting them as new records.
// Appends related to o.R.FilesLabels.
// Sets related.R.File appropriately.
func (o *File) AddFilesLabels(exec boil.Executor, insert bool, related ...*FilesLabel) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.FileID = o.ID
			if err = rel.Insert(exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `files_labels` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"file_id"}),
				strmangle.WhereClause("`", "`", 0, filesLabelPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.FileID = o.ID
		}
	}

	if o.R == nil {
		o.R = &fileR{
			FilesLabels: related,
		}
	} else {
		o.R.FilesLabels = append(o.R.FilesLabels, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &filesLabelR{
				File: o,
			}
		} else {
			rel.R.File = o
		}
	}
	return nil
}

// AddFilesStorages adds the given related objects to the existing relationships
// of the file, optionally inserting them as new records.
// Appends related to o.R.FilesStorages.
// Sets related.R.File appropriately.
func (o *File) AddFilesStorages(exec boil.Executor, insert bool, related ...*FilesStorage) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.FileID = o.ID
			if err = rel.Insert(exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `files_storages` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"file_id"}),
				strmangle.WhereClause("`", "`", 0, filesStoragePrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.FileID = o.ID
		}
	}

	if o.R == nil {
		o.R = &fileR{
			FilesStorages: related,
		}
	} else {
		o.R.FilesStorages = append(o.R.FilesStorages, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &filesStorageR{
				File: o,
			}
		} else {
			rel.R.File = o
		}
	}
	return nil
}

// AddNotifications adds the given related objects to the existing relationships
// of the file, optionally inserting them as new records.
// Appends related to o.R.Notifications.
// Sets related.R.File appropriately.
func (o *File) AddNotifications(exec boil.Executor, insert bool, related ...*Notification) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.FileID = o.ID
			if err = rel.Insert(exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `notifications` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"file_id"}),
				strmangle.WhereClause("`", "`", 0, notificationPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.FileID = o.ID
		}
	}

	if o.R == nil {
		o.R = &fileR{
			Notifications: related,
		}
	} else {
		o.R.Notifications = append(o.R.Notifications, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &notificationR{
				File: o,
			}
		} else {
			rel.R.File = o
		}
	}
	return nil
}

// Files retrieves all the records using an executor.
func Files(mods ...qm.QueryMod) fileQuery {
	mods = append(mods, qm.From("`files`"))
	return fileQuery{NewQuery(mods...)}
}

// FindFile retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFile(exec boil.Executor, iD int, selectCols ...string) (*File, error) {
	fileObj := &File{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `files` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, fileObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from files")
	}

	return fileObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *File) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no files provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(fileColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	fileInsertCacheMut.RLock()
	cache, cached := fileInsertCache[key]
	fileInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			fileColumns,
			fileColumnsWithDefault,
			fileColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(fileType, fileMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(fileType, fileMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `files` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `files` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `files` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, filePrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into files")
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

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == fileMapping["ID"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for files")
	}

CacheNoHooks:
	if !cached {
		fileInsertCacheMut.Lock()
		fileInsertCache[key] = cache
		fileInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// Update uses an executor to update the File.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *File) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	fileUpdateCacheMut.RLock()
	cache, cached := fileUpdateCache[key]
	fileUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			fileColumns,
			filePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update files, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `files` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, filePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(fileType, fileMapping, append(wl, filePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.Exec(cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update files row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for files")
	}

	if !cached {
		fileUpdateCacheMut.Lock()
		fileUpdateCache[key] = cache
		fileUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(exec)
}

// UpdateAll updates all rows with the specified column values.
func (q fileQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for files")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for files")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FileSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), filePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `files` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, filePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in file slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all file")
	}
	return rowsAff, nil
}

var mySQLFileUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *File) Upsert(exec boil.Executor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no files provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(fileColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLFileUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	fileUpsertCacheMut.RLock()
	cache, cached := fileUpsertCache[key]
	fileUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			fileColumns,
			fileColumnsWithDefault,
			fileColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			fileColumns,
			filePrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert files, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "files", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `files` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(fileType, fileMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(fileType, fileMapping, ret)
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

	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for files")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == fileMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(fileType, fileMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for files")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRow(cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for files")
	}

CacheNoHooks:
	if !cached {
		fileUpsertCacheMut.Lock()
		fileUpsertCache[key] = cache
		fileUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// Delete deletes a single File record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *File) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no File provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), filePrimaryKeyMapping)
	sql := "DELETE FROM `files` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from files")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for files")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q fileQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no fileQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from files")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for files")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FileSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no File slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(fileBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), filePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `files` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, filePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from file slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for files")
	}

	if len(fileAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *File) Reload(exec boil.Executor) error {
	ret, err := FindFile(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FileSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := FileSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), filePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `files`.* FROM `files` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, filePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in FileSlice")
	}

	*o = slice

	return nil
}

// FileExists checks if the File row exists.
func FileExists(exec boil.Executor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `files` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if files exists")
	}

	return exists, nil
}