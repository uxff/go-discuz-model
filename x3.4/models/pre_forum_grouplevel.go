package models

type PreForumGrouplevel struct {
	Levelid       int    `xorm:"not null pk autoincr SMALLINT(6)"`
	Type          string `xorm:"not null default ''default'' ENUM('default','special')"`
	Leveltitle    string `xorm:"not null default '''' VARCHAR(255)"`
	Creditshigher int    `xorm:"not null default 0 index(creditsrange) INT(10)"`
	Creditslower  int    `xorm:"not null default 0 index(creditsrange) INT(10)"`
	Icon          string `xorm:"not null default '''' VARCHAR(255)"`
	Creditspolicy string `xorm:"not null TEXT"`
	Postpolicy    string `xorm:"not null TEXT"`
	Specialswitch string `xorm:"not null TEXT"`
}

func (t *PreForumGrouplevel) TableName() string {
	return "pre_forum_grouplevel"
}
