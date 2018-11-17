package models

type PrePortalTopicPic struct {
	Picid    int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Topicid  int    `xorm:"not null default 0 index MEDIUMINT(8)"`
	Uid      int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Username string `xorm:"not null default '''' VARCHAR(15)"`
	Dateline int    `xorm:"not null default 0 INT(10)"`
	Filename string `xorm:"not null default '''' VARCHAR(255)"`
	Title    string `xorm:"not null default '''' VARCHAR(255)"`
	Size     int    `xorm:"not null default 0 INT(10)"`
	Filepath string `xorm:"not null default '''' VARCHAR(255)"`
	Thumb    int    `xorm:"not null default 0 TINYINT(1)"`
	Remote   int    `xorm:"not null default 0 TINYINT(1)"`
}
