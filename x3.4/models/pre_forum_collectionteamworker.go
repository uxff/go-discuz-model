package models

type PreForumCollectionteamworker struct {
	Ctid      int    `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Uid       int    `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Name      string `xorm:"not null default '''' VARCHAR(50)"`
	Username  string `xorm:"not null default '''' VARCHAR(15)"`
	Lastvisit int    `xorm:"not null default 0 INT(8)"`
}

func (t *PreForumCollectionteamworker) TableName() string {
	return "pre_forum_collectionteamworker"
}
