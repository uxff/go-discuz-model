package models

type PreCommonMagiclog struct {
	Uid       int    `xorm:"not null default 0 index(uid) MEDIUMINT(8)"`
	Magicid   int    `xorm:"not null default 0 index(magicid) SMALLINT(6)"`
	Action    int    `xorm:"not null default 0 index TINYINT(1)"`
	Dateline  int    `xorm:"not null default 0 index(magicid) index(targetuid) index(uid) INT(10)"`
	Amount    int    `xorm:"not null default 0 SMALLINT(6)"`
	Credit    int    `xorm:"not null default 0 TINYINT(1)"`
	Price     int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Targetid  int    `xorm:"not null default 0 INT(10)"`
	Idtype    string `xorm:"default 'NULL' CHAR(6)"`
	Targetuid int    `xorm:"not null default 0 index(targetuid) MEDIUMINT(8)"`
}

func (t *PreCommonMagiclog) TableName() string {
	return "pre_common_magiclog"
}
