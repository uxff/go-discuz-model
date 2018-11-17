package models

type PreCommonMemberLog struct {
	Uid      int    `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Action   string `xorm:"not null default '''' CHAR(10)"`
	Dateline int    `xorm:"not null default 0 INT(10)"`
}
