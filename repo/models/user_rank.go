package models

type UserRank struct {
	RankId     int    `xorm:"not null pk autoincr comment('自增ID') MEDIUMINT(5)"`
	UserRankId int    `xorm:"not null comment('等级ID') SMALLINT(5)"`
	RankMark   int    `xorm:"not null comment('等级积分') MEDIUMINT(6)"`
	RankName   string `xorm:"not null comment('等级名称') VARCHAR(32)"`
}
