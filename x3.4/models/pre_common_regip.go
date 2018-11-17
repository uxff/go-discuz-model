package models

type PreCommonRegip struct {
	Ip       string `xorm:"not null default '''' index CHAR(15)"`
	Dateline int    `xorm:"not null default 0 INT(10)"`
	Count    int    `xorm:"not null default 0 SMALLINT(6)"`
}
