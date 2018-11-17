package models

type PreHomeCommentModerate struct {
	Id       int    `xorm:"not null pk default 0 INT(10)"`
	Idtype   string `xorm:"not null default '''' index(idtype) VARCHAR(15)"`
	Status   int    `xorm:"not null default 0 index(idtype) TINYINT(3)"`
	Dateline int    `xorm:"not null default 0 index(idtype) INT(10)"`
}
