package models

type PreCommonAdmincpPerm struct {
	Cpgroupid int    `xorm:"not null pk unique(cpgroupperm) SMALLINT(6)"`
	Perm      string `xorm:"not null pk unique(cpgroupperm) VARCHAR(255)"`
}
