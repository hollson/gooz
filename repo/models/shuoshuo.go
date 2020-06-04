package models

type Shuoshuo struct {
	ShuoId   int    `xorm:"not null pk autoincr comment('说说记录ID') MEDIUMINT(8)"`
	UserId   int    `xorm:"not null comment('用户ID') MEDIUMINT(8)"`
	ShuoTime int    `xorm:"not null default 0 comment('发布时间') INT(13)"`
	ShuoIp   string `xorm:"not null comment('说说发布时的IP地址') VARCHAR(15)"`
	Shuoshuo string `xorm:"not null comment('说说内容') VARCHAR(255)"`
	TypeId   int    `xorm:"not null default 3 comment('栏目ID,默认为3') TINYINT(3)"`
}
