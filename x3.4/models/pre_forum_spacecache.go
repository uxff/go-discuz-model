package models

type PreForumSpacecache struct {
	Uid        int    `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Variable   string `xorm:"not null pk VARCHAR(20)"`
	Value      string `xorm:"not null TEXT"`
	Expiration int    `xorm:"not null default 0 INT(10)"`
}
