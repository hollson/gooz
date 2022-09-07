// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jwt

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var JWT_STATUS = map[int]string{
	200: "OK",
	400: "Token is required",
	401: "Token verification failed",
	504: "Token verification timeout",
}

// Gin中间件：
//	验证JWT的中间件
// 	从Http头部的Authorization(忽略大小写)中获取Token信息，并验证Token的有效性。
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code = 200
		token := c.GetHeader("Authorization")
		if len(token) == 0 {
			code = 400
		} else {
			_, err := Resolve(token)
			if err != nil {
				logrus.Errorln(err)
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = 504 // 超时
				default:
					code = 511 // 失败
				}
			}
		}

		if code != http.StatusOK {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  JWT_STATUS[code],
			})
			c.Abort()
			return
		}
		// c.Set("token", token)  //继续传递
		c.Next()
	}
}
