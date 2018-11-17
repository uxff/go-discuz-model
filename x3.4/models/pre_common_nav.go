package models

type PreCommonNav struct {
	Id           int    `xorm:"not null pk autoincr SMALLINT(6)"`
	Parentid     int    `xorm:"not null default 0 SMALLINT(6)"`
	Name         string `xorm:"not null VARCHAR(255)"`
	Title        string `xorm:"not null VARCHAR(255)"`
	Url          string `xorm:"not null VARCHAR(255)"`
	Identifier   string `xorm:"not null VARCHAR(255)"`
	Target       int    `xorm:"not null default 0 TINYINT(1)"`
	Type         int    `xorm:"not null default 0 TINYINT(1)"`
	Available    int    `xorm:"not null default 0 TINYINT(1)"`
	Displayorder int    `xorm:"not null TINYINT(3)"`
	Highlight    int    `xorm:"not null default 0 TINYINT(1)"`
	Level        int    `xorm:"not null default 0 TINYINT(1)"`
	Subtype      int    `xorm:"not null default 0 TINYINT(1)"`
	Subcols      int    `xorm:"not null default 0 TINYINT(1)"`
	Icon         string `xorm:"not null VARCHAR(255)"`
	Subname      string `xorm:"not null VARCHAR(255)"`
	Suburl       string `xorm:"not null VARCHAR(255)"`
	Navtype      int    `xorm:"not null default 0 index TINYINT(1)"`
	Logo         string `xorm:"not null VARCHAR(255)"`
}
