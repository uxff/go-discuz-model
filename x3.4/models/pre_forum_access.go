package models

type PreForumAccess struct {
	Uid             int `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Fid             int `xorm:"not null pk default 0 index(listorder) MEDIUMINT(8)"`
	Allowview       int `xorm:"not null default 0 TINYINT(1)"`
	Allowpost       int `xorm:"not null default 0 TINYINT(1)"`
	Allowreply      int `xorm:"not null default 0 TINYINT(1)"`
	Allowgetattach  int `xorm:"not null default 0 TINYINT(1)"`
	Allowgetimage   int `xorm:"not null default 0 TINYINT(1)"`
	Allowpostattach int `xorm:"not null default 0 TINYINT(1)"`
	Allowpostimage  int `xorm:"not null default 0 TINYINT(1)"`
	Adminuser       int `xorm:"not null default 0 MEDIUMINT(8)"`
	Dateline        int `xorm:"not null default 0 index(listorder) INT(10)"`
}
