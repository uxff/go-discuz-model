package models

type PreHomeFriend struct {
	Uid       int    `xorm:"not null pk default 0 index(uid) MEDIUMINT(8)"`
	Fuid      int    `xorm:"not null pk default 0 index MEDIUMINT(8)"`
	Fusername string `xorm:"not null default '''' VARCHAR(15)"`
	Gid       int    `xorm:"not null default 0 SMALLINT(6)"`
	Num       int    `xorm:"not null default 0 index(uid) MEDIUMINT(8)"`
	Dateline  int    `xorm:"not null default 0 index(uid) INT(10)"`
	Note      string `xorm:"not null default '''' VARCHAR(255)"`
}
