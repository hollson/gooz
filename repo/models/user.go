package models

type User struct {
	UserId             int    `xorm:"not null pk autoincr comment('用户ID') MEDIUMINT(8)"`
	GroupId            int    `xorm:"not null comment('用户组ID') MEDIUMINT(8)"`
	UserName           string `xorm:"not null comment('用户名') VARCHAR(32)"`
	UserPwd            string `xorm:"not null comment('用户密码') VARCHAR(32)"`
	UserPhone          int    `xorm:"not null comment('用户手机号码') INT(12)"`
	UserSex            string `xorm:"not null comment('用户性别') VARCHAR(6)"`
	UserQq             int    `xorm:"not null comment('用户QQ号码') MEDIUMINT(9)"`
	UserEmail          string `xorm:"not null comment('用户EMAIL地址') VARCHAR(64)"`
	UserAddress        string `xorm:"not null comment('用户地址') VARCHAR(255)"`
	UserMark           int    `xorm:"not null comment('用户积分') MEDIUMINT(9)"`
	UserRankId         int    `xorm:"not null comment('用户等级') TINYINT(3)"`
	UserLastLoginIp    string `xorm:"not null comment('用户上一次登录IP地址') VARCHAR(15)"`
	UserBirthday       int    `xorm:"not null comment('用户生日') INT(13)"`
	UserDescription    string `xorm:"not null comment('自我描述') VARCHAR(255)"`
	UserImageUrl       string `xorm:"not null comment('用户头像存储路径') VARCHAR(255)"`
	UserSchool         string `xorm:"not null comment('毕业学校') VARCHAR(255)"`
	UserRegisterTime   int    `xorm:"not null comment('用户注册时间') INT(13)"`
	UserRegisterIp     string `xorm:"not null comment('用户注册时IP地址') VARCHAR(15)"`
	UserLastUpdateTime int    `xorm:"not null comment('用户上次更新博客时间') INT(13)"`
	UserWeibo          string `xorm:"not null comment('用户微博') VARCHAR(255)"`
	UserBloodType      string `xorm:"not null comment('用户血型') CHAR(3)"`
	UserSays           string `xorm:"not null comment('用户语录') VARCHAR(255)"`
	UserLock           int    `xorm:"not null comment('是否锁定，0为不锁定，1为锁定') TINYINT(3)"`
	UserFreeze         int    `xorm:"not null comment('是否冻结，0为不冻结，1为冻结') TINYINT(3)"`
	UserPower          string `xorm:"not null comment('拥有权限') VARCHAR(255)"`
}
