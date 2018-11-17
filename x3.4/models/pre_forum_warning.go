package models

type PreForumWarning struct {
	Wid        int    `xorm:"not null pk autoincr INT(10)"`
	Pid        int    `xorm:"not null unique INT(10)"`
	Operatorid int    `xorm:"not null MEDIUMINT(8)"`
	Operator   string `xorm:"not null CHAR(15)"`
	Authorid   int    `xorm:"not null index MEDIUMINT(8)"`
	Author     string `xorm:"not null CHAR(15)"`
	Dateline   int    `xorm:"not null INT(10)"`
	Reason     string `xorm:"not null CHAR(40)"`
}

func (t *PreForumWarning) TableName() string {
	return "pre_forum_warning"
}
