package models

type PrePortalArticleCount struct {
	Aid        int `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Catid      int `xorm:"not null default 0 MEDIUMINT(8)"`
	Viewnum    int `xorm:"not null default 0 MEDIUMINT(8)"`
	Commentnum int `xorm:"not null default 0 MEDIUMINT(8)"`
	Favtimes   int `xorm:"not null default 0 MEDIUMINT(8)"`
	Sharetimes int `xorm:"not null default 0 MEDIUMINT(8)"`
}

func (t *PrePortalArticleCount) TableName() string {
	return "pre_portal_article_count"
}
