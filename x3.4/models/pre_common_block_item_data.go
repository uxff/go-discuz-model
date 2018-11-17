package models

type PreCommonBlockItemData struct {
	Dataid       int    `xorm:"not null pk autoincr INT(10)"`
	Bid          int    `xorm:"not null default 0 index(bid) MEDIUMINT(8)"`
	Id           int    `xorm:"not null default 0 INT(10)"`
	Idtype       string `xorm:"not null default '''' VARCHAR(255)"`
	Itemtype     int    `xorm:"not null default 0 TINYINT(1)"`
	Title        string `xorm:"not null default '''' VARCHAR(255)"`
	Url          string `xorm:"not null default '''' VARCHAR(255)"`
	Pic          string `xorm:"not null default '''' VARCHAR(255)"`
	Picflag      int    `xorm:"not null default 0 TINYINT(1)"`
	Makethumb    int    `xorm:"not null default 0 TINYINT(1)"`
	Summary      string `xorm:"not null TEXT"`
	Showstyle    string `xorm:"not null TEXT"`
	Related      string `xorm:"not null TEXT"`
	Fields       string `xorm:"not null TEXT"`
	Displayorder int    `xorm:"not null default 0 index(bid) SMALLINT(6)"`
	Startdate    int    `xorm:"not null default 0 INT(10)"`
	Enddate      int    `xorm:"not null default 0 INT(10)"`
	Uid          int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Username     string `xorm:"not null default '''' VARCHAR(255)"`
	Dateline     int    `xorm:"not null default 0 INT(10)"`
	Isverified   int    `xorm:"not null default 0 TINYINT(1)"`
	Verifiedtime int    `xorm:"not null default 0 index(bid) INT(10)"`
	Stickgrade   int    `xorm:"not null default 0 index(bid) TINYINT(2)"`
}

func (t *PreCommonBlockItemData) TableName() string {
	return "pre_common_block_item_data"
}
