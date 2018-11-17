package models

type PreCommonMemberNewprompt struct {
	Uid  int    `xorm:"not null pk MEDIUMINT(8)"`
	Data string `xorm:"not null default '''' VARCHAR(255)"`
}
