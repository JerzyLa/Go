package xorm_models

import (
	"time"
)

type FilesLabels struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	FileId     int       `xorm:"not null index INT(11)"`
	LabelId    int       `xorm:"index INT(11)"`
	CreateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
