package models

type PreCommonCardLog struct {
	Id          int    `xorm:"not null pk autoincr SMALLINT(6)"`
	Uid         int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Username    string `xorm:"not null default '''' VARCHAR(20)"`
	Cardrule    string `xorm:"not null default '''' VARCHAR(255)"`
	Info        string `xorm:"not null TEXT"`
	Dateline    int    `xorm:"not null default 0 index index(operation_dateline) INT(10)"`
	Description string `xorm:"not null MEDIUMTEXT"`
	Operation   int    `xorm:"not null default 0 index(operation_dateline) TINYINT(1)"`
}

func (t *PreCommonCardLog) TableName() string {
	return "pre_common_card_log"
}
