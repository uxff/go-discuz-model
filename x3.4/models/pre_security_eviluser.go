package models

type PreSecurityEviluser struct {
	Uid           int `xorm:"not null pk INT(10)"`
	Evilcount     int `xorm:"not null default 0 INT(10)"`
	Eviltype      int `xorm:"not null default 0 MEDIUMINT(8)"`
	Createtime    int `xorm:"not null default 0 index(operateresult) INT(10)"`
	Operateresult int `xorm:"not null default 0 index(operateresult) TINYINT(1)"`
	Isreported    int `xorm:"not null default 0 TINYINT(1)"`
}
