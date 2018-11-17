package models

type PreCommonMemberCount struct {
	Uid             int `xorm:"not null pk MEDIUMINT(8)"`
	Extcredits1     int `xorm:"not null default 0 INT(10)"`
	Extcredits2     int `xorm:"not null default 0 INT(10)"`
	Extcredits3     int `xorm:"not null default 0 INT(10)"`
	Extcredits4     int `xorm:"not null default 0 INT(10)"`
	Extcredits5     int `xorm:"not null default 0 INT(10)"`
	Extcredits6     int `xorm:"not null default 0 INT(10)"`
	Extcredits7     int `xorm:"not null default 0 INT(10)"`
	Extcredits8     int `xorm:"not null default 0 INT(10)"`
	Friends         int `xorm:"not null default 0 SMALLINT(6)"`
	Posts           int `xorm:"not null default 0 index MEDIUMINT(8)"`
	Threads         int `xorm:"not null default 0 MEDIUMINT(8)"`
	Digestposts     int `xorm:"not null default 0 SMALLINT(6)"`
	Doings          int `xorm:"not null default 0 SMALLINT(6)"`
	Blogs           int `xorm:"not null default 0 SMALLINT(6)"`
	Albums          int `xorm:"not null default 0 SMALLINT(6)"`
	Sharings        int `xorm:"not null default 0 SMALLINT(6)"`
	Attachsize      int `xorm:"not null default 0 INT(10)"`
	Views           int `xorm:"not null default 0 MEDIUMINT(8)"`
	Oltime          int `xorm:"not null default 0 SMALLINT(6)"`
	Todayattachs    int `xorm:"not null default 0 SMALLINT(6)"`
	Todayattachsize int `xorm:"not null default 0 INT(10)"`
	Feeds           int `xorm:"not null default 0 MEDIUMINT(8)"`
	Follower        int `xorm:"not null default 0 MEDIUMINT(8)"`
	Following       int `xorm:"not null default 0 MEDIUMINT(8)"`
	Newfollower     int `xorm:"not null default 0 MEDIUMINT(8)"`
	Blacklist       int `xorm:"not null default 0 MEDIUMINT(8)"`
}
