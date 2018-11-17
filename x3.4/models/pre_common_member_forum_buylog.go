package models

type PreCommonMemberForumBuylog struct {
	Uid     int `xorm:"not null pk MEDIUMINT(8)"`
	Fid     int `xorm:"not null pk default 0 index MEDIUMINT(8)"`
	Credits int `xorm:"not null default 0 INT(10)"`
}
