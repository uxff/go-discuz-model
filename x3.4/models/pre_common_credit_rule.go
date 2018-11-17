package models

type PreCommonCreditRule struct {
	Rid         int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Rulename    string `xorm:"not null default '''' VARCHAR(20)"`
	Action      string `xorm:"not null default '''' unique VARCHAR(20)"`
	Cycletype   int    `xorm:"not null default 0 TINYINT(1)"`
	Cycletime   int    `xorm:"not null default 0 INT(10)"`
	Rewardnum   int    `xorm:"not null default 1 TINYINT(2)"`
	Norepeat    int    `xorm:"not null default 0 TINYINT(1)"`
	Extcredits1 int    `xorm:"not null default 0 INT(10)"`
	Extcredits2 int    `xorm:"not null default 0 INT(10)"`
	Extcredits3 int    `xorm:"not null default 0 INT(10)"`
	Extcredits4 int    `xorm:"not null default 0 INT(10)"`
	Extcredits5 int    `xorm:"not null default 0 INT(10)"`
	Extcredits6 int    `xorm:"not null default 0 INT(10)"`
	Extcredits7 int    `xorm:"not null default 0 INT(10)"`
	Extcredits8 int    `xorm:"not null default 0 INT(10)"`
	Fids        string `xorm:"not null TEXT"`
}

func (t *PreCommonCreditRule) TableName() string {
	return "pre_common_credit_rule"
}
