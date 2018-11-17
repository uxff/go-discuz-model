package models

type PreCommonAdminnote struct {
	Id         int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Admin      string `xorm:"not null default '''' VARCHAR(15)"`
	Access     int    `xorm:"not null default 0 TINYINT(3)"`
	Adminid    int    `xorm:"not null default 0 TINYINT(3)"`
	Dateline   int    `xorm:"not null default 0 INT(10)"`
	Expiration int    `xorm:"not null default 0 INT(10)"`
	Message    string `xorm:"not null TEXT"`
}

func (t *PreCommonAdminnote) TableName() string {
	return "pre_common_adminnote"
}
