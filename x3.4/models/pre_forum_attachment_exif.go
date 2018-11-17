package models

type PreForumAttachmentExif struct {
	Aid  int    `xorm:"not null pk MEDIUMINT(8)"`
	Exif string `xorm:"not null TEXT"`
}
