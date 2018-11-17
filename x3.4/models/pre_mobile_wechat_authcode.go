package models

type PreMobileWechatAuthcode struct {
	Sid        string `xorm:"not null pk CHAR(6)"`
	Code       int    `xorm:"not null unique INT(10)"`
	Uid        int    `xorm:"not null MEDIUMINT(8)"`
	Status     int    `xorm:"not null default 0 TINYINT(1)"`
	Createtime int    `xorm:"not null index INT(10)"`
}
