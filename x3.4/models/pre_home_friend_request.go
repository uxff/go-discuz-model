package models

type PreHomeFriendRequest struct {
	Uid       int    `xorm:"not null pk default 0 index(dateline) MEDIUMINT(8)"`
	Fuid      int    `xorm:"not null pk default 0 index MEDIUMINT(8)"`
	Fusername string `xorm:"not null default '''' CHAR(15)"`
	Gid       int    `xorm:"not null default 0 SMALLINT(6)"`
	Note      string `xorm:"not null default '''' CHAR(60)"`
	Dateline  int    `xorm:"not null default 0 index(dateline) INT(10)"`
}

func (t *PreHomeFriendRequest) TableName() string {
	return "pre_home_friend_request"
}
