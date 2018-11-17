package models

type PreCommonAdmincpGroup struct {
	Cpgroupid   int    `xorm:"not null pk autoincr SMALLINT(6)"`
	Cpgroupname string `xorm:"not null VARCHAR(255)"`
}
