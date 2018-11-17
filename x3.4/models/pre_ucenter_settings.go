package models

type PreUcenterSettings struct {
	K string `xorm:"not null pk default '''' VARCHAR(32)"`
	V string `xorm:"not null TEXT"`
}
