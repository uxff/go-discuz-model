package models

type PreHomeVisitor struct {
	Uid       int    `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Vuid      int    `xorm:"not null pk default 0 index MEDIUMINT(8)"`
	Vusername string `xorm:"not null default '''' CHAR(15)"`
	Dateline  int    `xorm:"not null default 0 index INT(10)"`
}

func (t *PreHomeVisitor) TableName() string {
	return "pre_home_visitor"
}
