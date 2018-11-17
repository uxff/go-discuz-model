package models

type PreCommonMemberVerifyInfo struct {
	Vid        int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Uid        int    `xorm:"not null default 0 index(uid) MEDIUMINT(8)"`
	Username   string `xorm:"not null default '''' VARCHAR(30)"`
	Verifytype int    `xorm:"not null default 0 index(uid) index(verifytype) TINYINT(1)"`
	Flag       int    `xorm:"not null default 0 index(verifytype) TINYINT(1)"`
	Field      string `xorm:"not null TEXT"`
	Dateline   int    `xorm:"not null default 0 index(uid) INT(10)"`
}
