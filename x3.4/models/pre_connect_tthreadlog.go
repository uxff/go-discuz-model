package models

type PreConnectTthreadlog struct {
	Twid       string `xorm:"not null pk CHAR(16)"`
	Tid        int    `xorm:"not null default 0 index(nexttime) index(updatetime) MEDIUMINT(8)"`
	Conopenid  string `xorm:"not null CHAR(32)"`
	Pagetime   int    `xorm:"default 0 INT(10)"`
	Lasttwid   string `xorm:"default 'NULL' CHAR(16)"`
	Nexttime   int    `xorm:"default 0 index(nexttime) INT(10)"`
	Updatetime int    `xorm:"default 0 index(updatetime) INT(10)"`
	Dateline   int    `xorm:"default 0 INT(10)"`
}

func (t *PreConnectTthreadlog) TableName() string {
	return "pre_connect_tthreadlog"
}
