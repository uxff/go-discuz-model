package models

type PreForumCollectionfollow struct {
	Uid       int    `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Username  string `xorm:"not null default '''' CHAR(15)"`
	Ctid      int    `xorm:"not null pk default 0 index(ctid) MEDIUMINT(8)"`
	Dateline  int    `xorm:"not null default 0 index(ctid) INT(10)"`
	Lastvisit int    `xorm:"not null default 0 INT(10)"`
}

func (t *PreForumCollectionfollow) TableName() string {
	return "pre_forum_collectionfollow"
}
