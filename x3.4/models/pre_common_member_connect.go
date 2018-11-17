package models

type PreCommonMemberConnect struct {
	Uid              int    `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Conuin           string `xorm:"not null default '''' index CHAR(40)"`
	Conuinsecret     string `xorm:"not null default '''' CHAR(16)"`
	Conopenid        string `xorm:"not null default '''' index CHAR(32)"`
	Conisfeed        int    `xorm:"not null default 0 TINYINT(1)"`
	Conispublishfeed int    `xorm:"not null default 0 TINYINT(1)"`
	Conispublisht    int    `xorm:"not null default 0 TINYINT(1)"`
	Conisregister    int    `xorm:"not null default 0 TINYINT(1)"`
	Conisqzoneavatar int    `xorm:"not null default 0 TINYINT(1)"`
	Conisqqshow      int    `xorm:"not null default 0 TINYINT(1)"`
	Conuintoken      string `xorm:"not null default '''' CHAR(32)"`
}
