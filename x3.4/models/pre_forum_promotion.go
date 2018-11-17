package models

type PreForumPromotion struct {
	Ip       string `xorm:"not null pk default '''' CHAR(15)"`
	Uid      int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Username string `xorm:"not null default '''' CHAR(15)"`
}
