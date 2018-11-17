package models

type PreUcenterMemberfields struct {
	Uid       int    `xorm:"not null pk MEDIUMINT(8)"`
	Blacklist string `xorm:"not null TEXT"`
}
