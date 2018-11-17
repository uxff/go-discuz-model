package models

type PreHomeAppcreditlog struct {
	Logid    int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Uid      int    `xorm:"not null default 0 index(uid) MEDIUMINT(8)"`
	Appid    int    `xorm:"not null default 0 index MEDIUMINT(8)"`
	Appname  string `xorm:"not null default '''' VARCHAR(60)"`
	Type     int    `xorm:"not null default 0 TINYINT(1)"`
	Credit   int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Note     string `xorm:"not null TEXT"`
	Dateline int    `xorm:"not null default 0 index(uid) INT(10)"`
}

func (t *PreHomeAppcreditlog) TableName() string {
	return "pre_home_appcreditlog"
}
