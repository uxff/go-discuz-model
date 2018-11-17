package models

type PrePortalArticleRelated struct {
	Aid          int `xorm:"not null pk autoincr index(aid) MEDIUMINT(8)"`
	Raid         int `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Displayorder int `xorm:"not null default 0 index(aid) MEDIUMINT(8)"`
}

func (t *PrePortalArticleRelated) TableName() string {
	return "pre_portal_article_related"
}
