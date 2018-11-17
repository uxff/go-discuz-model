package models

type PreForumThread struct {
	Tid          int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Fid          int    `xorm:"not null default 0 index(displayorder) index(typeid) MEDIUMINT(8)"`
	Posttableid  int    `xorm:"not null default 0 SMALLINT(6)"`
	Typeid       int    `xorm:"not null default 0 index(typeid) SMALLINT(6)"`
	Sortid       int    `xorm:"not null default 0 index SMALLINT(6)"`
	Readperm     int    `xorm:"not null default 0 TINYINT(3)"`
	Price        int    `xorm:"not null default 0 SMALLINT(6)"`
	Author       string `xorm:"not null default '''' CHAR(15)"`
	Authorid     int    `xorm:"not null default 0 index MEDIUMINT(8)"`
	Subject      string `xorm:"not null default '''' CHAR(80)"`
	Dateline     int    `xorm:"not null default 0 INT(10)"`
	Lastpost     int    `xorm:"not null default 0 index(displayorder) index(isgroup) index(typeid) INT(10)"`
	Lastposter   string `xorm:"not null default '''' CHAR(15)"`
	Views        int    `xorm:"not null default 0 INT(10)"`
	Replies      int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Displayorder int    `xorm:"not null default 0 index(displayorder) index(typeid) TINYINT(1)"`
	Highlight    int    `xorm:"not null default 0 TINYINT(1)"`
	Digest       int    `xorm:"not null default 0 index TINYINT(1)"`
	Rate         int    `xorm:"not null default 0 TINYINT(1)"`
	Special      int    `xorm:"not null default 0 index TINYINT(1)"`
	Attachment   int    `xorm:"not null default 0 TINYINT(1)"`
	Moderated    int    `xorm:"not null default 0 TINYINT(1)"`
	Closed       int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Stickreply   int    `xorm:"not null default 0 TINYINT(1)"`
	Recommends   int    `xorm:"not null default 0 index SMALLINT(6)"`
	RecommendAdd int    `xorm:"not null default 0 SMALLINT(6)"`
	RecommendSub int    `xorm:"not null default 0 SMALLINT(6)"`
	Heats        int    `xorm:"not null default 0 index INT(10)"`
	Status       int    `xorm:"not null default 0 SMALLINT(6)"`
	Isgroup      int    `xorm:"not null default 0 index(isgroup) TINYINT(1)"`
	Favtimes     int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Sharetimes   int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Stamp        int    `xorm:"not null default -1 TINYINT(3)"`
	Icon         int    `xorm:"not null default -1 TINYINT(3)"`
	Pushedaid    int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Cover        int    `xorm:"not null default 0 SMALLINT(6)"`
	Replycredit  int    `xorm:"not null default 0 SMALLINT(6)"`
	Relatebytag  string `xorm:"not null default ''0'' CHAR(255)"`
	Maxposition  int    `xorm:"not null default 0 INT(8)"`
	Bgcolor      string `xorm:"not null default '''' CHAR(8)"`
	Comments     int    `xorm:"not null default 0 INT(10)"`
	Hidden       int    `xorm:"not null default 0 SMALLINT(6)"`
}
