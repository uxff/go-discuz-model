package models

type PreForumReplycredit struct {
	Tid            int `xorm:"not null pk MEDIUMINT(6)"`
	Extcredits     int `xorm:"not null default 0 MEDIUMINT(6)"`
	Extcreditstype int `xorm:"not null default 0 TINYINT(1)"`
	Times          int `xorm:"not null default 0 SMALLINT(6)"`
	Membertimes    int `xorm:"not null default 0 SMALLINT(6)"`
	Random         int `xorm:"not null default 0 TINYINT(1)"`
}
