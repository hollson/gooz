package models

type PowerList struct {
	PId       int    `xorm:"not null pk autoincr comment('自增ID') INT(10)"`
	PowerId   int    `xorm:"not null comment('权限ID') INT(10)"`
	PowerName string `xorm:"not null comment('权限描述') VARCHAR(36)"`
}
