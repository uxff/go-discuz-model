package models

type PreForumThreadrush struct {
	Tid           int    `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Stopfloor     int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Starttimefrom int    `xorm:"not null default 0 INT(10)"`
	Starttimeto   int    `xorm:"not null default 0 INT(10)"`
	Rewardfloor   string `xorm:"not null TEXT"`
	Creditlimit   int    `xorm:"not null default -996 INT(10)"`
	Replylimit    int    `xorm:"not null default 0 SMALLINT(6)"`
}
