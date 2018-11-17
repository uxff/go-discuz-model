package models

type PreForumRatelog struct {
	Pid        int    `xorm:"not null default 0 index(pid) INT(10)"`
	Uid        int    `xorm:"not null default 0 index MEDIUMINT(8)"`
	Username   string `xorm:"not null default '''' CHAR(15)"`
	Extcredits int    `xorm:"not null default 0 TINYINT(1)"`
	Dateline   int    `xorm:"not null default 0 index index(pid) INT(10)"`
	Score      int    `xorm:"not null default 0 SMALLINT(6)"`
	Reason     string `xorm:"not null default '''' CHAR(40)"`
}

func (t *PreForumRatelog) TableName() string {
	return "pre_forum_ratelog"
}
