package models

type PreCommonFailedlogin struct {
	Ip         string `xorm:"not null pk default '''' CHAR(15)"`
	Username   string `xorm:"not null pk default '''' CHAR(32)"`
	Count      int    `xorm:"not null default 0 TINYINT(1)"`
	Lastupdate int    `xorm:"not null default 0 INT(10)"`
}
