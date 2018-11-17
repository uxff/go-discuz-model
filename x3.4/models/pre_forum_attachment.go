package models

type PreForumAttachment struct {
	Aid       int `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Tid       int `xorm:"not null default 0 index MEDIUMINT(8)"`
	Pid       int `xorm:"not null default 0 index INT(10)"`
	Uid       int `xorm:"not null default 0 index MEDIUMINT(8)"`
	Tableid   int `xorm:"not null default 0 TINYINT(1)"`
	Downloads int `xorm:"not null default 0 MEDIUMINT(8)"`
}

func (t *PreForumAttachment) TableName() string {
	return "pre_forum_attachment"
}
