package models

type Ad struct {
	AdId       int    `xorm:"not null pk autoincr comment('自增ID') SMALLINT(5)"`
	PositionId int    `xorm:"not null comment('0,站外广告;从1开始代表的是该广告所处的广告位,同表ad_postition中的字段position_id的值') SMALLINT(5)"`
	MediaType  int    `xorm:"not null default 0 comment('广告类型,0图片;1flash;2代码3文字') TINYINT(3)"`
	AdName     string `xorm:"not null comment('该条广告记录的广告名称') VARCHAR(60)"`
	AdLink     string `xorm:"not null comment('广告链接地址') VARCHAR(255)"`
	AdCode     string `xorm:"not null comment('广告链接的表现,文字广告就是文字或图片和flash就是它们的地址') TEXT"`
	StartTime  int    `xorm:"not null default 0 comment('广告开始时间') INT(13)"`
	EndTime    int    `xorm:"not null default 0 comment('广告结束时间') INT(13)"`
	LinkMan    string `xorm:"not null comment('广告联系人') VARCHAR(60)"`
	LinkEmail  string `xorm:"not null comment('广告联系人的邮箱') VARCHAR(60)"`
	LinkPhone  string `xorm:"not null comment('广告联系人得电话') VARCHAR(60)"`
	ClickCount int    `xorm:"not null default 0 comment('广告点击次数') MEDIUMINT(8)"`
	Enabled    int    `xorm:"not null default 1 comment('该广告是否关闭;1开启; 0关闭; 关闭后广告将不再有效') TINYINT(3)"`
}
