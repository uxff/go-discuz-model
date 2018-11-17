package models

type PrePortalTopic struct {
	Topicid       int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Title         string `xorm:"not null default '''' VARCHAR(255)"`
	Name          string `xorm:"not null default '''' index VARCHAR(255)"`
	Domain        string `xorm:"not null default '''' VARCHAR(255)"`
	Summary       string `xorm:"not null TEXT"`
	Keyword       string `xorm:"not null TEXT"`
	Cover         string `xorm:"not null default '''' VARCHAR(255)"`
	Picflag       int    `xorm:"not null default 0 TINYINT(1)"`
	Primaltplname string `xorm:"not null default '''' VARCHAR(255)"`
	Useheader     int    `xorm:"not null default 0 TINYINT(1)"`
	Usefooter     int    `xorm:"not null default 0 TINYINT(1)"`
	Uid           int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Username      string `xorm:"not null default '''' VARCHAR(255)"`
	Viewnum       int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Dateline      int    `xorm:"not null default 0 INT(10)"`
	Closed        int    `xorm:"not null default 0 TINYINT(1)"`
	Allowcomment  int    `xorm:"not null default 0 TINYINT(1)"`
	Commentnum    int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Htmlmade      int    `xorm:"not null default 0 TINYINT(1)"`
	Htmldir       string `xorm:"not null default '''' VARCHAR(255)"`
}

func (t *PrePortalTopic) TableName() string {
	return "pre_portal_topic"
}
