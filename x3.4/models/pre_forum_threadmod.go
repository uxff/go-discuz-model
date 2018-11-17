package models

type PreForumThreadmod struct {
	Tid        int    `xorm:"not null default 0 index(tid) MEDIUMINT(8)"`
	Uid        int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Username   string `xorm:"not null default '''' CHAR(15)"`
	Dateline   int    `xorm:"not null default 0 index(tid) INT(10)"`
	Expiration int    `xorm:"not null default 0 index(expiration) INT(10)"`
	Action     string `xorm:"not null default '''' CHAR(5)"`
	Status     int    `xorm:"not null default 0 index(expiration) TINYINT(1)"`
	Magicid    int    `xorm:"not null SMALLINT(6)"`
	Stamp      int    `xorm:"not null TINYINT(3)"`
	Reason     string `xorm:"not null default '''' CHAR(40)"`
}

func (t *PreForumThreadmod) TableName() string {
	return "pre_forum_threadmod"
}
