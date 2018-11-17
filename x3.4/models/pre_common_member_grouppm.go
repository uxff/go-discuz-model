package models

type PreCommonMemberGrouppm struct {
	Uid      int `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Gpmid    int `xorm:"not null pk autoincr SMALLINT(6)"`
	Status   int `xorm:"not null default 0 TINYINT(1)"`
	Dateline int `xorm:"not null default 0 INT(10)"`
}
