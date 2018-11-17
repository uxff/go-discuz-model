package models

type PreCommonDiyData struct {
	Targettplname string `xorm:"not null pk default '''' VARCHAR(100)"`
	Tpldirectory  string `xorm:"not null pk default '''' VARCHAR(80)"`
	Primaltplname string `xorm:"not null default '''' VARCHAR(255)"`
	Diycontent    string `xorm:"not null MEDIUMTEXT"`
	Name          string `xorm:"not null default '''' VARCHAR(255)"`
	Uid           int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Username      string `xorm:"not null default '''' VARCHAR(15)"`
	Dateline      int    `xorm:"not null default 0 INT(10)"`
}

func (t *PreCommonDiyData) TableName() string {
	return "pre_common_diy_data"
}
