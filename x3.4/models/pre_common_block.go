package models

type PreCommonBlock struct {
	Bid            int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Blockclass     string `xorm:"not null default ''0'' VARCHAR(255)"`
	Blocktype      int    `xorm:"not null default 0 TINYINT(1)"`
	Name           string `xorm:"not null default '''' VARCHAR(255)"`
	Title          string `xorm:"not null TEXT"`
	Classname      string `xorm:"not null default '''' VARCHAR(255)"`
	Summary        string `xorm:"not null TEXT"`
	Uid            int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Username       string `xorm:"not null default '''' VARCHAR(255)"`
	Styleid        int    `xorm:"not null default 0 SMALLINT(6)"`
	Blockstyle     string `xorm:"not null TEXT"`
	Picwidth       int    `xorm:"not null default 0 SMALLINT(6)"`
	Picheight      int    `xorm:"not null default 0 SMALLINT(6)"`
	Target         string `xorm:"not null default '''' VARCHAR(255)"`
	Dateformat     string `xorm:"not null default '''' VARCHAR(255)"`
	Dateuformat    int    `xorm:"not null default 0 TINYINT(1)"`
	Script         string `xorm:"not null default '''' VARCHAR(255)"`
	Param          string `xorm:"not null TEXT"`
	Shownum        int    `xorm:"not null default 0 SMALLINT(6)"`
	Cachetime      int    `xorm:"not null default 0 INT(10)"`
	Cachetimerange string `xorm:"not null default '''' CHAR(5)"`
	Punctualupdate int    `xorm:"not null default 0 TINYINT(1)"`
	Hidedisplay    int    `xorm:"not null default 0 TINYINT(1)"`
	Dateline       int    `xorm:"not null default 0 INT(10)"`
	Notinherited   int    `xorm:"not null default 0 TINYINT(1)"`
	Isblank        int    `xorm:"not null default 0 TINYINT(1)"`
}
