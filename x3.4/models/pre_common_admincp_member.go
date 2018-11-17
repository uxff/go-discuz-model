package models

type PreCommonAdmincpMember struct {
	Uid        int    `xorm:"not null pk INT(10)"`
	Cpgroupid  int    `xorm:"not null INT(10)"`
	Customperm string `xorm:"not null TEXT"`
}
