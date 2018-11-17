package models

type PreCommonMyapp struct {
	Appid            int    `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Appname          string `xorm:"not null default '''' VARCHAR(60)"`
	Narrow           int    `xorm:"not null default 0 TINYINT(1)"`
	Flag             int    `xorm:"not null default 0 index(flag) TINYINT(1)"`
	Version          int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Userpanelarea    int    `xorm:"not null default 0 TINYINT(1)"`
	Canvastitle      string `xorm:"not null default '''' VARCHAR(60)"`
	Fullscreen       int    `xorm:"not null default 0 TINYINT(1)"`
	Displayuserpanel int    `xorm:"not null default 0 TINYINT(1)"`
	Displaymethod    int    `xorm:"not null default 0 TINYINT(1)"`
	Displayorder     int    `xorm:"not null default 0 index(flag) SMALLINT(6)"`
	Appstatus        int    `xorm:"not null default 0 TINYINT(2)"`
	Iconstatus       int    `xorm:"not null default 0 TINYINT(2)"`
	Icondowntime     int    `xorm:"not null default 0 INT(10)"`
}
