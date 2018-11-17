package models

type PreCommonAdmincpSession struct {
	Uid        int    `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Adminid    int    `xorm:"not null default 0 SMALLINT(6)"`
	Panel      int    `xorm:"not null pk default 0 TINYINT(1)"`
	Ip         string `xorm:"not null default '''' VARCHAR(15)"`
	Dateline   int    `xorm:"not null default 0 INT(10)"`
	Errorcount int    `xorm:"not null default 0 TINYINT(1)"`
	Storage    string `xorm:"not null MEDIUMTEXT"`
}

func (t *PreCommonAdmincpSession) TableName() string {
	return "pre_common_admincp_session"
}
