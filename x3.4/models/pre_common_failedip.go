package models

type PreCommonFailedip struct {
	Ip         string `xorm:"not null pk default '''' CHAR(7)"`
	Lastupdate int    `xorm:"not null pk default 0 index INT(10)"`
	Count      int    `xorm:"not null default 0 TINYINT(1)"`
}

func (t *PreCommonFailedip) TableName() string {
	return "pre_common_failedip"
}
