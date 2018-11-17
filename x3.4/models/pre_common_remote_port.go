package models

type PreCommonRemotePort struct {
	Id     int    `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Idtype string `xorm:"not null pk default '''' CHAR(15)"`
	Useip  string `xorm:"not null default '''' CHAR(15)"`
	Port   int    `xorm:"not null default 0 SMALLINT(6)"`
}
