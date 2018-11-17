package models

type PreCommonOptimizer struct {
	K string `xorm:"not null pk default '''' CHAR(100)"`
	V string `xorm:"not null default '''' CHAR(255)"`
}
