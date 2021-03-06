package models

type PreForumRsscache struct {
	Lastupdate  int    `xorm:"not null default 0 INT(10)"`
	Fid         int    `xorm:"not null default 0 index(fid) MEDIUMINT(8)"`
	Tid         int    `xorm:"not null pk default 0 unique MEDIUMINT(8)"`
	Dateline    int    `xorm:"not null default 0 index(fid) INT(10)"`
	Forum       string `xorm:"not null default '''' CHAR(50)"`
	Author      string `xorm:"not null default '''' CHAR(15)"`
	Subject     string `xorm:"not null default '''' CHAR(80)"`
	Description string `xorm:"not null default '''' CHAR(255)"`
	Guidetype   string `xorm:"not null default '''' CHAR(10)"`
}

func (t *PreForumRsscache) TableName() string {
	return "pre_forum_rsscache"
}
