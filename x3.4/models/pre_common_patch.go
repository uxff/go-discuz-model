package models

type PreCommonPatch struct {
	Serial   string `xorm:"not null pk default '''' VARCHAR(10)"`
	Rule     string `xorm:"not null TEXT"`
	Note     string `xorm:"not null TEXT"`
	Status   int    `xorm:"not null default 0 TINYINT(1)"`
	Dateline int    `xorm:"not null default 0 INT(10)"`
}

func (t *PreCommonPatch) TableName() string {
	return "pre_common_patch"
}
