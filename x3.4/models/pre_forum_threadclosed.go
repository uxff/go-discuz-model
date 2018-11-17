package models

type PreForumThreadclosed struct {
	Tid      int `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Redirect int `xorm:"not null default 0 MEDIUMINT(8)"`
}

func (t *PreForumThreadclosed) TableName() string {
	return "pre_forum_threadclosed"
}
