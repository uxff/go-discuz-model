package models

type PreForumMedal struct {
	Medalid      int    `xorm:"not null pk autoincr SMALLINT(6)"`
	Name         string `xorm:"not null default '''' VARCHAR(50)"`
	Available    int    `xorm:"not null default 0 index(available) TINYINT(1)"`
	Image        string `xorm:"not null default '''' VARCHAR(255)"`
	Type         int    `xorm:"not null default 0 TINYINT(1)"`
	Displayorder int    `xorm:"not null default 0 index(available) index TINYINT(3)"`
	Description  string `xorm:"not null VARCHAR(255)"`
	Expiration   int    `xorm:"not null default 0 SMALLINT(6)"`
	Permission   string `xorm:"not null MEDIUMTEXT"`
	Credit       int    `xorm:"not null default 0 TINYINT(1)"`
	Price        int    `xorm:"not null default 0 MEDIUMINT(8)"`
}

func (t *PreForumMedal) TableName() string {
	return "pre_forum_medal"
}
