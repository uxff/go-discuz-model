package models

type PrePortalRsscache struct {
	Lastupdate  int    `xorm:"not null default 0 INT(10)"`
	Catid       int    `xorm:"not null default 0 index(catid) MEDIUMINT(8)"`
	Aid         int    `xorm:"not null pk default 0 unique MEDIUMINT(8)"`
	Dateline    int    `xorm:"not null default 0 index(catid) INT(10)"`
	Catname     string `xorm:"not null default '''' CHAR(50)"`
	Author      string `xorm:"not null default '''' CHAR(15)"`
	Subject     string `xorm:"not null default '''' CHAR(80)"`
	Description string `xorm:"not null default '''' CHAR(255)"`
}
