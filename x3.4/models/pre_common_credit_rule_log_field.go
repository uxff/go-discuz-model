package models

type PreCommonCreditRuleLogField struct {
	Clid int    `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Uid  int    `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Info string `xorm:"not null TEXT"`
	User string `xorm:"not null TEXT"`
	App  string `xorm:"not null TEXT"`
}
