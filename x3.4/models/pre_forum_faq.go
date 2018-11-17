package models

type PreForumFaq struct {
	Id           int    `xorm:"not null pk autoincr SMALLINT(6)"`
	Fpid         int    `xorm:"not null default 0 SMALLINT(6)"`
	Displayorder int    `xorm:"not null default 0 index TINYINT(3)"`
	Identifier   string `xorm:"not null VARCHAR(20)"`
	Keyword      string `xorm:"not null VARCHAR(50)"`
	Title        string `xorm:"not null VARCHAR(50)"`
	Message      string `xorm:"not null TEXT"`
}

func (t *PreForumFaq) TableName() string {
	return "pre_forum_faq"
}
