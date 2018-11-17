package models

type PreCommonMemberProfileSetting struct {
	Fieldid        string `xorm:"not null pk default '''' VARCHAR(255)"`
	Available      int    `xorm:"not null default 0 TINYINT(1)"`
	Invisible      int    `xorm:"not null default 0 TINYINT(1)"`
	Needverify     int    `xorm:"not null default 0 TINYINT(1)"`
	Title          string `xorm:"not null default '''' VARCHAR(255)"`
	Description    string `xorm:"not null default '''' VARCHAR(255)"`
	Displayorder   int    `xorm:"not null default 0 SMALLINT(6)"`
	Required       int    `xorm:"not null default 0 TINYINT(1)"`
	Unchangeable   int    `xorm:"not null default 0 TINYINT(1)"`
	Showincard     int    `xorm:"not null default 0 TINYINT(1)"`
	Showinthread   int    `xorm:"not null default 0 TINYINT(1)"`
	Showinregister int    `xorm:"not null default 0 TINYINT(1)"`
	Allowsearch    int    `xorm:"not null default 0 TINYINT(1)"`
	Formtype       string `xorm:"not null VARCHAR(255)"`
	Size           int    `xorm:"not null default 0 SMALLINT(6)"`
	Choices        string `xorm:"not null TEXT"`
	Validate       string `xorm:"not null TEXT"`
}

func (t *PreCommonMemberProfileSetting) TableName() string {
	return "pre_common_member_profile_setting"
}
