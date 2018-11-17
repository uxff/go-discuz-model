package models

type PreForumPostTableid struct {
	Pid int `xorm:"not null pk autoincr INT(10)"`
}

func (t *PreForumPostTableid) TableName() string {
	return "pre_forum_post_tableid"
}
