package models

type Friend struct {
	FId      int `xorm:"not null pk autoincr comment('自增ID') SMALLINT(5)"`
	UserId   int `xorm:"not null comment('用户ID') MEDIUMINT(8)"`
	FriendId int `xorm:"not null comment('好友ID') MEDIUMINT(8)"`
}
