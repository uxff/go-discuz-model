package models

type PreUcenterTags struct {
	Tagname    string `xorm:"not null index(tagname) CHAR(20)"`
	Appid      int    `xorm:"not null default 0 index(tagname) SMALLINT(6)"`
	Data       string `xorm:"default 'NULL' MEDIUMTEXT"`
	Expiration int    `xorm:"not null INT(10)"`
}

func (t *PreUcenterTags) TableName() string {
	return "pre_ucenter_tags"
}
