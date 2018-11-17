package models

type PreForumGroupuser struct {
	Fid          int    `xorm:"not null pk default 0 index(userlist) MEDIUMINT(8)"`
	Uid          int    `xorm:"not null pk default 0 index(uid_lastupdate) MEDIUMINT(8)"`
	Username     string `xorm:"not null CHAR(15)"`
	Level        int    `xorm:"not null default 0 index(userlist) TINYINT(3)"`
	Threads      int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Replies      int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Joindateline int    `xorm:"not null default 0 INT(10)"`
	Lastupdate   int    `xorm:"not null default 0 index(uid_lastupdate) index(userlist) INT(10)"`
	Privacy      int    `xorm:"not null default 0 TINYINT(1)"`
}

func (t *PreForumGroupuser) TableName() string {
	return "pre_forum_groupuser"
}
