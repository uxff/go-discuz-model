package models

type PreCommonMytask struct {
	Uid      int    `xorm:"not null pk MEDIUMINT(8)"`
	Username string `xorm:"not null default '''' CHAR(15)"`
	Taskid   int    `xorm:"not null pk index(parter) SMALLINT(6)"`
	Status   int    `xorm:"not null default 0 TINYINT(1)"`
	Csc      string `xorm:"not null default '''' CHAR(255)"`
	Dateline int    `xorm:"not null default 0 index(parter) INT(10)"`
}

func (t *PreCommonMytask) TableName() string {
	return "pre_common_mytask"
}
