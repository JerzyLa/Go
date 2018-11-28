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

type FilesLabel struct {
	ID         int       `gorm:"column:id;primary_key" json:"id"`
	FileID     int       `gorm:"column:file_id" json:"file_id"`
	LabelID    null.Int  `gorm:"column:label_id" json:"label_id"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (f *FilesLabel) TableName() string {
	return "files_labels"
}
