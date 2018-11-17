package models

type PreCommonMemberWechat struct {
	Uid        int    `xorm:"not null pk MEDIUMINT(8)"`
	Openid     string `xorm:"not null default '''' unique CHAR(32)"`
	Status     int    `xorm:"not null default 0 TINYINT(1)"`
	Isregister int    `xorm:"not null default 0 TINYINT(1)"`
}
