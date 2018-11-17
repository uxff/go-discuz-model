package models

type PreUcenterMergemembers struct {
	Appid    int    `xorm:"not null pk SMALLINT(6)"`
	Username string `xorm:"not null pk CHAR(15)"`
}
