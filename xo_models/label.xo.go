// Package xo_models contains the types for schema 'dev_themiseum'.
package xo_models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"
)

// Label represents a row from 'dev_themiseum.labels'.
type Label struct {
	ID         int          `json:"id"`          // id
	Name       string       `json:"name"`        // name
	UUID       string       `json:"uuid"`        // uuid
	IsActive   sql.NullBool `json:"is_active"`   // is_active
	CreateTime time.Time    `json:"create_time"` // create_time
	UpdateTime time.Time    `json:"update_time"` // update_time

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Label exists in the database.
func (l *Label) Exists() bool {
	return l._exists
}

// Deleted provides information if the Label has been deleted from the database.
func (l *Label) Deleted() bool {
	return l._deleted
}

// Insert inserts the Label to the database.
func (l *Label) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if l._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO dev_themiseum.labels (` +
		`name, uuid, is_active, create_time, update_time` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, l.Name, l.UUID, l.IsActive, l.CreateTime, l.UpdateTime)
	res, err := db.Exec(sqlstr, l.Name, l.UUID, l.IsActive, l.CreateTime, l.UpdateTime)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	l.ID = int(id)
	l._exists = true

	return nil
}

// Update updates the Label in the database.
func (l *Label) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !l._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if l._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE dev_themiseum.labels SET ` +
		`name = ?, uuid = ?, is_active = ?, create_time = ?, update_time = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, l.Name, l.UUID, l.IsActive, l.CreateTime, l.UpdateTime, l.ID)
	_, err = db.Exec(sqlstr, l.Name, l.UUID, l.IsActive, l.CreateTime, l.UpdateTime, l.ID)
	return err
}

// Save saves the Label to the database.
func (l *Label) Save(db XODB) error {
	if l.Exists() {
		return l.Update(db)
	}

	return l.Insert(db)
}

// Delete deletes the Label from the database.
func (l *Label) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !l._exists {
		return nil
	}

	// if deleted, bail
	if l._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM dev_themiseum.labels WHERE id = ?`

	// run query
	XOLog(sqlstr, l.ID)
	_, err = db.Exec(sqlstr, l.ID)
	if err != nil {
		return err
	}

	// set deleted
	l._deleted = true

	return nil
}

// LabelByID retrieves a row from 'dev_themiseum.labels' as a Label.
//
// Generated from index 'labels_id_pkey'.
func LabelByID(db XODB, id int) (*Label, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, uuid, is_active, create_time, update_time ` +
		`FROM dev_themiseum.labels ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	l := Label{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&l.ID, &l.Name, &l.UUID, &l.IsActive, &l.CreateTime, &l.UpdateTime)
	if err != nil {
		return nil, err
	}

	return &l, nil
}

// LabelByUUIDName retrieves a row from 'dev_themiseum.labels' as a Label.
//
// Generated from index 'uuid'.
func LabelByUUIDName(db XODB, uuid string, name string) (*Label, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, uuid, is_active, create_time, update_time ` +
		`FROM dev_themiseum.labels ` +
		`WHERE uuid = ? AND name = ?`

	// run query
	XOLog(sqlstr, uuid, name)
	l := Label{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, uuid, name).Scan(&l.ID, &l.Name, &l.UUID, &l.IsActive, &l.CreateTime, &l.UpdateTime)
	if err != nil {
		return nil, err
	}

	return &l, nil
}

// LabelsByUUIDName retrieves a row from 'dev_themiseum.labels' as a Label.
//
// Generated from index 'uuid_2'.
func LabelsByUUIDName(db XODB, uuid string, name string) ([]*Label, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, uuid, is_active, create_time, update_time ` +
		`FROM dev_themiseum.labels ` +
		`WHERE uuid = ? AND name = ?`

	// run query
	XOLog(sqlstr, uuid, name)
	q, err := db.Query(sqlstr, uuid, name)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Label{}
	for q.Next() {
		l := Label{
			_exists: true,
		}

		// scan
		err = q.Scan(&l.ID, &l.Name, &l.UUID, &l.IsActive, &l.CreateTime, &l.UpdateTime)
		if err != nil {
			return nil, err
		}

		res = append(res, &l)
	}

	return res, nil
}
