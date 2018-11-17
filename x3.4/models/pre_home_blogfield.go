package models

type PreHomeBlogfield struct {
	Blogid      int    `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Uid         int    `xorm:"not null default 0 index MEDIUMINT(8)"`
	Pic         string `xorm:"not null default '''' VARCHAR(255)"`
	Tag         string `xorm:"not null default '''' VARCHAR(255)"`
	Message     string `xorm:"not null MEDIUMTEXT"`
	Postip      string `xorm:"not null default '''' VARCHAR(255)"`
	Port        int    `xorm:"not null default 0 SMALLINT(6)"`
	Related     string `xorm:"not null TEXT"`
	Relatedtime int    `xorm:"not null default 0 INT(10)"`
	TargetIds   string `xorm:"not null TEXT"`
	Hotuser     string `xorm:"not null TEXT"`
	Magiccolor  int    `xorm:"not null default 0 TINYINT(6)"`
	Magicpaper  int    `xorm:"not null default 0 TINYINT(6)"`
	Pushedaid   int    `xorm:"not null default 0 MEDIUMINT(8)"`
}

func (t *PreHomeBlogfield) TableName() string {
	return "pre_home_blogfield"
}
