package models

type PreCommonCard struct {
	Id            string `xorm:"not null pk default '''' CHAR(255)"`
	Typeid        int    `xorm:"not null default 1 SMALLINT(6)"`
	Maketype      int    `xorm:"not null default 0 TINYINT(1)"`
	Makeruid      int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Price         int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Extcreditskey int    `xorm:"not null default 0 TINYINT(1)"`
	Extcreditsval int    `xorm:"not null default 0 INT(10)"`
	Status        int    `xorm:"not null default 1 TINYINT(1)"`
	Dateline      int    `xorm:"not null default 0 index INT(10)"`
	Cleardateline int    `xorm:"not null default 0 INT(10)"`
	Useddateline  int    `xorm:"not null default 0 INT(10)"`
	Uid           int    `xorm:"not null default 0 MEDIUMINT(8)"`
}
