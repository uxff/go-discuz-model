package models

type PrePortalArticleTitle struct {
	Aid          int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Catid        int    `xorm:"not null default 0 index(catid) MEDIUMINT(8)"`
	Bid          int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Uid          int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Username     string `xorm:"not null default '''' VARCHAR(255)"`
	Title        string `xorm:"not null default '''' VARCHAR(255)"`
	Highlight    string `xorm:"not null default '''' VARCHAR(255)"`
	Author       string `xorm:"not null default '''' VARCHAR(255)"`
	From         string `xorm:"not null default '''' VARCHAR(255)"`
	Fromurl      string `xorm:"not null default '''' VARCHAR(255)"`
	Url          string `xorm:"not null default '''' VARCHAR(255)"`
	Summary      string `xorm:"not null default '''' VARCHAR(255)"`
	Pic          string `xorm:"not null default '''' VARCHAR(255)"`
	Thumb        int    `xorm:"not null default 0 TINYINT(1)"`
	Remote       int    `xorm:"not null default 0 TINYINT(1)"`
	Id           int    `xorm:"not null default 0 INT(10)"`
	Idtype       string `xorm:"not null default '''' VARCHAR(255)"`
	Contents     int    `xorm:"not null default 0 SMALLINT(6)"`
	Allowcomment int    `xorm:"not null default 0 TINYINT(1)"`
	Owncomment   int    `xorm:"not null default 0 TINYINT(1)"`
	Click1       int    `xorm:"not null default 0 SMALLINT(6)"`
	Click2       int    `xorm:"not null default 0 SMALLINT(6)"`
	Click3       int    `xorm:"not null default 0 SMALLINT(6)"`
	Click4       int    `xorm:"not null default 0 SMALLINT(6)"`
	Click5       int    `xorm:"not null default 0 SMALLINT(6)"`
	Click6       int    `xorm:"not null default 0 SMALLINT(6)"`
	Click7       int    `xorm:"not null default 0 SMALLINT(6)"`
	Click8       int    `xorm:"not null default 0 SMALLINT(6)"`
	Tag          int    `xorm:"not null default 0 TINYINT(8)"`
	Dateline     int    `xorm:"not null default 0 index(catid) INT(10)"`
	Status       int    `xorm:"not null default 0 TINYINT(1)"`
	Showinnernav int    `xorm:"not null default 0 TINYINT(1)"`
	Preaid       int    `xorm:"not null MEDIUMINT(8)"`
	Nextaid      int    `xorm:"not null MEDIUMINT(8)"`
	Htmlmade     int    `xorm:"not null default 0 TINYINT(1)"`
	Htmlname     string `xorm:"not null default '''' VARCHAR(255)"`
	Htmldir      string `xorm:"not null default '''' VARCHAR(255)"`
}

func (t *PrePortalArticleTitle) TableName() string {
	return "pre_portal_article_title"
}
