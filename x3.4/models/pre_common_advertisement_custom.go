package models

type PreCommonAdvertisementCustom struct {
	Id   int    `xorm:"not null pk autoincr SMALLINT(5)"`
	Name string `xorm:"not null index VARCHAR(255)"`
}

func (t *PreCommonAdvertisementCustom) TableName() string {
	return "pre_common_advertisement_custom"
}
