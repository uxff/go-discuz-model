package models

type PreCommonDomain struct {
	Domain     string `xorm:"not null default '''' index(domain) CHAR(30)"`
	Domainroot string `xorm:"not null default '''' index(domain) CHAR(60)"`
	Id         int    `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Idtype     string `xorm:"not null pk default '''' index CHAR(15)"`
}

func (t *PreCommonDomain) TableName() string {
	return "pre_common_domain"
}
