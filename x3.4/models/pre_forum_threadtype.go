package models

type PreForumThreadtype struct {
	Typeid       int    `xorm:"not null pk autoincr SMALLINT(6)"`
	Fid          int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Displayorder int    `xorm:"not null default 0 SMALLINT(6)"`
	Name         string `xorm:"not null default '''' VARCHAR(255)"`
	Description  string `xorm:"not null default '''' VARCHAR(255)"`
	Icon         string `xorm:"not null default '''' VARCHAR(255)"`
	Special      int    `xorm:"not null default 0 SMALLINT(6)"`
	Modelid      int    `xorm:"not null default 0 SMALLINT(6)"`
	Expiration   int    `xorm:"not null default 0 TINYINT(1)"`
	Template     string `xorm:"not null TEXT"`
	Stemplate    string `xorm:"not null TEXT"`
	Ptemplate    string `xorm:"not null TEXT"`
	Btemplate    string `xorm:"not null TEXT"`
}

func (t *PreForumThreadtype) TableName() string {
	return "pre_forum_threadtype"
}
