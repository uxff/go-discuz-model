package models

type PrePortalArticleContent struct {
	Cid       int    `xorm:"not null pk autoincr INT(10)"`
	Aid       int    `xorm:"not null default 0 index(aid) MEDIUMINT(8)"`
	Id        int    `xorm:"not null default 0 INT(10)"`
	Idtype    string `xorm:"not null default '''' VARCHAR(255)"`
	Title     string `xorm:"not null default '''' VARCHAR(255)"`
	Content   string `xorm:"not null TEXT"`
	Pageorder int    `xorm:"not null default 0 index(aid) index SMALLINT(6)"`
	Dateline  int    `xorm:"not null default 0 INT(10)"`
}

func (t *PrePortalArticleContent) TableName() string {
	return "pre_portal_article_content"
}
