package models

type PreForumTradelog struct {
	Tid           int    `xorm:"not null index(tid) MEDIUMINT(8)"`
	Pid           int    `xorm:"not null index index(tid) INT(10)"`
	Orderid       string `xorm:"not null pk unique VARCHAR(32)"`
	Tradeno       string `xorm:"not null VARCHAR(32)"`
	Paytype       int    `xorm:"not null default 0 TINYINT(1)"`
	Subject       string `xorm:"not null VARCHAR(100)"`
	Price         string `xorm:"not null default 0.00 DECIMAL(8,2)"`
	Quality       int    `xorm:"not null default 0 TINYINT(1)"`
	Itemtype      int    `xorm:"not null default 0 TINYINT(1)"`
	Number        int    `xorm:"not null default 0 SMALLINT(5)"`
	Tax           string `xorm:"not null default 0.00 DECIMAL(6,2)"`
	Locus         string `xorm:"not null VARCHAR(100)"`
	Sellerid      int    `xorm:"not null index index(sellerlog) MEDIUMINT(8)"`
	Seller        string `xorm:"not null VARCHAR(15)"`
	Selleraccount string `xorm:"not null VARCHAR(50)"`
	Tenpayaccount string `xorm:"not null default ''0'' VARCHAR(20)"`
	Buyerid       int    `xorm:"not null index index(buyerlog) MEDIUMINT(8)"`
	Buyer         string `xorm:"not null VARCHAR(15)"`
	Buyercontact  string `xorm:"not null VARCHAR(50)"`
	Buyercredits  int    `xorm:"not null default 0 SMALLINT(5)"`
	Buyermsg      string `xorm:"default 'NULL' VARCHAR(200)"`
	Status        int    `xorm:"not null default 0 index(buyerlog) index(sellerlog) index TINYINT(1)"`
	Lastupdate    int    `xorm:"not null default 0 index(buyerlog) index(sellerlog) INT(10)"`
	Offline       int    `xorm:"not null default 0 TINYINT(1)"`
	Buyername     string `xorm:"not null VARCHAR(50)"`
	Buyerzip      string `xorm:"not null VARCHAR(10)"`
	Buyerphone    string `xorm:"not null VARCHAR(20)"`
	Buyermobile   string `xorm:"not null VARCHAR(20)"`
	Transport     int    `xorm:"not null default 0 TINYINT(1)"`
	Transportfee  int    `xorm:"not null default 0 SMALLINT(6)"`
	Baseprice     string `xorm:"not null DECIMAL(8,2)"`
	Discount      int    `xorm:"not null default 0 TINYINT(1)"`
	Ratestatus    int    `xorm:"not null default 0 TINYINT(1)"`
	Message       string `xorm:"not null TEXT"`
	Credit        int    `xorm:"not null default 0 INT(10)"`
	Basecredit    int    `xorm:"not null default 0 INT(10)"`
}

func (t *PreForumTradelog) TableName() string {
	return "pre_forum_tradelog"
}
