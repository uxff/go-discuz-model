package models

type PreCommonCardType struct {
	Id       int    `xorm:"not null pk autoincr SMALLINT(6)"`
	Typename string `xorm:"not null default '''' CHAR(20)"`
}
