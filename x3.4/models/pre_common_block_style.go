package models

type PreCommonBlockStyle struct {
	Styleid    int    `xorm:"not null pk autoincr SMALLINT(6)"`
	Blockclass string `xorm:"not null default '''' index VARCHAR(255)"`
	Name       string `xorm:"not null default '''' VARCHAR(255)"`
	Template   string `xorm:"not null TEXT"`
	Hash       string `xorm:"not null default '''' index VARCHAR(255)"`
	Getpic     int    `xorm:"not null default 0 TINYINT(1)"`
	Getsummary int    `xorm:"not null default 0 TINYINT(1)"`
	Makethumb  int    `xorm:"not null default 0 TINYINT(1)"`
	Settarget  int    `xorm:"not null default 0 TINYINT(1)"`
	Fields     string `xorm:"not null TEXT"`
	Moreurl    int    `xorm:"not null default 0 TINYINT(1)"`
}
