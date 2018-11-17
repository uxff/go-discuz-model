package models

type PreCommonMemberMedal struct {
	Uid     int `xorm:"not null pk MEDIUMINT(8)"`
	Medalid int `xorm:"not null pk SMALLINT(6)"`
}
