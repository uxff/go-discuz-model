package models

type PreCommonUinBlack struct {
	Uin      string `xorm:"not null pk CHAR(40)"`
	Uid      int    `xorm:"not null default 0 unique MEDIUMINT(8)"`
	Dateline int    `xorm:"not null default 0 INT(10)"`
}

func (t *PreCommonUinBlack) TableName() string {
	return "pre_common_uin_black"
}
