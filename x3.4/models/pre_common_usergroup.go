package models

type PreCommonUsergroup struct {
	Groupid         int    `xorm:"not null pk autoincr SMALLINT(6)"`
	Radminid        int    `xorm:"not null default 0 TINYINT(3)"`
	Type            string `xorm:"not null default ''member'' ENUM('member','special','system')"`
	System          string `xorm:"not null default ''private'' VARCHAR(255)"`
	Grouptitle      string `xorm:"not null default '''' VARCHAR(255)"`
	Creditshigher   int    `xorm:"not null default 0 index(creditsrange) INT(10)"`
	Creditslower    int    `xorm:"not null default 0 index(creditsrange) INT(10)"`
	Stars           int    `xorm:"not null default 0 TINYINT(3)"`
	Color           string `xorm:"not null default '''' VARCHAR(255)"`
	Icon            string `xorm:"not null default '''' VARCHAR(255)"`
	Allowvisit      int    `xorm:"not null default 0 TINYINT(1)"`
	Allowsendpm     int    `xorm:"not null default 1 TINYINT(1)"`
	Allowinvite     int    `xorm:"not null default 0 TINYINT(1)"`
	Allowmailinvite int    `xorm:"not null default 0 TINYINT(1)"`
	Maxinvitenum    int    `xorm:"not null default 0 TINYINT(3)"`
	Inviteprice     int    `xorm:"not null default 0 SMALLINT(6)"`
	Maxinviteday    int    `xorm:"not null default 0 SMALLINT(6)"`
}
