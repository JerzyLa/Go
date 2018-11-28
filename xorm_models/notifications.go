package xorm_models

import (
	"time"
)

type Notifications struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	Email      string    `xorm:"not null unique(file_id) index(file_id_2) VARCHAR(100)"`
	FileId     int       `xorm:"not null unique(file_id) index(file_id_2) INT(11)"`
	WasSent    int       `xorm:"TINYINT(1)"`
	IsActive   int       `xorm:"TINYINT(1)"`
	CreateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
