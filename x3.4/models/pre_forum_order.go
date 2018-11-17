package models

type PreForumOrder struct {
	Orderid     string  `xorm:"not null pk default '''' unique CHAR(32)"`
	Status      string  `xorm:"not null default '''' CHAR(3)"`
	Buyer       string  `xorm:"not null default '''' CHAR(50)"`
	Admin       string  `xorm:"not null default '''' CHAR(15)"`
	Uid         int     `xorm:"not null default 0 index(uid) MEDIUMINT(8)"`
	Amount      int     `xorm:"not null default 0 INT(10)"`
	Price       float32 `xorm:"not null default 0.00 FLOAT(7,2)"`
	Submitdate  int     `xorm:"not null default 0 index index(uid) INT(10)"`
	Confirmdate int     `xorm:"not null default 0 INT(10)"`
	Email       string  `xorm:"not null default '''' CHAR(40)"`
	Ip          string  `xorm:"not null default '''' CHAR(15)"`
}
