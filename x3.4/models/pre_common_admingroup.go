package models

type PreCommonAdmingroup struct {
	Admingid              int `xorm:"not null pk default 0 SMALLINT(6)"`
	Alloweditpost         int `xorm:"not null default 0 TINYINT(1)"`
	Alloweditpoll         int `xorm:"not null default 0 TINYINT(1)"`
	Allowstickthread      int `xorm:"not null default 0 TINYINT(1)"`
	Allowmodpost          int `xorm:"not null default 0 TINYINT(1)"`
	Allowdelpost          int `xorm:"not null default 0 TINYINT(1)"`
	Allowmassprune        int `xorm:"not null default 0 TINYINT(1)"`
	Allowrefund           int `xorm:"not null default 0 TINYINT(1)"`
	Allowcensorword       int `xorm:"not null default 0 TINYINT(1)"`
	Allowviewip           int `xorm:"not null default 0 TINYINT(1)"`
	Allowbanip            int `xorm:"not null default 0 TINYINT(1)"`
	Allowedituser         int `xorm:"not null default 0 TINYINT(1)"`
	Allowmoduser          int `xorm:"not null default 0 TINYINT(1)"`
	Allowbanuser          int `xorm:"not null default 0 TINYINT(1)"`
	Allowbanvisituser     int `xorm:"not null default 0 TINYINT(1)"`
	Allowpostannounce     int `xorm:"not null default 0 TINYINT(1)"`
	Allowviewlog          int `xorm:"not null default 0 TINYINT(1)"`
	Allowbanpost          int `xorm:"not null default 0 TINYINT(1)"`
	SupeAllowpushthread   int `xorm:"not null default 0 TINYINT(1)"`
	Allowhighlightthread  int `xorm:"not null default 0 TINYINT(1)"`
	Allowlivethread       int `xorm:"not null default 0 TINYINT(1)"`
	Allowdigestthread     int `xorm:"not null default 0 TINYINT(1)"`
	Allowrecommendthread  int `xorm:"not null default 0 TINYINT(1)"`
	Allowbumpthread       int `xorm:"not null default 0 TINYINT(1)"`
	Allowclosethread      int `xorm:"not null default 0 TINYINT(1)"`
	Allowmovethread       int `xorm:"not null default 0 TINYINT(1)"`
	Allowedittypethread   int `xorm:"not null default 0 TINYINT(1)"`
	Allowstampthread      int `xorm:"not null default 0 TINYINT(1)"`
	Allowstamplist        int `xorm:"not null default 0 TINYINT(1)"`
	Allowcopythread       int `xorm:"not null default 0 TINYINT(1)"`
	Allowmergethread      int `xorm:"not null default 0 TINYINT(1)"`
	Allowsplitthread      int `xorm:"not null default 0 TINYINT(1)"`
	Allowrepairthread     int `xorm:"not null default 0 TINYINT(1)"`
	Allowwarnpost         int `xorm:"not null default 0 TINYINT(1)"`
	Allowviewreport       int `xorm:"not null default 0 TINYINT(1)"`
	Alloweditforum        int `xorm:"not null default 0 TINYINT(1)"`
	Allowremovereward     int `xorm:"not null default 0 TINYINT(1)"`
	Allowedittrade        int `xorm:"not null default 0 TINYINT(1)"`
	Alloweditactivity     int `xorm:"not null default 0 TINYINT(1)"`
	Allowstickreply       int `xorm:"not null default 0 TINYINT(1)"`
	Allowmanagearticle    int `xorm:"not null default 0 TINYINT(1)"`
	Allowaddtopic         int `xorm:"not null default 0 TINYINT(1)"`
	Allowmanagetopic      int `xorm:"not null default 0 TINYINT(1)"`
	Allowdiy              int `xorm:"not null default 0 TINYINT(1)"`
	Allowclearrecycle     int `xorm:"not null default 0 TINYINT(1)"`
	Allowmanagetag        int `xorm:"not null default 0 TINYINT(1)"`
	Alloweditusertag      int `xorm:"not null default 0 TINYINT(1)"`
	Managefeed            int `xorm:"not null default 0 TINYINT(1)"`
	Managedoing           int `xorm:"not null default 0 TINYINT(1)"`
	Manageshare           int `xorm:"not null default 0 TINYINT(1)"`
	Manageblog            int `xorm:"not null default 0 TINYINT(1)"`
	Managealbum           int `xorm:"not null default 0 TINYINT(1)"`
	Managecomment         int `xorm:"not null default 0 TINYINT(1)"`
	Managemagiclog        int `xorm:"not null default 0 TINYINT(1)"`
	Managereport          int `xorm:"not null default 0 TINYINT(1)"`
	Managehotuser         int `xorm:"not null default 0 TINYINT(1)"`
	Managedefaultuser     int `xorm:"not null default 0 TINYINT(1)"`
	Managevideophoto      int `xorm:"not null default 0 TINYINT(1)"`
	Managemagic           int `xorm:"not null default 0 TINYINT(1)"`
	Manageclick           int `xorm:"not null default 0 TINYINT(1)"`
	Allowmanagecollection int `xorm:"not null default 0 TINYINT(1)"`
	Allowmakehtml         int `xorm:"not null default 0 TINYINT(1)"`
}

func (t *PreCommonAdmingroup) TableName() string {
	return "pre_common_admingroup"
}
