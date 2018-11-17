package models

type PreCommonMemberMedal struct {
	Uid     int `xorm:"not null pk MEDIUMINT(8)"`
	Medalid int `xorm:"not null pk SMALLINT(6)"`
}

func (t *PreCommonMemberMedal) TableName() string {
	return "pre_common_member_medal"
}
