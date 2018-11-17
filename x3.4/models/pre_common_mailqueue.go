package models

type PreCommonMailqueue struct {
	Qid      int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Cid      int    `xorm:"not null default 0 index(mcid) MEDIUMINT(8)"`
	Subject  string `xorm:"not null TEXT"`
	Message  string `xorm:"not null TEXT"`
	Dateline int    `xorm:"not null default 0 index(mcid) INT(10)"`
}

func (t *PreCommonMailqueue) TableName() string {
	return "pre_common_mailqueue"
}
