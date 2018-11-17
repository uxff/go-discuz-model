package models

type PrePortalComment struct {
	Cid      int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Uid      int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Username string `xorm:"not null default '''' VARCHAR(255)"`
	Id       int    `xorm:"not null default 0 index(idtype) MEDIUMINT(8)"`
	Idtype   string `xorm:"not null default '''' index(idtype) VARCHAR(20)"`
	Postip   string `xorm:"not null default '''' VARCHAR(255)"`
	Port     int    `xorm:"not null default 0 SMALLINT(6)"`
	Dateline int    `xorm:"not null default 0 index(idtype) INT(10)"`
	Status   int    `xorm:"not null default 0 TINYINT(1)"`
	Message  string `xorm:"not null TEXT"`
}

func (t *PrePortalComment) TableName() string {
	return "pre_portal_comment"
}
