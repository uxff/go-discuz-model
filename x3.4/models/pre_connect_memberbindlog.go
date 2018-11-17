package models

type PreConnectMemberbindlog struct {
	Mblid    int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Uid      int    `xorm:"not null default 0 index MEDIUMINT(8)"`
	Uin      string `xorm:"not null index CHAR(40)"`
	Type     int    `xorm:"not null default 0 TINYINT(1)"`
	Dateline int    `xorm:"not null default 0 index INT(10)"`
}
