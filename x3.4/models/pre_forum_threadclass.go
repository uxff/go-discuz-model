package models

type PreForumThreadclass struct {
	Typeid       int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Fid          int    `xorm:"not null index(fid) MEDIUMINT(8)"`
	Name         string `xorm:"not null VARCHAR(255)"`
	Displayorder int    `xorm:"not null index(fid) MEDIUMINT(9)"`
	Icon         string `xorm:"not null VARCHAR(255)"`
	Moderators   int    `xorm:"not null default 0 TINYINT(1)"`
}

func (t *PreForumThreadclass) TableName() string {
	return "pre_forum_threadclass"
}
