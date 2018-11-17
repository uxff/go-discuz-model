package models

type PreCommonSmiley struct {
	Id           int    `xorm:"not null pk autoincr SMALLINT(6)"`
	Typeid       int    `xorm:"not null SMALLINT(6)"`
	Displayorder int    `xorm:"not null default 0 index(type) TINYINT(1)"`
	Type         string `xorm:"not null default ''smiley'' index(type) ENUM('smiley','stamp','stamplist')"`
	Code         string `xorm:"not null default '''' VARCHAR(30)"`
	Url          string `xorm:"not null default '''' VARCHAR(30)"`
}
