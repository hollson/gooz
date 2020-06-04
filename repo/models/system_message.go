package models

type SystemMessage struct {
	SystemId      int    `xorm:"not null pk autoincr comment('系统通知ID') MEDIUMINT(8)"`
	SendId        int    `xorm:"not null comment('接受者ID') MEDIUMINT(8)"`
	GroupId       int    `xorm:"not null comment('用户组ID') TINYINT(3)"`
	SendDefault   int    `xorm:"not null comment('1时发送所有用户，0时则不采用') MEDIUMINT(8)"`
	SystemTopic   string `xorm:"not null comment('通知内容') VARCHAR(60)"`
	SystemContent string `xorm:"not null comment('通知内容') VARCHAR(255)"`
}
