package models

type PreHomeAlbumCategory struct {
	Catid        int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Upid         int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Catname      string `xorm:"not null default '''' VARCHAR(255)"`
	Num          int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Displayorder int    `xorm:"not null default 0 SMALLINT(6)"`
}

func (t *PreHomeAlbumCategory) TableName() string {
	return "pre_home_album_category"
}
