package models

type PreCommonStyle struct {
	Styleid    int    `xorm:"not null pk autoincr SMALLINT(6)"`
	Name       string `xorm:"not null default '''' VARCHAR(20)"`
	Available  int    `xorm:"not null default 1 TINYINT(1)"`
	Templateid int    `xorm:"not null default 0 SMALLINT(6)"`
	Extstyle   string `xorm:"not null default '''' VARCHAR(255)"`
}

func (t *PreCommonStyle) TableName() string {
	return "pre_common_style"
}
