package models

type PreForumAnnouncement struct {
	Id           int    `xorm:"not null pk autoincr SMALLINT(6)"`
	Author       string `xorm:"not null default '''' VARCHAR(15)"`
	Subject      string `xorm:"not null default '''' VARCHAR(255)"`
	Type         int    `xorm:"not null default 0 TINYINT(1)"`
	Displayorder int    `xorm:"not null default 0 TINYINT(3)"`
	Starttime    int    `xorm:"not null default 0 index(timespan) INT(10)"`
	Endtime      int    `xorm:"not null default 0 index(timespan) INT(10)"`
	Message      string `xorm:"not null TEXT"`
	Groups       string `xorm:"not null TEXT"`
}
