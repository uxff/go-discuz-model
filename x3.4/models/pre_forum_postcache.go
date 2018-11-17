package models

type PreForumPostcache struct {
	Pid      int    `xorm:"not null pk INT(10)"`
	Comment  string `xorm:"not null TEXT"`
	Rate     string `xorm:"not null TEXT"`
	Dateline int    `xorm:"not null default 0 index INT(10)"`
}
