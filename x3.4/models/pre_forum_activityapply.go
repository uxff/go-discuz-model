package models

type PreForumActivityapply struct {
	Applyid    int    `xorm:"not null pk autoincr INT(10)"`
	Tid        int    `xorm:"not null default 0 index(dateline) index MEDIUMINT(8)"`
	Username   string `xorm:"not null default '''' VARCHAR(255)"`
	Uid        int    `xorm:"not null default 0 index MEDIUMINT(8)"`
	Message    string `xorm:"not null default '''' VARCHAR(255)"`
	Verified   int    `xorm:"not null default 0 TINYINT(1)"`
	Dateline   int    `xorm:"not null default 0 index(dateline) INT(10)"`
	Payment    int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Ufielddata string `xorm:"not null TEXT"`
}
