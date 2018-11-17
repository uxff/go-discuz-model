package models

type PreMobileWechatMasssend struct {
	Id             int    `xorm:"not null pk autoincr INT(10)"`
	Type           string `xorm:"not null CHAR(5)"`
	Name           string `xorm:"not null VARCHAR(255)"`
	ResourceId     int    `xorm:"not null INT(10)"`
	GroupId        int    `xorm:"not null INT(10)"`
	Text           string `xorm:"default 'NULL' TEXT"`
	MediaId        string `xorm:"default '''' CHAR(64)"`
	CreatedAt      int    `xorm:"not null INT(10)"`
	SentAt         int    `xorm:"default NULL INT(10)"`
	MsgId          int    `xorm:"default NULL INT(10)"`
	ResStatus      string `xorm:"default 'NULL' VARCHAR(50)"`
	ResTotalcount  int    `xorm:"default NULL INT(10)"`
	ResFiltercount int    `xorm:"default NULL INT(10)"`
	ResSentcount   int    `xorm:"default NULL INT(10)"`
	ResErrorcount  int    `xorm:"default NULL INT(10)"`
	ResFinishAt    int    `xorm:"default NULL INT(10)"`
}
