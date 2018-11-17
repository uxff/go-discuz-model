package models

type PreForumCreditslog struct {
	Uid            int    `xorm:"not null default 0 index(uid) MEDIUMINT(8)"`
	Fromto         string `xorm:"not null default '''' CHAR(15)"`
	Sendcredits    int    `xorm:"not null default 0 TINYINT(1)"`
	Receivecredits int    `xorm:"not null default 0 TINYINT(1)"`
	Send           int    `xorm:"not null default 0 INT(10)"`
	Receive        int    `xorm:"not null default 0 INT(10)"`
	Dateline       int    `xorm:"not null default 0 index(uid) INT(10)"`
	Operation      string `xorm:"not null default '''' CHAR(3)"`
}

func (t *PreForumCreditslog) TableName() string {
	return "pre_forum_creditslog"
}
