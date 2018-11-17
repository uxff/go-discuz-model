package models

type PreHomeClick struct {
	Clickid      int    `xorm:"not null pk autoincr SMALLINT(6)"`
	Name         string `xorm:"not null default '''' CHAR(50)"`
	Icon         string `xorm:"not null default '''' CHAR(100)"`
	Idtype       string `xorm:"not null default '''' index(idtype) CHAR(15)"`
	Available    int    `xorm:"not null default 0 TINYINT(1)"`
	Displayorder int    `xorm:"not null default 0 index(idtype) TINYINT(6)"`
}
