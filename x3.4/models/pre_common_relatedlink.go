package models

type PreCommonRelatedlink struct {
	Id     int    `xorm:"not null pk autoincr SMALLINT(6)"`
	Name   string `xorm:"not null default '''' VARCHAR(100)"`
	Url    string `xorm:"not null default '''' VARCHAR(255)"`
	Extent int    `xorm:"not null default 0 TINYINT(3)"`
}

func (t *PreCommonRelatedlink) TableName() string {
	return "pre_common_relatedlink"
}
