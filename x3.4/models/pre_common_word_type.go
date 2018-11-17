package models

type PreCommonWordType struct {
	Id       int    `xorm:"not null pk autoincr SMALLINT(6)"`
	Typename string `xorm:"not null default '''' VARCHAR(15)"`
}
