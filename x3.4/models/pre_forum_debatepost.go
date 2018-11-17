package models

type PreForumDebatepost struct {
	Pid      int    `xorm:"not null pk default 0 index(pid) INT(10)"`
	Stand    int    `xorm:"not null default 0 index(pid) TINYINT(1)"`
	Tid      int    `xorm:"not null default 0 index(tid) index(voters) MEDIUMINT(8)"`
	Uid      int    `xorm:"not null default 0 index(tid) MEDIUMINT(8)"`
	Dateline int    `xorm:"not null default 0 INT(10)"`
	Voters   int    `xorm:"not null default 0 index(voters) MEDIUMINT(10)"`
	Voterids string `xorm:"not null TEXT"`
}

func (t *PreForumDebatepost) TableName() string {
	return "pre_forum_debatepost"
}
