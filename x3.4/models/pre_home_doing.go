package models

type PreHomeDoing struct {
	Doid     int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Uid      int    `xorm:"not null default 0 index(uid) MEDIUMINT(8)"`
	Username string `xorm:"not null default '''' VARCHAR(15)"`
	From     string `xorm:"not null default '''' VARCHAR(20)"`
	Dateline int    `xorm:"not null default 0 index index(uid) INT(10)"`
	Message  string `xorm:"not null TEXT"`
	Ip       string `xorm:"not null default '''' VARCHAR(20)"`
	Port     int    `xorm:"not null default 0 SMALLINT(6)"`
	Replynum int    `xorm:"not null default 0 INT(10)"`
	Status   int    `xorm:"not null default 0 TINYINT(1)"`
}
