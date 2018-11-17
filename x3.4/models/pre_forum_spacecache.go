package models

type PreForumSpacecache struct {
	Uid        int    `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Variable   string `xorm:"not null pk VARCHAR(20)"`
	Value      string `xorm:"not null TEXT"`
	Expiration int    `xorm:"not null default 0 INT(10)"`
}

func (t *PreForumSpacecache) TableName() string {
	return "pre_forum_spacecache"
}
