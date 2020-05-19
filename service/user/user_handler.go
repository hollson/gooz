// ---------------------------------------------------------------------------
// @ Copyright (C) free license,without warranty of any kind .
// @ Author: hollson <hollson@live.cn>
// @ Date: 2020-05-07
// @ Version: 1.0.0
//
// 用户相关的业务处理逻辑
// ----------------------------------------------------------------------------

package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hollson/deeplink/repo/domain"
	"github.com/sirupsen/logrus"
)

var repo domain.IUser

// 请求参数
type ReqUser struct {
	Id int `json:"id"`
}

// 响应参数(Proto定义)
type RepUser struct {
	Id   int  `json:"id"`
	Name string `json:"name"`
}

// 获取用户信息
func GetUserHandler(ctx *gin.Context) {
	logrus.Infoln("接收到请求...")
	// req := &ReqUser{}
	// if err := ctx.ShouldBind(req); err != nil {
	//	ctx.String(http.StatusBadRequest, err.Error())
	//	ctx.Abort()
	// }
	id := ctx.DefaultQuery("id", "0")
	n, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "id error")
		return
	}

	// 逻辑处理
	repo = &domain.UserRepo{}
	u, err := repo.GetUser(int64(n))
	if err != nil {
		logrus.Errorln(err)
		return
	} else {
		ctx.JSON(http.StatusOK, RepUser{
			Id:   u.Id,
			Name: u.Name,
		})
	}
}
