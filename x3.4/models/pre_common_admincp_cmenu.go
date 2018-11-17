package models

type PreCommonAdmincpCmenu struct {
	Id           int    `xorm:"not null pk autoincr SMALLINT(6)"`
	Title        string `xorm:"not null VARCHAR(255)"`
	Url          string `xorm:"not null VARCHAR(255)"`
	Sort         int    `xorm:"not null default 0 TINYINT(1)"`
	Displayorder int    `xorm:"not null index TINYINT(3)"`
	Clicks       int    `xorm:"not null default 1 SMALLINT(6)"`
	Uid          int    `xorm:"not null index MEDIUMINT(8)"`
	Dateline     int    `xorm:"not null INT(10)"`
}

func (t *PreCommonAdmincpCmenu) TableName() string {
	return "pre_common_admincp_cmenu"
}
