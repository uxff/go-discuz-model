package models

type PreForumDebate struct {
	Tid            int    `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Uid            int    `xorm:"not null default 0 index(uid) MEDIUMINT(8)"`
	Starttime      int    `xorm:"not null default 0 index(uid) INT(10)"`
	Endtime        int    `xorm:"not null default 0 INT(10)"`
	Affirmdebaters int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Negadebaters   int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Affirmvotes    int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Negavotes      int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Umpire         string `xorm:"not null default '''' VARCHAR(15)"`
	Winner         int    `xorm:"not null default 0 TINYINT(1)"`
	Bestdebater    string `xorm:"not null default '''' VARCHAR(50)"`
	Affirmpoint    string `xorm:"not null TEXT"`
	Negapoint      string `xorm:"not null TEXT"`
	Umpirepoint    string `xorm:"not null TEXT"`
	Affirmvoterids string `xorm:"not null TEXT"`
	Negavoterids   string `xorm:"not null TEXT"`
	Affirmreplies  int    `xorm:"not null MEDIUMINT(8)"`
	Negareplies    int    `xorm:"not null MEDIUMINT(8)"`
}
