package models

type PreHomeAlbum struct {
	Albumid    int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Albumname  string `xorm:"not null default '''' VARCHAR(50)"`
	Catid      int    `xorm:"not null default 0 SMALLINT(6)"`
	Uid        int    `xorm:"not null default 0 index(uid) MEDIUMINT(8)"`
	Username   string `xorm:"not null default '''' VARCHAR(15)"`
	Dateline   int    `xorm:"not null default 0 INT(10)"`
	Updatetime int    `xorm:"not null default 0 index(uid) index INT(10)"`
	Picnum     int    `xorm:"not null default 0 SMALLINT(6)"`
	Pic        string `xorm:"not null default '''' VARCHAR(255)"`
	Picflag    int    `xorm:"not null default 0 TINYINT(1)"`
	Friend     int    `xorm:"not null default 0 TINYINT(1)"`
	Password   string `xorm:"not null default '''' VARCHAR(10)"`
	TargetIds  string `xorm:"not null TEXT"`
	Favtimes   int    `xorm:"not null MEDIUMINT(8)"`
	Sharetimes int    `xorm:"not null MEDIUMINT(8)"`
	Depict     string `xorm:"not null TEXT"`
}
