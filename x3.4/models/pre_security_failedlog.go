package models

type PreSecurityFailedlog struct {
	Id           int    `xorm:"not null pk autoincr INT(11)"`
	Reporttype   string `xorm:"not null CHAR(20)"`
	Tid          int    `xorm:"not null default 0 INT(10)"`
	Pid          int    `xorm:"not null default 0 index INT(10)"`
	Uid          int    `xorm:"not null default 0 index INT(10)"`
	Failcount    int    `xorm:"not null default 0 INT(10)"`
	Createtime   int    `xorm:"not null default 0 INT(10)"`
	Posttime     int    `xorm:"not null default 0 INT(10)"`
	Delreason    string `xorm:"not null CHAR(255)"`
	Scheduletime int    `xorm:"not null default 0 INT(10)"`
	Lastfailtime int    `xorm:"not null default 0 INT(10)"`
	Extra1       int    `xorm:"not null INT(10)"`
	Extra2       string `xorm:"not null default ''0'' CHAR(255)"`
}

func (t *PreSecurityFailedlog) TableName() string {
	return "pre_security_failedlog"
}
