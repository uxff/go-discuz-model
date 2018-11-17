package models

type PreCommonCron struct {
	Cronid    int    `xorm:"not null pk autoincr SMALLINT(6)"`
	Available int    `xorm:"not null default 0 index(nextrun) TINYINT(1)"`
	Type      string `xorm:"not null default ''user'' ENUM('plugin','system','user')"`
	Name      string `xorm:"not null default '''' CHAR(50)"`
	Filename  string `xorm:"not null default '''' CHAR(50)"`
	Lastrun   int    `xorm:"not null default 0 INT(10)"`
	Nextrun   int    `xorm:"not null default 0 index(nextrun) INT(10)"`
	Weekday   int    `xorm:"not null default 0 TINYINT(1)"`
	Day       int    `xorm:"not null default 0 TINYINT(2)"`
	Hour      int    `xorm:"not null default 0 TINYINT(2)"`
	Minute    string `xorm:"not null default '''' CHAR(36)"`
}
