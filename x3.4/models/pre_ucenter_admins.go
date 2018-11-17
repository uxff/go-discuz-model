package models

type PreUcenterAdmins struct {
	Uid               int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Username          string `xorm:"not null default '''' unique CHAR(15)"`
	Allowadminsetting int    `xorm:"not null default 0 TINYINT(1)"`
	Allowadminapp     int    `xorm:"not null default 0 TINYINT(1)"`
	Allowadminuser    int    `xorm:"not null default 0 TINYINT(1)"`
	Allowadminbadword int    `xorm:"not null default 0 TINYINT(1)"`
	Allowadmintag     int    `xorm:"not null default 0 TINYINT(1)"`
	Allowadminpm      int    `xorm:"not null default 0 TINYINT(1)"`
	Allowadmincredits int    `xorm:"not null default 0 TINYINT(1)"`
	Allowadmindomain  int    `xorm:"not null default 0 TINYINT(1)"`
	Allowadmindb      int    `xorm:"not null default 0 TINYINT(1)"`
	Allowadminnote    int    `xorm:"not null default 0 TINYINT(1)"`
	Allowadmincache   int    `xorm:"not null default 0 TINYINT(1)"`
	Allowadminlog     int    `xorm:"not null default 0 TINYINT(1)"`
}

func (t *PreUcenterAdmins) TableName() string {
	return "pre_ucenter_admins"
}
