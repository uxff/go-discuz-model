package models

type PreCommonUsergroupField struct {
	Groupid                int    `xorm:"not null pk SMALLINT(6)"`
	Readaccess             int    `xorm:"not null default 0 TINYINT(3)"`
	Allowpost              int    `xorm:"not null default 0 TINYINT(1)"`
	Allowreply             int    `xorm:"not null default 0 TINYINT(1)"`
	Allowpostpoll          int    `xorm:"not null default 0 TINYINT(1)"`
	Allowpostreward        int    `xorm:"not null default 0 TINYINT(1)"`
	Allowposttrade         int    `xorm:"not null default 0 TINYINT(1)"`
	Allowpostactivity      int    `xorm:"not null default 0 TINYINT(1)"`
	Allowdirectpost        int    `xorm:"not null default 0 TINYINT(1)"`
	Allowgetattach         int    `xorm:"not null default 0 TINYINT(1)"`
	Allowgetimage          int    `xorm:"not null default 0 TINYINT(1)"`
	Allowpostattach        int    `xorm:"not null default 0 TINYINT(1)"`
	Allowpostimage         int    `xorm:"not null default 0 TINYINT(1)"`
	Allowvote              int    `xorm:"not null default 0 TINYINT(1)"`
	Allowsearch            int    `xorm:"not null default 0 TINYINT(1)"`
	Allowcstatus           int    `xorm:"not null default 0 TINYINT(1)"`
	Allowinvisible         int    `xorm:"not null default 0 TINYINT(1)"`
	Allowtransfer          int    `xorm:"not null default 0 TINYINT(1)"`
	Allowsetreadperm       int    `xorm:"not null default 0 TINYINT(1)"`
	Allowsetattachperm     int    `xorm:"not null default 0 TINYINT(1)"`
	Allowposttag           int    `xorm:"not null default 0 TINYINT(1)"`
	Allowhidecode          int    `xorm:"not null default 0 TINYINT(1)"`
	Allowhtml              int    `xorm:"not null default 0 TINYINT(1)"`
	Allowanonymous         int    `xorm:"not null default 0 TINYINT(1)"`
	Allowsigbbcode         int    `xorm:"not null default 0 TINYINT(1)"`
	Allowsigimgcode        int    `xorm:"not null default 0 TINYINT(1)"`
	Allowmagics            int    `xorm:"not null TINYINT(1)"`
	Disableperiodctrl      int    `xorm:"not null default 0 TINYINT(1)"`
	Reasonpm               int    `xorm:"not null default 0 TINYINT(1)"`
	Maxprice               int    `xorm:"not null default 0 SMALLINT(6)"`
	Maxsigsize             int    `xorm:"not null default 0 SMALLINT(6)"`
	Maxattachsize          int    `xorm:"not null default 0 INT(10)"`
	Maxsizeperday          int    `xorm:"not null default 0 INT(10)"`
	Maxthreadsperhour      int    `xorm:"not null default 0 TINYINT(3)"`
	Maxpostsperhour        int    `xorm:"not null default 0 TINYINT(3)"`
	Attachextensions       string `xorm:"not null default '''' CHAR(100)"`
	Raterange              string `xorm:"not null default '''' CHAR(150)"`
	Loginreward            string `xorm:"not null default '''' CHAR(150)"`
	Mintradeprice          int    `xorm:"not null default 1 SMALLINT(6)"`
	Maxtradeprice          int    `xorm:"not null default 0 SMALLINT(6)"`
	Minrewardprice         int    `xorm:"not null default 1 SMALLINT(6)"`
	Maxrewardprice         int    `xorm:"not null default 0 SMALLINT(6)"`
	Magicsdiscount         int    `xorm:"not null TINYINT(1)"`
	Maxmagicsweight        int    `xorm:"not null SMALLINT(6)"`
	Allowpostdebate        int    `xorm:"not null default 0 TINYINT(1)"`
	Tradestick             int    `xorm:"not null TINYINT(1)"`
	Exempt                 int    `xorm:"not null TINYINT(1)"`
	Maxattachnum           int    `xorm:"not null default 0 SMALLINT(6)"`
	Allowposturl           int    `xorm:"not null default 3 TINYINT(1)"`
	Allowrecommend         int    `xorm:"not null default 1 TINYINT(1)"`
	Allowpostrushreply     int    `xorm:"not null default 0 TINYINT(1)"`
	Maxfriendnum           int    `xorm:"not null default 0 SMALLINT(6)"`
	Maxspacesize           int    `xorm:"not null default 0 INT(10)"`
	Allowcomment           int    `xorm:"not null default 0 TINYINT(1)"`
	Allowcommentarticle    int    `xorm:"not null default 0 SMALLINT(6)"`
	Searchinterval         int    `xorm:"not null default 0 SMALLINT(6)"`
	Searchignore           int    `xorm:"not null default 0 TINYINT(1)"`
	Allowblog              int    `xorm:"not null default 0 TINYINT(1)"`
	Allowdoing             int    `xorm:"not null default 0 TINYINT(1)"`
	Allowupload            int    `xorm:"not null default 0 TINYINT(1)"`
	Allowshare             int    `xorm:"not null default 0 TINYINT(1)"`
	Allowblogmod           int    `xorm:"not null default 0 TINYINT(1)"`
	Allowdoingmod          int    `xorm:"not null default 0 TINYINT(1)"`
	Allowuploadmod         int    `xorm:"not null default 0 TINYINT(1)"`
	Allowsharemod          int    `xorm:"not null default 0 TINYINT(1)"`
	Allowcss               int    `xorm:"not null default 0 TINYINT(1)"`
	Allowpoke              int    `xorm:"not null default 0 TINYINT(1)"`
	Allowfriend            int    `xorm:"not null default 0 TINYINT(1)"`
	Allowclick             int    `xorm:"not null default 0 TINYINT(1)"`
	Allowmagic             int    `xorm:"not null default 0 TINYINT(1)"`
	Allowstat              int    `xorm:"not null default 0 TINYINT(1)"`
	Allowstatdata          int    `xorm:"not null default 0 TINYINT(1)"`
	Videophotoignore       int    `xorm:"not null default 0 TINYINT(1)"`
	Allowviewvideophoto    int    `xorm:"not null default 0 TINYINT(1)"`
	Allowmyop              int    `xorm:"not null default 0 TINYINT(1)"`
	Magicdiscount          int    `xorm:"not null default 0 TINYINT(1)"`
	Domainlength           int    `xorm:"not null default 0 SMALLINT(6)"`
	Seccode                int    `xorm:"not null default 1 TINYINT(1)"`
	Disablepostctrl        int    `xorm:"not null default 0 TINYINT(1)"`
	Allowbuildgroup        int    `xorm:"not null default 0 TINYINT(1)"`
	Allowgroupdirectpost   int    `xorm:"not null default 0 TINYINT(1)"`
	Allowgroupposturl      int    `xorm:"not null default 0 TINYINT(1)"`
	Edittimelimit          int    `xorm:"not null default 0 SMALLINT(6)"`
	Allowpostarticle       int    `xorm:"not null default 0 TINYINT(1)"`
	Allowdownlocalimg      int    `xorm:"not null default 0 TINYINT(1)"`
	Allowdownremoteimg     int    `xorm:"not null default 0 TINYINT(1)"`
	Allowpostarticlemod    int    `xorm:"not null default 0 TINYINT(1)"`
	Allowspacediyhtml      int    `xorm:"not null default 0 TINYINT(1)"`
	Allowspacediybbcode    int    `xorm:"not null default 0 TINYINT(1)"`
	Allowspacediyimgcode   int    `xorm:"not null default 0 TINYINT(1)"`
	Allowcommentpost       int    `xorm:"not null default 2 TINYINT(1)"`
	Allowcommentitem       int    `xorm:"not null default 0 TINYINT(1)"`
	Allowcommentreply      int    `xorm:"not null default 0 TINYINT(1)"`
	Allowreplycredit       int    `xorm:"not null default 0 TINYINT(1)"`
	Ignorecensor           int    `xorm:"not null default 0 TINYINT(1)"`
	Allowsendallpm         int    `xorm:"not null default 0 TINYINT(1)"`
	Allowsendpmmaxnum      int    `xorm:"not null default 0 SMALLINT(6)"`
	Maximagesize           int    `xorm:"not null default 0 MEDIUMINT(8)"`
	Allowmediacode         int    `xorm:"not null default 0 TINYINT(1)"`
	Allowbegincode         int    `xorm:"not null default 0 TINYINT(1)"`
	Allowat                int    `xorm:"not null default 0 SMALLINT(6)"`
	Allowsetpublishdate    int    `xorm:"not null default 0 TINYINT(1)"`
	Allowfollowcollection  int    `xorm:"not null default 0 TINYINT(1)"`
	Allowcommentcollection int    `xorm:"not null default 0 TINYINT(1)"`
	Allowcreatecollection  int    `xorm:"not null default 0 SMALLINT(6)"`
	Forcesecques           int    `xorm:"not null default 0 TINYINT(1)"`
	Forcelogin             int    `xorm:"not null default 0 TINYINT(1)"`
	Closead                int    `xorm:"not null default 0 TINYINT(1)"`
	Buildgroupcredits      int    `xorm:"not null default 0 SMALLINT(6)"`
	Allowimgcontent        int    `xorm:"not null default 0 TINYINT(1)"`
}

func (t *PreCommonUsergroupField) TableName() string {
	return "pre_common_usergroup_field"
}
