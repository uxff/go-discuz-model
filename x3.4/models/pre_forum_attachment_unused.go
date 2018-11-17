package models

type PreForumAttachmentUnused struct {
	Aid        int    `xorm:"not null pk MEDIUMINT(8)"`
	Uid        int    `xorm:"not null default 0 index MEDIUMINT(8)"`
	Dateline   int    `xorm:"not null default 0 INT(10)"`
	Filename   string `xorm:"not null default '''' VARCHAR(255)"`
	Filesize   int    `xorm:"not null default 0 INT(10)"`
	Attachment string `xorm:"not null default '''' VARCHAR(255)"`
	Remote     int    `xorm:"not null default 0 TINYINT(1)"`
	Isimage    int    `xorm:"not null default 0 TINYINT(1)"`
	Width      int    `xorm:"not null default 0 SMALLINT(6)"`
	Thumb      int    `xorm:"not null default 0 TINYINT(1)"`
}

func (t *PreForumAttachmentUnused) TableName() string {
	return "pre_forum_attachment_unused"
}
