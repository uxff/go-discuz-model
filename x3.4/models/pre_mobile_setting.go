package models

type PreMobileSetting struct {
	Skey   string `xorm:"not null pk default '''' VARCHAR(255)"`
	Svalue string `xorm:"not null TEXT"`
}

func (t *PreMobileSetting) TableName() string {
	return "pre_mobile_setting"
}
