package models

type PreCommonStylevar struct {
	Stylevarid int    `xorm:"not null pk autoincr SMALLINT(6)"`
	Styleid    int    `xorm:"not null default 0 index SMALLINT(6)"`
	Variable   string `xorm:"not null TEXT"`
	Substitute string `xorm:"not null TEXT"`
}
