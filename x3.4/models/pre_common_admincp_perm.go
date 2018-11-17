package models

type PreCommonAdmincpPerm struct {
	Cpgroupid int    `xorm:"not null pk unique(cpgroupperm) SMALLINT(6)"`
	Perm      string `xorm:"not null pk unique(cpgroupperm) VARCHAR(255)"`
}

func (t *PreCommonAdmincpPerm) TableName() string {
	return "pre_common_admincp_perm"
}
