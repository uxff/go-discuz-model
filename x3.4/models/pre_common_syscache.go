package models

type PreCommonSyscache struct {
	Cname    string `xorm:"not null pk VARCHAR(32)"`
	Ctype    int    `xorm:"not null TINYINT(3)"`
	Dateline int    `xorm:"not null INT(10)"`
	Data     []byte `xorm:"not null MEDIUMBLOB"`
}

func (t *PreCommonSyscache) TableName() string {
	return "pre_common_syscache"
}
