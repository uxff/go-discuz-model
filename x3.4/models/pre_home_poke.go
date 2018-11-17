package models

type PreHomePoke struct {
	Uid          int    `xorm:"not null pk default 0 index(uid) MEDIUMINT(8)"`
	Fromuid      int    `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Fromusername string `xorm:"not null default '''' VARCHAR(15)"`
	Note         string `xorm:"not null default '''' VARCHAR(255)"`
	Dateline     int    `xorm:"not null default 0 index(uid) INT(10)"`
	Iconid       int    `xorm:"not null default 0 SMALLINT(6)"`
}

func (t *PreHomePoke) TableName() string {
	return "pre_home_poke"
}
