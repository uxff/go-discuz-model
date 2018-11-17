package models

type PreForumCollectionrelated struct {
	Tid        int    `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Collection string `xorm:"not null TEXT"`
}
