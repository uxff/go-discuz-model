package models

type PreHomeNotification struct {
	Id         int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Uid        int    `xorm:"not null default 0 index(by_type) index(category) index(uid) MEDIUMINT(8)"`
	Type       string `xorm:"not null default '''' index(by_type) VARCHAR(20)"`
	New        int    `xorm:"not null default 0 index(uid) TINYINT(1)"`
	Authorid   int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Author     string `xorm:"not null default '''' VARCHAR(15)"`
	Note       string `xorm:"not null TEXT"`
	Dateline   int    `xorm:"not null default 0 index(by_type) index(category) INT(10)"`
	FromId     int    `xorm:"not null default 0 index(from_id) MEDIUMINT(8)"`
	FromIdtype string `xorm:"not null default '''' index(from_id) VARCHAR(20)"`
	FromNum    int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Category   int    `xorm:"not null default 0 index(category) TINYINT(1)"`
}
