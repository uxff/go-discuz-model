package models

import (
	"time"
)

type PreForumModwork struct {
	Uid       int       `xorm:"not null default 0 index(uid) MEDIUMINT(8)"`
	Modaction string    `xorm:"not null default '''' CHAR(3)"`
	Dateline  time.Time `xorm:"not null default ''2006-01-01'' index(uid) DATE"`
	Count     int       `xorm:"not null default 0 SMALLINT(6)"`
	Posts     int       `xorm:"not null default 0 SMALLINT(6)"`
}

func (t *PreForumModwork) TableName() string {
	return "pre_forum_modwork"
}
