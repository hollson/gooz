package models

type Visitor struct {
	VId         int    `xorm:"not null pk autoincr comment('访客记录ID') MEDIUMINT(8)"`
	VisitorId   int    `xorm:"not null comment('访客ID') MEDIUMINT(8)"`
	VisitorTime int    `xorm:"not null comment('来访时间') INT(13)"`
	UserId      int    `xorm:"not null comment('被访用户ID') MEDIUMINT(8)"`
	VisitorIp   string `xorm:"not null comment('访客IP地址') VARCHAR(15)"`
	TypeId      int    `xorm:"not null comment('访问板块ID') INT(3)"`
	WhereId     int    `xorm:"not null comment('查看某板块的某个子项目，如查看相册板块的第3个相册，该ID对应该相册的ID号') MEDIUMINT(8)"`
}
