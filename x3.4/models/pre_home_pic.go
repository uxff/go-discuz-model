package models

type PreHomePic struct {
	Picid      int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Albumid    int    `xorm:"not null default 0 index(albumid) MEDIUMINT(8)"`
	Uid        int    `xorm:"not null default 0 index MEDIUMINT(8)"`
	Username   string `xorm:"not null default '''' VARCHAR(15)"`
	Dateline   int    `xorm:"not null default 0 index(albumid) INT(10)"`
	Postip     string `xorm:"not null default '''' VARCHAR(255)"`
	Port       int    `xorm:"not null default 0 SMALLINT(6)"`
	Filename   string `xorm:"not null default '''' VARCHAR(255)"`
	Title      string `xorm:"not null default '''' VARCHAR(255)"`
	Type       string `xorm:"not null default '''' VARCHAR(255)"`
	Size       int    `xorm:"not null default 0 INT(10)"`
	Filepath   string `xorm:"not null default '''' VARCHAR(255)"`
	Thumb      int    `xorm:"not null default 0 TINYINT(1)"`
	Remote     int    `xorm:"not null default 0 TINYINT(1)"`
	Hot        int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Sharetimes int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Click1     int    `xorm:"not null default 0 SMALLINT(6)"`
	Click2     int    `xorm:"not null default 0 SMALLINT(6)"`
	Click3     int    `xorm:"not null default 0 SMALLINT(6)"`
	Click4     int    `xorm:"not null default 0 SMALLINT(6)"`
	Click5     int    `xorm:"not null default 0 SMALLINT(6)"`
	Click6     int    `xorm:"not null default 0 SMALLINT(6)"`
	Click7     int    `xorm:"not null default 0 SMALLINT(6)"`
	Click8     int    `xorm:"not null default 0 SMALLINT(6)"`
	Magicframe int    `xorm:"not null default 0 TINYINT(6)"`
	Status     int    `xorm:"not null default 0 TINYINT(1)"`
}
