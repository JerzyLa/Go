package xorm_models

import (
	"time"
)

type Labels struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	Name       string    `xorm:"not null unique(uuid) index(uuid_2) VARCHAR(40)"`
	Uuid       string    `xorm:"not null unique(uuid) index(uuid_2) VARCHAR(36)"`
	IsActive   int       `xorm:"default 1 TINYINT(1)"`
	CreateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
