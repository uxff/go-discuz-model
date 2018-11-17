package models

type PreCommonAdvertisement struct {
	Advid        int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Available    int    `xorm:"not null default 0 TINYINT(1)"`
	Type         string `xorm:"not null default ''0'' VARCHAR(50)"`
	Displayorder int    `xorm:"not null default 0 TINYINT(3)"`
	Title        string `xorm:"not null default '''' VARCHAR(255)"`
	Targets      string `xorm:"not null TEXT"`
	Parameters   string `xorm:"not null TEXT"`
	Code         string `xorm:"not null TEXT"`
	Starttime    int    `xorm:"not null default 0 INT(10)"`
	Endtime      int    `xorm:"not null default 0 INT(10)"`
}
