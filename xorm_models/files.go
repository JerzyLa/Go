package xorm_models

import (
	"time"
)

type Files struct {
	Id              int       `xorm:"not null pk autoincr INT(11)"`
	FileHash        string    `xorm:"not null unique(file_hash) index(file_hash_2) VARCHAR(64)"`
	Signature       string    `xorm:"not null unique(file_hash) index(file_hash_2) VARCHAR(130)"`
	TransactionHash string    `xorm:"VARCHAR(64)"`
	Name            string    `xorm:"VARCHAR(255)"`
	Description     string    `xorm:"VARCHAR(255)"`
	BlockId         int64     `xorm:"BIGINT(20)"`
	Address         string    `xorm:"VARCHAR(40)"`
	Confirmed       int       `xorm:"default 0 TINYINT(1)"`
	Uuid            string    `xorm:"VARCHAR(36)"`
	Email           string    `xorm:"VARCHAR(255)"`
	FileIndex       int       `xorm:"not null INT(11)"`
	CreateTime      time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdateTime      time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
