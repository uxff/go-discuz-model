package models

type PreHomePokearchive struct {
	Pid      int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Pokeuid  int    `xorm:"not null default 0 index INT(10)"`
	Uid      int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Fromuid  int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Note     string `xorm:"not null default '''' VARCHAR(255)"`
	Dateline int    `xorm:"not null default 0 INT(10)"`
	Iconid   int    `xorm:"not null default 0 SMALLINT(6)"`
}
