package models

type PrePortalCategoryPermission struct {
	Catid          int `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Uid            int `xorm:"not null pk default 0 index MEDIUMINT(8)"`
	Allowpublish   int `xorm:"not null default 0 TINYINT(1)"`
	Allowmanage    int `xorm:"not null default 0 TINYINT(1)"`
	Inheritedcatid int `xorm:"not null default 0 MEDIUMINT(8)"`
}

func (t *PrePortalCategoryPermission) TableName() string {
	return "pre_portal_category_permission"
}
