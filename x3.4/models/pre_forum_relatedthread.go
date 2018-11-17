package models

type PreForumRelatedthread struct {
	Tid            int    `xorm:"not null pk default 0 MEDIUMINT(8)"`
	Type           string `xorm:"not null pk default ''general'' ENUM('general','trade')"`
	Expiration     int    `xorm:"not null default 0 INT(10)"`
	Keywords       string `xorm:"not null default '''' VARCHAR(255)"`
	Relatedthreads string `xorm:"not null TEXT"`
}

func (t *PreForumRelatedthread) TableName() string {
	return "pre_forum_relatedthread"
}
