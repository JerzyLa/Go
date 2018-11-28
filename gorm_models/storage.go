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

type Storage struct {
	ID         int         `gorm:"column:id;primary_key" json:"id"`
	Name       null.String `gorm:"column:name" json:"name"`
	FilePath   null.String `gorm:"column:file_path" json:"file_path"`
	Access     null.String `gorm:"column:access" json:"access"`
	Encrypted  null.Int    `gorm:"column:encrypted" json:"encrypted"`
	CreateTime time.Time   `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time   `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (s *Storage) TableName() string {
	return "storages"
}
