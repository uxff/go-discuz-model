package models

type PreForumThreadaddviews struct {
	Tid      int `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Addviews int `xorm:"not null default 0 INT(10)"`
}
