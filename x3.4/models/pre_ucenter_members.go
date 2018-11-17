package models

type PreUcenterMembers struct {
	Uid           int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Username      string `xorm:"not null default '''' unique CHAR(15)"`
	Password      string `xorm:"not null default '''' CHAR(32)"`
	Email         string `xorm:"not null default '''' index CHAR(32)"`
	Myid          string `xorm:"not null default '''' CHAR(30)"`
	Myidkey       string `xorm:"not null default '''' CHAR(16)"`
	Regip         string `xorm:"not null default '''' CHAR(15)"`
	Regdate       int    `xorm:"not null default 0 INT(10)"`
	Lastloginip   int    `xorm:"not null default 0 INT(10)"`
	Lastlogintime int    `xorm:"not null default 0 INT(10)"`
	Salt          string `xorm:"not null CHAR(6)"`
	Secques       string `xorm:"not null default '''' CHAR(8)"`
}
