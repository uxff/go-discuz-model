package models

type PreHomeClickuser struct {
	Uid      int    `xorm:"not null default 0 index(uid) MEDIUMINT(8)"`
	Username string `xorm:"not null default '''' VARCHAR(15)"`
	Id       int    `xorm:"not null default 0 index(id) MEDIUMINT(8)"`
	Idtype   string `xorm:"not null default '''' index(id) index(uid) VARCHAR(15)"`
	Clickid  int    `xorm:"not null default 0 SMALLINT(6)"`
	Dateline int    `xorm:"not null default 0 index(id) index(uid) INT(10)"`
}

func (t *PreHomeClickuser) TableName() string {
	return "pre_home_clickuser"
}
