package models

type PreCommonCreditRuleLog struct {
	Clid        int `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Uid         int `xorm:"not null default 0 index(uid) MEDIUMINT(8)"`
	Rid         int `xorm:"not null default 0 index(uid) MEDIUMINT(8)"`
	Fid         int `xorm:"not null default 0 index(uid) MEDIUMINT(8)"`
	Total       int `xorm:"not null default 0 MEDIUMINT(8)"`
	Cyclenum    int `xorm:"not null default 0 MEDIUMINT(8)"`
	Extcredits1 int `xorm:"not null default 0 INT(10)"`
	Extcredits2 int `xorm:"not null default 0 INT(10)"`
	Extcredits3 int `xorm:"not null default 0 INT(10)"`
	Extcredits4 int `xorm:"not null default 0 INT(10)"`
	Extcredits5 int `xorm:"not null default 0 INT(10)"`
	Extcredits6 int `xorm:"not null default 0 INT(10)"`
	Extcredits7 int `xorm:"not null default 0 INT(10)"`
	Extcredits8 int `xorm:"not null default 0 INT(10)"`
	Starttime   int `xorm:"not null default 0 INT(10)"`
	Dateline    int `xorm:"not null default 0 index INT(10)"`
}
