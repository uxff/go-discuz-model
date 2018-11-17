package models

type PreMobileWsqThreadlist struct {
	Skey   int    `xorm:"not null pk INT(10)"`
	Svalue string `xorm:"not null TEXT"`
}
