package models

type PreCommonPlugin struct {
	Pluginid    int    `xorm:"not null pk autoincr SMALLINT(6)"`
	Available   int    `xorm:"not null default 0 TINYINT(1)"`
	Adminid     int    `xorm:"not null default 0 TINYINT(1)"`
	Name        string `xorm:"not null default '''' VARCHAR(40)"`
	Identifier  string `xorm:"not null default '''' unique VARCHAR(40)"`
	Description string `xorm:"not null default '''' VARCHAR(255)"`
	Datatables  string `xorm:"not null default '''' VARCHAR(255)"`
	Directory   string `xorm:"not null default '''' VARCHAR(100)"`
	Copyright   string `xorm:"not null default '''' VARCHAR(100)"`
	Modules     string `xorm:"not null TEXT"`
	Version     string `xorm:"not null default '''' VARCHAR(20)"`
}
