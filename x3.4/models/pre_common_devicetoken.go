package models

type PreCommonDevicetoken struct {
	Uid   int    `xorm:"not null pk MEDIUMINT(8)"`
	Token string `xorm:"not null index CHAR(50)"`
}

func (t *PreCommonDevicetoken) TableName() string {
	return "pre_common_devicetoken"
}
