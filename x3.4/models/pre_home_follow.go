package models

type PreHomeFollow struct {
	Uid       int    `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Username  string `xorm:"not null default '''' CHAR(15)"`
	Followuid int    `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Fusername string `xorm:"not null default '''' CHAR(15)"`
	Bkname    string `xorm:"not null default '''' VARCHAR(255)"`
	Status    int    `xorm:"not null default 0 TINYINT(1)"`
	Mutual    int    `xorm:"not null default 0 TINYINT(1)"`
	Dateline  int    `xorm:"not null default 0 INT(10)"`
}
