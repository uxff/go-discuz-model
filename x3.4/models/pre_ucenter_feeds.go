package models

type PreUcenterFeeds struct {
	Feedid        int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Appid         string `xorm:"not null default '''' VARCHAR(30)"`
	Icon          string `xorm:"not null default '''' VARCHAR(30)"`
	Uid           int    `xorm:"not null default 0 index(uid) MEDIUMINT(8)"`
	Username      string `xorm:"not null default '''' VARCHAR(15)"`
	Dateline      int    `xorm:"not null default 0 index(uid) INT(10)"`
	HashTemplate  string `xorm:"not null default '''' VARCHAR(32)"`
	HashData      string `xorm:"not null default '''' VARCHAR(32)"`
	TitleTemplate string `xorm:"not null default '''' TEXT"`
	TitleData     string `xorm:"not null default '''' TEXT"`
	BodyTemplate  string `xorm:"not null TEXT"`
	BodyData      string `xorm:"not null TEXT"`
	BodyGeneral   string `xorm:"not null TEXT"`
	Image1        string `xorm:"not null default '''' VARCHAR(255)"`
	Image1Link    string `xorm:"not null default '''' VARCHAR(255)"`
	Image2        string `xorm:"not null default '''' VARCHAR(255)"`
	Image2Link    string `xorm:"not null default '''' VARCHAR(255)"`
	Image3        string `xorm:"not null default '''' VARCHAR(255)"`
	Image3Link    string `xorm:"not null default '''' VARCHAR(255)"`
	Image4        string `xorm:"not null default '''' VARCHAR(255)"`
	Image4Link    string `xorm:"not null default '''' VARCHAR(255)"`
	TargetIds     string `xorm:"not null default '''' VARCHAR(255)"`
}

func (t *PreUcenterFeeds) TableName() string {
	return "pre_ucenter_feeds"
}
