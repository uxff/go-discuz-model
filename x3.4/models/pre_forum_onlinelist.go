package models

type PreForumOnlinelist struct {
	Groupid      int    `xorm:"not null default 0 SMALLINT(6)"`
	Displayorder int    `xorm:"not null default 0 TINYINT(3)"`
	Title        string `xorm:"not null default '''' VARCHAR(30)"`
	Url          string `xorm:"not null default '''' VARCHAR(30)"`
}

func (t *PreForumOnlinelist) TableName() string {
	return "pre_forum_onlinelist"
}
