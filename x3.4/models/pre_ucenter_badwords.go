package models

type PreUcenterBadwords struct {
	Id          int    `xorm:"not null pk autoincr SMALLINT(6)"`
	Admin       string `xorm:"not null default '''' VARCHAR(15)"`
	Find        string `xorm:"not null default '''' unique VARCHAR(255)"`
	Replacement string `xorm:"not null default '''' VARCHAR(255)"`
	Findpattern string `xorm:"not null default '''' VARCHAR(255)"`
}
