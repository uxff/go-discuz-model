package models

type PreHomePicfield struct {
	Picid   int    `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Hotuser string `xorm:"not null TEXT"`
}
