package models

type PreHomeFollowFeedArchiver struct {
	Feedid   int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Uid      int    `xorm:"not null default 0 index(uid) MEDIUMINT(8)"`
	Username string `xorm:"not null default '''' VARCHAR(15)"`
	Tid      int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Note     string `xorm:"not null TEXT"`
	Dateline int    `xorm:"not null default 0 index(uid) INT(10)"`
}

func (t *PreHomeFollowFeedArchiver) TableName() string {
	return "pre_home_follow_feed_archiver"
}
