package models

type SecretMessage struct {
	SecretId       int    `xorm:"not null pk autoincr comment('自增私信ID') MEDIUMINT(8)"`
	SendId         int    `xorm:"not null comment('发信者ID') MEDIUMINT(8)"`
	ReceiveId      int    `xorm:"not null comment('收信者ID') MEDIUMINT(8)"`
	MessageTopic   string `xorm:"not null comment('私信标题') VARCHAR(64)"`
	MessageContent string `xorm:"not null comment('私信内容') VARCHAR(255)"`
}
