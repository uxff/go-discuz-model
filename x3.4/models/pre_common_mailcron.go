package models

type PreCommonMailcron struct {
	Cid      int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Touid    int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Email    string `xorm:"not null default '''' VARCHAR(100)"`
	Sendtime int    `xorm:"not null default 0 index INT(10)"`
}

func (t *PreCommonMailcron) TableName() string {
	return "pre_common_mailcron"
}
