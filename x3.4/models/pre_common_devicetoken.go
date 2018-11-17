package models

type PreCommonDevicetoken struct {
	Uid   int    `xorm:"not null pk MEDIUMINT(8)"`
	Token string `xorm:"not null index CHAR(50)"`
}
