package models

type PreForumTypeoption struct {
	Optionid     int    `xorm:"not null pk autoincr SMALLINT(6)"`
	Classid      int    `xorm:"not null default 0 index SMALLINT(6)"`
	Displayorder int    `xorm:"not null default 0 TINYINT(3)"`
	Expiration   int    `xorm:"not null TINYINT(1)"`
	Protect      string `xorm:"not null VARCHAR(255)"`
	Title        string `xorm:"not null default '''' VARCHAR(255)"`
	Description  string `xorm:"not null default '''' VARCHAR(255)"`
	Identifier   string `xorm:"not null default '''' VARCHAR(255)"`
	Type         string `xorm:"not null default '''' VARCHAR(255)"`
	Unit         string `xorm:"not null VARCHAR(255)"`
	Rules        string `xorm:"not null MEDIUMTEXT"`
	Permprompt   string `xorm:"not null MEDIUMTEXT"`
}

func (t *PreForumTypeoption) TableName() string {
	return "pre_forum_typeoption"
}
