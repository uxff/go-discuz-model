package models

type PreUcenterPmMessages5 struct {
	Pmid      int    `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Plid      int    `xorm:"not null default 0 index(dateline) index(plid) MEDIUMINT(8)"`
	Authorid  int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Message   string `xorm:"not null TEXT"`
	Delstatus int    `xorm:"not null default 0 index(plid) TINYINT(1)"`
	Dateline  int    `xorm:"not null default 0 index(dateline) index(plid) INT(10)"`
}
