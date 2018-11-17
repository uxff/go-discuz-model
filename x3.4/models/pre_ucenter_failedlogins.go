package models

type PreUcenterFailedlogins struct {
	Ip         string `xorm:"not null pk default '''' CHAR(15)"`
	Count      int    `xorm:"not null default 0 TINYINT(1)"`
	Lastupdate int    `xorm:"not null default 0 INT(10)"`
}
