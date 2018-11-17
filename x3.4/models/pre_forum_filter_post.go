package models

type PreForumFilterPost struct {
	Tid        int `xorm:"not null pk default 0 index(tid) MEDIUMINT(8)"`
	Pid        int `xorm:"not null pk default 0 INT(10)"`
	Postlength int `xorm:"not null default 0 index(tid) INT(10)"`
}
