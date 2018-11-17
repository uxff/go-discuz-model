package models

type PreForumForumThreadtable struct {
	Fid           int `xorm:"not null pk SMALLINT(6)"`
	Threadtableid int `xorm:"not null pk SMALLINT(6)"`
	Threads       int `xorm:"not null default 0 INT(11)"`
	Posts         int `xorm:"not null default 0 INT(11)"`
}

func (t *PreForumForumThreadtable) TableName() string {
	return "pre_forum_forum_threadtable"
}
