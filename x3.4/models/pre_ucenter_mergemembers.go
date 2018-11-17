package models

type PreUcenterMergemembers struct {
	Appid    int    `xorm:"not null pk SMALLINT(6)"`
	Username string `xorm:"not null pk CHAR(15)"`
}

func (t *PreUcenterMergemembers) TableName() string {
	return "pre_ucenter_mergemembers"
}
