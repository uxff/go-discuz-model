package models

type PreForumTrade struct {
	Tid            int    `xorm:"not null pk index(displayorder) MEDIUMINT(8)"`
	Pid            int    `xorm:"not null pk index INT(10)"`
	Typeid         int    `xorm:"not null index SMALLINT(6)"`
	Sellerid       int    `xorm:"not null index index(sellertrades) MEDIUMINT(8)"`
	Seller         string `xorm:"not null CHAR(15)"`
	Account        string `xorm:"not null CHAR(50)"`
	Tenpayaccount  string `xorm:"not null default '''' CHAR(20)"`
	Subject        string `xorm:"not null CHAR(100)"`
	Price          string `xorm:"not null DECIMAL(8,2)"`
	Amount         int    `xorm:"not null default 1 SMALLINT(6)"`
	Quality        int    `xorm:"not null default 0 TINYINT(1)"`
	Locus          string `xorm:"not null CHAR(20)"`
	Transport      int    `xorm:"not null default 0 TINYINT(1)"`
	Ordinaryfee    int    `xorm:"not null default 0 SMALLINT(4)"`
	Expressfee     int    `xorm:"not null default 0 SMALLINT(4)"`
	Emsfee         int    `xorm:"not null default 0 SMALLINT(4)"`
	Itemtype       int    `xorm:"not null default 0 TINYINT(1)"`
	Dateline       int    `xorm:"not null default 0 INT(10)"`
	Expiration     int    `xorm:"not null default 0 index INT(10)"`
	Lastbuyer      string `xorm:"not null CHAR(15)"`
	Lastupdate     int    `xorm:"not null default 0 INT(10)"`
	Totalitems     int    `xorm:"not null default 0 index(sellertrades) index SMALLINT(5)"`
	Tradesum       string `xorm:"not null default 0.00 index(sellertrades) index DECIMAL(8,2)"`
	Closed         int    `xorm:"not null default 0 TINYINT(1)"`
	Aid            int    `xorm:"not null MEDIUMINT(8)"`
	Displayorder   int    `xorm:"not null index(displayorder) TINYINT(1)"`
	Costprice      string `xorm:"not null DECIMAL(8,2)"`
	Credit         int    `xorm:"not null default 0 INT(10)"`
	Costcredit     int    `xorm:"not null default 0 INT(10)"`
	Credittradesum int    `xorm:"not null default 0 index INT(10)"`
}
