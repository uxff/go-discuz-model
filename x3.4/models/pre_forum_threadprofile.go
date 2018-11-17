package models

type PreForumThreadprofile struct {
	Id       int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Name     string `xorm:"not null default '''' CHAR(100)"`
	Template string `xorm:"not null TEXT"`
	Global   int    `xorm:"not null default 0 index TINYINT(1)"`
}
