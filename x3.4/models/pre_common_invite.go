package models

type PreCommonInvite struct {
	Id          int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Uid         int    `xorm:"not null default 0 index(uid) MEDIUMINT(8)"`
	Code        string `xorm:"not null default '''' CHAR(20)"`
	Fuid        int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Fusername   string `xorm:"not null default '''' CHAR(20)"`
	Type        int    `xorm:"not null default 0 TINYINT(1)"`
	Email       string `xorm:"not null default '''' CHAR(40)"`
	Inviteip    string `xorm:"not null default '''' CHAR(15)"`
	Appid       int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Dateline    int    `xorm:"not null default 0 index(uid) INT(10)"`
	Endtime     int    `xorm:"not null default 0 INT(10)"`
	Regdateline int    `xorm:"not null default 0 INT(10)"`
	Status      int    `xorm:"not null default 1 TINYINT(1)"`
	Orderid     string `xorm:"not null default '''' CHAR(32)"`
}

func (t *PreCommonInvite) TableName() string {
	return "pre_common_invite"
}
