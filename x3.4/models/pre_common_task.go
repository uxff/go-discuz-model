package models

type PreCommonTask struct {
	Taskid        int    `xorm:"not null pk autoincr SMALLINT(6)"`
	Relatedtaskid int    `xorm:"not null default 0 SMALLINT(6)"`
	Available     int    `xorm:"not null default 0 TINYINT(1)"`
	Name          string `xorm:"not null default '''' VARCHAR(50)"`
	Description   string `xorm:"not null TEXT"`
	Icon          string `xorm:"not null default '''' VARCHAR(150)"`
	Applicants    int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Achievers     int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Tasklimits    int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Applyperm     string `xorm:"not null TEXT"`
	Scriptname    string `xorm:"not null default '''' VARCHAR(50)"`
	Starttime     int    `xorm:"not null default 0 INT(10)"`
	Endtime       int    `xorm:"not null default 0 INT(10)"`
	Period        int    `xorm:"not null default 0 INT(10)"`
	Periodtype    int    `xorm:"not null default 0 TINYINT(1)"`
	Reward        string `xorm:"not null default ''credit'' ENUM('credit','group','invite','magic','medal')"`
	Prize         string `xorm:"not null default '''' VARCHAR(15)"`
	Bonus         int    `xorm:"not null default 0 INT(10)"`
	Displayorder  int    `xorm:"not null default 0 SMALLINT(6)"`
	Version       string `xorm:"not null default '''' VARCHAR(15)"`
}

func (t *PreCommonTask) TableName() string {
	return "pre_common_task"
}
