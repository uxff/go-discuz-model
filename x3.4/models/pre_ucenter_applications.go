package models

type PreUcenterApplications struct {
	Appid        int    `xorm:"not null pk autoincr SMALLINT(6)"`
	Type         string `xorm:"not null default '''' VARCHAR(16)"`
	Name         string `xorm:"not null default '''' VARCHAR(20)"`
	Url          string `xorm:"not null default '''' VARCHAR(255)"`
	Authkey      string `xorm:"not null default '''' VARCHAR(255)"`
	Ip           string `xorm:"not null default '''' VARCHAR(15)"`
	Viewprourl   string `xorm:"not null VARCHAR(255)"`
	Apifilename  string `xorm:"not null default ''uc.php'' VARCHAR(30)"`
	Charset      string `xorm:"not null default '''' VARCHAR(8)"`
	Dbcharset    string `xorm:"not null default '''' VARCHAR(8)"`
	Synlogin     int    `xorm:"not null default 0 TINYINT(1)"`
	Recvnote     int    `xorm:"default 0 TINYINT(1)"`
	Extra        string `xorm:"not null TEXT"`
	Tagtemplates string `xorm:"not null TEXT"`
	Allowips     string `xorm:"not null TEXT"`
}

func (t *PreUcenterApplications) TableName() string {
	return "pre_ucenter_applications"
}
