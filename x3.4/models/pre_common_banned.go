package models

type PreCommonBanned struct {
	Id         int    `xorm:"not null pk autoincr SMALLINT(6)"`
	Ip1        int    `xorm:"not null default 0 SMALLINT(3)"`
	Ip2        int    `xorm:"not null default 0 SMALLINT(3)"`
	Ip3        int    `xorm:"not null default 0 SMALLINT(3)"`
	Ip4        int    `xorm:"not null default 0 SMALLINT(3)"`
	Admin      string `xorm:"not null default '''' VARCHAR(15)"`
	Dateline   int    `xorm:"not null default 0 INT(10)"`
	Expiration int    `xorm:"not null default 0 INT(10)"`
}

func (t *PreCommonBanned) TableName() string {
	return "pre_common_banned"
}
