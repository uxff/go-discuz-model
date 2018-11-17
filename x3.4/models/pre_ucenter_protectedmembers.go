package models

type PreUcenterProtectedmembers struct {
	Uid      int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Username string `xorm:"not null pk default '''' unique(username) CHAR(15)"`
	Appid    int    `xorm:"not null pk default 0 unique(username) TINYINT(1)"`
	Dateline int    `xorm:"not null default 0 INT(10)"`
	Admin    string `xorm:"not null default ''0'' CHAR(15)"`
}

func (t *PreUcenterProtectedmembers) TableName() string {
	return "pre_ucenter_protectedmembers"
}
