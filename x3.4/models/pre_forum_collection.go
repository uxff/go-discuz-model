package models

type PreForumCollection struct {
	Ctid         int     `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Uid          int     `xorm:"not null default 0 index MEDIUMINT(8)"`
	Username     string  `xorm:"not null default '''' VARCHAR(15)"`
	Name         string  `xorm:"not null default '''' VARCHAR(50)"`
	Dateline     int     `xorm:"not null default 0 index INT(10)"`
	Follownum    int     `xorm:"not null default 0 index MEDIUMINT(8)"`
	Threadnum    int     `xorm:"not null default 0 index(hotcollection) MEDIUMINT(8)"`
	Commentnum   int     `xorm:"not null default 0 MEDIUMINT(8)"`
	Desc         string  `xorm:"not null default '''' VARCHAR(255)"`
	Lastupdate   int     `xorm:"not null default 0 index(hotcollection) INT(10)"`
	Rate         float32 `xorm:"not null default 0 FLOAT"`
	Ratenum      int     `xorm:"not null default 0 MEDIUMINT(8)"`
	Lastpost     int     `xorm:"not null default 0 MEDIUMINT(8)"`
	Lastsubject  string  `xorm:"not null default '''' VARCHAR(80)"`
	Lastposttime int     `xorm:"not null default 0 INT(10)"`
	Lastposter   string  `xorm:"not null default '''' VARCHAR(15)"`
	Lastvisit    int     `xorm:"not null default 0 INT(10)"`
	Keyword      string  `xorm:"not null default '''' VARCHAR(255)"`
}

func (t *PreForumCollection) TableName() string {
	return "pre_forum_collection"
}
