package models

type PreCommonBlockPermission struct {
	Bid              int    `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Uid              int    `xorm:"not null pk default 0 index MEDIUMINT(8)"`
	Allowmanage      int    `xorm:"not null default 0 TINYINT(1)"`
	Allowrecommend   int    `xorm:"not null default 0 TINYINT(1)"`
	Needverify       int    `xorm:"not null default 0 TINYINT(1)"`
	Inheritedtplname string `xorm:"not null default '''' VARCHAR(255)"`
}
