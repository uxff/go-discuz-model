package models

type PreUcenterPmMembers struct {
	Plid         int `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Uid          int `xorm:"not null pk default 0 index(lastdateline) index(lastupdate) MEDIUMINT(8)"`
	Isnew        int `xorm:"not null default 0 index TINYINT(1)"`
	Pmnum        int `xorm:"not null default 0 INT(10)"`
	Lastupdate   int `xorm:"not null default 0 index(lastupdate) INT(10)"`
	Lastdateline int `xorm:"not null default 0 index(lastdateline) INT(10)"`
}

func (t *PreUcenterPmMembers) TableName() string {
	return "pre_ucenter_pm_members"
}
