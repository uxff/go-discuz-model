package models

type PreCommonStatuser struct {
	Uid     int    `xorm:"not null default 0 index MEDIUMINT(8)"`
	Daytime int    `xorm:"not null default 0 INT(10)"`
	Type    string `xorm:"not null default '''' CHAR(20)"`
}

func (t *PreCommonStatuser) TableName() string {
	return "pre_common_statuser"
}
