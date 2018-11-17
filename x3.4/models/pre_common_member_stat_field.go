package models

type PreCommonMemberStatField struct {
	Optionid   int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Fieldid    string `xorm:"not null default '''' index VARCHAR(255)"`
	Fieldvalue string `xorm:"not null default '''' VARCHAR(255)"`
	Hash       string `xorm:"not null default '''' VARCHAR(255)"`
	Users      int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Updatetime int    `xorm:"not null default 0 INT(10)"`
}
