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

type Notification struct {
	ID         int       `gorm:"column:id;primary_key" json:"id"`
	Email      string    `gorm:"column:email" json:"email"`
	FileID     int       `gorm:"column:file_id" json:"file_id"`
	WasSent    null.Int  `gorm:"column:was_sent" json:"was_sent"`
	IsActive   null.Int  `gorm:"column:is_active" json:"is_active"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (n *Notification) TableName() string {
	return "notifications"
}
