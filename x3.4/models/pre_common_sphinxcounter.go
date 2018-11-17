package models

type PreCommonSphinxcounter struct {
	Indexid int `xorm:"not null pk TINYINT(1)"`
	Maxid   int `xorm:"not null INT(10)"`
}

func (t *PreCommonSphinxcounter) TableName() string {
	return "pre_common_sphinxcounter"
}
