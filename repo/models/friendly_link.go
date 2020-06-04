package models

type FriendlyLink struct {
	LinkId    int    `xorm:"not null pk autoincr comment('友情链接自增ID') SMALLINT(5)"`
	LinkName  string `xorm:"not null comment('友情链接名称') VARCHAR(60)"`
	LinkUrl   string `xorm:"not null comment('链接地址') VARCHAR(255)"`
	LinkLogo  string `xorm:"not null comment('LOGO图片') VARCHAR(255)"`
	ShowOrder int    `xorm:"not null comment('在页面显示的顺序') TINYINT(3)"`
}
