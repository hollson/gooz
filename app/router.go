//------------------------------------------------------------------
// @ Author: hollson <hollson@live.com>
// @ Date: 2019-12-01
// @ Version: 1.0.0
//------------------------------------------------------------------

// 注册路由
package app

import (
	"github.com/hollson/deeplink/handler/help"
	"github.com/hollson/deeplink/handler/user"
)

func Route() {
	UserRoute()
	HelpRoute()
}

// 用户模块
func UserRoute() {
	// 获取用户信息 curl http://127.0.0.1:8080/user/detail?id=1
	router.GET("user/detail", user.GetUserHandler)
}

// 帮助/测试路由
func HelpRoute() {
	// 测试连通性
	router.GET("/help/ping", help.PingHandler)
}
