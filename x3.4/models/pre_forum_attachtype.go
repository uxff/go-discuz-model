package models

type PreForumAttachtype struct {
	Id        int    `xorm:"not null pk autoincr SMALLINT(6)"`
	Fid       int    `xorm:"not null default 0 index MEDIUMINT(8)"`
	Extension string `xorm:"not null default '''' CHAR(12)"`
	Maxsize   int    `xorm:"not null default 0 INT(10)"`
}

func (t *PreForumAttachtype) TableName() string {
	return "pre_forum_attachtype"
}
