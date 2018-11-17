package models

type PreCommonTemplateBlock struct {
	Targettplname string `xorm:"not null pk default '''' VARCHAR(100)"`
	Tpldirectory  string `xorm:"not null pk default '''' VARCHAR(80)"`
	Bid           int    `xorm:"not null pk default 0 index MEDIUMINT(8)"`
}
