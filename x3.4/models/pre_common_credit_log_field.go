package models

type PreCommonCreditLogField struct {
	Logid int    `xorm:"not null index MEDIUMINT(8)"`
	Title string `xorm:"not null VARCHAR(100)"`
	Text  string `xorm:"not null TEXT"`
}
