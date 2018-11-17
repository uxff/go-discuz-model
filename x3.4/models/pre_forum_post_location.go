package models

type PreForumPostLocation struct {
	Pid      int    `xorm:"not null pk default 0 INT(10)"`
	Tid      int    `xorm:"default 0 index MEDIUMINT(8)"`
	Uid      int    `xorm:"default 0 index MEDIUMINT(8)"`
	Mapx     string `xorm:"not null VARCHAR(255)"`
	Mapy     string `xorm:"not null VARCHAR(255)"`
	Location string `xorm:"not null VARCHAR(255)"`
}
