package models

type UserGroup struct {
	GId        int    `xorm:"not null pk autoincr comment('自增ID号') TINYINT(3)"`
	GroupId    int    `xorm:"not null comment('用户组ID') TINYINT(3)"`
	GroupName  string `xorm:"not null comment('用户组名') VARCHAR(20)"`
	GroupPower string `xorm:"not null comment('用户权限') VARCHAR(20)"`
}
