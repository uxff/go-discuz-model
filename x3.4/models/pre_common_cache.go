package models

type PreCommonCache struct {
	Cachekey   string `xorm:"not null pk default '''' VARCHAR(255)"`
	Cachevalue []byte `xorm:"not null MEDIUMBLOB"`
	Dateline   int    `xorm:"not null default 0 INT(10)"`
}
