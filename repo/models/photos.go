package models

type Photos struct {
	PhotoId          int    `xorm:"not null pk autoincr comment('相片ID') MEDIUMINT(8)"`
	PhotoName        string `xorm:"not null comment('相片名称') VARCHAR(255)"`
	PhotoSrc         string `xorm:"not null comment('图片路径') VARCHAR(255)"`
	PhotoDescription string `xorm:"not null comment('图片描述') VARCHAR(255)"`
	UserId           int    `xorm:"not null comment('所属用户ID') MEDIUMINT(8)"`
	SortId           int    `xorm:"not null comment('所属相册ID') MEDIUMINT(8)"`
	UploadTime       int    `xorm:"not null comment('图片上传时间') INT(13)"`
	UploadIp         string `xorm:"not null comment('图片操作上传IP地址') VARCHAR(15)"`
}
