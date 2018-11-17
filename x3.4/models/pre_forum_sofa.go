package models

type PreForumSofa struct {
	Tid int `xorm:"not null pk default 0 index(ftid) MEDIUMINT(8)"`
	Fid int `xorm:"not null default 0 index(ftid) MEDIUMINT(8)"`
}

func (t *PreForumSofa) TableName() string {
	return "pre_forum_sofa"
}
