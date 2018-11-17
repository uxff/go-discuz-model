package models

type PreForumPollvoter struct {
	Tid      int    `xorm:"not null default 0 index MEDIUMINT(8)"`
	Uid      int    `xorm:"not null default 0 index(uid) MEDIUMINT(8)"`
	Username string `xorm:"not null default '''' VARCHAR(15)"`
	Options  string `xorm:"not null TEXT"`
	Dateline int    `xorm:"not null default 0 index(uid) INT(10)"`
}
