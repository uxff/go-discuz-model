package models

type PreCommonCreditRuleLogField struct {
	Clid int    `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Uid  int    `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Info string `xorm:"not null TEXT"`
	User string `xorm:"not null TEXT"`
	App  string `xorm:"not null TEXT"`
}

func (t *PreCommonCreditRuleLogField) TableName() string {
	return "pre_common_credit_rule_log_field"
}
