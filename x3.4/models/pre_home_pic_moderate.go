package models

type PreHomePicModerate struct {
	Id       int `xorm:"not null pk default 0 INT(10)"`
	Status   int `xorm:"not null default 0 index(status) TINYINT(3)"`
	Dateline int `xorm:"not null default 0 index(status) INT(10)"`
}

func (t *PreHomePicModerate) TableName() string {
	return "pre_home_pic_moderate"
}
