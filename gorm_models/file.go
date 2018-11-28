package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

type File struct {
	ID              int         `gorm:"column:id;primary_key" json:"id"`
	FileHash        string      `gorm:"column:file_hash" json:"file_hash"`
	Signature       string      `gorm:"column:signature" json:"signature"`
	TransactionHash null.String `gorm:"column:transaction_hash" json:"transaction_hash"`
	Name            null.String `gorm:"column:name" json:"name"`
	Description     null.String `gorm:"column:description" json:"description"`
	BlockID         null.Int    `gorm:"column:block_id" json:"block_id"`
	Address         null.String `gorm:"column:address" json:"address"`
	Confirmed       null.Int    `gorm:"column:confirmed" json:"confirmed"`
	UUID            null.String `gorm:"column:uuid" json:"uuid"`
	Email           null.String `gorm:"column:email" json:"email"`
	FileIndex       int         `gorm:"column:file_index" json:"file_index"`
	CreateTime      time.Time   `gorm:"column:create_time" json:"create_time"`
	UpdateTime      time.Time   `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (f *File) TableName() string {
	return "files"
}
