package models

type PreForumPostcomment struct {
	Id       int    `xorm:"not null pk autoincr INT(10)"`
	Tid      int    `xorm:"not null default 0 index MEDIUMINT(8)"`
	Pid      int    `xorm:"not null default 0 index(pid) INT(10)"`
	Author   string `xorm:"not null default '''' VARCHAR(15)"`
	Authorid int    `xorm:"not null default 0 index MEDIUMINT(8)"`
	Dateline int    `xorm:"not null default 0 index(pid) INT(10)"`
	Comment  string `xorm:"not null default '''' VARCHAR(255)"`
	Score    int    `xorm:"not null default 0 index TINYINT(1)"`
	Useip    string `xorm:"not null default '''' VARCHAR(15)"`
	Port     int    `xorm:"not null default 0 SMALLINT(6)"`
	Rpid     int    `xorm:"not null default 0 index INT(10)"`
}
