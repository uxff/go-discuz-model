package models

type PreCommonBlockXml struct {
	Id       int    `xorm:"not null pk autoincr SMALLINT(5)"`
	Name     string `xorm:"not null VARCHAR(255)"`
	Version  string `xorm:"not null VARCHAR(255)"`
	Url      string `xorm:"not null VARCHAR(255)"`
	Clientid string `xorm:"not null VARCHAR(255)"`
	Key      string `xorm:"not null VARCHAR(255)"`
	Signtype string `xorm:"not null VARCHAR(255)"`
	Data     string `xorm:"not null TEXT"`
}
