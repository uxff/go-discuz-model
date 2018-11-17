package models

type PreHomeComment struct {
	Cid          int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Uid          int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Id           int    `xorm:"not null default 0 index(id) MEDIUMINT(8)"`
	Idtype       string `xorm:"not null default '''' index(authorid) index(id) VARCHAR(20)"`
	Authorid     int    `xorm:"not null default 0 index(authorid) MEDIUMINT(8)"`
	Author       string `xorm:"not null default '''' VARCHAR(15)"`
	Ip           string `xorm:"not null default '''' VARCHAR(20)"`
	Port         int    `xorm:"not null default 0 SMALLINT(6)"`
	Dateline     int    `xorm:"not null default 0 index(id) INT(10)"`
	Message      string `xorm:"not null TEXT"`
	Magicflicker int    `xorm:"not null default 0 TINYINT(1)"`
	Status       int    `xorm:"not null default 0 TINYINT(1)"`
}

func (t *PreHomeComment) TableName() string {
	return "pre_home_comment"
}
