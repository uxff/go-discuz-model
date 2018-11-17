package models

type PreCommonConnectGuest struct {
	Conopenid    string `xorm:"not null pk default '''' CHAR(32)"`
	Conuin       string `xorm:"not null default '''' CHAR(40)"`
	Conuinsecret string `xorm:"not null default '''' CHAR(16)"`
	Conqqnick    string `xorm:"not null default '''' CHAR(100)"`
	Conuintoken  string `xorm:"not null default '''' CHAR(32)"`
}

func (t *PreCommonConnectGuest) TableName() string {
	return "pre_common_connect_guest"
}
