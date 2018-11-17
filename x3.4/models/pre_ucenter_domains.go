package models

type PreUcenterDomains struct {
	Id     int    `xorm:"not null pk autoincr INT(10)"`
	Domain string `xorm:"not null default '''' CHAR(40)"`
	Ip     string `xorm:"not null default '''' CHAR(15)"`
}
