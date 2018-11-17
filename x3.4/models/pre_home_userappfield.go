package models

type PreHomeUserappfield struct {
	Uid         int    `xorm:"not null default 0 index(uid) MEDIUMINT(8)"`
	Appid       int    `xorm:"not null default 0 index(uid) MEDIUMINT(8)"`
	Profilelink string `xorm:"not null TEXT"`
	Myml        string `xorm:"not null TEXT"`
}

func (t *PreHomeUserappfield) TableName() string {
	return "pre_home_userappfield"
}
