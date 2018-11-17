package models

type PreCommonCreditLog struct {
	Logid       int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Uid         int    `xorm:"not null default 0 index MEDIUMINT(8)"`
	Operation   string `xorm:"not null default '''' index CHAR(3)"`
	Relatedid   int    `xorm:"not null index INT(10)"`
	Dateline    int    `xorm:"not null index INT(10)"`
	Extcredits1 int    `xorm:"not null INT(10)"`
	Extcredits2 int    `xorm:"not null INT(10)"`
	Extcredits3 int    `xorm:"not null INT(10)"`
	Extcredits4 int    `xorm:"not null INT(10)"`
	Extcredits5 int    `xorm:"not null INT(10)"`
	Extcredits6 int    `xorm:"not null INT(10)"`
	Extcredits7 int    `xorm:"not null INT(10)"`
	Extcredits8 int    `xorm:"not null INT(10)"`
}
