package models

type PreUcenterMailqueue struct {
	Mailid   int    `xorm:"not null pk autoincr INT(10)"`
	Touid    int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Tomail   string `xorm:"not null VARCHAR(32)"`
	Frommail string `xorm:"not null VARCHAR(100)"`
	Subject  string `xorm:"not null VARCHAR(255)"`
	Message  string `xorm:"not null TEXT"`
	Charset  string `xorm:"not null VARCHAR(15)"`
	Htmlon   int    `xorm:"not null default 0 TINYINT(1)"`
	Level    int    `xorm:"not null default 1 index(level) TINYINT(1)"`
	Dateline int    `xorm:"not null default 0 INT(10)"`
	Failures int    `xorm:"not null default 0 index(level) TINYINT(3)"`
	Appid    int    `xorm:"not null default 0 index SMALLINT(6)"`
}
