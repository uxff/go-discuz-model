package models

type PreCommonAdmincpGroup struct {
	Cpgroupid   int    `xorm:"not null pk autoincr SMALLINT(6)"`
	Cpgroupname string `xorm:"not null VARCHAR(255)"`
}

func (t *PreCommonAdmincpGroup) TableName() string {
	return "pre_common_admincp_group"
}
