package models

type PhotoSort struct {
	SortImgId       int    `xorm:"not null pk autoincr comment('相册ID') MEDIUMINT(8)"`
	SortImgName     string `xorm:"not null comment('相册名') VARCHAR(20)"`
	SortImgType     string `xorm:"not null comment('展示方式 0->仅主人可见,1->输入密码即可查看,2->仅好友能查看,3->回答问题即可查看') VARCHAR(20)"`
	ImgPassword     string `xorm:"not null comment('查看密码') VARCHAR(32)"`
	UserId          int    `xorm:"not null comment('所属用户ID') MEDIUMINT(8)"`
	ImgSortQuestion string `xorm:"not null comment('访问问题') VARCHAR(255)"`
	ImgSortAnswer   string `xorm:"not null comment('访问问题的答案') VARCHAR(128)"`
	TypeId          int    `xorm:"not null default 1 comment('默认1表示相册板块') INT(3)"`
	TopPicSrc       int    `xorm:"not null comment('封面图片的路径') MEDIUMINT(8)"`
}
