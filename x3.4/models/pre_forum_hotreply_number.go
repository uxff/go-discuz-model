package models

type PreForumHotreplyNumber struct {
	Pid     int `xorm:"not null pk default 0 INT(10)"`
	Tid     int `xorm:"not null default 0 index(tid) MEDIUMINT(8)"`
	Support int `xorm:"not null default 0 SMALLINT(6)"`
	Against int `xorm:"not null default 0 SMALLINT(6)"`
	Total   int `xorm:"not null default 0 index(tid) MEDIUMINT(8)"`
}
