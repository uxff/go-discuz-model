package models

type PreCommonTaskvar struct {
	Taskvarid   int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Taskid      int    `xorm:"not null default 0 index SMALLINT(6)"`
	Sort        string `xorm:"not null default ''complete'' ENUM('apply','complete')"`
	Name        string `xorm:"not null default '''' VARCHAR(100)"`
	Description string `xorm:"not null default '''' VARCHAR(255)"`
	Variable    string `xorm:"not null default '''' VARCHAR(40)"`
	Type        string `xorm:"not null default ''text'' VARCHAR(20)"`
	Value       string `xorm:"not null TEXT"`
}
