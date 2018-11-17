package models

type PreHomeBlacklist struct {
	Uid      int `xorm:"not null pk default 0 index(uid) MEDIUMINT(8)"`
	Buid     int `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Dateline int `xorm:"not null default 0 index(uid) INT(10)"`
}
