package models

type PreCommonTemplatePermission struct {
	Targettplname    string `xorm:"not null pk default '''' VARCHAR(100)"`
	Uid              int    `xorm:"not null pk default 0 index MEDIUMINT(8)"`
	Allowmanage      int    `xorm:"not null default 0 TINYINT(1)"`
	Allowrecommend   int    `xorm:"not null default 0 TINYINT(1)"`
	Needverify       int    `xorm:"not null default 0 TINYINT(1)"`
	Inheritedtplname string `xorm:"not null default '''' VARCHAR(255)"`
}

func (t *PreCommonTemplatePermission) TableName() string {
	return "pre_common_template_permission"
}
