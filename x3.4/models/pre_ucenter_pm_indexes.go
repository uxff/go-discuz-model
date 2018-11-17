package models

type PreUcenterPmIndexes struct {
	Pmid int `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Plid int `xorm:"not null default 0 index MEDIUMINT(8)"`
}

func (t *PreUcenterPmIndexes) TableName() string {
	return "pre_ucenter_pm_indexes"
}
