package models

type PreForumPostlog struct {
	Pid      int    `xorm:"not null pk default 0 INT(10)"`
	Tid      int    `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Fid      int    `xorm:"not null default 0 index SMALLINT(6)"`
	Uid      int    `xorm:"not null default 0 index MEDIUMINT(8)"`
	Action   string `xorm:"not null default '''' CHAR(10)"`
	Dateline int    `xorm:"not null default 0 index INT(10)"`
}

func (t *PreForumPostlog) TableName() string {
	return "pre_forum_postlog"
}
