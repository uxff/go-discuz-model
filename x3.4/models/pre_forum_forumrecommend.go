package models

type PreForumForumrecommend struct {
	Fid          int    `xorm:"not null default 0 index(displayorder) MEDIUMINT(8)"`
	Tid          int    `xorm:"not null pk MEDIUMINT(8)"`
	Typeid       int    `xorm:"not null SMALLINT(6)"`
	Displayorder int    `xorm:"not null index(displayorder) TINYINT(1)"`
	Subject      string `xorm:"not null CHAR(80)"`
	Author       string `xorm:"not null CHAR(15)"`
	Authorid     int    `xorm:"not null MEDIUMINT(8)"`
	Moderatorid  int    `xorm:"not null MEDIUMINT(8)"`
	Expiration   int    `xorm:"not null INT(10)"`
	Position     int    `xorm:"not null default 0 index TINYINT(1)"`
	Highlight    int    `xorm:"not null default 0 TINYINT(1)"`
	Aid          int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Filename     string `xorm:"not null default '''' CHAR(100)"`
}
