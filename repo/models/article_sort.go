package models

type ArticleSort struct {
	SortArticleId   int    `xorm:"not null pk autoincr comment('文章自增ID') MEDIUMINT(8)"`
	UserId          int    `xorm:"not null comment('该分类所属用户') MEDIUMINT(8)"`
	SortArticleName string `xorm:"not null comment('分类名称') VARCHAR(60)"`
}
