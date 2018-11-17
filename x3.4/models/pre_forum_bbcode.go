package models

type PreForumBbcode struct {
	Id           int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Available    int    `xorm:"not null default 0 TINYINT(1)"`
	Tag          string `xorm:"not null default '''' VARCHAR(100)"`
	Icon         string `xorm:"not null VARCHAR(255)"`
	Replacement  string `xorm:"not null TEXT"`
	Example      string `xorm:"not null default '''' VARCHAR(255)"`
	Explanation  string `xorm:"not null TEXT"`
	Params       int    `xorm:"not null default 1 TINYINT(1)"`
	Prompt       string `xorm:"not null TEXT"`
	Nest         int    `xorm:"not null default 1 TINYINT(3)"`
	Displayorder int    `xorm:"not null default 0 TINYINT(3)"`
	Perm         string `xorm:"not null TEXT"`
}
