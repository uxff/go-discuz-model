package models

type PreHomeShow struct {
	Uid       int    `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Username  string `xorm:"not null default '''' VARCHAR(15)"`
	Unitprice int    `xorm:"not null default 1 index INT(10)"`
	Credit    int    `xorm:"not null default 0 index INT(10)"`
	Note      string `xorm:"not null default '''' VARCHAR(100)"`
}
