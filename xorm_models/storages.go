package xorm_models

import (
	"time"
)

type Storages struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	Name       string    `xorm:"ENUM('s3','swarm')"`
	FilePath   string    `xorm:"VARCHAR(255)"`
	Access     string    `xorm:"default 'public' ENUM('private','public')"`
	Encrypted  int       `xorm:"TINYINT(1)"`
	CreateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
