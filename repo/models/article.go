package models

type Article struct {
	ArticleId      int    `xorm:"not null pk autoincr comment('日志自增ID号') SMALLINT(5)"`
	ArticleName    string `xorm:"not null comment('文章名称') VARCHAR(128)"`
	ArticleTime    int    `xorm:"not null comment('发布时间') INT(13)"`
	ArticleIp      string `xorm:"not null comment('发布IP') VARCHAR(15)"`
	ArticleClick   int    `xorm:"not null comment('查看人数') INT(10)"`
	SortArticleId  int    `xorm:"not null comment('所属分类') MEDIUMINT(8)"`
	UserId         int    `xorm:"not null comment('所属用户ID') MEDIUMINT(8)"`
	TypeId         int    `xorm:"not null default 1 comment('栏目ID') TINYINT(3)"`
	ArticleType    int    `xorm:"not null default 1 comment('文章的模式:0为私有，1为公开，2为仅好友查看') INT(13)"`
	ArticleContent string `xorm:"not null comment('文章内容') TEXT"`
	ArticleUp      int    `xorm:"not null default 0 comment('是否置顶:0为否，1为是') TINYINT(3)"`
	ArticleSupport int    `xorm:"not null default 0 comment('是否博主推荐:0为否，1为是') TINYINT(3)"`
}
