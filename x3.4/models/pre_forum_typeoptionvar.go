package models

type PreForumTypeoptionvar struct {
	Sortid     int    `xorm:"not null default 0 index SMALLINT(6)"`
	Tid        int    `xorm:"not null default 0 index MEDIUMINT(8)"`
	Fid        int    `xorm:"not null default 0 index MEDIUMINT(8)"`
	Optionid   int    `xorm:"not null default 0 SMALLINT(6)"`
	Expiration int    `xorm:"not null default 0 INT(10)"`
	Value      string `xorm:"not null MEDIUMTEXT"`
}
