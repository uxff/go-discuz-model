package models

type PreCommonMemberValidate struct {
	Uid         int    `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Submitdate  int    `xorm:"not null default 0 INT(10)"`
	Moddate     int    `xorm:"not null default 0 INT(10)"`
	Admin       string `xorm:"not null default '''' VARCHAR(15)"`
	Submittimes int    `xorm:"not null default 0 TINYINT(3)"`
	Status      int    `xorm:"not null default 0 index TINYINT(1)"`
	Message     string `xorm:"not null TEXT"`
	Remark      string `xorm:"not null TEXT"`
}
