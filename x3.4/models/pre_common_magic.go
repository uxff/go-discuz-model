package models

type PreCommonMagic struct {
	Magicid      int    `xorm:"not null pk autoincr SMALLINT(6)"`
	Available    int    `xorm:"not null default 0 index(displayorder) TINYINT(1)"`
	Name         string `xorm:"not null VARCHAR(50)"`
	Identifier   string `xorm:"not null unique VARCHAR(40)"`
	Description  string `xorm:"not null VARCHAR(255)"`
	Displayorder int    `xorm:"not null default 0 index(displayorder) TINYINT(3)"`
	Credit       int    `xorm:"not null default 0 TINYINT(1)"`
	Price        int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Num          int    `xorm:"not null default 0 SMALLINT(6)"`
	Salevolume   int    `xorm:"not null default 0 SMALLINT(6)"`
	Supplytype   int    `xorm:"not null default 0 TINYINT(1)"`
	Supplynum    int    `xorm:"not null default 0 SMALLINT(6)"`
	Useperoid    int    `xorm:"not null default 0 TINYINT(1)"`
	Usenum       int    `xorm:"not null default 0 SMALLINT(6)"`
	Weight       int    `xorm:"not null default 1 TINYINT(3)"`
	Magicperm    string `xorm:"not null TEXT"`
	Useevent     int    `xorm:"not null default 0 TINYINT(1)"`
}
