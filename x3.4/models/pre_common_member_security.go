package models

type PreCommonMemberSecurity struct {
	Securityid int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Uid        int    `xorm:"not null default 0 index(uid) MEDIUMINT(8)"`
	Username   string `xorm:"not null default '''' VARCHAR(255)"`
	Fieldid    string `xorm:"not null default '''' index(uid) VARCHAR(255)"`
	Oldvalue   string `xorm:"not null TEXT"`
	Newvalue   string `xorm:"not null TEXT"`
	Dateline   int    `xorm:"not null default 0 index INT(10)"`
}
