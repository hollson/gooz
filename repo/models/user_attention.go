package models

type UserAttention struct {
	AId         int `xorm:"not null pk autoincr comment('自增ID') SMALLINT(5)"`
	UserId      int `xorm:"not null comment('用户ID') MEDIUMINT(8)"`
	AttentionId int `xorm:"not null comment('关注ID') MEDIUMINT(8)"`
}
