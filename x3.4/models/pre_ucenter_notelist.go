package models

type PreUcenterNotelist struct {
	Noteid     int    `xorm:"not null pk autoincr index(closed) INT(10)"`
	Operation  string `xorm:"not null CHAR(32)"`
	Closed     int    `xorm:"not null default 0 index(closed) TINYINT(4)"`
	Totalnum   int    `xorm:"not null default 0 SMALLINT(6)"`
	Succeednum int    `xorm:"not null default 0 SMALLINT(6)"`
	Getdata    string `xorm:"not null MEDIUMTEXT"`
	Postdata   string `xorm:"not null MEDIUMTEXT"`
	Dateline   int    `xorm:"not null default 0 index INT(10)"`
	Pri        int    `xorm:"not null default 0 index(closed) TINYINT(3)"`
	App1       int    `xorm:"not null TINYINT(4)"`
}
