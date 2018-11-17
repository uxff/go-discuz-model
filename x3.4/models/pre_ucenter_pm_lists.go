package models

type PreUcenterPmLists struct {
	Plid        int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Authorid    int    `xorm:"not null default 0 index(authorid) MEDIUMINT(8)"`
	Pmtype      int    `xorm:"not null default 0 index TINYINT(1)"`
	Subject     string `xorm:"not null VARCHAR(80)"`
	Members     int    `xorm:"not null default 0 SMALLINT(5)"`
	MinMax      string `xorm:"not null index VARCHAR(17)"`
	Dateline    int    `xorm:"not null default 0 index(authorid) INT(10)"`
	Lastmessage string `xorm:"not null TEXT"`
}
