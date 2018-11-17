package models

type PreHomeFeedApp struct {
	Feedid        int    `xorm:"not null pk autoincr INT(10)"`
	Appid         int    `xorm:"not null default 0 SMALLINT(6)"`
	Icon          string `xorm:"not null default '''' VARCHAR(30)"`
	Uid           int    `xorm:"not null default 0 index(uid) MEDIUMINT(8)"`
	Username      string `xorm:"not null default '''' VARCHAR(15)"`
	Dateline      int    `xorm:"not null default 0 index index(uid) INT(10)"`
	Friend        int    `xorm:"not null default 0 TINYINT(1)"`
	HashTemplate  string `xorm:"not null default '''' VARCHAR(32)"`
	HashData      string `xorm:"not null default '''' VARCHAR(32)"`
	TitleTemplate string `xorm:"not null TEXT"`
	TitleData     string `xorm:"not null TEXT"`
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
	TargetIds     string `xorm:"not null TEXT"`
}

func (t *PreHomeFeedApp) TableName() string {
	return "pre_home_feed_app"
}
