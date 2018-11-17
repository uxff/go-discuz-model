package models

type PreCommonSetting struct {
	Skey   string `xorm:"not null pk default '''' VARCHAR(255)"`
	Svalue string `xorm:"not null TEXT"`
}

func (t *PreCommonSetting) TableName() string {
	return "pre_common_setting"
}
