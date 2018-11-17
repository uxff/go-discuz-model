package models

type PreConnectFeedlog struct {
	Flid          int `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Tid           int `xorm:"not null default 0 unique MEDIUMINT(8)"`
	Uid           int `xorm:"not null default 0 MEDIUMINT(8)"`
	Publishtimes  int `xorm:"not null default 0 MEDIUMINT(8)"`
	Lastpublished int `xorm:"not null default 0 INT(10)"`
	Dateline      int `xorm:"not null default 0 INT(10)"`
	Status        int `xorm:"not null default 1 TINYINT(1)"`
}

func (t *PreConnectFeedlog) TableName() string {
	return "pre_connect_feedlog"
}
