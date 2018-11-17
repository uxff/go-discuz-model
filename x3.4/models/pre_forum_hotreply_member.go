package models

type PreForumHotreplyMember struct {
	Tid      int `xorm:"not null default 0 MEDIUMINT(8)"`
	Pid      int `xorm:"not null pk default 0 INT(10)"`
	Uid      int `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Attitude int `xorm:"not null default 0 TINYINT(1)"`
}

func (t *PreForumHotreplyMember) TableName() string {
	return "pre_forum_hotreply_member"
}
