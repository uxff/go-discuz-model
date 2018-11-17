package models

type PreForumGroupfield struct {
	Fid      int    `xorm:"not null pk default 0 index unique(types) MEDIUMINT(8)"`
	Privacy  int    `xorm:"not null default 0 TINYINT(1)"`
	Dateline int    `xorm:"not null default 0 INT(10)"`
	Type     string `xorm:"not null pk index unique(types) VARCHAR(255)"`
	Data     string `xorm:"not null TEXT"`
}
