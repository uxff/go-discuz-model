package models

type PreCommonMyinvite struct {
	Id       int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Typename string `xorm:"not null default '''' VARCHAR(100)"`
	Appid    int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Type     int    `xorm:"not null default 0 TINYINT(1)"`
	Fromuid  int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Touid    int    `xorm:"not null default 0 index(uid) MEDIUMINT(8)"`
	Myml     string `xorm:"not null TEXT"`
	Dateline int    `xorm:"not null default 0 index(uid) INT(10)"`
	Hash     int    `xorm:"not null default 0 index INT(10)"`
}
