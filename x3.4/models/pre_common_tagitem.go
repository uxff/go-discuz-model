package models

type PreCommonTagitem struct {
	Tagid  int    `xorm:"not null pk default 0 unique(item) MEDIUMINT(8)"`
	Itemid int    `xorm:"not null pk default 0 index(idtype) unique(item) MEDIUMINT(8)"`
	Idtype string `xorm:"not null pk default '''' index(idtype) unique(item) CHAR(10)"`
}

func (t *PreCommonTagitem) TableName() string {
	return "pre_common_tagitem"
}
