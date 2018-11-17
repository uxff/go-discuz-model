package models

type PreForumCollectioncomment struct {
	Cid      int     `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Ctid     int     `xorm:"not null default 0 index(ctid) index(userrate) MEDIUMINT(8)"`
	Uid      int     `xorm:"not null default 0 index(userrate) MEDIUMINT(8)"`
	Username string  `xorm:"not null default '''' VARCHAR(15)"`
	Message  string  `xorm:"not null TEXT"`
	Dateline int     `xorm:"not null default 0 index(ctid) INT(10)"`
	Useip    string  `xorm:"not null default '''' VARCHAR(16)"`
	Port     int     `xorm:"not null default 0 SMALLINT(6)"`
	Rate     float32 `xorm:"not null default 0 index(userrate) FLOAT"`
}
