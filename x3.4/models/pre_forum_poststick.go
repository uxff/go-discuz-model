package models

type PreForumPoststick struct {
	Tid      int `xorm:"not null pk index(dateline) MEDIUMINT(8)"`
	Pid      int `xorm:"not null pk INT(10)"`
	Position int `xorm:"not null INT(10)"`
	Dateline int `xorm:"not null index(dateline) INT(10)"`
}
