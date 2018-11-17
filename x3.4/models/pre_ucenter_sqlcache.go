package models

type PreUcenterSqlcache struct {
	Sqlid  string `xorm:"not null pk default '''' CHAR(6)"`
	Data   string `xorm:"not null CHAR(100)"`
	Expiry int    `xorm:"not null index INT(10)"`
}

func (t *PreUcenterSqlcache) TableName() string {
	return "pre_ucenter_sqlcache"
}
