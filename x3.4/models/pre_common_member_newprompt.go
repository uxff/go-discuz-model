package models

type PreCommonMemberNewprompt struct {
	Uid  int    `xorm:"not null pk MEDIUMINT(8)"`
	Data string `xorm:"not null default '''' VARCHAR(255)"`
}

func (t *PreCommonMemberNewprompt) TableName() string {
	return "pre_common_member_newprompt"
}
