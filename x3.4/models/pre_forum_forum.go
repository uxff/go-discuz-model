package models

type PreForumForum struct {
	Fid              int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Fup              int    `xorm:"not null default 0 index index(fup_type) MEDIUMINT(8)"`
	Type             string `xorm:"not null default ''forum'' index(forum) index(fup_type) ENUM('forum','group','sub')"`
	Name             string `xorm:"not null default '''' CHAR(50)"`
	Status           int    `xorm:"not null default 0 index(forum) TINYINT(1)"`
	Displayorder     int    `xorm:"not null default 0 index(forum) index(fup_type) SMALLINT(6)"`
	Styleid          int    `xorm:"not null default 0 SMALLINT(6)"`
	Threads          int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Posts            int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Todayposts       int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Yesterdayposts   int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Rank             int    `xorm:"not null default 0 SMALLINT(6)"`
	Oldrank          int    `xorm:"not null default 0 SMALLINT(6)"`
	Lastpost         string `xorm:"not null default '''' CHAR(110)"`
	Domain           string `xorm:"not null default '''' CHAR(15)"`
	Allowsmilies     int    `xorm:"not null default 0 TINYINT(1)"`
	Allowhtml        int    `xorm:"not null default 0 TINYINT(1)"`
	Allowbbcode      int    `xorm:"not null default 0 TINYINT(1)"`
	Allowimgcode     int    `xorm:"not null default 0 TINYINT(1)"`
	Allowmediacode   int    `xorm:"not null default 0 TINYINT(1)"`
	Allowanonymous   int    `xorm:"not null default 0 TINYINT(1)"`
	Allowpostspecial int    `xorm:"not null default 0 SMALLINT(6)"`
	Allowspecialonly int    `xorm:"not null default 0 TINYINT(1)"`
	Allowappend      int    `xorm:"not null default 0 TINYINT(1)"`
	Alloweditrules   int    `xorm:"not null default 0 TINYINT(1)"`
	Allowfeed        int    `xorm:"not null default 1 TINYINT(1)"`
	Allowside        int    `xorm:"not null default 0 TINYINT(1)"`
	Recyclebin       int    `xorm:"not null default 0 TINYINT(1)"`
	Modnewposts      int    `xorm:"not null default 0 TINYINT(1)"`
	Jammer           int    `xorm:"not null default 0 TINYINT(1)"`
	Disablewatermark int    `xorm:"not null default 0 TINYINT(1)"`
	Inheritedmod     int    `xorm:"not null default 0 TINYINT(1)"`
	Autoclose        int    `xorm:"not null default 0 SMALLINT(6)"`
	Forumcolumns     int    `xorm:"not null default 0 TINYINT(3)"`
	Catforumcolumns  int    `xorm:"not null default 0 TINYINT(3)"`
	Threadcaches     int    `xorm:"not null default 0 TINYINT(1)"`
	Alloweditpost    int    `xorm:"not null default 1 TINYINT(1)"`
	Simple           int    `xorm:"not null default 0 TINYINT(1)"`
	Modworks         int    `xorm:"not null default 0 TINYINT(1)"`
	Allowglobalstick int    `xorm:"not null default 1 TINYINT(1)"`
	Level            int    `xorm:"not null default 0 SMALLINT(6)"`
	Commoncredits    int    `xorm:"not null default 0 INT(10)"`
	Archive          int    `xorm:"not null default 0 TINYINT(1)"`
	Recommend        int    `xorm:"not null default 0 SMALLINT(6)"`
	Favtimes         int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Sharetimes       int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Disablethumb     int    `xorm:"not null default 0 TINYINT(1)"`
	Disablecollect   int    `xorm:"not null default 0 TINYINT(1)"`
}
