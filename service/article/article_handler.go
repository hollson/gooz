// ---------------------------------------------------------------------------
// @ Copyright (C) free license,without warranty of any kind .
// @ Author: hollson <hollson@live.cn>
// @ Date: 2020-05-07
// @ Version: 1.0.0
//
// 文章相关的业务处理逻辑
// ----------------------------------------------------------------------------

package article

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hollson/deeplink/repo/domain"
	"github.com/sirupsen/logrus"
)

var repo domain.Article

// // 请求参数
// type ReqArticle struct {
// 	Id int `json:"id"`
// }

// 响应参数(Proto定义)
type RepArticle struct {
	Id         int    `json:"id"`
	Title      string `json:"name"`
	Content    string `json:"content"`
	Cat        int    `json:"cat"`
	AuthorId   int    `json:"author_id"`
	AuthorName string `json:"author_name"`
	Created    string `json:"created"`
}

// 获取用户信息
func GetArticleDetailHandler(ctx *gin.Context) {
	id := ctx.DefaultQuery("id", "3")

	n, err := strconv.Atoi(id)
	if err != nil {
		logrus.Errorln(err)
		ctx.JSON(http.StatusBadRequest, "id error")
		return
	}
	repo = domain.Article{}
	pong, err := repo.GetArticleById(n)
	if err != nil {
		logrus.Errorln(err)
		return
	}

	ctx.JSON(http.StatusOK, RepArticle{
		Id:         pong.ArticleId,
		Title:      pong.ArticleName,
		Content:    pong.ArticleContent,
		Cat:        pong.ArticleType,
		AuthorId:   pong.User.UserId,
		AuthorName: pong.UserName,
		Created:    time.Unix(int64(pong.ArticleTime), 0).Format("2006-01-02 15:04:05"),
	})

}
