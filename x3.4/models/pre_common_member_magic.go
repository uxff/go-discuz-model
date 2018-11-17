package models

type PreCommonMemberMagic struct {
	Uid     int `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Magicid int `xorm:"not null pk default 0 SMALLINT(6)"`
	Num     int `xorm:"not null default 0 SMALLINT(6)"`
}

func (t *PreCommonMemberMagic) TableName() string {
	return "pre_common_member_magic"
}
