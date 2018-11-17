package models

type PreForumTypevar struct {
	Sortid       int `xorm:"not null pk default 0 unique(optionid) index SMALLINT(6)"`
	Optionid     int `xorm:"not null pk default 0 unique(optionid) SMALLINT(6)"`
	Available    int `xorm:"not null default 0 TINYINT(1)"`
	Required     int `xorm:"not null default 0 TINYINT(1)"`
	Unchangeable int `xorm:"not null default 0 TINYINT(1)"`
	Search       int `xorm:"not null default 0 TINYINT(1)"`
	Displayorder int `xorm:"not null default 0 TINYINT(3)"`
	Subjectshow  int `xorm:"not null default 0 TINYINT(1)"`
}
