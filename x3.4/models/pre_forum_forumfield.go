package models

type PreForumForumfield struct {
	Fid              int    `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Description      string `xorm:"not null TEXT"`
	Password         string `xorm:"not null default '''' VARCHAR(12)"`
	Icon             string `xorm:"not null default '''' VARCHAR(255)"`
	Redirect         string `xorm:"not null default '''' VARCHAR(255)"`
	Attachextensions string `xorm:"not null default '''' VARCHAR(255)"`
	Creditspolicy    string `xorm:"not null MEDIUMTEXT"`
	Formulaperm      string `xorm:"not null TEXT"`
	Moderators       string `xorm:"not null TEXT"`
	Rules            string `xorm:"not null TEXT"`
	Threadtypes      string `xorm:"not null TEXT"`
	Threadsorts      string `xorm:"not null TEXT"`
	Viewperm         string `xorm:"not null TEXT"`
	Postperm         string `xorm:"not null TEXT"`
	Replyperm        string `xorm:"not null TEXT"`
	Getattachperm    string `xorm:"not null TEXT"`
	Postattachperm   string `xorm:"not null TEXT"`
	Postimageperm    string `xorm:"not null TEXT"`
	Spviewperm       string `xorm:"not null TEXT"`
	Seotitle         string `xorm:"not null TEXT"`
	Keywords         string `xorm:"not null TEXT"`
	Seodescription   string `xorm:"not null TEXT"`
	SupePushsetting  string `xorm:"not null TEXT"`
	Modrecommend     string `xorm:"not null TEXT"`
	Threadplugin     string `xorm:"not null TEXT"`
	Replybg          string `xorm:"not null TEXT"`
	Extra            string `xorm:"not null TEXT"`
	Jointype         int    `xorm:"not null default 0 TINYINT(1)"`
	Gviewperm        int    `xorm:"not null default 0 TINYINT(1)"`
	Membernum        int    `xorm:"not null default 0 index SMALLINT(6)"`
	Dateline         int    `xorm:"not null default 0 index INT(10)"`
	Lastupdate       int    `xorm:"not null default 0 index INT(10)"`
	Activity         int    `xorm:"not null default 0 index INT(10)"`
	Founderuid       int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Foundername      string `xorm:"not null default '''' VARCHAR(255)"`
	Banner           string `xorm:"not null default '''' VARCHAR(255)"`
	Groupnum         int    `xorm:"not null default 0 SMALLINT(6)"`
	Commentitem      string `xorm:"not null TEXT"`
	Relatedgroup     string `xorm:"not null TEXT"`
	Picstyle         int    `xorm:"not null default 0 TINYINT(1)"`
	Widthauto        int    `xorm:"not null default 0 TINYINT(1)"`
	Noantitheft      int    `xorm:"not null default 0 TINYINT(1)"`
	Noforumhidewater int    `xorm:"not null default 0 TINYINT(1)"`
	Noforumrecommend int    `xorm:"not null default 0 TINYINT(1)"`
	Livetid          int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Price            int    `xorm:"not null default 0 MEDIUMINT(8)"`
}

func (t *PreForumForumfield) TableName() string {
	return "pre_forum_forumfield"
}
