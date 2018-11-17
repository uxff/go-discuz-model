package models

type PreHomeDocomment struct {
	Id       int    `xorm:"not null pk autoincr INT(10)"`
	Upid     int    `xorm:"not null default 0 INT(10)"`
	Doid     int    `xorm:"not null default 0 index(doid) MEDIUMINT(8)"`
	Uid      int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Username string `xorm:"not null default '''' VARCHAR(15)"`
	Dateline int    `xorm:"not null default 0 index index(doid) INT(10)"`
	Message  string `xorm:"not null TEXT"`
	Ip       string `xorm:"not null default '''' VARCHAR(20)"`
	Grade    int    `xorm:"not null default 0 SMALLINT(6)"`
}

func (t *PreHomeDocomment) TableName() string {
	return "pre_home_docomment"
}
