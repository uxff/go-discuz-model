package models

import (
	"time"
)

type PreForumStatlog struct {
	Logdate time.Time `xorm:"not null pk DATE"`
	Fid     int       `xorm:"not null pk MEDIUMINT(8)"`
	Type    int       `xorm:"not null default 0 SMALLINT(5)"`
	Value   int       `xorm:"not null default 0 INT(10)"`
}
