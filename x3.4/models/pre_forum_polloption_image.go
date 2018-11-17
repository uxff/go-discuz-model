package models

type PreForumPolloptionImage struct {
	Aid        int    `xorm:"not null pk autoincr INT(10)"`
	Poid       int    `xorm:"not null default 0 index INT(10)"`
	Tid        int    `xorm:"not null default 0 index MEDIUMINT(8)"`
	Pid        int    `xorm:"not null default 0 INT(10)"`
	Uid        int    `xorm:"not null default 0 index MEDIUMINT(8)"`
	Filename   string `xorm:"not null default '''' VARCHAR(255)"`
	Filesize   int    `xorm:"not null default 0 INT(10)"`
	Attachment string `xorm:"not null default '''' VARCHAR(255)"`
	Remote     int    `xorm:"not null default 0 TINYINT(1)"`
	Width      int    `xorm:"not null default 0 SMALLINT(6)"`
	Thumb      int    `xorm:"not null default 0 TINYINT(1)"`
	Dateline   int    `xorm:"not null default 0 INT(10)"`
}
