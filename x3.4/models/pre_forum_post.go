package models

type PreForumPost struct {
	Pid         int    `xorm:"not null unique INT(10)"`
	Fid         int    `xorm:"not null default 0 index MEDIUMINT(8)"`
	Tid         int    `xorm:"not null pk default 0 index(displayorder) index(first) MEDIUMINT(8)"`
	First       int    `xorm:"not null default 0 index(first) TINYINT(1)"`
	Author      string `xorm:"not null default '''' VARCHAR(15)"`
	Authorid    int    `xorm:"not null default 0 index(authorid) MEDIUMINT(8)"`
	Subject     string `xorm:"not null default '''' VARCHAR(80)"`
	Dateline    int    `xorm:"not null default 0 index index(displayorder) INT(10)"`
	Message     string `xorm:"not null MEDIUMTEXT"`
	Useip       string `xorm:"not null default '''' VARCHAR(15)"`
	Port        int    `xorm:"not null default 0 SMALLINT(6)"`
	Invisible   int    `xorm:"not null default 0 index(authorid) index(displayorder) index TINYINT(1)"`
	Anonymous   int    `xorm:"not null default 0 TINYINT(1)"`
	Usesig      int    `xorm:"not null default 0 TINYINT(1)"`
	Htmlon      int    `xorm:"not null default 0 TINYINT(1)"`
	Bbcodeoff   int    `xorm:"not null default 0 TINYINT(1)"`
	Smileyoff   int    `xorm:"not null default 0 TINYINT(1)"`
	Parseurloff int    `xorm:"not null default 0 TINYINT(1)"`
	Attachment  int    `xorm:"not null default 0 TINYINT(1)"`
	Rate        int    `xorm:"not null default 0 SMALLINT(6)"`
	Ratetimes   int    `xorm:"not null default 0 TINYINT(3)"`
	Status      int    `xorm:"not null default 0 INT(10)"`
	Tags        string `xorm:"not null default ''0'' VARCHAR(255)"`
	Comment     int    `xorm:"not null default 0 TINYINT(1)"`
	Replycredit int    `xorm:"not null default 0 INT(10)"`
	Position    int    `xorm:"not null pk autoincr INT(8)"`
}

func (t *PreForumPost) TableName() string {
	return "pre_forum_post"
}
