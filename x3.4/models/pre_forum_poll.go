package models

type PreForumPoll struct {
	Tid         int    `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Overt       int    `xorm:"not null default 0 TINYINT(1)"`
	Multiple    int    `xorm:"not null default 0 TINYINT(1)"`
	Visible     int    `xorm:"not null default 0 TINYINT(1)"`
	Maxchoices  int    `xorm:"not null default 0 TINYINT(3)"`
	Isimage     int    `xorm:"not null default 0 TINYINT(1)"`
	Expiration  int    `xorm:"not null default 0 INT(10)"`
	Pollpreview string `xorm:"not null default '''' VARCHAR(255)"`
	Voters      int    `xorm:"not null default 0 MEDIUMINT(8)"`
}
