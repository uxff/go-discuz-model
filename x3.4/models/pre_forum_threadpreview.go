package models

type PreForumThreadpreview struct {
	Tid     int    `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Relay   int    `xorm:"not null default 0 INT(10)"`
	Content string `xorm:"not null TEXT"`
}
