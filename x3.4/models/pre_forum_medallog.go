package models

type PreForumMedallog struct {
	Id         int `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Uid        int `xorm:"not null default 0 index(uid) MEDIUMINT(8)"`
	Medalid    int `xorm:"not null default 0 index(uid) SMALLINT(6)"`
	Type       int `xorm:"not null default 0 index index(uid) TINYINT(1)"`
	Dateline   int `xorm:"not null default 0 index INT(10)"`
	Expiration int `xorm:"not null default 0 index(status) INT(10)"`
	Status     int `xorm:"not null default 0 index(status) TINYINT(1)"`
}

func (t *PreForumMedallog) TableName() string {
	return "pre_forum_medallog"
}
