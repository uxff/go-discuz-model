package models

type PreHomeBlog struct {
	Blogid     int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Uid        int    `xorm:"not null default 0 index(uid) MEDIUMINT(8)"`
	Username   string `xorm:"not null default '''' CHAR(15)"`
	Subject    string `xorm:"not null default '''' CHAR(80)"`
	Classid    int    `xorm:"not null default 0 SMALLINT(6)"`
	Catid      int    `xorm:"not null default 0 SMALLINT(6)"`
	Viewnum    int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Replynum   int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Hot        int    `xorm:"not null default 0 index MEDIUMINT(8)"`
	Dateline   int    `xorm:"not null default 0 index index(uid) INT(10)"`
	Picflag    int    `xorm:"not null default 0 TINYINT(1)"`
	Noreply    int    `xorm:"not null default 0 TINYINT(1)"`
	Friend     int    `xorm:"not null default 0 TINYINT(1)"`
	Password   string `xorm:"not null default '''' CHAR(10)"`
	Favtimes   int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Sharetimes int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Status     int    `xorm:"not null default 0 TINYINT(1)"`
	Click1     int    `xorm:"not null default 0 SMALLINT(6)"`
	Click2     int    `xorm:"not null default 0 SMALLINT(6)"`
	Click3     int    `xorm:"not null default 0 SMALLINT(6)"`
	Click4     int    `xorm:"not null default 0 SMALLINT(6)"`
	Click5     int    `xorm:"not null default 0 SMALLINT(6)"`
	Click6     int    `xorm:"not null default 0 SMALLINT(6)"`
	Click7     int    `xorm:"not null default 0 SMALLINT(6)"`
	Click8     int    `xorm:"not null default 0 SMALLINT(6)"`
}

func (t *PreHomeBlog) TableName() string {
	return "pre_home_blog"
}
