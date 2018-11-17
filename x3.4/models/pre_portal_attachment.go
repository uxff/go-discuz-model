package models

type PrePortalAttachment struct {
	Attachid   int    `xorm:"not null pk autoincr index(aid) MEDIUMINT(8)"`
	Uid        int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Dateline   int    `xorm:"not null default 0 INT(10)"`
	Filename   string `xorm:"not null default '''' VARCHAR(255)"`
	Filetype   string `xorm:"not null default '''' VARCHAR(255)"`
	Filesize   int    `xorm:"not null default 0 INT(10)"`
	Attachment string `xorm:"not null default '''' VARCHAR(255)"`
	Isimage    int    `xorm:"not null default 0 TINYINT(1)"`
	Thumb      int    `xorm:"not null default 0 TINYINT(1)"`
	Remote     int    `xorm:"not null default 0 TINYINT(1)"`
	Aid        int    `xorm:"not null default 0 index(aid) MEDIUMINT(8)"`
}
