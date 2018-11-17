package models

type PreCommonBlockFavorite struct {
	Favid    int `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Uid      int `xorm:"not null default 0 index(uid) MEDIUMINT(8)"`
	Bid      int `xorm:"not null default 0 MEDIUMINT(8)"`
	Dateline int `xorm:"not null default 0 index(uid) INT(10)"`
}

func (t *PreCommonBlockFavorite) TableName() string {
	return "pre_common_block_favorite"
}
