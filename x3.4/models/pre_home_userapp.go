package models

type PreHomeUserapp struct {
	Uid              int    `xorm:"not null default 0 index(displayorder) index(menuorder) index(uid) MEDIUMINT(8)"`
	Appid            int    `xorm:"not null default 0 index(uid) MEDIUMINT(8)"`
	Appname          string `xorm:"not null default '''' VARCHAR(60)"`
	Privacy          int    `xorm:"not null default 0 TINYINT(1)"`
	Allowsidenav     int    `xorm:"not null default 0 TINYINT(1)"`
	Allowfeed        int    `xorm:"not null default 0 TINYINT(1)"`
	Allowprofilelink int    `xorm:"not null default 0 TINYINT(1)"`
	Narrow           int    `xorm:"not null default 0 TINYINT(1)"`
	Menuorder        int    `xorm:"not null default 0 index(menuorder) SMALLINT(6)"`
	Displayorder     int    `xorm:"not null default 0 index(displayorder) SMALLINT(6)"`
}
