package models

type PrePortalCategory struct {
	Catid                int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Upid                 int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Catname              string `xorm:"not null default '''' VARCHAR(255)"`
	Articles             int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Allowcomment         int    `xorm:"not null default 1 TINYINT(1)"`
	Displayorder         int    `xorm:"not null default 0 SMALLINT(6)"`
	Notinheritedarticle  int    `xorm:"not null default 0 TINYINT(1)"`
	Notinheritedblock    int    `xorm:"not null default 0 TINYINT(1)"`
	Domain               string `xorm:"not null default '''' VARCHAR(255)"`
	Url                  string `xorm:"not null default '''' VARCHAR(255)"`
	Uid                  int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Username             string `xorm:"not null default '''' VARCHAR(255)"`
	Dateline             int    `xorm:"not null default 0 INT(10)"`
	Closed               int    `xorm:"not null default 0 TINYINT(1)"`
	Shownav              int    `xorm:"not null default 0 TINYINT(1)"`
	Description          string `xorm:"not null TEXT"`
	Seotitle             string `xorm:"not null TEXT"`
	Keyword              string `xorm:"not null TEXT"`
	Primaltplname        string `xorm:"not null default '''' VARCHAR(255)"`
	Articleprimaltplname string `xorm:"not null default '''' VARCHAR(255)"`
	Disallowpublish      int    `xorm:"not null default 0 TINYINT(1)"`
	Foldername           string `xorm:"not null default '''' VARCHAR(255)"`
	Notshowarticlesummay string `xorm:"not null default '''' VARCHAR(255)"`
	Perpage              int    `xorm:"not null default 0 SMALLINT(6)"`
	Maxpages             int    `xorm:"not null default 0 SMALLINT(6)"`
	Noantitheft          int    `xorm:"not null default 0 TINYINT(1)"`
	Lastpublish          int    `xorm:"not null default 0 INT(10)"`
}
