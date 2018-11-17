package models

type PreHomeShare struct {
	Sid           int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Itemid        int    `xorm:"not null MEDIUMINT(8)"`
	Type          string `xorm:"not null default '''' VARCHAR(30)"`
	Uid           int    `xorm:"not null default 0 index(uid) MEDIUMINT(8)"`
	Username      string `xorm:"not null default '''' VARCHAR(15)"`
	Fromuid       int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Dateline      int    `xorm:"not null default 0 index index(uid) INT(10)"`
	TitleTemplate string `xorm:"not null TEXT"`
	BodyTemplate  string `xorm:"not null TEXT"`
	BodyData      string `xorm:"not null TEXT"`
	BodyGeneral   string `xorm:"not null TEXT"`
	Image         string `xorm:"not null default '''' VARCHAR(255)"`
	ImageLink     string `xorm:"not null default '''' VARCHAR(255)"`
	Hot           int    `xorm:"not null default 0 index MEDIUMINT(8)"`
	Hotuser       string `xorm:"not null TEXT"`
	Status        int    `xorm:"not null TINYINT(1)"`
}

func (t *PreHomeShare) TableName() string {
	return "pre_home_share"
}
