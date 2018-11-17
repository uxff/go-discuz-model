package models

type PreUcenterNewpm struct {
	Uid int `xorm:"not null pk MEDIUMINT(8)"`
}

func (t *PreUcenterNewpm) TableName() string {
	return "pre_ucenter_newpm"
}
