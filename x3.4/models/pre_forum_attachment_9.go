package models

type PreForumAttachment9 struct {
	Aid         int    `xorm:"not null pk MEDIUMINT(8)"`
	Tid         int    `xorm:"not null default 0 index MEDIUMINT(8)"`
	Pid         int    `xorm:"not null default 0 index INT(10)"`
	Uid         int    `xorm:"not null default 0 index MEDIUMINT(8)"`
	Dateline    int    `xorm:"not null default 0 INT(10)"`
	Filename    string `xorm:"not null default '''' VARCHAR(255)"`
	Filesize    int    `xorm:"not null default 0 INT(10)"`
	Attachment  string `xorm:"not null default '''' VARCHAR(255)"`
	Remote      int    `xorm:"not null default 0 TINYINT(1)"`
	Description string `xorm:"not null VARCHAR(255)"`
	Readperm    int    `xorm:"not null default 0 TINYINT(3)"`
	Price       int    `xorm:"not null default 0 SMALLINT(6)"`
	Isimage     int    `xorm:"not null default 0 TINYINT(1)"`
	Width       int    `xorm:"not null default 0 SMALLINT(6)"`
	Thumb       int    `xorm:"not null default 0 TINYINT(1)"`
	Picid       int    `xorm:"not null default 0 MEDIUMINT(8)"`
}

func (t *PreForumAttachment9) TableName() string {
	return "pre_forum_attachment_9"
}
