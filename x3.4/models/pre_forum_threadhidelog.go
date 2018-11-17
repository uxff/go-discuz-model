package models

type PreForumThreadhidelog struct {
	Tid int `xorm:"not null pk default 0 unique(uid) MEDIUMINT(8)"`
	Uid int `xorm:"not null pk default 0 unique(uid) MEDIUMINT(8)"`
}
