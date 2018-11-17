package models

type PreCommonSecquestion struct {
	Id       int    `xorm:"not null pk autoincr SMALLINT(6)"`
	Type     int    `xorm:"not null TINYINT(1)"`
	Question string `xorm:"not null TEXT"`
	Answer   string `xorm:"not null VARCHAR(255)"`
}
