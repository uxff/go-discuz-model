package models

type PreConnectPostfeedlog struct {
	Flid          int `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Pid           int `xorm:"not null default 0 unique INT(10)"`
	Uid           int `xorm:"not null default 0 MEDIUMINT(8)"`
	Publishtimes  int `xorm:"not null default 0 MEDIUMINT(8)"`
	Lastpublished int `xorm:"not null default 0 INT(10)"`
	Dateline      int `xorm:"not null default 0 INT(10)"`
	Status        int `xorm:"not null default 1 TINYINT(1)"`
}
