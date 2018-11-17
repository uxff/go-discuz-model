package models

type PreCommonSearchindex struct {
	Searchid     int    `xorm:"not null pk autoincr INT(10)"`
	Srchmod      int    `xorm:"not null index TINYINT(3)"`
	Keywords     string `xorm:"not null default '''' VARCHAR(255)"`
	Searchstring string `xorm:"not null TEXT"`
	Useip        string `xorm:"not null default '''' VARCHAR(15)"`
	Uid          int    `xorm:"not null default 0 MEDIUMINT(10)"`
	Dateline     int    `xorm:"not null default 0 INT(10)"`
	Expiration   int    `xorm:"not null default 0 INT(10)"`
	Threadsortid int    `xorm:"not null default 0 SMALLINT(6)"`
	Num          int    `xorm:"not null default 0 SMALLINT(6)"`
	Ids          string `xorm:"not null TEXT"`
}
