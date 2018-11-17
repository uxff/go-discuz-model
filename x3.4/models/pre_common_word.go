package models

type PreCommonWord struct {
	Id          int    `xorm:"not null pk autoincr SMALLINT(6)"`
	Admin       string `xorm:"not null default '''' VARCHAR(15)"`
	Type        int    `xorm:"not null default 1 SMALLINT(6)"`
	Find        string `xorm:"not null default '''' VARCHAR(255)"`
	Replacement string `xorm:"not null default '''' VARCHAR(255)"`
	Extra       string `xorm:"not null default '''' VARCHAR(255)"`
}

func (t *PreCommonWord) TableName() string {
	return "pre_common_word"
}
