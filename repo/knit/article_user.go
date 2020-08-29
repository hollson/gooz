// -----------------------------------------------------------------------
// @ Copyright (C) free license,without warranty of any kind .
// @ Author: hollson <hollson@live.cn>
// @ Date: 2020-05-18
// @ Version: 1.0.0
//
// UserClass:自定义的Join联合对象
// -----------------------------------------------------------------------

package knit

import (
	"github.com/hollson/gooz/repo/models"
)

// 自定义的Join联合对象
type ArticleUser struct {
	models.Article `xorm:"extends"`
	models.User    `xorm:"extends"`
}

func (ArticleUser) TableName() string {
	//指定使用该结构体对象 进行数据库查询时，使用的表名，这里返回子表名称
	return "article_user"
}
