package models

type PreHomeClass struct {
	Classid   int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Classname string `xorm:"not null default '''' CHAR(40)"`
	Uid       int    `xorm:"not null default 0 index MEDIUMINT(8)"`
	Dateline  int    `xorm:"not null default 0 INT(10)"`
}

func (t *PreHomeClass) TableName() string {
	return "pre_home_class"
}
