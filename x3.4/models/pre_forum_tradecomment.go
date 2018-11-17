package models

type PreForumTradecomment struct {
	Id          int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Orderid     string `xorm:"not null index CHAR(32)"`
	Pid         int    `xorm:"not null INT(10)"`
	Type        int    `xorm:"not null index(rateeid) index(raterid) TINYINT(1)"`
	Raterid     int    `xorm:"not null index(raterid) MEDIUMINT(8)"`
	Rater       string `xorm:"not null CHAR(15)"`
	Rateeid     int    `xorm:"not null index(rateeid) MEDIUMINT(8)"`
	Ratee       string `xorm:"not null CHAR(15)"`
	Message     string `xorm:"not null CHAR(200)"`
	Explanation string `xorm:"not null CHAR(200)"`
	Score       int    `xorm:"not null TINYINT(1)"`
	Dateline    int    `xorm:"not null index(rateeid) index(raterid) INT(10)"`
}

func (t *PreForumTradecomment) TableName() string {
	return "pre_forum_tradecomment"
}
