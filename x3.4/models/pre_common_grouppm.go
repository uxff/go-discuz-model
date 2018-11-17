package models

type PreCommonGrouppm struct {
	Id       int    `xorm:"not null pk autoincr SMALLINT(6)"`
	Authorid int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Author   string `xorm:"not null default '''' VARCHAR(15)"`
	Dateline int    `xorm:"not null default 0 INT(10)"`
	Message  string `xorm:"not null TEXT"`
	Numbers  int    `xorm:"not null default 0 MEDIUMINT(8)"`
}

func (t *PreCommonGrouppm) TableName() string {
	return "pre_common_grouppm"
}
