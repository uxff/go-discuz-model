package models

type PreForumCollectioninvite struct {
	Ctid     int `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Uid      int `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Dateline int `xorm:"not null default 0 index INT(10)"`
}

func (t *PreForumCollectioninvite) TableName() string {
	return "pre_forum_collectioninvite"
}
