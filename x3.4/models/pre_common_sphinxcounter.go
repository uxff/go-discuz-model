package models

type PreCommonSphinxcounter struct {
	Indexid int `xorm:"not null pk TINYINT(1)"`
	Maxid   int `xorm:"not null INT(10)"`
}
