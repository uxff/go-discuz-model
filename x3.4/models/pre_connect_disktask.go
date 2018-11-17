package models

type PreConnectDisktask struct {
	Taskid       int    `xorm:"not null pk autoincr INT(10)"`
	Aid          int    `xorm:"not null default 0 INT(10)"`
	Uid          int    `xorm:"not null default 0 INT(10)"`
	Openid       string `xorm:"not null default '''' index CHAR(32)"`
	Filename     string `xorm:"not null default '''' VARCHAR(255)"`
	Verifycode   string `xorm:"not null default '''' CHAR(32)"`
	Status       int    `xorm:"not null default 0 index SMALLINT(6)"`
	Dateline     int    `xorm:"not null default 0 INT(10)"`
	Downloadtime int    `xorm:"not null default 0 INT(10)"`
	Extra        string `xorm:"default 'NULL' TEXT"`
}
