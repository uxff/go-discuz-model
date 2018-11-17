package models

type PreCommonMemberWechatmp struct {
	Uid    int    `xorm:"not null pk MEDIUMINT(8)"`
	Openid string `xorm:"not null default '''' index CHAR(32)"`
	Status int    `xorm:"not null default 0 TINYINT(1)"`
}

func (t *PreCommonMemberWechatmp) TableName() string {
	return "pre_common_member_wechatmp"
}
