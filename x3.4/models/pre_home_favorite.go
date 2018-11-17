package models

type PreHomeFavorite struct {
	Favid       int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Uid         int    `xorm:"not null default 0 index(uid) MEDIUMINT(8)"`
	Id          int    `xorm:"not null default 0 index(idtype) MEDIUMINT(8)"`
	Idtype      string `xorm:"not null default '''' index(idtype) index(uid) VARCHAR(255)"`
	Spaceuid    int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Title       string `xorm:"not null default '''' VARCHAR(255)"`
	Description string `xorm:"not null TEXT"`
	Dateline    int    `xorm:"not null default 0 index(uid) INT(10)"`
}
