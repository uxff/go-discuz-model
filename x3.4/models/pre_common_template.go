package models

type PreCommonTemplate struct {
	Templateid int    `xorm:"not null pk autoincr SMALLINT(6)"`
	Name       string `xorm:"not null default '''' VARCHAR(30)"`
	Directory  string `xorm:"not null default '''' VARCHAR(100)"`
	Copyright  string `xorm:"not null default '''' VARCHAR(100)"`
}

func (t *PreCommonTemplate) TableName() string {
	return "pre_common_template"
}
