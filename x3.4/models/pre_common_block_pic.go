package models

type PreCommonBlockPic struct {
	Picid   int    `xorm:"not null pk autoincr INT(10)"`
	Bid     int    `xorm:"not null default 0 index(bid) MEDIUMINT(8)"`
	Itemid  int    `xorm:"not null default 0 index(bid) INT(10)"`
	Pic     string `xorm:"not null default '''' VARCHAR(255)"`
	Picflag int    `xorm:"not null default 0 TINYINT(1)"`
	Type    int    `xorm:"not null default 0 TINYINT(1)"`
}

func (t *PreCommonBlockPic) TableName() string {
	return "pre_common_block_pic"
}
