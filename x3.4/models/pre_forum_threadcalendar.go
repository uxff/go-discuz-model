package models

type PreForumThreadcalendar struct {
	Cid      int `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Fid      int `xorm:"not null default 0 index(fid) MEDIUMINT(8)"`
	Dateline int `xorm:"not null default 0 index(fid) INT(10)"`
	Hotnum   int `xorm:"not null default 0 INT(10)"`
}
