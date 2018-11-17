package models

type PreForumGroupinvite struct {
	Fid       int `xorm:"not null pk default 0 unique(ids) MEDIUMINT(8)"`
	Uid       int `xorm:"not null default 0 MEDIUMINT(8)"`
	Inviteuid int `xorm:"not null pk default 0 unique(ids) MEDIUMINT(8)"`
	Dateline  int `xorm:"not null default 0 index INT(10)"`
}
