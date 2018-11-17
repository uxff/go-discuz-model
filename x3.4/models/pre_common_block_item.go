package models

type PreCommonBlockItem struct {
	Itemid       int    `xorm:"not null pk autoincr INT(10)"`
	Bid          int    `xorm:"not null default 0 index MEDIUMINT(8)"`
	Id           int    `xorm:"not null default 0 INT(10)"`
	Idtype       string `xorm:"not null default '''' VARCHAR(255)"`
	Itemtype     int    `xorm:"not null default 0 TINYINT(1)"`
	Title        string `xorm:"not null default '''' VARCHAR(255)"`
	Url          string `xorm:"not null default '''' VARCHAR(255)"`
	Pic          string `xorm:"not null default '''' VARCHAR(255)"`
	Picflag      int    `xorm:"not null default 0 TINYINT(1)"`
	Makethumb    int    `xorm:"not null default 0 TINYINT(1)"`
	Thumbpath    string `xorm:"not null default '''' VARCHAR(255)"`
	Summary      string `xorm:"not null TEXT"`
	Showstyle    string `xorm:"not null TEXT"`
	Related      string `xorm:"not null TEXT"`
	Fields       string `xorm:"not null TEXT"`
	Displayorder int    `xorm:"not null default 0 SMALLINT(6)"`
	Startdate    int    `xorm:"not null default 0 INT(10)"`
	Enddate      int    `xorm:"not null default 0 INT(10)"`
}
