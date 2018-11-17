package models

type PreForumThreadimage struct {
	Tid        int    `xorm:"not null default 0 index MEDIUMINT(8)"`
	Attachment string `xorm:"not null default '''' VARCHAR(255)"`
	Remote     int    `xorm:"not null default 0 TINYINT(1)"`
}

func (t *PreForumThreadimage) TableName() string {
	return "pre_forum_threadimage"
}
