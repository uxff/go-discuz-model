package models

type PreUcenterFriends struct {
	Uid       int    `xorm:"not null default 0 index MEDIUMINT(8)"`
	Friendid  int    `xorm:"not null default 0 index MEDIUMINT(8)"`
	Direction int    `xorm:"not null default 0 TINYINT(1)"`
	Version   int    `xorm:"not null pk autoincr INT(10)"`
	Delstatus int    `xorm:"not null default 0 TINYINT(1)"`
	Comment   string `xorm:"not null default '''' CHAR(255)"`
}

func (t *PreUcenterFriends) TableName() string {
	return "pre_ucenter_friends"
}
