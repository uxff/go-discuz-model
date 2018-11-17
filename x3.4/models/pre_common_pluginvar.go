package models

type PreCommonPluginvar struct {
	Pluginvarid  int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Pluginid     int    `xorm:"not null default 0 index SMALLINT(6)"`
	Displayorder int    `xorm:"not null default 0 TINYINT(3)"`
	Title        string `xorm:"not null default '''' VARCHAR(100)"`
	Description  string `xorm:"not null default '''' VARCHAR(255)"`
	Variable     string `xorm:"not null default '''' VARCHAR(40)"`
	Type         string `xorm:"not null default ''text'' VARCHAR(20)"`
	Value        string `xorm:"not null TEXT"`
	Extra        string `xorm:"not null TEXT"`
}
