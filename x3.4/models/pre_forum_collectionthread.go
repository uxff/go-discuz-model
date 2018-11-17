package models

type PreForumCollectionthread struct {
	Ctid     int    `xorm:"not null pk default 0 index(ctid) MEDIUMINT(8)"`
	Tid      int    `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Dateline int    `xorm:"not null default 0 index(ctid) INT(10)"`
	Reason   string `xorm:"not null default '''' VARCHAR(255)"`
}

func (t *PreForumCollectionthread) TableName() string {
	return "pre_forum_collectionthread"
}
