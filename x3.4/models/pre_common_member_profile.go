package models

type PreCommonMemberProfile struct {
	Uid             int    `xorm:"not null pk MEDIUMINT(8)"`
	Realname        string `xorm:"not null default '''' VARCHAR(255)"`
	Gender          int    `xorm:"not null default 0 TINYINT(1)"`
	Birthyear       int    `xorm:"not null default 0 SMALLINT(6)"`
	Birthmonth      int    `xorm:"not null default 0 TINYINT(3)"`
	Birthday        int    `xorm:"not null default 0 TINYINT(3)"`
	Constellation   string `xorm:"not null default '''' VARCHAR(255)"`
	Zodiac          string `xorm:"not null default '''' VARCHAR(255)"`
	Telephone       string `xorm:"not null default '''' VARCHAR(255)"`
	Mobile          string `xorm:"not null default '''' VARCHAR(255)"`
	Idcardtype      string `xorm:"not null default '''' VARCHAR(255)"`
	Idcard          string `xorm:"not null default '''' VARCHAR(255)"`
	Address         string `xorm:"not null default '''' VARCHAR(255)"`
	Zipcode         string `xorm:"not null default '''' VARCHAR(255)"`
	Nationality     string `xorm:"not null default '''' VARCHAR(255)"`
	Birthprovince   string `xorm:"not null default '''' VARCHAR(255)"`
	Birthcity       string `xorm:"not null default '''' VARCHAR(255)"`
	Birthdist       string `xorm:"not null default '''' VARCHAR(20)"`
	Birthcommunity  string `xorm:"not null default '''' VARCHAR(255)"`
	Resideprovince  string `xorm:"not null default '''' VARCHAR(255)"`
	Residecity      string `xorm:"not null default '''' VARCHAR(255)"`
	Residedist      string `xorm:"not null default '''' VARCHAR(20)"`
	Residecommunity string `xorm:"not null default '''' VARCHAR(255)"`
	Residesuite     string `xorm:"not null default '''' VARCHAR(255)"`
	Graduateschool  string `xorm:"not null default '''' VARCHAR(255)"`
	Company         string `xorm:"not null default '''' VARCHAR(255)"`
	Education       string `xorm:"not null default '''' VARCHAR(255)"`
	Occupation      string `xorm:"not null default '''' VARCHAR(255)"`
	Position        string `xorm:"not null default '''' VARCHAR(255)"`
	Revenue         string `xorm:"not null default '''' VARCHAR(255)"`
	Affectivestatus string `xorm:"not null default '''' VARCHAR(255)"`
	Lookingfor      string `xorm:"not null default '''' VARCHAR(255)"`
	Bloodtype       string `xorm:"not null default '''' VARCHAR(255)"`
	Height          string `xorm:"not null default '''' VARCHAR(255)"`
	Weight          string `xorm:"not null default '''' VARCHAR(255)"`
	Alipay          string `xorm:"not null default '''' VARCHAR(255)"`
	Icq             string `xorm:"not null default '''' VARCHAR(255)"`
	Qq              string `xorm:"not null default '''' VARCHAR(255)"`
	Yahoo           string `xorm:"not null default '''' VARCHAR(255)"`
	Msn             string `xorm:"not null default '''' VARCHAR(255)"`
	Taobao          string `xorm:"not null default '''' VARCHAR(255)"`
	Site            string `xorm:"not null default '''' VARCHAR(255)"`
	Bio             string `xorm:"not null TEXT"`
	Interest        string `xorm:"not null TEXT"`
	Field1          string `xorm:"not null TEXT"`
	Field2          string `xorm:"not null TEXT"`
	Field3          string `xorm:"not null TEXT"`
	Field4          string `xorm:"not null TEXT"`
	Field5          string `xorm:"not null TEXT"`
	Field6          string `xorm:"not null TEXT"`
	Field7          string `xorm:"not null TEXT"`
	Field8          string `xorm:"not null TEXT"`
}
