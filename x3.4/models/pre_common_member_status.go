package models

type PreCommonMemberStatus struct {
	Uid             int    `xorm:"not null pk MEDIUMINT(8)"`
	Regip           string `xorm:"not null default '''' CHAR(15)"`
	Lastip          string `xorm:"not null default '''' CHAR(15)"`
	Port            int    `xorm:"not null default 0 SMALLINT(6)"`
	Lastvisit       int    `xorm:"not null default 0 INT(10)"`
	Lastactivity    int    `xorm:"not null default 0 index(lastactivity) INT(10)"`
	Lastpost        int    `xorm:"not null default 0 INT(10)"`
	Lastsendmail    int    `xorm:"not null default 0 INT(10)"`
	Invisible       int    `xorm:"not null default 0 index(lastactivity) TINYINT(1)"`
	Buyercredit     int    `xorm:"not null default 0 SMALLINT(6)"`
	Sellercredit    int    `xorm:"not null default 0 SMALLINT(6)"`
	Favtimes        int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Sharetimes      int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Profileprogress int    `xorm:"not null default 0 TINYINT(2)"`
}

func (t *PreCommonMemberStatus) TableName() string {
	return "pre_common_member_status"
}
