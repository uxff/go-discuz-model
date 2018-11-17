package models

type PreCommonMemberFieldHome struct {
	Uid              int    `xorm:"not null pk MEDIUMINT(8)"`
	Videophoto       string `xorm:"not null default '''' VARCHAR(255)"`
	Spacename        string `xorm:"not null default '''' VARCHAR(255)"`
	Spacedescription string `xorm:"not null default '''' VARCHAR(255)"`
	Domain           string `xorm:"not null default '''' index CHAR(15)"`
	Addsize          int    `xorm:"not null default 0 INT(10)"`
	Addfriend        int    `xorm:"not null default 0 SMALLINT(6)"`
	Menunum          int    `xorm:"not null default 0 SMALLINT(6)"`
	Theme            string `xorm:"not null default '''' VARCHAR(20)"`
	Spacecss         string `xorm:"not null TEXT"`
	Blockposition    string `xorm:"not null TEXT"`
	Recentnote       string `xorm:"not null TEXT"`
	Spacenote        string `xorm:"not null TEXT"`
	Privacy          string `xorm:"not null TEXT"`
	Feedfriend       string `xorm:"not null MEDIUMTEXT"`
	Acceptemail      string `xorm:"not null TEXT"`
	Magicgift        string `xorm:"not null TEXT"`
	Stickblogs       string `xorm:"not null TEXT"`
}

func (t *PreCommonMemberFieldHome) TableName() string {
	return "pre_common_member_field_home"
}
