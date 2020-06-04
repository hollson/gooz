package models

type UserComment struct {
	CId           int    `xorm:"not null pk autoincr comment('评论自增ID号') MEDIUMINT(8)"`
	UserId        int    `xorm:"not null comment('收到评论的用户ID') MEDIUMINT(8)"`
	TypeId        int    `xorm:"not null comment('评论栏目ID') TINYINT(3)"`
	CommitId      int    `xorm:"not null comment('评论内容的ID') MEDIUMINT(8)"`
	CommitContent string `xorm:"not null comment('评论内容') VARCHAR(255)"`
	CommitUserId  int    `xorm:"not null comment('评论者ID') MEDIUMINT(8)"`
	CommitTime    int    `xorm:"not null comment('评论时间') INT(13)"`
	CommitIp      string `xorm:"not null comment('评论时的IP地址') VARCHAR(15)"`
}
