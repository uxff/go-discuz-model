package models

type PreForumActivity struct {
	Tid           int    `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Uid           int    `xorm:"not null default 0 index(uid) MEDIUMINT(8)"`
	Aid           int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Cost          int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Starttimefrom int    `xorm:"not null default 0 index index(uid) INT(10)"`
	Starttimeto   int    `xorm:"not null default 0 INT(10)"`
	Place         string `xorm:"not null default '''' VARCHAR(255)"`
	Class         string `xorm:"not null default '''' VARCHAR(255)"`
	Gender        int    `xorm:"not null default 0 TINYINT(1)"`
	Number        int    `xorm:"not null default 0 SMALLINT(5)"`
	Applynumber   int    `xorm:"not null default 0 index SMALLINT(5)"`
	Expiration    int    `xorm:"not null default 0 index INT(10)"`
	Ufield        string `xorm:"not null TEXT"`
	Credit        int    `xorm:"not null default 0 SMALLINT(6)"`
}

func (t *PreForumActivity) TableName() string {
	return "pre_forum_activity"
}
