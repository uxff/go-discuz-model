package models

type PreCommonMember struct {
	Uid                int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Email              string `xorm:"not null default '''' index CHAR(40)"`
	Username           string `xorm:"not null default '''' unique CHAR(15)"`
	Password           string `xorm:"not null default '''' CHAR(32)"`
	Status             int    `xorm:"not null default 0 TINYINT(1)"`
	Emailstatus        int    `xorm:"not null default 0 TINYINT(1)"`
	Avatarstatus       int    `xorm:"not null default 0 TINYINT(1)"`
	Videophotostatus   int    `xorm:"not null default 0 TINYINT(1)"`
	Adminid            int    `xorm:"not null default 0 TINYINT(1)"`
	Groupid            int    `xorm:"not null default 0 index SMALLINT(6)"`
	Groupexpiry        int    `xorm:"not null default 0 INT(10)"`
	Extgroupids        string `xorm:"not null default '''' CHAR(20)"`
	Regdate            int    `xorm:"not null default 0 index INT(10)"`
	Credits            int    `xorm:"not null default 0 INT(10)"`
	Notifysound        int    `xorm:"not null default 0 TINYINT(1)"`
	Timeoffset         string `xorm:"not null default '''' CHAR(4)"`
	Newpm              int    `xorm:"not null default 0 SMALLINT(6)"`
	Newprompt          int    `xorm:"not null default 0 SMALLINT(6)"`
	Accessmasks        int    `xorm:"not null default 0 TINYINT(1)"`
	Allowadmincp       int    `xorm:"not null default 0 TINYINT(1)"`
	Onlyacceptfriendpm int    `xorm:"not null default 0 TINYINT(1)"`
	Conisbind          int    `xorm:"not null default 0 index TINYINT(1)"`
	Freeze             int    `xorm:"not null default 0 TINYINT(1)"`
}

func (t *PreCommonMember) TableName() string {
	return "pre_common_member"
}
