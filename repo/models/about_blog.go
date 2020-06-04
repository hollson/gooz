package models

type AboutBlog struct {
	BlogId          int    `xorm:"not null pk comment('用户ID') MEDIUMINT(8)"`
	BlogKeyword     string `xorm:"not null comment('博客关键字') VARCHAR(255)"`
	BlogDescription string `xorm:"not null comment('博客描述') VARCHAR(255)"`
	BlogName        string `xorm:"not null comment('博客名称') VARCHAR(36)"`
	BlogTitle       string `xorm:"not null comment('博客标题') VARCHAR(128)"`
}
