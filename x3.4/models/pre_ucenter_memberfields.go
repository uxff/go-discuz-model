package models

type PreUcenterMemberfields struct {
	Uid       int    `xorm:"not null pk MEDIUMINT(8)"`
	Blacklist string `xorm:"not null TEXT"`
}

func (t *PreUcenterMemberfields) TableName() string {
	return "pre_ucenter_memberfields"
}
