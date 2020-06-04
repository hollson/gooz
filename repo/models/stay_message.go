package models

type StayMessage struct {
	StayId          int    `xorm:"not null pk autoincr comment('留言表自增ID') SMALLINT(5)"`
	UserId          int    `xorm:"not null comment('用户ID') MEDIUMINT(8)"`
	StayUserId      int    `xorm:"not null comment('留言者ID') MEDIUMINT(8)"`
	MessageContent  string `xorm:"not null comment('留言内容') VARCHAR(255)"`
	StayUserIp      string `xorm:"not null comment('留言用户的IP地址') VARCHAR(15)"`
	MessageStayTime int    `xorm:"not null comment('留言时间') INT(13)"`
	Place           string `xorm:"not null comment('地区') VARCHAR(64)"`
}
