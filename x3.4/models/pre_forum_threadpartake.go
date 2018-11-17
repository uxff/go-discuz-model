package models

type PreForumThreadpartake struct {
	Tid      int `xorm:"not null default 0 index(tid) MEDIUMINT(8)"`
	Uid      int `xorm:"not null default 0 index(tid) MEDIUMINT(8)"`
	Dateline int `xorm:"not null default 0 INT(10)"`
}

func (t *PreForumThreadpartake) TableName() string {
	return "pre_forum_threadpartake"
}
