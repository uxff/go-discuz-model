package models

type PreForumThreadprofileGroup struct {
	Gid  int `xorm:"not null pk MEDIUMINT(8)"`
	Tpid int `xorm:"not null MEDIUMINT(8)"`
}
