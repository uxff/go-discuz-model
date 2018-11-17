package models

type PreForumModerator struct {
	Uid          int `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Fid          int `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Displayorder int `xorm:"not null default 0 TINYINT(3)"`
	Inherited    int `xorm:"not null default 0 TINYINT(1)"`
}
