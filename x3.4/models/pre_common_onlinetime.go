package models

type PreCommonOnlinetime struct {
	Uid        int `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Thismonth  int `xorm:"not null default 0 SMALLINT(6)"`
	Total      int `xorm:"not null default 0 MEDIUMINT(8)"`
	Lastupdate int `xorm:"not null default 0 INT(10)"`
}

func (t *PreCommonOnlinetime) TableName() string {
	return "pre_common_onlinetime"
}
