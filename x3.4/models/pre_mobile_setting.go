package models

type PreMobileSetting struct {
	Skey   string `xorm:"not null pk default '''' VARCHAR(255)"`
	Svalue string `xorm:"not null TEXT"`
}
