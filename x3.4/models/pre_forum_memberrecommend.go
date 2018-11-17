package models

type PreForumMemberrecommend struct {
	Tid          int `xorm:"not null index MEDIUMINT(8)"`
	Recommenduid int `xorm:"not null index MEDIUMINT(8)"`
	Dateline     int `xorm:"not null INT(10)"`
}

func (t *PreForumMemberrecommend) TableName() string {
	return "pre_forum_memberrecommend"
}
