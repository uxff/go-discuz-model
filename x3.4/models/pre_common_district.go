package models

type PreCommonDistrict struct {
	Id           int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Name         string `xorm:"not null default '''' VARCHAR(255)"`
	Level        int    `xorm:"not null default 0 TINYINT(4)"`
	Usetype      int    `xorm:"not null default 0 TINYINT(1)"`
	Upid         int    `xorm:"not null default 0 index(upid) MEDIUMINT(8)"`
	Displayorder int    `xorm:"not null default 0 index(upid) SMALLINT(6)"`
}
