package models

type PrePortalArticleTrash struct {
	Aid     int    `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Content string `xorm:"not null TEXT"`
}
