package models

type PreForumNewthread struct {
	Tid      int `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Fid      int `xorm:"not null default 0 index MEDIUMINT(8)"`
	Dateline int `xorm:"not null default 0 index INT(10)"`
}
