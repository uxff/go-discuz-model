package models

type PreCommonMemberActionLog struct {
	Id       int `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Uid      int `xorm:"not null default 0 index(dateline) MEDIUMINT(8)"`
	Action   int `xorm:"not null default 0 index(dateline) TINYINT(5)"`
	Dateline int `xorm:"not null default 0 index(dateline) INT(10)"`
}
