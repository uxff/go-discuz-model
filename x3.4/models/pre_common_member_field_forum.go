package models

type PreCommonMemberFieldForum struct {
	Uid            int    `xorm:"not null pk MEDIUMINT(8)"`
	Publishfeed    int    `xorm:"not null default 0 TINYINT(3)"`
	Customshow     int    `xorm:"not null default 26 TINYINT(1)"`
	Customstatus   string `xorm:"not null default '''' VARCHAR(30)"`
	Medals         string `xorm:"not null TEXT"`
	Sightml        string `xorm:"not null TEXT"`
	Groupterms     string `xorm:"not null TEXT"`
	Authstr        string `xorm:"not null default '''' VARCHAR(20)"`
	Groups         string `xorm:"not null MEDIUMTEXT"`
	Attentiongroup string `xorm:"not null default '''' VARCHAR(255)"`
}
