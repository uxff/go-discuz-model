package models

type PreForumPolloption struct {
	Polloptionid int    `xorm:"not null pk autoincr INT(10)"`
	Tid          int    `xorm:"not null default 0 index(tid) MEDIUMINT(8)"`
	Votes        int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Displayorder int    `xorm:"not null default 0 index(tid) TINYINT(3)"`
	Polloption   string `xorm:"not null default '''' VARCHAR(80)"`
	Voterids     string `xorm:"not null MEDIUMTEXT"`
}
