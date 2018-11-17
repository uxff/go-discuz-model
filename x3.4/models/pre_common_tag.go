package models

type PreCommonTag struct {
	Tagid   int    `xorm:"not null pk autoincr index(status) MEDIUMINT(8)"`
	Tagname string `xorm:"not null default '''' index CHAR(20)"`
	Status  int    `xorm:"not null default 0 index(status) TINYINT(1)"`
}
