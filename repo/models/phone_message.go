package models

type PhoneMessage struct {
	PhoneId  int    `xorm:"not null pk autoincr comment('自增ID号') MEDIUMINT(8)"`
	PhoneNum string `xorm:"not null comment('用户手机号码') VARCHAR(12)"`
	Contents string `xorm:"not null comment('发送内容') VARCHAR(255)"`
	SendTime int    `xorm:"not null comment('发送时间') INT(13)"`
	UserId   int    `xorm:"not null comment('用户ID') MEDIUMINT(8)"`
}
