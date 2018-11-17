package models

type PreMobileWechatResource struct {
	Id       int    `xorm:"not null pk autoincr INT(10)"`
	Name     string `xorm:"not null default '''' VARCHAR(255)"`
	Dateline int    `xorm:"not null INT(10)"`
	Type     int    `xorm:"not null default 0 index TINYINT(1)"`
	Data     string `xorm:"not null TEXT"`
}

func (t *PreMobileWechatResource) TableName() string {
	return "pre_mobile_wechat_resource"
}
