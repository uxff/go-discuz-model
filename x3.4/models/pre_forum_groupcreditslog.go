package models

type PreForumGroupcreditslog struct {
	Fid     int `xorm:"not null pk MEDIUMINT(8)"`
	Uid     int `xorm:"not null pk MEDIUMINT(8)"`
	Logdate int `xorm:"not null pk default 0 INT(8)"`
}

func (t *PreForumGroupcreditslog) TableName() string {
	return "pre_forum_groupcreditslog"
}
