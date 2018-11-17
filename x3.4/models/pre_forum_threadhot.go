package models

type PreForumThreadhot struct {
	Cid int `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Fid int `xorm:"not null default 0 index MEDIUMINT(8)"`
	Tid int `xorm:"not null pk default 0 MEDIUMINT(8)"`
}
