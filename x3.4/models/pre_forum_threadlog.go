package models

type PreForumThreadlog struct {
	Tid      int    `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Fid      int    `xorm:"not null pk default 0 SMALLINT(6)"`
	Uid      int    `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Otherid  int    `xorm:"not null default 0 SMALLINT(6)"`
	Action   string `xorm:"not null CHAR(10)"`
	Expiry   int    `xorm:"not null default 0 INT(10)"`
	Dateline int    `xorm:"not null default 0 index INT(10)"`
}

func (t *PreForumThreadlog) TableName() string {
	return "pre_forum_threadlog"
}
