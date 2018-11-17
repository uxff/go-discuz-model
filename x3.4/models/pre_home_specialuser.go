package models

type PreHomeSpecialuser struct {
	Uid          int    `xorm:"not null default 0 index(uid) MEDIUMINT(8)"`
	Username     string `xorm:"not null default '''' VARCHAR(15)"`
	Status       int    `xorm:"not null default 0 index(displayorder) index(uid) TINYINT(1)"`
	Dateline     int    `xorm:"not null default 0 INT(10)"`
	Reason       string `xorm:"not null TEXT"`
	Opuid        int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Opusername   string `xorm:"not null default '''' VARCHAR(15)"`
	Displayorder int    `xorm:"not null default 0 index(displayorder) MEDIUMINT(8)"`
}

func (t *PreHomeSpecialuser) TableName() string {
	return "pre_home_specialuser"
}
