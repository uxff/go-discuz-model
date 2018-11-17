package models

type PreCommonMemberVerify struct {
	Uid     int `xorm:"not null pk MEDIUMINT(8)"`
	Verify1 int `xorm:"not null default 0 index TINYINT(1)"`
	Verify2 int `xorm:"not null default 0 index TINYINT(1)"`
	Verify3 int `xorm:"not null default 0 index TINYINT(1)"`
	Verify4 int `xorm:"not null default 0 index TINYINT(1)"`
	Verify5 int `xorm:"not null default 0 index TINYINT(1)"`
	Verify6 int `xorm:"not null default 0 index TINYINT(1)"`
	Verify7 int `xorm:"not null default 0 index TINYINT(1)"`
}

func (t *PreCommonMemberVerify) TableName() string {
	return "pre_common_member_verify"
}
