package models

type PreCommonReport struct {
	Id       int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Urlkey   string `xorm:"not null default '''' index CHAR(32)"`
	Url      string `xorm:"not null default '''' VARCHAR(255)"`
	Message  string `xorm:"not null TEXT"`
	Uid      int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Username string `xorm:"not null default '''' VARCHAR(15)"`
	Dateline int    `xorm:"not null default 0 INT(10)"`
	Num      int    `xorm:"not null default 1 SMALLINT(6)"`
	Opuid    int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Opname   string `xorm:"not null default '''' VARCHAR(15)"`
	Optime   int    `xorm:"not null default 0 INT(10)"`
	Opresult string `xorm:"not null default '''' VARCHAR(255)"`
	Fid      int    `xorm:"not null default 0 index MEDIUMINT(8)"`
}

func (t *PreCommonReport) TableName() string {
	return "pre_common_report"
}
