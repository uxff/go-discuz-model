package models

type PreCommonStat struct {
	Daytime      int `xorm:"not null pk default 0 INT(10)"`
	Login        int `xorm:"not null default 0 INT(10)"`
	Mobilelogin  int `xorm:"not null default 0 INT(10)"`
	Connectlogin int `xorm:"not null default 0 INT(10)"`
	Register     int `xorm:"not null default 0 INT(10)"`
	Invite       int `xorm:"not null default 0 INT(10)"`
	Appinvite    int `xorm:"not null default 0 INT(10)"`
	Doing        int `xorm:"not null default 0 INT(10)"`
	Blog         int `xorm:"not null default 0 INT(10)"`
	Pic          int `xorm:"not null default 0 INT(10)"`
	Poll         int `xorm:"not null default 0 INT(10)"`
	Activity     int `xorm:"not null default 0 INT(10)"`
	Share        int `xorm:"not null default 0 INT(10)"`
	Thread       int `xorm:"not null default 0 INT(10)"`
	Docomment    int `xorm:"not null default 0 INT(10)"`
	Blogcomment  int `xorm:"not null default 0 INT(10)"`
	Piccomment   int `xorm:"not null default 0 INT(10)"`
	Sharecomment int `xorm:"not null default 0 INT(10)"`
	Reward       int `xorm:"not null default 0 INT(10)"`
	Debate       int `xorm:"not null default 0 INT(10)"`
	Trade        int `xorm:"not null default 0 INT(10)"`
	Group        int `xorm:"not null default 0 INT(10)"`
	Groupjoin    int `xorm:"not null default 0 INT(10)"`
	Groupthread  int `xorm:"not null default 0 INT(10)"`
	Grouppost    int `xorm:"not null default 0 INT(10)"`
	Post         int `xorm:"not null default 0 INT(10)"`
	Wall         int `xorm:"not null default 0 INT(10)"`
	Poke         int `xorm:"not null default 0 INT(10)"`
	Click        int `xorm:"not null default 0 INT(10)"`
	Sendpm       int `xorm:"not null default 0 INT(10)"`
	Friend       int `xorm:"not null default 0 INT(10)"`
	Addfriend    int `xorm:"not null default 0 INT(10)"`
}
