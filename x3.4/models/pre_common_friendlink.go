package models

type PreCommonFriendlink struct {
	Id           int    `xorm:"not null pk autoincr SMALLINT(6)"`
	Displayorder int    `xorm:"not null default 0 TINYINT(3)"`
	Name         string `xorm:"not null default '''' VARCHAR(100)"`
	Url          string `xorm:"not null default '''' VARCHAR(255)"`
	Description  string `xorm:"not null MEDIUMTEXT"`
	Logo         string `xorm:"not null default '''' VARCHAR(255)"`
	Type         int    `xorm:"not null default 0 TINYINT(3)"`
}

func (t *PreCommonFriendlink) TableName() string {
	return "pre_common_friendlink"
}
