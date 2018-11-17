package models

type PreCommonMemberCrime struct {
	Cid        int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Uid        int    `xorm:"not null default 0 index(uid) MEDIUMINT(8)"`
	Operatorid int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Operator   string `xorm:"not null VARCHAR(15)"`
	Action     int    `xorm:"not null index(uid) TINYINT(5)"`
	Reason     string `xorm:"not null TEXT"`
	Dateline   int    `xorm:"not null default 0 index(uid) INT(10)"`
}
