// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package account

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hollson/deeplink/app/auth/jwt"
	"github.com/sirupsen/logrus"
)

// 请求参数
type ReqLogin struct {
	Account  string `json:"account"` // 用户名||手机号||邮箱
	Password string `json:"password"`
}

// 响应参数(Proto定义)
type RepArticle struct {
	Token string `json:"token"`
}

// 请求验证
func check() {
	// validate 插件
}

func LoginHandler(c *gin.Context) {
	var req = new(ReqLogin)
	err := c.BindJSON(req)
	if err != nil {
		logrus.Errorln(err.Error())
		c.String(400, "参数错误", )
		return
	}

	u, err := UserDemo(req)
	if err != nil || u == nil {
		logrus.Errorln(err, u) // todo
		c.String(http.StatusOK, "账号或密码错误")
	}

	tk, err := jwt.Generate(strconv.Itoa(int(u.Id)))
	if err != nil {
		c.String(400, err.Error())
	}
	c.Header("Authorization",fmt.Sprintf("Bearer %s",tk))
	c.Header("Location","/home/index")
	c.JSON(200, tk)
}
