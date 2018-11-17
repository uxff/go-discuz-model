package models

type PreForumImagetype struct {
	Typeid       int    `xorm:"not null pk autoincr SMALLINT(6)"`
	Available    int    `xorm:"not null default 0 TINYINT(1)"`
	Name         string `xorm:"not null CHAR(20)"`
	Type         string `xorm:"not null default ''smiley'' ENUM('avatar','icon','smiley')"`
	Displayorder int    `xorm:"not null default 0 TINYINT(3)"`
	Directory    string `xorm:"not null CHAR(100)"`
}
