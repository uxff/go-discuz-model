package models

type PreCommonCardType struct {
	Id       int    `xorm:"not null pk autoincr SMALLINT(6)"`
	Typename string `xorm:"not null default '''' CHAR(20)"`
}

func (t *PreCommonCardType) TableName() string {
	return "pre_common_card_type"
}
